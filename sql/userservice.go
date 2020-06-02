package sql

import (
	"crypto/sha256"
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
	ID         uuid.UUID
	Username   string
	Hashpass   string
	Registered time.Time
	Banned     bool
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
		ID:         uid,
		Username:   u,
		Hashpass:   *hp,
		Registered: createdAt,
		Banned:     false,
	}

	return &user, nil
}
