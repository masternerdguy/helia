package sql

import "github.com/google/uuid"

// Facility for interacting with the processinputs table
type ProcessInputService struct{}

// Gets a processinput service for interacting with processinputs in the database
func GetProcessInputService() *ProcessInputService {
	return &ProcessInputService{}
}

// Structure representing a row in the processinputs table
type ProcessInput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
}

// Retrieves all processinputs from the database
func (s ProcessInputService) GetAllProcessInputs() ([]ProcessInput, error) {
	processinputs := make([]ProcessInput, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load processinputs
	sql := `
				SELECT id, itemtypeid, quantity, meta, processid
				FROM public.processinputs;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := ProcessInput{}

		// scan into processinput structure
		rows.Scan(&r.ID, &r.ItemTypeID, &r.Quantity, &r.Meta, &r.ProcessID)

		// append to processinput slice
		processinputs = append(processinputs, r)
	}

	return processinputs, err
}

// Retrieves all processinputs for a given process from the database
func (s ProcessInputService) GetProcessInputsByProcess(processID uuid.UUID) ([]ProcessInput, error) {
	processinputs := make([]ProcessInput, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load processinputs
	sql := `
				SELECT id, itemtypeid, quantity, meta, processid
				FROM public.processinputs
				WHERE processid = $1;
			`

	rows, err := db.Query(sql, processID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := ProcessInput{}

		// scan into processinput structure
		rows.Scan(&r.ID, &r.ItemTypeID, &r.Quantity, &r.Meta, &r.ProcessID)

		// append to processinput slice
		processinputs = append(processinputs, r)
	}

	return processinputs, err
}

// Used by worldfiller to create process inputs
func (s ProcessInputService) NewProcessInputForWorldFiller(e ProcessInput) (*ProcessInput, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert
	sql := `
				INSERT INTO public.processinputs(
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
