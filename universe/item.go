package universe

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Structure representing the most basic features of an item in the running game simulation
type Item struct {
	ID            uuid.UUID
	ItemTypeID    uuid.UUID
	Meta          Meta
	Created       time.Time
	CreatedBy     *uuid.UUID
	CreatedReason string
	ContainerID   uuid.UUID
	Quantity      int
	IsPackaged    bool
	// in-memory only
	Lock           sync.Mutex
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
	Process        *Process
	CoreDirty      bool
	SchematicInUse bool
}

// Structure representing a request to create a new item by item type name, used for devhax only
type NewItemFromNameTicketDevHax struct {
	ItemTypeName string
	Quantity     int
	Container    *Container
	ContainerID  uuid.UUID
	UserID       uuid.UUID
}

// Structure representing a request to create a new ship by item type name, used for devhax only
type NewShipFromNameTicketDevHax struct {
	ItemTypeName string
	Quantity     int
	StationID    uuid.UUID
	UserID       uuid.UUID
}

// List of module attributes that can be percentage mutated by mod kits
var MutableModuleAttributes = [...]string{
	"range",
	"shield_damage",
	"armor_damage",
	"hull_damage",
	"ore_mining_volume",
	"ice_mining_volume",
	"volume",
	"fault_tolerance",
	"flight_time",
	"shield_boost_amount",
	"activation_energy",
	"cooldown",
	"drag_multiplier",
	"energy_siphon_amount",
	"salvage_volume",
	"salvage_chance",
	"tracking",
	"armor_max_add",
	"energy_max_add",
	"energy_regen_max_add",
	"heat_sink_add",
	"fuel_max_add",
	"max_fuel_volume",
	"leakage",
	"heat_damage",
	"missile_destruction_chance",
	"signal_flux",
	"signal_gain",
	"guidance_drift",
	"tracking_drift",
	"mask_radius",
	"heat_cap_max_add",
}

// Helper function to determine whether or not a given item meta attribute can be mutated by a mod kit
func itemMetaIsMutable(key string) bool {
	for _, v := range MutableModuleAttributes {
		if v == key {
			return true
		}
	}

	return false
}

// Structure mirroring the SQL ItemType structure - should not be used outside of the universe caches and loader
type ItemTypeRaw struct {
	ID     uuid.UUID
	Family string
	Name   string
	Meta   Meta `json:"meta"`
}

// Structure mirroring the SQL ItemFamily structure - should not be used outside of the universe caches and loader
type ItemFamilyRaw struct {
	ID           string
	FriendlyName string
	Meta         Meta `json:"meta"`
}
