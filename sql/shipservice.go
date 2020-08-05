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
	ID                uuid.UUID
	UserID            uuid.UUID
	Created           time.Time
	ShipName          string
	PosX              float64
	PosY              float64
	SystemID          uuid.UUID
	Texture           string
	Theta             float64
	VelX              float64
	VelY              float64
	Shield            float64
	Armor             float64
	Hull              float64
	Fuel              float64
	Heat              float64
	Energy            float64
	ShipTemplateID    uuid.UUID
	DockedAtStationID *uuid.UUID
}

//NewShip Creates a new ship
func (s ShipService) NewShip(e Ship) (*Ship, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//insert ship
	sql := `
				INSERT INTO public.ships(
					id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta,
					vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
					docketat_stationid)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.SystemID, e.UserID, e.PosX, e.PosY, createdAt, e.ShipName, e.Texture, e.Theta,
		e.VelX, e.VelY, e.Shield, e.Armor, e.Hull, e.Fuel, e.Heat, e.Energy, e.ShipTemplateID, e.DockedAtStationID)

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

	//find ship with this id
	ship := Ship{}

	sqlStatement :=
		`
			SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, 
				   theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
				   docketat_stationid
			FROM public.ships
			WHERE id = $1
			`

	row := db.QueryRow(sqlStatement, shipID)

	switch err := row.Scan(&ship.ID, &ship.SystemID, &ship.UserID, &ship.PosX, &ship.PosY, &ship.Created, &ship.ShipName, &ship.Texture,
		&ship.Theta, &ship.VelX, &ship.VelY, &ship.Shield, &ship.Armor, &ship.Hull, &ship.Fuel, &ship.Heat, &ship.Energy, &ship.ShipTemplateID,
		&ship.DockedAtStationID); err {
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

	//load solar systems
	sql := `
				SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, 
					   theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
					   docketat_stationid
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
			&s.Theta, &s.VelX, &s.VelY, &s.Shield, &s.Armor, &s.Hull, &s.Fuel, &s.Heat, &s.Energy, &s.ShipTemplateID,
			&s.DockedAtStationID)

		//append to ship slice
		systems = append(systems, s)
	}

	return systems, err
}

//UpdateShip Updates a ship in the database
func (s ShipService) UpdateShip(ship Ship) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//update ship in database
	sqlStatement :=
		`
			UPDATE public.ships
			SET universe_systemid=$2, userid=$3, pos_x=$4, pos_y=$5, created=$6, shipname=$7, texture=$8, theta=$9, vel_x=$10,
				vel_y=$11, shield=$12, armor=$13, hull=$14, fuel=$15, heat=$16, energy=$17, shiptemplateid=$18, docketat_stationid=$19
			WHERE id = $1
		`

	_, err = db.Exec(sqlStatement, ship.ID, ship.SystemID, ship.UserID, ship.PosX, ship.PosY, ship.Created, ship.ShipName, ship.Texture, ship.Theta, ship.VelX,
		ship.VelY, ship.Shield, ship.Armor, ship.Hull, ship.Fuel, ship.Heat, ship.Energy, ship.ShipTemplateID, ship.DockedAtStationID)

	return err
}
