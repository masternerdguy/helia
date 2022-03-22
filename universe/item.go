package universe

import (
	"helia/shared"
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
	Lock           shared.LabeledMutex
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
	Process        *Process
	CoreDirty      bool
	SchematicInUse bool
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
