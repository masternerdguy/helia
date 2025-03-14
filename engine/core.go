package engine

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// shared services
var regionSvc = sql.GetRegionService()
var systemSvc = sql.GetSolarSystemService()
var shipSvc = sql.GetShipService()
var starSvc = sql.GetStarService()
var planetSvc = sql.GetPlanetService()
var stationSvc = sql.GetStationService()
var stationProcessSvc = sql.GetStationProcessService()
var processSvc = sql.GetProcessService()
var jumpholeSvc = sql.GetJumpholeService()
var asteroidSvc = sql.GetAsteroidService()
var itemTypeSvc = sql.GetItemTypeService()
var itemFamilySvc = sql.GetItemFamilyService()
var sellOrderSvc = sql.GetSellOrderService()
var itemSvc = sql.GetItemService()
var factionSvc = sql.GetFactionService()
var schematicRunSvc = sql.GetSchematicRunService()
var actionReportSvc = sql.GetActionReportService()
var userSvc = sql.GetUserService()
var startSvc = sql.GetStartService()
var processInputSvc = sql.GetProcessInputService()
var processOutputSvc = sql.GetProcessOutputService()
var shipTemplateSvc = sql.GetShipTemplateService()
var containerSvc = sql.GetContainerService()
var outpostSvc = sql.GetOutpostService()
var outpostTemplateSvc = sql.GetOutpostTemplateService()
var artifactSvc = sql.GetArtifactService()

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
	shared.TeeLog("Starting cluster goroutines...")

	// get core count
	cores := runtime.NumCPU()

	if cores <= 0 {
		cores = 1
	}

	// aggregate systems into a master slice
	e.Universe.AllSystems = make([]*universe.SolarSystem, 0)

	for _, r := range e.Universe.Regions {
		for _, s := range r.Systems {
			e.Universe.AllSystems = append(e.Universe.AllSystems, s)
		}
	}

	// create slices for cluster groups
	clusterGroups := make([][]*universe.SolarSystem, 0)

	for x := 0; x < cores; x++ {
		clusterGroups = append(clusterGroups, make([]*universe.SolarSystem, 0))
	}

	// assign systems to cluster groups (one per core)
	coreCounter := 0

	for _, s := range e.Universe.AllSystems {
		// bound core counter
		if coreCounter >= cores {
			coreCounter = 0
		}

		// assign system to cluster
		clusterGroups[coreCounter] = append(clusterGroups[coreCounter], s)

		// increment core counter
		coreCounter++
	}

	// start cluster group goroutines
	for x, cg := range clusterGroups {
		go func(x int, cg []*universe.SolarSystem) {
			// set up time keeping variables for cluster group
			var tpf int = 0
			lastFrame := makeTimestamp()

			// game loop
			for {
				// check for shutdown signal
				if shutdownSignal {
					// pass to systems
					defer func() {
						for _, s := range cg {
							s.HandleShutdownSignal()
						}
					}()

					// exit routine
					break
				}

				// sleep for residual tick duration
				time.Sleep(time.Duration(universe.Heartbeat-tpf) * time.Millisecond)

				// process each system in the cluster group
				for _, s := range cg {
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
			shared.TeeLog(fmt.Sprintf("Cluster group %d has halted.", x))
		}(x, cg)
	}

	// log progress
	shared.TeeLog("Cluster goroutines started!")

	// automatic scheduled restart goroutine
	go func(e *HeliaEngine) {
		needsReboot := false

		for {
			if shutdownSignal || needsReboot {
				shared.TeeLog("! Automatic restart check [goroutine] Halted!")
				break
			}

			// get UTC time
			utcNow := time.Now().UTC()

			if utcNow.Minute() == 0 {
				// check for afternoon reboot (noon EDT)
				if utcNow.Hour() == 16 {
					needsReboot = true
				}

				// handle reboot if needed
				if needsReboot {
					// log reboot scheduled
					shared.TeeLog("Automatic scheduled reboot will occur soon!")

					// message explaining situation to players
					sm := "Helia reboots once a day at noon EDT for scheduled maintainence. " +
						"A reboot will occur in ~10 minutes - please ensure you have gotten to a safe place " +
						"before then. The server is expected to take ~60 minutes to reboot."

					// send message to connected clients informing them of shutdown
					b, _ := json.Marshal(models.ServerPushInfoMessage{
						Message: sm,
					})

					msg := models.GameMessage{
						MessageType: models.SharedMessageRegistry.PushInfo,
						MessageBody: string(b),
					}

					go func(e *HeliaEngine) {
						e.Universe.SendGlobalMessage(&msg)
					}(e)

					// schedule reboot
					go func(e *HeliaEngine) {
						// wait 10 minutes
						time.Sleep(10 * time.Minute)

						// do reboot (containers are automatically restarted on Azure)
						shared.TeeLog("Automatic scheduled reboot is happening!")
						e.Shutdown()
					}(e)
				}
			}

			// don't peg cpu
			time.Sleep(250 * time.Millisecond)
		}
	}(e)

	shared.TeeLog("Automatic scheduled restart goroutine started!")
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
	handleEscalations(sol, e)
}

// Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	// guard against duplicate requests
	if shutdownSignal {
		return
	}

	// shut down simulation
	shared.TeeLog("Stopping simulation...")
	shutdownSignal = true

	// get blob storage service
	bss, err := shared.LoadBlobStorageConfiguration()
	bssReady := true

	if err != nil {
		shared.TeeLog(fmt.Sprintf("Unable to load blob storage configuration: %v", err))
		bssReady = false
	}

	// log progress
	shared.TeeLog("! Server shutdown initiated")

	// dump heap profile before saving
	if shared.HeapProfileFile != nil {
		// write data
		pprof.WriteHeapProfile(shared.HeapProfileFile)

		// flush to disk
		shared.HeapProfileFile.Close()
	}

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

	// load executable and cpu/heap profiles
	cpuProf, f := shared.ReadFileBytes("cpu.prof")
	heliaEx, g := shared.ReadFileBytes("main")
	heapProf, h := shared.ReadFileBytes("heap.prof")

	// make timestamp
	ts := makeTimestamp()
	exeUploaded := false

	// upload profiling files
	if f && g && bssReady {

		// profile
		shared.TeeLog("Uploading cpu profile...")
		v, err := bss.UploadBytesToBlob(*cpuProf, fmt.Sprintf("%v-cpu.prof", ts), "application/octet-stream")

		if err != nil {
			shared.TeeLog(fmt.Sprintf("Error uploading cpu profile: %v", err))
		} else {
			shared.TeeLog(fmt.Sprintf("Profile uploaded: %v", v))
		}

		if !exeUploaded {
			// executable
			shared.TeeLog("Uploading executable...")
			v, err = bss.UploadBytesToBlob(*heliaEx, fmt.Sprintf("%v-helia", ts), "application/octet-stream")

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Error uploading executable: %v", err))
			} else {
				shared.TeeLog(fmt.Sprintf("Executable uploaded: %v", v))
			}

			exeUploaded = true
		}
	}

	if h && g && bssReady {
		// profile
		shared.TeeLog("Uploading heap profile...")
		v, err := bss.UploadBytesToBlob(*heapProf, fmt.Sprintf("%v-heap.prof", ts), "application/octet-stream")

		if err != nil {
			shared.TeeLog(fmt.Sprintf("Error uploading heap profile: %v", err))
		} else {
			shared.TeeLog(fmt.Sprintf("Profile uploaded: %v", v))
		}

		if !exeUploaded {
			// executable
			shared.TeeLog("Uploading executable...")
			v, err = bss.UploadBytesToBlob(*heliaEx, fmt.Sprintf("%v-helia", ts), "application/octet-stream")

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Error uploading executable: %v", err))
			} else {
				shared.TeeLog(fmt.Sprintf("Executable uploaded: %v", v))
			}

			exeUploaded = true
		}
	}

	// end program
	shared.TeeLog("Shutdown complete! Goodbye :)")

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}

// Returns the shutdown signal flag
func (e *HeliaEngine) IsShuttingDown() bool {
	return shutdownSignal
}

// Helper function to handle panics caused by escalation goroutines
func escalationRecover(sol *universe.SolarSystem, e *HeliaEngine) {
	if r := recover(); r != nil {
		// log error for inspection
		shared.TeeLog(fmt.Sprintf("solar system [escalation] %v panicked: %v", sol.SystemName, r))

		// include stack trace
		shared.TeeLog(fmt.Sprintf("stacktrace from panic: \n" + string(debug.Stack())))

		// emergency shutdown
		shared.TeeLog("! Emergency shutdown initiated due to solar system [escalation] panic!")
		e.Shutdown()
	}
}

// Helper function to handle escalations from a solar system
func handleEscalations(sol *universe.SolarSystem, e *HeliaEngine) {
	// obtain lock
	sol.Lock.Lock()
	defer sol.Lock.Unlock()

	// iterate over moved items
	for _, mi := range sol.MovedItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
		// capture reference
		mi := sol.ChangedMetaItems[id]

		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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

	// clear changed meta items
	sol.ChangedMetaItems = make([]*universe.Item, 0)

	// iterate over new items
	for _, mi := range sol.NewItems {
		// handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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

				if err != nil || ni == nil {
					shared.TeeLog(fmt.Sprintf("Unable to load new item %v: %v", mi.ID, err))
				} else {
					fi, err := LoadItem(ni)

					if err != nil || fi == nil {
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

	// iterate over new items for devhax
	for _, mi := range sol.NewItemsDevHax {
		// handle escalation on another goroutine
		go func(mi *universe.NewItemFromNameTicketDevHax, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// find item type by name
			itemType, err := itemTypeSvc.GetItemTypeByName(mi.ItemTypeName)

			if err != nil || itemType == nil || itemType.Family == "ship" {
				shared.TeeLog(fmt.Sprintf("Unable to find non-ship item type [devhax] %v: %v", mi.ItemTypeName, err))
			} else {
				// target container
				tc := mi.Container

				// create item
				mi := &universe.Item{
					ItemTypeID:    itemType.ID,
					Meta:          universe.Meta(itemType.Meta),
					Created:       time.Now(),
					CreatedReason: "[devhax] :party parrot:",
					CreatedBy:     &mi.UserID,
					ContainerID:   mi.ContainerID,
					Quantity:      mi.Quantity,
					IsPackaged:    true,
					ItemTypeName:  mi.ItemTypeName,
				}

				// save new item to db
				id, err := newItem(mi)

				// error check
				if err != nil || id == nil {
					shared.TeeLog(fmt.Sprintf("Unable to save new item [devhax] %v: %v", mi.ID, err))
				} else {
					// store corrected id from db insert
					mi.ID = *id

					// load new item
					ni, err := itemSvc.GetItemByID(*id)

					if err != nil || ni == nil {
						shared.TeeLog(fmt.Sprintf("Unable to load new item [devhax] %v: %v", mi.ID, err))
					} else {
						fi, err := LoadItem(ni)

						if err != nil || fi == nil {
							shared.TeeLog(fmt.Sprintf("Unable to integrate new item [devhax] %v: %v", mi.ID, err))
						} else {
							// copy loaded values
							mi.Meta = fi.Meta
							mi.ItemTypeMeta = fi.ItemTypeMeta
							mi.ItemTypeName = fi.ItemTypeName
							mi.ItemFamilyID = fi.ItemFamilyID
							mi.ItemFamilyName = fi.ItemFamilyName
							mi.Process = fi.Process

							// store in container
							tc.Lock.Lock()
							defer tc.Lock.Unlock()

							tc.Items = append(tc.Items, mi)

							// mark as clean
							mi.CoreDirty = false
						}
					}
				}
			}
		}(mi, sol)
	}

	// clear new items from devhax
	sol.NewItemsDevHax = make([]*universe.NewItemFromNameTicketDevHax, 0)

	// iterate over new ships for devhax
	for _, rs := range sol.NewShipsDevHax {
		// handle escalation on another goroutine
		go func(rs *universe.NewShipFromNameTicketDevHax, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// find item type by name
			itemType, err := itemTypeSvc.GetItemTypeByName(rs.ItemTypeName)

			if err != nil || itemType == nil || itemType.Family != "ship" {
				shared.TeeLog(fmt.Sprintf("Unable to find ship item type [devhax] %v: %v", rs.ItemTypeName, err))
			} else {
				tStr, _ := universe.Meta(itemType.Meta).GetString("shiptemplateid")

				// create new ship for player
				ps, err := CreateShipForPlayer(
					rs.UserID,
					uuid.MustParse(tStr),
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
			}
		}(rs, sol)
	}

	// clear devhax new ships
	sol.NewShipsDevHax = make([]*universe.NewShipFromNameTicketDevHax, 0)

	// iterate over new sell orders
	for _, mi := range sol.NewSellOrders {
		// make sure we have waited long enough
		mi.CoreWait--

		if mi.CoreWait > -1 {
			continue
		}

		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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

				// for good measure
				mi.CoreWait = -999
			}
		}(mi, sol)
	}

	// clear handled new sell orders
	nso := make([]*universe.SellOrder, 0)

	for _, mi := range sol.NewSellOrders {
		if mi.CoreWait > -1 {
			// keep sell order
			nso = append(nso, mi)
		}
	}

	sol.NewSellOrders = nso

	// iterate over bought sell orders
	for _, mi := range sol.BoughtSellOrders {
		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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

			if err != nil {
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

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

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
			// handle escalation failure
			defer escalationRecover(sol, e)

			// remove client from system
			sol.RemoveClient(rs, true)

			// find their starting conditions
			start, err := startSvc.GetStartByID(rs.StartID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to respawn player %v - no starting conditions!", rs.UID))
				return
			}

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

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

			if err != nil {
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
			// handle escalation failure
			defer escalationRecover(sol, e)

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

	// iterate over new outpost tickets
	for _, rs := range sol.NewOutpostTickets {
		// handle escalation on another goroutine
		go func(rs *universe.NewOutpostTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// create new outpost for player
			po, err := CreateOutpostForPlayer(
				rs.UserID,
				rs.OutpostTemplateID,
				sol.ID,
				rs.PosX,
				rs.PosY,
				rs.Theta,
			)

			if err != nil || po == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to complete outpost deployment for %v - failure saving (%v)!", rs.UserID, err))
				return
			}

			// load into universe
			eo, err := LoadOutpost(po, sol.Universe)

			if err != nil || eo == nil {
				shared.TeeLog(fmt.Sprintf("! Unable to complete outpost deployment for %v - failure loading (%v)!", rs.UserID, err))
				return
			}

			// put outpost in system
			sol.AddOutpost(eo, true)
		}(rs, sol)
	}

	// clear new ship tickets
	sol.NewOutpostTickets = make([]*universe.NewOutpostTicket, 0)

	// iterate over ship switches
	for _, rs := range sol.ShipSwitches {
		// handle escalation on another goroutine
		go func(rs *universe.ShipSwitch, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// update currently flown ship in database
			err := userSvc.UpdateCurrentShipID(*rs.Client.UID, &rs.Target.ID)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

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
			// handle escalation failure
			defer escalationRecover(sol, e)

			// update name in database
			err := shipSvc.Rename(rs.ShipID, rs.Name)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to rename ship %v - failure saving (%v)!", rs.ShipID, err))
				return
			}
		}(rs, sol)
	}

	// clear ship renames
	sol.ShipRenames = make([]*universe.ShipRename, 0)

	// iterate over outpost renames
	for _, ro := range sol.OutpostRenames {
		// handle escalation on another goroutine
		go func(ro *universe.OutpostRename, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// update name in database
			err := outpostSvc.Rename(ro.OutpostID, ro.Name)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("! Unable to rename outpost %v - failure saving (%v)!", ro.OutpostID, err))
				return
			}
		}(ro, sol)
	}

	// clear outpost renames
	sol.OutpostRenames = make([]*universe.OutpostRename, 0)

	// iterate over clients in need of a schematic runs update
	for id := range sol.SchematicRunViews {
		// capture reference
		rs := sol.SchematicRunViews[id]

		// handle escalation on another goroutine
		go func(rs *shared.GameClient) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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

	// clear schematic run views
	sol.SchematicRunViews = make([]*shared.GameClient, 0)

	// iterate over newly invoked schematics
	for id := range sol.NewSchematicRuns {
		// capture reference
		rs := sol.NewSchematicRuns[id]

		// handle escalation on another goroutine
		go func(rs *universe.NewSchematicRunTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

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

	// clear new schematic runs
	sol.NewSchematicRuns = make([]*universe.NewSchematicRunTicket, 0)

	// iterate over new faction requests
	for id := range sol.NewFactions {
		// capture reference
		mi := sol.NewFactions[id]

		// handle escalation on another goroutine
		go func(mi *universe.NewFactionTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// obtain lock on solar system
			sol.Lock.Lock()
			defer sol.Lock.Unlock()

			// obtain lock on game client
			mi.Client.Lock.Lock()
			defer mi.Client.Lock.Unlock()

			// obtain lock on ship attached to client
			sh := sol.Universe.FindShip(mi.Client.CurrentShipID, &sol.ID)

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

			// obtain factions write lock
			sol.Universe.FactionsLock.Lock()
			defer sol.Universe.FactionsLock.Unlock()

			// load faction into universe
			uf := FactionFromSQL(f)
			sol.Universe.Factions[uf.ID.String()] = uf

			// put founder in their new faction
			err = userSvc.UpdateCurrentFactionID(*mi.Client.UID, &f.ID)

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

	// clear new faction requests
	sol.NewFactions = make([]*universe.NewFactionTicket, 0)

	// iterate over leave faction requests
	for id := range sol.LeaveFactions {
		// capture reference
		mi := sol.LeaveFactions[id]

		// handle escalation on another goroutine
		go func(mi *universe.LeaveFactionTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// obtain lock on solar system
			sol.Lock.Lock()
			defer sol.Lock.Unlock()

			// obtain lock on game client
			mi.Client.Lock.Lock()
			defer mi.Client.Lock.Unlock()

			// obtain lock on ship attached to client
			sh := sol.Universe.FindShip(mi.Client.CurrentShipID, &sol.ID)

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
			err = userSvc.UpdateCurrentFactionID(*mi.Client.UID, &fID)

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

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

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

	// clear leave faction requests
	sol.LeaveFactions = make([]*universe.LeaveFactionTicket, 0)

	// iterate over join faction requests
	for id := range sol.JoinFactions {
		// capture reference
		mi := sol.JoinFactions[id]

		// handle escalation on another goroutine
		go func(mi *universe.JoinFactionTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// null check
			if mi.OwnerClient == nil {
				shared.TeeLog(fmt.Sprintf("No owner client attached to join request: %v", mi))
				return
			}

			// obtain lock on solar system
			sol.Lock.Lock()
			defer sol.Lock.Unlock()

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

			// try to find the target faction
			faction := sol.Universe.Factions[mi.FactionID.String()]

			if faction == nil {
				shared.TeeLog(fmt.Sprintf("Unable to find target faction to join: %v", mi))
				return
			}

			// try to find applicant's game client
			appClient := sol.Universe.FindCurrentPlayerClient(mi.UserID, &sol.ID)

			if appClient != nil {
				// obtain lock on applicant's game client
				appClient.Lock.Lock()
				defer appClient.Lock.Unlock()
			}

			// try to find applicant's ship
			appShip := sol.Universe.FindCurrentPlayerShip(mi.UserID, &sol.ID)

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
			err := userSvc.UpdateCurrentFactionID(mi.UserID, &fID)

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

	// clear join faction requests
	sol.JoinFactions = make([]*universe.JoinFactionTicket, 0)

	// iterate over clients in need of a faction member list
	for id := range sol.ViewMembers {
		// capture reference
		rs := sol.ViewMembers[id]

		// handle escalation on another goroutine
		go func(rs *universe.ViewMembersTicket) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// null check
			if rs.OwnerClient == nil {
				return
			}

			// obtain lock on solar system
			sol.Lock.Lock()
			defer sol.Lock.Unlock()

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

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

	// clear view members
	sol.ViewMembers = make([]*universe.ViewMembersTicket, 0)

	// iterate over kick member requests
	for id := range sol.KickMembers {
		// capture reference
		mi := sol.KickMembers[id]

		// handle escalation on another goroutine
		go func(mi *universe.KickMemberTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// null check
			if mi.OwnerClient == nil {
				shared.TeeLog(fmt.Sprintf("No owner client attached to kick request: %v", mi))
				return
			}

			// obtain lock on solar system
			sol.Lock.Lock()
			defer sol.Lock.Unlock()

			// obtain factions read lock
			sol.Universe.FactionsLock.RLock()
			defer sol.Universe.FactionsLock.RUnlock()

			// try to find the source faction
			kickingFaction := sol.Universe.Factions[mi.FactionID.String()]

			if kickingFaction == nil {
				shared.TeeLog(fmt.Sprintf("Unable to find kicking faction: %v", mi))
				return
			}

			// try to find the kickee's game client
			kickeeClient := sol.Universe.FindCurrentPlayerClient(mi.UserID, &sol.ID)

			if kickeeClient != nil {
				// obtain lock on kickee's game client
				kickeeClient.Lock.Lock()
				defer kickeeClient.Lock.Unlock()
			}

			// try to find kickee's ship
			kickeeShip := sol.Universe.FindCurrentPlayerShip(mi.UserID, &sol.ID)

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
			err = userSvc.UpdateCurrentFactionID(mi.UserID, &sfID)

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

	// clear kick member requests
	sol.KickMembers = make([]*universe.KickMemberTicket, 0)

	// iterate over view action report page requests
	for id := range sol.ActionReportPages {
		// capture reference
		mi := sol.ActionReportPages[id]

		// handle escalation on another goroutine
		go func(mi *shared.ViewActionReportPageTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// null check
			if mi.Client == nil {
				shared.TeeLog(fmt.Sprintf("No client attached to action report request: %v", mi))
				return
			}

			// pull report summaries
			summaries, err := actionReportSvc.GetActionReportSummariesByUserID(*mi.Client.UID, mi.Page*mi.Take, mi.Take)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Error retrieving action report summary page: %v", err))
				return
			}

			// map for transmission
			page := models.ServerActionReportsPage{
				Page: mi.Page,
			}

			for _, v := range summaries {
				page.Logs = append(page.Logs, models.ServerActionReportSummary{
					ID:                     v.ID,
					VictimName:             v.VictimName,
					VictimShipTemplateName: v.VictimShipTemplateName,
					VictimTicker:           v.VictimTicker,
					Timestamp:              v.Timestamp,
					SystemName:             v.SolarSystemName,
					RegionName:             v.RegionName,
					Parties:                v.Parties,
				})
			}

			// send message to client containing page
			b, _ := json.Marshal(page)

			msg := models.GameMessage{
				MessageType: models.SharedMessageRegistry.ActionReportsPage,
				MessageBody: string(b),
			}

			go func() {
				mi.Client.WriteMessage(&msg)
			}()
		}(mi, sol)
	}

	// clear action report page requests
	sol.ActionReportPages = make([]*shared.ViewActionReportPageTicket, 0)

	// iterate over view action report detail requests
	for id := range sol.ActionReportDetails {
		// capture reference
		mi := sol.ActionReportDetails[id]

		// handle escalation on another goroutine
		go func(mi *shared.ViewActionReportDetailTicket, sol *universe.SolarSystem) {
			// handle escalation failure
			defer escalationRecover(sol, e)

			// null check
			if mi.Client == nil {
				shared.TeeLog(fmt.Sprintf("No client attached to action report detail request: %v", mi))
				return
			}

			// pull report
			report, err := actionReportSvc.GetActionReportByID(mi.KillID)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Error retrieving action report details: %v", err))
				return
			}

			// map for transmission
			log := models.ServerActionReportDetail{
				ID: report.ID,
				ServerKillLog: models.ServerKillLog{
					Wallet: report.Report.Wallet,
				},
			}

			// map header
			log.ServerKillLog.Header.VictimID = report.Report.Header.VictimID
			log.ServerKillLog.Header.VictimName = report.Report.Header.VictimName
			log.ServerKillLog.Header.VictimFactionID = report.Report.Header.VictimFactionID
			log.ServerKillLog.Header.VictimFactionName = report.Report.Header.VictimFactionName
			log.ServerKillLog.Header.VictimShipTemplateID = report.Report.Header.VictimShipTemplateID
			log.ServerKillLog.Header.VictimShipTemplateName = report.Report.Header.VictimShipTemplateName
			log.ServerKillLog.Header.VictimShipID = report.Report.Header.VictimShipID
			log.ServerKillLog.Header.VictimShipName = report.Report.Header.VictimShipName
			log.ServerKillLog.Header.Timestamp = report.Report.Header.Timestamp
			log.ServerKillLog.Header.SolarSystemID = report.Report.Header.SolarSystemID
			log.ServerKillLog.Header.SolarSystemName = report.Report.Header.SolarSystemName
			log.ServerKillLog.Header.RegionID = report.Report.Header.RegionID
			log.ServerKillLog.Header.RegionName = report.Report.Header.RegionName
			log.ServerKillLog.Header.HoldingFactionID = report.Report.Header.HoldingFactionID
			log.ServerKillLog.Header.HoldingFactionName = report.Report.Header.HoldingFactionName
			log.ServerKillLog.Header.InvolvedParties = report.Report.Header.InvolvedParties
			log.ServerKillLog.Header.IsNPC = report.Report.Header.IsNPC
			log.ServerKillLog.Header.PosX = report.Report.Header.PosX
			log.ServerKillLog.Header.PosY = report.Report.Header.PosY

			// map fitting
			for _, s := range report.Report.Fitting.ARack {
				log.ServerKillLog.Fitting.ARack = append(log.ServerKillLog.Fitting.ARack, models.ServerKillLogSlot{
					ItemID:              s.ItemID,
					ItemTypeID:          s.ItemTypeID,
					ItemFamilyID:        s.ItemFamilyID,
					ItemTypeName:        s.ItemTypeName,
					ItemFamilyName:      s.ItemFamilyName,
					IsModified:          s.IsModified,
					CustomizationFactor: s.CustomizationFactor,
				})
			}

			for _, s := range report.Report.Fitting.BRack {
				log.ServerKillLog.Fitting.BRack = append(log.ServerKillLog.Fitting.BRack, models.ServerKillLogSlot{
					ItemID:              s.ItemID,
					ItemTypeID:          s.ItemTypeID,
					ItemFamilyID:        s.ItemFamilyID,
					ItemTypeName:        s.ItemTypeName,
					ItemFamilyName:      s.ItemFamilyName,
					IsModified:          s.IsModified,
					CustomizationFactor: s.CustomizationFactor,
				})
			}

			for _, s := range report.Report.Fitting.CRack {
				log.ServerKillLog.Fitting.CRack = append(log.ServerKillLog.Fitting.CRack, models.ServerKillLogSlot{
					ItemID:              s.ItemID,
					ItemTypeID:          s.ItemTypeID,
					ItemFamilyID:        s.ItemFamilyID,
					ItemTypeName:        s.ItemTypeName,
					ItemFamilyName:      s.ItemFamilyName,
					IsModified:          s.IsModified,
					CustomizationFactor: s.CustomizationFactor,
				})
			}

			// map cargo
			for _, s := range report.Report.Cargo {
				log.ServerKillLog.Cargo = append(log.ServerKillLog.Cargo, models.ServerKillLogCargoItem{
					ItemID:         s.ItemID,
					ItemTypeID:     s.ItemTypeID,
					ItemFamilyID:   s.ItemFamilyID,
					ItemTypeName:   s.ItemTypeName,
					ItemFamilyName: s.ItemFamilyName,
					Quantity:       s.Quantity,
					IsPackaged:     s.IsPackaged,
				})
			}

			// map involved parties
			for _, p := range report.Report.InvolvedParties {
				q := models.ServerKillLogInvolvedParty{
					UserID:              p.UserID,
					FactionID:           p.FactionID,
					CharacterName:       p.CharacterName,
					FactionName:         p.FactionName,
					IsNPC:               p.IsNPC,
					LastAggressed:       p.LastAggressed,
					ShipID:              p.ShipID,
					ShipName:            p.ShipName,
					ShipTemplateID:      p.ShipTemplateID,
					ShipTemplateName:    p.ShipTemplateName,
					LastSolarSystemID:   p.LastSolarSystemID,
					LastSolarSystemName: p.LastSolarSystemName,
					LastRegionID:        p.LastRegionID,
					LastRegionName:      p.LastRegionName,
					LastPosX:            p.LastPosX,
					LastPosY:            p.LastPosY,
					WeaponUse:           make(map[string]*models.ServerKillLogWeaponUse),
				}

				for k, x := range p.WeaponUse {
					q.WeaponUse[k] = &models.ServerKillLogWeaponUse{
						ItemID:          x.ItemID,
						ItemTypeID:      x.ItemTypeID,
						ItemFamilyID:    x.ItemFamilyID,
						ItemFamilyName:  x.ItemFamilyName,
						ItemTypeName:    x.ItemTypeName,
						LastUsed:        x.LastUsed,
						DamageInflicted: x.DamageInflicted,
					}
				}

				log.ServerKillLog.InvolvedParties = append(log.ServerKillLog.InvolvedParties, q)
			}

			// send message to client containing report
			b, _ := json.Marshal(log)

			msg := models.GameMessage{
				MessageType: models.SharedMessageRegistry.ActionReportDetail,
				MessageBody: string(b),
			}

			go func() {
				mi.Client.WriteMessage(&msg)
			}()
		}(mi, sol)
	}

	// clear action report detail requests
	sol.ActionReportDetails = make([]*shared.ViewActionReportDetailTicket, 0)
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
