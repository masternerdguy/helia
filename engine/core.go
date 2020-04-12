package engine

import (
	"fmt"
	"helia/universe"
)

//HeliaEngine Structure representing the core server-side game engine
type HeliaEngine struct {
	universe *universe.Universe
}

//Initialize Initializes a new instance of the game engine
func (e HeliaEngine) Initialize() *HeliaEngine {
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
func (e HeliaEngine) Start() {

}
