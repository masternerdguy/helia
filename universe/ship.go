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
	Dock      int
	Undock    int
}

//NewAutopilotRegistry Returns a populated AutopilotRegistry struct for use as an enum
func NewAutopilotRegistry() *AutopilotRegistry {
	return &AutopilotRegistry{
		None:      0,
		ManualNav: 1,
		Goto:      2,
		Orbit:     3,
		Dock:      4,
		Undock:    5,
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
	Distance float64
}

//DockData Container structure for arguments of the Dock autopilot mode
type DockData struct {
	TargetID uuid.UUID
	Type     int
}

//UndockData Container structure for arguments of the Undock autopilot mode
type UndockData struct {
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
	//docking
	DockedAtStationID *uuid.UUID
	//in-memory only
	AutopilotMode      int
	AutopilotManualNav ManualNavData
	AutopilotGoto      GotoData
	AutopilotOrbit     OrbitData
	AutopilotDock      DockData
	AutopilotUndock    UndockData
	CurrentSystem      *SolarSystem
	DockedAtStation    *Station
	Lock               sync.Mutex
}

//CopyShip Returns a copy of the ship
func (s *Ship) CopyShip() *Ship {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	sc := Ship{
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
		Lock:               sync.Mutex{},
		AutopilotMode:      s.AutopilotMode,
		AutopilotManualNav: s.AutopilotManualNav,
		AutopilotGoto:      s.AutopilotGoto,
		AutopilotOrbit:     s.AutopilotOrbit,
		AutopilotDock:      s.AutopilotDock,
		AutopilotUndock:    s.AutopilotUndock,
	}

	if s.DockedAtStationID != nil {
		sc.DockedAtStationID = *&s.DockedAtStationID
	}

	return &sc
}

//PeriodicUpdate Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// check if docked or undocked at a station (docking with other objects not yet supported)
	if s.DockedAtStationID == nil {
		/* Is Undocked */

		// check autopilot
		s.doUndockedAutopilot()

		// update position
		s.PosX += s.VelX
		s.PosY += s.VelY

		// clamp theta
		if s.Theta > 360 {
			s.Theta -= 360
		} else if s.Theta < 0 {
			s.Theta += 360
		}

		// apply dampening
		dampX := SpaceDrag * s.VelX
		dampY := SpaceDrag * s.VelY

		s.VelX -= dampX
		s.VelY -= dampY
	} else {
		/* Is Docked */

		// validate station pointer
		if s.DockedAtStation == nil {
			// find station
			station := s.CurrentSystem.stations[s.DockedAtStationID.String()]

			// we aren't really docked i guess
			if station == nil {
				s.DockedAtStationID = nil
			} else {
				s.DockedAtStation = station
			}
		}

		// clamp to station
		s.VelX = 0
		s.VelY = 0
		s.PosX = s.DockedAtStation.PosX
		s.PosY = s.DockedAtStation.PosY

		// check autopilot
		s.doDockedAutopilot()
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

//CmdDock Invokes dock autopilot on the ship
func (s *Ship) CmdDock(targetID uuid.UUID, targetType int) {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash dock and activate autopilot
	s.AutopilotDock = DockData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Dock
}

//CmdUndock Invokes undock autopilot on the ship
func (s *Ship) CmdUndock() {
	// get registry
	registry := NewAutopilotRegistry()

	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//stash dock and activate autopilot
	s.AutopilotUndock = UndockData{}

	s.AutopilotMode = registry.Undock
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
	return s.TemplateData.BaseAccel
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

//doUndockedAutopilot Flies the ship for you
func (s *Ship) doUndockedAutopilot() {
	// get registry
	registry := NewAutopilotRegistry()

	switch s.AutopilotMode {
	case registry.None:
		return
	case registry.ManualNav:
		s.doAutopilotManualNav()
	case registry.Goto:
		s.doAutopilotGoto()
	case registry.Orbit:
		s.doAutopilotOrbit()
	case registry.Dock:
		s.doAutopilotDock()
	}
}

//doUndockedAutopilot Flies the ship for you
func (s *Ship) doDockedAutopilot() {
	// get registry
	registry := NewAutopilotRegistry()

	switch s.AutopilotMode {
	case registry.None:
		return
	case registry.Undock:
		s.doAutopilotUndock()
	}
}

//doAutopilotManualNav Causes ship to turn to face a target angle while accelerating
func (s *Ship) doAutopilotManualNav() {
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
	s.AutopilotManualNav.Magnitude -= s.AutopilotManualNav.Magnitude * SpaceDrag

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
	s.flyToPoint(tX, tY, hold, 30)
}

//doAutopilotOrbit Causes ship to fly a circle around the target
func (s *Ship) doAutopilotOrbit() {
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	// target details
	var tX float64 = 0
	var tY float64 = 0

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
	}

	if s.AutopilotOrbit.Distance <= 0 {
		// stash current distance
		s.AutopilotOrbit.Distance = (physics.Distance(s.ToPhysicsDummy(), physics.Dummy{PosX: tX, PosY: tY}))
	}

	// get angle between ship and target
	rX := s.PosX - tX
	rY := s.PosY - tY
	pAngle := physics.ToDegrees(math.Atan2(rY, rX))

	// find point 5 degree ahead
	pAngle += 5
	nX := s.AutopilotOrbit.Distance * math.Cos(physics.ToRadians(pAngle))
	nY := s.AutopilotOrbit.Distance * math.Sin(physics.ToRadians(pAngle))

	// fly to that point
	s.flyToPoint(nX+tX, nY+tY, 0, 3)
}

//doAutopilotDock Causes ship to dock with a target
func (s *Ship) doAutopilotDock() {
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	if s.AutopilotDock.Type == targetTypeReg.Station {
		// find station
		station := s.CurrentSystem.stations[s.AutopilotDock.TargetID.String()]

		if station == nil {
			s.CmdAbort()
			return
		}

		// get distance to station
		d := physics.Distance(s.ToPhysicsDummy(), station.ToPhysicsDummy())
		hold := station.Radius * 0.75

		if d > hold {
			// get closer
			s.flyToPoint(station.PosX, station.PosY, hold, 20)
		} else {
			// dock with station
			s.DockedAtStation = station
			s.DockedAtStationID = &station.ID
			s.AutopilotMode = NewAutopilotRegistry().None
		}
	}
}

//doAutopilotUndock Causes ship to undock from a target
func (s *Ship) doAutopilotUndock() {
	// verify that we are docked (currently only supports stations)
	if s.DockedAtStationID != nil && s.DockedAtStation != nil {
		// remove references
		s.DockedAtStationID = nil
		s.DockedAtStation = nil

		// not docked - cancel autopilot
		s.AutopilotMode = NewAutopilotRegistry().None
	} else {
		// not docked - cancel autopilot
		s.AutopilotMode = NewAutopilotRegistry().None
	}
}

//flyToPoint Reusable function to fly a ship towards a point
func (s *Ship) flyToPoint(tX float64, tY float64, hold float64, caution float64) {
	// face towards target
	turnMag := s.facePoint(tX, tY)

	// determine whether to thrust forward and by how much
	scale := ((s.GetRealAccel() * (caution / SpaceDrag)) / 0.175)
	d := (physics.Distance(s.ToPhysicsDummy(), physics.Dummy{PosX: tX, PosY: tY}) - hold)

	if turnMag < 1 {
		// thrust forward
		s.forwardThrust(d / scale)
	}
}

//facePoint Reusable function to turn a ship towards a point (returns the turn magnitude needed in degrees)
func (s *Ship) facePoint(tX float64, tY float64) float64 {
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

	return turnMag
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
	s.Theta += s.GetRealTurn() * scale
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
	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (s.GetRealAccel() * scale)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (s.GetRealAccel() * scale)
}
