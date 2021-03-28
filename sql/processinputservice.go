package sql

import "github.com/google/uuid"

//ProcessInputService Facility for interacting with the processinputs table
type ProcessInputService struct{}

//GetProcessInputService Gets a processinput service for interacting with processinputs in the database
func GetProcessInputService() *ProcessInputService {
	return &ProcessInputService{}
}

//ProcessInput Structure representing a row in the processinputs table
type ProcessInput struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Quantity   int
	Meta       Meta
	ProcessID  uuid.UUID
}

//GetAllProcessInputs Retrieves all processinputs from the database
func (s ProcessInputService) GetAllProcessInputs() ([]ProcessInput, error) {
	processinputs := make([]ProcessInput, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load processinputs
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

		//scan into processinput structure
		rows.Scan(&r.ID, &r.ItemTypeID, r.Quantity, r.Meta, r.ProcessID)

		//append to processinput slice
		processinputs = append(processinputs, r)
	}

	return processinputs, err
}

//GetProcessInputsByProcess Retrieves all processinputs for a given process from the database
func (s ProcessInputService) GetProcessInputsByProcess(processID uuid.UUID) ([]ProcessInput, error) {
	processinputs := make([]ProcessInput, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load processinputs
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

		//scan into processinput structure
		rows.Scan(&r.ID, &r.ItemTypeID, r.Quantity, r.Meta, r.ProcessID)

		//append to processinput slice
		processinputs = append(processinputs, r)
	}

	return processinputs, err
}
