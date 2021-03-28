package universe

import (
	"helia/physics"

	"github.com/google/uuid"
)

// Structure representing an asteroid
type Asteroid struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	Name     string
	Texture  string
	Radius   float64
	Theta    float64
	PosX     float64
	PosY     float64
	Mass     float64
	// secret, do not expose to player in global update
	Yield float64
	// secret (ore details)
	ItemTypeID     uuid.UUID
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
}

// Returns a new physics dummy structure representing this station
func (a *Asteroid) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: a.PosX,
		PosY: a.PosY,
		Mass: a.Mass,
	}
}
