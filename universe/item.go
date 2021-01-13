package universe

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

//Item Structure representing the most basic features of an item in the running game simulation
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
	//in-memory only
	Lock           sync.Mutex
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
	CoreDirty      bool
}
