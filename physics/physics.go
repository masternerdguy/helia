package physics

import (
	"math"
)

//Dummy Dummy structure used to pass key values in physics calculations
type Dummy struct {
	PosX float64
	PosY float64
	VelX float64
	VelY float64
	Mass float64
}

//ElasticCollide Calculates the result of a 2D elastic collission between two dummies and stores the result in those dummies
func ElasticCollide(dummyA *Dummy, dummyB *Dummy, tpf float64) {
	// safety check
	if dummyA == nil || dummyB == nil {
		return
	}

	//get velocity and mass
	aVx := dummyA.VelX
	aVy := dummyA.VelY
	aM := dummyA.Mass
	bVx := dummyB.VelX
	bVy := dummyB.VelY
	bM := dummyB.Mass

	//push them apart to avoid double counting and overlap
	dummyA.PosX = (dummyA.PosX - aVx*tpf*2.0)
	dummyA.PosY = (dummyA.PosY - aVy*tpf*2.0)
	dummyB.PosX = (dummyB.PosX - bVx*tpf*2.0)
	dummyB.PosY = (dummyB.PosY - bVy*tpf*2.0)

	//determine center of mass's velocity
	cVx := (aVx*aM + bVx*bM) / (aM + bM)
	cVy := (aVy*aM + bVy*bM) / (aM + bM)

	//reverse directions and de-reference frame
	aVx2 := -aVx + cVx
	aVy2 := -aVy + cVy
	bVx2 := -bVx + cVx
	bVy2 := -bVy + cVy

	//store
	dummyA.VelX = aVx2
	dummyA.VelY = aVy2
	dummyB.VelX = bVx2
	dummyB.VelY = bVy2
}

//Distance Calculates the distance between 2 physics dummies and returns the result
func Distance(dummyA Dummy, dummyB Dummy) float64 {
	dx := dummyA.PosX - dummyB.PosX
	dy := dummyA.PosY - dummyB.PosY

	dx2 := dx * dx
	dy2 := dy * dy

	return math.Sqrt(dx2 + dy2)
}
