package sql

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//UserService Facility for interacting with the users table in the database
type UserService struct{}

//GetUserService Returns a user service for interacting with users in the database
func GetUserService() UserService {
	return UserService{}
}

//User Structure representing a row in the users table
type User struct {
	ID            uuid.UUID
	Username      string
	Hashpass      string
	Registered    time.Time
	Banned        bool
	CurrentShipID *uuid.UUID
}

//Hashpass Hashes a user's password using their username and an internal constant as the salt
func (s UserService) Hashpass(username string, pwd string) (hash *string, err error) {
	const salt = "_4ppl3j4ck!_"
	token := []byte(fmt.Sprintf("%s-xiwmg-%s-dnjij-%s", username, pwd, salt))

	if err != nil {
		return nil, err
	}

	hp := fmt.Sprintf("%x", sha256.Sum256(token))

	return &hp, nil
}

//NewUser Creates a new user
func (s UserService) NewUser(u string, p string) (*User, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	//hash password
	hp, err := s.Hashpass(u, p)

	if err != nil {
		return nil, err
	}

	//insert user
	sql := `
				INSERT INTO public.users(id, username, hashpass, registered, banned)
				VALUES ($1, $2, $3, $4, $5);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, u, *hp, createdAt, 0)

	if err != nil {
		return nil, err
	}

	//return user with inserted data
	user := User{
		ID:            uid,
		Username:      u,
		Hashpass:      *hp,
		Registered:    createdAt,
		Banned:        false,
		CurrentShipID: nil,
	}

	return &user, nil
}

//SetCurrentShipID Sets current_shipid on a user in the database
func (s UserService) SetCurrentShipID(uid uuid.UUID, shipID *uuid.UUID) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//defer close
	defer db.Close()

	//update user
	sql := `
				UPDATE public.users SET current_shipid = $1 WHERE id = $2;
			`

	_, err = db.Query(sql, *shipID, uid)

	if err != nil {
		return err
	}

	return nil
}

//GetUserByLogin Finds the user with the supplied credentials
func (s UserService) GetUserByLogin(username string, pwd string) (*User, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	//hash password
	hp, err := s.Hashpass(username, pwd)

	if err != nil {
		return nil, err
	}

	//check for user with these credentials
	user := User{}

	sqlStatement := `SELECT id, username, hashpass, registered, banned, current_shipid
					 FROM users
					 WHERE username=$1 AND hashpass=$2`

	row := db.QueryRow(sqlStatement, username, *hp)

	switch err := row.Scan(&user.ID, &user.Username, &user.Hashpass, &user.Registered, &user.Banned, &user.CurrentShipID); err {
	case sql.ErrNoRows:
		return nil, errors.New("User does not exist or invalid credentials")
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

//GetUserByID Finds a user by its id
func (s UserService) GetUserByID(uid uuid.UUID) (*User, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//defer close
	defer db.Close()

	//check for user with these credentials
	user := User{}

	sqlStatement := `SELECT id, username, hashpass, registered, banned, current_shipid
					 FROM users
					 WHERE id=$1`

	row := db.QueryRow(sqlStatement, uid)

	switch err := row.Scan(&user.ID, &user.Username, &user.Hashpass, &user.Registered, &user.Banned, &user.CurrentShipID); err {
	case sql.ErrNoRows:
		return nil, errors.New("User does not exist or invalid credentials")
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}
