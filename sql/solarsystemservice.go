package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_systems table in the database
type SolarSystemService struct{}

// Returns a solar system service for interacting with solar systems in the database
func GetSolarSystemService() SolarSystemService {
	return SolarSystemService{}
}

// Structure representing a row in the universe_systems table
type SolarSystem struct {
	ID               uuid.UUID
	SystemName       string
	RegionID         uuid.UUID
	HoldingFactionID uuid.UUID
}

// Retrieves all solar systems in a region from the database
func (s SolarSystemService) GetSolarSystemsByRegion(regionID uuid.UUID) ([]SolarSystem, error) {
	systems := make([]SolarSystem, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load solar systems
	sql := "select id, systemname, regionid, holding_factionid from universe_systems where regionid = $1"

	rows, err := db.Query(sql, regionID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := SolarSystem{}

		// scan into system structure
		rows.Scan(&r.ID, &r.SystemName, &r.RegionID, &r.HoldingFactionID)

		// append to system slice
		systems = append(systems, r)
	}

	return systems, err
}
