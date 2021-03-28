package sql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// Facility for interacting with the stationprocesses table
type StationProcessService struct{}

// Gets a stationprocess service for interacting with stationProcesses in the database
func GetStationProcessService() *StationProcessService {
	return &StationProcessService{}
}

// Structure representing a row in the stationprocesses table
type StationProcess struct {
	ID            uuid.UUID
	StationID     uuid.UUID
	ProcessID     uuid.UUID
	Progress      int
	Installed     bool
	InternalState StationProcessInternalState
	Meta          Meta
}

// JSON structure representing the internal state of the ware silos involved in the process
type StationProcessInternalState struct {
	IsRunning bool                                         `json:"isRunning"`
	Inputs    map[string]StationProcessInternalStateFactor `json:"inputs"`
	Outputs   map[string]StationProcessInternalStateFactor `json:"outputs"`
}

// JSON structure representing an input or output factor in a station process's internal state
type StationProcessInternalStateFactor struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

// Converts from a StationProcessInternalState to JSON
func (a StationProcessInternalState) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a StationProcessInternalState
func (a *StationProcessInternalState) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Retrieves all station processes from the database
func (s StationProcessService) GetAllStationProcesses() ([]StationProcess, error) {
	stationProcesses := make([]StationProcess, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load station processes
	sql := `
				SELECT id, universe_stationid, processid, progress, installed, internalstate, meta
				FROM public.stationprocesses;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := StationProcess{}

		// scan into stationprocess structure
		rows.Scan(&r.ID, &r.StationID, &r.ProcessID, &r.Progress, &r.Installed, &r.InternalState, &r.Meta)

		// append to stationprocess slice
		stationProcesses = append(stationProcesses, r)
	}

	return stationProcesses, err
}

// Retrieves all station processes in a given station from the database
func (s StationProcessService) GetStationProcessesByStation(stationID uuid.UUID) ([]StationProcess, error) {
	stationProcesses := make([]StationProcess, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load stationProcesses
	sql := `
				SELECT id, universe_stationid, processid, progress, installed, internalstate, meta
				FROM public.stationprocesses
				WHERE universe_stationid = $1;
			`

	rows, err := db.Query(sql, stationID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := StationProcess{}

		// scan into stationprocess structure
		rows.Scan(&r.ID, &r.StationID, &r.ProcessID, &r.Progress, &r.Installed, &r.InternalState, &r.Meta)

		// append to stationprocess slice
		stationProcesses = append(stationProcesses, r)
	}

	return stationProcesses, err
}
