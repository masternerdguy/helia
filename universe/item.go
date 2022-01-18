package universe

import (
	"helia/shared"
	"time"

	"github.com/google/uuid"
)

// Structure representing the most basic features of an item in the running game simulation
type Item struct {
	ID            uuid.UUID
	ItemTypeID    uuid.UUID
	Meta          Meta
	Created       time.Time
	CreatedBy     *uuid.UUID
	CreatedReason string
	ContainerID   uuid.UUID
	Quantity      int
	IsPackaged    bool
	// in-memory only
	Lock           shared.LabeledMutex
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
	Process        *Process
	CoreDirty      bool
}
