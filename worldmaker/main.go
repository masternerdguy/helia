package main

import (
	"fmt"
	"helia/physics"
	"log"
	"math"
	"math/rand"
	"sort"

	"github.com/google/uuid"
)

/*
 * Contains routines for procedurally generating the universe
 */

const MinSystemCount = 1250
const MaxSystemCount = 1500
const MinRegions = 27
const MaxRegions = 45
const MinExtent = 3
const MaxExtent = 10
const connectivity = 6

const SpiralFactor = 30

func main() {
	// initialize RNG
	rand.Seed(101)

	// determine size of map
	extent := physics.RandInRange(MinExtent, MaxExtent)

	// determine how many systems to generate
	systemCount := physics.RandInRange(MinSystemCount, MaxSystemCount)

	// determine how many regions to sort them into
	regionCount := physics.RandInRange(MinRegions, MaxRegions)

	// output chosen values
	log.Println(fmt.Sprintf("Extent: %v, Regions: %v, Systems: %v", extent, regionCount, systemCount))

	// generate empty systems in a spiral
	systems := generateEmptySystems(extent, systemCount)

	// generate empty regions in map
	regions := generateEmptyRegions(systems, regionCount)

	// sort each system into the region closest to it
	for _, s := range systems {
		if s.RegionID == nil {
			var closest *Regionling = nil
			var distance float64 = 0.0

			for _, r := range regions {
				// initialize if needed
				if closest == nil {
					closest = r
					distance = physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})
				}

				// check distance to region
				d := physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})

				if d < distance {
					// new closest region
					closest = r
					distance = d
				}
			}

			s.RegionID = &closest.ID
			closest.Systems = append(closest.Systems, s)
		}
	}

	/*// dump
	for _, r := range regions {
		dumpAcc := ""

		for _, s := range r.Systems {
			dumpAcc = fmt.Sprintf("%v(%v,%v)", dumpAcc, s.PosX, s.PosY)
		}

		log.Println("==========")
		log.Println(dumpAcc)
	}*/

	// name systems
	usedSystemNames := make(map[string]*Sysling)

	for _, s := range systems {
		for {
			n := randomPlaceholderName()

			if usedSystemNames[n] == nil {
				s.Name = n
				usedSystemNames[n] = s

				break
			}
		}
	}

	// name regions
	usedRegionNames := make(map[string]*Regionling)

	for _, s := range regions {
		for {
			n := randomPlaceholderName()

			if usedRegionNames[n] == nil {
				s.Name = n
				usedRegionNames[n] = s

				break
			}
		}
	}

	// generate jumphole connections
	jumpNetworkEdges := make([]Edgeling, 0)

	for _, s := range systems {
		systemsTmp := systems

		// sort by distance from this system
		sort.SliceStable(systemsTmp, func(i, j int) bool {
			ie := systemsTmp[i]
			je := systemsTmp[j]

			d1 := physics.Distance(physics.Dummy{PosX: s.PosX, PosY: s.PosY}, physics.Dummy{PosX: ie.PosX, PosY: ie.PosY})
			d2 := physics.Distance(physics.Dummy{PosX: s.PosX, PosY: s.PosY}, physics.Dummy{PosX: je.PosX, PosY: je.PosY})

			return d1 < d2
		})

		// get top entries
		linkCount := physics.RandInRange(1, connectivity)
		links := make([]*Sysling, 0)

		for v := 1; v < linkCount+1; v++ {
			links = append(links, systemsTmp[v])
		}

		// insert edges into network map
		for _, l := range links {
			// make sure there isn't already an equivalent edge
			safe := true

			for _, e := range jumpNetworkEdges {
				if e.A == s && e.B == l {
					safe = false
					break
				}

				if e.A == l && e.B == s {
					safe = false
					break
				}
			}

			if safe {
				// insert edge
				jumpNetworkEdges = append(jumpNetworkEdges, Edgeling{
					A: s,
					B: l,
				})
			}
		}

		// dump
		for _, e := range jumpNetworkEdges {
			log.Println(fmt.Sprintf("%v :: %v", e.A.Name, e.B.Name))
		}
	}
}

func randomPlaceholderName() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("1234567890")

	acc := ""

	firstLength := physics.RandInRange(1, 5)
	secondLength := physics.RandInRange(1, 5)

	for i := 0; i < firstLength; i++ {
		idx := physics.RandInRange(0, len(letters))
		acc = fmt.Sprintf("%v%v", acc, string(letters[idx]))
	}

	acc = fmt.Sprintf("%v%v", acc, "-")

	for i := 0; i < secondLength; i++ {
		idx := physics.RandInRange(0, len(numbers))
		acc = fmt.Sprintf("%v%v", acc, string(numbers[idx]))
	}

	return acc
}

func generateEmptyRegions(systems []*Sysling, regionCount int) []*Regionling {
	o := make([]*Regionling, 0)

	for p := 0; p < regionCount; p++ {
		id := uuid.New()

		// pick a system at random that doesn't have a region
		for {
			exit := false

			sIdx := physics.RandInRange(0, len(systems))
			s := systems[sIdx]

			if s.RegionID == nil {
				s.RegionID = &id

				// this will be the first system in the region
				r := Regionling{
					ID:      s.ID,
					PosX:    s.PosX,
					PosY:    s.PosY,
					Systems: make([]*Sysling, 0),
				}

				r.Systems = append(r.Systems, s)

				// store region
				o = append(o, &r)

				// done
				exit = true
			}

			// exit if done
			if exit {
				break
			}
		}
	}

	return o
}

func generateEmptySystems(extent int, systemCount int) []*Sysling {
	o := make([]*Sysling, 0)

	for p := 0; p < systemCount; p++ {
		r := math.Sqrt(rand.Float64()) * float64(extent)

		x := float64(SpiralFactor)*r*math.Cos(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))
		y := float64(SpiralFactor)*r*math.Sin(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))

		o = append(o, &Sysling{
			ID:   uuid.New(),
			PosX: x,
			PosY: y,
		})
	}

	return o
}

// Represents a scaffolding for a solar system
type Sysling struct {
	ID       uuid.UUID
	PosX     float64
	PosY     float64
	RegionID *uuid.UUID
	Name     string
}

// Represents a scaffolding for a region
type Regionling struct {
	ID      uuid.UUID
	PosX    float64
	PosY    float64
	Systems []*Sysling
	Name    string
}

// Represents an edge in the jumphole network
type Edgeling struct {
	A *Sysling
	B *Sysling
}
