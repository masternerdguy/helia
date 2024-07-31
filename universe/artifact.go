package universe

import (
	"helia/physics"

	"github.com/google/uuid"
)

// Structure representing a artifact
type Artifact struct {
	ID           uuid.UUID
	SystemID     uuid.UUID
	ArtifactName string
	PosX         float64
	PosY         float64
	Texture      string
	Radius       float64
	Mass         float64
	Theta        float64
	Meta         Meta
}

// Returns a copy of the artifact
func (p *Artifact) CopyArtifact() Artifact {
	return Artifact{
		ID:       p.ID,
		PosX:     p.PosX,
		PosY:     p.PosY,
		SystemID: p.SystemID,
		Texture:  p.Texture,
		Theta:    p.Theta,
		Radius:   p.Radius,
		Mass:     p.Mass,
		Meta:     p.Meta,
	}
}

// Returns a new physics dummy structure representing this artifact
func (p *Artifact) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: p.PosX,
		PosY: p.PosY,
		Mass: p.Mass,
	}
}
