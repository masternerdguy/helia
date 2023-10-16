package universe

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
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

// Shared registries - do not modify at runtime!
var SharedAutopilotRegistry = NewAutopilotRegistry()
var SharedBehaviourRegistry = NewBehaviourRegistry()

// Autopilot states for ships
type AutopilotRegistry struct {
	None      int
	ManualNav int
	Goto      int
	Orbit     int
	Dock      int
	Undock    int
	Fight     int
	Mine      int
	Salvage   int
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
		Mine:      7,
		Salvage:   8,
	}
}

// Autopilot states for ships
type BehaviourRegistry struct {
	None         int
	Wander       int
	Patrol       int
	PatchTrade   int
	PatchMine    int
	PatchSalvage int
}

// Returns a populated AutopilotRegistry struct for use as an enum
func NewBehaviourRegistry() *BehaviourRegistry {
	return &BehaviourRegistry{
		None:         0,
		Wander:       1,
		Patrol:       2,
		PatchTrade:   3,
		PatchMine:    4,
		PatchSalvage: 5,
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

// Container structure for arguments of the Fight autopilot mode
type FightData struct {
	TargetID uuid.UUID
	Type     int
}

// Container structure for arguments of the Mine autopilot mode
type MineData struct {
	TargetID uuid.UUID
	Type     int
}

// Container structure for arguments of the Salvage autopilot mode
type SalvageData struct {
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
	IsNPC                  bool
	IsDocked               bool
	Faction                *Faction
	AutopilotMode          int
	AutopilotManualNav     ManualNavData
	AutopilotGoto          GotoData
	AutopilotOrbit         OrbitData
	AutopilotDock          DockData
	AutopilotUndock        UndockData
	AutopilotFight         FightData
	AutopilotMine          MineData
	AutopilotSalvage       SalvageData
	BehaviourMode          *int
	CurrentSystem          *SolarSystem
	DockedAtStation        *Station
	CargoBay               *Container
	FittingBay             *Container
	CachedHeatSink         float64 // cache of output from GetRealHeatSink
	CachedMaxHeat          float64 // cache of output from GetRealMaxHeat
	CachedRealAccel        float64 // cache of output from GetRealAccel
	CachedRealTurn         float64 // cache of output from GetRealTurn
	CachedRealSpaceDrag    float64 // cache of output from GetRealSpaceDrag
	CachedMaxFuel          float64 // cache of output from GetRealMaxFuel
	CachedMaxEnergy        float64 // cache of output from GetRealMaxEnergy
	CachedMaxShield        float64 // cache of output from GetRealMaxShield
	CachedMaxArmor         float64 // cache of output from GetRealMaxArmor
	CachedMaxHull          float64 // cache of output from GetRealMaxHull
	CachedEnergyRegen      float64 // cache of output from GetRealEnergyRegen
	CachedShieldRegen      float64 // cache of output from GetRealShieldRegen
	CachedCargoBayVolume   float64 // cache of output from GetRealCargoBayVolume
	SumCloaking            float64 // cache of sum of cloaking power
	SumVeiling             float64 // cache of sum of veiling power
	EscrowContainerID      *uuid.UUID
	BeingFlownByPlayer     bool
	ReputationSheet        *shared.PlayerReputationSheet
	ExperienceSheet        *shared.PlayerExperienceSheet
	DestructArmed          bool
	LeaveFactionArmed      bool
	TemporaryModifiers     []TemporaryShipModifier
	IsCloaked              bool
	Aggressors             map[string]*shared.PlayerReputationSheet // reputation sheets for players who have attacked this ship
	AggressionLog          map[string]*shared.AggressionLog         // details of aggression
	aiIncompatibleOreFault bool                                     // true when mining autopilot failed due to incompatible ore (for patch miners)
	aiNoOrePulledFault     bool                                     // true when mining autopilot failed due to pulling no ore (for patch miners)
	aiNoWreckFault         bool                                     // true when salvaging autopilot failed due to wreck disappearing (for patch salvagers)
	WreckReady             bool                                     // true when a dead ship has been saved to the db and the wreck can be looted
	InLimbo                bool                                     // true when ship is being migrated between systems
	Lock                   sync.Mutex
}

// Structure representing a percentage or raw modifier applied to a ship for a limit number of ticks
type TemporaryShipModifier struct {
	Attribute      string
	Percentage     float64 // if set, raw should be unset
	Raw            float64 // if set, percentage should be unset
	RemainingTicks int
}

// Structure representing a newly purchased ship, not yet materialized
type NewShipTicket struct {
	ID             uuid.UUID
	ShipTemplateID uuid.UUID
	UserID         uuid.UUID
	StationID      uuid.UUID
	Client         *shared.GameClient
}

// Structure represending an invoked schematic
type NewSchematicRunTicket struct {
	SchematicItem *Item
	Client        *shared.GameClient
	Ship          *Ship
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
	modFamily := getModuleFamily(itemFamilyID)

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

// Helper function to map an item family to its module group, if it is a module
func getModuleFamily(itemFamilyID string) string {
	var modFamily string = ""

	if itemFamilyID == "gun_turret" {
		modFamily = "gun"
	} else if itemFamilyID == "missile_launcher" {
		modFamily = "missile"
	} else if itemFamilyID == "shield_booster" ||
		itemFamilyID == "armor_plate" {
		modFamily = "tank"
	} else if itemFamilyID == "fuel_tank" ||
		itemFamilyID == "fuel_loader" {
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
	} else if itemFamilyID == "utility_add" {
		modFamily = "utility"
	} else if itemFamilyID == "utility_veil" {
		modFamily = "utility"
	} else if itemFamilyID == "battery_pack" {
		modFamily = "power"
	} else if itemFamilyID == "aux_generator" {
		modFamily = "power"
	} else if itemFamilyID == "cargo_expander" {
		modFamily = "cargo"
	} else if itemFamilyID == "salvager" {
		modFamily = "utility"
	} else if itemFamilyID == "ewar_cycle" {
		modFamily = "ewar"
	} else if itemFamilyID == "ewar_fcj" {
		modFamily = "ewar"
	} else if itemFamilyID == "ewar_r_mask" {
		modFamily = "ewar"
	} else if itemFamilyID == "ewar_d_mask" {
		modFamily = "ewar"
	} else if itemFamilyID == "therm_cap" {
		modFamily = "heat"
	} else if itemFamilyID == "burst_reactor" {
		modFamily = "power"
	} else if itemFamilyID == "xfer_heat" {
		modFamily = "utility"
	} else if itemFamilyID == "xfer_energy" {
		modFamily = "utility"
	} else if itemFamilyID == "xfer_shield" {
		modFamily = "utility"
	}

	return modFamily
}

// Returns a copy of the ship
func (s *Ship) CopyShip(lock bool) *Ship {
	if lock {
		s.Lock.Lock()
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
		Lock:               sync.Mutex{},
		IsDocked:           s.IsDocked,
		AutopilotMode:      s.AutopilotMode,
		AutopilotManualNav: s.AutopilotManualNav,
		AutopilotGoto:      s.AutopilotGoto,
		AutopilotOrbit:     s.AutopilotOrbit,
		AutopilotDock:      s.AutopilotDock,
		AutopilotUndock:    s.AutopilotUndock,
		AutopilotFight:     s.AutopilotFight,
		AutopilotMine:      s.AutopilotMine,
		AutopilotSalvage:   s.AutopilotSalvage,
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
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// perform guaranteed updates
	s.alwaysPeriodicUpdate()

	// check if docked or undocked at a station (docking with other objects not yet supported)
	if s.DockedAtStationID == nil {
		s.undockedPeriodicUpdate()
	} else {
		s.dockedPeriodicUpdate()
	}

	// fix any NaN weirdness
	s.fixLocalNaNs()
}

// Perform periodic update steps that will always happen
func (s *Ship) alwaysPeriodicUpdate() {
	// update experience modifier
	if s.ExperienceSheet != nil {
		s.FlightExperienceModifier = s.GetExperienceModifier()
	}

	// cache system name
	if s.CurrentSystem != nil {
		s.SystemName = s.CurrentSystem.SystemName
	}

	// cache player faction id on rep sheet
	if s.Faction != nil && s.ReputationSheet != nil {
		s.ReputationSheet.FactionID = s.FactionID
	}

	// remax some stats if needed for spawning
	if s.ReMaxDirty {
		s.ReMaxStatsForSpawn()
	}

	// special cheats NPCs get to make things easier to implement for now (i do want to eliminate these)
	if s.IsNPC {
		if s.Fuel <= 0 {
			// infinite fuel via refill at 0
			s.Fuel = s.GetRealMaxFuel(false)
		}
	}

	if len(s.TemporaryModifiers) > 0 {
		// handle temporary modifiers
		keptTemporaryModifiers := make([]TemporaryShipModifier, 0)

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
	}

	// remove zero-quantity items from cargo bay
	if s.CurrentSystem != nil && s.CurrentSystem.tickCounter%50 == 0 {
		if len(s.CargoBay.Items) > 0 {
			s.removeZeroQuantityItemsFromCargo()
		}
	}

	// null check
	if s.CurrentSystem == nil {
		return
	}

	// update cloaking (safe to do every few ticks)
	if s.CurrentSystem.tickCounter%6 == 0 {
		s.updateCloaking()
	}

	// update veiling (safe to do every few ticks)
	if s.CurrentSystem.tickCounter%7 == 0 {
		s.updateVeiling()
	}

	// determine whether to recalculate cargo capacity
	rcc := s.CurrentSystem.tickCounter%22 == 0

	// chance to update cargo capacity
	s.GetRealCargoBayVolume(rcc)

	// update energy
	s.updateEnergy()

	// update shields
	s.updateShield()

	// update heat
	s.updateHeat()

	// run behaviour routine if applicable
	s.behave()
}

// Perform periodic update steps that will only happen if undocked
func (s *Ship) undockedPeriodicUpdate() {
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

	// determine whether to recalculate real drag
	rc := s.CurrentSystem.tickCounter%8 == 0

	// calculate felt drag
	drag := s.GetRealSpaceDrag(rc)

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
}

// Perform periodic update steps that will only happen if docked
func (s *Ship) dockedPeriodicUpdate() {
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

// Helper function to reset any strange NaN values to zero
func (s *Ship) fixLocalNaNs() {
	if math.IsNaN(s.PosX) {
		s.PosX = 0
	}

	if math.IsNaN(s.PosY) {
		s.PosY = 0
	}

	if math.IsNaN(s.VelX) {
		s.VelX = 0
	}

	if math.IsNaN(s.VelY) {
		s.VelY = 0
	}

	if math.IsNaN(s.Shield) {
		s.Shield = 0
	}

	if math.IsNaN(s.Armor) {
		s.Armor = 0
	}

	if math.IsNaN(s.Hull) {
		s.Hull = 0
	}

	if math.IsNaN(s.Fuel) {
		s.Fuel = 0
	}

	if math.IsNaN(s.Heat) {
		s.Heat = 0
	}

	if math.IsNaN(s.Energy) {
		s.Energy = 0
	}

	if math.IsNaN(s.Theta) {
		s.Theta = 0
	}
}

// Determines whether or not a ship is considered cloaked for a tick
func (s *Ship) updateCloaking() {
	// aggregate cloaking percentage
	cloaked := false
	cloakPercentage := 0.0

	for _, e := range s.TemporaryModifiers {
		if e.Attribute == "cloak" {
			cloakPercentage += e.Percentage
		}
	}

	// store cloak percentage sum
	s.SumCloaking = cloakPercentage

	// determine whether cloaked for tick
	if cloakPercentage >= 1 {
		// ship is cloaked
		cloaked = true
	} else if cloakPercentage > 0 {
		// ship is intermittently cloaked
		r := rand.Float64()

		if r <= cloakPercentage {
			cloaked = true
		}
	}

	s.IsCloaked = cloaked
}

// Determines veil damage absorption for a tick
func (s *Ship) updateVeiling() {
	// check for active omni hardeners
	genericHardening := 0.0

	for _, x := range s.TemporaryModifiers {
		if x.RemainingTicks > 0 && x.Attribute == "veil" {
			genericHardening += x.Percentage
		}
	}

	// clamp active omni hardening
	if genericHardening > 0.99 {
		genericHardening = 0.99
	}

	if genericHardening < 0 {
		genericHardening = 0
	}

	// store factor
	s.SumVeiling = genericHardening
}

// Updates the ship's energy level for a tick
func (s *Ship) updateEnergy() {
	// determine whether to recalculate max energy
	re := s.CurrentSystem.tickCounter%14 == 0

	// get max energy
	maxEnergy := s.GetRealMaxEnergy(re)

	// calculate scaler for energy regen based on current energy percentage
	x := math.Abs(s.Energy / maxEnergy)
	u := math.Pow(ShipMinEnergyRegenPercent, x)

	// determine whether to recalculate energy regen
	rme := s.CurrentSystem.tickCounter%11 == 0

	// get energy regen amount for tick taking energy percentage scaling into account
	tickRegen := ((s.GetRealEnergyRegen(rme) / 1000) * Heartbeat) * u

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

	// determine whether to recalculate max fuel
	rf := s.CurrentSystem.tickCounter%34 == 0

	maxFuel := s.GetRealMaxFuel(rf)

	// clamp fuel
	if s.Fuel > maxFuel {
		s.Fuel = maxFuel
	}
}

// Updates the ship's shield level for a tick
func (s *Ship) updateShield() {
	// determine whether to recalculate max shield
	rs := s.CurrentSystem.tickCounter%15 == 0

	// get max shield
	max := s.GetRealMaxShield(rs)

	// calculate scaler for shield regen based on current shield percentage
	x := math.Abs(s.Shield / max)
	u := math.Pow(ShipMinShieldRegenPercent, 1.0-x)

	// determine whether to recalculate shield regen
	rme := s.CurrentSystem.tickCounter%11 == 0

	// get shield regen amount for tick taking shield percentage scaling into account
	tickRegen := ((s.GetRealShieldRegen(rme) / 1000) * Heartbeat) * u

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

	// determine whether to recalculate max armor
	ra := s.CurrentSystem.tickCounter%16 == 0

	// get max armor
	maxArmor := s.GetRealMaxArmor(ra)

	// clamp armor
	if s.Armor > maxArmor {
		s.Armor = maxArmor
	}

	// clamp hull
	if s.Hull < 0 {
		s.Hull = 0
	}

	// determine whether to recalculate max hull
	rh := s.CurrentSystem.tickCounter%17 == 0

	// get max hull
	maxHull := s.GetRealMaxHull(rh)

	// clamp hull
	if s.Hull > maxHull {
		s.Hull = maxHull
	}
}

// Updates the ship's heat level for a tick
func (s *Ship) updateHeat() {
	// determine whether to recalculate heat cap amount
	rhc := s.CurrentSystem.tickCounter%13 == 0

	// get max heat
	maxHeat := s.GetRealMaxHeat(rhc)

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
		// determine whether to recalculate heat sink amount
		rc := s.CurrentSystem.tickCounter%12 == 0

		// dissipate heat taking efficiency modifier into account
		s.Heat -= ((s.GetRealHeatSink(rc) / 1000) * Heartbeat) * u
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
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// stop autopilot
	s.AutopilotMode = SharedAutopilotRegistry.None

	// reset autopilot parameters
	s.AutopilotManualNav = ManualNavData{}
	s.AutopilotGoto = GotoData{}
	s.AutopilotOrbit = OrbitData{}
	s.AutopilotDock = DockData{}
	s.AutopilotUndock = UndockData{}
	s.AutopilotFight = FightData{}
	s.AutopilotMine = MineData{}
	s.AutopilotSalvage = SalvageData{}

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
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
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
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
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
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
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
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
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
	maxBay := s.GetRealCargoBayVolume(true)

	if usedBay > maxBay {
		return
	}

	// get registry
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// stash dock and activate autopilot
	s.AutopilotUndock = UndockData{}

	s.AutopilotMode = registry.Undock
}

// Invokes fight autopilot on the ship
func (s *Ship) CmdFight(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// stash fight and activate autopilot
	s.AutopilotFight = FightData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.AutopilotMode = registry.Fight
}

// Invokes mine autopilot on the ship
func (s *Ship) CmdMine(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// stash mine and activate autopilot
	s.AutopilotMine = MineData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.aiIncompatibleOreFault = false
	s.aiNoOrePulledFault = false

	s.AutopilotMode = registry.Mine
}

// Invokes salvage autopilot on the ship
func (s *Ship) CmdSalvage(targetID uuid.UUID, targetType int, lock bool) {
	// get registry
	registry := SharedAutopilotRegistry

	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// stash salvage and activate autopilot
	s.AutopilotSalvage = SalvageData{
		TargetID: targetID,
		Type:     targetType,
	}

	s.aiNoWreckFault = false

	s.AutopilotMode = registry.Salvage
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
		// remax selected stats
		s.Shield = s.GetRealMaxShield(true)
		s.Armor = s.GetRealMaxArmor(true)
		s.Hull = s.GetRealMaxHull(true)
		s.Fuel = s.GetRealMaxFuel(true)
		s.Energy = s.GetRealMaxEnergy(true)
		s.ReMaxDirty = false

		// recalculate cached stats
		RecalcAllStatCaches(s)
	}
}

// Returns the real acceleration capability of a ship after modifiers
func (s *Ship) GetRealAccel(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedRealAccel
	}

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

	// calculate and cache real accel
	s.CachedRealAccel = s.TemplateData.BaseAccel * s.FlightExperienceModifier * tpm

	// return result
	return s.CachedRealAccel
}

// Returns the real drag felt by a ship after modifiers
func (s *Ship) GetRealSpaceDrag(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedRealSpaceDrag
	}

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

	// calculate and cache true drag
	s.CachedRealSpaceDrag = SpaceDrag * tpm

	// return result
	return s.CachedRealSpaceDrag
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
func (s *Ship) GetRealTurn(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedRealTurn
	}

	// calculate real turning
	a := s.TemplateData.BaseTurn * s.FlightExperienceModifier

	// store in cache
	s.CachedRealTurn = a

	// return result
	return s.CachedRealTurn
}

// Returns the real mass of a ship after modifiers
func (s *Ship) GetRealMass() float64 {
	return s.TemplateData.BaseMass
}

// Returns the real max shield of the ship after modifiers
func (s *Ship) GetRealMaxShield(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedMaxShield
	}

	// calculate max shield
	a := s.TemplateData.BaseShield * s.FlightExperienceModifier

	// store in cache
	s.CachedMaxShield = a

	// return result
	return s.CachedMaxShield
}

// Returns the real shield regen rate after modifiers
func (s *Ship) GetRealShieldRegen(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedShieldRegen
	}

	// get base shield regen
	a := s.TemplateData.BaseShieldRegen * s.FlightExperienceModifier

	// apply temporary modifiers
	for _, e := range s.TemporaryModifiers {
		// regeneration mask (does not accumulate)
		if e.Attribute == "regeneration_mask" {
			a *= (1.0 - e.Percentage)
		}
	}

	// store in cache
	s.CachedShieldRegen = a

	// return result
	return s.CachedShieldRegen
}

// Returns the real max armor of the ship after modifiers
func (s *Ship) GetRealMaxArmor(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedMaxArmor
	}

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

	// store result in cache
	s.CachedMaxArmor = a

	// return result
	return s.CachedMaxArmor
}

// Returns the real max hull of the ship after modifiers
func (s *Ship) GetRealMaxHull(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedMaxHull
	}

	// recalculate max hull
	a := s.TemplateData.BaseHull * s.FlightExperienceModifier

	// store result in cache
	s.CachedMaxHull = a

	// return result
	return s.CachedMaxHull
}

// Returns the real max energy of the ship after modifiers
func (s *Ship) GetRealMaxEnergy(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedMaxEnergy
	}

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

	// store result in cache
	s.CachedMaxEnergy = a

	// return result
	return s.CachedMaxEnergy
}

// Returns the real energy regeneration rate of the ship after modifiers
func (s *Ship) GetRealEnergyRegen(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedEnergyRegen
	}

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

	// apply temporary modifiers
	for _, e := range s.TemporaryModifiers {
		// regeneration mask (does not accumulate)
		if e.Attribute == "regeneration_mask" {
			a *= (1.0 - e.Percentage)
		}
	}

	// store in cache
	s.CachedEnergyRegen = a

	// return result
	return s.CachedEnergyRegen
}

// Returns the real max heat of the ship after modifiers
func (s *Ship) GetRealMaxHeat(recalculate bool) float64 {
	if !recalculate {
		return s.CachedMaxHeat
	}

	// get base heat cap
	a := s.TemplateData.BaseHeatCap * s.FlightExperienceModifier

	// add bonuses from passive modules in rack c
	for _, e := range s.Fitting.CRack {
		heatCapAdd, s := e.ItemMeta.GetFloat64("heat_cap_max_add")

		if s {
			// include in real max
			a += heatCapAdd
		}
	}

	// calculate and cache final heat cap
	s.CachedMaxHeat = a

	// return result
	return s.CachedMaxHeat
}

// Returns the real heat dissipation rate of the ship after modifiers
func (s *Ship) GetRealHeatSink(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedHeatSink
	}

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
		// heat sink
		if e.Attribute == "heat_sink" {
			tpm += e.Percentage
		}

		// dissipation mask (does not accumulate)
		if e.Attribute == "dissipation_mask" {
			a *= (1.0 - e.Percentage)
		}
	}

	// floor percentage modifier at 0
	if tpm < 0 {
		tpm = 0
	}

	// calculate and cache final heat sink
	s.CachedHeatSink = a * tpm

	// return result
	return s.CachedHeatSink
}

// Returns the real max fuel of the ship after modifiers
func (s *Ship) GetRealMaxFuel(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedMaxFuel
	}

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

	// cache final max fuel
	s.CachedMaxFuel = f

	// return result
	return s.CachedMaxFuel
}

// Returns the real max cargo bay volume of the ship after modifiers
func (s *Ship) GetRealCargoBayVolume(recalculate bool) float64 {
	// return cache if no recalculation
	if !recalculate {
		return s.CachedCargoBayVolume
	}

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

	// store result in cache
	s.CachedCargoBayVolume = cv

	// return result
	return s.CachedCargoBayVolume
}

// Returns the total amount of cargo bay space currently in use
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
		if i.Quantity == 0 {
			continue
		}

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
func (s *Ship) dealDamage(
	shieldDmg float64,
	armorDmg float64,
	hullDmg float64,
	attackerRS *shared.PlayerReputationSheet,
	attackerModule *FittedSlot,
) {
	// update aggression tables
	s.updateAggressionTables(
		shieldDmg,
		armorDmg,
		hullDmg,
		attackerRS,
		attackerModule,
	)

	// get generic hardening (veiling)
	genericHardening := s.SumVeiling

	// apply active omni hardeners
	shieldDmg *= 1.0 - genericHardening
	armorDmg *= 1.0 - genericHardening
	hullDmg *= 1.0 - genericHardening

	// apply shield damage
	s.Shield -= shieldDmg

	// clamp shield
	if s.Shield < 0 {
		s.Shield = 0
	}

	// determine shield percentage
	shieldP := s.Shield / s.GetRealMaxShield(false)

	// apply armor damage if shields below 25% scaling for remaining shields
	if shieldP < 0.25 {
		s.Armor -= armorDmg * (1 - shieldP)
	}

	// clamp armor
	if s.Armor < 0 {
		s.Armor = 0
	}

	// determine armor percentage
	armorP := s.Armor / s.GetRealMaxArmor(false)

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
func (s *Ship) siphonEnergy(
	maxSiphonAmount float64,
	attackerRS *shared.PlayerReputationSheet,
	attackerModule *FittedSlot,
) float64 {
	// update aggression tables
	s.updateAggressionTables(
		0,
		0,
		0,
		attackerRS,
		attackerModule,
	)

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

// Siphons heat from the ship, returns the actual amount siphoned
func (s *Ship) siphonHeat(
	maxSiphonAmount float64,
	assisterModule *FittedSlot,
) float64 {
	// limit amount to siphon so that the remaining amount is positive
	actualSiphon := 0.0

	if s.Heat-maxSiphonAmount >= 0 {
		actualSiphon = maxSiphonAmount
	} else {
		actualSiphon = maxSiphonAmount - s.Heat
	}

	// apply siphon
	s.Heat -= actualSiphon

	// clamp heat
	if s.Heat < 0 {
		s.Heat = 0
	}

	// return amount siphoned
	return actualSiphon
}

// Receives energy and returns the excess that could not be stored
func (s *Ship) receiveEnergy(
	maxRecAmount float64,
	assisterModule *FittedSlot,
) float64 {
	// limit amount to receive so that max is not exceeded
	actualReceived := 0.0

	if s.Energy+maxRecAmount <= s.CachedMaxEnergy {
		actualReceived = maxRecAmount
	} else {
		actualReceived = s.CachedMaxEnergy - s.Energy
	}

	// apply transfer
	s.Energy += actualReceived

	// return actual amount received
	return actualReceived
}

// Receives shield and returns the excess that could not be stored
func (s *Ship) receiveShield(
	maxRecAmount float64,
	assisterModule *FittedSlot,
) float64 {
	// limit amount to receive so that max is not exceeded
	actualReceived := 0.0

	if s.Shield+maxRecAmount <= s.CachedMaxShield {
		actualReceived = maxRecAmount
	} else {
		actualReceived = s.CachedMaxShield - s.Shield
	}

	// apply transfer
	s.Shield += actualReceived

	// return actual amount received
	return actualReceived
}

// Given a faction to compare against, returns the standing and whether they have declared open hostilities
func (s *Ship) checkStandings(factionID uuid.UUID) (float64, bool) {
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
	s.ReputationSheet.Lock.Lock()
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
		ms := s.GetRealMaxShield(false)
		sr := s.Shield / ms

		if sr < 0.75 {
			s.behaviourPatrol()
		} else {
			// get registry
			registry := SharedBehaviourRegistry

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
			case registry.PatchMine:
				s.behaviourPatchMine()
			case registry.PatchSalvage:
				s.behaviourPatchSalvage()
			}
		}
	}
}

// wanders around the universe aimlessly
func (s *Ship) behaviourWander() {
	// pause if heat too high
	maxHeat := s.GetRealMaxHeat(false)
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := SharedAutopilotRegistry

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil && s.DockedAtStation != nil {
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
	autoReg := SharedAutopilotRegistry

	// get heat level
	maxHeat := s.GetRealMaxHeat(false)
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
		if s.DockedAtStationID != nil && s.DockedAtStation != nil {
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
			tgtReg := models.SharedTargetTypeRegistry

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
				standing, openlyHostile := sx.checkStandings(s.FactionID)

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
	maxHeat := s.GetRealMaxHeat(false)
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := SharedAutopilotRegistry

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil && s.DockedAtStation != nil {
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
								q := physics.RandInRange(1, i.Quantity+2)

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
				cv := s.GetRealCargoBayVolume(false)
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

// wanders around the universe randomly mining and selling what it mined to patch the economy
func (s *Ship) behaviourPatchMine() {
	// pause if heat too high
	maxHeat := s.GetRealMaxHeat(false)
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := SharedAutopilotRegistry

	// check for faults
	if s.aiIncompatibleOreFault {
		// stop autopilot
		s.CmdAbort(false)

		// clear flag
		s.aiIncompatibleOreFault = false

		// wander
		s.gotoNextWanderDestination(15)
		return
	}

	if s.aiNoOrePulledFault {
		// stop autopilot
		s.CmdAbort(false)

		// clear flag
		s.aiNoOrePulledFault = false

		// wander
		s.gotoNextWanderDestination(95)
		return
	}

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil && s.DockedAtStation != nil {
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

			// check if sell attempts should be made
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
								q := physics.RandInRange(1, i.Quantity+2)

								// try to sell item to silo
								s.SellItemToSilo(p.ID, i.ID, q, false)
							}
						}
					}
				}
			}
		} else {
			// check if cargo bay is almost full (>80%)
			max := s.GetRealCargoBayVolume(false)
			used := s.TotalCargoBayVolumeUsed(false)

			if used/max > 0.8 {
				// go somewhere to try and sell it
				s.gotoNextWanderDestination(85)
			} else {
				// get and count asteroids in system
				asteroids := s.CurrentSystem.asteroids
				count := len(asteroids)

				// verify there are candidates
				if count == 0 {
					// no asteroids here? wander
					s.gotoNextWanderDestination(15)

					return
				}

				// pick random asteroid to mine
				tgt := physics.RandInRange(0, count)
				var tgtAst *Asteroid = nil

				idx := 0
				for _, v := range asteroids {
					if idx == tgt {
						tgtAst = v
						break
					}

					idx++
				}

				if tgtAst != nil {
					// go mine it
					s.CmdMine(tgtAst.ID, models.SharedTargetTypeRegistry.Asteroid, false)
				} else {
					// no asteroids here? wander
					s.gotoNextWanderDestination(15)
				}
			}
		}
	}
}

// wanders around the universe randomly salaving wrecks and selling what it salvaged to patch the economy
func (s *Ship) behaviourPatchSalvage() {
	// pause if heat too high
	maxHeat := s.GetRealMaxHeat(false)
	heatLevel := s.Heat / maxHeat

	if heatLevel > 0.95 {
		s.CmdAbort(false)
	}

	// get registry
	autoReg := SharedAutopilotRegistry

	if s.aiNoWreckFault {
		// stop autopilot
		s.CmdAbort(false)

		// clear flag
		s.aiNoWreckFault = false

		// wander
		s.gotoNextWanderDestination(95)
		return
	}

	// check if idle
	if s.AutopilotMode == autoReg.None {
		// allow time to cool off :)
		if heatLevel > 0.25 {
			return
		}

		// check if docked
		if s.DockedAtStationID != nil && s.DockedAtStation != nil {
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

			// check if sell attempts should be made
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
								q := physics.RandInRange(1, i.Quantity+2)

								// try to sell item to silo
								s.SellItemToSilo(p.ID, i.ID, q, false)
							}
						}
					}
				}
			}
		} else {
			// check if cargo bay is almost full (>80%)
			max := s.GetRealCargoBayVolume(false)
			used := s.TotalCargoBayVolumeUsed(false)

			if used/max > 0.8 {
				// go somewhere to try and sell it
				s.gotoNextWanderDestination(85)
			} else {
				// get and count wrecks in system
				wrecks := s.CurrentSystem.wrecks
				count := len(wrecks)

				// verify there are candidates
				if count == 0 {
					// no wrecks here? wander
					s.gotoNextWanderDestination(15)

					return
				}

				// pick random asteroid to mine
				tgt := physics.RandInRange(0, count)
				var tgtWre *Wreck = nil

				idx := 0
				for _, v := range wrecks {
					if idx == tgt {
						tgtWre = v
						break
					}

					idx++
				}

				if tgtWre != nil {
					// go salvage it
					s.CmdSalvage(tgtWre.ID, models.SharedTargetTypeRegistry.Wreck, false)
				} else {
					// no wrecks here? wander
					s.gotoNextWanderDestination(15)
				}
			}
		}
	}
}

// helper for behaviour routines that need to wander around the universe
func (s *Ship) gotoNextWanderDestination(stationDockChance int) {
	// get registry
	tgtReg := models.SharedTargetTypeRegistry

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
				v, oh := s.checkStandings(e.FactionID)

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
	registry := SharedAutopilotRegistry

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
	case registry.Mine:
		s.doAutopilotMine()
	case registry.Salvage:
		s.doAutopilotSalvage()
	}
}

// Flies the ship automatically when docked
func (s *Ship) doDockedAutopilot() {
	// get registry
	registry := SharedAutopilotRegistry

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
		s.rotate(turnMag / s.GetRealTurn(false))
	} else if a < 0 {
		s.rotate(turnMag / -s.GetRealTurn(false))
	}

	// thrust forward
	s.forwardThrust(s.AutopilotManualNav.Magnitude)

	// determine whether to recalculate real drag
	rc := s.CurrentSystem.tickCounter%8 == 0

	// decrease magnitude (this is to allow this to expire and require another move order from the player)
	s.AutopilotManualNav.Magnitude -= s.AutopilotManualNav.Magnitude * s.GetRealSpaceDrag(rc)

	// stop when magnitude is low
	if s.AutopilotManualNav.Magnitude < 0.0001 {
		s.AutopilotMode = SharedAutopilotRegistry.None
	}
}

// Causes ship to turn to move towards a target and stop when within range
func (s *Ship) doAutopilotGoto() {
	// get registry
	targetTypeReg := models.SharedTargetTypeRegistry

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
	} else if s.AutopilotGoto.Type == targetTypeReg.Outpost {
		// find outpost with id
		tgt := s.CurrentSystem.outposts[s.AutopilotGoto.TargetID.String()]

		if tgt == nil {
			s.CmdAbort(false)
			return
		}

		// store target details
		tX = tgt.PosX
		tY = tgt.PosY
		tR = tgt.TemplateData.Radius
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
	} else if s.AutopilotGoto.Type == targetTypeReg.Wreck {
		// find wreck with id
		tgt := s.CurrentSystem.wrecks[s.AutopilotGoto.TargetID.String()]

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
	targetTypeReg := models.SharedTargetTypeRegistry

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
	} else if s.AutopilotOrbit.Type == targetTypeReg.Outpost {
		// find outpost with id
		tgt := s.CurrentSystem.outposts[s.AutopilotOrbit.TargetID.String()]

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
	} else if s.AutopilotOrbit.Type == targetTypeReg.Wreck {
		// find wreck with id
		tgt := s.CurrentSystem.wrecks[s.AutopilotOrbit.TargetID.String()]

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
	// disarm leave faction
	s.LeaveFactionArmed = false

	// get registry
	targetTypeReg := models.SharedTargetTypeRegistry

	// check if dockable
	isDockable := s.AutopilotDock.Type == targetTypeReg.Station || s.AutopilotDock.Type == targetTypeReg.Outpost

	if isDockable {
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
			s.AutopilotMode = SharedAutopilotRegistry.None
		}
	} else {
		s.CmdAbort(false)
		return
	}
}

// Causes ship to undock from a target
func (s *Ship) doAutopilotUndock() {
	// disarm leave faction
	s.LeaveFactionArmed = false

	// verify that we are docked (currently only supports stations)
	if s.DockedAtStationID != nil && s.DockedAtStation != nil {
		// remove references
		s.DockedAtStationID = nil
		s.DockedAtStation = nil

		// not docked - cancel autopilot
		s.AutopilotMode = SharedAutopilotRegistry.None
	} else {
		// not docked - cancel autopilot
		s.AutopilotMode = SharedAutopilotRegistry.None
	}
}

// Causes ship to fight with a target
func (s *Ship) doAutopilotFight() {
	// get registry
	targetTypeReg := models.SharedTargetTypeRegistry

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
			maxHeat := s.GetRealMaxHeat(false)
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
			maxHeat := s.GetRealMaxHeat(false)
			heatAdd := 0.0

			for i, v := range s.Fitting.BRack {
				// special check for shield boosters
				if v.ItemTypeFamily == "shield_booster" {
					// make sure enough shield has been lost for this to be worth it
					shieldBoost, _ := v.ItemMeta.GetFloat64("shield_boost_amount")
					maxShield := s.GetRealMaxShield(false)

					if s.Shield+shieldBoost >= 0.75*maxShield {
						// deactivate and continue
						s.Fitting.BRack[i].WillRepeat = false
						continue
					}
				}

				// special check for burst fusion reactors
				if v.ItemTypeFamily == "burst_reactor" {
					// make sure energy is below 15% for this to be worth it
					if s.Energy/s.GetRealMaxEnergy(false) >= 0.15 {
						// deactivate and continue
						s.Fitting.BRack[i].WillRepeat = false
						continue
					}
				}

				// get heat
				h, _ := v.ItemMeta.GetFloat64("activation_heat")

				// get heat ratio
				hr := (s.Heat + heatAdd) / maxHeat

				// special check for engine overchargers
				if v.ItemTypeFamily == "eng_oc" {
					// do not activate if heat is too high
					if hr > 0.25 {
						// deactivate and continue
						s.Fitting.BRack[i].WillRepeat = false
						continue
					}
				}

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

// Causes ship to mine a target
func (s *Ship) doAutopilotMine() {
	// get registry
	targetTypeReg := models.SharedTargetTypeRegistry

	if s.AutopilotMine.Type == targetTypeReg.Asteroid {
		// find asteroid
		targetAsteroid := s.CurrentSystem.asteroids[s.AutopilotMine.TargetID.String()]

		if targetAsteroid == nil {
			s.CmdAbort(false)
			return
		}

		// determine type of asteroid
		isOre := targetAsteroid.ItemFamilyID == "ore"
		isIce := targetAsteroid.ItemFamilyID == "ice"

		// use average mining range to determine stand-off distance (this can be improved a lot with more specific mining routines)
		totalRange := 0.0
		rangedMods := 0

		for _, m := range s.Fitting.ARack {
			// only take appropriate equipment into account
			if isOre {
				f, c := m.ItemMeta.GetBool("can_mine_ore")

				if !f || !c {
					continue
				}
			}

			if isIce {
				f, c := m.ItemMeta.GetBool("can_mine_ice")

				if !f || !c {
					continue
				}
			}

			// accumulate range
			r, f := m.ItemMeta.GetFloat64("range")

			if f {
				totalRange += r
				rangedMods += 1
			}
		}

		if rangedMods == 0 {
			// can't mine this type
			s.CmdAbort(false)
			s.aiIncompatibleOreFault = true

			return
		}

		avgRange := totalRange / (float64(rangedMods) + Epsilon)
		standOff := int(avgRange / 2.0)

		// fill autopilot data
		s.AutopilotGoto = GotoData{
			TargetID: s.AutopilotMine.TargetID,
			Type:     s.AutopilotMine.Type,
			Hold:     &standOff,
		}

		// reuse orbit autopilot routine to keep distance with target
		s.doAutopilotGoto()

		if s.CurrentSystem.tickCounter%45 == 0 {
			// try to activate rack A mining modules
			maxHeat := s.GetRealMaxHeat(false)
			heatAdd := 0.0

			for i, v := range s.Fitting.ARack {
				if isOre {
					f, c := v.ItemMeta.GetBool("can_mine_ore")

					if !f || !c {
						continue
					}
				}

				if isIce {
					f, c := v.ItemMeta.GetBool("can_mine_ice")

					if !f || !c {
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
					s.Fitting.ARack[i].TargetID = &s.AutopilotMine.TargetID
					s.Fitting.ARack[i].TargetType = &s.AutopilotMine.Type
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
			// check if cargo bay is almost full (>80%)
			max := s.GetRealCargoBayVolume(false)
			used := s.TotalCargoBayVolumeUsed(false)

			if used/max > 0.8 {
				// stop mining
				s.CmdAbort(false)
			}
		}
	} else {
		s.CmdAbort(false)
		return
	}
}

// Causes ship to salvage a target
func (s *Ship) doAutopilotSalvage() {
	// get registry
	targetTypeReg := models.SharedTargetTypeRegistry

	if s.AutopilotSalvage.Type == targetTypeReg.Wreck {
		// find wreck
		targetWreck := s.CurrentSystem.wrecks[s.AutopilotSalvage.TargetID.String()]

		if targetWreck == nil {
			s.CmdAbort(false)
			return
		}

		// use average salvage range to determine stand-off distance (this can be improved a lot with more specific salvaging routines)
		totalRange := 0.0
		rangedMods := 0

		for _, m := range s.Fitting.ARack {
			// only take appropriate equipment into account
			if m.ItemTypeFamily != "salvager" {
				continue
			}

			// accumulate range
			r, f := m.ItemMeta.GetFloat64("range")

			if f {
				totalRange += r
				rangedMods += 1
			}
		}

		if rangedMods == 0 {
			// can't salvage
			s.CmdAbort(false)

			return
		}

		avgRange := totalRange / (float64(rangedMods) + Epsilon)
		standOff := int(avgRange / 2.0)

		// fill autopilot data
		s.AutopilotGoto = GotoData{
			TargetID: s.AutopilotSalvage.TargetID,
			Type:     s.AutopilotSalvage.Type,
			Hold:     &standOff,
		}

		// reuse orbit autopilot routine to keep distance with target
		s.doAutopilotGoto()

		if s.CurrentSystem.tickCounter%45 == 0 {
			// try to activate rack A salvagers
			maxHeat := s.GetRealMaxHeat(false)
			heatAdd := 0.0

			for i, v := range s.Fitting.ARack {
				// only take appropriate equipment into account
				if v.ItemTypeFamily != "salvager" {
					continue
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
					s.Fitting.ARack[i].TargetID = &s.AutopilotSalvage.TargetID
					s.Fitting.ARack[i].TargetType = &s.AutopilotSalvage.Type
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
			// check if cargo bay is almost full (>80%)
			max := s.GetRealCargoBayVolume(false)
			used := s.TotalCargoBayVolumeUsed(false)

			if used/max > 0.8 {
				// stop salvaging
				s.CmdAbort(false)
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

	// determine whether to recalculate real accel + drag + turning
	rc := s.CurrentSystem.tickCounter%8 == 0

	// chance to recalculate turning
	s.GetRealTurn(rc)

	// determine whether to thrust forward and by how much
	scale := ((s.GetRealAccel(rc) * (caution / s.GetRealSpaceDrag(rc))) / 0.175)
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
		s.rotate(turnMag / s.GetRealTurn(false))
	} else if a < 0 {
		s.rotate(turnMag / -s.GetRealTurn(false))
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
	burn := s.GetRealTurn(false) * scale

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

	// determine whether to recalculate real accel + drag
	rc := s.CurrentSystem.tickCounter%8 == 0

	// calculate burn
	burn := s.GetRealAccel(rc) * scale

	// account for additional drag effects
	drag := s.GetRealSpaceDrag(rc)
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

// Returns an item in the ship's cargo bay if it is present and clean
func (s *Ship) FindItemInCargo(id uuid.UUID) *Item {
	// look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		// skip if unclean
		if item.CoreDirty {
			continue
		}

		if item.SchematicInUse {
			continue
		}

		if item.Quantity <= 0 {
			continue
		}

		if item.ID == id {
			// return item
			return item
		}
	}

	// nothing found
	return nil
}

// Returns the first available, clean, item of a given type and packaging state in the cargo bay
func (s *Ship) FindFirstAvailableItemOfTypeInCargo(typeID uuid.UUID, packaged bool) *Item {
	// look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		if item.CoreDirty {
			continue
		}

		if item.SchematicInUse {
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

// Returns the first available, clean, packaged item of a given type and stack size in the cargo bay
func (s *Ship) FindFirstAvailablePackagedStackOfSizeInCargo(typeID uuid.UUID, size int) *Item {
	// look for item
	for i := range s.CargoBay.Items {
		item := s.CargoBay.Items[i]

		if item.CoreDirty {
			continue
		}

		if item.SchematicInUse {
			continue
		}

		if item.Quantity <= 0 {
			continue
		}

		if item.Quantity < size {
			continue
		}

		if item.ItemTypeID == typeID && item.IsPackaged {
			// return item
			return item
		}
	}

	// nothing found
	return nil
}

// Removes all items with a zero quantity from the cargo of the ship
func (s *Ship) removeZeroQuantityItemsFromCargo() {
	// lock container
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	// remove zero-quantity items from cargo bay
	newCB := make([]*Item, 0)

	for _, i := range s.CargoBay.Items {
		if i.Quantity > 0 {
			// keep in cargo bay
			newCB = append(newCB, i)
		}
	}

	s.CargoBay.Items = newCB
}

// Removes an item from the cargo hold and fits it to the ship
func (s *Ship) FitModule(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock containers
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	s.FittingBay.Lock.Lock()
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
	v, _ := item.Meta.GetFloat64("volume")

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
			s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, i)
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

	// recalculate cached stats
	RecalcAllStatCaches(s)

	// success!
	return nil
}

// Removes a module from a ship and places it in the cargo hold
func (s *Ship) UnfitModule(m *FittedSlot, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock containers
	m.shipMountedOn.CargoBay.Lock.Lock()
	defer m.shipMountedOn.CargoBay.Lock.Unlock()

	m.shipMountedOn.FittingBay.Lock.Lock()
	defer m.shipMountedOn.FittingBay.Lock.Unlock()

	// make sure ship is docked
	if s.DockedAtStationID == nil {
		return errors.New("you must be docked to unfit a module")
	}

	// get module volume
	v, _ := m.ItemMeta.GetFloat64("volume")

	// make sure there is sufficient space in the cargo bay
	if s.TotalCargoBayVolumeUsed(lock)+v > s.GetRealCargoBayVolume(true) {
		return errors.New("insufficient room in cargo bay to unfit module")
	}

	// make sure the module is not cycling
	if m.IsCycling || m.WillRepeat {
		return errors.New("modules must be offline to be unfit")
	}

	// if the module is in rack c, make sure the ship is fully repaired
	if m.Rack == "C" {
		if s.Armor < s.GetRealMaxArmor(true) || s.Hull < s.GetRealMaxHull(true) {
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
		o.Lock.Lock()
		defer o.Lock.Unlock()

		// skip if not this module
		if o.ID != m.ItemID {
			newFB = append(newFB, o)
		} else {
			// move to cargo bay if there is still room
			if s.TotalCargoBayVolumeUsed(lock)+v <= s.GetRealCargoBayVolume(true) {
				m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, o)
				o.ContainerID = m.shipMountedOn.CargoBayContainerID

				// escalate to core to save to db
				s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, o)
			} else {
				return errors.New("insufficient room in cargo bay to unfit module")
			}
		}
	}

	s.FittingBay.Items = newFB

	// recalculate cached stats
	RecalcAllStatCaches(s)

	// success!
	return nil
}

// Trashes an item in the ship's cargo bay if it exists
func (s *Ship) TrashItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
			s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, o)
		}
	}

	s.CargoBay.Items = newCB

	return nil
}

// Removes an item from the ship's cargo bay if it exists, without updating its container id
func (s *Ship) RemoveItemFromCargo(id uuid.UUID, lock bool) (*Item, error) {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	// null check
	if item == nil {
		return errors.New("item must not be null")
	}

	// make sure there is enough space
	used := s.TotalCargoBayVolumeUsed(lock)
	max := s.GetRealCargoBayVolume(true)

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
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	s.CurrentSystem.PackagedItems = append(s.CurrentSystem.PackagedItems, item)

	return nil
}

// Attempts to buy an output from a silo and places it in cargo if successful
func (s *Ship) BuyItemFromSilo(siloID uuid.UUID, itemTypeID uuid.UUID, quantity int, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
		silo.Lock.Lock()
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
		maxBay := s.GetRealCargoBayVolume(true)

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
			ID:             nid,
			ItemTypeID:     output.ItemTypeID,
			Meta:           output.ItemTypeMeta,
			Created:        time.Now(),
			CreatedBy:      &s.UserID,
			CreatedReason:  "Bought from silo",
			ContainerID:    s.CargoBayContainerID,
			Quantity:       quantity,
			IsPackaged:     true,
			Lock:           sync.Mutex{},
			ItemTypeName:   output.ItemTypeName,
			ItemFamilyID:   output.ItemFamilyID,
			ItemFamilyName: output.ItemFamilyName,
			ItemTypeMeta:   output.ItemTypeMeta,
			CoreDirty:      true,
		}

		// escalate order save request to core
		s.CurrentSystem.NewItems = append(s.CurrentSystem.NewItems, &newItem)

		// place item in cargo bay
		s.CargoBay.Items = append(s.CargoBay.Items, &newItem)
	} else {
		// parse template id
		stID, err := uuid.Parse(stIDStr)

		if err != nil {
			return err
		}

		// request a new ship to be generated from this purchase
		r := NewShipTicket{
			UserID:         s.UserID,
			ShipTemplateID: stID,
			StationID:      *s.DockedAtStationID,
		}

		c, ok := s.CurrentSystem.clients[s.UserID.String()]

		if ok {
			r.Client = c
		}

		// escalate order save request to core
		s.CurrentSystem.NewShipTickets = append(s.CurrentSystem.NewShipTickets, &r)
	}

	// log buy to console if player
	if !s.IsNPC {
		bm := 0

		if s.BehaviourMode != nil {
			bm = *s.BehaviourMode
		}

		shared.TeeLog(
			fmt.Sprintf(
				"[%v] %v (%v::%v) bought %v %v from %v (silo)",
				s.CurrentSystem.SystemName,
				s.CharacterName,
				s.Texture,
				bm,
				quantity,
				output.ItemTypeName,
				s.DockedAtStation.StationName,
			),
		)
	}

	// success
	return nil
}

// Attempts to sell an item in the cargo bay to a silo and removes it if successful
func (s *Ship) SellItemToSilo(siloID uuid.UUID, itemId uuid.UUID, quantity int, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
		item.Lock.Lock()
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
	s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)

	// log sell to console if player
	if !s.IsNPC {
		bm := 0

		if s.BehaviourMode != nil {
			bm = *s.BehaviourMode
		}

		shared.TeeLog(
			fmt.Sprintf(
				"[%v] %v (%v::%v) sold %v %v to %v (silo)",
				s.CurrentSystem.SystemName,
				s.CharacterName,
				s.Texture,
				bm,
				quantity,
				item.ItemTypeName,
				s.DockedAtStation.StationName,
			),
		)
	}

	// success
	return nil
}

// Attempts to fulfill a sell order and place the item in cargo if successful
func (s *Ship) BuyItemFromOrder(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
		order.Lock.Lock()
		defer order.Lock.Unlock()

		order.Item.Lock.Lock()
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
		maxBay := s.GetRealCargoBayVolume(true)

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
		seller.Lock.Lock()
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
	s.CurrentSystem.BoughtSellOrders = append(s.CurrentSystem.BoughtSellOrders, order)

	if order.Item.ItemFamilyID != "ship" {
		// mark item as dirty and place it in the cargo container
		order.Item.CoreDirty = true
		order.Item.ContainerID = s.CargoBayContainerID

		// escalate item for saving in db
		s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, order.Item)

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

		s.CurrentSystem.UsedShipPurchases = append(s.CurrentSystem.UsedShipPurchases, &r)

		// mark stub item as dirty and place it in the trash container
		order.Item.CoreDirty = true
		order.Item.ContainerID = s.TrashContainerID

		// escalate stub item for saving in db
		s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, order.Item)
	}

	// log buy to console if player
	if !s.IsNPC {
		bm := 0

		if s.BehaviourMode != nil {
			bm = *s.BehaviourMode
		}

		shared.TeeLog(
			fmt.Sprintf(
				"[%v] %v (%v::%v) bought %v %v at %v (order)",
				s.CurrentSystem.SystemName,
				s.CharacterName,
				s.Texture,
				bm,
				order.Item.Quantity,
				order.Item.ItemTypeName,
				s.DockedAtStation.StationName,
			),
		)
	}

	// success
	return nil
}

// Lists this ship on the station order exchange
func (s *Ship) SellSelfAsOrder(price float64, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
		Lock:          sync.Mutex{},
	}

	// escalate ship and stub item for saving in db
	s.CurrentSystem.NewItems = append(s.CurrentSystem.NewItems, &stub)
	s.CurrentSystem.SetNoLoad = append(s.CurrentSystem.SetNoLoad, &ShipNoLoadSet{
		ID:   s.ID,
		Flag: true,
	})

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
		Lock:         sync.Mutex{},
	}

	// link stub into sell order
	newOrder.Item = &stub
	newOrder.GetItemIDFromItem = true // we won't know the real item id until after it saves

	// escalate to core for saving in db
	s.CurrentSystem.NewSellOrders = append(s.CurrentSystem.NewSellOrders, &newOrder)

	// log listing to console if player
	if !s.IsNPC {
		bm := 0

		if s.BehaviourMode != nil {
			bm = *s.BehaviourMode
		}

		shared.TeeLog(
			fmt.Sprintf(
				"[%v] %v (%v::%v) listed itself for sale at %v (order)",
				s.CurrentSystem.SystemName,
				s.CharacterName,
				s.Texture,
				bm,
				s.DockedAtStation.StationName,
			),
		)
	}

	// success
	return nil
}

// Lists an item in the ship's cargo bay on the station order exchange if it exists
func (s *Ship) SellItemAsOrder(id uuid.UUID, price float64, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	s.CurrentSystem.MovedItems = append(s.CurrentSystem.MovedItems, item)

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
		Lock:         sync.Mutex{},
	}

	// link item into sell order
	newOrder.Item = item

	// escalate to core for saving in db
	s.CurrentSystem.NewSellOrders = append(s.CurrentSystem.NewSellOrders, &newOrder)

	// log listing to console if player
	if !s.IsNPC {
		bm := 0

		if s.BehaviourMode != nil {
			bm = *s.BehaviourMode
		}

		shared.TeeLog(
			fmt.Sprintf(
				"[%v] %v (%v::%v) listed %v %v at %v (order)",
				s.CurrentSystem.SystemName,
				s.CharacterName,
				s.Texture,
				bm,
				newOrder.Item.Quantity,
				newOrder.Item.ItemTypeName,
				s.DockedAtStation.StationName,
			),
		)
	}

	// success
	return nil
}

// Packages an item in the ship's cargo bay if it exists
func (s *Ship) UnpackageItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	s.CurrentSystem.UnpackagedItems = append(s.CurrentSystem.UnpackagedItems, item)

	return nil
}

// Stacks an item in the ship's cargo bay if it exists
func (s *Ship) StackItemInCargo(id uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
				// verify we won't overflow
				var testQ int64 = 0
				testQ += int64(o.Quantity) + int64(item.Quantity)

				if testQ > math.MaxInt32 {
					continue
				}

				// merge stacks
				q := item.Quantity

				o.Quantity += q
				item.Quantity -= q

				o.CoreDirty = true
				item.CoreDirty = true

				// escalate to core for saving in db
				s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)
				s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, o)

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
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)

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
		Lock:          sync.Mutex{}, ItemTypeName: item.ItemTypeName,
		ItemFamilyID:   item.ItemFamilyID,
		ItemFamilyName: item.ItemFamilyName,
		ItemTypeMeta:   item.ItemTypeMeta,
		CoreDirty:      true,
	}

	// escalate to core for saving in db
	s.CurrentSystem.NewItems = append(s.CurrentSystem.NewItems, &newItem)

	// add new item to cargo hold
	s.CargoBay.Items = append(s.CargoBay.Items, &newItem)

	return nil
}

func (s *Ship) SelfDestruct(lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
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
	s.dealDamage(s.GetRealMaxShield(false)+1, s.GetRealMaxArmor(false)+1, s.GetRealMaxHull(false)+1, nil, nil)

	return nil
}

// Attempts to consume a fuel pellet from cargo and add it to the fuel tank
func (s *Ship) ConsumeFuelFromCargo(itemID uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)

	return nil
}

// Attempts to consume a repair kit from cargo and apply it to health
func (s *Ship) ConsumeRepairKitFromCargo(itemID uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
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
	maxArmor := s.GetRealMaxArmor(true)
	maxHull := s.GetRealMaxHull(true)

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
	s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)

	return nil
}

// Attempts to consume an outpost kit from cargo and spawn a new outpost
func (s *Ship) ConsumeOutpostKitFromCargo(itemID uuid.UUID, lock bool) error {
	if lock {
		// lock entity
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// lock cargo bay
	s.CargoBay.Lock.Lock()
	defer s.CargoBay.Lock.Unlock()

	// make sure ship is undocked
	if s.DockedAtStationID != nil {
		return errors.New("you must be undocked to consume an outpost kit")
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

	// make sure item is an outpost kit
	if item.ItemFamilyID != "outpost_kit" {
		return errors.New("item is not an outpost kit")
	}

	// make sure item is unpackaged
	if item.IsPackaged {
		return errors.New("only unpackaged outpost kits can be consumed")
	}

	// get outpost template id
	otIDStr, isOutpost := item.ItemTypeMeta.GetString("outposttemplateid")

	if !isOutpost {
		return errors.New("unable to find linked outpost template id")
	}

	// parse template id
	otID, err := uuid.Parse(otIDStr)

	if err != nil {
		return err
	}

	// todo: check for obstacles

	// reduce quantity of item to 0 (always unpackaged)
	item.Quantity = 0
	item.CoreDirty = true

	// escalate to core for saving in db
	s.CurrentSystem.ChangedQuantityItems = append(s.CurrentSystem.ChangedQuantityItems, item)

	// generate new outpost ticket
	nt := NewOutpostTicket{
		OutpostTemplateID: otID,
		UserID:            s.UserID,
		PosX:              s.PosX,
		PosY:              s.PosY,
		Theta:             s.Theta,
	}

	c, ok := s.CurrentSystem.clients[s.UserID.String()]

	if ok {
		nt.Client = c
	}

	// escalate to core for spawning new outpost
	s.CurrentSystem.NewOutpostTickets = append(s.CurrentSystem.NewOutpostTickets, &nt)

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
				tgtReg := models.SharedTargetTypeRegistry

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
				} else if *m.TargetType == tgtReg.Wreck {
					// find wreck
					_, f := m.shipMountedOn.CurrentSystem.wrecks[m.TargetID.String()]

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
			} else if m.ItemTypeFamily == "salvager" {
				canActivate = m.activateAsSalvager()
			} else if m.ItemTypeFamily == "utility_veil" {
				canActivate = m.activateAsUtilityVeil()
			} else if m.ItemTypeFamily == "fuel_loader" {
				canActivate = m.activateAsFuelLoader()
			} else if m.ItemTypeFamily == "utility_add" {
				canActivate = m.activateAsAreaDenialDevice()
			} else if m.ItemTypeFamily == "ewar_cycle" {
				canActivate = m.activateAsCycleDisruptor()
			} else if m.ItemTypeFamily == "ewar_fcj" {
				canActivate = m.activateAsFireControlJammer()
			} else if m.ItemTypeFamily == "ewar_r_mask" {
				canActivate = m.activateAsRegenerationMask()
			} else if m.ItemTypeFamily == "ewar_d_mask" {
				canActivate = m.activateAsDissipationMask()
			} else if m.ItemTypeFamily == "burst_reactor" {
				canActivate = m.activateAsBurstReactor()
			} else if m.ItemTypeFamily == "xfer_heat" {
				canActivate = m.activateAsHeatXfer()
			} else if m.ItemTypeFamily == "xfer_energy" {
				canActivate = m.activateAsEnergyXfer()
			} else if m.ItemTypeFamily == "xfer_shield" {
				canActivate = m.activateAsShieldXfer()
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

// Helper function to recalculate all stat caches
func RecalcAllStatCaches(s *Ship) {
	// recalculate caches
	s.GetRealHeatSink(true)
	s.GetRealMaxHeat(true)
	s.GetRealAccel(true)
	s.GetRealTurn(true)
	s.GetRealSpaceDrag(true)
	s.GetRealMaxFuel(true)
	s.GetRealMaxEnergy(true)
	s.GetRealMaxShield(true)
	s.GetRealMaxArmor(true)
	s.GetRealMaxHull(true)
	s.GetRealEnergyRegen(true)
	s.GetRealShieldRegen(true)
	s.GetRealCargoBayVolume(true)

	// zero sums
	s.SumCloaking = 0
	s.SumVeiling = 0
}
