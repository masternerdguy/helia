package universe

import "github.com/google/uuid"

// Structure representing a star
type Star struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	PosX     float64
	PosY     float64
	Texture  string
	Radius   float64
	Mass     float64
	Theta    float64
}
