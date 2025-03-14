package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_stars table
type StarService struct{}

// Gets a star service for interacting with stars in the database
func GetStarService() *StarService {
	return &StarService{}
}

// Structure representing a row in the universe_stars table
type Star struct {
	ID       uuid.UUID
	SystemID uuid.UUID
	PosX     float64
	PosY     float64
	Texture  string
	Radius   float64
	Mass     float64
	Theta    float64
	Meta     Meta
}

// Retrieves all stars from the database
func (s StarService) GetAllStars() ([]Star, error) {
	stars := make([]Star, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load stars
	sql := `
				SELECT id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta, meta
				FROM public.universe_stars;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Star{}

		// scan into star structure
		rows.Scan(&r.ID, &r.SystemID, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta, &r.Meta)

		// append to star slice
		stars = append(stars, r)
	}

	return stars, err
}

// Retrieves all stars in a given solar system from the database
func (s StarService) GetStarsBySolarSystem(systemID uuid.UUID) ([]Star, error) {
	stars := make([]Star, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load stars
	sql := `
				SELECT id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta, meta
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

		// scan into star structure
		rows.Scan(&r.ID, &r.SystemID, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta, &r.Meta)

		// append to star slice
		stars = append(stars, r)
	}

	return stars, err
}

// Creates a new star in the database (for worldmaker)
func (s StarService) NewStarWorldMaker(r *Star) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert star
	sql := `
			INSERT INTO public.universe_stars(
				id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
			`

	q, err := db.Query(sql, r.ID, r.SystemID, r.PosX, r.PosY, r.Texture, r.Radius, r.Mass, r.Theta)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Updates meta on a planet in the database (for worldfiller)
func (s StarService) UpdateMetaWorldfiller(id uuid.UUID, m *Meta) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update planet
	sql := `
				UPDATE public.universe_stars
					SET meta=$2
					WHERE id=$1;
			`

	q, err := db.Query(sql, id, m)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
