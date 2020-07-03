package universe

import (
	"math"
	"sync"
	"time"

	"helia/physics"

	"github.com/google/uuid"
)

//Ship Structure representing a player ship in the game universe
type Ship struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Created  time.Time
	ShipName string
	PosX     float64
	PosY     float64
	SystemID uuid.UUID
	Texture  string
	Theta    float64
	VelX     float64
	VelY     float64
	Accel    float64
	Radius   float64
	Mass     float64
	Lock     sync.Mutex
}

//PeriodicUpdate Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// update position
	s.PosX += s.VelX * TimeModifier
	s.PosY += s.VelY * TimeModifier

	// clamp theta
	if s.Theta > 360 {
		s.Theta -= 360
	} else if s.Theta < 0 {
		s.Theta += 360
	}
}

//ManualTurn Test function for manual turn
func (s *Ship) ManualTurn(screenT float64, screenM float64) {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// accelerate along new angle for debugging (this whole function needs a redo)
	s.Theta = screenT

	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (s.Accel * TimeModifier)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (s.Accel * TimeModifier)
}

//ToPhysicsDummy Returns a new physics dummy structure representing this ship
func (s *Ship) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: s.VelX,
		VelY: s.VelY,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.Mass,
	}
}

//ApplyPhysicsDummy Applies the values of a physics dummy to this ship
func (s *Ship) ApplyPhysicsDummy(dummy physics.Dummy) {
	s.VelX = dummy.VelX
	s.VelY = dummy.VelY
	s.PosX = dummy.PosX
	s.PosY = dummy.PosY
	s.Mass = dummy.Mass
}
