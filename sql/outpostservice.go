package sql

import (
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the outposts table
type OutpostService struct{}

// Gets a outpost service for interacting with outposts in the database
func GetOutpostService() *OutpostService {
	return &OutpostService{}
}

// Structure representing a row in the universe_outposts table
type Outpost struct {
	ID                uuid.UUID
	SystemID          uuid.UUID
	OutpostName       string
	PosX              float64
	PosY              float64
	Theta             float64
	Shield            float64
	Armor             float64
	Hull              float64
	Wallet            float64
	UserID            uuid.UUID
	OutpostTemplateId uuid.UUID
	Created           time.Time
	Destroyed         bool
	DestroyedAt       *time.Time
}

// Retrieves all outposts in a given solar system from the database
func (s OutpostService) GetOutpostsBySolarSystem(systemID uuid.UUID, isDestroyed bool) ([]Outpost, error) {
	outposts := make([]Outpost, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load outposts
	sql := `
				SELECT 
					id, universe_systemid, outpostname, pos_x, pos_y, theta, userid, 
					shield, armor, hull, wallet, outposttemplateid
				FROM public.outposts
				WHERE universe_systemid = $1 and destroyed = $2
			`

	rows, err := db.Query(sql, systemID, isDestroyed)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Outpost{}

		// scan into outpost structure
		rows.Scan(&r.ID, &r.SystemID, &r.OutpostName, &r.PosX, &r.PosY, &r.Theta, &r.UserID,
			&r.Shield, &r.Armor, &r.Hull, &r.Wallet, &r.OutpostTemplateId)

		// append to outpost slice
		outposts = append(outposts, r)
	}

	return outposts, err
}

// Creates a new outpost
func (s OutpostService) NewOutpost(e Outpost) (*Outpost, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert outpost
	sql := `
				INSERT INTO public.outposts(
					id, universe_systemid, outpostname, pos_x, pos_y, theta, userid, shield, armor, hull, wallet, 
					destroyed, destroyedat, outposttemplateid, created, universe_stationid
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $1);
			`
	createdAt := time.Now()

	q, err := db.Query(sql, e.ID, e.SystemID, e.OutpostName, e.PosX, e.PosY, e.Theta, e.UserID, e.Shield, e.Armor, e.Hull, e.Wallet,
		false, nil, e.OutpostTemplateId, createdAt)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update in model
	e.Created = createdAt

	// return pointer to inserted outpost model
	return &e, nil
}

// Updates an outpost in the database
func (s OutpostService) UpdateOutpost(o Outpost) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update outpost in database
	sqlStatement :=
		`
			UPDATE public.outposts
			SET universe_systemid=$2, outpostname=$3, pos_x=$4, pos_y=$5, 
				theta=$6, userid=$7, shield=$8, armor=$9, hull=$10, wallet=$11, 
				destroyed=$12, destroyedat=$13, outposttemplateid=$14, created=$15
			WHERE id=$1;
		`

	_, err = db.Exec(sqlStatement, o.ID,
		o.SystemID, o.OutpostName, o.PosX, o.PosY,
		o.Theta, o.UserID, o.Shield, o.Armor, o.Hull, o.Wallet,
		o.Destroyed, o.DestroyedAt, o.OutpostTemplateId, o.Created)

	return err
}

// Updates the name of an outpost in the database
func (s OutpostService) Rename(id uuid.UUID, name string) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update outpost in database
	sqlStatement :=
		`
			UPDATE public.outposts
			SET outpostname = $2
			WHERE id = $1
		`

	_, err = db.Exec(sqlStatement, id, name)

	return err
}
