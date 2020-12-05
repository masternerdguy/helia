package sql

import (
	"database/sql"
	"errors"
	"time"

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
	ID            uuid.UUID
	ItemTypeID    uuid.UUID
	Meta          Meta `json:"meta"`
	Created       time.Time
	CreatedBy     *uuid.UUID
	CreatedReason string
	ContainerID   uuid.UUID
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
			SELECT id, itemtypeid, meta, created, createdby, createdreason, containerid
			FROM public.Items
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemID)

	switch err := row.Scan(&item.ID, &item.ItemTypeID, &item.Meta, &item.Created, &item.CreatedBy, &item.CreatedReason, &item.ContainerID); err {
	case sql.ErrNoRows:
		return nil, errors.New("Item not found")
	case nil:
		return &item, nil
	default:
		return nil, err
	}
}

//NewItem Creates a new item
func (s ItemService) NewItem(e Item) (*Item, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//insert item
	sql := `
				INSERT INTO public.items(
					id, itemtypeid, meta, created, createdby, createdreason, containerid
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.ItemTypeID, e.Meta, createdAt, e.CreatedBy, e.CreatedReason, e.ContainerID)

	if err != nil {
		return nil, err
	}

	//update id in model
	e.ID = uid
	e.Created = createdAt

	//return pointer to inserted item model
	return &e, nil
}
