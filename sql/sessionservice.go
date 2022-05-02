package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// Facility for interacting with the sessions table in the database
type SessionService struct{}

// Returns a Session service for interacting with sessions in the database
func GetSessionService() SessionService {
	return SessionService{}
}

// Structure representing a row in the sessions table
type Session struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	RemoteIP string
}

// Creates a new session
func (s SessionService) NewSession(userid uuid.UUID, remoteIP string) (*Session, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	// insert Session
	sql := `
				INSERT INTO sessions(id, userid, remoteip)
				VALUES ($1, $2, $3);
		   `

	sid := uuid.New()
	q, err := db.Query(sql, sid, userid, remoteIP)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// return session with inserted data
	Session := Session{
		ID:     sid,
		UserID: userid,
	}

	return &Session, nil
}

// Finds a session by its id
func (s SessionService) GetSessionByID(sessionid uuid.UUID) (*Session, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	// find session by id
	Session := Session{}

	sqlStatement := `SELECT id, userid
					 FROM sessions
					 WHERE id=$1`

	row := db.QueryRow(sqlStatement, sessionid)

	switch err := row.Scan(&Session.ID, &Session.UserID); err {
	case sql.ErrNoRows:
		return nil, errors.New("session does not exist")
	case nil:
		return &Session, nil
	default:
		return nil, err
	}
}

// Deletes existing sessions by userid
func (s SessionService) DeleteSession(userid uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// delete sessions
	sql := `DELETE
			FROM sessions
			WHERE userid=$1`

	q, err := db.Query(sql, userid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
