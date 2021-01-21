package sql

import "github.com/google/uuid"

//StarService Facility for interacting with the universe_stars table
type StarService struct{}

//GetStarService Gets a star service for interacting with stars in the database
func GetStarService() *StarService {
	return &StarService{}
}

//Star Structure representing a row in the universe_stars table
type Star struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	PosX     float64
	PosY     float64
	Texture  string
	Radius   float64
	Mass     float64
	Theta    float64
}

//GetAllStars Retrieves all stars from the database
func (s StarService) GetAllStars() ([]Star, error) {
	stars := make([]Star, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load stars
	sql := `
				SELECT id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_stars;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Star{}

		//scan into star structure
		rows.Scan(&r.ID, &r.SystemID, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		//append to star slice
		stars = append(stars, r)
	}

	return stars, err
}

//GetStarsBySolarSystem Retrieves all stars in a given solar system from the database
func (s StarService) GetStarsBySolarSystem(systemID uuid.UUID) ([]Star, error) {
	stars := make([]Star, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load stars
	sql := `
				SELECT id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta
				FROM public.universe_stars
				WHERE universe_systemid = $1;
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Star{}

		//scan into star structure
		rows.Scan(&r.ID, &r.SystemID, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta)

		//append to star slice
		stars = append(stars, r)
	}

	return stars, err
}
