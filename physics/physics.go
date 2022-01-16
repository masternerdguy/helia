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
func ElasticCollide(dummyA *Dummy, dummyB *Dummy) {
	// safety check
	if dummyA == nil || dummyB == nil {
		return
	}

	// get velocity and mass
	aVx := dummyA.VelX
	aVy := dummyA.VelY
	aM := dummyA.Mass
	bVx := dummyB.VelX
	bVy := dummyB.VelY
	bM := dummyB.Mass

	// guarantee some push effect in event of zero velocity
	if aVx < 1 {
		aVx = rand.Float64()
	}

	if aVx < 1 {
		aVy = rand.Float64()
	}

	if aVx < 1 {
		bVx = rand.Float64()
	}

	if aVx < 1 {
		bVy = rand.Float64()
	}

	// push them apart to avoid double counting and overlap
	dummyA.PosX = (dummyA.PosX - aVx*(2.0+rand.Float64()))
	dummyA.PosY = (dummyA.PosY - aVy*(2.0+rand.Float64()))
	dummyB.PosX = (dummyB.PosX - bVx*(2.0+rand.Float64()))
	dummyB.PosY = (dummyB.PosY - bVy*(2.0+rand.Float64()))

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
}

// Calculates the distance between the centers of 2 physics dummies and returns the result
func Distance(dummyA Dummy, dummyB Dummy) float64 {
	dx := dummyA.PosX - dummyB.PosX
	dy := dummyA.PosY - dummyB.PosY

	dx2 := dx * dx
	dy2 := dy * dy

	return math.Sqrt(dx2 + dy2)
}
