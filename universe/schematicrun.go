package universe

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Structure representing a user-initiated manufacturing job
type SchematicRun struct {
	ID              uuid.UUID
	Created         time.Time
	ProcessID       uuid.UUID
	StatusID        string
	Progress        int
	SchematicItemID uuid.UUID
	UserID          uuid.UUID
	// in-memory only
	Initialized   bool
	Lock          sync.Mutex
	SchematicItem *Item
	Process       *Process
	Ship          *Ship
}
