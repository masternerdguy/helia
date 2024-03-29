package universe

import (
	"helia/shared"
	"sync"

	"github.com/google/uuid"
)

// Structure representing a faction that something can be a member of
type Faction struct {
	ID              uuid.UUID
	Name            string
	Description     string
	IsNPC           bool
	IsJoinable      bool
	CanHoldSov      bool
	Meta            Meta
	ReputationSheet shared.FactionReputationSheet
	Ticker          string
	OwnerID         *uuid.UUID
	HomeStationID   *uuid.UUID
	// in-memory only
	Lock         sync.Mutex
	Applications map[string]FactionApplicationTicket
}

// Structure representing a partially validated request to create a new faction and join the creator into it
type NewFactionTicket struct {
	Name          string
	Description   string
	Ticker        string
	Client        *shared.GameClient
	HomeStationID uuid.UUID
}

// Structure representing a validated request to leave a player faction and rejoin the starter faction
type LeaveFactionTicket struct {
	Client *shared.GameClient
}

// Structure representing a validated request to join a player faction
type FactionApplicationTicket struct {
	UserID         uuid.UUID
	CurrentFaction *Faction
	CharacterName  string
}

// Structure representing a partially validated request to join a player into a player faction
type JoinFactionTicket struct {
	UserID      uuid.UUID
	FactionID   uuid.UUID
	OwnerClient *shared.GameClient
}

// Structure representing a validated request to view the current member list
type ViewMembersTicket struct {
	FactionID   uuid.UUID
	OwnerClient *shared.GameClient
}

// Structure representing a partially validated request to kick a member from a player faction
type KickMemberTicket struct {
	UserID      uuid.UUID
	FactionID   uuid.UUID
	OwnerClient *shared.GameClient
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
