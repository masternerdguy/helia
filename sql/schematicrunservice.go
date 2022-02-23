package sql

import (
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the schematicruns table
type SchematicRunService struct{}

// Gets a schematicRun service for interacting with schematicruns in the database
func GetSchematicRunService() *SchematicRunService {
	return &SchematicRunService{}
}

// Structure representing a row in the schematicruns table
type SchematicRun struct {
	ID              uuid.UUID
	Created         time.Time
	ProcessID       uuid.UUID
	StatusID        string
	Progress        int
	SchematicItemID uuid.UUID
	UserID          uuid.UUID
}

// Retrieves all undelivered schematicruns from the database
func (s SchematicRunService) GetUndeliveredSchematicRuns() ([]SchematicRun, error) {
	schematicruns := make([]SchematicRun, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load schematicruns
	sql := `
				SELECT id, created, processid, statusid, progress, schematic_itemid, userid
				FROM public.schematicruns
				WHERE statusid != 'delivered';
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := SchematicRun{}

		// scan into schematicRun structure
		rows.Scan(&r.ID, &r.Created, &r.ProcessID, &r.StatusID, &r.Progress, &r.SchematicItemID, &r.UserID)

		// append to schematicRun slice
		schematicruns = append(schematicruns, r)
	}

	return schematicruns, err
}

// Creates a new schematic run
func (s SchematicRunService) NewSchematicRun(e SchematicRun) (*SchematicRun, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert sell order
	sql := `
			INSERT INTO public.schematicruns(
				id, created, processid, statusid, progress, schematic_itemid, userid)
				VALUES ($1, $2, $3, $4, $5, $6, $7);
		   `

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, createdAt, e.ProcessID, e.StatusID, e.Progress, e.SchematicItemID, e.UserID)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid
	e.Created = createdAt

	// return pointer to inserted schematic run model
	return &e, nil
}

// Saves the progress and status of a schematic run
func (s SchematicRunService) UpdateSchematicRun(e SchematicRun) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update schematic run
	sql := `
			UPDATE public.schematicruns
			SET StatusID=$2, Progress=$3
			WHERE id = $1
		   `

	q, err := db.Query(sql, e.ID, e.StatusID, e.Progress)

	if err != nil {
		return err
	}

	defer q.Close()

	// success
	return nil
}
