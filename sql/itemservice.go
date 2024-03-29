package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the Items table in the database
type ItemService struct{}

// Returns a Item service for interacting with Items in the database
func GetItemService() ItemService {
	return ItemService{}
}

// Structure representing a row in the Items table
type Item struct {
	ID            uuid.UUID
	ItemTypeID    uuid.UUID
	Meta          Meta `json:"meta"`
	Created       time.Time
	CreatedBy     *uuid.UUID
	CreatedReason string
	ContainerID   uuid.UUID
	Quantity      int
	IsPackaged    bool
}

// Finds and returns an Item by its id
func (s ItemService) GetItemByID(ItemID uuid.UUID) (*Item, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find Item with this id
	item := Item{}

	sqlStatement :=
		`
			SELECT id, itemtypeid, meta, created, createdby, createdreason, containerid, 
				   quantity, ispackaged
			FROM public.Items
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemID)

	switch err := row.Scan(&item.ID, &item.ItemTypeID, &item.Meta, &item.Created, &item.CreatedBy, &item.CreatedReason, &item.ContainerID,
		&item.Quantity, &item.IsPackaged); err {
	case sql.ErrNoRows:
		return nil, errors.New("item not found")
	case nil:
		return &item, nil
	default:
		return nil, err
	}
}

// Retrieves all items in a given container
func (s ItemService) GetItemsByContainer(containerID uuid.UUID) ([]Item, error) {
	items := make([]Item, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load items
	sql :=
		`
			SELECT id, itemtypeid, meta, created, createdby, createdreason, containerid,
				   quantity, ispackaged
			FROM public.Items
			WHERE containerid = $1
			AND quantity > 0
		`

	rows, err := db.Query(sql, containerID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := Item{}

		// scan into item structure
		rows.Scan(&i.ID, &i.ItemTypeID, &i.Meta, &i.Created, &i.CreatedBy, &i.CreatedReason, &i.ContainerID,
			&i.Quantity, &i.IsPackaged)

		// append to item slice
		items = append(items, i)
	}

	return items, err
}

// Creates a new item
func (s ItemService) NewItem(e Item) (*Item, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert item
	sql := `
				INSERT INTO public.items(
					id, itemtypeid, meta, created, createdby, createdreason, containerid,
					quantity, ispackaged
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
			`

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, e.ItemTypeID, e.Meta, createdAt, e.CreatedBy, e.CreatedReason, e.ContainerID,
		e.Quantity, e.IsPackaged)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid
	e.Created = createdAt

	// return pointer to inserted item model
	return &e, nil
}

// Updates the database with a new storage location for an item
func (s ItemService) SetContainerID(id uuid.UUID, containerID uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update item
	sql := `
				UPDATE public.items SET containerid = $1 WHERE id = $2;
			`

	q, err := db.Query(sql, containerID, id)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Updates the database to make an item packaged
func (s ItemService) PackageItem(id uuid.UUID) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update item
	sql := `
				UPDATE public.items SET meta='{}', ispackaged='t' WHERE id = $1;
			`

	q, err := db.Query(sql, id)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Updates the database to make an item unpackaged
func (s ItemService) UnpackageItem(id uuid.UUID, meta Meta) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update item
	sql := `
				UPDATE public.items SET meta=$2, ispackaged='f' WHERE id = $1;
			`

	q, err := db.Query(sql, id, meta)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Updates the database to change the quantity of an item stack
func (s ItemService) ChangeQuantity(id uuid.UUID, quantity int) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update item
	sql := `
				UPDATE public.items SET quantity=$2 WHERE id = $1;
			`

	q, err := db.Query(sql, id, quantity)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Updates the Metadata on an item
func (s ItemService) ChangeMeta(id uuid.UUID, meta Meta) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update item
	sql := `
				UPDATE public.items SET meta=$2 WHERE id = $1;
			`

	q, err := db.Query(sql, id, meta)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
