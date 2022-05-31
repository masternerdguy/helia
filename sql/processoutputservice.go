package sql

import "github.com/google/uuid"

// Facility for interacting with the processoutputs table
type ProcessOutputService struct{}

// Gets a processoutput service for interacting with processoutputs in the database
func GetProcessOutputService() *ProcessOutputService {
	return &ProcessOutputService{}
}

// Structure representing a row in the processoutputs table
type ProcessOutput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
}

// Retrieves all processoutputs from the database
func (s ProcessOutputService) GetAllProcessOutputs() ([]ProcessOutput, error) {
	processoutputs := make([]ProcessOutput, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load processoutputs
	sql := `
				SELECT id, itemtypeid, quantity, meta, processid
				FROM public.processoutputs;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := ProcessOutput{}

		// scan into processoutput structure
		rows.Scan(&r.ID, &r.ItemTypeID, &r.Quantity, &r.Meta, &r.ProcessID)

		// append to processoutput slice
		processoutputs = append(processoutputs, r)
	}

	return processoutputs, err
}

// Retrieves all processoutputs for a given process from the database
func (s ProcessOutputService) GetProcessOutputsByProcess(processID uuid.UUID) ([]ProcessOutput, error) {
	processoutputs := make([]ProcessOutput, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load processoutputs
	sql := `
				SELECT id, itemtypeid, quantity, meta, processid
				FROM public.processoutputs
				WHERE processid = $1;
			`

	rows, err := db.Query(sql, processID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := ProcessOutput{}

		// scan into processoutput structure
		rows.Scan(&r.ID, &r.ItemTypeID, &r.Quantity, &r.Meta, &r.ProcessID)

		// append to processoutput slice
		processoutputs = append(processoutputs, r)
	}

	return processoutputs, err
}

// Used by worldfiller to create process outputs
func (s ProcessOutputService) NewProcessOutputForWorldFiller(e ProcessOutput) (*ProcessOutput, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert
	sql := `
				INSERT INTO public.processoutputs(
					id, itemtypeid, quantity, meta, processid)
				VALUES ($1, $2, $3, $4, $5);
			`

	q, err := db.Query(sql, e.ID, e.ItemTypeID, e.Quantity, e.Meta, e.ProcessID)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// return pointer to inserted model
	return &e, nil
}
