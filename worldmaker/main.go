package main

import (
	"fmt"
	"helia/physics"
	"log"
	"math"
	"math/rand"
)

/*
 * Contains routines for procedurally generating the universe
 */

const MinSystemCount = 250
const MaxSystemCount = 500
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

	// output chosen values
	log.Println(fmt.Sprintf("Extent: %v, Systems: %v", extent, systemCount))

	// generate empty systems in a spiral
	emptySystems := generateEmptySystems(extent, systemCount)

	// dump
	dumpAcc := ""

	for _, e := range emptySystems {
		dumpAcc = fmt.Sprintf("%v(%v, %v)", dumpAcc, e.PosX, e.PosY)
	}

	log.Println(dumpAcc)
}

func generateEmptySystems(extent int, systemCount int) []Sysling {
	o := make([]Sysling, 0)

	for p := 0; p < systemCount; p++ {
		r := math.Sqrt(rand.Float64()) * float64(extent)

		x := float64(SpiralFactor)*r*math.Cos(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))
		y := float64(SpiralFactor)*r*math.Sin(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))

		o = append(o, Sysling{
			PosX: x,
			PosY: y,
		})
	}

	return o
}

// Represents a scaffolding for a solar system
type Sysling struct {
	PosX float64
	PosY float64
}
