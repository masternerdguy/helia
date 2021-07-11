package sql

import (
	"database/sql"
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
	ID          uuid.UUID
	Name        string
	Description string
	IsNPC       bool
	IsJoinable  bool
	IsClosed    bool
	CanHoldSov  bool
	Meta        Meta
	Ticker      string
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
			SELECT id, name, description, isnpc, isjoinable, isclosed, canholdsov, meta, ticker
			FROM public.Factions
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, FactionID)

	switch err := row.Scan(&t.ID, &t.Name, &t.Description, &t.IsNPC, &t.IsJoinable, &t.IsClosed, &t.CanHoldSov,
		&t.Meta, &t.Ticker); err {
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
				SELECT id, name, description, isnpc, isjoinable, isclosed, canholdsov, meta, ticker
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
		rows.Scan(&s.ID, &s.Name, &s.Description, &s.IsNPC, &s.IsJoinable, &s.IsClosed, &s.CanHoldSov,
			&s.Meta, &s.Ticker)

		// append to ship slice
		factions = append(factions, s)
	}

	return factions, err
}
