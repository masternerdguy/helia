package sql

import (
	"crypto/sha256"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the users table in the database
type UserService struct{}

// Returns a user service for interacting with users in the database
func GetUserService() UserService {
	return UserService{}
}

// Structure representing a row in the users table
type User struct {
	ID                uuid.UUID
	CharacterName     string
	EmailAddress      *string
	Hashpass          string
	Registered        time.Time
	Banned            bool
	CurrentShipID     *uuid.UUID
	StartID           uuid.UUID
	EscrowContainerID uuid.UUID
	CurrentFactionID  uuid.UUID
	ReputationSheet   PlayerReputationSheet
	ExperienceSheet   PlayerExperienceSheet
	// for special NPC "users"
	IsNPC         bool
	BehaviourMode *int
	// for devs
	IsDev bool
}

type PlayerExperienceSheet struct {
	ShipExperience   map[string]PlayerShipExperienceEntry   `json:"shipEntries"`
	ModuleExperience map[string]PlayerModuleExperienceEntry `json:"moduleEntries"`
}

type PlayerShipExperienceEntry struct {
	SecondsOfExperience float64   `json:"secondsOfExperience"`
	ShipTemplateID      uuid.UUID `json:"shipTemplateID"`
	ShipTemplateName    string    `json:"shipTemplateName"`
}

type PlayerModuleExperienceEntry struct {
	SecondsOfExperience float64   `json:"secondsOfExperience"`
	ItemTypeID          uuid.UUID `json:"itemTypeID"`
	ItemTypeName        string    `json:"itemTypeName"`
}

// Converts from a PlayerShipExperienceEntry to JSON
func (a PlayerShipExperienceEntry) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a PlayerShipExperienceEntry
func (a *PlayerShipExperienceEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Converts from a PlayerModuleExperienceEntry to JSON
func (a PlayerModuleExperienceEntry) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a PlayerModuleExperienceEntry
func (a *PlayerModuleExperienceEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Converts from a PlayerExperienceSheet to JSON
func (a PlayerExperienceSheet) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a PlayerExperienceSheet
func (a *PlayerExperienceSheet) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type PlayerReputationSheetFactionEntry struct {
	FactionID        uuid.UUID
	StandingValue    float64
	AreOpenlyHostile bool
}

type PlayerReputationSheet struct {
	FactionEntries map[string]PlayerReputationSheetFactionEntry `json:"factionEntries"`
}

// Converts from a PlayerReputationSheetEntry to JSON
func (a PlayerReputationSheetFactionEntry) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a PlayerReputationSheetEntry
func (a *PlayerReputationSheetFactionEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Converts from a PlayerReputationSheet to JSON
func (a PlayerReputationSheet) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a PlayerReputationSheet
func (a *PlayerReputationSheet) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Hashes a user's password using their email address and an internal constant as the salt
func (s UserService) Hashpass(emailaddress string, pwd string) (hash *string, err error) {
	const salt = "_4ppl3j4ck!_"
	token := []byte(fmt.Sprintf("%s-xiwmg-%s-dnjij-%s", emailaddress, pwd, salt))

	if err != nil {
		return nil, err
	}

	hp := fmt.Sprintf("%x", sha256.Sum256(token))

	return &hp, nil
}

// Creates a new user
func (s UserService) NewUser(
	u string,
	p string,
	startID uuid.UUID,
	escrowContainerID uuid.UUID,
	factionID uuid.UUID,
	emailAddress string,
	registeredIP string,
) (*User, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// hash password
	hp, err := s.Hashpass(emailAddress, p)

	if err != nil {
		return nil, err
	}

	// insert user
	sql := `
				INSERT INTO public.users(id, charactername, hashpass, registered, banned, startid, escrow_containerid, current_factionid, emailaddress, registeredip)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
			`

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, u, *hp, createdAt, 0, startID, escrowContainerID, factionID, emailAddress, registeredIP)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// return user with inserted data
	user := User{
		ID:               uid,
		CharacterName:    u,
		Hashpass:         *hp,
		Registered:       createdAt,
		Banned:           false,
		CurrentShipID:    nil,
		StartID:          startID,
		CurrentFactionID: factionID,
	}

	return &user, nil
}

// Sets current_shipid on a user in the database
func (s UserService) UpdateCurrentShipID(uid uuid.UUID, shipID *uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users SET current_shipid = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, *shipID, uid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Sets current_factionid on a user in the database
func (s UserService) UpdateCurrentFactionID(uid uuid.UUID, factionID *uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users SET current_factionid = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, *factionID, uid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Sets reputationsheet on a user in the database
func (s UserService) UpdateReputationSheet(uid uuid.UUID, repsheet PlayerReputationSheet) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users SET reputationsheet = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, repsheet, uid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Sets resettoken on a user in the database
func (s UserService) UpdateResetToken(uid uuid.UUID, token *uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users SET resettoken = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, token, uid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Retrieves the password reset token of a user
func (s UserService) GetPasswordResetToken(uid uuid.UUID) (*uuid.UUID, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// get reset token for this user id
	sqlStatement := `SELECT resettoken
					 FROM users
					 WHERE id=$1 and resettoken is not null`

	row := db.QueryRow(sqlStatement, uid)

	var token uuid.UUID

	switch err := row.Scan(&token); err {
	case sql.ErrNoRows:
		return nil, errors.New("user does not exist or does not have a reset token")
	case nil:
		return &token, nil
	default:
		return nil, err
	}
}

// Changes a password for a given user and token and then clears the token
func (s UserService) UpdatePassword(uid uuid.UUID, token uuid.UUID, password string) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// get user
	u, err := s.GetUserByID(uid)

	if err != nil {
		return err
	}

	// hash password
	hp, err := s.Hashpass(*u.EmailAddress, password)

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users 
				SET hashpass = $1, resettoken = null 
				WHERE id = $2 and resettoken = $3;
			`

	q, err := db.Query(sql, hp, uid, token)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Sets experiencesheet on a user in the database
func (s UserService) UpdateExperienceSheet(uid uuid.UUID, expsheet PlayerExperienceSheet) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update user
	sql := `
				UPDATE public.users SET experiencesheet = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, expsheet, uid)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Finds the user with the supplied credentials
func (s UserService) GetUserByLogin(emailaddress string, pwd string) (*User, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// hash password
	hp, err := s.Hashpass(emailaddress, pwd)

	if err != nil {
		return nil, err
	}

	// check for user with these credentials
	user := User{}

	sqlStatement := `SELECT id, charactername, hashpass, registered, banned, current_shipid, escrow_containerid, current_factionid, reputationsheet,
							isnpc, behaviourmode, emailaddress, experiencesheet, isdev
					 FROM users
					 WHERE emailaddress=$1 AND hashpass=$2`

	row := db.QueryRow(sqlStatement, emailaddress, *hp)

	switch err := row.Scan(&user.ID, &user.CharacterName, &user.Hashpass, &user.Registered, &user.Banned, &user.CurrentShipID,
		&user.EscrowContainerID, &user.CurrentFactionID, &user.ReputationSheet, &user.IsNPC, &user.BehaviourMode, &user.EmailAddress,
		&user.ExperienceSheet, &user.IsDev); err {
	case sql.ErrNoRows:
		return nil, errors.New("user does not exist or invalid credentials")
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

// Finds the id of the user with the supplied email address
func (s UserService) GetUserIdByEmailAddress(emailaddress string) (*uuid.UUID, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// retrieve user id for user with this email address
	user := User{}

	sqlStatement := `SELECT id
					 FROM users
					 WHERE emailaddress=$1 and emailaddress is not null`

	row := db.QueryRow(sqlStatement, emailaddress)

	switch err := row.Scan(&user.ID); err {
	case sql.ErrNoRows:
		return nil, errors.New("user does not exist")
	case nil:
		return &user.ID, nil
	default:
		return nil, err
	}
}

// Finds a user by its id
func (s UserService) GetUserByID(uid uuid.UUID) (*User, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// check for user with this id
	user := User{}

	sqlStatement := `SELECT id, charactername, hashpass, registered, banned, current_shipid, startid, escrow_containerid, current_factionid, reputationsheet,
							isnpc, behaviourmode, emailaddress, experiencesheet, isdev
					 FROM users
					 WHERE id=$1`

	row := db.QueryRow(sqlStatement, uid)

	switch err := row.Scan(&user.ID, &user.CharacterName, &user.Hashpass, &user.Registered, &user.Banned, &user.CurrentShipID,
		&user.StartID, &user.EscrowContainerID, &user.CurrentFactionID, &user.ReputationSheet, &user.IsNPC, &user.BehaviourMode,
		&user.EmailAddress, &user.ExperienceSheet, &user.IsDev); err {
	case sql.ErrNoRows:
		return nil, errors.New("user does not exist or invalid credentials")
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

// Retrieves all users in a given faction from the database
func (s UserService) GetUsersByFactionID(factionID uuid.UUID) ([]User, error) {
	users := make([]User, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load users
	sql := `SELECT id, charactername, hashpass, registered, banned, current_shipid, escrow_containerid, current_factionid, reputationsheet,
				   isnpc, behaviourmode, emailaddress, experiencesheet, startid, isdev
			FROM users
			WHERE current_factionid=$1`

	rows, err := db.Query(sql, factionID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := User{}

		// scan into user structure
		rows.Scan(&s.ID, &s.CharacterName, &s.Hashpass, &s.Registered, &s.Banned, &s.CurrentShipID,
			&s.EscrowContainerID, &s.CurrentFactionID, &s.ReputationSheet, &s.IsNPC, &s.BehaviourMode, &s.EmailAddress,
			&s.ExperienceSheet, &s.StartID, &s.IsDev)

		// append to user slice
		users = append(users, s)
	}

	return users, err
}
