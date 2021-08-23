package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_planets table
type PlanetService struct{}

// Gets a planet service for interacting with planets in the database
func GetPlanetService() *PlanetService {
	return &PlanetService{}
}

// Structure representing a row in the universe_planets table
type Planet struct {
	ID         uuid.UUID
	SystemID   uuid.UUID
	PlanetName string
	PosX       float64
	PosY       float64
	Texture    string
	Radius     float64
	Mass       float64
	Theta      float64
}

// Retrieves all planets from the database
func (s PlanetService) GetAllPlanets() ([]Planet, error) {
	planets := make([]Planet, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load planets
	sql := `
				SELECT id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_planets;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Planet{}

		// scan into planet structure
		rows.Scan(&r.ID, &r.SystemID, &r.PlanetName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		// append to planet slice
		planets = append(planets, r)
	}

	return planets, err
}

// Retrieves all planets in a given solar system from the database
func (s PlanetService) GetPlanetsBySolarSystem(systemID uuid.UUID) ([]Planet, error) {
	planets := make([]Planet, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load planets
	sql := `
				SELECT id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_planets
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Planet{}

		// scan into planet structure
		rows.Scan(&r.ID, &r.SystemID, &r.PlanetName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		// append to planet slice
		planets = append(planets, r)
	}

	return planets, err
}

// Creates a new star in the database (for worldmaker)
func (s PlanetService) NewPlanetWorldMaker(r *Planet) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert planet
	sql := `
			INSERT INTO public.universe_planets(
				id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
			`

	q, err := db.Query(sql, r.ID, r.SystemID, r.PlanetName, r.PosX, r.PosY, r.Texture, r.Radius, r.Mass, r.Theta)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
