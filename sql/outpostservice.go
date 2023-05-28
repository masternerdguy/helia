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
					destroyed, destroyedat, outposttemplateid, created
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);
			`

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, e.SystemID, e.OutpostName, e.PosX, e.PosY, e.Theta, e.UserID, e.Shield, e.Armor, e.Hull, e.Wallet,
		false, nil, e.OutpostTemplateId, createdAt)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid
	e.Created = createdAt

	// return pointer to inserted ship model
	return &e, nil
}
