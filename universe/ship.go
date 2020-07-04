package universe

import (
	"math"
	"sync"
	"time"

	"helia/physics"

	"github.com/google/uuid"
)

//AutopilotRegistry Autopilot states for ships
type AutopilotRegistry struct {
	None       int
	ManualTurn int
}

//NewAutopilotRegistry Returns a populated AutopilotRegistry struct for use as an enum
func NewAutopilotRegistry() *AutopilotRegistry {
	return &AutopilotRegistry{
		None:       0,
		ManualTurn: 1,
	}
}

//ManualTurnData Container structure for arguments of the ManualTurn autopilot mode
type ManualTurnData struct {
	Magnitude float64
	Theta     float64
}

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
	Turn     float64
	//in-memory only
	AutopilotMode           int
	AutopilotSeekManualTurn ManualTurnData
	Lock                    sync.Mutex
}

//CopyShip Returns a copy of the ship
func (s *Ship) CopyShip() Ship {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	return Ship{
		ID:       s.ID,
		UserID:   s.UserID,
		Created:  s.Created,
		ShipName: s.ShipName,
		PosX:     s.PosX,
		PosY:     s.PosY,
		SystemID: s.SystemID,
		Texture:  s.Texture,
		Theta:    s.Theta,
		VelX:     s.VelX,
		VelY:     s.VelY,
		Accel:    s.Accel,
		Radius:   s.Radius,
		Mass:     s.Mass,
		Turn:     s.Turn,
		//in-memory only
		AutopilotMode:           s.AutopilotMode,
		AutopilotSeekManualTurn: s.AutopilotSeekManualTurn,
		Lock:                    sync.Mutex{},
	}
}

//PeriodicUpdate Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// check autopilot
	s.doAutopilot()

	// update position
	s.PosX += s.VelX * TimeModifier
	s.PosY += s.VelY * TimeModifier

	// clamp theta
	if s.Theta > 360 {
		s.Theta -= 360
	} else if s.Theta < 0 {
		s.Theta += 360
	}

	// apply dampening
	dampX := SpaceDrag * s.VelX * TimeModifier
	dampY := SpaceDrag * s.VelY * TimeModifier

	s.VelX -= dampX
	s.VelY -= dampY
}

//doAutopilot Flies the ship for you
func (s *Ship) doAutopilot() {
	// get registry
	registry := NewAutopilotRegistry()

	switch s.AutopilotMode {
	case registry.None:
		return
	case registry.ManualTurn:
		s.doAutopilotSeekManualNav()
	}
}

//doAutopilotSeekManualNav Causes ship to turn to face a target angle while accelerating
func (s *Ship) doAutopilotSeekManualNav() {
	screenT := s.AutopilotSeekManualTurn.Theta

	//calculate magnitude of requested turn
	turnMag := math.Sqrt((screenT - s.Theta) * (screenT - s.Theta))

	a := screenT - s.Theta
	a = physics.FMod(a+180, 360) - 180

	//apply turn with ship limits
	if a > 0 {
		s.rotate(turnMag / s.Turn)
	} else if a < 0 {
		s.rotate(turnMag / -s.Turn)
	}

	//thrust forward
	s.forwardThrust(s.AutopilotSeekManualTurn.Magnitude)

	//decrease magnitude (this is to allow this to expire and require another move order from the player)
	s.AutopilotSeekManualTurn.Magnitude -= s.AutopilotSeekManualTurn.Magnitude * SpaceDrag * TimeModifier

	//stop when magnitude is low
	if s.AutopilotSeekManualTurn.Magnitude < 0.0001 {
		s.AutopilotMode = NewAutopilotRegistry().None
	}
}

//ManualTurn Invokes manual turn autopilot on the ship
func (s *Ship) ManualTurn(screenT float64, screenM float64) {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash manual turn and activate autopilot
	s.AutopilotSeekManualTurn = ManualTurnData{
		Magnitude: screenM,
		Theta:     screenT,
	}

	s.AutopilotMode = registry.ManualTurn
}

//rotate Turn the ship
func (s *Ship) rotate(scale float64) {
	// bound requested turn magnitude
	if scale > 1 {
		scale = 1
	}

	if scale < -1 {
		scale = -1
	}

	// turn
	s.Theta += s.Turn * scale * TimeModifier
}

//forwardThrust Fire the ship's thrusters
func (s *Ship) forwardThrust(scale float64) {
	// bound requested thrust magnitude
	if scale > 1 {
		scale = 1
	}

	if scale < 0 {
		scale = 0
	}

	// accelerate along theta using thrust proportional to bounded magnitude
	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (s.Accel * scale * TimeModifier)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (s.Accel * scale * TimeModifier)
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
