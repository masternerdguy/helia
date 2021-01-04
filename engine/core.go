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

//Will cause all system goroutines to stop when true
var shutdownSignal = false

//HeliaEngine Structure representing the core server-side game engine
type HeliaEngine struct {
	Universe *universe.Universe
}

//Initialize Initializes a new instance of the game engine
func (e *HeliaEngine) Initialize() *HeliaEngine {
	log.Println("Loading game universe from database...")

	//load universe
	u, err := loadUniverse()

	if err != nil {
		panic(fmt.Sprintf("Unable to load game universe: %v", err))
	}

	e.Universe = u

	//instantiate engine
	log.Println("Universe loaded!")
	engine := HeliaEngine{}

	return &engine
}

//Start Start the goroutines for each system
func (e *HeliaEngine) Start() {
	log.Println("Starting system goroutines...")

	for _, r := range e.Universe.Regions {
		for _, s := range r.Systems {
			go func(sol *universe.SolarSystem) {
				var tpf int = 0
				lastFrame := makeTimestamp()

				//game loop
				for {
					//check for shutdown signal
					if shutdownSignal {
						break
					}

					//for monitoring spawn exit
					c := make(chan struct{}, 1)

					//spawn goroutine so defer works as expected
					go func(sol *universe.SolarSystem) {
						//update system
						sol.PeriodicUpdate()

						//handle escalations from system
						handleEscalations(sol)

						//get time of last frame
						now := makeTimestamp()
						tpf = int(now - lastFrame)

						//find remaining portion of server heatbeat
						if tpf < universe.Heartbeat {
							//sleep for remainder of server heartbeat
							time.Sleep(time.Duration(universe.Heartbeat-tpf) * time.Millisecond)
						}

						//done!
						c <- struct{}{}
					}(sol)

					//wait for goroutine to return
					<-c

					//update last frame time
					lastFrame = makeTimestamp()
				}

				log.Println(fmt.Sprintf("System %s has halted.", sol.SystemName))
			}(s)
		}
	}

	log.Println("System goroutines started!")
}

//Shutdown Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	log.Println("! Server shutdown initiated")

	//shut down simulation
	log.Println("Stopping simulation...")
	shutdownSignal = true

	//wait for 30 seconds to give everything a chance to exit
	sleepStart := time.Now()

	for {
		if time.Now().Sub(sleepStart).Seconds() > 30 {
			break
		}

		time.Sleep(1000)
	}

	log.Println("Halt success assumed")

	//save progress
	log.Println("Saving world state...")
	saveUniverse(e.Universe)
	log.Println("World state saved!")

	//end program
	log.Println("Shutdown complete! Goodbye :)")
	os.Exit(0)
}

func handleEscalations(sol *universe.SolarSystem) {
	//get services
	startSvc := sql.GetStartService()
	shipSvc := sql.GetShipService()

	//obtain lock
	sol.Lock.Lock()
	defer sol.Lock.Unlock()

	//iterate over moved items
	for id := range sol.MovedItems {
		//capture reference and remove from map
		mi, _ := sol.MovedItems[id]
		delete(sol.MovedItems, id)

		//handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			//lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			//save new location of item to db
			err := saveItemLocation(mi.ID, mi.ContainerID)

			//error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to relocate item %v: %v", mi.ID, err))
			}
		}(mi, sol)
	}

	//iterate over packaged items
	for id := range sol.PackagedItems {
		//capture reference and remove from map
		mi, _ := sol.PackagedItems[id]
		delete(sol.PackagedItems, id)

		//handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			//lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			//mark item as packaged in the db
			err := packageItem(mi.ID)

			//error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to package item %v: %v", mi.ID, err))
			}
		}(mi, sol)
	}

	//iterate over unpackaged items
	for id := range sol.UnpackagedItems {
		//capture reference and remove from map
		mi, _ := sol.UnpackagedItems[id]
		delete(sol.UnpackagedItems, id)

		//handle escalation on another goroutine
		go func(mi *universe.Item, sol *universe.SolarSystem) {
			//lock item
			mi.Lock.Lock()
			defer mi.Lock.Unlock()

			//save unpackaged item to db
			err := unpackageItem(mi.ID, mi.Meta)

			//error check
			if err != nil {
				log.Println(fmt.Sprintf("Unable to unpackage item %v: %v", mi.ID, err))
			}
		}(mi, sol)
	}

	//iterate over dead ships
	for id := range sol.DeadShips {
		//capture reference and remove from map
		ds, _ := sol.DeadShips[id]
		delete(sol.DeadShips, id)

		//handle escalation on another goroutine
		go func(ds *universe.Ship, sol *universe.SolarSystem) {
			//remove ship from system
			sol.RemoveShip(ds, true)

			//make sure everything is set
			ds.Destroyed = true

			if ds.DestroyedAt == nil {
				now := time.Now()
				ds.DestroyedAt = &now
			}

			//update dead ship in db
			err := saveShip(ds)

			if err != nil {
				log.Println(fmt.Sprintf("! Unable to mark ship %v as dead in db (%v)!", ds.ID, err))
				return
			}
		}(ds, sol)
	}

	//iterate over clients in need of respawn
	for id := range sol.NeedRespawn {
		//capture reference and remove from map
		rs, _ := sol.NeedRespawn[id]
		delete(sol.NeedRespawn, id)

		//handle escalation on another goroutine
		go func(rs *shared.GameClient, sol *universe.SolarSystem) {
			//remove client from system
			sol.RemoveClient(rs, true)

			//find their starting conditions
			start, err := startSvc.GetStartByID(rs.StartID)

			if err != nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no starting conditions!", rs.UID))
				return
			}

			//find their home station
			home := sol.Universe.FindStation(start.HomeStationID)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - no home station!", rs.UID))
				return
			}

			//create their noob ship docked in that station
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

			ns.DockedAtStationID = &start.HomeStationID

			//save noob ship
			err = shipSvc.UpdateShip(*ns)

			if home == nil {
				log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't save noob ship changes (%v)!", rs.UID, err))
				return
			}

			//load the noob ship into the home station's system
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

			//put ship in home system
			home.CurrentSystem.AddShip(es, true)

			//set client current ship to new noob ship
			rs.CurrentShipID = es.ID

			//put the client in that system
			home.CurrentSystem.AddClient(rs, true)
		}(rs, sol)
	}
}
