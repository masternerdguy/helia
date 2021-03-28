package universe

import "github.com/google/uuid"

//StationProcess Structure representing the status of a manufaturing process in a specific station
type StationProcess struct {
	ID            uuid.UUID
	StationID     uuid.UUID
	ProcessID     uuid.UUID
	Progress      int
	Installed     bool
	InternalState Meta // todo: make a proper model for this
	Meta          Meta
	// in-memory only
	Process Process
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
}

// Structure representing an output product from a manufacturing process
type ProcessOutput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
}
