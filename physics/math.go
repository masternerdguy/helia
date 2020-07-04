package physics

import (
	"math"
	"math/rand"
)

//RandInRange Generates a random number in a specified range
func RandInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

//FMod Modulo operator that doesn't retain negative sign
func FMod(a float64, n float64) float64 {
	//mod = (a, n) -> a - floor(a/n) * n
	return a - math.Floor(a/n)*n
}
