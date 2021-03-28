package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// Facility for interacting with the processes table
type ProcessService struct{}

// Gets a process service for interacting with processes in the database
func GetProcessService() *ProcessService {
	return &ProcessService{}
}

// Structure representing a row in the processes table
type Process struct {
	ID   uuid.UUID
	Name string
	Meta Meta
	Time int
}

// Retrieves all processes from the database
func (s ProcessService) GetAllProcesses() ([]Process, error) {
	processes := make([]Process, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load processes
	sql := `
				SELECT id, name, meta, time
				FROM public.processes;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Process{}

		// scan into process structure
		rows.Scan(&r.ID, &r.Name, &r.Meta, &r.Time)

		// append to process slice
		processes = append(processes, r)
	}

	return processes, err
}

// Finds and returns a process by its id
func (s ProcessService) GetProcessByID(processID uuid.UUID) (*Process, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find process with this id
	process := Process{}

	sqlStatement :=
		`
			SELECT id, name, meta, time
			FROM public.processes
			WHERE id = $1;
		`

	row := db.QueryRow(sqlStatement, processID)

	switch err := row.Scan(&process.ID, &process.Name, &process.Meta, &process.Time); err {
	case sql.ErrNoRows:
		return nil, errors.New("Process not found")
	case nil:
		return &process, nil
	default:
		return nil, err
	}
}
