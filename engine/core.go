package engine

import (
	"fmt"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"log"
	"os"
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
	log.Println("Loading game universe from database...")

	// load universe
	u, err := LoadUniverse()

	if err != nil {
		panic(fmt.Sprintf("Unable to load game universe: %v", err))
	}

	// generate transient celestials
	log.Println("Generating transient celestials...")
	u.BuildTransientCelestials()

	// cache starmap
	log.Println("Building starmap...")
	err = u.BuildMapWithCache()

	if err != nil {
		panic(fmt.Sprintf("Unable to build starmap: %v", err))
	}

	e.Universe = u

	// instantiate engine
	log.Println("Universe ready!")
	engine := HeliaEngine{}

	return &engine
}

// Start the goroutines for each system
func (e *HeliaEngine) Start() {
	log.Println("Starting system goroutines...")

	for _, r := range e.Universe.Regions {
		for _, s := range r.Systems {
			go func(sol *universe.SolarSystem) {
				var tpf int = 0
				lastFrame := makeTimestamp()

				// game loop
				for {
					// check for shutdown signal
					if shutdownSignal {
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

				log.Println(fmt.Sprintf("System %s has halted.", sol.SystemName))
			}(s)
		}
	}

	log.Println("System goroutines started!")

	// start watchdog goroutine (to alert of any deadlocks for debugging purposes)
	go func(e *HeliaEngine) {
		for {
			log.Println("* Deadlock Check Starting")

			// iterate over systems
			for _, r := range e.Universe.Regions {
				for _, s := range r.Systems {
					log.Println(fmt.Sprintf("* Testing [%v]", s.SystemName))

					// test locks
					s.TestLocks()
					log.Println(fmt.Sprintf("* [%v] Passed", s.SystemName))
				}

				// small sleep between systems
				time.Sleep(20 * time.Millisecond)
			}

			log.Println("* All systems passed deadlock check!")

			// wait 60 seconds
			time.Sleep(time.Second * 60)
		}
	}(e)
}

// Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	log.Println("! Server shutdown initiated")

	// shut down simulation
	log.Println("Stopping simulation...")
	shutdownSignal = true

	// wait for 30 seconds to give everything a chance to exit
	sleepStart := time.Now()

	for {
		if time.Since(sleepStart).Seconds() > 30 {
			break
		}

		time.Sleep(1000)
	}

	log.Println("Halt success assumed")

	// save progress
	log.Println("Saving world state...")
	saveUniverse(e.Universe)
	log.Println("World state saved!")

	// end program
	log.Println("Shutdown complete! Goodbye :)")
	os.Exit(0)
}

func handleEscalations(sol *universe.SolarSystem) {
	// get services
	startSvc := sql.GetStartService()
	shipSvc := sql.GetShipService()
	userSvc := sql.GetUserService()

	// obtain lock
	sol.Lock.Lock()
	defer sol.Lock.Unlock()

	// iterate over moved items
	for id := range sol.MovedItems {
		// capture reference and remove from map
		mi := sol.MovedItems[id]
		delete(sol.MovedItems, id)

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
				log.Println(fmt.Sprintf("Unable to relocate item %v: %v", mi.ID, err))
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
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// mark item as packaged in the db
			err := packageItem(mi.ID)

			// error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to package item %v: %v", mi.ID, err))
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
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save unpackaged item to db
			err := unpackageItem(mi.ID, mi.Meta)

			// error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to unpackage item %v: %v", mi.ID, err))
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
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark as dirty if not marked already
			mi.CoreDirty = true

			// save quantity of item to db
			err := changeQuantity(mi.ID, mi.Quantity)

			// error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to change quantity of item %v: %v", mi.ID, err))
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
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// save new item to db
			id, err := newItem(mi)

			// error check
			if err != nil || id == nil {
				log.Println(fmt.Sprintf("Unable to save new item %v: %v", mi.ID, err))
			} else {
				// store corrected id from db insert
				mi.ID = *id

				// mark as clean
				mi.CoreDirty = false
			}
		}(mi, sol)
	}

	// iterate over new sell orders
	for id := range sol.NewSellOrders {
		// capture reference and remove from map
		mi := sol.NewSellOrders[id]
		delete(sol.NewSellOrders, id)

		// handle escalation on another goroutine
		go func(mi *universe.SellOrder, sol *universe.SolarSystem) {
			// lock sell order
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// save new sell order to db
			id, err := newSellOrder(mi)

			// error check
			if err != nil || id == nil {
				log.Println(fmt.Sprintf("Unable to save new sell order %v: %v", mi.ID, err))
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
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			// mark sell order as bought in db
			err := markSellOrderAsBought(mi)

			// error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to update bought sell order %v: %v", mi.ID, err))
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

			// update dead ship in db
			err := saveShip(ds)

			if err != nil {
				log.Println(fmt.Sprintf("! Unable to mark ship %v as dead in db (%v)!", ds.ID, err))
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
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - no associated user!", rs.UserID))
				return
			}

			start, err := startSvc.GetStartByID(user.StartID)

			if err != nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - no associated start!", rs.UserID))
				return
			}

			// find their home station
			home := sol.Universe.FindStation(start.HomeStationID)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - no home station!", rs.UserID))
				return
			}

			// create their ship docked in that station
			u, err := CreateNoobShipForPlayer(start, rs.UserID)

			if err != nil || u.CurrentShipID == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - failed to create noob ship (%v | %v)!", rs.UserID, err, u.CurrentShipID))
				return
			}

			ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - no noob ship!", rs.UserID))
				return
			}

			// save ship
			err = shipSvc.UpdateShip(*ns)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - couldn't save noob ship changes (%v)!", rs.UserID, err))
				return
			}

			// load the ship into the home station's system
			ns, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - no noob ship again!", rs.UserID))
				return
			}

			es, err := LoadShip(ns, sol.Universe)

			if err != nil || es == nil {
				log.Println(fmt.Sprintf("! Unable to respawn NPC %v - couldn't load new noobship into universe (%v)!", rs.UserID, err))
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
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no starting conditions!", rs.UID))
				return
			}

			// find their home station
			home := sol.Universe.FindStation(start.HomeStationID)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no home station!", rs.UID))
				return
			}

			// create their noob ship docked in that station
			u, err := CreateNoobShipForPlayer(start, *rs.UID)

			if err != nil || u.CurrentShipID == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - failed to create noob ship (%v | %v)!", rs.UID, err, u.CurrentShipID))
				return
			}

			ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no noob ship!", rs.UID))
				return
			}

			// save noob ship
			err = shipSvc.UpdateShip(*ns)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't save noob ship changes (%v)!", rs.UID, err))
				return
			}

			// load the noob ship into the home station's system
			ns, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no noob ship again!", rs.UID))
				return
			}

			es, err := LoadShip(ns, sol.Universe)

			if err != nil || es == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't load new noobship into universe (%v)!", rs.UID, err))
				return
			}

			// put ship in home system
			home.CurrentSystem.AddShip(es, true)

			// link player's faction into ship
			es.FactionID = u.CurrentFactionID

			// link player's reputation sheet into ship
			es.ReputationSheet = &rs.ReputationSheet

			// set client current ship to new noob ship
			rs.CurrentShipID = es.ID
			es.BeingFlownByPlayer = true

			// put the client in that system
			home.CurrentSystem.AddClient(rs, true)
		}(rs, sol)
	}
}
