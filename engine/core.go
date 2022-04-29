package engine

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// shared services
var startSvc = sql.GetStartService()
var shipSvc = sql.GetShipService()
var userSvc = sql.GetUserService()
var itemSvc = sql.GetItemService()
var schematicRunSvc = sql.GetSchematicRunService()
var factionSvc = sql.GetFactionService()
var actionReportSvc = sql.GetActionReportService()

// Will cause all region goroutines to stop when true
var shutdownSignal = false

// Structure representing the core server-side game engine
type HeliaEngine struct {
	Universe *universe.Universe
}

// Initializes a new instance of the game engine
func (e *HeliaEngine) Initialize() *HeliaEngine {
	// load universe
	shared.TeeLog("Loading game universe from database...")

	u, err := LoadUniverse()

	if err != nil {
		panic(fmt.Sprintf("Unable to load game universe: %v", err))
	}

	// generate transient celestials
	shared.TeeLog("Generating transient celestials...")
	u.BuildTransientCelestials()

	// cache starmap
	shared.TeeLog("Building starmap...")
	err = u.BuildMapWithCache()

	if err != nil {
		panic(fmt.Sprintf("Unable to build starmap: %v", err))
	}

	e.Universe = u

	// setup schematic runs watcher
	shared.TeeLog("Starting schematic runs watcher...")

	startSchematics()

	if !schematicRunnerStarted {
		panic("Schematics watcher not running!")
	}

	// instantiate engine
	shared.TeeLog("Universe ready!")
	engine := HeliaEngine{}

	// return pointer to engine
	return &engine
}

// Start a goroutine for each region
func (e *HeliaEngine) Start() {
	// log progress
	shared.TeeLog("Starting region goroutines...")

	// iterate over each region
	for _, r := range e.Universe.Regions {
		go func(r *universe.Region) {
			// set up time keeping variables for region
			var tpf int = 0
			lastFrame := makeTimestamp()

			// game loop
			for {
				// check for shutdown signal
				if shutdownSignal {
					// pass to systems
					defer func() {
						for _, s := range r.Systems {
							s.HandleShutdownSignal()
						}
					}()

					// exit routine
					break
				}

				// sleep for residual tick duration
				time.Sleep(time.Duration(universe.Heartbeat-tpf) * time.Millisecond)

				// process each system in the region
				for _, s := range r.Systems {
					// call periodic update with a wrapper function so defer works properly
					wrapSystemPeriodicUpdate(s, e)
				}

				// calculate residual
				now := makeTimestamp()
				tpf = int(now - lastFrame)

				// update last frame time
				lastFrame = now

				// guarantee routine yields
				time.Sleep(250 * time.Microsecond)
			}

			// detect routine exit
			shared.TeeLog(fmt.Sprintf("Region %s has halted.", r.RegionName))
		}(r)
	}

	// log progress
	shared.TeeLog("Region goroutines started!")

	// start watchdog goroutine (to alert of any deadlocks for debugging purposes)
	go func(e *HeliaEngine) {
		for {
			if shutdownSignal {
				shared.TeeLog("! Deadlock Check [goroutine] Halted!")
				break
			}

			shared.TeeLog("* Deadlock Check Starting")

			// iterate over systems
			for _, r := range e.Universe.Regions {
				for _, s := range r.Systems {
					shared.TeeLog(fmt.Sprintf("* Testing [%v]", s.SystemName))

					wgi := 500 // 5 second limit
					done := false

					go func(s *universe.SolarSystem) {
						// test locks
						s.TestLocks()
						shared.TeeLog(fmt.Sprintf("* [%v] Passed", s.SystemName))

						// small sleep between systems
						time.Sleep(50 * time.Millisecond)

						// mark done
						done = true
					}(s)

					// wait for exit
					for {
						// short sleep
						time.Sleep(10 * time.Millisecond)

						// check for exit
						if done {
							break
						}

						// check for too much time passing
						if wgi <= 0 {
							shared.TeeLog("! Deadlock check failed - initiating shutdown")
							shutdownSignal = true

							break
						}

						// decrement counter
						wgi--
					}
				}

				// larger sleep between regions
				time.Sleep(500 * time.Millisecond)
			}

			shared.TeeLog("* All systems passed deadlock check!")

			// wait 30 minutes
			time.Sleep(time.Minute * 30)
		}
	}(e)
}

// Wrapper so defer works as expected when updating a solar system
func wrapSystemPeriodicUpdate(sol *universe.SolarSystem, e *HeliaEngine) {
	// handle panics caused by solar system
	defer func(sol *universe.SolarSystem) {
		if r := recover(); r != nil {
			// log error for inspection
			shared.TeeLog(fmt.Sprintf("solar system %v panicked: %v", sol.SystemName, r))

			// include stack trace
			shared.TeeLog(fmt.Sprintf("stacktrace from panic: \n" + string(debug.Stack())))

			// emergency shutdown
			shared.TeeLog("! Emergency shutdown initiated due to solar system panic!")
			e.Shutdown()
		}
	}(sol)

	// update system
	sol.PeriodicUpdate()

	// handle escalations from system
	handleEscalations(sol)
}

// Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	// guard against duplicate requests
	if shutdownSignal {
		return
	}

	// get blob storage service
	bss, err := shared.LoadBlobStorageConfiguration()
	bssReady := true

	if err != nil {
		shared.TeeLog(fmt.Sprintf("Unable to load blob storage configuration: %v", err))
		bssReady = false
	}

	// log progress
	shared.TeeLog("! Server shutdown initiated")

	// shut down simulation
	shared.TeeLog("Stopping simulation...")
	shutdownSignal = true

	// wait for 30 seconds to give everything a chance to exit
	time.Sleep(30 * time.Second)

	shared.TeeLog("Halt success assumed")

	// save progress
	shared.TeeLog("Saving world state...")
	saveUniverse(e.Universe)
	shared.TeeLog("World state saved!")

	// stop profiling now
	pprof.StopCPUProfile()

	if shared.CpuProfileFile != nil {
		// close to flush
		shared.CpuProfileFile.Close()
	}

	// upload executable and cpu profile
	cpuProf, f := shared.ReadFileBytes("cpu.prof")
	heliaEx, g := shared.ReadFileBytes("main")

	if f && g && bssReady {
		// make timestamp
		ts := makeTimestamp()

		// profile
		shared.TeeLog("Uploading cpu profile...")
		v, err := bss.UploadBytesToBlob(*cpuProf, fmt.Sprintf("%v-cpu.prof", ts), "application/octet-stream")

		if err != nil {
			shared.TeeLog(fmt.Sprintf("Error uploading cpu profile: %v", err))
		} else {
			shared.TeeLog(fmt.Sprintf("Profile uploaded: %v", v))
		}

		// executable
		shared.TeeLog("Uploading executable...")
		v, err = bss.UploadBytesToBlob(*heliaEx, fmt.Sprintf("%v-helia", ts), "application/octet-stream")

		if err != nil {
			shared.TeeLog(fmt.Sprintf("Error uploading executable: %v", err))
		} else {
			shared.TeeLog(fmt.Sprintf("Executable uploaded: %v", v))
		}
	}

	// end program
	shared.TeeLog("Shutdown complete! Goodbye :)")

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}

func handleEscalations(sol *universe.SolarSystem) {
	// obtain lock
	sol.Lock.Lock()
	defer sol.Lock.Unlock()

	// iterate over moved items
	for _, mi := range sol.MovedItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save new location of item to db
			err := saveItemLocation(mi.ID, mi.ContainerID)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to relocate item %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear moved items
	sol.MovedItems = make([]*universe.Item, 0)

	// iterate over packaged items
	for _, mi := range sol.PackagedItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// mark item as packaged in the db
			err := packageItem(mi.ID)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to package item %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear packaged items
	sol.PackagedItems = make([]*universe.Item, 0)

	// iterate over unpackaged items
	for _, mi := range sol.UnpackagedItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save unpackaged item to db
			err := unpackageItem(mi.ID, mi.Meta)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to unpackage item %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear unpackaged items
	sol.UnpackagedItems = make([]*universe.Item, 0)

	// iterate over changed quantity items
	for _, mi := range sol.ChangedQuantityItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save quantity of item to db
			err := changeQuantity(mi.ID, mi.Quantity)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to change quantity of item %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear changed quantity items
	sol.ChangedQuantityItems = make([]*universe.Item, 0)

	// iterate over changed meta items
	for id := range sol.ChangedMetaItems {
		// capture reference and remove from map
		mi := sol.ChangedMetaItems[id]
		delete(sol.ChangedMetaItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save metadata of item to db
			err := changeMeta(mi.ID, mi.Meta)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to change metadata of item %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// iterate over new items
	for _, mi := range sol.NewItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// save new item to db
			id, err := newItem(mi)

			// error check
			if err != nil || id == nil {
				shared.TeeLog(fmt.Sprintf("Unable to save new item %v: %v", mi.ID, err))
			} else {
				// store corrected id from db insert
				mi.ID = *id

				// load new item
				ni, err := itemSvc.GetItemByID(*id)

				if err != nil || id == nil {
					shared.TeeLog(fmt.Sprintf("Unable to load new item %v: %v", mi.ID, err))
				} else {
					fi, err := LoadItem(ni)

					if err != nil || id == nil {
						shared.TeeLog(fmt.Sprintf("Unable to integrate new item %v: %v", mi.ID, err))
					} else {
						// copy loaded values
						mi.Meta = fi.Meta
						mi.ItemTypeMeta = fi.ItemTypeMeta
						mi.ItemTypeName = fi.ItemTypeName
						mi.ItemFamilyID = fi.ItemFamilyID
						mi.ItemFamilyName = fi.ItemFamilyName
						mi.Process = fi.Process

						// mark as clean
						mi.CoreDirty = false
					}
				}
			}
		}(mi, sol)
	}

	// clear new items
	sol.NewItems = make([]*universe.Item, 0)

	// iterate over new sell orders
	for _, mi := range sol.NewSellOrders {
		// make sure we have waited long enough
		mi.CoreWait--

		if mi.CoreWait > -1 {
			continue
		}

		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// lock sell order
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// get item id from item if linked and flag set
			if mi.Item != nil && mi.GetItemIDFromItem {
				mi.ItemID = mi.Item.ID
			}

			// save new sell order to db
			id, err := newSellOrder(mi)

			// error check
			if err != nil || id == nil {
				shared.TeeLog(fmt.Sprintf("Unable to save new sell order %v: %v", mi.ID, err))
			} else {
				// store corrected id from db insert
				mi.ID = *id

				// add to station's open orders
				sol.StoreOpenSellOrder(mi, true)

				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear new sell orders
	sol.NewSellOrders = make([]*universe.SellOrder, 0)

	// iterate over bought sell orders
	for _, mi := range sol.BoughtSellOrders {
		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// lock sell order
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark sell order as bought in db
			err := markSellOrderAsBought(mi)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to update bought sell order %v: %v", mi.ID, err))
			} else {
				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// clear bought sell orders
	sol.BoughtSellOrders = make([]*universe.SellOrder, 0)

	// iterate over dead ships
	for _, ds := range sol.DeadShips {
		// handle escalation on another goroutine
		go func(ds *universe.Ship, sol *universe.SolarSystem) {
			// remove ship from system
			sol.RemoveShip(ds, true)

			// make sure everything is set
			ds.Destroyed = true

			if ds.DestroyedAt == nil {
				now := time.Now()
				ds.DestroyedAt = &now
			}

			// apply standings change for kill
			for _, v := range ds.PlayerAggressors {
				if v != nil {
					// combat adjustment
					if ds.Faction.IsNPC {
						v.AdjustStandingNPC(ds.FactionID, ds.Faction.ReputationSheet, -1, true)
					} else {
						v.AdjustStandingPlayer(ds.ReputationSheet, -0.25, true)
					}

					// bound check
					v.EnforceBounds(true)
				}
			}

			// update dead ship in db
			err := saveShip(ds)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to mark ship %v as dead in db (%v)!", ds.ID, err))
				return
			}

			// generate action report
			kl := generateKillLog(ds)

			if kl == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to generate kill report from ship %v!", ds.ID))
				return
			}

			// get involved party user ids
			ips := make(map[string]interface{})

			for i, p := range kl.InvolvedParties {
				ips[fmt.Sprint(i+1)] = p.UserID.String()
			}

			ips[fmt.Sprint(0)] = ds.UserID.String()

			// save action report
			_, err = actionReportSvc.NewActionReport(
				sql.ActionReport{
					Report:          *kl,
					Timestamp:       *ds.DestroyedAt,
					InvolvedParties: ips,
				},
			)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to save action log for ship %v - (%v)!", ds.ID, err))
				return
			}

			// make wreck available
			ds.WreckReady = true
		}(ds, sol)
	}

	// clear dead ships
	sol.DeadShips = make([]*universe.Ship, 0)

	// iterate over NPCs in need of respawn
	for _, rs := range sol.NPCNeedRespawn {
		// handle escalation on another goroutine
		go func(rs *universe.Ship, sol *universe.SolarSystem) {
			// find their starting conditions
			user, err := userSvc.GetUserByID(rs.UserID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - no associated user!", rs.UserID))
				return
			}

			start, err := startSvc.GetStartByID(user.StartID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - no associated start!", rs.UserID))
				return
			}

			// find their home station
			home := sol.Universe.FindStation(start.HomeStationID, &sol.ID)

			if home == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - no home station!", rs.UserID))
				return
			}

			// create their ship docked in that station
			u, err := CreateNoobShipForPlayer(start, rs.UserID)

			if err != nil || u.CurrentShipID == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - failed to create noob ship (%v | %v)!", rs.UserID, err, u.CurrentShipID))
				return
			}

			ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - no noob ship!", rs.UserID))
				return
			}

			// save ship
			err = shipSvc.UpdateShip(*ns)

			if home == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - couldn't save noob ship changes (%v)!", rs.UserID, err))
				return
			}

			// load the ship into the home station's system
			ns, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - no noob ship again!", rs.UserID))
				return
			}

			es, err := LoadShip(ns, sol.Universe)

			if err != nil || es == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn NPC %v - couldn't load new noobship into universe (%v)!", rs.UserID, err))
				return
			}

			// set up as NPC
			es.BehaviourMode = user.BehaviourMode
			es.IsNPC = true
			es.BeingFlownByPlayer = false

			// link NPC's faction into ship
			es.FactionID = u.CurrentFactionID
			es.Faction = sol.Universe.Factions[u.CurrentFactionID.String()]

			// link NPC's reputation sheet into ship (not that this will be used by anything...)
			es.ReputationSheet = LoadReputationSheet(user)

			// put ship in home system
			home.CurrentSystem.AddShip(es, true)

			// log respawn to console
			shared.TeeLog(
				fmt.Sprintf(
					"[%v] %v was respawned at %v (NPC).",
					home.CurrentSystem.SystemName,
					es.CharacterName,
					home.StationName,
				),
			)
		}(rs, sol)
	}

	// clear npcs in need of respawn
	sol.NPCNeedRespawn = make([]*universe.Ship, 0)

	// iterate over clients in need of respawn
	for _, rs := range sol.PlayerNeedRespawn {
		// handle escalation on another goroutine
		go func(rs *shared.GameClient, sol *universe.SolarSystem) {
			// remove client from system
			sol.RemoveClient(rs, true)

			// find their starting conditions
			start, err := startSvc.GetStartByID(rs.StartID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no starting conditions!", rs.UID))
				return
			}

			// check if their home station is overriden by faction membership
			ur, err := userSvc.GetUserByID(*rs.UID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no user!", rs.UID))
				return
			}

			uf := sol.Universe.Factions[ur.CurrentFactionID.String()]

			if uf == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no faction!", rs.UID))
				return
			}

			if !uf.IsNPC && uf.HomeStationID != nil {
				// override home station on start in-memory only
				start.HomeStationID = *uf.HomeStationID
			}

			// find their home station
			home := sol.Universe.FindStation(start.HomeStationID, &sol.ID)

			if home == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no home station!", rs.UID))
				return
			}

			// make sure correct systemid is stored in-memory
			if !uf.IsNPC && uf.HomeStationID != nil {
				start.SystemID = home.SystemID
			}

			// create their noob ship docked in that station
			u, err := CreateNoobShipForPlayer(start, *rs.UID)

			if err != nil || u.CurrentShipID == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - failed to create noob ship (%v | %v)!", rs.UID, err, u.CurrentShipID))
				return
			}

			ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no noob ship!", rs.UID))
				return
			}

			// save noob ship
			err = shipSvc.UpdateShip(*ns)

			if home == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - couldn't save noob ship changes (%v)!", rs.UID, err))
				return
			}

			// load the noob ship into the home station's system
			ns, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no noob ship again!", rs.UID))
				return
			}

			es, err := LoadShip(ns, sol.Universe)

			if err != nil || es == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - couldn't load new noobship into universe (%v)!", rs.UID, err))
				return
			}

			// put ship in home system
			home.CurrentSystem.AddShip(es, true)

			// link player's faction into ship
			es.FactionID = u.CurrentFactionID

			// link player's reputation sheet into ship
			es.ReputationSheet = &rs.ReputationSheet

			// link player's experience sheet into ship
			es.ExperienceSheet = &rs.ExperienceSheet

			// set client current ship to new noob ship
			rs.CurrentShipID = es.ID
			es.BeingFlownByPlayer = true

			// put the client in that system
			home.CurrentSystem.AddClient(rs, true)

			// log respawn to console
			shared.TeeLog(
				fmt.Sprintf(
					"[%v] %v was respawned at %v.",
					home.CurrentSystem.SystemName,
					es.CharacterName,
					home.StationName,
				),
			)
		}(rs, sol)
	}

	// clear clients in need of respawn
	sol.PlayerNeedRespawn = make([]*shared.GameClient, 0)

	// iterate over new ship tickets
	for _, rs := range sol.NewShipTickets {
		// handle escalation on another goroutine
		go func(rs *universe.NewShipTicket, sol *universe.SolarSystem) {
			// create new ship for player
			ps, err := CreateShipForPlayer(
				rs.UserID,
				rs.ShipTemplateID,
				rs.StationID,
				sol.ID,
			)

			if err != nil || ps == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to complete ship purchase for %v - failure saving (%v)!", rs.UserID, err))
				return
			}

			// load into universe
			es, err := LoadShip(ps, sol.Universe)

			if err != nil || es == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to complete ship purchase for %v - failure loading (%v)!", rs.UserID, err))
				return
			}

			// put ship in system
			sol.AddShip(es, true)

			// is this a station workshop/warehouse?
			if !es.TemplateData.CanUndock {
				if rs.Client != nil {
					// inform player they will be charged for storing items in the cargo bay
					notifyClientOfWorkshopFee(rs.Client)
				}
			}
		}(rs, sol)
	}

	// clear new ship tickets
	sol.NewShipTickets = make([]*universe.NewShipTicket, 0)

	// iterate over ship switches
	for _, rs := range sol.ShipSwitches {
		// handle escalation on another goroutine
		go func(rs *universe.ShipSwitch, sol *universe.SolarSystem) {
			// update currently flown ship in database
			err := userSvc.SetCurrentShipID(*rs.Client.UID, &rs.Target.ID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to complete ship switch for %v - failure saving (%v)!", rs.Client.UID, err))
				return
			}

			tgt := rs.Target
			src := rs.Source

			// link player's faction into target ship
			tgt.FactionID = src.FactionID

			// link player's reputation sheet into target ship
			tgt.ReputationSheet = &rs.Client.ReputationSheet

			// link player's experience sheet into target ship
			tgt.ExperienceSheet = &rs.Client.ExperienceSheet

			// set client current ship to target ship
			rs.Client.CurrentShipID = tgt.ID
			tgt.BeingFlownByPlayer = true

			// unmark source ship as being flown by player
			src.BeingFlownByPlayer = false
		}(rs, sol)
	}

	// clear ship switches
	sol.ShipSwitches = make([]*universe.ShipSwitch, 0)

	// iterate over ship no load flag sets
	for _, rs := range sol.SetNoLoad {
		// handle escalation on another goroutine
		go func(rs *universe.ShipNoLoadSet, sol *universe.SolarSystem) {
			// update flag in database
			err := shipSvc.SetNoLoad(rs.ID, rs.Flag)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to set no load switch for %v - failure saving (%v)!", rs.ID, err))
				return
			}
		}(rs, sol)
	}

	// clear set no loads
	sol.SetNoLoad = make([]*universe.ShipNoLoadSet, 0)

	// iterate over used ship purchases
	for _, rs := range sol.UsedShipPurchases {
		// handle escalation on another goroutine
		go func(rs *universe.UsedShipPurchase, sol *universe.SolarSystem) {
			// update owner in database
			err := shipSvc.SetOwner(rs.ShipID, rs.UserID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to set new owner for %v - failure saving (%v)!", rs.ShipID, err))
				return

			}

			// update noload flag in database
			err = shipSvc.SetNoLoad(rs.ShipID, false)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to set no load switch for %v - failure saving (%v)!", rs.ShipID, err))
				return

			}

			// load ship
			shd, err := shipSvc.GetShipByID(rs.ShipID, false)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to load used ship for %v - failure saving (%v)! [part 1]", rs.ShipID, err))
				return
			}

			sh, err := LoadShip(shd, sol.Universe)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to load used ship for %v - failure saving (%v)! [part 2]", rs.ShipID, err))
				return
			}

			// inject into system
			sol.AddShip(sh, true)

			// is this a station workshop/warehouse?
			if !sh.TemplateData.CanUndock {
				if rs.Client != nil {
					// inform player they will be charged for storing items in the cargo bay
					notifyClientOfWorkshopFee(rs.Client)
				}
			}
		}(rs, sol)
	}

	// clear used ship purchases
	sol.UsedShipPurchases = make([]*universe.UsedShipPurchase, 0)

	// iterate over ship renames
	for _, rs := range sol.ShipRenames {
		// handle escalation on another goroutine
		go func(rs *universe.ShipRename, sol *universe.SolarSystem) {
			// update name in database
			err := shipSvc.Rename(rs.ShipID, rs.Name)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to rename ship %v - failure saving (%v)!", rs.ShipID.ID(), err))
				return
			}
		}(rs, sol)
	}

	// clear ship renames
	sol.ShipRenames = make([]*universe.ShipRename, 0)

	// iterate over clients in need of a schematic runs update
	for id := range sol.SchematicRunViews {
		// capture reference and remove from map
		rs := sol.SchematicRunViews[id]
		delete(sol.SchematicRunViews, id)

		// handle escalation on another goroutine
		go func(rs *shared.GameClient) {
			// get runs
			runs := getSchematicRunsByUser(*rs.UID)

			// copy to update message
			msg := models.ServerSchematicRunsUpdateBody{
				Runs: make([]models.ServerSchematicRunEntryBody, 0),
			}

			for _, e := range runs {
				// obtain lock
				e.Lock.Lock()
				defer e.Lock.Unlock()

				// skip if uninitialized
				if !e.Initialized {
					continue
				}

				// null checks
				if e.Process == nil {
					continue
				}

				if e.Ship == nil {
					continue
				}

				if e.Ship.DockedAtStation == nil {
					continue
				}

				if e.Ship.CurrentSystem == nil {
					continue
				}

				// skip if dirty
				if e.SchematicItem.CoreDirty {
					continue
				}

				// copy fields
				o := models.ServerSchematicRunEntryBody{
					SchematicRunID:     e.ID,
					SchematicName:      e.Process.Name,
					HostShipName:       e.Ship.ShipName,
					HostStationName:    e.Ship.DockedAtStation.StationName,
					SolarSystemName:    e.Ship.CurrentSystem.SystemName,
					StatusID:           e.StatusID,
					PercentageComplete: (float64(e.Progress) / float64(e.Process.Time)) + universe.Epsilon,
				}

				// store in message
				msg.Runs = append(msg.Runs, o)
			}

			// get message registry
			msgRegistry := models.SharedMessageRegistry

			// serialize message
			b, _ := json.Marshal(&msg)

			sct := models.GameMessage{
				MessageType: msgRegistry.SchematicRunsUpdate,
				MessageBody: string(b),
			}

			// write message to client
			rs.WriteMessage(&sct)
		}(rs)
	}

	// iterate over newly invoked schematics
	for id := range sol.NewSchematicRuns {
		// capture reference and remove from map
		rs := sol.NewSchematicRuns[id]
		delete(sol.NewSchematicRuns, id)

		// handle escalation on another goroutine
		go func(rs *universe.NewSchematicRunTicket, sol *universe.SolarSystem) {
			// obtain locks
			rs.Client.Lock.Lock()
			defer rs.Client.Lock.Unlock()

			rs.Ship.Lock.Lock()
			defer rs.Ship.Lock.Unlock()

			rs.SchematicItem.Lock.Lock()
			defer rs.SchematicItem.Lock.Unlock()

			// create new run in database
			r, err := schematicRunSvc.NewSchematicRun(sql.SchematicRun{
				ProcessID:       rs.SchematicItem.Process.ID,
				StatusID:        "new",
				SchematicItemID: rs.SchematicItem.ID,
				UserID:          *rs.Client.UID,
			})

			if err != nil || r == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to run schematic %v - failure saving (%v)!", rs.SchematicItem.ID, err))
				return
			}

			// convert to universe model
			usr := universe.SchematicRun{
				ID:              r.ID,
				Created:         r.Created,
				ProcessID:       r.ProcessID,
				StatusID:        r.StatusID,
				Progress:        r.Progress,
				SchematicItemID: r.SchematicItemID,
				UserID:          r.UserID,
				Lock:            sync.Mutex{},
			}

			// hook schematic
			addSchematicRunForUser(*rs.Client.UID, &usr)
			hookSchematics(rs.Ship)

			// mark as running
			rs.SchematicItem.SchematicInUse = true

			// unmark dirty
			rs.SchematicItem.CoreDirty = false
		}(rs, sol)
	}

	// iterate over new faction requests
	for id := range sol.NewFactions {
		// capture reference and remove from map
		mi := sol.NewFactions[id]
		delete(sol.NewFactions, id)

		// handle escalation on another goroutine
		go func(mi *universe.NewFactionTicket, sol *universe.SolarSystem) {
			// obtain lock on game client
			mi.Client.Lock.Lock()
			defer mi.Client.Lock.Unlock()

			// obtain lock on ship attached to client
			sh := sol.Universe.FindShip(mi.Client.CurrentShipID, nil)

			// null check
			if sh == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to create faction %v - no ship!", mi))
				return
			}

			sh.Lock.Lock()
			defer sh.Lock.Unlock()

			// one last docked check
			if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to create faction %v - creator not docked!", mi))
				return
			}

			// try to create faction (most likely to fail if name or ticker are taken)
			f, err := factionSvc.NewFaction(sql.Faction{
				Name:            mi.Name,
				Description:     mi.Description,
				IsNPC:           false,
				IsJoinable:      true,
				CanHoldSov:      false,
				Meta:            make(sql.Meta),
				ReputationSheet: sql.FactionReputationSheet{},
				Ticker:          mi.Ticker,
				OwnerID:         mi.Client.UID,
				HomeStationID:   &mi.HomeStationID,
			})

			// error check
			if err != nil {
				// log
				shared.TeeLog(fmt.Sprintf("Unable to create faction %v: %v", mi, err))

				// notify client of failure
				go func(c *shared.GameClient) {
					c.WriteErrorMessage("unable to create your faction - are your name and ticker unique?")
				}(mi.Client)

				// exit
				return
			}

			// copy founder's reputation sheet to new faction
			playerRep := SQLFromPlayerReputationSheet(&mi.Client.ReputationSheet)

			f.ReputationSheet = sql.FactionReputationSheet{
				Entries:        make(map[string]sql.ReputationSheetEntry),
				HostFactionIDs: make([]uuid.UUID, 0),
				WorldPercent:   0,
			}

			for _, l := range playerRep.FactionEntries {
				f.ReputationSheet.Entries[l.FactionID.String()] = sql.ReputationSheetEntry{
					SourceFactionID:  f.ID,
					TargetFactionID:  l.FactionID,
					StandingValue:    l.StandingValue,
					AreOpenlyHostile: l.AreOpenlyHostile,
				}
			}

			// save reputation sheet
			err = factionSvc.SaveFaction(*f)

			// error check
			if err != nil {
				// log
				shared.TeeLog(fmt.Sprintf("Unable to copy reputation sheet to new faction %v: %v", mi, err))

				// notify client of failure
				go func(c *shared.GameClient) {
					c.WriteErrorMessage("unable to copy reputation sheet!")
				}(mi.Client)

				// exit
				return
			}

			// load faction into universe
			uf := FactionFromSQL(f)
			sol.Universe.Factions[uf.ID.String()] = uf

			// put founder in their new faction
			err = userSvc.SetCurrentFactionID(*mi.Client.UID, &f.ID)

			// error check
			if err != nil {
				// log
				shared.TeeLog(fmt.Sprintf("Unable to assign creator to new faction %v: %v", mi, err))

				// notify client of failure
				go func(c *shared.GameClient) {
					c.WriteErrorMessage("unable to join you to your new faction!")
				}(mi.Client)

				// exit
				return
			}

			// notify clients of new faction
			af := models.ServerFactionUpdateBody{
				Factions: make([]models.ServerFactionBody, 0),
			}

			// include relationship data
			rels := make([]models.ServerFactionRelationship, 0)

			for _, rel := range f.ReputationSheet.Entries {
				rels = append(rels, models.ServerFactionRelationship{
					FactionID:        rel.TargetFactionID,
					AreOpenlyHostile: rel.AreOpenlyHostile,
					StandingValue:    rel.StandingValue,
				})
			}

			ax := models.ServerFactionBody{
				ID:            f.ID,
				Name:          f.Name,
				Description:   f.Description,
				IsNPC:         f.IsNPC,
				IsJoinable:    f.IsJoinable,
				CanHoldSov:    f.CanHoldSov,
				Ticker:        f.Ticker,
				Relationships: rels,
			}

			if f.OwnerID != nil {
				ax.OwnerID = f.OwnerID.String()
			}

			af.Factions = append(af.Factions, ax)

			// package faction data
			b, _ := json.Marshal(&af)

			msg := models.GameMessage{
				MessageType: models.SharedMessageRegistry.FactionUpdate,
				MessageBody: string(b),
			}

			// write message to connected clients
			sol.Universe.SendGlobalMessage(&msg)

			// update ship with new faction
			sh.Faction = uf
			sh.FactionID = uf.ID

			// send welcome message to player
			go func(c *shared.GameClient) {
				c.WriteInfoMessage(fmt.Sprintf("welcome to %v !", uf.Name))
			}(mi.Client)
		}(mi, sol)
	}

	// iterate over leave faction requests
	for id := range sol.LeaveFactions {
		// capture reference and remove from map
		mi := sol.LeaveFactions[id]
		delete(sol.LeaveFactions, id)

		// handle escalation on another goroutine
		go func(mi *universe.LeaveFactionTicket, sol *universe.SolarSystem) {
			// obtain lock on game client
			mi.Client.Lock.Lock()
			defer mi.Client.Lock.Unlock()

			// obtain lock on ship attached to client
			sh := sol.Universe.FindShip(mi.Client.CurrentShipID, nil)

			// null check
			if sh == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to leave faction %v - no ship!", mi))
				return
			}

			sh.Lock.Lock()
			defer sh.Lock.Unlock()

			// one last docked check
			if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to leave faction %v - not docked!", mi))
				return
			}

			// find their starting conditions
			user, err := userSvc.GetUserByID(*mi.Client.UID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to leave faction %v - no associated user!", mi.Client.UID))
				return
			}

			start, err := startSvc.GetStartByID(user.StartID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to leave faction %v - no associated start!", mi.Client.UID))
				return
			}

			// put player in their starter faction
			fID := start.FactionID
			err = userSvc.SetCurrentFactionID(*mi.Client.UID, &fID)

			// error check
			if err != nil {
				// log
				shared.TeeLog(fmt.Sprintf("Unable to join creator to new faction %v: %v", mi, err))

				// notify client of failure
				go func(c *shared.GameClient) {
					c.WriteErrorMessage("unable to join you to your starter faction!")
				}(mi.Client)

				// exit
				return
			}

			// get faction from cache
			uf := sol.Universe.Factions[fID.String()]

			// update ship with new faction
			sh.Faction = uf
			sh.FactionID = uf.ID

			// send welcome message to player
			go func(c *shared.GameClient) {
				c.WriteInfoMessage(fmt.Sprintf("welcome to %v !", uf.Name))
			}(mi.Client)
		}(mi, sol)
	}

	// iterate over join faction requests
	for id := range sol.JoinFactions {
		// capture reference and remove from map
		mi := sol.JoinFactions[id]
		delete(sol.JoinFactions, id)

		// handle escalation on another goroutine
		go func(mi *universe.JoinFactionTicket, sol *universe.SolarSystem) {
			// null check
			if mi.OwnerClient == nil {
				shared.TeeLog(fmt.Sprintf("No owner client attached to join request: %v", mi))
				return
			}

			// try to find the target faction
			faction := sol.Universe.Factions[mi.FactionID.String()]

			if faction == nil {
				shared.TeeLog(fmt.Sprintf("Unable to find target faction to join: %v", mi))
				return
			}

			// try to find applicant's game client
			appClient := sol.Universe.FindCurrentPlayerClient(mi.UserID, nil)

			if appClient != nil {
				// obtain lock on applicant's game client
				appClient.Lock.Lock()
				defer appClient.Lock.Unlock()
			}

			// try to find applicant's ship
			appShip := sol.Universe.FindCurrentPlayerShip(mi.UserID, nil)

			// null check
			if appShip == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to join faction %v - no ship!", mi))
				return
			}

			// obtain lock on applicant's ship
			appShip.Lock.Lock()
			defer appShip.Lock.Unlock()

			// put player in the target faction
			fID := faction.ID
			err := userSvc.SetCurrentFactionID(mi.UserID, &fID)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to join player to target faction %v: %v", mi, err))
				return
			}

			// verify applicant is still in an NPC faction
			if appShip.Faction == nil || !appShip.Faction.IsNPC {
				// send failure message to owner
				go func(c *shared.GameClient) {
					if c != nil {
						c.WriteErrorMessage("unable to approve - applicant is no longer in an NPC faction")
					}
				}(mi.OwnerClient)

				// exit
				return
			}

			// update ship with new faction
			appShip.Faction = faction
			appShip.FactionID = faction.ID

			// send welcome message to applicant
			go func(c *shared.GameClient) {
				if c != nil {
					c.WriteInfoMessage(fmt.Sprintf("welcome to %v !", faction.Name))
				}
			}(appClient)

			// send confirmation message to owner
			go func(c *shared.GameClient) {
				if c != nil {
					c.WriteInfoMessage("approval success!")
				}
			}(mi.OwnerClient)
		}(mi, sol)
	}

	// iterate over clients in need of a faction member list
	for id := range sol.ViewMembers {
		// capture reference and remove from map
		rs := sol.ViewMembers[id]
		delete(sol.ViewMembers, id)

		// handle escalation on another goroutine
		go func(rs *universe.ViewMembersTicket) {
			// null check
			if rs.OwnerClient == nil {
				return
			}

			// get members
			members, err := userSvc.GetUsersByFactionID(rs.FactionID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to retrieve faction member users: %v, %v", rs, err))
				return
			}

			// copy to update message
			msg := models.ServerMembersUpdateBody{
				Members: make([]models.ServerMemberEntry, 0),
			}

			for _, e := range members {
				msg.Members = append(msg.Members, models.ServerMemberEntry{
					UserID:        e.ID,
					CharacterName: e.CharacterName,
				})
			}

			// get message registry
			msgRegistry := models.SharedMessageRegistry

			// serialize message
			b, _ := json.Marshal(&msg)

			sct := models.GameMessage{
				MessageType: msgRegistry.MembersUpdate,
				MessageBody: string(b),
			}

			// write message to client
			rs.OwnerClient.WriteMessage(&sct)
		}(rs)
	}

	// iterate over kick member requests
	for id := range sol.KickMembers {
		// capture reference and remove from map
		mi := sol.KickMembers[id]
		delete(sol.KickMembers, id)

		// handle escalation on another goroutine
		go func(mi *universe.KickMemberTicket, sol *universe.SolarSystem) {
			// null check
			if mi.OwnerClient == nil {
				shared.TeeLog(fmt.Sprintf("No owner client attached to kick request: %v", mi))
				return
			}

			// try to find the source faction
			kickingFaction := sol.Universe.Factions[mi.FactionID.String()]

			if kickingFaction == nil {
				shared.TeeLog(fmt.Sprintf("Unable to find kicking faction: %v", mi))
				return
			}

			// try to find the kickee's game client
			kickeeClient := sol.Universe.FindCurrentPlayerClient(mi.UserID, nil)

			if kickeeClient != nil {
				// obtain lock on kickee's game client
				kickeeClient.Lock.Lock()
				defer kickeeClient.Lock.Unlock()
			}

			// try to find kickee's ship
			kickeeShip := sol.Universe.FindCurrentPlayerShip(mi.UserID, nil)

			// null check
			if kickeeShip == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to kick member %v - no ship!", mi))
				return
			}

			// obtain lock on kickee's ship
			kickeeShip.Lock.Lock()
			defer kickeeShip.Lock.Unlock()

			// verify kickee is in the kicker's faction
			if kickeeShip.Faction == nil || kickeeShip.Faction.IsNPC || kickeeShip.FactionID != mi.FactionID {
				// send failure message to owner
				go func(c *shared.GameClient) {
					if c != nil {
						c.WriteErrorMessage("unable to kick - not a member of your faction")
					}
				}(mi.OwnerClient)

				// exit
				return
			}

			// docked check
			if kickeeShip.DockedAtStation == nil || kickeeShip.DockedAtStationID == nil {
				// send failure message to owner
				go func(c *shared.GameClient) {
					if c != nil {
						c.WriteErrorMessage("unable to kick - member is not currently docked")
					}
				}(mi.OwnerClient)

				// exit
				return
			}

			// find kickee's starting conditions
			kickee, err := userSvc.GetUserByID(mi.UserID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to kick member %v - no associated user!", mi.UserID))
				return
			}

			kickeeStart, err := startSvc.GetStartByID(kickee.StartID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to kick member %v - no associated start!", mi.UserID))
				return
			}

			// put kickee in their starter faction
			sfID := kickeeStart.FactionID
			err = userSvc.SetCurrentFactionID(mi.UserID, &sfID)

			// error check
			if err != nil {
				shared.TeeLog(fmt.Sprintf("Unable to kick member to starter faction %v: %v", mi, err))
				return
			}

			// get faction from cache
			uf := sol.Universe.Factions[sfID.String()]

			// update kickee's ship with new faction
			kickeeShip.Faction = uf
			kickeeShip.FactionID = uf.ID

			// send welcome message to kickee
			go func(c *shared.GameClient) {
				if c != nil {
					c.WriteInfoMessage(fmt.Sprintf("welcome to %v !", uf.Name))
				}
			}(kickeeClient)

			// send confirmation message to owner
			go func(c *shared.GameClient) {
				if c != nil {
					c.WriteInfoMessage("kick success!")
				}
			}(mi.OwnerClient)
		}(mi, sol)
	}
}

func notifyClientOfWorkshopFee(c *shared.GameClient) {
	if c == nil {
		return
	}

	lns := []string{
		"you have purchased a station workship! you will be charged a small fee ",
		"for every unit of volume you consume in its cargo bay. the more you store ",
		"in its cargo bay, the higher the fee will become. the fee is deducted from ",
		"the workshop's wallet - if you go into debt you will need to transfer CBN ",
		"to settle it before you can retrieve any stored items.",
	}

	infoMsg := strings.Join(lns, "")
	c.WriteInfoMessage(infoMsg)
}
