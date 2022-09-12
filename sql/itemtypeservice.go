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

// Structure representing a row from the vw_itemtypes_industrial view
type VwItemTypeIndustrial struct {
	ID       uuid.UUID
	Family   string
	Name     string
	MinPrice float64
	MaxPrice float64
	Volume   float64
	SiloSize float64
}

// Structure representing a row from the vw_modules_needsschematics view
type VwModuleNeedSchematic struct {
	ID     uuid.UUID
	Family string
	Name   string
	Meta   Meta `json:"meta"`
}

// Structure representing a row from the vw_ships_needsschematics view
type VwShipNeedSchematic struct {
	ID     uuid.UUID
	Family string
	Name   string
	Meta   Meta `json:"meta"`
}

// Retrieves all item types from the database
func (s ItemTypeService) GetAllItemTypes() ([]ItemType, error) {
	o := make([]ItemType, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load item types
	sql := `
				SELECT id, family, name, meta
				FROM public.ItemTypes;
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := ItemType{}

		// scan into item type structure
		rows.Scan(&i.ID, &i.Family, &i.Name, &i.Meta)

		// append to output slice
		o = append(o, i)
	}

	return o, err
}

// Finds and returns an ItemType by its id
func (s ItemTypeService) GetItemTypeByID(itemTypeID uuid.UUID) (*ItemType, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find item type with this id
	i := ItemType{}

	sqlStatement :=
		`
			SELECT id, family, name, meta
			FROM public.ItemTypes
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, itemTypeID)

	switch err := row.Scan(&i.ID, &i.Family, &i.Name, &i.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("itemType not found")
	case nil:
		return &i, nil
	default:
		return nil, err
	}
}

// Finds and returns an ItemType by its name
func (s ItemTypeService) GetItemTypeByName(name string) (*ItemType, error) {
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
			WHERE name = $1
		`

	row := db.QueryRow(sqlStatement, name)

	switch err := row.Scan(&ItemType.ID, &ItemType.Family, &ItemType.Name, &ItemType.Meta); err {
	case sql.ErrNoRows:
		return nil, errors.New("itemType not found")
	case nil:
		return &ItemType, nil
	default:
		return nil, err
	}
}

// Retrieves all item types with industrial view
func (s ItemTypeService) GetVwItemTypeIndustrials() ([]VwItemTypeIndustrial, error) {
	vws := make([]VwItemTypeIndustrial, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load item types
	sql :=
		`
			SELECT id, family, name, volume, minprice, maxprice, silosize
			FROM public.vw_itemtypes_industrial;
		`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := VwItemTypeIndustrial{}

		// scan into structure
		rows.Scan(&i.ID, &i.Family, &i.Name, &i.Volume, &i.MinPrice, &i.MaxPrice, &i.SiloSize)

		// append to slice
		vws = append(vws, i)
	}

	return vws, err
}

// Retrieves all module item types in need of schematics
func (s ItemTypeService) GetVwModuleNeedSchematics() ([]VwModuleNeedSchematic, error) {
	vws := make([]VwModuleNeedSchematic, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load item types
	sql :=
		`
			SELECT id, family, name, meta
			FROM public.vw_modules_needsschematics;
		`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := VwModuleNeedSchematic{}

		// scan into structure
		rows.Scan(&i.ID, &i.Family, &i.Name, &i.Meta)

		// append to slice
		vws = append(vws, i)
	}

	return vws, err
}

// Retrieves all ship item types in need of schematics
func (s ItemTypeService) GetVwShipNeedSchematics() ([]VwShipNeedSchematic, error) {
	vws := make([]VwShipNeedSchematic, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load item types
	sql :=
		`
			SELECT id, family, name, meta
			FROM public.vw_ships_needsschematics;
		`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := VwShipNeedSchematic{}

		// scan into structure
		rows.Scan(&i.ID, &i.Family, &i.Name, &i.Meta)

		// append to slice
		vws = append(vws, i)
	}

	return vws, err
}

// Used to save generated item types in worldfiller
func (s ItemTypeService) NewItemTypeForWorldFiller(e ItemType) (*ItemType, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert
	sql := `
				INSERT INTO public.itemtypes(
					id, family, name, meta)
				VALUES ($1, $2, $3, $4);
			`

	q, err := db.Query(sql, e.ID, e.Family, e.Name, e.Meta)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// return pointer to inserted model
	return &e, nil
}
