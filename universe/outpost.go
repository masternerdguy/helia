package universe

import (
	"helia/physics"
	"helia/shared"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Minimum deployment distance between an outpost and the nearest station
const OUTPOST_STATION_DEPLOY_MARGIN = 175000

// Minimum deployment distance between an outpost and the nearest outpost
const OUTPOST_OUTPOST_DEPLOY_MARGIN = 325000

// Minimum deployment distance between an outpost and the nearest asteroid
const OUTPOST_ASTEROID_DEPLOY_MARGIN = 75000

// Minimum deployment distance between an outpost and the nearest jumphole
const OUTPOST_JUMPHOLE_DEPLOY_MARGIN = 125000

// Structure representing an player-owned space station
type Outpost struct {
	ID          uuid.UUID
	SystemID    uuid.UUID
	OutpostName string
	PosX        float64
	PosY        float64
	Theta       float64
	Shield      float64
	Armor       float64
	Hull        float64
	UserID      uuid.UUID
	Wallet      float64
	Created     time.Time
	Destroyed   bool
	DestroyedAt *time.Time
	// cache of base template
	TemplateData OutpostTemplate
	// cache from controlling user
	FactionID uuid.UUID
	// in-memory only
	Lock                   sync.Mutex
	CurrentSystem          *SolarSystem
	SystemName             string
	CharacterName          string
	Faction                *Faction
	lastPeriodicUpdateTime time.Time
	dt                     int64
}

// Structure representing a newly deployed outpost, not yet materialized
type NewOutpostTicket struct {
	ID                uuid.UUID
	OutpostTemplateID uuid.UUID
	UserID            uuid.UUID
	PosX              float64
	PosY              float64
	Theta             float64
	Client            *shared.GameClient
}

// Structure representing a renamed outpost, not yet materialized
type OutpostRename struct {
	OutpostID uuid.UUID
	Name      string
}

// Processes the outpost for a tick
func (s *Outpost) PeriodicUpdate() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// calculate delta and store time
	s.dt = time.Since(s.lastPeriodicUpdateTime).Milliseconds()
	s.lastPeriodicUpdateTime = time.Now()
}

// Returns a copy of the outpost
func (s *Outpost) CopyOutpost() Outpost {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	return Outpost{
		ID:           s.ID,
		OutpostName:  s.OutpostName,
		PosX:         s.PosX,
		PosY:         s.PosY,
		SystemID:     s.SystemID,
		Theta:        s.Theta,
		Shield:       s.Shield,
		Armor:        s.Armor,
		Hull:         s.Hull,
		TemplateData: s.TemplateData,
		UserID:       s.UserID,
		FactionID:    s.FactionID,
		// in-memory only
		Lock: sync.Mutex{},
	}
}

// Returns a new physics dummy structure representing this outpost
func (s *Outpost) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.TemplateData.BaseMass,
	}
}

// Stub to absorb damage inflicted on outpost
func (s *Outpost) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	// todo: not yet implemented
}

// Returns the real max shield of the ship after modifiers
func (s *Outpost) GetRealMaxShield() float64 {
	// no modifiers yet
	return s.TemplateData.BaseShield
}

// Returns the real max armor of the ship after modifiers
func (s *Outpost) GetRealMaxArmor() float64 {
	// no modifiers yet
	return s.TemplateData.BaseArmor
}

// Returns the real max hull of the ship after modifiers
func (s *Outpost) GetRealMaxHull() float64 {
	// no modifiers yet
	return s.TemplateData.BaseHull
}
