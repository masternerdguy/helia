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
	Quantity      int
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
			SELECT id, itemtypeid, meta, created, createdby, createdreason, containerid, 
				   quantity
			FROM public.Items
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemID)

	switch err := row.Scan(&item.ID, &item.ItemTypeID, &item.Meta, &item.Created, &item.CreatedBy, &item.CreatedReason, &item.ContainerID,
		&item.Quantity); err {
	case sql.ErrNoRows:
		return nil, errors.New("Item not found")
	case nil:
		return &item, nil
	default:
		return nil, err
	}
}

//GetItemsByContainer Retrieves all items in a given container
func (s ItemService) GetItemsByContainer(containerID uuid.UUID) ([]Item, error) {
	items := make([]Item, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load items
	sql :=
		`
			SELECT id, itemtypeid, meta, created, createdby, createdreason, containerid,
				   quantity
			FROM public.Items
			WHERE containerid = $1
		`

	rows, err := db.Query(sql, containerID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		i := Item{}

		//scan into item structure
		rows.Scan(&i.ID, &i.ItemTypeID, &i.Meta, &i.Created, &i.CreatedBy, &i.CreatedReason, &i.ContainerID,
			&i.Quantity)

		//append to item slice
		items = append(items, i)
	}

	return items, err
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
					id, itemtypeid, meta, created, createdby, createdreason, containerid,
					quantity
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
			`

	uid := uuid.New()
	createdAt := time.Now()

	_, err = db.Query(sql, uid, e.ItemTypeID, e.Meta, createdAt, e.CreatedBy, e.CreatedReason, e.ContainerID,
		e.Quantity)

	if err != nil {
		return nil, err
	}

	//update id in model
	e.ID = uid
	e.Created = createdAt

	//return pointer to inserted item model
	return &e, nil
}

//SetContainerID Updates the database with a new storage location for an item
func (s ItemService) SetContainerID(id uuid.UUID, containerID uuid.UUID) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//update user
	sql := `
				UPDATE public.items SET containerid = $1 WHERE id = $2;
			`

	_, err = db.Query(sql, containerID, id)

	if err != nil {
		return err
	}

	return nil
}
