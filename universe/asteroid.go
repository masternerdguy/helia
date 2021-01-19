package universe

import "github.com/google/uuid"

//Asteroid Structure representing an asteroid
type Asteroid struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	Name     string
	Texture  string
	Radius   float64
	Theta    float64
	PosX     float64
	PosY     float64
	Mass     float64
	// secret, do not expose to player in global update
	ItemTypeID uuid.UUID
	Yield      float64
}
