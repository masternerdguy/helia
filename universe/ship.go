package universe

import (
	"math"
	"sync"
	"time"

	"helia/listener/models"
	"helia/physics"

	"github.com/google/uuid"
)

//AutopilotRegistry Autopilot states for ships
type AutopilotRegistry struct {
	None      int
	ManualNav int
	Goto      int
	Orbit     int
}

//NewAutopilotRegistry Returns a populated AutopilotRegistry struct for use as an enum
func NewAutopilotRegistry() *AutopilotRegistry {
	return &AutopilotRegistry{
		None:      0,
		ManualNav: 1,
		Goto:      2,
		Orbit:     3,
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

//OrbitData Container structure for arguments of the Orbit autopilot mode
type OrbitData struct {
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
	AutopilotOrbit     OrbitData
	CurrentSystem      *SolarSystem
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
		AutopilotOrbit:     s.AutopilotOrbit,
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
	case registry.Orbit:
		s.doAutopilotOrbit()
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
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	// target details
	var tX float64 = 0
	var tY float64 = 0
	var tR float64 = 0

	// get target
	if s.AutopilotGoto.Type == targetTypeReg.Ship {
		// find ship with id
		tgt := s.CurrentSystem.ships[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.TemplateData.Radius
	} else if s.AutopilotGoto.Type == targetTypeReg.Station {
		// find station with id
		tgt := s.CurrentSystem.stations[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	}

	// fly towards target
	hold := (s.TemplateData.Radius + tR)
	s.flyToPoint(tX, tY, hold)
}

//doAutopilotOrbit Causes ship to fly a circle around the target
func (s *Ship) doAutopilotOrbit() {
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	// target details
	var tX float64 = 0
	var tY float64 = 0
	var tR float64 = 0

	// get target
	if s.AutopilotOrbit.Type == targetTypeReg.Ship {
		// find ship with id
		tgt := s.CurrentSystem.ships[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.TemplateData.Radius
	} else if s.AutopilotOrbit.Type == targetTypeReg.Station {
		// find station with id
		tgt := s.CurrentSystem.stations[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	}

	// fly towards target
	hold := (s.TemplateData.Radius + tR)
	s.flyToPoint(tX, tY, hold)
}

//flyToPoint Reusable function to fly a ship towards a point
func (s *Ship) flyToPoint(tX float64, tY float64, hold float64) {
	// get relative position of target to ship
	rX := s.PosX - tX
	rY := s.PosY - tY

	// get angle between ship and target
	pAngle := -physics.ToDegrees(math.Atan2(rY, rX)) + 180

	// calculate magnitude of requested turn
	turnMag := math.Sqrt((pAngle - s.Theta) * (pAngle - s.Theta))

	a := pAngle - s.Theta
	a = physics.FMod(a+180, 360) - 180

	// apply turn with ship limits
	if a > 0 {
		s.rotate(turnMag / s.GetRealTurn())
	} else if a < 0 {
		s.rotate(turnMag / -s.GetRealTurn())
	}

	scale := (s.GetRealAccel() * (3 / SpaceDrag)) / TimeModifier
	d := (physics.Distance(s.ToPhysicsDummy(), physics.Dummy{PosX: tX, PosY: tY}) - hold)

	if turnMag < 1 && d > hold {
		// thrust forward
		s.forwardThrust(d / scale)
	}
}

//CmdAbort Abruptly ends the current autopilot mode
func (s *Ship) CmdAbort() {
	s.AutopilotMode = NewAutopilotRegistry().None
}

//CmdManualNav Invokes manual nav autopilot on the ship
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

//CmdGoto Invokes goto autopilot on the ship
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

//CmdOrbit Invokes orbit autopilot on the ship
func (s *Ship) CmdOrbit(targetID uuid.UUID, targetType int) {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash orbit and activate autopilot
	s.AutopilotOrbit = OrbitData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Orbit
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
