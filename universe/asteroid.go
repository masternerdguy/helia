package universe

import "github.com/google/uuid"

//Asteroid Structure representing an asteroid
type Asteroid struct {
	ID         uuid.UUID
	SystemID   uuid.UUID
	ItemTypeID uuid.UUID
	Name       string
	Texture    string
	Radius     float64
	Theta      float64
	PosX       float64
	PosY       float64
	Yield      float64
	Mass       float64
}
