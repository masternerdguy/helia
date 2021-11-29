package universe

import (
	"helia/shared"
	"time"

	"github.com/google/uuid"
)

// Structure representing a container in the running game simulation.
type Container struct {
	ID      uuid.UUID
	Meta    Meta
	Created time.Time
	// in-memory only
	Lock  shared.LabeledMutex
	Items []*Item
}
