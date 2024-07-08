package universe

import "github.com/google/uuid"

// Structure representing a star
type Star struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	PosX     float64
	PosY     float64
	Texture  string
	Radius   float64
	Mass     float64
	Theta    float64
	Meta     Meta
	// in-memory only
	GasMiningMetadata GasMiningMetadata
}

// Returns a copy of the star
func (s *Star) CopyStar() Star {
	return Star{
		ID:       s.ID,
		PosX:     s.PosX,
		PosY:     s.PosY,
		SystemID: s.SystemID,
		Texture:  s.Texture,
		Theta:    s.Theta,
		Radius:   s.Radius,
		Mass:     s.Mass,
		Meta:     s.Meta,
		// in-memory only
		GasMiningMetadata: s.GasMiningMetadata,
	}
}
