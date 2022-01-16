package physics

import (
	"math"
	"math/rand"
)

// Dummy structure used to pass key values in physics calculations
type Dummy struct {
	PosX float64
	PosY float64
	VelX float64
	VelY float64
	Mass float64
}

// Calculates the result of a 2D elastic collission between two dummies and stores the result in those dummies
func ElasticCollide(dummyA *Dummy, dummyB *Dummy, systemRadius float64) float64 {
	// safety check
	if dummyA == nil || dummyB == nil {
		return 0
	}

	// get velocity and mass
	aVx := dummyA.VelX
	aVy := dummyA.VelY
	aM := dummyA.Mass
	bVx := dummyB.VelX
	bVy := dummyB.VelY
	bM := dummyB.Mass

	// push them apart to avoid double counting and overlap
	iter := 0

	for {
		dummyA.PosX = (dummyA.PosX - Sign(aVx)*rand.Float64()*5.0)
		dummyA.PosY = (dummyA.PosY - Sign(aVy)*rand.Float64()*5.0)
		dummyB.PosX = (dummyB.PosX - Sign(bVx)*rand.Float64()*5.0)
		dummyB.PosY = (dummyB.PosY - Sign(bVy)*rand.Float64()*5.0)

		// exit if no longer touching
		nD := Distance(*dummyA, *dummyB)

		if nD > systemRadius*1.01 {
			break
		}

		// prevents infinite loop
		if iter > 5 {
			break
		}

		iter++
	}

	// determine center of mass's velocity
	cVx := (aVx*aM + bVx*bM) / (aM + bM)
	cVy := (aVy*aM + bVy*bM) / (aM + bM)

	// randomize mix angle
	mix := rand.Float64()
	antiMix := 1.0 - mix

	// reverse directions and de-reference frame
	aVx2 := -aVx + (cVx * mix)
	aVy2 := -aVy + (cVy * mix)
	bVx2 := -bVx + (cVx * antiMix)
	bVy2 := -bVy + (cVy * antiMix)

	// store
	dummyA.VelX = aVx2
	dummyA.VelY = aVy2
	dummyB.VelX = bVx2
	dummyB.VelY = bVy2

	// return mixing angle
	return mix
}

// Calculates the distance between the centers of 2 physics dummies and returns the result
func Distance(dummyA Dummy, dummyB Dummy) float64 {
	dx := dummyA.PosX - dummyB.PosX
	dy := dummyA.PosY - dummyB.PosY

	dx2 := dx * dx
	dy2 := dy * dy

	return math.Sqrt(dx2 + dy2)
}

func Sign(a float64) float64 {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}

	return 1
}
