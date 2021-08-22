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
	PosX             float64
	PosY             float64
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
	sql := "select id, systemname, regionid, holding_factionid, pos_x, pos_y from universe_systems where regionid = $1"

	rows, err := db.Query(sql, regionID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := SolarSystem{}

		// scan into system structure
		rows.Scan(&r.ID, &r.SystemName, &r.RegionID, &r.HoldingFactionID, &r.PosX, &r.PosY)

		// append to system slice
		systems = append(systems, r)
	}

	return systems, err
}

// Creates a new solar system in the database (for worldmaker)
func (s SolarSystemService) NewSolarSystemWorldMaker(r *SolarSystem) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert system
	sql := `
			INSERT INTO public.universe_systems(
				id, systemname, regionid, holding_factionid, pos_x, pos_y)
				VALUES ($1, $2, $3, $4, $5, $6);
			`

	q, err := db.Query(sql, r.ID, r.SystemName, r.RegionID, r.HoldingFactionID, r.PosX, r.PosY)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
