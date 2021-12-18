package shared

import (
	"math"

	"github.com/google/uuid"
)

// Structure representing a player's amount of experience with a given kind of game entity
type PlayerExperienceSheet struct {
	ShipExperience map[string]*ShipExperienceEntry
	// in-memory only
	Lock LabeledMutex
}

// Structure representing a player's experience flying ships of a ship template
type ShipExperienceEntry struct {
	SecondsOfExperience float64
	ShipTemplateID      uuid.UUID
	ShipTemplateName    string
}

// Returns the unrounded experience level represented by a ShipExperienceEntry
func (e *ShipExperienceEntry) GetExperience() float64 {
	return secondsToExperienceLevel(e.SecondsOfExperience)
}

func secondsToExperienceLevel(s float64) float64 {
	// convert seconds to minutes
	m := s / 60.0

	// calculate experience level
	return math.Log((math.Pow(m, 0.85)) + (m / 4) + (math.Pow(m, 0.25)) + 1)
}
