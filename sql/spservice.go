package sql

import (
	"github.com/google/uuid"
)

// Facility for for calling stored procedures in the database
type SPService struct{}

// Returns a service for calling stored procedures in the database
func GetSPService() SPService {
	return SPService{}
}

// Deletes logs, sessions, dead ships, and untracked items/containers from the database
func (s SPService) Cleanup(userid uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// run cleanup procedure
	sql := `call public.sp_cleanup();`

	q, err := db.Query(sql, userid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
