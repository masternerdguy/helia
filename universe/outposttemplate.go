package universe

import (
	"time"

	"github.com/google/uuid"
)

// Structure containing pre-modifier stats for a outpost class
type OutpostTemplate struct {
	ID                  uuid.UUID
	Created             time.Time
	OutpostTemplateName string
	Texture             string
	WreckTexture        string
	ExplosionTexture    string
	Radius              float64
	BaseMass            float64
	BaseShield          float64
	BaseShieldRegen     float64
	BaseArmor           float64
	BaseHull            float64
	ItemTypeID          uuid.UUID
}
