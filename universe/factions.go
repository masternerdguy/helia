package universe

import (
	"github.com/google/uuid"
)

// Structure representing a faction that something can be a member of
type Faction struct {
	ID              uuid.UUID
	Name            string
	Description     string
	IsNPC           bool
	IsJoinable      bool
	IsClosed        bool
	CanHoldSov      bool
	Meta            Meta
	ReputationSheet FactionReputationSheet
	Ticker          string
}

type ReputationSheetEntry struct {
	SourceFactionID  uuid.UUID
	TargetFactionID  uuid.UUID
	StandingValue    float64
	AreOpenlyHostile bool
}

type FactionReputationSheet struct {
	Entries        map[string]ReputationSheetEntry
	HostFactionIDs []uuid.UUID
	WorldPercent   float64
}
