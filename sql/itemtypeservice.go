package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// Facility for interacting with the ItemTypes table in the database
type ItemTypeService struct{}

// Returns a ItemType service for interacting with ItemTypes in the database
func GetItemTypeService() ItemTypeService {
	return ItemTypeService{}
}

// Structure representing a row in the ItemTypes table
type ItemType struct {
	ID     uuid.UUID
	Family string
	Name   string
	Meta   Meta `json:"meta"`
}

// Finds and returns an ItemType by its id
func (s ItemTypeService) GetItemTypeByID(ItemTypeID uuid.UUID) (*ItemType, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find ItemType with this id
	ItemType := ItemType{}

	sqlStatement :=
		`
			SELECT id, family, name, meta
			FROM public.ItemTypes
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemTypeID)

	switch err := row.Scan(&ItemType.ID, &ItemType.Family, &ItemType.Name, &ItemType.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("ItemType not found")
	case nil:
		return &ItemType, nil
	default:
		return nil, err
	}
}
