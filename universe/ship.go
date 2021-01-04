package universe

import (
	"errors"
	"math"
	"sync"
	"time"

	"helia/listener/models"
	"helia/physics"

	"github.com/google/uuid"
)

//ShipFuelTurn Scaler for the amount of fuel used turning
const ShipFuelTurn = 0.001

//ShipHeatTurn Scaler for the amount of heat generated turning
const ShipHeatTurn = 0.003

//ShipFuelBurn Scaler for the amount of fuel used thrusting
const ShipFuelBurn = 0.003

//ShipHeatBurn Scaler for the amount of heat generated thrusting
const ShipHeatBurn = 0.06

//ShipFuelEnergyRegen Scaler for the amount of fuel used regenerating energy
const ShipFuelEnergyRegen = 0.09

//ShipHeatEnergyRegen Scaler for the amount of heat generated regenerating energy
const ShipHeatEnergyRegen = 0.1

//ShipHeatDamage Scaler for damage inflicted by excess heat
const ShipHeatDamage = 0.01

//ShipShieldRegenEnergyBurn Scaler for the amount of energy used regenerating shields
const ShipShieldRegenEnergyBurn = 0.5

//ShipShieldRegenHeat Scaler for the amount of heat generated regenerating shields
const ShipShieldRegenHeat = 1.5

//ShipMinShieldRegenPercent Percentage of shield regen to be applied to a ship at 0% shields
const ShipMinShieldRegenPercent = 0.05

//ShipMinEnergyRegenPercent Percentage of energy regen to be applied to a ship at 100% energy
const ShipMinEnergyRegenPercent = 0.07

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
	ID                    uuid.UUID
	UserID                uuid.UUID
	Created               time.Time
	ShipName              string
	OwnerName             string
	PosX                  float64
	PosY                  float64
	SystemID              uuid.UUID
	Texture               string
	Theta                 float64
	VelX                  float64
	VelY                  float64
	Shield                float64
	Armor                 float64
	Hull                  float64
	Fuel                  float64
	Heat                  float64
	Energy                float64
	Fitting               Fitting
	Destroyed             bool
	DestroyedAt           *time.Time
	CargoBayContainerID   uuid.UUID
	FittingBayContainerID uuid.UUID
	TrashContainerID      uuid.UUID
	ReMaxDirty            bool
	//cache of base template
	TemplateData ShipTemplate
	//docking
	DockedAtStationID *uuid.UUID
	//in-memory only
	IsDocked           bool
	AutopilotMode      int
	AutopilotManualNav ManualNavData
	AutopilotGoto      GotoData
	AutopilotOrbit     OrbitData
	AutopilotDock      DockData
	AutopilotUndock    UndockData
	CurrentSystem      *SolarSystem
	DockedAtStation    *Station
	CargoBay           *Container
	FittingBay         *Container
	Lock               sync.Mutex
}

//Fitting Structure representing the module racks of a ship and what is fitted to them
type Fitting struct {
	ARack []FittedSlot
	BRack []FittedSlot
	CRack []FittedSlot
}

//FittedSlot Structure representing a slot within a ship's fitting rack
type FittedSlot struct {
	ItemTypeID uuid.UUID
	ItemID     uuid.UUID
	//in-memory only, exposable to player
	ItemTypeFamily string
	ItemTypeName   string
	ItemMeta       Meta
	ItemTypeMeta   Meta
	IsCycling      bool
	WillRepeat     bool
	CyclePercent   int
	TargetID       *uuid.UUID
	TargetType     *int
	Rack           string
	//in-memory only, secret
	shipMountedOn    *Ship
	cooldownProgress int
}

//LinkShip Links the ship the slot is on into the slot
func (m *FittedSlot) LinkShip(sp *Ship) {
	m.shipMountedOn = sp
}

//stripModuleFromFitting Removes the fitted slot containing a module from the fitting
func (f *Fitting) stripModuleFromFitting(itemID uuid.UUID) {
	// stub new empty racks
	newA := make([]FittedSlot, 0)
	newB := make([]FittedSlot, 0)
	newC := make([]FittedSlot, 0)

	// copy all except module being unfit
	for _, i := range f.ARack {
		if i.ItemID != itemID {
			newA = append(newA, i)
		}
	}

	for _, i := range f.BRack {
		if i.ItemID != itemID {
			newB = append(newB, i)
		}
	}

	for _, i := range f.CRack {
		if i.ItemID != itemID {
			newC = append(newC, i)
		}
	}

	// replace racks in fitting
	f.ARack = newA
	f.BRack = newB
	f.CRack = newC
}

//CopyShip Returns a copy of the ship
func (s *Ship) CopyShip() *Ship {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	sc := Ship{
		ID:                    s.ID,
		UserID:                s.UserID,
		Created:               s.Created,
		ShipName:              s.ShipName,
		OwnerName:             s.OwnerName,
		PosX:                  s.PosX,
		PosY:                  s.PosY,
		SystemID:              s.SystemID,
		Texture:               s.Texture,
		Theta:                 s.Theta,
		VelX:                  s.VelX,
		VelY:                  s.VelY,
		Shield:                s.Shield,
		Armor:                 s.Armor,
		Hull:                  s.Hull,
		Fuel:                  s.Fuel,
		Heat:                  s.Heat,
		Energy:                s.Energy,
		Fitting:               s.Fitting,
		Destroyed:             s.Destroyed,
		CargoBayContainerID:   s.CargoBayContainerID,
		FittingBayContainerID: s.FittingBayContainerID,
		TrashContainerID:      s.TrashContainerID,
		ReMaxDirty:            s.ReMaxDirty,
		TemplateData: ShipTemplate{
			ID:                 s.TemplateData.ID,
			Created:            s.TemplateData.Created,
			ShipTemplateName:   s.TemplateData.ShipTemplateName,
			Texture:            s.TemplateData.Texture,
			Radius:             s.TemplateData.Radius,
			BaseAccel:          s.TemplateData.BaseAccel,
			BaseMass:           s.TemplateData.BaseMass,
			BaseTurn:           s.TemplateData.BaseTurn,
			BaseShield:         s.TemplateData.BaseShield,
			BaseShieldRegen:    s.TemplateData.BaseShieldRegen,
			BaseArmor:          s.TemplateData.BaseArmor,
			BaseHull:           s.TemplateData.BaseHull,
			BaseFuel:           s.TemplateData.BaseFuel,
			BaseHeatCap:        s.TemplateData.BaseHeatCap,
			BaseHeatSink:       s.TemplateData.BaseHeatSink,
			BaseEnergy:         s.TemplateData.BaseEnergy,
			BaseEnergyRegen:    s.TemplateData.BaseEnergyRegen,
			ShipTypeID:         s.TemplateData.ShipTypeID,
			BaseCargoBayVolume: s.TemplateData.BaseCargoBayVolume,
		},
		//in-memory only
		Lock:               sync.Mutex{},
		IsDocked:           s.IsDocked,
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

	if s.DestroyedAt != nil {
		sc.DestroyedAt = *&s.DestroyedAt
	}

	return &sc
}

//PeriodicUpdate Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// remax some stats if needed for spawning
	if s.ReMaxDirty {
		s.ReMaxStatsForSpawn()
	}

	// update energy
	s.updateEnergy()

	// update shields
	s.updateShield()

	// update heat
	s.updateHeat()

	// check if docked or undocked at a station (docking with other objects not yet supported)
	if s.DockedAtStationID == nil {
		/* Is Undocked */
		s.IsDocked = false

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

		// update modules
		for i := range s.Fitting.ARack {
			s.Fitting.ARack[i].shipMountedOn = s
			s.Fitting.ARack[i].Rack = "A"
			s.Fitting.ARack[i].PeriodicUpdate()
		}

		for i := range s.Fitting.BRack {
			s.Fitting.BRack[i].shipMountedOn = s
			s.Fitting.BRack[i].Rack = "B"
			s.Fitting.BRack[i].PeriodicUpdate()
		}

		for i := range s.Fitting.CRack {
			s.Fitting.CRack[i].shipMountedOn = s
			s.Fitting.CRack[i].Rack = "C"
			s.Fitting.CRack[i].PeriodicUpdate()
		}
	} else {
		/* Is Docked */
		s.IsDocked = true

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

		// reset fuel
		s.Fuel = math.Abs(s.GetRealMaxFuel())

		// check autopilot
		s.doDockedAutopilot()
	}
}

//updateEnergy Updates the ship's energy level for a tick
func (s *Ship) updateEnergy() {
	maxEnergy := s.GetRealMaxEnergy()

	// calculate scaler for energy regen based on current energy percentage
	x := math.Abs(s.Energy / maxEnergy)
	u := math.Pow(ShipMinEnergyRegenPercent, x)

	// get energy regen amount for tick taking energy percentage scaling into account
	tickRegen := ((s.GetRealEnergyRegen() / 1000) * Heartbeat) * u

	if s.Energy < (maxEnergy - tickRegen) {
		// regenerate energy
		energyRegenFuelCost := tickRegen * ShipFuelEnergyRegen

		if s.Fuel-energyRegenFuelCost >= 0 {
			// deduct fuel
			s.Fuel -= energyRegenFuelCost

			// generate heat
			s.Heat += tickRegen * ShipHeatEnergyRegen

			// increase energy level
			s.Energy += tickRegen
		}
	}

	// clamp energy
	if s.Energy < 0 {
		s.Energy = 0
	}

	if s.Energy > maxEnergy {
		s.Energy = maxEnergy
	}

	// clamp fuel
	if s.Fuel < 0 {
		s.Fuel = 0
	}

	maxFuel := s.GetRealMaxFuel()

	if s.Fuel > maxFuel {
		s.Fuel = maxFuel
	}
}

//UpdateShield Updates the ship's shield level for a tick
func (s *Ship) updateShield() {
	// get max shield
	max := s.GetRealMaxShield()

	// calculate scaler for shield regen based on current shield percentage
	x := math.Abs(s.Shield / max)
	u := math.Pow(ShipMinShieldRegenPercent, 1.0-x)

	// get shield regen amount for tick taking shield percentage scaling into account
	tickRegen := ((s.GetRealShieldRegen() / 1000) * Heartbeat) * u

	if s.Shield < (max - tickRegen) {
		// calculate shield regen energy use
		burn := tickRegen * ShipShieldRegenEnergyBurn

		if s.Energy-burn >= 0 {
			// use energy
			s.Energy -= burn

			// generate heat
			s.Heat += tickRegen * ShipShieldRegenHeat

			// regenerate shield
			s.Shield += tickRegen
		}
	}

	// clamp shield
	if s.Shield < 0 {
		s.Shield = 0
	}

	if s.Shield > max {
		s.Shield = max
	}

	// clamp armor
	if s.Armor < 0 {
		s.Armor = 0
	}

	maxArmor := s.GetRealMaxArmor()

	if s.Armor > maxArmor {
		s.Armor = maxArmor
	}

	// clamp hull
	if s.Hull < 0 {
		s.Hull = 0
	}

	maxHull := s.GetRealMaxHull()

	if s.Hull > maxHull {
		s.Hull = maxHull
	}
}

//updateHeat Updates the ship's heat level for a tick
func (s *Ship) updateHeat() {
	// get max heat
	maxHeat := s.GetRealMaxHeat()

	// calculate dissipation efficiency modifier based on heat percentage
	x := s.Heat / maxHeat
	g := math.Cos(1.45 + math.Log(x+0.6))
	u := math.Abs(math.Pow(x, g) + 0.1)

	/*
	 * the above formula will reach an optimal dissipation modifier of ~1.15x (+15% over ship max)
	 * at ~73% capacity then fall off from there until ~800% capacity when it bottoms out at ~0.25x
	 * then strictly increases, exceeding 1x again at ~2500% capacity. if their ship is still alive
	 * at 2500% heat they deserve the >1x strictly increasing modifier :)
	 */

	// check for excess heat
	if s.Heat > maxHeat {
		// damage ship with excess heat
		s.Hull -= (((s.Heat - maxHeat) / 1000) * Heartbeat) * ShipHeatDamage
	}

	// dissipate heat taking efficiency modifier into account
	s.Heat -= ((s.GetRealHeatSink() / 1000) * Heartbeat) * u

	// clamp heat
	if s.Heat < 0 {
		s.Heat = 0
	}
}

//CmdAbort Abruptly ends the current autopilot mode
func (s *Ship) CmdAbort() {
	//stop autopilot
	s.AutopilotMode = NewAutopilotRegistry().None

	//reset autopilot parameters
	s.AutopilotManualNav = ManualNavData{}
	s.AutopilotGoto = GotoData{}
	s.AutopilotOrbit = OrbitData{}
	s.AutopilotDock = DockData{}
	s.AutopilotUndock = UndockData{}
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

//ReMaxStatsForSpawn Resets some stats to their maximum (for use when spawning a new ship)
func (s *Ship) ReMaxStatsForSpawn() {
	if s.ReMaxDirty {
		s.Shield = s.GetRealMaxShield()
		s.Armor = s.GetRealMaxArmor()
		s.Hull = s.GetRealMaxHull()
		s.Fuel = s.GetRealMaxFuel()
		s.Energy = s.GetRealMaxEnergy()
		s.ReMaxDirty = false
	}
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

//GetRealShieldRegen Returns the real shield regen rate after modifiers
func (s *Ship) GetRealShieldRegen() float64 {
	return s.TemplateData.BaseShieldRegen
}

//GetRealMaxArmor Returns the real max armor of the ship after modifiers
func (s *Ship) GetRealMaxArmor() float64 {
	// get base max armor
	a := s.TemplateData.BaseArmor

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		armorMaxAdd, s := e.ItemTypeMeta.GetFloat64("armor_max_add")

		if s {
			// include in real max
			a += armorMaxAdd
		}
	}

	return a
}

//GetRealMaxHull Returns the real max hull of the ship after modifiers
func (s *Ship) GetRealMaxHull() float64 {
	return s.TemplateData.BaseHull
}

//GetRealMaxEnergy Returns the real max energy of the ship after modifiers
func (s *Ship) GetRealMaxEnergy() float64 {
	return s.TemplateData.BaseEnergy
}

//GetRealEnergyRegen Returns the real energy regeneration rate of the ship after modifiers
func (s *Ship) GetRealEnergyRegen() float64 {
	return s.TemplateData.BaseEnergyRegen
}

//GetRealMaxHeat Returns the real max heat of the ship after modifiers
func (s *Ship) GetRealMaxHeat() float64 {
	return s.TemplateData.BaseHeatCap
}

//GetRealHeatSink Returns the real heat dissipation rate of the ship after modifiers
func (s *Ship) GetRealHeatSink() float64 {
	return s.TemplateData.BaseHeatSink
}

//GetRealMaxFuel Returns the real max fuel of the ship after modifiers
func (s *Ship) GetRealMaxFuel() float64 {
	// get base max fuel
	f := s.TemplateData.BaseFuel

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		fuelMaxAdd, s := e.ItemTypeMeta.GetFloat64("fuel_max_add")

		if s {
			// include in real max
			f += fuelMaxAdd
		}
	}

	return f
}

//GetRealCargoBayVolume Returns the real max cargo bay volume of the ship after modifiers
func (s *Ship) GetRealCargoBayVolume() float64 {
	return s.TemplateData.BaseCargoBayVolume
}

//TotalCargoBayVolumeUsed Returns the total amount of cargo bay space currently in use
func (s *Ship) TotalCargoBayVolumeUsed(lock bool) float64 {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// accumulator
	var tV = 0.0

	// loop over items in cargo hold
	for _, i := range s.CargoBay.Items {
		// get their volume metadata
		volume, f := i.Meta.GetFloat64("volume")

		if f {
			tV += (volume * float64(i.Quantity))
		}
	}

	// return total
	return tV
}

//DealDamage Deals damage to the ship
func (s *Ship) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	//apply shield damage
	s.Shield -= shieldDmg

	//clamp shield
	if s.Shield < 0 {
		s.Shield = 0
	}

	//determine shield percentage
	shieldP := s.Shield / s.GetRealMaxShield()

	//apply armor damage if shields below 25% scaling for remaining shields
	if shieldP < 0.25 {
		s.Armor -= armorDmg * (1 - shieldP)
	}

	//clamp armor
	if s.Armor < 0 {
		s.Armor = 0
	}

	//determine armor percentage
	armorP := s.Armor / s.GetRealMaxArmor()

	//apply hull damage if armor below 25% scaling for remaining shield and armor
	if armorP < 0.25 {
		s.Hull -= hullDmg * (1 - armorP) * (1 - shieldP)
	}

	//clamp hull
	if s.Hull < 0 {
		s.Hull = 0

		//todo: handle death of player and respawn in noob ship at nearest station
	}
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

		// abort if docked
		if tgt.IsDocked {
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
	} else if s.AutopilotGoto.Type == targetTypeReg.Star {
		// find star with id
		tgt := s.CurrentSystem.stars[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	} else if s.AutopilotGoto.Type == targetTypeReg.Planet {
		// find planet with id
		tgt := s.CurrentSystem.planets[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	} else if s.AutopilotGoto.Type == targetTypeReg.Jumphole {
		// find jumphole with id
		tgt := s.CurrentSystem.jumpholes[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	} else {
		s.CmdAbort()
		return
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

		// abort if docked
		if tgt.IsDocked {
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
	} else if s.AutopilotOrbit.Type == targetTypeReg.Star {
		// find star with id
		tgt := s.CurrentSystem.stars[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Planet {
		// find planet with id
		tgt := s.CurrentSystem.planets[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Jumphole {
		// find jumphole with id
		tgt := s.CurrentSystem.jumpholes[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort()
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else {
		s.CmdAbort()
		return
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
	} else {
		s.CmdAbort()
		return
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
	// do nothing if out of fuel
	if s.Fuel <= 0 {
		return
	}

	// bound requested turn magnitude
	if scale > 1 {
		scale = 1
	}

	if scale < -1 {
		scale = -1
	}

	// calculate burn
	burn := s.GetRealTurn() * scale

	// turn
	s.Theta += burn

	// expend fuel
	s.Fuel -= math.Abs(burn) * ShipFuelTurn

	// accumulate heat
	s.Heat += math.Abs(burn) * ShipHeatTurn
}

//forwardThrust Fire the ship's thrusters
func (s *Ship) forwardThrust(scale float64) {
	// do nothing if out of fuel
	if s.Fuel <= 0 {
		return
	}

	// bound requested thrust magnitude
	if scale > 1 {
		scale = 1
	}

	if scale < 0 {
		scale = 0
	}

	// calculate burn
	burn := s.GetRealAccel() * scale

	// accelerate along theta using thrust proportional to bounded magnitude
	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (burn)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (burn)

	// consume fuel
	s.Fuel -= math.Abs(burn) * ShipFuelBurn

	// accumulate heat
	s.Heat += math.Abs(burn) * ShipHeatBurn
}

//FindModule Finds a module fitted on this ship
func (s *Ship) FindModule(id uuid.UUID, rack string) *FittedSlot {
	if rack == "A" {
		for i, v := range s.Fitting.ARack {
			if v.ItemID == id {
				return &s.Fitting.ARack[i]
			}
		}
	} else if rack == "B" {
		for i, v := range s.Fitting.BRack {
			if v.ItemID == id {
				return &s.Fitting.BRack[i]
			}
		}
	} else if rack == "C" {
		for i, v := range s.Fitting.CRack {
			if v.ItemID == id {
				return &s.Fitting.CRack[i]
			}
		}
	}

	return nil
}

//FindItemInCargo Returns an item in the ship's cargo bay if it is present
func (s *Ship) FindItemInCargo(id uuid.UUID) *Item {
	//look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		if item.ID == id {
			//return item
			return item
		}
	}

	//nothing found
	return nil
}

//UnfitModule Removes a module from a ship and places it in the cargo hold
func (s *Ship) UnfitModule(m *FittedSlot, lock bool) error {
	if lock {
		//lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//lock containers
	m.shipMountedOn.CargoBay.Lock.Lock()
	defer m.shipMountedOn.CargoBay.Lock.Unlock()

	m.shipMountedOn.FittingBay.Lock.Lock()
	defer m.shipMountedOn.FittingBay.Lock.Unlock()

	//make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("You must be docked to unfit a module")
	}

	//get module volume
	v, _ := m.ItemTypeMeta.GetFloat64("volume")

	//make sure there is sufficient space in the cargo bay
	if s.TotalCargoBayVolumeUsed(lock)+v > s.GetRealCargoBayVolume() {
		return errors.New("Insufficient room in cargo bay to unfit module")
	}

	//make sure the module is not cycling
	if m.IsCycling || m.WillRepeat {
		return errors.New("Modules must be offline to be unfit")
	}

	//if the module is in rack c, make sure the ship is fully repaired
	if m.Rack == "C" {
		if s.Armor < s.GetRealMaxArmor() || s.Hull < s.GetRealMaxHull() {
			return errors.New("Armor and hull must be fully repaired before unfitting modules in rack c")
		}
	}

	//remove from fitting data
	m.shipMountedOn.Fitting.stripModuleFromFitting(m.ItemID)

	//reassign item to cargo bay
	newFB := make([]*Item, 0)

	for i := range m.shipMountedOn.FittingBay.Items {
		o := m.shipMountedOn.FittingBay.Items[i]

		//lock item
		o.Lock.Lock()
		defer o.Lock.Unlock()

		//skip if not this module
		if o.ID != m.ItemID {
			newFB = append(newFB, o)
		} else {
			//move to cargo bay if there is still room
			if s.TotalCargoBayVolumeUsed(lock)+v <= s.GetRealCargoBayVolume() {
				m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, o)
				o.ContainerID = m.shipMountedOn.CargoBayContainerID

				//escalate to core to save to db
				s.CurrentSystem.MovedItems[o.ID.String()] = o
			} else {
				return errors.New("Insufficient room in cargo bay to unfit module")
			}
		}
	}

	s.FittingBay.Items = newFB

	//success!
	return nil
}

//TrashItemInCargo Trashes an item in the ship's cargo bay if it exists
func (s *Ship) TrashItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		//lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	//make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("You must be docked to trash an item")
	}

	//remove from cargo bay
	newCB := make([]*Item, 0)

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		//skip if not this item
		if o.ID != id {
			newCB = append(newCB, o)
		} else {
			//move to trash
			o.ContainerID = s.TrashContainerID

			//escalate to core to save to db
			s.CurrentSystem.MovedItems[o.ID.String()] = o
		}
	}

	s.CargoBay.Items = newCB

	return nil
}

//PackageItemInCargo Packages an item in the ship's cargo bay if it exists
func (s *Ship) PackageItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		//lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	//make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("You must be docked to package an item")
	}

	//get the item to be packaged
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("Item not found in cargo bay")
	}

	//make sure the item is unpackaged
	if item.IsPackaged {
		return errors.New("Item is already packaged")
	}

	//make sure the item is fully repaired
	iHp, f := item.Meta.GetFloat64("hp")
	tHp, g := item.ItemTypeMeta.GetFloat64("hp")

	if f && g {
		if iHp < tHp {
			return errors.New("Item must be fully repaired before packaging")
		}
	}

	//package item in-memory
	item.IsPackaged = true

	//wipe out item metadata
	item.Meta = Meta{}

	//escalate item for packaging in db
	s.CurrentSystem.PackagedItems[item.ID.String()] = item

	return nil
}

//UnpackageItemInCargo Packages an item in the ship's cargo bay if it exists
func (s *Ship) UnpackageItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		//lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	//make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("You must be docked to unpackage an item")
	}

	//get the item to be packaged
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("Item not found in cargo bay")
	}

	//make sure the item is packaged
	if !item.IsPackaged {
		return errors.New("Item is already unpackaged")
	}

	//make sure there is only one in the stack
	if item.Quantity != 1 {
		return errors.New("Must be a stack of 1 to unpackage")
	}

	//unpackage item in-memory
	item.IsPackaged = false

	//copy item type metadata as initial item metadata
	item.Meta = item.ItemTypeMeta

	//escalate item for unpackaging in db
	s.CurrentSystem.UnpackagedItems[item.ID.String()] = item

	return nil
}

//StackItemInCargo Stacks an item in the ship's cargo bay if it exists
func (s *Ship) StackItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		//lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	//get the item to be stacked
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("Item not found in cargo bay")
	}

	//make sure the item is packaged
	if !item.IsPackaged {
		return errors.New("Only packaged items can be stacked")
	}

	//merge stack into next stack
	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		//skip if this item
		if o.ID == id {
			continue
		} else {
			//see if we can merge into this stack
			if o.IsPackaged && o.ItemTypeID == item.ItemTypeID {
				//merge stacks
				q := item.Quantity

				o.Quantity += q
				item.Quantity -= q

				//escalate to core for saving in db
				s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item
				s.CurrentSystem.ChangedQuantityItems[o.ID.String()] = o

				//exit loop
				break
			}
		}
	}

	//remove 0 quantity stacks
	newCB := make([]*Item, 0)

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		//only retain if non empty
		if o.Quantity > 0 {
			newCB = append(newCB, o)
		}
	}

	//update cargo bay
	s.CargoBay.Items = newCB

	return nil
}

//PeriodicUpdate Updates a fitted slot on a ship
func (m *FittedSlot) PeriodicUpdate() {
	if m.IsCycling {
		//update cycle timer
		cooldown, found := m.ItemTypeMeta.GetFloat64("cooldown")

		if !found {
			//module has no cooldown - deactivate
			m.IsCycling = false
			m.WillRepeat = false

			return
		}

		cooldownMs := int(cooldown * 1000)
		m.cooldownProgress += Heartbeat

		if m.cooldownProgress > cooldownMs {
			//cycle completed
			m.IsCycling = false
			m.cooldownProgress = 0
			m.CyclePercent = 0
		} else {
			//update percentage
			m.CyclePercent = ((m.cooldownProgress * 100) / cooldownMs)
		}
	} else {
		//check for activation intent
		if m.WillRepeat {
			//check if a target is required
			needsTarget, found := m.ItemTypeMeta.GetBool("needs_target")

			if needsTarget {
				//check for a target
				if m.TargetID == nil || m.TargetType == nil {
					//no target - can't activate
					m.WillRepeat = false
					return
				}

				//make sure the target actually exists in this solar system
				tgtReg := models.NewTargetTypeRegistry()

				if *m.TargetType == tgtReg.Ship {
					// find ship
					_, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

					if !f {
						//target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Station {
					// find station
					_, f := m.shipMountedOn.CurrentSystem.stations[m.TargetID.String()]

					if !f {
						//target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Planet {
					// find planet
					_, f := m.shipMountedOn.CurrentSystem.planets[m.TargetID.String()]

					if !f {
						//target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Jumphole {
					// find jumphole
					_, f := m.shipMountedOn.CurrentSystem.jumpholes[m.TargetID.String()]

					if !f {
						//target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Star {
					// find star
					_, f := m.shipMountedOn.CurrentSystem.stars[m.TargetID.String()]

					if !f {
						//target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else {
					// unsupported target type - can't activate
					m.TargetID = nil
					m.TargetType = nil

					return
				}
			}

			//check for sufficient activation energy
			activationEnergy, found := m.ItemTypeMeta.GetFloat64("activation_energy")

			if found {
				if m.shipMountedOn.Energy-activationEnergy < 0 {
					//insufficient energy - can't activate
					m.WillRepeat = false
					return
				}
			}

			//to determine whether activation succeeds later
			canActivate := false

			//handle module family effects
			if m.ItemTypeFamily == "gun_turret" {
				canActivate = m.activateAsGunTurret()
			} else if m.ItemTypeFamily == "shield_booster" {
				canActivate = m.activateAsShieldBooster()
			}

			if canActivate {
				//activate module
				m.shipMountedOn.Energy -= activationEnergy
				m.IsCycling = true

				//apply activation heating
				activationHeat, found := m.ItemTypeMeta.GetFloat64("activation_heat")

				if found {
					m.shipMountedOn.Heat += activationHeat
				}
			}
		}
	}
}

func (m *FittedSlot) activateAsGunTurret() bool {
	//safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	//get target
	tgtReg := models.NewTargetTypeRegistry()

	//target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		//find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			//target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else if *m.TargetType == tgtReg.Station {
		// find station
		tgt, f := m.shipMountedOn.CurrentSystem.stations[m.TargetID.String()]

		if !f {
			//target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		//store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		//unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	//check for max range
	modRange, found := m.ItemTypeMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		//verify target is in range
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		if d > modRange {
			//out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	//get damage values
	shieldDmg, _ := m.ItemTypeMeta.GetFloat64("shield_damage")
	armorDmg, _ := m.ItemTypeMeta.GetFloat64("armor_damage")
	hullDmg, _ := m.ItemTypeMeta.GetFloat64("hull_damage")

	//account for falloff if present
	falloff, found := m.ItemTypeMeta.GetString("falloff")

	if found {
		//adjust damage based on falloff style
		if falloff == "linear" {
			//damage dealt is a proportion of the distance to target over max range (closer is higher)
			rangeRatio := 1 - (d / modRange)

			shieldDmg *= rangeRatio
			armorDmg *= rangeRatio
			hullDmg *= rangeRatio
		}
	}

	//apply damage to target
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)
		c.DealDamage(shieldDmg, armorDmg, hullDmg)
	} else if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Station)
		c.DealDamage(shieldDmg, armorDmg, hullDmg)
	}

	//include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		//build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		//push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	//module activates!
	return true
}

func (m *FittedSlot) activateAsShieldBooster() bool {
	//get shield boost amount
	shieldBoost, _ := m.ItemTypeMeta.GetFloat64("shield_boost_amount")

	//apply boost to mounting ship
	m.shipMountedOn.DealDamage(-shieldBoost, 0, 0)

	//include visual effect if present
	activationPGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		tgtReg := models.NewTargetTypeRegistry()

		//build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationPGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
		}

		//push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	//module activates!
	return true
}
