package universe

import (
	"time"

	"github.com/google/uuid"
)

// Structure representing an initial start a player has chosen
type Start struct {
	ID             uuid.UUID
	Name           string
	ShipTemplateID uuid.UUID
	ShipFitting    StartFitting
	Created        time.Time
	Available      bool
	SystemID       uuid.UUID
	HomeStationID  uuid.UUID
}

// Structure representing the initial fitting of a starter ship of a given start
type StartFitting struct {
	ARack []StartFittedSlot
	BRack []StartFittedSlot
	CRack []StartFittedSlot
}

// Structure representing a slot within the initial fitting of a starter ship of a given start
type StartFittedSlot struct {
	ItemTypeID uuid.UUID
}
