package universe

import (
	"helia/physics"

	"github.com/google/uuid"
)

type Missile struct {
	ID         uuid.UUID
	TargetID   uuid.UUID
	TargetType int
	PosX       float64
	PosY       float64
	Texture    string
	Radius     float64
	Module     *FittedSlot
}

// Returns a new physics dummy structure representing this missile
func (s *Missile) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: 0,
	}
}
