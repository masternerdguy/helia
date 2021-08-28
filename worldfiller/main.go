package main

import (
	"helia/engine"
	"helia/universe"
	"log"
	"math/rand"
)

/*
 * Contains routines for procedurally filling a scaffolded universe
 * (Designed to be run after world maker)
 */

func main() {
	// initialize RNG
	rand.Seed(102)

	// load universe from database
	log.Println("Loading universe from database...")
	universe, err := engine.LoadUniverse()

	if err != nil {
		panic(err)
	}

	log.Println("Loaded universe!")

	/*
	 * COMMENT AND UNCOMMENT THE BELOW ROUTINES AS NEEDED
	 */

	dropAsteroids(universe)
}

// Inserts minable asteroids into the universe
func dropAsteroids(universe *universe.Universe) {

}
