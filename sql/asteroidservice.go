package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_asteroids table
type AsteroidService struct{}

// Gets a asteroid service for interacting with asteroids in the database
func GetAsteroidService() *AsteroidService {
	return &AsteroidService{}
}

// Structure representing a row in the universe_asteroids table
type Asteroid struct {
	ID         uuid.UUID
	SystemID   uuid.UUID
	ItemTypeID uuid.UUID
	Name       string
	Texture    string
	Radius     float64
	Theta      float64
	PosX       float64
	PosY       float64
	Yield      float64
	Mass       float64
}

// Retrieves all asteroids in a given solar system from the database
func (s AsteroidService) GetAsteroidsBySolarSystem(systemID uuid.UUID) ([]Asteroid, error) {
	asteroids := make([]Asteroid, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load asteroids
	sql := `
				SELECT id, universe_systemid, ore_itemtypeid, name, texture, radius, theta, pos_x, pos_y, yield, mass
				FROM public.universe_asteroids
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Asteroid{}

		//scan into asteroid structure
		rows.Scan(&r.ID, &r.SystemID, &r.ItemTypeID, &r.Name, &r.Texture, &r.Radius, &r.Theta, &r.PosX, &r.PosY, &r.Yield, &r.Mass)

		//append to asteroid slice
		asteroids = append(asteroids, r)
	}

	return asteroids, err
}
