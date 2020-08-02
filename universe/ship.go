package universe

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"helia/physics"

	"github.com/google/uuid"
)

//AutopilotRegistry Autopilot states for ships
type AutopilotRegistry struct {
	None      int
	ManualNav int
	Goto      int
}

//NewAutopilotRegistry Returns a populated AutopilotRegistry struct for use as an enum
func NewAutopilotRegistry() *AutopilotRegistry {
	return &AutopilotRegistry{
		None:      0,
		ManualNav: 1,
		Goto:      2,
	}
}

//ManualNavData Container structure for arguments of the ManualTurn autopilot mode
type ManualNavData struct {
	Magnitude float64
	Theta     float64
}

//GotoData Container structure for arguments of the Goto autopilot mode
type GotoData struct {
	TargetID uuid.UUID
	Type     int
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
	Shield   float64
	Armor    float64
	Hull     float64
	Fuel     float64
	Heat     float64
	Energy   float64
	//cache of base template
	TemplateData ShipTemplate
	//in-memory only
	AutopilotMode      int
	AutopilotManualNav ManualNavData
	AutopilotGoto      GotoData
	Lock               sync.Mutex
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
		Shield:   s.Shield,
		Armor:    s.Armor,
		Hull:     s.Hull,
		Fuel:     s.Fuel,
		Heat:     s.Heat,
		Energy:   s.Energy,
		TemplateData: ShipTemplate{
			ID:               s.TemplateData.ID,
			Created:          s.TemplateData.Created,
			ShipTemplateName: s.TemplateData.ShipTemplateName,
			Texture:          s.TemplateData.Texture,
			Radius:           s.TemplateData.Radius,
			BaseAccel:        s.TemplateData.BaseAccel,
			BaseMass:         s.TemplateData.BaseMass,
			BaseTurn:         s.TemplateData.BaseTurn,
			BaseShield:       s.TemplateData.BaseShield,
			BaseShieldRegen:  s.TemplateData.BaseShieldRegen,
			BaseArmor:        s.TemplateData.BaseArmor,
			BaseHull:         s.TemplateData.BaseHull,
			BaseFuel:         s.TemplateData.BaseFuel,
			BaseHeatCap:      s.TemplateData.BaseHeatCap,
			BaseHeatSink:     s.TemplateData.BaseHeatSink,
			BaseEnergy:       s.TemplateData.BaseEnergy,
			BaseEnergyRegen:  s.TemplateData.BaseEnergyRegen,
			ShipTypeID:       s.TemplateData.ShipTypeID,
		},
		//in-memory only
		AutopilotMode:      s.AutopilotMode,
		AutopilotManualNav: s.AutopilotManualNav,
		AutopilotGoto:      s.AutopilotGoto,
		Lock:               sync.Mutex{},
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
	case registry.ManualNav:
		s.doAutopilotSeekManualNav()
	case registry.Goto:
		s.doAutopilotGoto()
	}
}

//doAutopilotSeekManualNav Causes ship to turn to face a target angle while accelerating
func (s *Ship) doAutopilotSeekManualNav() {
	screenT := s.AutopilotManualNav.Theta

	//calculate magnitude of requested turn
	turnMag := math.Sqrt((screenT - s.Theta) * (screenT - s.Theta))

	a := screenT - s.Theta
	a = physics.FMod(a+180, 360) - 180

	//apply turn with ship limits
	if a > 0 {
		s.rotate(turnMag / s.GetRealTurn())
	} else if a < 0 {
		s.rotate(turnMag / -s.GetRealTurn())
	}

	//thrust forward
	s.forwardThrust(s.AutopilotManualNav.Magnitude)

	//decrease magnitude (this is to allow this to expire and require another move order from the player)
	s.AutopilotManualNav.Magnitude -= s.AutopilotManualNav.Magnitude * SpaceDrag * TimeModifier

	//stop when magnitude is low
	if s.AutopilotManualNav.Magnitude < 0.0001 {
		s.AutopilotMode = NewAutopilotRegistry().None
	}
}

//doAutopilotGoto Causes ship to turn to move towards a target and stop when within range
func (s *Ship) doAutopilotGoto() {
	log.Println(fmt.Sprintf("goto: %v", s.AutopilotGoto))
}

//CmdManualNav Invokes manual turn autopilot on the ship
func (s *Ship) CmdManualNav(screenT float64, screenM float64) {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash manual nav and activate autopilot
	s.AutopilotManualNav = ManualNavData{
		Magnitude: screenM,
		Theta:     screenT,
	}

	s.AutopilotMode = registry.ManualNav
}

//CmdGoto Invokes manual turn autopilot on the ship
func (s *Ship) CmdGoto(targetID uuid.UUID, targetType int) {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash goto and activate autopilot
	s.AutopilotGoto = GotoData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Goto
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
	s.Theta += s.GetRealTurn() * scale * TimeModifier
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
	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (s.GetRealAccel() * scale * TimeModifier)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (s.GetRealAccel() * scale * TimeModifier)
}

//ToPhysicsDummy Returns a new physics dummy structure representing this ship
func (s *Ship) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: s.VelX,
		VelY: s.VelY,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.GetRealMass(),
	}
}

//ApplyPhysicsDummy Applies the values of a physics dummy to this ship
func (s *Ship) ApplyPhysicsDummy(dummy physics.Dummy) {
	s.VelX = dummy.VelX
	s.VelY = dummy.VelY
	s.PosX = dummy.PosX
	s.PosY = dummy.PosY
}

//GetRealAccel Returns the real acceleration capability of a ship after modifiers
func (s *Ship) GetRealAccel() float64 {
	return s.TemplateData.BaseAccel / 10
}

//GetRealTurn Returns the real turning capability of a ship after modifiers
func (s *Ship) GetRealTurn() float64 {
	return s.TemplateData.BaseTurn
}

//GetRealMass Returns the real mass of a ship after modifiers
func (s *Ship) GetRealMass() float64 {
	return s.TemplateData.BaseMass
}

//GetRealMaxShield Returns the real max shield of the ship after modifiers
func (s *Ship) GetRealMaxShield() float64 {
	return s.TemplateData.BaseShield
}

//GetRealMaxArmor Returns the real max armor of the ship after modifiers
func (s *Ship) GetRealMaxArmor() float64 {
	return s.TemplateData.BaseArmor
}

//GetRealMaxHull Returns the real max hull of the ship after modifiers
func (s *Ship) GetRealMaxHull() float64 {
	return s.TemplateData.BaseHull
}

//GetRealMaxEnergy Returns the real max energy of the ship after modifiers
func (s *Ship) GetRealMaxEnergy() float64 {
	return s.TemplateData.BaseEnergy
}

//GetRealMaxHeat Returns the real max heat of the ship after modifiers
func (s *Ship) GetRealMaxHeat() float64 {
	return s.TemplateData.BaseHeatCap
}

//GetRealMaxFuel Returns the real max fuel of the ship after modifiers
func (s *Ship) GetRealMaxFuel() float64 {
	return s.TemplateData.BaseFuel
}
