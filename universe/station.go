package universe

import (
	"fmt"
	"helia/physics"
	"helia/shared"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Structure representing an NPC space station
type Station struct {
	ID          uuid.UUID
	SystemID    uuid.UUID
	StationName string
	PosX        float64
	PosY        float64
	Texture     string
	Radius      float64
	Mass        float64
	Theta       float64
	FactionID   uuid.UUID
	// in-memory only
	Lock                   shared.LabeledMutex
	CurrentSystem          *SolarSystem
	OpenSellOrders         map[string]*SellOrder
	Processes              map[string]*StationProcess
	Faction                Faction
	lastPeriodicUpdateTime time.Time
}

// Initializes internal aspects of Station
func (s *Station) Initialize() {
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

			/*
			 * In Helia, "stations" are always NPC operated and indestructible. This is so that players
			 * have a safe refuge to park their stuff and base out of. The equivalent, destructible,
			 * structures built by players will be called "Outposts". Since "stations" are NPC-owned,
			 * we want to pre-seed process inputs and outputs at installation to create the feeling of
			 * a used universe.
			 */

			// iterate over inputs
			for _, x := range process.Process.Inputs {
				// get industrial market metadata
				marketLimits := x.GetIndustrialMetadata()

				// randomize stack size based on market limit
				sf := StationProcessInternalStateFactor{
					Quantity: physics.RandInRange(0, marketLimits.SiloSize),
				}

				// store in state
				is.Inputs[x.ItemTypeID.String()] = &sf
			}

			for _, x := range process.Process.Outputs {
				// get industrial market metadata
				marketLimits := x.GetIndustrialMetadata()

				// randomize stack size based on market limit
				sf := StationProcessInternalStateFactor{
					Quantity: physics.RandInRange(0, marketLimits.SiloSize),
				}

				// store in state
				is.Outputs[x.ItemTypeID.String()] = &sf
			}

			// randomize job progress
			process.Progress = physics.RandInRange(0, process.Process.Time)

			if process.Progress > 0 {
				is.IsRunning = true
			} else {
				is.IsRunning = false
			}

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

// Processes the station for a tick
func (s *Station) PeriodicUpdate() {
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

// Returns a copy of the station
func (s *Station) CopyStation() Station {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	copiedProcesses := make(map[string]*StationProcess)

	for i, p := range s.Processes {
		copy := p.CopyStationProcess()
		copiedProcesses[i] = copy
	}

	return Station{
		ID:          s.ID,
		StationName: s.StationName,
		PosX:        s.PosX,
		PosY:        s.PosY,
		SystemID:    s.SystemID,
		Texture:     s.Texture,
		Theta:       s.Theta,
		Radius:      s.Radius,
		Mass:        s.Mass,
		FactionID:   s.FactionID,
		// in-memory only
		Lock: shared.LabeledMutex{
			Structure: "Station",
			UID:       fmt.Sprintf("%v :: %v :: %v", s.ID, time.Now(), rand.Float64()),
		},
		Processes: copiedProcesses,
	}
}

// Returns a new physics dummy structure representing this station
func (s *Station) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.Mass,
	}
}

// Stub to absorb damage inflicted on station
func (s *Station) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	// do nothing - NPC owned stations can't be destroyed
}

// Recalculates the prices of production factors based on a linear demand curve
func (s *Station) calculateIndustrialMarketPrices() {
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
