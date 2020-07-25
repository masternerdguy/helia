package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

//ShipTemplateService Facility for interacting with the shiptemplates table
type ShipTemplateService struct{}

//GetShipTemplateService Gets a shipTemplate service for interacting with shiptemplates in the database
func GetShipTemplateService() *ShipTemplateService {
	return &ShipTemplateService{}
}

//ShipTemplate Structure representing a row in the shiptemplates table
type ShipTemplate struct {
	ID               uuid.UUID
	Created          time.Time
	ShipTemplateName string
	Texture          string
	Radius           float64
	BaseAccel        float64
	BaseMass         float64
	BaseTurn         float64
	BaseShield       float64
	BaseShieldRegen  float64
	BaseArmor        float64
	BaseHull         float64
	BaseFuel         float64
	BaseHeatCap      float64
	BaseHeatSink     float64
	BaseEnergy       float64
	BaseEnergyRegen  float64
	ShipTypeID       uuid.UUID
}

//GetShipTemplateByID Finds and returns a ship template by its id
func (s ShipTemplateService) GetShipTemplateByID(shipTemplateID uuid.UUID) (*ShipTemplate, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find ship with this id
	t := ShipTemplate{}

	sqlStatement :=
		`
			SELECT id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, 
				   baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, 
				   baseenergy, baseenergyregen, shiptypeid
			FROM public.shiptemplates
			WHERE id = $1
			`

	row := db.QueryRow(sqlStatement, shipTemplateID)

	switch err := row.Scan(&t.ID, &t.Created, &t.ShipTemplateName, &t.Texture, &t.Radius, &t.BaseAccel, &t.BaseMass, &t.BaseTurn,
		&t.BaseShield, &t.BaseShieldRegen, &t.BaseArmor, &t.BaseHull, &t.BaseFuel, &t.BaseHeatCap, &t.BaseHeatSink,
		&t.BaseEnergy, &t.BaseEnergyRegen, &t.ShipTypeID); err {
	case sql.ErrNoRows:
		return nil, errors.New("Ship not found")
	case nil:
		return &t, nil
	default:
		return nil, err
	}
}
