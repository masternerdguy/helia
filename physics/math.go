package physics

import "math/rand"

//RandInRange Generates a random number in a specified range
func RandInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
