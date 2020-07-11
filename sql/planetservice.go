package sql

import "github.com/google/uuid"

//PlanetService Facility for interacting with the universe_planets table
type PlanetService struct{}

//GetPlanetService Gets a planet service for interacting with planets in the database
func GetPlanetService() *PlanetService {
	return &PlanetService{}
}

//Planet Structure representing a row in the universe_planets table
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

//GetAllPlanets Retrieves all planets from the database
func (s PlanetService) GetAllPlanets() ([]Planet, error) {
	planets := make([]Planet, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load planets
	sql := `
				SELECT id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_planets;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := Planet{}

		//scan into planet structure
		rows.Scan(&r.ID, &r.SystemID, &r.PlanetName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		//append to planet slice
		planets = append(planets, r)
	}

	return planets, err
}

//GetPlanetsBySolarSystem Retrieves all planets in a given solar system from the database
func (s PlanetService) GetPlanetsBySolarSystem(systemID uuid.UUID) ([]Planet, error) {
	planets := make([]Planet, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load planets
	sql := `
				SELECT id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_planets
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := Planet{}

		//scan into planet structure
		rows.Scan(&r.ID, &r.SystemID, &r.PlanetName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		//append to planet slice
		planets = append(planets, r)
	}

	return planets, err
}
