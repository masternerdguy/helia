package sql

import (
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
				INSERT INTO public.ships(
				id, universe_systemid, userid, pos_x, pos_y, created, shipname)
				VALUES (?, ?, ?, ?, ?, ?, ?);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.SystemID, e.UserID, e.PosX, e.PosY, createdAt, e.ShipName)

	if err != nil {
		return nil, err
	}

	//return pointer to inserted ship model
	return &e, nil
}
