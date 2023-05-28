package universe

import (
	"helia/physics"
	"helia/shared"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Structure representing an player-owned space station
type Outpost struct {
	ID          uuid.UUID
	SystemID    uuid.UUID
	OutpostName string
	PosX        float64
	PosY        float64
	Theta       float64
	Shield      float64
	Armor       float64
	Hull        float64
	UserID      uuid.UUID
	Wallet      float64
	Created     time.Time
	Destroyed   bool
	DestroyedAt *time.Time
	// cache of base template
	TemplateData OutpostTemplate
	// cache from controlling user
	FactionID uuid.UUID
	// in-memory only
	Lock                   sync.Mutex
	CurrentSystem          *SolarSystem
	OpenSellOrders         map[string]*SellOrder
	Processes              map[string]*OutpostProcess
	CharacterName          string
	Faction                *Faction
	lastPeriodicUpdateTime time.Time
}

// Structure representing a newly deployed outpost, not yet materialized
type NewOutpostTicket struct {
	ID                uuid.UUID
	OutpostTemplateID uuid.UUID
	UserID            uuid.UUID
	PosX              float64
	PosY              float64
	Theta             float64
	Client            *shared.GameClient
}

// Initializes internal aspects of an outpost
func (s *Outpost) Initialize() {
	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// initialize maps
	s.OpenSellOrders = make(map[string]*SellOrder)

	// install processes if needed
	for i := range s.Processes {
		process := s.Processes[i]

		if !process.Installed {
			// set up process for first time
			is := StationProcessInternalState{}
			is.Inputs = make(map[string]*StationProcessInternalStateFactor)
			is.Outputs = make(map[string]*StationProcessInternalStateFactor)

			// store state
			process.InternalState = is

			// mark as installed
			process.Installed = true
		}
	}

	// do initial price calculation
	s.calculateIndustrialMarketPrices()

	// store time
	s.lastPeriodicUpdateTime = time.Now()
}

// Processes the outpost for a tick
func (s *Outpost) PeriodicUpdate() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// calculate delta and store time
	dT := time.Since(s.lastPeriodicUpdateTime).Milliseconds()
	s.lastPeriodicUpdateTime = time.Now()

	// update processes
	for _, p := range s.Processes {
		p.PeriodicUpdate(dT)
	}

	// recalculate industrial market prices
	s.calculateIndustrialMarketPrices()
}

// Returns a copy of the outpost
func (s *Outpost) CopyOutpost() Outpost {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	copiedProcesses := make(map[string]*OutpostProcess)

	for i, p := range s.Processes {
		copy := p.CopyOutpostProcess()
		copiedProcesses[i] = copy
	}

	return Outpost{
		ID:           s.ID,
		OutpostName:  s.OutpostName,
		PosX:         s.PosX,
		PosY:         s.PosY,
		SystemID:     s.SystemID,
		Theta:        s.Theta,
		Shield:       s.Shield,
		Armor:        s.Armor,
		Hull:         s.Hull,
		TemplateData: s.TemplateData,
		UserID:       s.UserID,
		FactionID:    s.FactionID,
		// in-memory only
		Lock:      sync.Mutex{},
		Processes: copiedProcesses,
	}
}

// Returns a new physics dummy structure representing this outpost
func (s *Outpost) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.TemplateData.BaseMass,
	}
}

// Stub to absorb damage inflicted on outpost
func (s *Outpost) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	// todo: not yet implemented
}

// Returns the real max shield of the ship after modifiers
func (s *Outpost) GetRealMaxShield() float64 {
	// no modifiers yet
	return s.TemplateData.BaseShield
}

// Returns the real max armor of the ship after modifiers
func (s *Outpost) GetRealMaxArmor() float64 {
	// no modifiers yet
	return s.TemplateData.BaseArmor
}

// Returns the real max hull of the ship after modifiers
func (s *Outpost) GetRealMaxHull() float64 {
	// no modifiers yet
	return s.TemplateData.BaseHull
}

// Recalculates the prices of production factors based on a linear demand curve
func (s *Outpost) calculateIndustrialMarketPrices() {
	for i := range s.Processes {
		// get process
		process := s.Processes[i]

		// skip if not set up
		if !process.Installed {
			continue
		}

		// iterate over inputs and outputs
		for x := range process.Process.Inputs {
			o := process.Process.Inputs[x]

			// get industrial data
			marketLimits := o.GetIndustrialMetadata()

			if marketLimits.SiloSize <= 0 {
				// skip due to invalid silo size
				continue
			}

			// get internal state
			is := process.InternalState.Inputs[o.ItemTypeID.String()]

			// get percentage of silo filled
			filled := float32(is.Quantity) / float32(marketLimits.SiloSize)

			// calculate current price
			spread := marketLimits.MaxPrice - marketLimits.MinPrice
			price := int(float32(spread)*(1.0-filled)) + marketLimits.MinPrice

			// store price
			is.Price = price

			// update internal state
			process.InternalState.Inputs[o.ItemTypeID.String()] = is
		}

		for x := range process.Process.Outputs {
			o := process.Process.Outputs[x]

			// get industrial data
			marketLimits := o.GetIndustrialMetadata()

			if marketLimits.SiloSize <= 0 {
				// skip due to invalid silo size
				continue
			}

			// get internal state
			is := process.InternalState.Outputs[o.ItemTypeID.String()]

			// get percentage of silo filled
			filled := float32(is.Quantity) / float32(marketLimits.SiloSize)

			// calculate current price
			spread := marketLimits.MaxPrice - marketLimits.MinPrice
			price := int(float32(spread)*(1.0-filled)) + marketLimits.MinPrice

			// store price
			is.Price = price

			// update internal state
			process.InternalState.Outputs[o.ItemTypeID.String()] = is
		}
	}
}
