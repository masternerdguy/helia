package shared

import (
	"time"

	"github.com/google/uuid"
)

// Structure representing any combat between two ships
type AggressionLog struct {
	// aggressor info
	UserID        uuid.UUID
	FactionID     uuid.UUID
	CharacterName string
	FactionName   string
	IsNPC         bool
	LastAggressed time.Time
	// aggressor ship info
	ShipID           uuid.UUID
	ShipName         string
	ShipTemplateID   uuid.UUID
	ShipTemplateName string
	// location info
	LastSolarSystemID   uuid.UUID
	LastSolarSystemName string
	LastRegionID        uuid.UUID
	LastRegionName      string
	LastPosX            float64
	LastPosY            float64
	// weapons used against victim
	WeaponUse map[string]*AggressionLogWeaponUse
}

// Structure representing a weapon used in combat between two ships
type AggressionLogWeaponUse struct {
	ItemID          uuid.UUID
	ItemTypeID      uuid.UUID
	ItemFamilyID    string
	ItemFamilyName  string
	ItemTypeName    string
	LastUsed        time.Time
	DamageInflicted float64
}

// Structure representing a validated request to view an action report summary page
type ViewActionReportPageTicket struct {
	Client *GameClient
	Page   int
	Take   int
}

// Structure representing a validated request to view a full action report
type ViewActionReportDetailTicket struct {
	Client *GameClient
	KillID uuid.UUID
}
