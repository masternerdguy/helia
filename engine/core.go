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
					go func() {
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
					}()

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

	//iterate over dead ships
	for id := range sol.DeadShips {
		ds, _ := sol.DeadShips[id]

		//remove ship from system
		sol.RemoveShip(ds, false)

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
			continue
		}
	}

	//reset dead ship list
	sol.DeadShips = make(map[string]*universe.Ship)

	//iterate over clients in need of respawn
	for id := range sol.NeedRespawn {
		rs, _ := sol.NeedRespawn[id]

		//remove client from system
		sol.RemoveClient(rs, false)

		//find their starting conditions
		start, err := startSvc.GetStartByID(rs.StartID)

		if err != nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - no starting conditions!", rs.UID))
			continue
		}

		//find their home station
		home := sol.Universe.FindStation(start.HomeStationID)

		if home == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - no home station!", rs.UID))
			continue
		}

		//create their noob ship docked in that station
		u, err := CreateNoobShipForPlayer(start, *rs.UID)

		if err != nil || u.CurrentShipID == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - failed to create noob ship (%v | %v)!", rs.UID, err, u.CurrentShipID))
			continue
		}

		ns, err := shipSvc.GetShipByID(*u.CurrentShipID, false)

		if err != nil || ns == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - no noob ship!", rs.UID))
			continue
		}

		ns.DockedAtStationID = &start.HomeStationID

		//save noob ship
		err = shipSvc.UpdateShip(*ns)

		if home == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't save noob ship changes (%v)!", rs.UID, err))
			continue
		}

		//load the noob ship into the home station's system
		ns, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

		if err != nil || ns == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - no noob ship again!", rs.UID))
			continue
		}

		es, err := loadShip(ns)

		if err != nil || es == nil {
			log.Println(fmt.Sprintf("! Unable to respawn player %v - couldn't load new noobship into universe (%v)!", rs.UID, err))
			continue
		}

		//don't lock if home station is in same system
		sameSystem := home.CurrentSystem.ID == sol.ID

		//put ship in home system
		home.CurrentSystem.AddShip(es, !sameSystem)

		//set client current ship to new noob ship
		rs.CurrentShipID = es.ID

		//put the client in that system
		home.CurrentSystem.AddClient(rs, !sameSystem)
	}

	//reset respawn client list
	sol.NeedRespawn = make(map[string]*shared.GameClient)
}
