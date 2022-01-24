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
