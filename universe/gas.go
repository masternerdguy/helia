package universe

import (
	"log"

	"github.com/google/uuid"
)

// Structure representing a gas and its yield for gas mining
type GasMiningYield struct {
	ItemTypeId uuid.UUID `json:"itemTypeId"`
	Yield      int       `json:"yield"`
	// in-memory only
}

// Structure representing gases ("medium wisps") that can be mined around a celestial
type GasMiningMetadata struct {
	Yields map[string]GasMiningYield `json:"yields"`
}

// Fetches gas mining metadata from an asteroid
func (a *Asteroid) GetGasMiningMetadata() GasMiningMetadata {
	// make empty metadata
	d := GasMiningMetadata{}

	// attempt to fetch from metadata
	l, f := a.Meta.GetMap("gasmining")

	if f {
		// get yields from metadata
		ys, f := l.GetMap("yields")

		if f {
			// debug out
			log.Printf("ys: %v", ys)
		}
	}

	return d
}
