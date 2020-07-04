package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

//ShipService Facility for interacting with the ships table in the database
type ShipService struct{}

//GetShipService Returns a ship service for interacting with ships in the database
func GetShipService() ShipService {
	return ShipService{}
}

//Ship Structure representing a row in the ships table
type Ship struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Created  time.Time
	ShipName string
	PosX     float64
	PosY     float64
	SystemID uuid.UUID
	Texture  string
	Theta    float64
	VelX     float64
	VelY     float64
	Accel    float64
	Radius   float64
	Mass     float64
	Turn     float64
}

//NewShip Creates a new ship
func (s ShipService) NewShip(e Ship) (*Ship, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	if err != nil {
		return nil, err
	}

	//insert ship
	sql := `
				INSERT INTO public.ships(id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture,
				                         theta, vel_x, vel_y, accel, radius, mass, turn)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8,
					    $9, $10, $11, $12, $13, $14, $15);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.SystemID, e.UserID, e.PosX, e.PosY, createdAt, e.ShipName, e.Texture,
		e.Theta, e.VelX, e.VelY, e.Accel, e.Radius, e.Mass, e.Turn)

	if err != nil {
		return nil, err
	}

	//update id in model
	e.ID = uid

	//return pointer to inserted ship model
	return &e, nil
}

//GetShipByID Finds and returns a ship by its id
func (s ShipService) GetShipByID(shipID uuid.UUID) (*Ship, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	//find ship with this id
	ship := Ship{}

	sqlStatement :=
		`
			SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture,
		           theta, vel_x, vel_y, accel, radius, mass, turn
			FROM public.ships
			WHERE id = $1
			`

	row := db.QueryRow(sqlStatement, shipID)

	switch err := row.Scan(&ship.ID, &ship.SystemID, &ship.UserID, &ship.PosX, &ship.PosY, &ship.Created, &ship.ShipName, &ship.Texture,
		&ship.Theta, &ship.VelX, &ship.VelY, &ship.Accel, &ship.Radius, &ship.Mass, &ship.Turn); err {
	case sql.ErrNoRows:
		return nil, errors.New("Ship not found")
	case nil:
		return &ship, nil
	default:
		return nil, err
	}
}

//GetShipsBySolarSystem Retrieves all ships in a solar system from the database
func (s ShipService) GetShipsBySolarSystem(systemID uuid.UUID) ([]Ship, error) {
	systems := make([]Ship, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	//load solar systems
	sql := `
				SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture,
					theta, vel_x, vel_y, accel, radius, mass, turn
				FROM public.ships
				WHERE universe_systemid = $1
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		s := Ship{}

		//scan into ship structure
		rows.Scan(&s.ID, &s.SystemID, &s.UserID, &s.PosX, &s.PosY, &s.Created, &s.ShipName, &s.Texture,
			&s.Theta, &s.VelX, &s.VelY, &s.Accel, &s.Radius, &s.Mass, &s.Turn)

		//append to ship slice
		systems = append(systems, s)
	}

	return systems, err
}
