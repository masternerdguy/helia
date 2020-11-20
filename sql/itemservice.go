package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

//ItemService Facility for interacting with the Items table in the database
type ItemService struct{}

//GetItemService Returns a Item service for interacting with Items in the database
func GetItemService() ItemService {
	return ItemService{}
}

//Item Structure representing a row in the Items table
type Item struct {
	ID         uuid.UUID
	ItemTypeID uuid.UUID
	Meta       Meta `json:"meta"`
}

//GetItemByID Finds and returns an Item by its id
func (s ItemService) GetItemByID(ItemID uuid.UUID) (*Item, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find Item with this id
	item := Item{}

	sqlStatement :=
		`
			SELECT id, itemtypeid, meta
			FROM public.Items
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemID)

	switch err := row.Scan(&item.ID, &item.ItemTypeID, &item.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("Item not found")
	case nil:
		return &item, nil
	default:
		return nil, err
	}
}
