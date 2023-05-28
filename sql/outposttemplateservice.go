package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the outposttemplates table
type OutpostTemplateService struct{}

// Gets a outpostTemplate service for interacting with outposttemplates in the database
func GetOutpostTemplateService() *OutpostTemplateService {
	return &OutpostTemplateService{}
}

// Structure representing a row in the outposttemplates table
type OutpostTemplate struct {
	ID                  uuid.UUID
	Created             time.Time
	OutpostTemplateName string
	Texture             string
	WreckTexture        string
	ExplosionTexture    string
	Radius              float64
	BaseMass            float64
	BaseShield          float64
	BaseShieldRegen     float64
	BaseArmor           float64
	BaseHull            float64
	ItemTypeID          uuid.UUID
}

// Finds and returns an outpost template by its id
func (s OutpostTemplateService) GetOutpostTemplateByID(outpostTemplateID uuid.UUID) (*OutpostTemplate, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find outpost template with this id
	t := OutpostTemplate{}

	sqlStatement :=
		`
			SELECT 
				id, created, outposttemplatename, texture, radius, basemass, baseshield, baseshieldregen, 
				basearmor, basehull, itemtypeid, wrecktexture, explosiontexture
			FROM public.outposttemplates;
		`

	row := db.QueryRow(sqlStatement, outpostTemplateID)

	switch err := row.Scan(&t.ID, &t.Created, &t.OutpostTemplateName, &t.Texture, &t.Radius, &t.BaseMass, &t.BaseShield, &t.BaseShieldRegen,
		&t.BaseArmor, &t.BaseHull, &t.ItemTypeID, &t.WreckTexture, &t.ExplosionTexture); err {
	case sql.ErrNoRows:
		return nil, errors.New("outpost template not found")
	case nil:
		return &t, nil
	default:
		return nil, err
	}
}
