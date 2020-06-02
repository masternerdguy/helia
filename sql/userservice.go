package sql

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//UserService Facility for interacting with the universe_systems table in the database
type UserService struct{}

//GetUserService Returns a solar system service for interacting with solar systems in the database
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
	password := []byte(pwd + salt + username)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		return nil, err
	}

	x := fmt.Sprintf("%x", sha256.Sum256(hashedPassword))

	return &x, nil
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
