package shared

import (
	"helia/listener/models"
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

func (e *PlayerExperienceSheet) CopyAsUpdate() models.ServerExperienceUpdateBody {
	e.Lock.Lock("playerexperiencesheet.CopyAsUpdate")
	defer e.Lock.Unlock()

	u := models.ServerExperienceUpdateBody{}

	for _, e := range e.ShipExperience {
		if e == nil {
			continue
		}

		u.ShipEntries = append(u.ShipEntries, models.ServerExperienceUpdateShipEntryBody{
			ExperienceLevel:  e.GetExperience(),
			ShipTemplateID:   e.ShipTemplateID,
			ShipTemplateName: e.ShipTemplateName,
		})
	}

	return u
}

// Returns a ship experience entry from the map or returns a blank one if not found
func (e *PlayerExperienceSheet) GetShipExperienceEntry(shipTemplateID uuid.UUID) ShipExperienceEntry {
	// obtain lock
	e.Lock.Lock("playerexperiencesheet.GetShipExperienceEntry")
	defer e.Lock.Unlock()

	// build empty entry
	x := ShipExperienceEntry{
		ShipTemplateID: shipTemplateID,
	}

	// copy if found
	v, f := e.ShipExperience[shipTemplateID.String()]

	if f {
		x.SecondsOfExperience = v.SecondsOfExperience
		x.ShipTemplateName = v.ShipTemplateName
	}

	// return result
	return x
}

// Overwrites a ship experience entry in the map
func (e *PlayerExperienceSheet) SetShipExperienceEntry(value ShipExperienceEntry) {
	// obtain lock
	e.Lock.Lock("playerexperiencesheet.SetShipExperienceEntry")
	defer e.Lock.Unlock()

	// update map
	e.ShipExperience[value.ShipTemplateID.String()] = &value
}

// Returns the unrounded experience level represented by a ShipExperienceEntry
func (e *ShipExperienceEntry) GetExperience() float64 {
	return secondsToExperienceLevel(e.SecondsOfExperience)
}

// Converts seconds to an experience level using a logarithmic function
func secondsToExperienceLevel(s float64) float64 {
	// convert seconds to minutes
	m := s / 60.0

	// calculate experience level
	return math.Log((math.Pow(m, 0.85)) + (m / 4) + (math.Pow(m, 0.25)) + 1)
}
