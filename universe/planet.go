package universe

import (
	"helia/physics"

	"github.com/google/uuid"
)

// Structure representing a planet
type Planet struct {
	ID         uuid.UUID
	SystemID   uuid.UUID
	PlanetName string
	PosX       float64
	PosY       float64
	Texture    string
	Radius     float64
	Mass       float64
	Theta      float64
	Meta       Meta
	// in-memory only
	GasMiningMetadata GasMiningMetadata
}

// Returns a copy of the planet
func (p *Planet) CopyPlanet() Planet {
	return Planet{
		ID:       p.ID,
		PosX:     p.PosX,
		PosY:     p.PosY,
		SystemID: p.SystemID,
		Texture:  p.Texture,
		Theta:    p.Theta,
		Radius:   p.Radius,
		Mass:     p.Mass,
		Meta:     p.Meta,
		// in-memory only
		GasMiningMetadata: p.GasMiningMetadata,
	}
}

// Returns a new physics dummy structure representing this planet
func (p *Planet) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: p.PosX,
		PosY: p.PosY,
		Mass: p.Mass,
	}
}
