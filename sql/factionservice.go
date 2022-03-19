package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// Facility for interacting with the Factions table
type FactionService struct{}

// Gets a Faction service for interacting with Factions in the database
func GetFactionService() *FactionService {
	return &FactionService{}
}

// Structure representing a row in the Factions table
type Faction struct {
	ID              uuid.UUID
	Name            string
	Description     string
	IsNPC           bool
	IsJoinable      bool
	CanHoldSov      bool
	Meta            Meta `json:"meta"`
	ReputationSheet FactionReputationSheet
	Ticker          string
	OwnerID         *uuid.UUID
	HomeStationID   *uuid.UUID
}

type ReputationSheetEntry struct {
	SourceFactionID  uuid.UUID `json:"sourceId"`
	TargetFactionID  uuid.UUID `json:"targetId"`
	StandingValue    float64   `json:"value"`
	AreOpenlyHostile bool      `json:"areOpenlyHostile"`
}

type FactionReputationSheet struct {
	Entries        map[string]ReputationSheetEntry `json:"entries"`
	HostFactionIDs []uuid.UUID                     `json:"hostFactionIds"`
	WorldPercent   float64                         `json:"worldPercent"`
}

// Converts from a ReputationSheetEntry to JSON
func (a ReputationSheetEntry) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a ReputationSheetEntry
func (a *ReputationSheetEntry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Converts from a FactionReputationSheet to JSON
func (a FactionReputationSheet) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a FactionReputationSheet
func (a *FactionReputationSheet) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Finds and returns a faction by its id
func (s FactionService) GetFactionByID(FactionID uuid.UUID) (*Faction, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find faction with this id
	t := Faction{}

	sqlStatement :=
		`
			SELECT id, name, description, isnpc, isjoinable, canholdsov, meta, ticker, reputationsheet,
			       ownerid, homestationid
			FROM public.Factions
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, FactionID)

	switch err := row.Scan(&t.ID, &t.Name, &t.Description, &t.IsNPC, &t.IsJoinable, &t.CanHoldSov,
		&t.Meta, &t.Ticker, &t.ReputationSheet, &t.OwnerID, &t.HomeStationID); err {
	case sql.ErrNoRows:
		return nil, errors.New("faction not found")
	case nil:
		return &t, nil
	default:
		return nil, err
	}
}

// Retrieves all factions from the database
func (s FactionService) GetAllFactions() ([]Faction, error) {
	factions := make([]Faction, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load factions
	sql := `
				SELECT id, name, description, isnpc, isjoinable, canholdsov, meta, ticker, reputationsheet,
				       ownerid, homestationid
				FROM public.Factions
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := Faction{}

		// scan into faction structure
		rows.Scan(&s.ID, &s.Name, &s.Description, &s.IsNPC, &s.IsJoinable, &s.CanHoldSov,
			&s.Meta, &s.Ticker, &s.ReputationSheet, &s.OwnerID, &s.HomeStationID)

		// append to ship slice
		factions = append(factions, s)
	}

	return factions, err
}

// Retrieves all player factions from the database
func (s FactionService) GetPlayerFactions() ([]Faction, error) {
	factions := make([]Faction, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load factions
	sql := `
				SELECT id, name, description, isnpc, isjoinable, canholdsov, meta, ticker, reputationsheet,
				       ownerid, homestationid
				FROM public.Factions
				WHERE isnpc = 'f' AND id != '42b937ad-0000-46e9-9af9-fc7dbf878e6a'
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := Faction{}

		// scan into faction structure
		rows.Scan(&s.ID, &s.Name, &s.Description, &s.IsNPC, &s.IsJoinable, &s.CanHoldSov,
			&s.Meta, &s.Ticker, &s.ReputationSheet, &s.OwnerID, &s.HomeStationID)

		// append to ship slice
		factions = append(factions, s)
	}

	return factions, err
}

// Creates a new faction
func (s FactionService) NewFaction(e Faction) (*Faction, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert faction
	sql := `
				INSERT INTO public.factions(
					id, name, description, isnpc, isjoinable, canholdsov, meta, ticker, reputationsheet,
				    ownerid, homestationid
				)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
			`

	uid := uuid.New()

	q, err := db.Query(sql, uid, e.Name, e.Description, e.IsNPC, e.IsJoinable, e.CanHoldSov, e.Meta, e.Ticker, e.ReputationSheet,
		e.OwnerID, e.HomeStationID)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid

	// return pointer to inserted faction model
	return &e, nil
}

// Save changes to a faction to the database
func (s FactionService) SaveFaction(e Faction) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update faction
	sql := `
				UPDATE public.factions
				SET name=$2, description=$3, meta=$4, ticker=$5, isnpc=$6, isjoinable=$7, canholdsov=$8, reputationsheet=$9,
				    ownerid=$10, homestationid=$11
				WHERE id=$1;
			`

	q, err := db.Query(sql, e.ID, e.Name, e.Description, e.Meta, e.Ticker, e.IsNPC, e.IsJoinable, e.CanHoldSov, e.ReputationSheet,
		e.OwnerID, e.HomeStationID)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
