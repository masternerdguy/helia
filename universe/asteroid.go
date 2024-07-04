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
	Transient bool
	Lock      sync.Mutex
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

// Returns a copy of the asteroid
func (s *Asteroid) CopyAsteroid() Asteroid {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	return Asteroid{
		ID:       s.ID,
		SystemID: s.SystemID,
		Name:     s.Name,
		Texture:  s.Texture,
		Radius:   s.Radius,
		Theta:    s.Theta,
		PosX:     s.PosX,
		PosY:     s.PosY,
		Mass:     s.Mass,
		Meta:     s.Meta,
		// secret, do not expose to player in global update
		Yield: s.Yield,
		// secret (ore details)
		ItemTypeID:     s.ItemTypeID,
		ItemTypeName:   s.ItemTypeName,
		ItemFamilyID:   s.ItemFamilyID,
		ItemFamilyName: s.ItemFamilyName,
		ItemTypeMeta:   s.ItemTypeMeta,
		// in-memory only
		Lock: sync.Mutex{},
	}
}
