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

const MinSystemCount = 250
const MaxSystemCount = 500
const MinRegions = 9
const MaxRegions = 27
const MinExtent = 30
const MaxExtent = 100

const SpiralFactor = 3

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

	// sort into regions
	regions := generateRegions(emptySystems, regionCount)

	// dump
	dumpAcc := ""

	for _, e := range regions {
		dumpAcc = fmt.Sprintf("%v(%v,%v)", dumpAcc, e.PosX, e.PosY)
	}

	log.Println(dumpAcc)
}

func generateRegions(systems []Sysling, regionCount int) []Regionling {
	o := make([]Regionling, 0)
	id := uuid.New()

	for p := 0; p < regionCount; p++ {
		// pick a system at random that doesn't have a region
		for {
			exit := false

			sIdx := physics.RandInRange(0, len(systems))
			s := systems[sIdx]

			if s.RegionID == nil {
				s.RegionID = &id
				systems[sIdx] = s

				// this will be the first system in the region
				r := Regionling{
					ID:   s.ID,
					PosX: s.PosX,
					PosY: s.PosY,
				}

				o = append(o, r)

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

func generateEmptySystems(extent int, systemCount int) []Sysling {
	o := make([]Sysling, 0)

	for p := 0; p < systemCount; p++ {
		r := math.Sqrt(rand.Float64()) * float64(extent)

		x := float64(SpiralFactor)*r*math.Cos(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))
		y := float64(SpiralFactor)*r*math.Sin(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))

		o = append(o, Sysling{
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
}

// Represents a scaffolding for a region
type Regionling struct {
	ID      uuid.UUID
	PosX    float64
	PosY    float64
	Systems []Sysling
}
