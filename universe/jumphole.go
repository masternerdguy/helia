package universe

import "github.com/google/uuid"

//Jumphole Structure representing a jumphole
type Jumphole struct {
	ID           uuid.UUID
	SystemID     uuid.UUID
	OutSystemID  uuid.UUID
	JumpholeName string
	PosX         float64
	PosY         float64
	Texture      string
	Radius       float64
	Mass         float64
	Theta        float64
}
