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
	u, err := loadUniverse()

	if err != nil {
		panic(fmt.Sprintf("Unable to load game universe: %v", err))
	}

	e.Universe = u

	// instantiate engine
	log.Println("Universe loaded!")
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

	// iterate over clients in need of respawn
	for id := range sol.NeedRespawn {
		// capture reference and remove from map
		rs := sol.NeedRespawn[id]
		delete(sol.NeedRespawn, id)

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
			u, err := CreateNoobShipForPlayer(start, *rs.UID, false)

			if err != nil || u.CurrentShipID == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - failed to create noob ship (%v | %v)!", rs.UID, err, u.CurrentShipID))
				return
			}

			ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if err != nil || ns == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no noob ship!", rs.UID))
				return
			}

			ns.DockedAtStationID = &start.HomeStationID

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

			es, err := LoadShip(ns)

			if err != nil || es == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't load new noobship into universe (%v)!", rs.UID, err))
				return
			}

			// put ship in home system
			home.CurrentSystem.AddShip(es, true)

			// set client current ship to new noob ship
			rs.CurrentShipID = es.ID
			es.BeingFlownByPlayer = true

			// put the client in that system
			home.CurrentSystem.AddClient(rs, true)
		}(rs, sol)
	}
}
