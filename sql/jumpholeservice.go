package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_jumpholes table
type JumpholeService struct{}

// Gets a jumphole service for interacting with jumpholes in the database
func GetJumpholeService() *JumpholeService {
	return &JumpholeService{}
}

// Structure representing a row in the universe_jumpholes table
type Jumphole struct {
	ID           uuid.UUID
	SystemID     uuid.UUID
	OutSystemID  uuid.UUID
	JumpholeName string
	PosX         float64
	PosY         float64
	Texture      string
	Radius       float64
	Mass         float64
	Theta        float64
}

// Retrieves all jumpholes from the database
func (s JumpholeService) GetAllJumpholes() ([]Jumphole, error) {
	jumpholes := make([]Jumphole, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load jumpholes
	sql := `
				SELECT id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_jumpholes;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Jumphole{}

		// scan into jumphole structure
		rows.Scan(&r.ID, &r.SystemID, &r.OutSystemID, &r.JumpholeName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		// append to jumphole slice
		jumpholes = append(jumpholes, r)
	}

	return jumpholes, err
}

// Retrieves all jumpholes in a given solar system from the database
func (s JumpholeService) GetJumpholesBySolarSystem(systemID uuid.UUID) ([]Jumphole, error) {
	jumpholes := make([]Jumphole, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load jumpholes
	sql := `
				SELECT id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_jumpholes
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Jumphole{}

		// scan into jumphole structure
		rows.Scan(&r.ID, &r.SystemID, &r.OutSystemID, &r.JumpholeName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		// append to jumphole slice
		jumpholes = append(jumpholes, r)
	}

	return jumpholes, err
}
