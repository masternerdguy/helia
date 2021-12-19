package shared

import (
	"helia/listener/models"
	"math"

	"github.com/google/uuid"
)

// Structure representing a player's amount of experience with a given kind of game entity
type PlayerExperienceSheet struct {
	ShipExperience   map[string]*ShipExperienceEntry
	ModuleExperience map[string]*ModuleExperienceEntry
	// in-memory only
	Lock LabeledMutex
}

// Structure representing a player's experience flying ships of a ship template
type ShipExperienceEntry struct {
	SecondsOfExperience float64
	ShipTemplateID      uuid.UUID
	ShipTemplateName    string
}

// Structure representing a player's experience using modules of a givem type
type ModuleExperienceEntry struct {
	SecondsOfExperience float64
	ItemTypeID          uuid.UUID
	ItemTypeName        string
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

	for _, e := range e.ModuleExperience {
		if e == nil {
			continue
		}

		u.ModuleEntries = append(u.ModuleEntries, models.ServerExperienceUpdateModuleEntryBody{
			ExperienceLevel: e.GetExperience(),
			ItemTypeID:      e.ItemTypeID,
			ItemTypeName:    e.ItemTypeName,
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
	return secondsToExperienceLevel(e.SecondsOfExperience, 0)
}

// Returns a module experience entry from the map or returns a blank one if not found
func (e *PlayerExperienceSheet) GetModuleExperienceEntry(itemTypeID uuid.UUID) ModuleExperienceEntry {
	// obtain lock
	e.Lock.Lock("playerexperiencesheet.GetModuleExperienceEntry")
	defer e.Lock.Unlock()

	// build empty entry
	x := ModuleExperienceEntry{
		ItemTypeID: itemTypeID,
	}

	// copy if found
	v, f := e.ModuleExperience[itemTypeID.String()]

	if f {
		x.SecondsOfExperience = v.SecondsOfExperience
		x.ItemTypeName = v.ItemTypeName
	}

	// return result
	return x
}

// Overwrites a module experience entry in the map
func (e *PlayerExperienceSheet) SetModuleExperienceEntry(value ModuleExperienceEntry) {
	// obtain lock
	e.Lock.Lock("playerexperiencesheet.SetModuleExperienceEntry")
	defer e.Lock.Unlock()

	// update map
	e.ModuleExperience[value.ItemTypeID.String()] = &value
}

// Returns the unrounded experience level represented by a ModuleExperienceEntry
func (e *ModuleExperienceEntry) GetExperience() float64 {
	return secondsToExperienceLevel(e.SecondsOfExperience, -0.05)
}

// Converts seconds to an experience level using a logarithmic function
func secondsToExperienceLevel(s float64, offset float64) float64 {
	// convert seconds to minutes
	m := s / 60.0

	// calculate experience level
	return math.Log((math.Pow(m, 0.85-offset)) + (m / (4 + offset)) + (math.Pow(m, (0.25 - offset))) + 1)
}
