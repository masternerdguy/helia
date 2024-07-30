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
func (o *Outpost) PeriodicUpdate() {
	o.Lock.Lock()
	defer o.Lock.Unlock()

	// calculate delta and store time
	o.dt = time.Since(o.lastPeriodicUpdateTime).Milliseconds()
	o.lastPeriodicUpdateTime = time.Now()
}

// Returns a copy of the outpost
func (o *Outpost) CopyOutpost() *Outpost {
	o.Lock.Lock()
	defer o.Lock.Unlock()

	op := Outpost{
		ID:          o.ID,
		SystemID:    o.SystemID,
		OutpostName: o.OutpostName,
		PosX:        o.PosX,
		PosY:        o.PosY,
		Theta:       o.Theta,
		Shield:      o.Shield,
		Armor:       o.Armor,
		Hull:        o.Hull,
		UserID:      o.UserID,
		Wallet:      o.Wallet,
		Created:     o.Created,
		Destroyed:   o.Destroyed,
		// copy base template
		TemplateData: OutpostTemplate{
			ID:                  o.TemplateData.ID,
			Created:             o.TemplateData.Created,
			OutpostTemplateName: o.TemplateData.OutpostTemplateName,
			Texture:             o.TemplateData.Texture,
			WreckTexture:        o.TemplateData.WreckTexture,
			ExplosionTexture:    o.TemplateData.ExplosionTexture,
			Radius:              o.TemplateData.Radius,
			BaseMass:            o.TemplateData.BaseMass,
			BaseShield:          o.TemplateData.BaseShield,
			BaseShieldRegen:     o.TemplateData.BaseShieldRegen,
			BaseArmor:           o.TemplateData.BaseArmor,
			BaseHull:            o.TemplateData.BaseHull,
			ItemTypeID:          o.TemplateData.ItemTypeID,
		},
		// cache from controlling user
		FactionID: o.FactionID,
		// in-memory only
		Lock:          sync.Mutex{},
		CurrentSystem: o.CurrentSystem,
		SystemName:    o.SystemName,
		CharacterName: o.CharacterName,
		Faction:       o.Faction,
	}

	// handle possibility of destruction
	if o.DestroyedAt != nil {
		op.DestroyedAt = o.DestroyedAt
	}

	// return result
	return &op
}

// Returns a new physics dummy structure representing this outpost
func (o *Outpost) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: o.PosX,
		PosY: o.PosY,
		Mass: o.TemplateData.BaseMass,
	}
}

// Stub to absorb damage inflicted on outpost
func (o *Outpost) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	// todo: not yet implemented
}

// Returns the real max shield of the ship after modifiers
func (o *Outpost) GetRealMaxShield() float64 {
	// no modifiers yet
	return o.TemplateData.BaseShield
}

// Returns the real max armor of the ship after modifiers
func (o *Outpost) GetRealMaxArmor() float64 {
	// no modifiers yet
	return o.TemplateData.BaseArmor
}

// Returns the real max hull of the ship after modifiers
func (o *Outpost) GetRealMaxHull() float64 {
	// no modifiers yet
	return o.TemplateData.BaseHull
}
