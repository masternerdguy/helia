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
	sql := "select id, regionname from universe_regions"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Region{}

		// scan into region structure
		rows.Scan(&r.ID, &r.RegionName)

		// append to region slice
		regions = append(regions, r)
	}

	return regions, err
}
