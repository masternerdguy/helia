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
	Theta    int
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
				                         theta)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8,
					    $9);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.SystemID, e.UserID, e.PosX, e.PosY, createdAt, e.ShipName, e.Texture,
		e.Theta)

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
		           theta
			FROM public.ships
			WHERE id = $1
			`

	row := db.QueryRow(sqlStatement, shipID)

	switch err := row.Scan(&ship.ID, &ship.SystemID, &ship.UserID, &ship.PosX, &ship.PosY, &ship.Created, &ship.ShipName, &ship.Texture,
		&ship.Theta); err {
	case sql.ErrNoRows:
		return nil, errors.New("Ship not found")
	case nil:
		return &ship, nil
	default:
		return nil, err
	}
}
