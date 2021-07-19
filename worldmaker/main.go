package main

import (
	"fmt"
	"helia/physics"
	"log"
	"math"
	"math/rand"

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
	emptySystems := generateEmptySystems(extent, systemCount)

	// generate empty regions in map
	regions := generateEmptyRegions(emptySystems, regionCount)

	// sort each system into the region closest to it
	for _, s := range emptySystems {
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

	for _, s := range emptySystems {
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
