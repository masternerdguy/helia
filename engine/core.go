package engine

import (
	"fmt"
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
				lastFrame := makeTimestamp()

				//game loop
				for {
					//check for shutdown signal
					if shutdownSignal {
						break
					}

					//update system
					sol.PeriodicUpdate()

					//get time of last frame
					now := makeTimestamp()
					tpf := int(now - lastFrame)

					//find remaining portion of server heatbeat
					if tpf < universe.Heartbeat {
						//sleep for remainder of server heartbeat
						time.Sleep(time.Duration(universe.Heartbeat-tpf) * time.Millisecond)
					}

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
