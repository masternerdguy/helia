package engine

import (
	"fmt"
	"helia/universe"
	"log"
	"os"
	"time"
)

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
				//game loop
				for {
					//update system
					sol.PeriodicUpdate()

					//sleep for server heartbeat
					time.Sleep(universe.Heartbeat)
				}
			}(s)
		}
	}

	log.Println("System goroutines started!")
}

//Shutdown Saves the current state of the simulation and halts
func (e *HeliaEngine) Shutdown() {
	log.Println("Server shutdown initiated")

	//end program

	log.Println("Shutdown complete! Goodbye :)")
	os.Exit(0)
}
