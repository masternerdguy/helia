package universe

import (
	"helia/physics"
	"sync"

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
	Meta     Meta
	// secret, do not expose to player in global update
	Yield float64
	// secret (ore details)
	ItemTypeID     uuid.UUID
	ItemTypeName   string
	ItemFamilyID   string
	ItemFamilyName string
	ItemTypeMeta   Meta
	// in-memory only
	GasMiningMetadata GasMiningMetadata
	Transient         bool
	Lock              sync.Mutex
}

// Returns a new physics dummy structure representing this asteroid
func (a *Asteroid) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: a.PosX,
		PosY: a.PosY,
		Mass: a.Mass,
	}
}

// Returns a copy of the asteroid
func (a *Asteroid) CopyAsteroid() Asteroid {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	return Asteroid{
		ID:       a.ID,
		SystemID: a.SystemID,
		Name:     a.Name,
		Texture:  a.Texture,
		Radius:   a.Radius,
		Theta:    a.Theta,
		PosX:     a.PosX,
		PosY:     a.PosY,
		Mass:     a.Mass,
		Meta:     a.Meta,
		// secret, do not expose to player in global update
		Yield: a.Yield,
		// secret (ore details)
		ItemTypeID:     a.ItemTypeID,
		ItemTypeName:   a.ItemTypeName,
		ItemFamilyID:   a.ItemFamilyID,
		ItemFamilyName: a.ItemFamilyName,
		ItemTypeMeta:   a.ItemTypeMeta,
		// in-memory only
		GasMiningMetadata: a.GasMiningMetadata,
		Lock:              sync.Mutex{},
	}
}
