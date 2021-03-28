package sql

import (
	"database/sql"
	"errors"
)

// Facility for interacting with the ItemFamilies table in the database
type ItemFamilyService struct{}

// Returns a ItemFamily service for interacting with ItemFamilies in the database
func GetItemFamilyService() ItemFamilyService {
	return ItemFamilyService{}
}

// Structure representing a row in the ItemFamilies table
type ItemFamily struct {
	ID           string
	FriendlyName string
	Meta         Meta `json:"meta"`
}

// Finds and returns an ItemFamily by its id
func (s ItemFamilyService) GetItemFamilyByID(ItemFamilyID string) (*ItemFamily, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find ItemFamily with this id
	ItemFamily := ItemFamily{}

	sqlStatement :=
		`
			SELECT id, friendlyname, meta
			FROM public.ItemFamilies
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemFamilyID)

	switch err := row.Scan(&ItemFamily.ID, &ItemFamily.FriendlyName, &ItemFamily.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("ItemFamily not found")
	case nil:
		return &ItemFamily, nil
	default:
		return nil, err
	}
}
