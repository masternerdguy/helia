package universe

import (
	"math"
	"sync"
	"time"

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
	Lock     sync.Mutex
}

//PeriodicUpdate Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// update position
	s.PosX += s.VelX
	s.PosY += s.VelY

	// clamp theta
	if s.Theta > 360 {
		s.Theta -= 360
	} else if s.Theta < 0 {
		s.Theta += 360
	}
}

//ManualTurn Test function for manual turn
func (s *Ship) ManualTurn(t float64, m float64) {
	//todo: redo this so that the player won't be able to spam events to turn faster, etc
	const tmpMaxAccel = 10

	// calculate dT
	dT := s.Theta - t

	// accelerate along new angle for debugging
	s.Theta = t

	s.VelX += math.Cos(dT*(math.Pi/180)) * tmpMaxAccel
	s.VelY += math.Sin(dT*(math.Pi/180)) * tmpMaxAccel
}
