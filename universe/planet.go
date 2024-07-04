package universe

import "github.com/google/uuid"

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
}

// Returns a copy of the planet
func (s *Planet) CopyPlanet() Planet {
	return Planet{
		ID:       s.ID,
		PosX:     s.PosX,
		PosY:     s.PosY,
		SystemID: s.SystemID,
		Texture:  s.Texture,
		Theta:    s.Theta,
		Radius:   s.Radius,
		Mass:     s.Mass,
		Meta:     s.Meta,
	}
}
