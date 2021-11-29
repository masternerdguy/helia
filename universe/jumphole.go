package universe

import (
	"fmt"
	"helia/physics"
	"helia/shared"
	"math/rand"
	"time"

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
	Transient    bool
	// in-memory only
	OutSystem   *SolarSystem
	OutJumphole *Jumphole
	Lock        shared.LabeledMutex
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
	s.Lock.Lock("jumphole.CopyJumphole")
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
		Transient:    s.Transient,
		// in-memory only
		Lock: shared.LabeledMutex{
			Structure: "Jumphole",
			UID:       fmt.Sprintf("%v :: %v :: %v", s.ID, time.Now(), rand.Float64()),
		},
		// intentionally not copying pointers
	}
}
