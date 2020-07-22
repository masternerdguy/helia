package universe

import (
	"helia/physics"

	"github.com/google/uuid"
)

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

//ToPhysicsDummy Returns a new physics dummy structure representing this jumphole
func (s *Jumphole) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.Mass,
	}
}
