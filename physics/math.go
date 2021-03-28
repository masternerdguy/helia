package physics

import (
	"math"
	"math/rand"
)

// Generates a random number in a specified range
func RandInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Modulo operator that doesn't retain negative sign
func FMod(a float64, n float64) float64 {
	//mod = (a, n) -> a - floor(a/n) * n
	return a - math.Floor(a/n)*n
}

// Converts degrees to radians
func ToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// Converts radians to degrees
func ToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
