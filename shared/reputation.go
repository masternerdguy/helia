package shared

import (
	"github.com/google/uuid"
)

const MIN_STANDING = -10
const MAX_STANDING = 10
const MIN_DOCK_STANDING = -1.999

// Structure representing a relationship this faction has to another faction
type FactionReputationSheetEntry struct {
	SourceFactionID  uuid.UUID
	TargetFactionID  uuid.UUID
	StandingValue    float64
	AreOpenlyHostile bool
}

// Structure containing information about this faction's relationship with the world
type FactionReputationSheet struct {
	Entries        map[string]FactionReputationSheetEntry
	HostFactionIDs []uuid.UUID
	WorldPercent   float64
	// factions do not need a lock because their reputations are static and won't be updated
}

// Structure representing a relationship this player has to a faction
type PlayerReputationSheetFactionEntry struct {
	FactionID        uuid.UUID
	StandingValue    float64
	AreOpenlyHostile bool
}

// Structure containing information about this player's relationship with factions
type PlayerReputationSheet struct {
	FactionEntries map[string]*PlayerReputationSheetFactionEntry // map key is faction id string
	Lock           LabeledMutex
}

// Enforces standing bounds on reputation entries
func (s *PlayerReputationSheet) EnforceBounds(lock bool) {
	if lock {
		s.Lock.Lock("reputation.EnforceBounds")
		defer s.Lock.Unlock()
	}

	for k := range s.FactionEntries {
		u := s.FactionEntries[k]

		if u.StandingValue < MIN_STANDING {
			u.StandingValue = MIN_STANDING
		}

		if u.StandingValue > MAX_STANDING {
			u.StandingValue = MAX_STANDING
		}
	}
}
