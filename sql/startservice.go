package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

//StartService Facility for interacting with the Starts table
type StartService struct{}

//GetStartService Gets a Start service for interacting with Starts in the database
func GetStartService() *StartService {
	return &StartService{}
}

//Start Structure representing a row in the Starts table
type Start struct {
	ID             uuid.UUID
	Name           string
	ShipTemplateID uuid.UUID
	ShipFitting    StartFitting
	Created        time.Time
	Available      bool
	SystemID       uuid.UUID
	HomeStationID  uuid.UUID
}

//StartFitting Structure representing the initial fitting of a starter ship of a given start
type StartFitting struct {
	ARack []StartFittedSlot `json:"a_rack"`
	BRack []StartFittedSlot `json:"b_rack"`
	CRack []StartFittedSlot `json:"c_rack"`
}

//StartFittedSlot Structure representing a slot within the initial fitting of a starter ship of a given start
type StartFittedSlot struct {
	ItemTypeID uuid.UUID `json:"item_type_id"`
}

//Value Converts from a SlotLayout to JSON
func (a StartFitting) Value() (driver.Value, error) {
	return json.Marshal(a)
}

//Scan Converts from JSON to a SlotLayout
func (a *StartFitting) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

//GetStartByID Finds and returns a start by its id
func (s StartService) GetStartByID(StartID uuid.UUID) (*Start, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find start with this id
	t := Start{}

	sqlStatement :=
		`
			SELECT id, name, shiptemplateid, shipfitting, created, available, systemid, 
				   homestationid
			FROM public.Starts
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, StartID)

	switch err := row.Scan(&t.ID, &t.Name, &t.ShipTemplateID, &t.ShipFitting, &t.Created, &t.Available, &t.SystemID,
		&t.HomeStationID); err {
	case sql.ErrNoRows:
		return nil, errors.New("start not found")
	case nil:
		return &t, nil
	default:
		return nil, err
	}
}

//GetAllStarts Retrieves all starts from the database
func (s StartService) GetAllStarts() ([]Start, error) {
	starts := make([]Start, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load starts
	sql := `
				SELECT id, name, shiptemplateid, shipfitting, created, available, systemid,
					   homestationid
				FROM public.Starts
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		s := Start{}

		//scan into start structure
		rows.Scan(&s.ID, &s.Name, &s.ShipTemplateID, &s.ShipFitting, &s.Created, &s.Available, &s.SystemID,
			&s.HomeStationID)

		//append to ship slice
		starts = append(starts, s)
	}

	return starts, err
}
