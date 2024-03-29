package shared

import (
	"math"
	"sync"
	"time"

	"github.com/google/uuid"
)

const MIN_STANDING = -10
const MAX_STANDING = 10
const MIN_DOCK_STANDING = -1.999
const BECOME_OPENLY_HOSTILE = -6
const CLEAR_OPENLY_HOSTILE = 6
const INDIRECT_ADJUSTMENT_MODIFIER = 0.37
const STANDING_GAIN_MODIFIER = 0.15
const POLE_ZERO_BOUND = 0.00001

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
	// in-memory only
	TemporarilyOpenlyHostileUntil *time.Time
}

// Structure containing information about this player's relationship with factions
type PlayerReputationSheet struct {
	FactionEntries map[string]*PlayerReputationSheetFactionEntry // map key is faction id string
	// in-memory only
	Lock          sync.Mutex
	UserID        uuid.UUID
	CharacterName string
	FactionID     uuid.UUID
}

// Adjusts standing relative to an NPC faction
func (s *PlayerReputationSheet) AdjustStandingNPC(factionID uuid.UUID, factionRS FactionReputationSheet, amount float64, lock bool) {
	if lock {
		// obtain lock on this reputation sheet
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// reduce standings to faction being attacked
	s.applyStandingChange(factionID, amount)

	// iterate over faction relationships
	for _, v := range factionRS.Entries {
		if v.TargetFactionID == factionID {
			continue
		}

		if v.SourceFactionID == v.TargetFactionID {
			continue
		}

		// get indirect adjustment amount
		rv := (amount * (v.StandingValue / MAX_STANDING)) * INDIRECT_ADJUSTMENT_MODIFIER

		// apply indirect adjustment
		s.applyStandingChange(v.TargetFactionID, rv)
	}
}

// Adjusts standing relative to a player created faction
func (s *PlayerReputationSheet) AdjustStandingPlayer(playerRS *PlayerReputationSheet, amount float64, lock bool) {
	// null check
	if playerRS == nil {
		return
	}

	if lock {
		// obtain lock on this reputation sheet
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// obtain lock on attacker sheet
	playerRS.Lock.Lock()
	defer playerRS.Lock.Unlock()

	// verify they aren't members of the same faction
	if playerRS.FactionID == s.FactionID {
		// no penalty for infighting
		return
	}

	// iterate over attacked player's faction relationships
	for _, v := range playerRS.FactionEntries {
		// get indirect adjustment amount
		rv := (amount * (v.StandingValue / MAX_STANDING)) * INDIRECT_ADJUSTMENT_MODIFIER

		// apply indirect adjustment
		s.applyStandingChange(v.FactionID, rv)
	}
}

func (s *PlayerReputationSheet) applyStandingChange(factionID uuid.UUID, amount float64) {
	// get player's reputation sheet entry for this faction
	f, ok := s.FactionEntries[factionID.String()]

	if !ok {
		// does not exist - create a neutral one
		ne := PlayerReputationSheetFactionEntry{
			FactionID:        factionID,
			StandingValue:    0,
			AreOpenlyHostile: false,
		}

		s.FactionEntries[factionID.String()] = &ne
		f = s.FactionEntries[factionID.String()]
	}

	// get absolute standing value
	aV := math.Abs(f.StandingValue)

	// get absolute difference from pole
	aP := (10.0 - aV)

	if aP > 10 {
		aP = 10
	}

	if aP < 0 {
		aP = 0
	}

	// divide absolute pole value by 10
	fP := (aP / 10.0)

	if amount > 0 {
		// penalize standings gains more harshly
		fP *= STANDING_GAIN_MODIFIER
	}

	// bound pole near 0
	if fP < POLE_ZERO_BOUND {
		fP = POLE_ZERO_BOUND
	}

	// adjust standing
	f.StandingValue += (amount * fP)

	// check new amount
	if f.StandingValue >= CLEAR_OPENLY_HOSTILE {
		// unset openly hostile flag
		f.AreOpenlyHostile = false
	} else if f.StandingValue <= BECOME_OPENLY_HOSTILE {
		// set openly hostile flag
		f.AreOpenlyHostile = true
	}
}

// Enforces standing bounds on reputation entries
func (s *PlayerReputationSheet) EnforceBounds(lock bool) {
	if lock {
		s.Lock.Lock()
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
