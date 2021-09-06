package main

import (
	"fmt"
	"helia/engine"
	"helia/physics"
	"helia/sql"
	"helia/universe"
	"log"
	"math"
	"math/rand"

	"github.com/google/uuid"
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

	// dropAsteroids(universe)
	dropSanctuaryStations(universe)
}

/* Parameters for asteroid generation */
const MinAsteroidsPerSystem = 0
const MaxAsteroidsPerSystem = 100
const MinAsteroidYield = 0.1
const MaxAsteroidYield = 5.0
const MinAsteroidRadius = 120
const MaxAsteroidRadius = 315
const MinAsteroidMass = 6000
const MaxAsteroidMass = 75000
const MinBeltRadius = 25000
const MaxBeltRadius = 1350000
const MinBeltWidth = 1000
const MaxBeltWidth = 10000

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
	ID      string
	Stop    float64
	Texture string
}

func GetOreStops() []OreStop {
	o := make([]OreStop, 0)

	o = append(o, OreStop{
		ID:   "dummy",
		Stop: 0,
	})

	o = append(o, OreStop{
		ID:      "dd522f03-2f52-4e82-b2f8-d7e0029cb82f",
		Stop:    0.1875,
		Texture: "Mega/asteroidR4",
	})

	o = append(o, OreStop{
		ID:      "56617d30-6c30-425c-84bf-2484ae8c1156",
		Stop:    0.3618,
		Texture: "Mega/asteroidR2",
	})

	o = append(o, OreStop{
		ID:      "26a3fc9e-db2f-439d-a929-ba755d11d09c",
		Stop:    0.5227,
		Texture: "Mega/asteroidR6",
	})

	o = append(o, OreStop{
		ID:      "1d0d344b-ef28-43c8-a7a6-3275936b2dea",
		Stop:    0.6070,
		Texture: "Mega/asteroidR3",
	})

	o = append(o, OreStop{
		ID:      "0cd04eea-a150-410c-91eb-6af00d8c6eae",
		Stop:    0.6684,
		Texture: "Mega/asteroidR9",
	})

	o = append(o, OreStop{
		ID:      "39b8eedf-ef80-4c29-a4bf-99abc4d84fa6",
		Stop:    0.7216,
		Texture: "Mega/asteroidR7",
	})

	o = append(o, OreStop{
		ID:      "dd0c9b0a-279e-418e-b3b6-2f569fda0186",
		Stop:    0.7500,
		Texture: "Mega/asteroidR1",
	})

	o = append(o, OreStop{
		ID:      "7dcd5138-d7e0-419f-867a-6f0f23b99b5b",
		Stop:    0.8333,
		Texture: "Mega/asteroidR12",
	})

	o = append(o, OreStop{
		ID:      "61f52ba3-654b-45cf-88e3-33399d12350d",
		Stop:    0.8954,
		Texture: "Mega/asteroidR5",
	})

	o = append(o, OreStop{
		ID:      "11688112-f3d4-4d30-864a-684a8b96ea23",
		Stop:    0.9336,
		Texture: "Mega/asteroidR11",
	})

	o = append(o, OreStop{
		ID:      "2ce48bef-f06b-4550-b20c-0e64864db051",
		Stop:    0.9634,
		Texture: "Mega/asteroidR8",
	})

	o = append(o, OreStop{
		ID:      "66b7a322-8cfc-4467-9410-492e6b58f159",
		Stop:    0.9865,
		Texture: "Mega/asteroidR13",
	})

	o = append(o, OreStop{
		ID:      "d1866be4-5c3e-4b95-b6d9-020832338014",
		Stop:    1.0000,
		Texture: "Mega/asteroidR10",
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

func dropSanctuaryStations(u *universe.Universe) {
	fID, err := uuid.Parse("b3d3fa9c-b21e-490f-b39e-128b3af12128")

	if err != nil {
		panic(err)
	}

	toSave := make([]sql.Station, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)))

			// is owned by sanctuary?
			if s.HoldingFactionID == fID {
				stars := s.CopyStars()
				planets := s.CopyPlanets()
				jumpholes := s.CopyJumpholes()
				asteroids := s.CopyAsteroids()
				stations := s.CopyStations()

				// build sanctuary station for the system
				stat := sql.Station{
					ID:          uuid.New(),
					SystemID:    s.ID,
					StationName: "Sanctuary Outpost",
					Texture:     "sanctuary-1",
					Radius:      525,
					Mass:        23700,
					Theta:       float64(physics.RandInRange(0, 360)),
					FactionID:   fID,
				}

				// determine position
				for {
					// generate random position in belt
					dW := float64(physics.RandInRange(25000, 75000))
					mag := dW
					the := 2.0 * math.Pi * rand.Float64()

					stat.PosX = (mag * math.Cos(the))
					stat.PosY = (mag * math.Sin(the))

					// check for overlap with forbidden objects
					sB := physics.Dummy{
						PosX: stat.PosX,
						PosY: stat.PosY,
					}

					for _, v := range stars {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range planets {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range stations {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range jumpholes {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range asteroids {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					// safe to store asteroid
					break
				}

				toSave = append(toSave, stat)
			}
		}
	}

	// get service
	stationSvc := sql.GetStationService()

	// save new stations
	for _, st := range toSave {
		err := stationSvc.NewStationWorldFiller(&st)

		if err != nil {
			panic(err)
		}
	}
}

// Inserts minable asteroids into the universe
func dropAsteroids(u *universe.Universe) {
	globalAsteroids := make([]Astling, 0)

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

			// get obstacles in system
			stars := s.CopyStars()
			planets := s.CopyPlanets()
			jumpholes := s.CopyJumpholes()

			// determine total number of asteroids in system
			potentialAsteroids := physics.RandInRange(MinAsteroidsPerSystem, MaxAsteroidsPerSystem)
			actualAsteroids := int((1.0 - math.Pow(scarcity, 0.35)) * float64(potentialAsteroids))

			// determine asteroid distribution scale
			beltRadius := physics.RandInRange(MinBeltRadius, MaxBeltRadius)

			// use first star as belt center
			beltX := 0.0
			beltY := 0.0

			for _, e := range stars {
				beltX = e.PosX
				beltY = e.PosY
			}

			// get ore stops
			stops := GetOreStops()

			// generated localAsteroids
			localAsteroids := make([]Astling, 0)

			// make asteroids
			for i := 0; i < actualAsteroids; i++ {
				// roll to determine yield
				y := math.Max(rand.Float64()*MaxAsteroidYield, MinAsteroidYield) * (1 - scarcity)

				// roll to determine ore type
				p := rand.Float64()

				for x := range stops {
					// skip dummy stop
					if x == 0 {
						continue
					}

					// determine if this stop captures the roll
					floor := stops[x-1]
					ceiling := stops[x]

					if floor.Stop < p && p <= ceiling.Stop {
						ast := Astling{
							ID:            uuid.New().String(),
							SystemID:      s.ID.String(),
							OreItemTypeID: stops[x].ID,
							Name:          fmt.Sprintf("A%v%v%v-%v", stops[x].ID[0], stops[x].ID[1], stops[x].ID[2], physics.RandInRange(1000, 9999)),
							Texture:       stops[x].Texture,
							Radius:        float64(physics.RandInRange(MinAsteroidRadius, MaxAsteroidRadius)),
							Theta:         float64(rand.Float64() * 360.0),
							Yield:         y,
							Mass:          float64(physics.RandInRange(MinAsteroidMass, MaxAsteroidMass)),
						}

						// determine position
						for {
							// generate random position in belt
							dW := float64(physics.RandInRange(MinBeltWidth, MaxBeltWidth))
							mag := dW + float64(beltRadius)
							the := 2.0 * math.Pi * rand.Float64()

							ast.PosX = (mag * math.Cos(the)) + beltX
							ast.PosY = (mag * math.Sin(the)) + beltY

							// check for overlap with forbidden objects
							sB := physics.Dummy{
								PosX: ast.PosX,
								PosY: ast.PosY,
							}

							for _, v := range stars {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range planets {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range jumpholes {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range localAsteroids {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							// safe to store asteroid
							break
						}

						// store asteroid
						localAsteroids = append(localAsteroids, ast)

						// next asteroid
						break
					}
				}
			}

			// store asteroids to save
			globalAsteroids = append(globalAsteroids, localAsteroids...)
		}
	}

	// save generated asteroids
	astSvc := sql.GetAsteroidService()

	for _, ast := range globalAsteroids {
		id, err := uuid.Parse(ast.ID)

		if err != nil {
			panic(err)
		}

		systemID, err := uuid.Parse(ast.SystemID)

		if err != nil {
			panic(err)
		}

		oreID, err := uuid.Parse(ast.OreItemTypeID)

		if err != nil {
			panic(err)
		}

		f := sql.Asteroid{
			ID:         id,
			SystemID:   systemID,
			ItemTypeID: oreID,
			Name:       ast.Name,
			Texture:    ast.Texture,
			Radius:     ast.Radius,
			Theta:      ast.Theta,
			PosX:       ast.PosX,
			PosY:       ast.PosY,
			Yield:      ast.Yield,
			Mass:       ast.Mass,
		}

		err = astSvc.NewAsteroidWorldFiller(&f)

		if err != nil {
			panic(err)
		}
	}
}

// Structure representing a scaffold for an asteroid for worldfiller
type Astling struct {
	ID            string
	SystemID      string
	OreItemTypeID string
	Name          string
	Texture       string
	Radius        float64
	Theta         float64
	PosX          float64
	PosY          float64
	Yield         float64
	Mass          float64
}
