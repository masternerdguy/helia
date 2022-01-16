package sql

import (
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the Logs table in the database
type LogService struct{}

// Returns a Log service for interacting with Logs in the database
func GetLogService() LogService {
	return LogService{}
}

// Writes a log entry to the database
func (s LogService) WriteLog(
	m string,
	t time.Time,
) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert Log
	sql := `
				INSERT INTO public.Logs(id, timestamp, message)
				VALUES ($1, $2, $3);
			`

	uid := uuid.New()

	q, err := db.Query(sql, uid, t, m)

	if err != nil {
		return err
	}

	defer q.Close()

	// return success
	return nil
}
