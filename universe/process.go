package universe

import (
	"fmt"
	"helia/shared"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Structure representing the status of a manufacturing process in a specific station
type StationProcess struct {
	ID            uuid.UUID
	StationID     uuid.UUID
	ProcessID     uuid.UUID
	Progress      int
	Installed     bool
	InternalState StationProcessInternalState
	Meta          Meta
	// in-memory only
	Lock            shared.LabeledMutex
	Process         Process
	StationName     string
	SolarSystemName string
	MSCounter       int64
}

// Updates a station manufacturing process for a tick
func (p *StationProcess) PeriodicUpdate(dT int64) {
	// obtain lock
	p.Lock.Lock("process.PeriodicUpdate")
	defer p.Lock.Unlock()

	// check process status
	if p.InternalState.IsRunning {
		if p.Progress >= p.Process.Time {
			p.InternalState.IsRunning = false

			// make sure there is enough room to deliver outputs
			for k := range p.InternalState.Outputs {
				o := p.InternalState.Outputs[k]
				s := p.Process.Outputs[k]
				m := s.GetIndustrialMetadata()

				if s.Quantity+o.Quantity > m.SiloSize {
					// no room - can't deliver
					return
				}
			}

			// deliver results
			for k := range p.InternalState.Outputs {
				o := p.InternalState.Outputs[k]
				s := p.Process.Outputs[k]

				// store updated factor
				o.Quantity += s.Quantity
				p.InternalState.Outputs[k] = o
			}

			// reset process
			p.Progress = 0
			p.MSCounter = 0
		} else {
			// advance clock
			p.MSCounter += dT

			// check for second tick
			if p.MSCounter >= 1000 {
				// add 1 second to clock
				p.Progress += 1

				// roll back ms counter
				p.MSCounter -= 1000
			}
		}
	} else {
		// check for all available inputs
		for k := range p.InternalState.Inputs {
			i := p.InternalState.Inputs[k]
			s := p.Process.Inputs[k]

			if i.Quantity-s.Quantity < 0 {
				// insufficient input resources - can't start
				return
			}
		}

		// collect input resources from silos
		for k := range p.InternalState.Inputs {
			i := p.InternalState.Inputs[k]
			s := p.Process.Inputs[k]

			// store updated factor
			i.Quantity -= s.Quantity
			p.InternalState.Inputs[k] = i
		}

		// start process
		p.InternalState.IsRunning = true
	}
}

// Structure representing the internal state of the ware silos involved in the process
type StationProcessInternalState struct {
	IsRunning bool
	Inputs    map[string]*StationProcessInternalStateFactor
	Outputs   map[string]*StationProcessInternalStateFactor
}

// Structure representing an input or output factor in a station process's internal state
type StationProcessInternalStateFactor struct {
	Quantity int
	Price    int
}

// Structure representing a manufacturing process
type Process struct {
	ID   uuid.UUID
	Name string
	Meta Meta
	Time int
	// in-memory only
	Inputs  map[string]ProcessInput
	Outputs map[string]ProcessOutput
}

// Structure representing an input resource in a manufacturing process
type ProcessInput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
	// in-memory only
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
}

// Structure representing an output product from a manufacturing process
type ProcessOutput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
	// in-memory only
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
}

// Structure representing the standardized metadata for item types on the industrial market
type IndustrialMetadata struct {
	MinPrice   int
	MaxPrice   int
	SiloSize   int
	ProcessID *uuid.UUID
}

// Fetches industrial market limits from item type metadata
func (p *ProcessInput) GetIndustrialMetadata() IndustrialMetadata {
	// make empty metadata
	d := IndustrialMetadata{}

	// attempt to fetch from metadata
	l, f := p.ItemTypeMeta.GetMap("industrialmarket")

	if f {
		// load from metadata
		minprice, _ := l.GetInt("minprice")
		maxprice, _ := l.GetInt("maxprice")
		silosize, _ := l.GetInt("silosize")
		processidStr, _ := l.GetString("process_id")

		d.MinPrice = minprice
		d.MaxPrice = maxprice
		d.SiloSize = silosize

		if len(processidStr) > 0 {
			x := uuid.MustParse(processidStr)
			d.ProcessID = &x
		}
	}

	return d
}

// Fetches industrial market limits from item type metadata
func (p *ProcessOutput) GetIndustrialMetadata() IndustrialMetadata {
	// make empty metadata
	d := IndustrialMetadata{}

	// attempt to fetch from metadata
	l, f := p.ItemTypeMeta.GetMap("industrialmarket")

	if f {
		// load from metadata
		minprice, _ := l.GetInt("minprice")
		maxprice, _ := l.GetInt("maxprice")
		silosize, _ := l.GetInt("silosize")

		d.MinPrice = minprice
		d.MaxPrice = maxprice
		d.SiloSize = silosize
	}

	return d
}

// Returns a copy of a station process
func (p *StationProcess) CopyStationProcess() *StationProcess {
	copy := StationProcess{
		ID:            p.ID,
		StationID:     p.StationID,
		ProcessID:     p.ProcessID,
		Progress:      p.Progress,
		Installed:     p.Installed,
		InternalState: p.InternalState,
		Meta:          p.Meta,
		Lock: shared.LabeledMutex{
			Structure: "StationProcess",
			UID:       fmt.Sprintf("%v :: %v :: %v", p.ID, time.Now(), rand.Float64()),
		},
		Process:         p.Process,
		MSCounter:       p.MSCounter,
		StationName:     p.StationName,
		SolarSystemName: p.SolarSystemName,
	}

	return &copy
}
