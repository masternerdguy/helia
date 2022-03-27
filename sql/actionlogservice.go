package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the ActionLogs table in the database
type ActionLogService struct{}

// Returns a ActionLog service for interacting with ActionLogs in the database
func GetActionLogService() ActionLogService {
	return ActionLogService{}
}

// Structure representing a row in the ActionLogs table
type ActionLog struct {
	ID              uuid.UUID
	Report          KillLog
	ShipID          uuid.UUID
	UserID          uuid.UUID
	FactionID       uuid.UUID
	SolarSystemID   uuid.UUID
	Timestamp       time.Time
	InvolvedUserIDs []uuid.UUID
}

// Finds and returns an ActionLog by its id
func (s ActionLogService) GetActionLogByID(id uuid.UUID) (*ActionLog, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find ActionLog with this id
	actionLog := ActionLog{}

	sqlStatement :=
		`
			SELECT id, "timestamp", shipid, systemid, actionreport, factionid, userid, involvedparties
			FROM public.actionreports
			WHERE id=$1;
		`

	row := db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&actionLog.ID, &actionLog.Timestamp, &actionLog.ShipID, &actionLog.SolarSystemID,
		&actionLog.Report, &actionLog.FactionID, &actionLog.UserID, &actionLog.InvolvedUserIDs); err {
	case sql.ErrNoRows:
		return nil, errors.New("actionLog not found")
	case nil:
		return &actionLog, nil
	default:
		return nil, err
	}
}

// Creates a new action log
func (s ActionLogService) NewActionLog(e ActionLog) (*ActionLog, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert action log
	sql :=
		`
			INSERT INTO public.actionreports(
				id, "timestamp", shipid, systemid, actionreport, factionid, userid, involvedparties)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`

	uid := uuid.New()

	q, err := db.Query(sql, e.ID, e.Timestamp, e.ShipID, e.SolarSystemID,
		e.Report, e.FactionID, e.UserID, e.InvolvedUserIDs)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid

	// return pointer to inserted action log model
	return &e, nil
}

// Structure represening a copy-pastable report of the death of a ship
type KillLog struct {
	Header KillLogHeader `json:"header"`
}

// Converts from a KillLog to JSON
func (a KillLog) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLog
func (a *KillLog) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing a summary header for a kill log
type KillLogHeader struct {
	VictimID           uuid.UUID `json:"victimID"`
	VictimName         string    `json:"victimName"`
	VictimFactionID    uuid.UUID `json:"victimFactionID"`
	VictimFactionName  string    `json:"victimFactionName"`
	VictimShipTypeID   uuid.UUID `json:"victimShipTypeID"`
	VictimShipTypeName string    `json:"victimShipTypeName"`
	VictimShipID       uuid.UUID `json:"victimShipID"`
	VictimShipName     string    `json:"victimShipName"`
	Timestamp          time.Time `json:"timestamp"`
	SolarSystemID      uuid.UUID `json:"solarSystemID"`
	SolarSystemName    string    `json:"solarSystemName"`
	RegionID           uuid.UUID `json:"regionID"`
	RegionName         string    `json:"regionName"`
	InvolvedParties    int       `json:"involvedParties"`
	IsNPC              bool      `json:"isNPC"`
}

// Converts from a KillLogHeader to JSON
func (a KillLogHeader) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogHeader
func (a *KillLogHeader) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
