package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_regions table
type RegionService struct{}

// Gets a region service for interacting with regions in the database
func GetRegionService() *RegionService {
	return &RegionService{}
}

// Structure representing a row in the universe_regions table
type Region struct {
	ID         uuid.UUID
	RegionName string
	PosX       float64
	PosY       float64
}

// Retrieves all regions from the database
func (s RegionService) GetAllRegions() ([]Region, error) {
	regions := make([]Region, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load regions
	sql := "select id, regionname, pos_x, pos_y from universe_regions"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Region{}

		// scan into region structure
		rows.Scan(&r.ID, &r.RegionName, &r.PosX, &r.PosY)

		// append to region slice
		regions = append(regions, r)
	}

	return regions, err
}
