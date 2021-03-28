package universe

import (
	"log"

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
	Process   Process
	MSCounter int64
}

func (p *StationProcess) PeriodicUpdate(dT int64) {
	if p.InternalState.IsRunning {
		if p.Progress >= p.Process.Time {
			// reset process
			p.Progress = 0
			p.MSCounter = 0
			p.InternalState.IsRunning = false

			// deliver result
			log.Println("process done!") // todo
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
	}
}

// Structure representing the internal state of the ware silos involved in the process
type StationProcessInternalState struct {
	IsRunning bool
	Inputs    map[string]StationProcessInternalStateFactor
	Outputs   map[string]StationProcessInternalStateFactor
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
	Inputs  []ProcessInput
	Outputs []ProcessOutput
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
	MinPrice int
	MaxPrice int
	SiloSize int
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

		d.MinPrice = minprice
		d.MaxPrice = maxprice
		d.SiloSize = silosize
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