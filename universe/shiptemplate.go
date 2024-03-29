package universe

import (
	"time"

	"github.com/google/uuid"
)

// Structure containing pre-modifier stats for a ship class
type ShipTemplate struct {
	ID                 uuid.UUID
	Created            time.Time
	ShipTemplateName   string
	Texture            string
	WreckTexture       string
	ExplosionTexture   string
	Radius             float64
	BaseAccel          float64
	BaseMass           float64
	BaseTurn           float64
	BaseShield         float64
	BaseShieldRegen    float64
	BaseArmor          float64
	BaseHull           float64
	BaseFuel           float64
	BaseHeatCap        float64
	BaseHeatSink       float64
	BaseEnergy         float64
	BaseEnergyRegen    float64
	ShipTypeID         uuid.UUID
	SlotLayout         SlotLayout
	BaseCargoBayVolume float64
	ItemTypeID         uuid.UUID
	CanUndock          bool
}

// Structure representing the slot layout available for fitting modules to this ship
type SlotLayout struct {
	ASlots []Slot
	BSlots []Slot
	CSlots []Slot
}

// Structure representing a slot in a SlotLayout
type Slot struct {
	Family          string
	Volume          int
	TexturePosition [2]float64
}
