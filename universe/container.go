package universe

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

//Container Structure representing a container in the running game simulation.
type Container struct {
	ID      uuid.UUID
	Meta    Meta
	Created time.Time
	//in-memory only
	Lock  sync.Mutex
	Items []Item
}
