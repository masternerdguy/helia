package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the ships table in the database
type ShipService struct{}

// Returns a ship service for interacting with ships in the database
func GetShipService() ShipService {
	return ShipService{}
}

// Structure representing a row in the ships table
type Ship struct {
	ID                    uuid.UUID
	UserID                uuid.UUID
	Created               time.Time
	ShipName              string
	PosX                  float64
	PosY                  float64
	SystemID              uuid.UUID
	Texture               string
	Theta                 float64
	VelX                  float64
	VelY                  float64
	Shield                float64
	Armor                 float64
	Hull                  float64
	Fuel                  float64
	Heat                  float64
	Energy                float64
	ShipTemplateID        uuid.UUID
	DockedAtStationID     *uuid.UUID
	Fitting               Fitting
	Destroyed             bool
	DestroyedAt           *time.Time
	CargoBayContainerID   uuid.UUID
	FittingBayContainerID uuid.UUID
	TrashContainerID      uuid.UUID
	ReMaxDirty            bool
	Wallet                float64
}

// JSON structure representing the module racks of a ship and what is fitted to them
type Fitting struct {
	ARack []FittedSlot `json:"a_rack"`
	BRack []FittedSlot `json:"b_rack"`
	CRack []FittedSlot `json:"c_rack"`
}

// JSON structure representing a slot within a ship's fitting rack
type FittedSlot struct {
	ItemTypeID uuid.UUID `json:"item_type_id"`
	ItemID     uuid.UUID `json:"item_id"`
}

// Converts from a Fitting to JSON
func (a Fitting) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a Fitting
func (a *Fitting) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Creates a new ship
func (s ShipService) NewShip(e Ship) (*Ship, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert ship
	sql := `
				INSERT INTO public.ships(
					id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta,
					vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
					dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid,
					fittingbay_containerid, remaxdirty, trash_containerid, wallet)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22,
						$23, $24, $25, $26, $27);
			`

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, e.SystemID, e.UserID, e.PosX, e.PosY, createdAt, e.ShipName, e.Texture, e.Theta,
		e.VelX, e.VelY, e.Shield, e.Armor, e.Hull, e.Fuel, e.Heat, e.Energy, e.ShipTemplateID, e.DockedAtStationID,
		e.Fitting, e.Destroyed, e.DestroyedAt, e.CargoBayContainerID, e.FittingBayContainerID, e.ReMaxDirty,
		e.TrashContainerID, e.Wallet)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid
	e.Created = createdAt

	// return pointer to inserted ship model
	return &e, nil
}

// Finds and returns a ship by its id
func (s ShipService) GetShipByID(shipID uuid.UUID, isDestroyed bool) (*Ship, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find ship with this id
	ship := Ship{}

	sqlStatement :=
		`
			SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, 
				   theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
				   dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid,
				   fittingbay_containerid, remaxdirty, trash_containerid, wallet
			FROM public.ships
			WHERE id = $1 and destroyed = $2
		`

	row := db.QueryRow(sqlStatement, shipID, isDestroyed)

	switch err := row.Scan(&ship.ID, &ship.SystemID, &ship.UserID, &ship.PosX, &ship.PosY, &ship.Created, &ship.ShipName, &ship.Texture,
		&ship.Theta, &ship.VelX, &ship.VelY, &ship.Shield, &ship.Armor, &ship.Hull, &ship.Fuel, &ship.Heat, &ship.Energy, &ship.ShipTemplateID,
		&ship.DockedAtStationID, &ship.Fitting, &ship.Destroyed, &ship.DestroyedAt, &ship.CargoBayContainerID,
		&ship.FittingBayContainerID, &ship.ReMaxDirty, &ship.TrashContainerID, &ship.Wallet); err {
	case sql.ErrNoRows:
		return nil, errors.New("ship not found")
	case nil:
		return &ship, nil
	default:
		return nil, err
	}
}

// Retrieves all ships in a solar system from the database
func (s ShipService) GetShipsBySolarSystem(systemID uuid.UUID, isDestroyed bool) ([]Ship, error) {
	systems := make([]Ship, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load ships
	sql := `
				SELECT id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, 
					   theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid,
					   dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid,
					   fittingbay_containerid, remaxdirty, trash_containerid, wallet
				FROM public.ships
				WHERE universe_systemid = $1 and destroyed = $2
			`

	rows, err := db.Query(sql, systemID, isDestroyed)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := Ship{}

		// scan into ship structure
		rows.Scan(&s.ID, &s.SystemID, &s.UserID, &s.PosX, &s.PosY, &s.Created, &s.ShipName, &s.Texture,
			&s.Theta, &s.VelX, &s.VelY, &s.Shield, &s.Armor, &s.Hull, &s.Fuel, &s.Heat, &s.Energy, &s.ShipTemplateID,
			&s.DockedAtStationID, &s.Fitting, &s.Destroyed, &s.DestroyedAt, &s.CargoBayContainerID,
			&s.FittingBayContainerID, &s.ReMaxDirty, &s.TrashContainerID, &s.Wallet)

		// append to ship slice
		systems = append(systems, s)
	}

	return systems, err
}

// Updates a ship in the database
func (s ShipService) UpdateShip(ship Ship) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update ship in database
	sqlStatement :=
		`
			UPDATE public.ships
			SET universe_systemid=$2, userid=$3, pos_x=$4, pos_y=$5, created=$6, shipname=$7, texture=$8, theta=$9, vel_x=$10,
				vel_y=$11, shield=$12, armor=$13, hull=$14, fuel=$15, heat=$16, energy=$17, shiptemplateid=$18, dockedat_stationid=$19,
				fitting=$20, destroyed=$21, destroyedat=$22, cargobay_containerid=$23, fittingbay_containerid=$24, remaxdirty=$25,
				trash_containerid=$26, wallet=$27
			WHERE id = $1
		`

	_, err = db.Exec(sqlStatement, ship.ID, ship.SystemID, ship.UserID, ship.PosX, ship.PosY, ship.Created, ship.ShipName, ship.Texture, ship.Theta, ship.VelX,
		ship.VelY, ship.Shield, ship.Armor, ship.Hull, ship.Fuel, ship.Heat, ship.Energy, ship.ShipTemplateID, ship.DockedAtStationID, ship.Fitting, ship.Destroyed,
		ship.DestroyedAt, ship.CargoBayContainerID, ship.FittingBayContainerID, ship.ReMaxDirty, ship.TrashContainerID, ship.Wallet)

	return err
}
