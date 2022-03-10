package universe

import (
	"helia/shared"

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
	ReputationSheet shared.FactionReputationSheet
	Ticker          string
	OwnerID         *uuid.UUID
	HomeStationID   *uuid.UUID
}

// Given a faction to compare against, returns the standing and whether they have declared open hostilities
func (s *Faction) CheckStandings(factionID uuid.UUID) (float64, bool) {
	// try to find faction relationship
	if val, ok := s.ReputationSheet.Entries[factionID.String()]; ok {
		return val.StandingValue, val.AreOpenlyHostile
	} else {
		return 0, false
	}
}
