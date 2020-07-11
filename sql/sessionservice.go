package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

//SessionService Facility for interacting with the sessions table in the database
type SessionService struct{}

//GetSessionService Returns a Session service for interacting with sessions in the database
func GetSessionService() SessionService {
	return SessionService{}
}

//Session Structure representing a row in the sessions table
type Session struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

//NewSession Creates a new session
func (s SessionService) NewSession(userid uuid.UUID) (*Session, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	//insert Session
	sql := `
				INSERT INTO sessions(id, userid)
				VALUES ($1, $2);
		   `

	sid := uuid.New()
	_, err = db.Query(sql, sid, userid)

	if err != nil {
		return nil, err
	}

	//return session with inserted data
	Session := Session{
		ID:     sid,
		UserID: userid,
	}

	return &Session, nil
}

//GetSessionByID Finds a session by its id
func (s SessionService) GetSessionByID(sessionid uuid.UUID) (*Session, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	//find session by id
	Session := Session{}

	sqlStatement := `SELECT id, userid
					 FROM sessions
					 WHERE id=$1`

	row := db.QueryRow(sqlStatement, sessionid)

	switch err := row.Scan(&Session.ID, &Session.UserID); err {
	case sql.ErrNoRows:
		return nil, errors.New("Session does not exist")
	case nil:
		return &Session, nil
	default:
		return nil, err
	}
}

//DeleteSession Deletes existing sessions by userid
func (s SessionService) DeleteSession(userid uuid.UUID) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//delete sessions
	sql := `DELETE
			FROM sessions
			WHERE userid=$1`

	_, err = db.Query(sql, userid)

	if err != nil {
		return err
	}

	return nil
}
