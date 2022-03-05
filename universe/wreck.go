package universe

import "github.com/google/uuid"

// Structure representing a wreck
type Wreck struct {
	ID        uuid.UUID
	SystemID  uuid.UUID
	WreckName string
	PosX      float64
	PosY      float64
	Texture   string
	Radius    float64
	Theta     float64
	Items     []*Item
	Ship      *Ship
}
