package sql

import "github.com/google/uuid"

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
	InternalState Meta // todo: make a proper model for this
	Meta          Meta
}

// Retrieves all station processes from the database
func (s StationProcessService) GetAllStationProcesses() ([]StationProcess, error) {
	stationProcesses := make([]StationProcess, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load station processes
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

		//scan into stationprocess structure
		rows.Scan(&r.ID, &r.StationID, &r.ProcessID, &r.Progress, &r.Installed, &r.InternalState, &r.Meta)

		//append to stationprocess slice
		stationProcesses = append(stationProcesses, r)
	}

	return stationProcesses, err
}

// Retrieves all station processes in a given station from the database
func (s StationProcessService) GetStationProcessesByStation(stationID uuid.UUID) ([]StationProcess, error) {
	stationProcesses := make([]StationProcess, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load stationProcesses
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

		//scan into stationprocess structure
		rows.Scan(&r.ID, &r.StationID, &r.ProcessID, &r.Progress, &r.Installed, &r.InternalState, &r.Meta)

		//append to stationprocess slice
		stationProcesses = append(stationProcesses, r)
	}

	return stationProcesses, err
}
