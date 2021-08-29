package main

import (
	"fmt"
	"helia/engine"
	"helia/physics"
	"helia/universe"
	"log"
	"math"
	"math/rand"
)

/*
 * Contains routines for procedurally filling a scaffolded universe
 * (Designed to be run after world maker)
 */

func main() {
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

/* Parameters for asteroid generation */
const MinAsteroidsPerSystem = 0
const MaxAsteroidsPerSystem = 100

/*
============ ORE RARITY TABLE ========================================
id	                                    name	    probability	stop
dd522f03-2f52-4e82-b2f8-d7e0029cb82f	Testite	    0.1875	    0.1875
56617d30-6c30-425c-84bf-2484ae8c1156	Alri  	    0.1743	    0.3618
26a3fc9e-db2f-439d-a929-ba755d11d09c	Feymar	    0.1609    	0.5227
1d0d344b-ef28-43c8-a7a6-3275936b2dea	Listine	    0.0843   	0.607
0cd04eea-a150-410c-91eb-6af00d8c6eae	Hetrone	    0.0614  	0.6684
39b8eedf-ef80-4c29-a4bf-99abc4d84fa6	Novum	    0.0532  	0.7216
dd0c9b0a-279e-418e-b3b6-2f569fda0186	Suemetrium	0.0284  	0.75
7dcd5138-d7e0-419f-867a-6f0f23b99b5b	Jutrick	    0.0833  	0.8333
61f52ba3-654b-45cf-88e3-33399d12350d	Ovan	    0.0621  	0.8954
11688112-f3d4-4d30-864a-684a8b96ea23	Caiqua	    0.0382  	0.9336
2ce48bef-f06b-4550-b20c-0e64864db051	Zvitis	    0.0298  	0.9634
66b7a322-8cfc-4467-9410-492e6b58f159	Ichre	    0.0231  	0.9865
d1866be4-5c3e-4b95-b6d9-020832338014	Betro	    0.0135  	1

============ ASTEROID COUNT FOR SYSTEM ================================
actual = (1-scarcity^0.35) * potential

*/

type OreStop struct {
	ID   string
	Stop float64
}

func GetOreStops() []OreStop {
	o := make([]OreStop, 0)

	o = append(o, OreStop{
		ID:   "dd522f03-2f52-4e82-b2f8-d7e0029cb82f",
		Stop: 0.1875,
	})

	o = append(o, OreStop{
		ID:   "56617d30-6c30-425c-84bf-2484ae8c1156",
		Stop: 0.3618,
	})

	o = append(o, OreStop{
		ID:   "26a3fc9e-db2f-439d-a929-ba755d11d09c",
		Stop: 0.5227,
	})

	o = append(o, OreStop{
		ID:   "1d0d344b-ef28-43c8-a7a6-3275936b2dea",
		Stop: 0.6070,
	})

	o = append(o, OreStop{
		ID:   "0cd04eea-a150-410c-91eb-6af00d8c6eae",
		Stop: 0.6684,
	})

	o = append(o, OreStop{
		ID:   "39b8eedf-ef80-4c29-a4bf-99abc4d84fa6",
		Stop: 0.7216,
	})

	o = append(o, OreStop{
		ID:   "dd0c9b0a-279e-418e-b3b6-2f569fda0186",
		Stop: 0.7500,
	})

	o = append(o, OreStop{
		ID:   "7dcd5138-d7e0-419f-867a-6f0f23b99b5b",
		Stop: 0.8333,
	})

	o = append(o, OreStop{
		ID:   "61f52ba3-654b-45cf-88e3-33399d12350d",
		Stop: 0.8954,
	})

	o = append(o, OreStop{
		ID:   "11688112-f3d4-4d30-864a-684a8b96ea23",
		Stop: 0.9336,
	})

	o = append(o, OreStop{
		ID:   "2ce48bef-f06b-4550-b20c-0e64864db051",
		Stop: 0.9634,
	})

	o = append(o, OreStop{
		ID:   "66b7a322-8cfc-4467-9410-492e6b58f159",
		Stop: 0.9865,
	})

	o = append(o, OreStop{
		ID:   "d1866be4-5c3e-4b95-b6d9-020832338014",
		Stop: 1.0000,
	})

	return o
}

// Generates a seed integer for use as a system-specific RNG seed for consistent internal generation
func calculateSystemSeed(s *universe.SolarSystem) int {
	// use the system's position and uuid timestamp as a seed
	seed := (int(s.PosX*10000.0)>>int(math.Abs(s.PosY)*10000.0) + s.ID.ClockSequence())

	if s.PosY < 0 {
		seed *= -1
	}

	return seed
}

// Inserts minable asteroids into the universe
func dropAsteroids(u *universe.Universe) {
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// get system internal seed
			seed := calculateSystemSeed(s)

			// introduce unique offset for this function
			seed = seed + 50912

			// initialize RNG with seed
			rand.Seed(int64(seed))

			// get scarcity level
			scarcity := rand.Float64()

			// determine total number of asteroids in system
			potentialAsteroids := physics.RandInRange(MinAsteroidsPerSystem, MaxAsteroidsPerSystem)
			actualAsteroids := int((1.0 - math.Pow(scarcity, 0.35)) * float64(potentialAsteroids))

			// print level
			log.Println(fmt.Sprintf("%v | asteroids: %v scarcity: %v", s.SystemName, actualAsteroids, scarcity))
		}
	}
}
