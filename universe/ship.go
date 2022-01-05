package universe

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"helia/listener/models"
	"helia/physics"
	"helia/shared"

	"github.com/google/uuid"
)

// Scalar for the base cost of warehousing a unit of volume per hour
const WarehouseCostPerHour = 0.25

// Scaler for the amount of fuel used turning
const ShipFuelTurn = 0.001

// Scaler for the amount of heat generated turning
const ShipHeatTurn = 0.003

// Scaler for the amount of fuel used thrusting
const ShipFuelBurn = 0.003

// Scaler for the amount of heat generated thrusting
const ShipHeatBurn = 0.06

// Scaler for the amount of fuel used regenerating energy
const ShipFuelEnergyRegen = 0.09

// Scaler for the amount of heat generated regenerating energy
const ShipHeatEnergyRegen = 0.1

// Scaler for damage inflicted by excess heat
const ShipHeatDamage = 0.01

// Scaler for the amount of energy used regenerating shields
const ShipShieldRegenEnergyBurn = 0.5

// Scaler for the amount of heat generated regenerating shields
const ShipShieldRegenHeat = 1.5

// Percentage of shield regen to be applied to a ship at 0% shields
const ShipMinShieldRegenPercent = 0.05

// Percentage of energy regen to be applied to a ship at 100% energy
const ShipMinEnergyRegenPercent = 0.07

// Autopilot states for ships
type AutopilotRegistry struct {
	None      int
	ManualNav int
	Goto      int
	Orbit     int
	Dock      int
	Undock    int
	Fight     int
}

// Returns a populated AutopilotRegistry struct for use as an enum
func NewAutopilotRegistry() *AutopilotRegistry {
	return &AutopilotRegistry{
		None:      0,
		ManualNav: 1,
		Goto:      2,
		Orbit:     3,
		Dock:      4,
		Undock:    5,
		Fight:     6,
	}
}

// Autopilot states for ships
type BehaviourRegistry struct {
	None       int
	Wander     int
	Patrol     int
	PatchTrade int
}

// Returns a populated AutopilotRegistry struct for use as an enum
func NewBehaviourRegistry() *BehaviourRegistry {
	return &BehaviourRegistry{
		None:       0,
		Wander:     1,
		Patrol:     2,
		PatchTrade: 3,
	}
}

// Container structure for arguments of the ManualTurn autopilot mode
type ManualNavData struct {
	Magnitude float64
	Theta     float64
}

// Container structure for arguments of the Goto autopilot mode
type GotoData struct {
	TargetID uuid.UUID
	Type     int
	Hold     *int
	Caution  *int
}

// Container structure for arguments of the Orbit autopilot mode
type OrbitData struct {
	TargetID uuid.UUID
	Type     int
	Distance float64
}

// DockData Container structure for arguments of the Dock autopilot mode
type DockData struct {
	TargetID uuid.UUID
	Type     int
}

// Container structure for arguments of the Undock autopilot mode
type UndockData struct {
}

// Container structure for arguments of the Goto autopilot mode
type FightData struct {
	TargetID uuid.UUID
	Type     int
}

// Structure representing a player ship in the game universe
type Ship struct {
	ID                       uuid.UUID
	UserID                   uuid.UUID
	Created                  time.Time
	ShipName                 string
	CharacterName            string
	PosX                     float64
	PosY                     float64
	SystemID                 uuid.UUID
	SystemName               string
	Texture                  string
	Theta                    float64
	VelX                     float64
	VelY                     float64
	Shield                   float64
	Armor                    float64
	Hull                     float64
	Fuel                     float64
	Heat                     float64
	Energy                   float64
	Fitting                  Fitting
	Destroyed                bool
	DestroyedAt              *time.Time
	CargoBayContainerID      uuid.UUID
	FittingBayContainerID    uuid.UUID
	TrashContainerID         uuid.UUID
	ReMaxDirty               bool
	Wallet                   float64
	FlightExperienceModifier float64
	// cache of base template
	TemplateData ShipTemplate
	// cache from controlling user
	FactionID uuid.UUID
	// docking
	DockedAtStationID *uuid.UUID
	// in-memory only
	IsNPC              bool
	IsDocked           bool
	Faction            *Faction
	AutopilotMode      int
	AutopilotManualNav ManualNavData
	AutopilotGoto      GotoData
	AutopilotOrbit     OrbitData
	AutopilotDock      DockData
	AutopilotUndock    UndockData
	AutopilotFight     FightData
	BehaviourMode      *int
	CurrentSystem      *SolarSystem
	DockedAtStation    *Station
	CargoBay           *Container
	FittingBay         *Container
	EscrowContainerID  *uuid.UUID
	BeingFlownByPlayer bool
	ReputationSheet    *shared.PlayerReputationSheet
	ExperienceSheet    *shared.PlayerExperienceSheet
	DestructArmed      bool
	TemporaryModifiers []TemporaryShipModifier
	IsCloaked          bool
	Aggressors         map[string]*shared.PlayerReputationSheet
	Lock               shared.LabeledMutex
}

type TemporaryShipModifier struct {
	Attribute      string
	Percentage     float64
	RemainingTicks int
}

// Structure representing a newly purchased ship, not yet materialized
type NewShipPurchase struct {
	ID             uuid.UUID
	ShipTemplateID uuid.UUID
	UserID         uuid.UUID
	StationID      uuid.UUID
	Client         *shared.GameClient
}

// Structure representing a purchased used ship, not yet materialized
type UsedShipPurchase struct {
	ShipID uuid.UUID
	UserID uuid.UUID
	Client *shared.GameClient
}

// Structure representing a renamed ship, not yet materialized
type ShipRename struct {
	ShipID uuid.UUID
	Name   string
}

// Structure representing the owner boarding a different ship, not yet materialized
type ShipSwitch struct {
	Client *shared.GameClient
	Source *Ship
	Target *Ship
}

// Structure representing a change to the no load flag, not yet materialized
type ShipNoLoadSet struct {
	ID   uuid.UUID
	Flag bool
}

// Structure representing the module racks of a ship and what is fitted to them
type Fitting struct {
	ARack []FittedSlot
	BRack []FittedSlot
	CRack []FittedSlot
}

// Structure representing a slot within a ship's fitting rack
type FittedSlot struct {
	ItemTypeID uuid.UUID
	ItemID     uuid.UUID
	// in-memory only, exposable to player
	ItemTypeFamily     string
	ItemTypeFamilyName string
	ItemTypeName       string
	ItemMeta           Meta
	ItemTypeMeta       Meta
	IsCycling          bool
	WillRepeat         bool
	CyclePercent       int
	TargetID           *uuid.UUID
	TargetType         *int
	Rack               string
	SlotIndex          *int
	// in-memory only, secret
	shipMountedOn           *Ship
	cooldownProgress        int
	usageExperienceModifier float64
}

// Links the ship the slot is on into the slot
func (m *FittedSlot) LinkShip(sp *Ship) {
	m.shipMountedOn = sp
}

// stripModuleFromFitting Removes the fitted slot containing a module from the fitting
func (f *Fitting) stripModuleFromFitting(itemID uuid.UUID) {
	// stub new empty racks
	newA := make([]FittedSlot, 0)
	newB := make([]FittedSlot, 0)
	newC := make([]FittedSlot, 0)

	// copy all except module being unfit
	for _, i := range f.ARack {
		if i.ItemID != itemID {
			newA = append(newA, i)
		} else {
			newA = append(newA, FittedSlot{})
		}
	}

	for _, i := range f.BRack {
		if i.ItemID != itemID {
			newB = append(newB, i)
		} else {
			newB = append(newB, FittedSlot{})
		}
	}

	for _, i := range f.CRack {
		if i.ItemID != itemID {
			newC = append(newC, i)
		} else {
			newC = append(newC, FittedSlot{})
		}
	}

	// replace racks in fitting
	f.ARack = newA
	f.BRack = newB
	f.CRack = newC
}

// Determines whether a rack has a free slot suitable for fitting a given module and returns the index if found
func (s *Ship) getFreeSlotIndex(itemFamilyID string, volume float64, rack string) (int, bool) {
	// get default uuid
	emptyUUID := uuid.UUID{}
	defaultUUID := emptyUUID.String()

	// convert item family to module family
	var modFamily string = ""

	if itemFamilyID == "gun_turret" {
		modFamily = "gun"
	} else if itemFamilyID == "missile_launcher" {
		modFamily = "missile"
	} else if itemFamilyID == "shield_booster" ||
		itemFamilyID == "armor_plate" {
		modFamily = "tank"
	} else if itemFamilyID == "fuel_tank" {
		modFamily = "fuel"
	} else if itemFamilyID == "eng_oc" {
		modFamily = "engine"
	} else if itemFamilyID == "active_sink" {
		modFamily = "heat"
	} else if itemFamilyID == "heat_sink" {
		modFamily = "heat"
	} else if itemFamilyID == "drag_amp" {
		modFamily = "tackle"
	} else if itemFamilyID == "utility_miner" {
		modFamily = "utility"
	} else if itemFamilyID == "utility_siphon" {
		modFamily = "utility"
	} else if itemFamilyID == "utility_cloak" {
		modFamily = "utility"
	} else if itemFamilyID == "battery_pack" {
		modFamily = "power"
	} else if itemFamilyID == "aux_generator" {
		modFamily = "power"
	} else if itemFamilyID == "cargo_expander" {
		modFamily = "cargo"
	}

	if modFamily == "" {
		return -1, false
	}

	// find rack in fitting
	var realRack *[]FittedSlot = nil

	if rack == "a" {
		realRack = &s.Fitting.ARack
	} else if rack == "b" {
		realRack = &s.Fitting.BRack
	} else if rack == "c" {
		realRack = &s.Fitting.CRack
	}

	if realRack == nil {
		return -1, false
	}

	// find rack in layout
	var layoutRack *[]Slot = nil

	if rack == "a" {
		layoutRack = &s.TemplateData.SlotLayout.ASlots
	} else if rack == "b" {
		layoutRack = &s.TemplateData.SlotLayout.BSlots
	} else if rack == "c" {
		layoutRack = &s.TemplateData.SlotLayout.CSlots
	}

	if layoutRack == nil {
		return -1, false
	}

	// iterate over slot layout
	for i, s := range *layoutRack {
		// check if this slot is used in the fitting
		if len(*realRack) > i {
			f := (*realRack)[i]

			if f.ItemID.String() != defaultUUID {
				// this slot is in use
				continue
			}
		}

		// check if this slot is compatible
		if float64(s.Volume) >= volume &&
			(s.Family == modFamily || s.Family == "any") {
			// this slot is compatible
			return i, true
		}
	}

	return -1, false
}

// Returns a copy of the ship
func (s *Ship) CopyShip(lock bool) *Ship {
	if lock {
		s.Lock.Lock("ship.CopyShip")
		defer s.Lock.Unlock()
	}

	sc := Ship{
		ID:                       s.ID,
		UserID:                   s.UserID,
		Created:                  s.Created,
		ShipName:                 s.ShipName,
		CharacterName:            s.CharacterName,
		PosX:                     s.PosX,
		PosY:                     s.PosY,
		SystemID:                 s.SystemID,
		SystemName:               s.SystemName,
		Texture:                  s.Texture,
		Theta:                    s.Theta,
		VelX:                     s.VelX,
		VelY:                     s.VelY,
		Shield:                   s.Shield,
		Armor:                    s.Armor,
		Hull:                     s.Hull,
		Fuel:                     s.Fuel,
		Heat:                     s.Heat,
		Energy:                   s.Energy,
		Fitting:                  s.Fitting,
		Destroyed:                s.Destroyed,
		CargoBayContainerID:      s.CargoBayContainerID,
		FittingBayContainerID:    s.FittingBayContainerID,
		TrashContainerID:         s.TrashContainerID,
		ReMaxDirty:               s.ReMaxDirty,
		Wallet:                   s.Wallet,
		FlightExperienceModifier: s.FlightExperienceModifier,
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
			ItemTypeID:         s.TemplateData.ItemTypeID,
			CanUndock:          s.TemplateData.CanUndock,
		},
		FactionID: s.FactionID,
		// in-memory only
		Lock: shared.LabeledMutex{
			Structure: "Ship",
			UID:       fmt.Sprintf("%v :: %v :: %v", s.ID, time.Now(), rand.Float64()),
		},
		IsDocked:           s.IsDocked,
		AutopilotMode:      s.AutopilotMode,
		AutopilotManualNav: s.AutopilotManualNav,
		AutopilotGoto:      s.AutopilotGoto,
		AutopilotOrbit:     s.AutopilotOrbit,
		AutopilotDock:      s.AutopilotDock,
		AutopilotUndock:    s.AutopilotUndock,
		AutopilotFight:     s.AutopilotFight,
		EscrowContainerID:  s.EscrowContainerID,
		BeingFlownByPlayer: s.BeingFlownByPlayer,
		BehaviourMode:      s.BehaviourMode,
		IsNPC:              s.IsNPC,
		TemporaryModifiers: s.TemporaryModifiers,
		IsCloaked:          s.IsCloaked,
	}

	// copy station if docked
	if s.DockedAtStationID != nil {
		g := &s.DockedAtStationID
		sc.DockedAtStationID = *g

		if s.DockedAtStation != nil {
			cs := s.DockedAtStation.CopyStation()
			sc.DockedAtStation = &cs
		}
	}

	// copy destroyed at if set
	if s.DestroyedAt != nil {
		g := &s.DestroyedAt
		sc.DestroyedAt = *g
	}

	return &sc
}

// Processes the ship for a tick
func (s *Ship) PeriodicUpdate() {
	// lock entity
	s.Lock.Lock("ship.PeriodicUpdate")
	defer s.Lock.Unlock()

	// update experience modifier
	if s.ExperienceSheet != nil {
		s.FlightExperienceModifier = s.GetExperienceModifier()
	}

	// cache system name
	if s.CurrentSystem != nil {
		s.SystemName = s.CurrentSystem.SystemName
	}

	// remax some stats if needed for spawning
	if s.ReMaxDirty {
		s.ReMaxStatsForSpawn()
	}

	// special cheats NPCs get to make things easier to implement for now (i do want to eliminate these)
	if s.IsNPC {
		if s.Fuel <= 0 {
			// infinite fuel via refill at 0
			s.Fuel = s.GetRealMaxFuel()
		}
	}

	// handle temporary modifiers
	keptTemporaryModifiers := make([]TemporaryShipModifier, len(s.TemporaryModifiers))

	for _, e := range s.TemporaryModifiers {
		// eliminate if expired
		if e.RemainingTicks <= 0 {
			continue
		}

		// tick down
		e.RemainingTicks--

		// store
		keptTemporaryModifiers = append(keptTemporaryModifiers, e)
	}

	s.TemporaryModifiers = keptTemporaryModifiers

	// update cloaking
	s.updateCloaking()

	// update energy
	s.updateEnergy()

	// update shields
	s.updateShield()

	// update heat
	s.updateHeat()

	// run behaviour routine if applicable
	s.behave()

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

		// calculate felt drag
		drag := s.GetRealSpaceDrag()

		dampX := drag * s.VelX
		dampY := drag * s.VelY

		// apply dampening
		if math.Abs(dampX*(1+Epsilon)) >= math.Abs(s.VelX) {
			s.VelX = 0
		} else {
			s.VelX -= dampX
		}

		if math.Abs(dampY*(1+Epsilon)) >= math.Abs(s.VelY) {
			s.VelY = 0
		} else {
			s.VelY -= dampY
		}

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

		// disarm self destruct
		s.DestructArmed = false

		// make sure ship is still linked into frozen modules
		for i := range s.Fitting.ARack {
			s.Fitting.ARack[i].shipMountedOn = s
			s.Fitting.ARack[i].Rack = "A"
		}

		for i := range s.Fitting.BRack {
			s.Fitting.BRack[i].shipMountedOn = s
			s.Fitting.BRack[i].Rack = "B"
		}

		for i := range s.Fitting.CRack {
			s.Fitting.CRack[i].shipMountedOn = s
			s.Fitting.CRack[i].Rack = "C"
		}

		// clamp to station
		s.VelX = 0
		s.VelY = 0
		s.PosX = s.DockedAtStation.PosX
		s.PosY = s.DockedAtStation.PosY

		// handle station workshop / warehouse
		if !s.TemplateData.CanUndock {
			/*
			 * these are a special type of ship that allows a player to store and work with a large volume of items
			 * in a station for a fee. this fee is deducted from the warehouse "ship"'s wallet and can run into the
			 * negative. in order to re-board this ship to work with or retrieve the items, any defecit must be
			 * settled with a cash transfer.
			 *
			 * the hourly fee is gradually assessed every tick.
			 *
			 * if there is nothing in in the warehouse's cargo bay, no fee is assesed.
			 */

			// get total volume stored
			warehousedVolume := s.TotalCargoBayVolumeUsed(false)

			// get cost per hour per unit
			baseHourlyCost := WarehouseCostPerHour

			// get fee to assess this tick
			costPerHour := warehousedVolume * baseHourlyCost
			secondsPerTick := float64(Heartbeat) / 1000.0

			costPerTick := (secondsPerTick * costPerHour) / 3600

			// deduct from wallet
			s.Wallet -= costPerTick
		} else {
			// check autopilot
			s.doDockedAutopilot()
		}
	}
}

func (s *Ship) updateCloaking() {
	// aggregate cloaking percentage
	cloaked := false
	cloakPercentage := 0.0

	for _, e := range s.TemporaryModifiers {
		if e.Attribute == "cloak" {
			cloakPercentage += e.Percentage
		}
	}

	// determine whether cloaked for tick
	if cloakPercentage >= 1 {
		// ship is cloaked
		cloaked = true
	} else {
		// ship is intermittently cloaked
		r := rand.Float64()

		if r <= cloakPercentage {
			cloaked = true
		}
	}

	s.IsCloaked = cloaked
}

// Updates the ship's energy level for a tick
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

// Updates the ship's shield level for a tick
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

// Updates the ship's heat level for a tick
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

	if !s.IsCloaked {
		// dissipate heat taking efficiency modifier into account
		s.Heat -= ((s.GetRealHeatSink() / 1000) * Heartbeat) * u
	}

	// clamp heat
	if s.Heat < 0 {
		s.Heat = 0
	}
}

// Abruptly ends the current autopilot mode
func (s *Ship) CmdAbort(lock bool) {
	if lock {
		// obtain lock
		s.Lock.Lock("ship.CmdAbort")
		defer s.Lock.Unlock()
	}

	// stop autopilot
	s.AutopilotMode = NewAutopilotRegistry().None

	// reset autopilot parameters
	s.AutopilotManualNav = ManualNavData{}
	s.AutopilotGoto = GotoData{}
	s.AutopilotOrbit = OrbitData{}
	s.AutopilotDock = DockData{}
	s.AutopilotUndock = UndockData{}
	s.AutopilotFight = FightData{}

	for i := range s.Fitting.ARack {
		s.Fitting.ARack[i].WillRepeat = false
		s.Fitting.ARack[i].TargetID = nil
		s.Fitting.ARack[i].TargetType = nil
	}

	for i := range s.Fitting.BRack {
		s.Fitting.BRack[i].WillRepeat = false
		s.Fitting.BRack[i].TargetID = nil
		s.Fitting.BRack[i].TargetType = nil
	}

	for i := range s.Fitting.CRack {
		s.Fitting.CRack[i].WillRepeat = false
	}
}

// Invokes manual nav autopilot on the ship
func (s *Ship) CmdManualNav(screenT float64, screenM float64, lock bool) {
	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdManualNav")
		defer s.Lock.Unlock()
	}

	// stash manual nav and activate autopilot
	s.AutopilotManualNav = ManualNavData{
		Magnitude: screenM,
		Theta:     screenT,
	}

	s.AutopilotMode = registry.ManualNav
}

// Invokes goto autopilot on the ship
func (s *Ship) CmdGoto(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdGoto")
		defer s.Lock.Unlock()
	}

	// stash goto and activate autopilot
	s.AutopilotGoto = GotoData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Goto
}

// Invokes orbit autopilot on the ship
func (s *Ship) CmdOrbit(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdOrbit")
		defer s.Lock.Unlock()
	}

	// stash orbit and activate autopilot
	s.AutopilotOrbit = OrbitData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Orbit
}

// Invokes dock autopilot on the ship
func (s *Ship) CmdDock(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdDock")
		defer s.Lock.Unlock()
	}

	// stash dock and activate autopilot
	s.AutopilotDock = DockData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Dock
}

// Invokes undock autopilot on the ship
func (s *Ship) CmdUndock(lock bool) {
	// ignore if incapable of undocking
	if !s.TemplateData.CanUndock {
		return
	}

	// make sure cargo isn't overloaded
	usedBay := s.TotalCargoBayVolumeUsed(lock)
	maxBay := s.GetRealCargoBayVolume()

	if usedBay > maxBay {
		return
	}

	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdUndock")
		defer s.Lock.Unlock()
	}

	// stash dock and activate autopilot
	s.AutopilotUndock = UndockData{}

	s.AutopilotMode = registry.Undock
}

// Invokes fight autopilot on the ship
func (s *Ship) CmdFight(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := NewAutopilotRegistry()

	if lock {
		// lock entity
		s.Lock.Lock("ship.CmdFight")
		defer s.Lock.Unlock()
	}

	// stash fight and activate autopilot
	s.AutopilotFight = FightData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Fight
}

// Returns a new physics dummy structure representing this ship
func (s *Ship) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: s.VelX,
		VelY: s.VelY,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.GetRealMass(),
	}
}

// Applies the values of a physics dummy to this ship
func (s *Ship) ApplyPhysicsDummy(dummy physics.Dummy) {
	s.VelX = dummy.VelX
	s.VelY = dummy.VelY
	s.PosX = dummy.PosX
	s.PosY = dummy.PosY
}

// Resets some stats to their maximum (for use when spawning a new ship)
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

// Returns the real acceleration capability of a ship after modifiers
func (s *Ship) GetRealAccel() float64 {
	// temporary modifier percentage accumulator
	tpm := 1.0

	// apply temporary modifiers
	for _, e := range s.TemporaryModifiers {
		if e.Attribute == "accel" {
			tpm += e.Percentage
		}
	}

	// floor percentage modifier at 0
	if tpm < 0 {
		tpm = 0
	}

	// return true acceleration
	return s.TemplateData.BaseAccel * s.FlightExperienceModifier * tpm
}

// Returns the real drag felt by a ship after modifiers
func (s *Ship) GetRealSpaceDrag() float64 {
	// temporary modifier percentage accumulator
	tpm := 1.0

	// apply temporary modifiers
	for _, e := range s.TemporaryModifiers {
		if e.Attribute == "drag" {
			tpm += e.Percentage
		}
	}

	// floor percentage modifier at 0
	if tpm < 0 {
		tpm = 0
	}

	// return true drag
	return SpaceDrag * tpm
}

// Calculates the experience percentage bonus to apply to some basic stats
func (s *Ship) GetExperienceModifier() float64 {
	m := 1.0

	if s.ExperienceSheet != nil {
		// get experience entry for this ship template
		v := s.ExperienceSheet.GetShipExperienceEntry(s.TemplateData.ID)

		// get truncated level
		l := math.Trunc(v.GetExperience())

		if l > 0 {
			// apply a dampening factor to get percentage
			b := math.Log(((math.Pow(l, 0.9)) / 9) + 1)

			if b > 0 {
				// add bonus
				m += b
			}
		}
	}

	return m
}

// Returns the real turning capability of a ship after modifiers
func (s *Ship) GetRealTurn() float64 {
	return s.TemplateData.BaseTurn * s.FlightExperienceModifier
}

// Returns the real mass of a ship after modifiers
func (s *Ship) GetRealMass() float64 {
	return s.TemplateData.BaseMass
}

// Returns the real max shield of the ship after modifiers
func (s *Ship) GetRealMaxShield() float64 {
	return s.TemplateData.BaseShield * s.FlightExperienceModifier
}

// Returns the real shield regen rate after modifiers
func (s *Ship) GetRealShieldRegen() float64 {
	return s.TemplateData.BaseShieldRegen * s.FlightExperienceModifier
}

// Returns the real max armor of the ship after modifiers
func (s *Ship) GetRealMaxArmor() float64 {
	// get base max armor
	a := s.TemplateData.BaseArmor * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		armorMaxAdd, s := e.ItemMeta.GetFloat64("armor_max_add")

		if s {
			// include in real max
			a += armorMaxAdd
		}
	}

	return a
}

// Returns the real max hull of the ship after modifiers
func (s *Ship) GetRealMaxHull() float64 {
	return s.TemplateData.BaseHull * s.FlightExperienceModifier
}

// Returns the real max energy of the ship after modifiers
func (s *Ship) GetRealMaxEnergy() float64 {
	// get base max energy
	a := s.TemplateData.BaseEnergy * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		energyMaxAdd, s := e.ItemMeta.GetFloat64("energy_max_add")

		if s {
			// include in real max
			a += energyMaxAdd
		}
	}

	return a
}

// Returns the real energy regeneration rate of the ship after modifiers
func (s *Ship) GetRealEnergyRegen() float64 {
	// get base energy regen
	a := s.TemplateData.BaseEnergyRegen * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		energyRegenAdd, s := e.ItemMeta.GetFloat64("energy_regen_max_add")

		if s {
			// include in real max
			a += energyRegenAdd
		}
	}

	return a
}

// Returns the real max heat of the ship after modifiers
func (s *Ship) GetRealMaxHeat() float64 {
	return s.TemplateData.BaseHeatCap * s.FlightExperienceModifier
}

// Returns the real heat dissipation rate of the ship after modifiers
func (s *Ship) GetRealHeatSink() float64 {
	// get base heat sink
	a := s.TemplateData.BaseHeatSink * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		heatSinkAdd, s := e.ItemMeta.GetFloat64("heat_sink_add")

		if s {
			// include in real max
			a += heatSinkAdd
		}
	}

	// temporary modifier percentage accumulator
	tpm := 1.0

	// apply temporary modifiers
	for _, e := range s.TemporaryModifiers {
		if e.Attribute == "heat_sink" {
			tpm += e.Percentage
		}
	}

	// floor percentage modifier at 0
	if tpm < 0 {
		tpm = 0
	}

	return a * tpm
}

// Returns the real max fuel of the ship after modifiers
func (s *Ship) GetRealMaxFuel() float64 {
	// get base max fuel
	f := s.TemplateData.BaseFuel * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		fuelMaxAdd, s := e.ItemMeta.GetFloat64("fuel_max_add")

		if s {
			// include in real max
			f += fuelMaxAdd
		}
	}

	return f
}

// Returns the real max cargo bay volume of the ship after modifiers
func (s *Ship) GetRealCargoBayVolume() float64 {
	// get base cargo volume
	cv := s.TemplateData.BaseCargoBayVolume * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		if e.SlotIndex == nil {
			continue
		}

		// check if cargo expander
		k, s := e.ItemMeta.GetBool("cargokit")

		if k && s {
			// get slot volume
			l := e.shipMountedOn.TemplateData.SlotLayout.CSlots[*e.SlotIndex]
			sv := float64(l.Volume)

			// include in real max
			cv += sv
		}
	}

	return cv
}

// Returns the total amount of cargo bay space currently in use
func (s *Ship) TotalCargoBayVolumeUsed(lock bool) float64 {
	if lock {
		// lock entity
		s.Lock.Lock("ship.TotalCargoBayVolumeUsed")
		defer s.Lock.Unlock()
	}

	// accumulator
	var tV = 0.0

	// loop over items in cargo hold
	for _, i := range s.CargoBay.Items {
		if i.IsPackaged {
			// get item type volume metadata
			volume, f := i.ItemTypeMeta.GetFloat64("volume")

			if f {
				tV += (volume * float64(i.Quantity))
			}
		} else {
			// get item volume metadata
			volume, f := i.Meta.GetFloat64("volume")

			if f {
				tV += (volume * float64(i.Quantity))
			}
		}
	}

	// return total
	return tV
}

// Deals damage to the ship
func (s *Ship) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64, attackerRS *shared.PlayerReputationSheet) {
	// update aggression table
	if attackerRS != nil {
		// obtain lock
		attackerRS.Lock.Lock("ship.DealDamage")
		defer attackerRS.Lock.Unlock()

		// get attacking player's reputation sheet entry for this ship's faction
		f, ok := attackerRS.FactionEntries[s.FactionID.String()]

		if !ok {
			// does not exist - create a neutral one
			ne := shared.PlayerReputationSheetFactionEntry{
				FactionID:        s.FactionID,
				StandingValue:    0,
				AreOpenlyHostile: false,
			}

			attackerRS.FactionEntries[s.FactionID.String()] = &ne
			f = attackerRS.FactionEntries[s.FactionID.String()]
		}

		// update temporary hostility due to aggro flag
		at := time.Now().Add(15 * time.Minute)
		f.TemporarilyOpenlyHostileUntil = &at

		// store entry
		s.Aggressors[attackerRS.UserID.String()] = attackerRS
	}

	// apply shield damage
	s.Shield -= shieldDmg

	// clamp shield
	if s.Shield < 0 {
		s.Shield = 0
	}

	// determine shield percentage
	shieldP := s.Shield / s.GetRealMaxShield()

	// apply armor damage if shields below 25% scaling for remaining shields
	if shieldP < 0.25 {
		s.Armor -= armorDmg * (1 - shieldP)
	}

	// clamp armor
	if s.Armor < 0 {
		s.Armor = 0
	}

	// determine armor percentage
	armorP := s.Armor / s.GetRealMaxArmor()

	// apply hull damage if armor below 25% scaling for remaining shield and armor
	if armorP < 0.25 {
		s.Hull -= hullDmg * (1 - armorP) * (1 - shieldP)
	}

	// clamp hull
	if s.Hull < 0 {
		s.Hull = 0
	}
}

// Siphons energy from the ship, returns the actual amount siphoned
func (s *Ship) SiphonEnergy(maxSiphonAmount float64, attackerRS *shared.PlayerReputationSheet) float64 {
	// update aggression table
	if attackerRS != nil {
		// obtain lock
		attackerRS.Lock.Lock("ship.SiphonEnergy")
		defer attackerRS.Lock.Unlock()

		// get attacking player's reputation sheet entry for this ship's faction
		f, ok := attackerRS.FactionEntries[s.FactionID.String()]

		if !ok {
			// does not exist - create a neutral one
			ne := shared.PlayerReputationSheetFactionEntry{
				FactionID:        s.FactionID,
				StandingValue:    0,
				AreOpenlyHostile: false,
			}

			attackerRS.FactionEntries[s.FactionID.String()] = &ne
			f = attackerRS.FactionEntries[s.FactionID.String()]
		}

		// update temporary hostility due to aggro flag
		at := time.Now().Add(15 * time.Minute)
		f.TemporarilyOpenlyHostileUntil = &at

		// store entry
		s.Aggressors[attackerRS.UserID.String()] = attackerRS
	}

	// limit amount to siphon so that the remaining amount is positive
	actualSiphon := 0.0

	if s.Energy-maxSiphonAmount >= 0 {
		actualSiphon = maxSiphonAmount
	} else {
		actualSiphon = maxSiphonAmount - s.Energy
	}

	// apply siphon
	s.Energy -= actualSiphon

	// clamp energy
	if s.Energy < 0 {
		s.Energy = 0
	}

	// return amount siphoned
	return actualSiphon
}

// Given a faction to compare against, returns the standing and whether they have declared open hostilities
func (s *Ship) CheckStandings(factionID uuid.UUID) (float64, bool) {
	// handle NPC case first
	if s.IsNPC {
		// null check
		if s.Faction != nil {
			// redirect to faction
			return s.Faction.CheckStandings(factionID)
		}

		// open season :)
		return shared.MIN_STANDING, true
	}

	// null check for human player case
	if s.ReputationSheet == nil {
		// open season :)
		return shared.MIN_STANDING, true
	}

	// obtain lock (pointer to player's entry on their game client)
	s.ReputationSheet.Lock.Lock("ship.CheckStandings")
	defer s.ReputationSheet.Lock.Unlock()

	// try to find faction relationship
	if val, ok := s.ReputationSheet.FactionEntries[factionID.String()]; ok {
		// determine if currently openly hostile
		oh := false

		if val.AreOpenlyHostile {
			// always openly hostile
			oh = true
		} else {
			if val.TemporarilyOpenlyHostileUntil != nil {
				ht := *val.TemporarilyOpenlyHostileUntil

				if time.Now().Before(ht) {
					// temporarily hostile until 15 minutes has elapsed
					oh = true
				}
			}
		}

		// return standings and hostility status
		return val.StandingValue, oh
	} else {
		return 0, false
	}
}

// Run routine to control ship using autopilot commands to achieve a goal
func (s *Ship) behave() {
	// NPC-only routine
	if !s.IsNPC {
		return
	}

	if s.BehaviourMode != nil {
		// get aggressive if under threat
		ms := s.GetRealMaxShield()
		sr := s.Shield / ms

		if sr < 0.75 {
			s.behaviourPatrol()
		} else {
			// get registry
			registry := NewBehaviourRegistry()

			switch *s.BehaviourMode {
			case registry.None:
				// unset mode
				s.BehaviourMode = nil
				return
			case registry.Wander:
				s.behaviourWander()
			case registry.Patrol:
				s.behaviourPatrol()
			case registry.PatchTrade:
				s.behaviourPatchTrade()
			}
		}
	}
}

// wanders around the universe aimlessly
func (s *Ship) behaviourWander() {
	// pause if heat too high
	maxHeat := s.GetRealMaxHeat()
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := NewAutopilotRegistry()

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil {
			// 1% chance of undocking per tick
			roll := physics.RandInRange(0, 100)

			if roll == 1 {
				// undock
				s.CmdUndock(false)
				return
			}
		} else {
			s.gotoNextWanderDestination(50)
		}
	}
}

// wanders around the universe looking for hostiles to attack
func (s *Ship) behaviourPatrol() {
	// get registry
	autoReg := NewAutopilotRegistry()

	// get heat level
	maxHeat := s.GetRealMaxHeat()
	heatLevel := s.Heat / maxHeat

	// don't worry about heat if actually fighting
	if s.AutopilotMode != autoReg.Fight {
		// pause if heat too high
		if heatLevel > 0.95 {
			s.CmdAbort(false)
		}
	}

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil {
			// 1% chance of undocking per tick
			roll := physics.RandInRange(0, 100)

			if roll == 1 {
				// undock
				s.CmdUndock(false)
				return
			}
		} else {
			s.gotoNextWanderDestination(50)
		}
	} else if s.AutopilotMode != autoReg.Fight {
		// look for any hostile ships to attack
		if s.CurrentSystem.tickCounter%16 == 0 {
			// get registry
			tgtReg := models.NewTargetTypeRegistry()

			// scan ships in system
			var tgtS *Ship = nil
			lowestStanding := math.MaxFloat64
			lowestDistance := math.MaxFloat64

			for _, sx := range s.CurrentSystem.ships {
				// skip if docked
				if sx.DockedAtStation != nil {
					continue
				}

				// get distance
				sA := s.ToPhysicsDummy()
				sB := s.ToPhysicsDummy()
				distance := physics.Distance(sA, sB)

				// check standings
				standing, openlyHostile := sx.CheckStandings(s.FactionID)

				// skip if self
				if sx.ID == s.ID {
					continue
				}

				// only attack openly hostile targets
				if openlyHostile {
					// prioritize targets by standing then distance
					if standing <= lowestStanding {
						tgtS = sx
						lowestStanding = standing
					} else if distance <= lowestDistance {
						tgtS = sx
						lowestDistance = distance
					}
				}
			}

			// issue attack order if found
			if tgtS != nil {
				s.CmdFight(tgtS.ID, tgtReg.Ship, false)
			}
		}
	}
}

// wanders around the universe randomly buying and selling things to patch the economy
func (s *Ship) behaviourPatchTrade() {
	// pause if heat too high
	maxHeat := s.GetRealMaxHeat()
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := NewAutopilotRegistry()

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil {
			// 1% chance of undocking per tick
			roll := physics.RandInRange(0, 100)

			if roll == 1 {
				// undock
				s.CmdUndock(false)
				return
			}

			// check if wallet should be randomized
			if roll%22 == 0 {
				// randomize wallet
				s.Wallet = float64(physics.RandInRange(0, math.MaxInt32/64))
			}

			// check if buy/sell/trash attempts should be made
			if roll%33 == 0 {
				// attempt to sell items in cargo bay on the industrial market
				for _, i := range s.CargoBay.Items {
					st := s.DockedAtStation

					// skip if unpackaged or 0 quantity
					if !i.IsPackaged || i.Quantity == 0 {
						continue
					}

					// skip if dirty
					if i.CoreDirty {
						continue
					}

					// iterate over processes
					for _, p := range st.Processes {
						for _, pi := range p.Process.Inputs {
							// skip if not buying this item type
							if pi.ItemTypeID != i.ItemTypeID {
								continue
							}

							// roll for sell chance
							sellRoll := physics.RandInRange(0, 100)

							if sellRoll%3 == 0 {
								// get random quantity
								q := physics.RandInRange(1, i.Quantity)

								// try to sell item to silo
								s.SellItemToSilo(p.ID, i.ID, q, false)
							}
						}
					}
				}
			} else if roll%32 == 0 {
				// attempt to buy items from the industrial market
				for _, p := range s.DockedAtStation.Processes {
					for _, po := range p.Process.Outputs {
						// skip if ship
						if po.ItemFamilyID == "ship" {
							continue
						}

						// roll for buy chance
						buyRoll := physics.RandInRange(0, 100)

						if buyRoll%3 == 0 {
							// get random quantity
							q := physics.RandInRange(1, 1000)

							// try to buy item from silo
							s.BuyItemFromSilo(p.ID, po.ItemTypeID, q, false)
						}
					}
				}
			} else if roll%84 == 0 {
				toTrash := make([]*Item, 0)

				// verify sufficient volume is being used to warrant trashing
				cv := s.GetRealCargoBayVolume()
				cu := s.TotalCargoBayVolumeUsed(false)

				if cu/cv > 0.75 {
					// trash random items in cargo bay
					for _, i := range s.CargoBay.Items {
						// skip if dirty
						if i.CoreDirty {
							continue
						}

						// roll for trash chance
						trashRoll := physics.RandInRange(0, 100)

						if trashRoll%3 == 0 {
							// mark for trash
							toTrash = append(toTrash, i)
						}
					}

					// commit trash
					for _, i := range toTrash {
						// skip if dirty
						if i.CoreDirty {
							continue
						}

						// trash item
						s.TrashItemInCargo(i.ID, false)
					}
				}
			}
		} else {
			s.gotoNextWanderDestination(85)
		}
	}
}

// helper for behaviour routines that need to wander around the universe
func (s *Ship) gotoNextWanderDestination(stationDockChance int) {
	// get registry
	tgtReg := models.NewTargetTypeRegistry()

	// chance of docking at a station vs flying through a jumphole is based on the stationDockChance parameter
	roll := physics.RandInRange(0, 100)

	if roll < stationDockChance {
		// get and count stations in system
		stations := s.CurrentSystem.stations
		count := len(stations)

		// verify there are candidates
		if count == 0 {
			return
		}

		// pick random station to dock at
		tgt := physics.RandInRange(0, count)

		idx := 0
		for _, e := range stations {
			if idx == tgt {
				// check standings
				v, oh := s.CheckStandings(e.FactionID)

				if v > shared.MIN_DOCK_STANDING && !oh {
					// dock at it
					s.CmdDock(e.ID, tgtReg.Station, false)
					return
				}
			}

			idx++
		}
	} else {
		// get and count jumholes in system
		jumpholes := s.CurrentSystem.jumpholes
		count := len(jumpholes)

		// verify there are candidates
		if count == 0 {
			return
		}

		// pick random jumphole to fly through
		tgt := physics.RandInRange(0, count)

		idx := 0
		for _, e := range jumpholes {
			if idx == tgt {
				// fly through it
				s.CmdGoto(e.ID, tgtReg.Jumphole, false)

				hold := 0
				caution := 0

				s.AutopilotGoto.Hold = &hold
				s.AutopilotGoto.Caution = &caution

				return
			}

			idx++
		}
	}
}

// Flies the ship automatically when undocked
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
	case registry.Fight:
		s.doAutopilotFight()
	}
}

// Flies the ship automatically when docked
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

// Causes ship to turn to face a target angle while accelerating
func (s *Ship) doAutopilotManualNav() {
	screenT := s.AutopilotManualNav.Theta

	// calculate magnitude of requested turn
	turnMag := math.Sqrt((screenT - s.Theta) * (screenT - s.Theta))

	a := screenT - s.Theta
	a = physics.FMod(a+180, 360) - 180

	// apply turn with ship limits
	if a > 0 {
		s.rotate(turnMag / s.GetRealTurn())
	} else if a < 0 {
		s.rotate(turnMag / -s.GetRealTurn())
	}

	// thrust forward
	s.forwardThrust(s.AutopilotManualNav.Magnitude)

	// decrease magnitude (this is to allow this to expire and require another move order from the player)
	s.AutopilotManualNav.Magnitude -= s.AutopilotManualNav.Magnitude * s.GetRealSpaceDrag()

	// stop when magnitude is low
	if s.AutopilotManualNav.Magnitude < 0.0001 {
		s.AutopilotMode = NewAutopilotRegistry().None
	}
}

// Causes ship to turn to move towards a target and stop when within range
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
			s.CmdAbort(false)
			return
		}

		// abort if docked
		if tgt.IsDocked {
			s.CmdAbort(false)
			return
		}

		// abort if cloaked
		if tgt.IsCloaked {
			s.CmdAbort(false)
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
			s.CmdAbort(false)
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
			s.CmdAbort(false)
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
			s.CmdAbort(false)
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
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	} else if s.AutopilotGoto.Type == targetTypeReg.Asteroid {
		// find asteroid with id
		tgt := s.CurrentSystem.asteroids[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.Radius
	} else {
		s.CmdAbort(false)
		return
	}

	// get hold / caution overrides if provided
	hold := (s.TemplateData.Radius + tR)

	if s.AutopilotGoto.Hold != nil {
		hold = float64(*s.AutopilotGoto.Hold)
	}

	caution := 30.0

	if s.AutopilotGoto.Caution != nil {
		caution = float64(*s.AutopilotGoto.Caution)
	}

	// fly towards target
	s.flyToPoint(tX, tY, hold, caution)
}

// Causes ship to fly a circle around the target
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
			s.CmdAbort(false)
			return
		}

		// abort if docked
		if tgt.IsDocked {
			s.CmdAbort(false)
			return
		}

		// abort if cloaked
		if tgt.IsCloaked {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Station {
		// find station with id
		tgt := s.CurrentSystem.stations[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Star {
		// find star with id
		tgt := s.CurrentSystem.stars[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Planet {
		// find planet with id
		tgt := s.CurrentSystem.planets[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Jumphole {
		// find jumphole with id
		tgt := s.CurrentSystem.jumpholes[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else if s.AutopilotOrbit.Type == targetTypeReg.Asteroid {
		// find asteroid with id
		tgt := s.CurrentSystem.asteroids[s.AutopilotOrbit.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
	} else {
		s.CmdAbort(false)
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

// Causes ship to dock with a target
func (s *Ship) doAutopilotDock() {
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	if s.AutopilotDock.Type == targetTypeReg.Station {
		// find station
		station := s.CurrentSystem.stations[s.AutopilotDock.TargetID.String()]

		if station == nil {
			s.CmdAbort(false)
			return
		}

		// get distance to station
		d := physics.Distance(s.ToPhysicsDummy(), station.ToPhysicsDummy())
		hold := station.Radius * 0.75

		if d > hold {
			// get closer
			s.flyToPoint(station.PosX, station.PosY, hold, 20)
		} else {
			// wait if cloaked
			if s.IsCloaked {
				return
			}

			// dock with station
			s.DockedAtStation = station
			s.DockedAtStationID = &station.ID
			s.AutopilotMode = NewAutopilotRegistry().None
		}
	} else {
		s.CmdAbort(false)
		return
	}
}

// Causes ship to undock from a target
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

// Causes ship to fight with a target
func (s *Ship) doAutopilotFight() {
	// get registry
	targetTypeReg := models.NewTargetTypeRegistry()

	if s.AutopilotFight.Type == targetTypeReg.Ship {
		// find ship
		targetShip := s.CurrentSystem.ships[s.AutopilotFight.TargetID.String()]

		if targetShip == nil {
			s.CmdAbort(false)
			return
		}

		// abort if docked
		if targetShip.IsDocked {
			s.CmdAbort(false)
			return
		}

		// abort if cloaked
		if targetShip.IsCloaked {
			s.CmdAbort(false)
			return
		}

		// use average weapon range to determine stand-off distance (this can be improved a lot with more specific fighting routines)
		totalRange := 0.0
		rangedMods := 0

		for _, m := range s.Fitting.ARack {
			r, f := m.ItemMeta.GetFloat64("range")

			if f {
				totalRange += r
				rangedMods += 1
			}
		}

		avgRange := totalRange / (float64(rangedMods) + Epsilon)
		standOff := avgRange / 2.0

		// fill autopilot data
		s.AutopilotOrbit = OrbitData{
			TargetID: s.AutopilotFight.TargetID,
			Type:     s.AutopilotFight.Type,
			Distance: standOff,
		}

		// reuse orbit autopilot routine to keep distance with target
		s.doAutopilotOrbit()

		if s.CurrentSystem.tickCounter%45 == 0 {
			// try to activate rack A modules
			maxHeat := s.GetRealMaxHeat()
			heatAdd := 0.0

			for i, v := range s.Fitting.ARack {
				// get heat
				h, _ := v.ItemMeta.GetFloat64("activation_heat")

				// get heat ratio
				hr := (s.Heat + heatAdd) / maxHeat

				// determine whether to activate
				roll := physics.RandInRange(0, 100)
				hit := int(hr * 100)

				if roll >= hit {
					// activate module
					s.Fitting.ARack[i].TargetID = &s.AutopilotFight.TargetID
					s.Fitting.ARack[i].TargetType = &s.AutopilotFight.Type
					s.Fitting.ARack[i].WillRepeat = true

					// track heat
					heatAdd += h
				} else {
					// deactivate module
					s.Fitting.ARack[i].TargetID = nil
					s.Fitting.ARack[i].TargetType = nil
					s.Fitting.ARack[i].WillRepeat = false
				}
			}
		} else if s.CurrentSystem.tickCounter%37 == 0 {
			// try to activate rack B modules
			maxHeat := s.GetRealMaxHeat()
			heatAdd := 0.0

			for i, v := range s.Fitting.BRack {
				if v.ItemTypeFamily == "shield_booster" {
					// make sure enough shield has been lost for this to be worth it
					shieldBoost, _ := v.ItemMeta.GetFloat64("shield_boost_amount")
					maxShield := s.GetRealMaxShield()

					if s.Shield+shieldBoost >= 0.75*maxShield {
						continue
					}
				}

				// get heat
				h, _ := v.ItemMeta.GetFloat64("activation_heat")

				// get heat ratio
				hr := (s.Heat + heatAdd) / maxHeat

				// determine whether to activate
				roll := physics.RandInRange(0, 100)
				hit := int(hr * 100)

				if roll >= hit {
					// activate module
					s.Fitting.BRack[i].TargetID = &s.AutopilotFight.TargetID
					s.Fitting.BRack[i].TargetType = &s.AutopilotFight.Type
					s.Fitting.BRack[i].WillRepeat = true

					// track heat
					heatAdd += h
				} else {
					// deactivate module
					s.Fitting.BRack[i].TargetID = nil
					s.Fitting.BRack[i].TargetType = nil
					s.Fitting.BRack[i].WillRepeat = false
				}
			}
		}
	} else {
		s.CmdAbort(false)
		return
	}
}

// Reusable function to fly a ship towards a point
func (s *Ship) flyToPoint(tX float64, tY float64, hold float64, caution float64) {
	// face towards target
	turnMag := s.facePoint(tX, tY)

	// determine whether to thrust forward and by how much
	scale := ((s.GetRealAccel() * (caution / s.GetRealSpaceDrag())) / 0.175)
	d := (physics.Distance(s.ToPhysicsDummy(), physics.Dummy{PosX: tX, PosY: tY}) - hold)

	if turnMag < 1 {
		// thrust forward
		s.forwardThrust(d / scale)
	}
}

// Reusable function to turn a ship towards a point (returns the turn magnitude needed in degrees)
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

// Turn the ship
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

// Fire the ship's thrusters
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

	// account for additional drag effects
	drag := s.GetRealSpaceDrag()
	dragRatio := SpaceDrag / drag

	// accelerate along theta using thrust proportional to bounded magnitude
	s.VelX += math.Cos(s.Theta*(math.Pi/-180)) * (burn * dragRatio)
	s.VelY += math.Sin(s.Theta*(math.Pi/-180)) * (burn * dragRatio)

	// consume fuel
	s.Fuel -= math.Abs(burn) * ShipFuelBurn

	// accumulate heat
	s.Heat += math.Abs(burn) * ShipHeatBurn
}

// Finds a module fitted on this ship
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

// Returns an item in the ship's cargo bay if it is present
func (s *Ship) FindItemInCargo(id uuid.UUID) *Item {
	// look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		if item.ID == id {
			// return item
			return item
		}
	}

	// nothing found
	return nil
}

// Returns the first available (clean and nonzero quantity) item of a given type and packaging state in the cargo bay
func (s *Ship) FindFirstAvailableItemOfTypeInCargo(typeID uuid.UUID, packaged bool) *Item {
	// look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		if item.CoreDirty {
			continue
		}

		if item.Quantity <= 0 {
			continue
		}

		if item.ItemTypeID == typeID && item.IsPackaged == packaged {
			// return item
			return item
		}
	}

	// nothing found
	return nil
}

// Removes an item from the cargo hold and fits it to the ship
func (s *Ship) FitModule(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.FitModule")
		defer s.Lock.Unlock()
	}

	// lock containers
	s.CargoBay.Lock.Lock("ship.FitModule")
	defer s.CargoBay.Lock.Unlock()

	s.FittingBay.Lock.Lock("ship.FitModule")
	defer s.FittingBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to fit a module")
	}

	// get the item to be packaged
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// get module rack
	r, _ := item.ItemTypeMeta.GetString("rack")
	var rack *[]FittedSlot = nil

	if r == "a" {
		rack = &s.Fitting.ARack
	} else if r == "b" {
		rack = &s.Fitting.BRack
	} else if r == "c" {
		rack = &s.Fitting.CRack
	}

	if rack == nil {
		return errors.New("rack not found")
	}

	// get module volume
	v, _ := item.ItemTypeMeta.GetFloat64("volume")

	// find a free slot
	idx, fnd := s.getFreeSlotIndex(item.ItemFamilyID, v, r)

	if !fnd || idx < 0 {
		return errors.New("no available and compatible slot found to fit this module")
	}

	// remove item from cargo bay and move to fitting bay
	newCB := make([]*Item, 0)

	for _, i := range s.CargoBay.Items {
		if i.ID == item.ID {
			// reassign to fitting bay and escalate to save
			i.ContainerID = s.FittingBayContainerID
			s.FittingBay.Items = append(s.FittingBay.Items, i)
			s.CurrentSystem.MovedItems[i.ID.String()] = i
		} else {
			// keep in cargo bay
			newCB = append(newCB, i)
		}
	}

	s.CargoBay.Items = newCB

	// copy loop index value
	cpyIdx := idx

	// create new fitted slot for item
	fs := FittedSlot{
		ItemID:             item.ID,
		ItemTypeID:         item.ItemTypeID,
		ItemMeta:           item.Meta,
		ItemTypeMeta:       item.ItemTypeMeta,
		ItemTypeFamily:     item.ItemFamilyID,
		ItemTypeFamilyName: item.ItemFamilyName,
		ItemTypeName:       item.ItemTypeName,
		SlotIndex:          &cpyIdx,
	}

	fs.shipMountedOn = s

	// fit to ship
	(*rack)[idx] = fs

	// increase max armor if needed
	armorMaxAdd, f := fs.ItemMeta.GetFloat64("armor_max_add")

	if f {
		s.Armor += armorMaxAdd
	}

	// success!
	return nil
}

// Removes a module from a ship and places it in the cargo hold
func (s *Ship) UnfitModule(m *FittedSlot, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.UnfitModule")
		defer s.Lock.Unlock()
	}

	// lock containers
	m.shipMountedOn.CargoBay.Lock.Lock("ship.UnfitModule")
	defer m.shipMountedOn.CargoBay.Lock.Unlock()

	m.shipMountedOn.FittingBay.Lock.Lock("ship.UnfitModule")
	defer m.shipMountedOn.FittingBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to unfit a module")
	}

	// get module volume
	v, _ := m.ItemTypeMeta.GetFloat64("volume")

	// make sure there is sufficient space in the cargo bay
	if s.TotalCargoBayVolumeUsed(lock)+v > s.GetRealCargoBayVolume() {
		return errors.New("insufficient room in cargo bay to unfit module")
	}

	// make sure the module is not cycling
	if m.IsCycling || m.WillRepeat {
		return errors.New("modules must be offline to be unfit")
	}

	// if the module is in rack c, make sure the ship is fully repaired
	if m.Rack == "C" {
		if s.Armor < s.GetRealMaxArmor() || s.Hull < s.GetRealMaxHull() {
			return errors.New("armor and hull must be fully repaired before unfitting modules in rack c")
		}
	}

	// remove from fitting data
	m.shipMountedOn.Fitting.stripModuleFromFitting(m.ItemID)

	// reassign item to cargo bay
	newFB := make([]*Item, 0)

	for i := range m.shipMountedOn.FittingBay.Items {
		o := m.shipMountedOn.FittingBay.Items[i]

		// lock item
		o.Lock.Lock("ship.UnfitModule")
		defer o.Lock.Unlock()

		// skip if not this module
		if o.ID != m.ItemID {
			newFB = append(newFB, o)
		} else {
			// move to cargo bay if there is still room
			if s.TotalCargoBayVolumeUsed(lock)+v <= s.GetRealCargoBayVolume() {
				m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, o)
				o.ContainerID = m.shipMountedOn.CargoBayContainerID

				// escalate to core to save to db
				s.CurrentSystem.MovedItems[o.ID.String()] = o
			} else {
				return errors.New("insufficient room in cargo bay to unfit module")
			}
		}
	}

	s.FittingBay.Items = newFB

	// success!
	return nil
}

// Trashes an item in the ship's cargo bay if it exists
func (s *Ship) TrashItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.TrashItemInCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.TrashItemInCargo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to trash an item")
	}

	// remove from cargo bay
	newCB := make([]*Item, 0)

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		// skip if not this item or is dirty
		if o.ID != id || o.CoreDirty {
			newCB = append(newCB, o)
		} else {
			// move to trash
			o.ContainerID = s.TrashContainerID
			o.CoreDirty = true

			// escalate to core to save to db
			s.CurrentSystem.MovedItems[o.ID.String()] = o
		}
	}

	s.CargoBay.Items = newCB

	return nil
}

// Removes an item from the ship's cargo bay if it exists, without updating its container id
func (s *Ship) RemoveItemFromCargo(id uuid.UUID, lock bool) (*Item, error) {
	if lock {
		// lock entity
		s.Lock.Lock("ship.RemoveItemFromCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.RemoveItemFromCargo")
	defer s.CargoBay.Lock.Unlock()

	// remove from cargo bay
	newCB := make([]*Item, 0)

	var pulled *Item = nil

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		// skip if not this item or is dirty
		if o.ID != id || o.CoreDirty {
			newCB = append(newCB, o)
		} else {
			// store pointer
			pulled = o
		}
	}

	s.CargoBay.Items = newCB

	// return pulled item
	return pulled, nil
}

// Places an item in the cargo bay without updating its container id
func (s *Ship) PutItemInCargo(item *Item, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.PutItemInCargoHold")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.PutItemInCargoHold")
	defer s.CargoBay.Lock.Unlock()

	// null check
	if item == nil {
		return errors.New("item must not be null")
	}

	// make sure there is enough space
	used := s.TotalCargoBayVolumeUsed(lock)
	max := s.GetRealCargoBayVolume()

	tV := 0.0

	if item.IsPackaged {
		// get item type volume metadata
		volume, f := item.ItemTypeMeta.GetFloat64("volume")

		if f {
			tV += (volume * float64(item.Quantity))
		}
	} else {
		// get item volume metadata
		volume, f := item.Meta.GetFloat64("volume")

		if f {
			tV += (volume * float64(item.Quantity))
		}
	}

	if tV+used > max {
		return errors.New("insufficient cargo space")
	}

	// place item in cargo bay slice
	s.CargoBay.Items = append(s.CargoBay.Items, item)

	return nil
}

// Packages an item in the ship's cargo bay if it exists
func (s *Ship) PackageItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.PackageItemInCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.PackageItemInCargo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to package an item")
	}

	// get the item to be packaged
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure the item is unpackaged
	if item.IsPackaged {
		return errors.New("item is already packaged")
	}

	// make sure the item is fully repaired
	iHp, f := item.Meta.GetFloat64("hp")
	tHp, g := item.ItemTypeMeta.GetFloat64("hp")

	if f && g {
		if iHp < tHp {
			return errors.New("item must be fully repaired before packaging")
		}
	}

	// package item in-memory
	item.IsPackaged = true
	item.CoreDirty = true

	// wipe out item metadata
	item.Meta = Meta{}

	// escalate item for packaging in db
	s.CurrentSystem.PackagedItems[item.ID.String()] = item

	return nil
}

// Attempts to buy an output from a silo and places it in cargo if successful
func (s *Ship) BuyItemFromSilo(siloID uuid.UUID, itemTypeID uuid.UUID, quantity int, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.BuyItemFromSilo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.BuyItemFromSilo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to buy an item from the station industrial market")
	}

	// make sure there is an escrow container attached to this ship
	if s.EscrowContainerID == nil {
		return errors.New("no escrow container associated with this ship")
	}

	// find the silo
	var silo *StationProcess = nil

	for _, px := range s.DockedAtStation.Processes {
		if px.ID == siloID {
			silo = px
			break
		}
	}

	// verify order exists
	if silo == nil {
		return errors.New("silo not found")
	}

	var output *ProcessOutput = nil

	// verify this silo is selling this item type
	for _, t := range silo.Process.Outputs {
		if t.ItemTypeID == itemTypeID {
			output = &t
			break
		}
	}

	if output == nil {
		return errors.New("silo does not produce this item")
	}

	// verify there is enough to satisfy the order
	state := silo.InternalState.Outputs[itemTypeID.String()]

	if state.Quantity < quantity {
		return errors.New("silo doesn't contain enough of this item type to fulfill this order")
	}

	// lock silo if needed
	if lock {
		silo.Lock.Lock("ship.BuyItemFromSilo")
		defer silo.Lock.Unlock()
	}

	// make sure quantity is positive
	if quantity <= 0 {
		return errors.New("quantity of stack to buy must be positive")
	}

	// verify sufficient funds
	cost := float64(state.Price * quantity)

	if s.Wallet-cost < 0 {
		return errors.New("insufficient CBN to fulfill order")
	}

	// calculate order volume
	unitVolume, _ := output.ItemTypeMeta.GetFloat64("volume")
	orderVolume := unitVolume * float64(quantity)

	// determine if this is a ship
	stIDStr, isShip := output.ItemTypeMeta.GetString("shiptemplateid")

	if !isShip {
		// verify sufficient cargo capacity
		usedBay := s.TotalCargoBayVolumeUsed(lock)
		maxBay := s.GetRealCargoBayVolume()

		if usedBay+orderVolume > maxBay {
			return errors.New("insufficient cargo space available")
		}
	} else {
		// verify quantity requested is 1
		if quantity != 1 {
			return errors.New("ships must be purchased one at a time")
		}

		// verify this won't put more than 5 ships belonging to this player in this station
		dockedCount := 0
		sysShips := s.CurrentSystem.CopyShips(lock)

		for _, ds := range sysShips {
			if ds.UserID == s.UserID {
				if ds.DockedAtStationID != nil {
					if *ds.DockedAtStationID == *s.DockedAtStationID {
						dockedCount++
					}
				}
			}
		}

		if dockedCount > 4 {
			return errors.New("purchase would exceed 5 of your ships docked in the same station")
		}
	}

	// adjust wallet
	s.Wallet -= cost

	// reduce quantity
	state.Quantity -= quantity

	// generate a new uuid
	nid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	if !isShip {
		// make a new item stack of the given size
		newItem := Item{
			ID:            nid,
			ItemTypeID:    output.ItemTypeID,
			Meta:          output.ItemTypeMeta,
			Created:       time.Now(),
			CreatedBy:     &s.UserID,
			CreatedReason: "Bought from silo",
			ContainerID:   s.CargoBayContainerID,
			Quantity:      quantity,
			IsPackaged:    true,
			Lock: shared.LabeledMutex{
				Structure: "Item",
				UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
			},
			ItemTypeName:   output.ItemTypeName,
			ItemFamilyID:   output.ItemFamilyID,
			ItemFamilyName: output.ItemFamilyName,
			ItemTypeMeta:   output.ItemTypeMeta,
			CoreDirty:      true,
		}

		// escalate order save request to core
		s.CurrentSystem.NewItems[nid.String()] = &newItem

		// place item in cargo bay
		s.CargoBay.Items = append(s.CargoBay.Items, &newItem)
	} else {
		// parse template id
		stID, err := uuid.Parse(stIDStr)

		if err != nil {
			return err
		}

		// request a new ship to be generated from this purchase
		r := NewShipPurchase{
			UserID:         s.UserID,
			ShipTemplateID: stID,
			StationID:      *s.DockedAtStationID,
		}

		c, ok := s.CurrentSystem.clients[s.UserID.String()]

		if ok {
			r.Client = c
		}

		// escalate order save request to core
		s.CurrentSystem.NewShipPurchases[nid.String()] = &r
	}

	// success
	return nil
}

// Attempts to sell an item in the cargo bay to a silo and removes it if successful
func (s *Ship) SellItemToSilo(siloID uuid.UUID, itemId uuid.UUID, quantity int, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.SellItemToSilo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.SellItemToSilo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to sell an item on the industrial market")
	}

	// make sure there is an escrow container attached to this ship
	if s.EscrowContainerID == nil {
		return errors.New("no escrow container associated with this ship")
	}

	// get the item to be listed for sale
	item := s.FindItemInCargo(itemId)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// lock item if needed
	if lock {
		item.Lock.Lock("ship.SellItemToSilo")
		defer item.Lock.Unlock()
	}

	// make sure item is packaged
	if !item.IsPackaged {
		return errors.New("item must be packaged before selling on the industrial market")
	}

	// make sure quantity is positive
	if quantity <= 0 {
		return errors.New("quantity of stack to sell must be positive")
	}

	// make sure quantity is bound by stack size
	if quantity > item.Quantity {
		return errors.New("order quantity cannot exceed stack size")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// find the silo
	var silo *StationProcess = nil

	for _, px := range s.DockedAtStation.Processes {
		if px.ID == siloID {
			silo = px
			break
		}
	}

	// verify order exists
	if silo == nil {
		return errors.New("silo not found")
	}

	var input *ProcessInput = nil

	// verify this silo is buying this item type
	for _, t := range silo.Process.Inputs {
		if t.ItemTypeID == item.ItemTypeID {
			input = &t
			break
		}
	}

	if input == nil {
		return errors.New("silo does not consume this item")
	}

	// verify there is enough room in the silo to deliver this order
	state := silo.InternalState.Inputs[item.ItemTypeID.String()]
	m := input.GetIndustrialMetadata()

	if state.Quantity+quantity > m.SiloSize {
		return errors.New("silo is too full to accept this order")
	}

	// adjust wallet
	price := quantity * state.Price
	s.Wallet += float64(price)

	// reduce item quantity
	item.Quantity -= quantity
	item.CoreDirty = true

	// save item
	s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item

	return nil
}

// Attempts to fulfill a sell order and place the item in cargo if successful
func (s *Ship) BuyItemFromOrder(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.BuyItemFromOrder")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.BuyItemFromOrder")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to buy an item on the station orders exchange")
	}

	// make sure there is an escrow container attached to this ship
	if s.EscrowContainerID == nil {
		return errors.New("no escrow container associated with this ship")
	}

	// find the sell order
	order := s.DockedAtStation.OpenSellOrders[id.String()]

	// verify order exists
	if order == nil {
		return errors.New("sell order not found")
	}

	// verify order is clean
	if order.CoreDirty {
		return errors.New("sell order is dirty")
	}

	// verify order is unfulfilled
	if order.Bought != nil || order.BuyerUserID != nil {
		return errors.New("sell order has already been fulfilled")
	}

	// lock order and item if needed
	if lock {
		order.Lock.Lock("ship.BuyItemFromOrder")
		defer order.Lock.Unlock()

		order.Item.Lock.Lock("ship.BuyItemFromOrder")
		defer order.Item.Lock.Unlock()
	}

	// verify sufficient funds
	if s.Wallet-order.AskPrice < 0 {
		return errors.New("insufficient CBN to fulfill order")
	}

	if order.Item.ItemFamilyID != "ship" {
		// calculate order volume
		var unitVolume float64 = 0.0

		if order.Item.IsPackaged {
			v, f := order.Item.ItemTypeMeta.GetFloat64("volume")

			if f {
				unitVolume = v
			}
		} else {
			v, f := order.Item.Meta.GetFloat64("volume")

			if f {
				unitVolume = v
			}
		}

		orderVolume := unitVolume * float64(order.Item.Quantity)

		// verify sufficient cargo capacity
		usedBay := s.TotalCargoBayVolumeUsed(lock)
		maxBay := s.GetRealCargoBayVolume()

		if usedBay+orderVolume > maxBay {
			return errors.New("insufficient cargo space available")
		}
	}

	// find the ship currently being flown by the seller so we can deposit funds in their wallet
	seller := s.CurrentSystem.Universe.FindCurrentPlayerShip(order.SellerUserID, &s.CurrentSystem.ID)

	if seller == nil {
		return errors.New("seller ship not found")
	}

	// check if we need to lock the seller
	if seller.CurrentSystem.ID != s.CurrentSystem.ID {
		// obtain lock
		seller.Lock.Lock("ship.BuyItemFromOrder")
		defer seller.Lock.Unlock()
	}

	// adjust wallets
	seller.Wallet += order.AskPrice
	s.Wallet -= order.AskPrice

	// stamp order as fulfilled
	now := time.Now()

	order.CoreDirty = true
	order.Bought = &now
	order.BuyerUserID = &s.UserID

	// remove order from station
	delete(s.DockedAtStation.OpenSellOrders, order.ID.String())

	// escalate order save request to core
	s.CurrentSystem.BoughtSellOrders[order.ID.String()] = order

	if order.Item.ItemFamilyID != "ship" {
		// mark item as dirty and place it in the cargo container
		order.Item.CoreDirty = true
		order.Item.ContainerID = s.CargoBayContainerID

		// escalate item for saving in db
		s.CurrentSystem.MovedItems[order.Item.ID.String()] = order.Item

		// place item in cargo bay
		s.CargoBay.Items = append(s.CargoBay.Items, order.Item)
	} else {
		// escalate request to complete purchase to core
		purShipIDStr, _ := order.Item.Meta.GetString("shipid")
		purShipID := uuid.MustParse(purShipIDStr)

		r := UsedShipPurchase{
			UserID: s.UserID,
			ShipID: purShipID,
		}

		c, ok := s.CurrentSystem.clients[s.UserID.String()]

		if ok {
			r.Client = c
		}

		s.CurrentSystem.UsedShipPurchases[order.ID.String()] = &r

		// mark stub item as dirty and place it in the trash container
		order.Item.CoreDirty = true
		order.Item.ContainerID = s.TrashContainerID

		// escalate stub item for saving in db
		s.CurrentSystem.MovedItems[order.Item.ID.String()] = order.Item
	}

	// success
	return nil
}

// Lists this ship on the station order exchange
func (s *Ship) SellSelfAsOrder(price float64, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.SellSelfAsOrder")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.SellSelfAsOrder")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to list an item on the station orders exchange")
	}

	// make sure there is an escrow container attached to this ship
	if s.EscrowContainerID == nil {
		return errors.New("no escrow container associated with this ship")
	}

	// make sure the ask price is > 0
	if price <= 0 {
		return errors.New("items must be sold at a price greater than 0")
	}

	// remove self from system
	s.CurrentSystem.RemoveShip(s, lock)

	// create a stub item for this ship
	stubID := uuid.New()

	meta := Meta{}
	meta["shipid"] = s.ID
	meta["wallet"] = s.Wallet
	meta["armor"] = s.Armor
	meta["hull"] = s.Hull
	meta["fuel"] = s.Fuel
	meta["seller"] = s.CharacterName

	for i, e := range s.Fitting.ARack {
		k := fmt.Sprintf("rackA[%v]", i)
		meta[k] = e.ItemTypeName
	}

	for i, e := range s.Fitting.BRack {
		k := fmt.Sprintf("rackB[%v]", i)
		meta[k] = e.ItemTypeName
	}

	for i, e := range s.Fitting.CRack {
		k := fmt.Sprintf("rackC[%v]", i)
		meta[k] = e.ItemTypeName
	}

	for i, e := range s.CargoBay.Items {
		v := fmt.Sprintf("%v x%v", e.ItemTypeName, e.Quantity)
		k := fmt.Sprintf("cargo[%v]", i)

		meta[k] = v
	}

	stub := Item{
		ID:            stubID,
		ItemTypeID:    s.TemplateData.ItemTypeID,
		Created:       time.Now(),
		CreatedBy:     &s.UserID,
		CreatedReason: "Ship listed for sale",
		ContainerID:   *s.EscrowContainerID,
		Meta:          meta,
		Quantity:      1,
		IsPackaged:    false,
		CoreDirty:     true,
		Lock: shared.LabeledMutex{
			Structure: "Item",
			UID:       fmt.Sprintf("%v :: %v :: %v", stubID, time.Now(), rand.Float64()),
		},
	}

	// escalate ship and stub item for saving in db
	s.CurrentSystem.NewItems[stub.ID.String()] = &stub
	s.CurrentSystem.SetNoLoad[s.ID.String()] = &ShipNoLoadSet{
		ID:   s.ID,
		Flag: true,
	}

	// create sell order for stub item
	nid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	newOrder := SellOrder{
		ID:           nid,
		StationID:    *s.DockedAtStationID,
		ItemID:       stub.ID,
		SellerUserID: s.UserID,
		AskPrice:     price,
		Created:      time.Now(),
		CoreDirty:    true,
		CoreWait:     20, // defer for 20 cycles so that the stub item can be saved first
		Lock: shared.LabeledMutex{
			Structure: "SellOrder",
			UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
		},
	}

	// link stub into sell order
	newOrder.Item = &stub
	newOrder.GetItemIDFromItem = true // we won't know the real item id until after it saves

	// escalate to core for saving in db
	s.CurrentSystem.NewSellOrders[newOrder.ID.String()] = &newOrder

	return nil
}

// Lists an item in the ship's cargo bay on the station order exchange if it exists
func (s *Ship) SellItemAsOrder(id uuid.UUID, price float64, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.SellItemAsOrder")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.SellItemAsOrder")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to list an item on the station orders exchange")
	}

	// make sure there is an escrow container attached to this ship
	if s.EscrowContainerID == nil {
		return errors.New("no escrow container associated with this ship")
	}

	// get the item to be listed for sale
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure the ask price is > 0
	if price <= 0 {
		return errors.New("items must be sold at a price greater than 0")
	}

	// remove from cargo bay
	newCB := make([]*Item, 0)

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		// keep if not this item or is dirty
		if o.ID != id {
			newCB = append(newCB, o)
		}
	}

	s.CargoBay.Items = newCB

	// mark item as dirty and place it in the escrow container
	item.CoreDirty = true
	item.ContainerID = *s.EscrowContainerID

	// escalate item for saving in db
	s.CurrentSystem.MovedItems[item.ID.String()] = item

	// create sell order for item
	nid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	newOrder := SellOrder{
		ID:           nid,
		StationID:    *s.DockedAtStationID,
		ItemID:       item.ID,
		SellerUserID: s.UserID,
		AskPrice:     price,
		Created:      time.Now(),
		CoreDirty:    true,
		Lock: shared.LabeledMutex{
			Structure: "SellOrder",
			UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
		},
	}

	// link item into sell order
	newOrder.Item = item

	// escalate to core for saving in db
	s.CurrentSystem.NewSellOrders[newOrder.ID.String()] = &newOrder

	return nil
}

// Packages an item in the ship's cargo bay if it exists
func (s *Ship) UnpackageItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.UnpackageItemInCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.UnpackageItemInCargo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to unpackage an item")
	}

	// get the item to be packaged
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure the item is packaged
	if !item.IsPackaged {
		return errors.New("item is already unpackaged")
	}

	// make sure there is only one in the stack
	if item.Quantity != 1 {
		return errors.New("must be a stack of 1 to unpackage")
	}

	// unpackage item in-memory
	item.IsPackaged = false
	item.CoreDirty = true

	// copy item type metadata as initial item metadata
	item.Meta = item.ItemTypeMeta

	// escalate item for unpackaging in db
	s.CurrentSystem.UnpackagedItems[item.ID.String()] = item

	return nil
}

// Stacks an item in the ship's cargo bay if it exists
func (s *Ship) StackItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.StackItemInCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.StackItemInCargo")
	defer s.CargoBay.Lock.Unlock()

	// get the item to be stacked
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure the item is packaged
	if !item.IsPackaged {
		return errors.New("only packaged items can be stacked")
	}

	// merge stack into next stack
	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		// skip if this item
		if o.ID == id {
			continue
		} else {
			// see if we can merge into this stack
			if o.IsPackaged && o.ItemTypeID == item.ItemTypeID && !o.CoreDirty && !item.CoreDirty {
				// merge stacks
				q := item.Quantity

				o.Quantity += q
				item.Quantity -= q

				o.CoreDirty = true
				item.CoreDirty = true

				// escalate to core for saving in db
				s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item
				s.CurrentSystem.ChangedQuantityItems[o.ID.String()] = o

				// exit loop
				break
			}
		}
	}

	// remove 0 quantity stacks
	newCB := make([]*Item, 0)

	for i := range s.CargoBay.Items {
		o := s.CargoBay.Items[i]

		// only retain if non empty
		if o.Quantity > 0 {
			newCB = append(newCB, o)
		}
	}

	// update cargo bay
	s.CargoBay.Items = newCB

	return nil
}

// Splits an item stack in the ship's cargo bay if it exists
func (s *Ship) SplitItemInCargo(id uuid.UUID, size int, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.SplitItemInCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.SplitItemInCargo")
	defer s.CargoBay.Lock.Unlock()

	// get the item to be split
	item := s.FindItemInCargo(id)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure the item is packaged
	if !item.IsPackaged {
		return errors.New("only packaged items can be split")
	}

	// make sure we are splitting a positive size
	if size <= 0 {
		return errors.New("new stack size must be positive")
	}

	// make sure we are splitting off less than the quantity - 1
	if size > item.Quantity-1 {
		return errors.New("both output stacks must have a positive size")
	}

	// reduce found stack by new stack size
	item.Quantity -= size
	item.CoreDirty = true

	// escalate to core for saving in db
	s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item

	// make a new item stack of the given size
	nid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	newItem := Item{
		ID:            nid,
		ItemTypeID:    item.ItemTypeID,
		Meta:          item.Meta,
		Created:       time.Now(),
		CreatedBy:     &s.UserID,
		CreatedReason: "Stack split",
		ContainerID:   s.CargoBayContainerID,
		Quantity:      size,
		IsPackaged:    true,
		Lock: shared.LabeledMutex{
			Structure: "Item",
			UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
		}, ItemTypeName: item.ItemTypeName,
		ItemFamilyID:   item.ItemFamilyID,
		ItemFamilyName: item.ItemFamilyName,
		ItemTypeMeta:   item.ItemTypeMeta,
		CoreDirty:      true,
	}

	// escalate to core for saving in db
	s.CurrentSystem.NewItems[newItem.ID.String()] = &newItem

	// add new item to cargo hold
	s.CargoBay.Items = append(s.CargoBay.Items, &newItem)

	return nil
}

func (s *Ship) SelfDestruct(lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.SelfDestruct")
		defer s.Lock.Unlock()
	}

	// make sure ship is undocked
	if s.DockedAtStationID != nil {
		return errors.New("you must be undocked to self destruct")
	}

	// arm if needed
	if !s.DestructArmed {
		s.DestructArmed = true
		return errors.New("issue self destruct again to detonate")
	}

	// blow it up
	s.DealDamage(s.GetRealMaxShield()+1, s.GetRealMaxArmor()+1, s.GetRealMaxHull()+1, nil)

	return nil
}

// Attempts to consume a fuel pellet from cargo and add it to the fuel tank
func (s *Ship) ConsumeFuelFromCargo(itemID uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.ConsumeFuelFromCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.ConsumeFuelFromCargo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to consume a fuel pellet")
	}

	// find item in cargo bay
	item := s.FindItemInCargo(itemID)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure item is a fuel pellet
	if item.ItemFamilyID != "fuel" {
		return errors.New("item is not a fuel pellet")
	}

	// make sure item is unpackaged
	if item.IsPackaged {
		return errors.New("only unpackaged pellets can be consumed")
	}

	// determine fuel quantity to add
	factor, f := item.Meta.GetInt("fuelconversion")

	if !f {
		return errors.New("missing conversion factor")
	}

	// add fuel to tank
	s.Fuel += float64(factor)

	// reduce quantity of item to 0 (always unpackaged)
	item.Quantity = 0
	item.CoreDirty = true

	// escalate to core for saving in db
	s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item

	return nil
}

// Attempts to consume a repair kit from cargo and apply it to health
func (s *Ship) ConsumeRepairKitFromCargo(itemID uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock("ship.ConsumeRepairKitFromCargo")
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock("ship.ConsumeRepairKitFromCargo")
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to consume a repair kit")
	}

	// find item in cargo bay
	item := s.FindItemInCargo(itemID)

	if item == nil {
		return errors.New("item not found in cargo bay")
	}

	// make sure item is clean
	if item.CoreDirty {
		return errors.New("item is dirty and waiting on an escalation to save its state")
	}

	// make sure item is a repair kit
	if item.ItemFamilyID != "repair_kit" {
		return errors.New("item is not a repair kit")
	}

	// make sure item is unpackaged
	if item.IsPackaged {
		return errors.New("only unpackaged repair kits can be consumed")
	}

	// determine armor and hull quantity to add
	armorFactor, _ := item.Meta.GetInt("armorconversion")
	hullFactor, _ := item.Meta.GetInt("hullconversion")

	// add to health
	s.Armor += float64(armorFactor)
	s.Hull += float64(hullFactor)

	// limit health
	maxArmor := s.GetRealMaxArmor()
	maxHull := s.GetRealMaxHull()

	if s.Armor > maxArmor {
		s.Armor = maxArmor
	}

	if s.Hull > maxHull {
		s.Hull = maxHull
	}

	// reduce quantity of item to 0 (always unpackaged)
	item.Quantity = 0
	item.CoreDirty = true

	// escalate to core for saving in db
	s.CurrentSystem.ChangedQuantityItems[item.ID.String()] = item

	return nil
}

// Updates a fitted slot on a ship
func (m *FittedSlot) PeriodicUpdate() {
	if m.IsCycling {
		// update cycle timer
		cooldown, found := m.ItemMeta.GetFloat64("cooldown")

		if !found {
			// module has no cooldown - deactivate
			m.IsCycling = false
			m.WillRepeat = false

			return
		}

		cooldownMs := int(cooldown * 1000)
		m.cooldownProgress += Heartbeat

		if m.cooldownProgress > cooldownMs {
			// cycle completed
			m.IsCycling = false
			m.cooldownProgress = 0
			m.CyclePercent = 0

			// update experience
			if m.shipMountedOn.ExperienceSheet != nil && m.shipMountedOn.BeingFlownByPlayer {
				// verify there is a client connected to this module's ship
				c := m.shipMountedOn.CurrentSystem.clients[m.shipMountedOn.UserID.String()]

				if c != nil {
					// get experience
					xp := m.shipMountedOn.ExperienceSheet.GetModuleExperienceEntry(m.ItemTypeID)

					// cache current level
					xl := math.Trunc(xp.GetExperience())

					// update experience
					xp.SecondsOfExperience += cooldown
					xp.ItemTypeName = m.ItemTypeName
					m.shipMountedOn.ExperienceSheet.SetModuleExperienceEntry(xp)

					// cache new level
					nl := math.Trunc(xp.GetExperience())

					if nl > xl {
						// notify player of the level up!
						c.WriteInfoMessage(
							fmt.Sprintf(
								"you have advanced to %v level %v!",
								m.ItemTypeName,
								nl,
							),
						)
					}
				}
			}
		} else {
			// update percentage
			m.CyclePercent = ((m.cooldownProgress * 100) / cooldownMs)
		}
	} else {
		// check for activation intent
		if m.WillRepeat {
			// cache usage experience modifier
			if m.shipMountedOn.ExperienceSheet != nil && m.shipMountedOn.BeingFlownByPlayer {
				m.usageExperienceModifier = m.GetExperienceModifier()
			} else {
				m.usageExperienceModifier = 1.0
			}

			// check if cloaked (exempting cloaking devices)
			if m.shipMountedOn.IsCloaked {
				// cloaked - can't active
				m.WillRepeat = false
				return
			}

			// check if a target is required
			needsTarget, _ := m.ItemTypeMeta.GetBool("needs_target")

			if needsTarget {
				// check for a target
				if m.TargetID == nil || m.TargetType == nil {
					// no target - can't activate
					m.WillRepeat = false
					return
				}

				// make sure the target actually exists in this solar system
				tgtReg := models.NewTargetTypeRegistry()

				if *m.TargetType == tgtReg.Ship {
					// find ship
					sx, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}

					if sx.IsCloaked {
						// target is cloaked - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}

					// make sure this isn't the ship it's mounted on
					if sx.ID == m.shipMountedOn.ID {
						// target is self - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Station {
					// find station
					_, f := m.shipMountedOn.CurrentSystem.stations[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Planet {
					// find planet
					_, f := m.shipMountedOn.CurrentSystem.planets[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Jumphole {
					// find jumphole
					_, f := m.shipMountedOn.CurrentSystem.jumpholes[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Star {
					// find star
					_, f := m.shipMountedOn.CurrentSystem.stars[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
						m.TargetID = nil
						m.TargetType = nil

						return
					}
				} else if *m.TargetType == tgtReg.Asteroid {
					// find asteroid
					_, f := m.shipMountedOn.CurrentSystem.asteroids[m.TargetID.String()]

					if !f {
						// target doesn't exist - can't activate
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

			// check for sufficient activation energy
			activationEnergy, found := m.ItemMeta.GetFloat64("activation_energy")

			if found {
				if m.shipMountedOn.Energy-activationEnergy < 0 {
					// insufficient energy - can't activate
					m.WillRepeat = false
					return
				}
			}

			// to determine whether activation succeeds later
			canActivate := false

			// handle module family effects
			if m.ItemTypeFamily == "gun_turret" {
				canActivate = m.activateAsGunTurret()
			} else if m.ItemTypeFamily == "shield_booster" {
				canActivate = m.activateAsShieldBooster()
			} else if m.ItemTypeFamily == "eng_oc" {
				canActivate = m.activateAsEngineOvercharger()
			} else if m.ItemTypeFamily == "active_sink" {
				canActivate = m.activateAsActiveRadiator()
			} else if m.ItemTypeFamily == "missile_launcher" {
				canActivate = m.activateAsMissileLauncher()
			} else if m.ItemTypeFamily == "drag_amp" {
				canActivate = m.activateAsAetherDragger()
			} else if m.ItemTypeFamily == "utility_miner" {
				canActivate = m.activateAsUtilityMiner()
			} else if m.ItemTypeFamily == "utility_siphon" {
				canActivate = m.activateAsUtilitySiphon()
			} else if m.ItemTypeFamily == "utility_cloak" {
				canActivate = m.activateAsUtilityCloak()
			}

			if canActivate {
				// activate module
				m.shipMountedOn.Energy -= activationEnergy
				m.IsCycling = true

				// apply activation heating
				activationHeat, found := m.ItemMeta.GetFloat64("activation_heat")

				if found {
					m.shipMountedOn.Heat += activationHeat
				}
			}
		}
	}
}

func (m *FittedSlot) activateAsGunTurret() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.NewTargetTypeRegistry()

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
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
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else if *m.TargetType == tgtReg.Asteroid {
		// find asteroid
		tgt, f := m.shipMountedOn.CurrentSystem.asteroids[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// check if ammunition required to fire (note: NPCs do not require ammunition - i want to fix this eventually)
	ammoTypeRaw, found := m.ItemMeta.GetString("ammunition_type")
	var ammoItem *Item = nil

	if found && !m.shipMountedOn.IsNPC {
		// parse type id
		typeID, _ := uuid.Parse(ammoTypeRaw)

		// verify there is enough ammunition to fire
		x := m.shipMountedOn.FindFirstAvailableItemOfTypeInCargo(typeID, true)

		if x == nil {
			return false
		}

		// store item to take from
		ammoItem = x
	}

	// get damage values
	shieldDmg, _ := m.ItemMeta.GetFloat64("shield_damage")
	armorDmg, _ := m.ItemMeta.GetFloat64("armor_damage")
	hullDmg, _ := m.ItemMeta.GetFloat64("hull_damage")

	// apply usage experience modifier
	shieldDmg *= m.usageExperienceModifier
	armorDmg *= m.usageExperienceModifier
	hullDmg *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust damage based on tracking
	shieldDmg *= trackingRatio
	armorDmg *= trackingRatio
	hullDmg *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% damage (or ore pull if asteroid)

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			shieldDmg *= rangeRatio
			armorDmg *= rangeRatio
			hullDmg *= rangeRatio
		} else if falloff == "reverse_linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			shieldDmg *= rangeRatio
			armorDmg *= rangeRatio
			hullDmg *= rangeRatio
		}
	}

	// reduce ammunition count if needed
	if ammoItem != nil {
		ammoItem.Quantity--
		ammoItem.CoreDirty = true

		// escalate for saving
		m.shipMountedOn.CurrentSystem.ChangedQuantityItems[ammoItem.ID.String()] = ammoItem
	}

	// apply damage (or ore / ice pulled if asteroid) to target
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)
		c.DealDamage(shieldDmg, armorDmg, hullDmg, m.shipMountedOn.ReputationSheet)
	} else if *m.TargetType == tgtReg.Station {
		c := tgtI.(*Station)
		c.DealDamage(shieldDmg, armorDmg, hullDmg)
	} else if *m.TargetType == tgtReg.Asteroid {
		// target is an asteroid - what can this module mine?
		canMineOre, _ := m.ItemTypeMeta.GetBool("can_mine_ore")
		canMineIce, _ := m.ItemTypeMeta.GetBool("can_mine_ice")

		// get ore / ice type and volume
		c := tgtI.(*Asteroid)

		// can this module mine this type?
		canMine := false

		if c.ItemFamilyID == "ore" && canMineOre {
			canMine = true
		} else if c.ItemFamilyID == "ice" && canMineIce {
			canMine = true
		}

		if canMine {
			// get mining volume
			oreMiningVolume, _ := m.ItemMeta.GetFloat64("ore_mining_volume")
			iceMiningVolume, _ := m.ItemMeta.GetFloat64("ice_mining_volume")

			miningVolume := 0.0

			if c.ItemFamilyID == "ore" && canMineOre {
				miningVolume = oreMiningVolume
			} else if c.ItemFamilyID == "ice" && canMineIce {
				miningVolume = iceMiningVolume
			}

			// get type and volume of ore / ice being collected
			unitType := c.ItemTypeID
			unitVol, _ := c.ItemTypeMeta.GetFloat64("volume")

			// get available space in cargo hold
			free := m.shipMountedOn.GetRealCargoBayVolume() - m.shipMountedOn.TotalCargoBayVolumeUsed(false)

			// calculate effective ore / ice volume pulled
			pulled := miningVolume * c.Yield * rangeRatio

			// apply usage experience modifier
			pulled *= m.usageExperienceModifier

			// make sure there is sufficient room to deposit the ore / ice
			if free-pulled >= 0 {
				found := false

				// quantity to be placed in cargo bay
				q := int((miningVolume * c.Yield) / unitVol)

				// is there already packaged ore / ice of this type in the hold?
				for idx := range m.shipMountedOn.CargoBay.Items {
					itm := m.shipMountedOn.CargoBay.Items[idx]

					if itm.ItemTypeID == unitType && itm.IsPackaged && !itm.CoreDirty {
						// increase the size of this stack
						itm.Quantity += q

						// escalate for saving
						m.shipMountedOn.CurrentSystem.ChangedQuantityItems[itm.ID.String()] = itm

						// mark as found
						found = true
						break
					}
				}

				if !found && q > 0 {
					// create a new stack of ore / ice
					nid := uuid.New()

					newItem := Item{
						ID:            nid,
						ItemTypeID:    unitType,
						Meta:          c.ItemTypeMeta,
						Created:       time.Now(),
						CreatedBy:     &m.shipMountedOn.UserID,
						CreatedReason: fmt.Sprintf("Mined %v", c.ItemFamilyID),
						ContainerID:   m.shipMountedOn.CargoBayContainerID,
						Quantity:      q,
						IsPackaged:    true,
						Lock: shared.LabeledMutex{
							Structure: "Item",
							UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
						}, ItemTypeName: c.ItemTypeName,
						ItemFamilyID:   c.ItemFamilyID,
						ItemFamilyName: c.ItemFamilyName,
						ItemTypeMeta:   c.ItemTypeMeta,
						CoreDirty:      true,
					}

					// escalate to core for saving in db
					m.shipMountedOn.CurrentSystem.NewItems[newItem.ID.String()] = &newItem

					// add new item to cargo hold
					m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, &newItem)
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsMissileLauncher() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.NewTargetTypeRegistry()

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
	} else if *m.TargetType == tgtReg.Station {
		// find station
		tgt, f := m.shipMountedOn.CurrentSystem.stations[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// check if ammunition required to fire (note: NPCs do not require ammunition - i want to fix this eventually)
	ammoTypeRaw, found := m.ItemMeta.GetString("ammunition_type")
	var ammoItem *Item = nil

	if found && !m.shipMountedOn.IsNPC {
		// parse type id
		typeID, _ := uuid.Parse(ammoTypeRaw)

		// verify there is enough ammunition to fire
		x := m.shipMountedOn.FindFirstAvailableItemOfTypeInCargo(typeID, true)

		if x == nil {
			return false
		}

		// store item to take from
		ammoItem = x
	}

	// reduce ammunition count if needed
	if ammoItem != nil {
		ammoItem.Quantity--
		ammoItem.CoreDirty = true

		// escalate for saving
		m.shipMountedOn.CurrentSystem.ChangedQuantityItems[ammoItem.ID.String()] = ammoItem
	}

	// build and hook missile projectile
	missileGfxEffect, _ := m.ItemTypeMeta.GetString("missile_gfx_effect")
	missileRadius, _ := m.ItemTypeMeta.GetFloat64("missile_radius")
	faultTolerance, _ := m.ItemMeta.GetFloat64("fault_tolerance")
	flightTime, _ := m.ItemMeta.GetFloat64("flight_time")
	maxVelocity := (modRange / flightTime)

	// apply usage experience modifiers
	flightTime *= m.usageExperienceModifier
	maxVelocity *= m.usageExperienceModifier

	stubID := uuid.New()
	flightTicks := int((flightTime * 1000) / Heartbeat)

	// determine launcher hardpoint location
	hpX := [...]float64{
		0,
		0,
	}

	if m.SlotIndex != nil {
		rack := m.Rack
		idx := *m.SlotIndex

		if rack == "A" {
			hpX = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
		}
	}

	sRad := physics.ToRadians(m.shipMountedOn.Theta)
	hRad := physics.ToRadians(hpX[1])

	cRad := sRad + hRad
	lx := math.Cos(cRad) * hpX[0]
	ly := math.Sin(cRad) * hpX[0]

	// store stub
	stub := Missile{
		ID:             stubID,
		PosX:           m.shipMountedOn.PosX + lx,
		PosY:           m.shipMountedOn.PosY + ly,
		Texture:        missileGfxEffect,
		Module:         m,
		TargetID:       *m.TargetID,
		TargetType:     *m.TargetType,
		Radius:         missileRadius,
		FiredByID:      m.shipMountedOn.ID,
		TicksRemaining: flightTicks,
		MaxVelocity:    maxVelocity,
		FaultTolerance: faultTolerance,
	}

	m.shipMountedOn.CurrentSystem.missiles[stub.ID.String()] = &stub

	// module activates!
	return true
}

func (m *FittedSlot) activateAsShieldBooster() bool {
	// get shield boost amount
	shieldBoost, _ := m.ItemMeta.GetFloat64("shield_boost_amount")

	// apply usage experience modifier
	shieldBoost *= m.usageExperienceModifier

	// apply boost to mounting ship
	m.shipMountedOn.DealDamage(-shieldBoost, 0, 0, nil)

	// include visual effect if present
	activationPGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		tgtReg := models.NewTargetTypeRegistry()

		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationPGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsEngineOvercharger() bool {
	// get activation energy and duration (same as cooldown for engine overchargers)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship mass
	mx := m.shipMountedOn.GetRealMass()

	// calculate engine boost amount
	dA := (activationEnergy / mx) * 10

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// apply usage experience modifier
	dT *= m.usageExperienceModifier

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "accel",
		Percentage:     dA,
		RemainingTicks: int(dT),
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsActiveRadiator() bool {
	// get activation energy and duration (same as cooldown for active radiators)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship heat capacity
	mx := m.shipMountedOn.GetRealMaxHeat()

	// calculate heat sink amount
	dA := (activationEnergy / mx) * 35

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// apply usage experience modifier
	dT *= m.usageExperienceModifier

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "heat_sink",
		Percentage:     dA,
		RemainingTicks: int(dT),
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsAetherDragger() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.NewTargetTypeRegistry()

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI interface{}

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get drag multiplier
	dragMul, _ := m.ItemMeta.GetFloat64("drag_multiplier")

	// apply usage experience modifier
	dragMul *= m.usageExperienceModifier

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// drag increase is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)
			dragMul *= rangeRatio
		} else if falloff == "reverse_linear" {
			// drag increase is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			dragMul *= rangeRatio
		}
	}

	// apply drag to target
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		// calculate effect duration in ticks\
		cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
		dT := (cooldown * 1000) / Heartbeat

		// add temporary modifier to target
		modifier := TemporaryShipModifier{
			Attribute:      "drag",
			Percentage:     dragMul,
			RemainingTicks: int(dT),
		}

		c.TemporaryModifiers = append(c.TemporaryModifiers, modifier)
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityMiner() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.NewTargetTypeRegistry()

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Asteroid {
		// find asteroid
		tgt, f := m.shipMountedOn.CurrentSystem.asteroids[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% efficiency

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)
		} else if falloff == "reverse_linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}
		}
	}

	// apply damage (or ore / ice pulled if asteroid) to target
	if *m.TargetType == tgtReg.Asteroid {
		// target is an asteroid - what can this module mine?
		canMineOre, _ := m.ItemTypeMeta.GetBool("can_mine_ore")
		canMineIce, _ := m.ItemTypeMeta.GetBool("can_mine_ice")

		// get ore / ice type and volume
		c := tgtI.(*Asteroid)

		// can this module mine this type?
		canMine := false

		if c.ItemFamilyID == "ore" && canMineOre {
			canMine = true
		} else if c.ItemFamilyID == "ice" && canMineIce {
			canMine = true
		}

		if canMine {
			// get mining volume
			oreMiningVolume, _ := m.ItemMeta.GetFloat64("ore_mining_volume")
			iceMiningVolume, _ := m.ItemMeta.GetFloat64("ice_mining_volume")

			miningVolume := 0.0

			if c.ItemFamilyID == "ore" && canMineOre {
				miningVolume = oreMiningVolume
			} else if c.ItemFamilyID == "ice" && canMineIce {
				miningVolume = iceMiningVolume
			}

			// get type and volume of ore / ice being collected
			unitType := c.ItemTypeID
			unitVol, _ := c.ItemTypeMeta.GetFloat64("volume")

			// get available space in cargo hold
			free := m.shipMountedOn.GetRealCargoBayVolume() - m.shipMountedOn.TotalCargoBayVolumeUsed(false)

			// calculate effective ore / ice volume pulled
			pulled := miningVolume * c.Yield * rangeRatio * trackingRatio

			// apply usage experience modifier
			pulled *= m.usageExperienceModifier

			// make sure there is sufficient room to deposit the ore / ice
			if free-pulled >= 0 {
				found := false

				// quantity to be placed in cargo bay
				q := int((miningVolume * c.Yield) / unitVol)

				// is there already packaged ore / ice of this type in the hold?
				for idx := range m.shipMountedOn.CargoBay.Items {
					itm := m.shipMountedOn.CargoBay.Items[idx]

					if itm.ItemTypeID == unitType && itm.IsPackaged && !itm.CoreDirty {
						// increase the size of this stack
						itm.Quantity += q

						// escalate for saving
						m.shipMountedOn.CurrentSystem.ChangedQuantityItems[itm.ID.String()] = itm

						// mark as found
						found = true
						break
					}
				}

				if !found && q > 0 {
					// create a new stack of ore / ice
					nid := uuid.New()

					newItem := Item{
						ID:            nid,
						ItemTypeID:    unitType,
						Meta:          c.ItemTypeMeta,
						Created:       time.Now(),
						CreatedBy:     &m.shipMountedOn.UserID,
						CreatedReason: fmt.Sprintf("Mined %v", c.ItemFamilyID),
						ContainerID:   m.shipMountedOn.CargoBayContainerID,
						Quantity:      q,
						IsPackaged:    true,
						Lock: shared.LabeledMutex{
							Structure: "Item",
							UID:       fmt.Sprintf("%v :: %v :: %v", nid, time.Now(), rand.Float64()),
						}, ItemTypeName: c.ItemTypeName,
						ItemFamilyID:   c.ItemFamilyID,
						ItemFamilyName: c.ItemFamilyName,
						ItemTypeMeta:   c.ItemTypeMeta,
						CoreDirty:      true,
					}

					// escalate to core for saving in db
					m.shipMountedOn.CurrentSystem.NewItems[newItem.ID.String()] = &newItem

					// add new item to cargo hold
					m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, &newItem)
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilitySiphon() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.NewTargetTypeRegistry()

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max siphon amount
	maxSiphonAmt, _ := m.ItemMeta.GetFloat64("energy_siphon_amount")

	// apply usage experience modifier
	maxSiphonAmt *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust siphon amount based on tracking
	maxSiphonAmt *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% damage (or ore pull if asteroid)

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount siphoned is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxSiphonAmt *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount siphoned is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxSiphonAmt *= rangeRatio
		}
	}

	// siphon energy from target ship
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)
		actualSiphon := c.SiphonEnergy(maxSiphonAmt, m.shipMountedOn.ReputationSheet)

		// add to energy
		m.shipMountedOn.Energy += actualSiphon

		// any excess becomes heat
		maxEnergy := m.shipMountedOn.GetRealMaxEnergy()
		excess := m.shipMountedOn.Energy - maxEnergy

		if excess > 0 {
			// apply heat
			m.shipMountedOn.Heat += excess

			// clamp energy
			m.shipMountedOn.Energy = maxEnergy
		}
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityCloak() bool {
	// make sure that no other modules are active
	for _, x := range m.shipMountedOn.Fitting.ARack {
		if x.IsCycling {
			return false
		}
	}

	for _, x := range m.shipMountedOn.Fitting.BRack {
		if x.IsCycling {
			return false
		}
	}

	// get activation energy and duration (approximately same as cooldown for cloaking devices)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship mass
	mx := m.shipMountedOn.GetRealMass()

	// calculate "cloak amount" (if percentage < 100% then cloaking will "flicker")
	dC := (activationEnergy / mx) * 7

	// apply usage experience modifier
	dC *= m.usageExperienceModifier

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "cloak",
		Percentage:     dC,
		RemainingTicks: int(dT) + int(dC*10), // duration bonus for lighter ships
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

// Reusable helper function to determine tracking ratio between a module and a target
func (m *FittedSlot) calculateTrackingRatioWithTarget(tgtDummy physics.Dummy) float64 {
	// calculate distance
	d := physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

	// determine angular velocity for tracking
	dvX := m.shipMountedOn.VelX - tgtDummy.VelX
	dvY := m.shipMountedOn.VelY - tgtDummy.VelY

	dv := math.Sqrt((dvX * dvX) + (dvY * dvY))
	w := 0.0

	if d > 0 {
		w = ((dv / d) * float64(Heartbeat)) * (180.0 / math.Pi)
	}

	// get tracking value
	tracking, _ := m.ItemMeta.GetFloat64("tracking")

	// calculate tracking ratio
	trackingRatio := 1.0 // default to 100% tracking

	if w > 0 {
		trackingRatio = tracking / w
	}

	// clamp tracking to 100%
	if trackingRatio > 1.0 {
		trackingRatio = 1.0
	}

	return trackingRatio
}

// Calculates the experience percentage bonus to apply to some active module stats
func (m *FittedSlot) GetExperienceModifier() float64 {
	mx := 1.0

	if m.shipMountedOn.ExperienceSheet != nil {
		// get experience entry for this item type as a module
		v := m.shipMountedOn.ExperienceSheet.GetModuleExperienceEntry(m.ItemTypeID)

		// get truncated level
		l := math.Trunc(v.GetExperience())

		if l > 0 {
			// apply a dampening factor to get percentage
			b := math.Log(((math.Pow(l, 0.75)) / 8.8) + 1)

			if b > 0 {
				// add bonus
				mx += b
			}
		}
	}

	return mx
}
