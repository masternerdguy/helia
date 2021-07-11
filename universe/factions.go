package universe

import (
	"github.com/google/uuid"
)

// Structure representing a faction that something can be a member of
type Faction struct {
	ID          uuid.UUID
	Name        string
	Description string
	IsNPC       bool
	IsJoinable  bool
	IsClosed    bool
	CanHoldSov  bool
	Meta        Meta
	Ticker      string
}
