package engine

import (
	"fmt"
	"helia/universe"
	"log"
	"time"
)

//HeliaEngine Structure representing the core server-side game engine
type HeliaEngine struct {
	universe *universe.Universe
}

//Initialize Initializes a new instance of the game engine
func (e *HeliaEngine) Initialize() *HeliaEngine {
	//load universe
	u, err := loadUniverse()

	if err != nil {
		panic(fmt.Sprintf("Unable to load game universe: %v", err))
	}

	e.universe = u

	//instantiate engine
	engine := HeliaEngine{}

	return &engine
}

//Start Start the goroutines for each system
func (e *HeliaEngine) Start() {
	log.Println(fmt.Sprintf("fupdating %v", e))
	for _, r := range e.universe.Regions {
		for _, s := range r.Systems {
			go func(sol *universe.SolarSystem) {
				//game loop
				for {
					//update system
					sol.PeriodicUpdate()

					//sleep for 1/4 of a second (this is the server heartbeat)
					time.Sleep(250)
				}
			}(s)
		}
	}
}
