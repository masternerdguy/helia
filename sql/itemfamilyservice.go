package sql

import (
	"database/sql"
	"errors"
)

// Facility for interacting with the ItemFamilies table in the database
type ItemFamilyService struct{}

// Returns a item family service for interacting with ItemFamilies in the database
func GetItemFamilyService() ItemFamilyService {
	return ItemFamilyService{}
}

// Structure representing a row in the ItemFamilies table
type ItemFamily struct {
	ID           string
	FriendlyName string
	Meta         Meta `json:"meta"`
}

// Retrieves all item families from the database
func (s ItemFamilyService) GetAllItemFamilies() ([]ItemFamily, error) {
	o := make([]ItemFamily, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load item families
	sql := `
				SELECT id, friendlyname, meta
				FROM public.ItemFamilies;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		f := ItemFamily{}

		// scan into structure
		rows.Scan(&f.ID, &f.FriendlyName, &f.Meta)

		// append to outout slice
		o = append(o, f)
	}

	return o, err
}

// Finds and returns an ItemFamily by its id
func (s ItemFamilyService) GetItemFamilyByID(ItemFamilyID string) (*ItemFamily, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find item family with this id
	f := ItemFamily{}

	sqlStatement :=
		`
			SELECT id, friendlyname, meta
			FROM public.ItemFamilies
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, ItemFamilyID)

	switch err := row.Scan(&f.ID, &f.FriendlyName, &f.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("itemFamily not found")
	case nil:
		return &f, nil
	default:
		return nil, err
	}
}
