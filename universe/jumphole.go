package universe

import (
	"helia/physics"
	"sync"

	"github.com/google/uuid"
)

// Structure representing a jumphole
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
	// in-memory only
	OutSystem   *SolarSystem
	OutJumphole *Jumphole
	Lock        sync.Mutex
}

// Returns a new physics dummy structure representing this jumphole
func (s *Jumphole) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.Mass,
	}
}

// Returns a copy of the jumphole
func (s *Jumphole) CopyJumphole() Jumphole {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	return Jumphole{
		ID:           s.ID,
		JumpholeName: s.JumpholeName,
		OutSystemID:  s.OutSystemID,
		PosX:         s.PosX,
		PosY:         s.PosY,
		SystemID:     s.SystemID,
		Texture:      s.Texture,
		Theta:        s.Theta,
		Radius:       s.Radius,
		Mass:         s.Mass,
		// in-memory only
		Lock: sync.Mutex{},
		// intentionally not copying pointers
	}
}
