package universe

import (
	"github.com/google/uuid"
)

// Structure representing a gas and its yield for gas mining
type GasMiningYield struct {
	ItemTypeID uuid.UUID `json:"itemTypeId"`
	Yield      int       `json:"yield"`
	// in-memory only
	ItemTypeName   string
	ItemFamilyName string
	ItemFamilyID   string
	ItemTypeMeta   Meta
}

// Structure representing gases ("medium wisps") that can be mined around a celestial
type GasMiningMetadata struct {
	Yields map[string]GasMiningYield `json:"yields"`
}
