package engine

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"os"
	"strings"
	"time"
)

// Will cause all system goroutines to stop when true
var shutdownSignal = false

// Structure representing the core server-side game engine
type HeliaEngine struct {
	Universe *universe.Universe
}

// Initializes a new instance of the game engine
func (e *HeliaEngine) Initialize() *HeliaEngine {
	// setup schematic runs watcher
	shared.TeeLog("Starting schematic runs watcher...")

	startSchematics()

	if !schematicRunnerStarted {
		panic("Schematics watcher not running!")
	}

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

	// instantiate engine
	shared.TeeLog("Universe ready!")
	engine := HeliaEngine{}

	return &engine
}

// Start the goroutines for each system
func (e *HeliaEngine) Start() {
	shared.TeeLog("Starting system goroutines...")

	for _, r := range e.Universe.Regions {
		for _, s := range r.Systems {
			go func(sol *universe.SolarSystem) {
				var tpf int = 0
				lastFrame := makeTimestamp()

				// game loop
				for {
					// check for shutdown signal
					if shutdownSignal {
						sol.HandleShutdownSignal()
						break
					}

					// for monitoring spawn exit
					c := make(chan struct{}, 1)

					// spawn goroutine so defer works as expected
					go func(sol *universe.SolarSystem) {
						// update system
						sol.PeriodicUpdate()

						// handle escalations from system
						handleEscalations(sol)

						// get time of last frame
						now := makeTimestamp()
						tpf = int(now - lastFrame)

						// find remaining portion of server heatbeat
						if tpf < universe.Heartbeat {
							// sleep for remainder of server heartbeat
							time.Sleep(time.Duration(universe.Heartbeat-tpf) * time.Millisecond)
						}

						// done!
						c <- struct{}{}
					}(sol)

					// wait for goroutine to return
					<-c

					// update last frame time
					lastFrame = makeTimestamp()
				}

				shared.TeeLog(fmt.Sprintf("System %s has halted.", sol.SystemName))
			}(s)
		}
	}

	shared.TeeLog("System goroutines started!")

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

					// test locks
					s.TestLocks()
					shared.TeeLog(fmt.Sprintf("* [%v] Passed", s.SystemName))

					// small sleep between systems
					time.Sleep(5 * time.Millisecond)
				}

				// larger sleep between regions
				time.Sleep(20 * time.Millisecond)
			}

			shared.TeeLog("* All systems passed deadlock check!")

			// wait 5 minutes
			time.Sleep(time.Second * 300)
		}
	}(e)

	go func() {
		for {
			// automatic shutdown in event of deadlock
			if shared.MutexFreeze {
				e.Shutdown()
			}

			// avoid pegging CPU
			time.Sleep(1 * time.Second)
		}
	}()
}

// Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	// guard against duplicate requests
	if shutdownSignal {
		return
	}

	shared.TeeLog("! Server shutdown initiated")

	// shut down simulation
	shared.TeeLog("Stopping simulation...")
	shutdownSignal = true
	shared.ShutdownSignal = true

	// wait for 30 seconds to give everything a chance to exit
	time.Sleep(30 * time.Second)

	shared.TeeLog("Halt success assumed")

	// save progress
	shared.TeeLog("Saving world state...")
	saveUniverse(e.Universe)
	shared.TeeLog("World state saved!")

	// end program
	shared.TeeLog("Shutdown complete! Goodbye :)")

	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func handleEscalations(sol *universe.SolarSystem) {
	// get services
	startSvc := sql.GetStartService()
	shipSvc := sql.GetShipService()
	userSvc := sql.GetUserService()
	itemSvc := sql.GetItemService()
	schematicRunSvc := sql.GetSchematicRunService()

	// obtain lock
	sol.Lock.Lock("core.handleEscalations")
	defer sol.Lock.Unlock()

	// iterate over moved items
	for id := range sol.MovedItems {
		// capture reference and remove from map
		mi := sol.MovedItems[id]
		delete(sol.MovedItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock("core.handleEscalations::MovedItems")
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

	// iterate over packaged items
	for id := range sol.PackagedItems {
		// capture reference and remove from map
		mi := sol.PackagedItems[id]
		delete(sol.PackagedItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock("core.handleEscalations::PackagedItems")
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

	// iterate over unpackaged items
	for id := range sol.UnpackagedItems {
		// capture reference and remove from map
		mi := sol.UnpackagedItems[id]
		delete(sol.UnpackagedItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock("core.handleEscalations::UnpackagedItems")
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

	// iterate over changed quantity items
	for id := range sol.ChangedQuantityItems {
		// capture reference and remove from map
		mi := sol.ChangedQuantityItems[id]
		delete(sol.ChangedQuantityItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock("core.handleEscalations::ChangedQuantityItems")
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

	// iterate over new items
	for id := range sol.NewItems {
		// capture reference and remove from map
		mi := sol.NewItems[id]
		delete(sol.NewItems, id)

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// lock item
			mi.Lock.Lock("core.handleEscalations::NewItems")
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

	// iterate over new sell orders
	for id := range sol.NewSellOrders {
		// capture reference
		mi := sol.NewSellOrders[id]

		// make sure we have waited long enough
		mi.CoreWait--

		if mi.CoreWait > -1 {
			continue
		}

		// remove from map
		delete(sol.NewSellOrders, id)

		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// lock sell order
			mi.Lock.Lock("core.handleEscalations::NewSellOrders")
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

	// iterate over bought sell orders
	for id := range sol.BoughtSellOrders {
		// capture reference and remove from map
		mi := sol.BoughtSellOrders[id]
		delete(sol.BoughtSellOrders, id)

		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// lock sell order
			mi.Lock.Lock("core.handleEscalations::BoughtSellOrders")
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

	// iterate over dead ships
	for id := range sol.DeadShips {
		// capture reference and remove from map
		ds := sol.DeadShips[id]
		delete(sol.DeadShips, id)

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
			for _, v := range ds.Aggressors {
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
		}(ds, sol)
	}

	// iterate over NPCs in need of respawn
	for id := range sol.NPCNeedRespawn {
		// capture reference and remove from map
		rs := sol.NPCNeedRespawn[id]
		delete(sol.NPCNeedRespawn, id)

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

	// iterate over clients in need of respawn
	for id := range sol.PlayerNeedRespawn {
		// capture reference and remove from map
		rs := sol.PlayerNeedRespawn[id]
		delete(sol.PlayerNeedRespawn, id)

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

			// find their home station
			home := sol.Universe.FindStation(start.HomeStationID, &sol.ID)

			if home == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no home station!", rs.UID))
				return
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

	// iterate over new ship tickets
	for id := range sol.NewShipTickets {
		// capture reference and remove from map
		rs := sol.NewShipTickets[id]
		delete(sol.NewShipTickets, id)

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

	// iterate over ship switches
	for id := range sol.ShipSwitches {
		// capture reference and remove from map
		rs := sol.ShipSwitches[id]
		delete(sol.ShipSwitches, id)

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

	// iterate over ship no load flag sets
	for id := range sol.SetNoLoad {
		// capture reference and remove from map
		rs := sol.SetNoLoad[id]
		delete(sol.SetNoLoad, id)

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

	// iterate over used ship purchases
	for id := range sol.UsedShipPurchases {
		// capture reference and remove from map
		rs := sol.UsedShipPurchases[id]
		delete(sol.UsedShipPurchases, id)

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

	// iterate over ship renames
	for id := range sol.ShipRenames {
		// capture reference and remove from map
		rs := sol.ShipRenames[id]
		delete(sol.ShipRenames, id)

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
				Runs: make([]models.ServerSchematicRunEntryBody, len(runs)),
			}

			for _, e := range runs {
				// obtain lock
				e.Lock.Lock("core.handleEscalations::SchematicRunViews::runs")
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
			msgRegistry := models.NewMessageRegistry()

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
			rs.Client.Lock.Lock("core.handleEscalations::NewSchematicRuns")
			defer rs.Client.Lock.Unlock()

			rs.Ship.Lock.Lock("core.handleEscalations::NewSchematicRuns")
			defer rs.Ship.Lock.Unlock()

			rs.SchematicItem.Lock.Lock("core.handleEscalations::NewSchematicRuns")
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
