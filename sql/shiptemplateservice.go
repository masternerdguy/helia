package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the shiptemplates table
type ShipTemplateService struct{}

// Gets a shipTemplate service for interacting with shiptemplates in the database
func GetShipTemplateService() *ShipTemplateService {
	return &ShipTemplateService{}
}

// Structure representing a row in the shiptemplates table
type ShipTemplate struct {
	ID                 uuid.UUID
	Created            time.Time
	ShipTemplateName   string
	Texture            string
	WreckTexture       string
	ExplosionTexture   string
	Radius             float64
	BaseAccel          float64
	BaseMass           float64
	BaseTurn           float64
	BaseShield         float64
	BaseShieldRegen    float64
	BaseArmor          float64
	BaseHull           float64
	BaseFuel           float64
	BaseHeatCap        float64
	BaseHeatSink       float64
	BaseEnergy         float64
	BaseEnergyRegen    float64
	ShipTypeID         uuid.UUID
	SlotLayout         SlotLayout
	BaseCargoBayVolume float64
	ItemTypeID         uuid.UUID
	CanUndock          bool
}

// JSON structure representing the slot layout available for fitting modules to this ship
type SlotLayout struct {
	ASlots []Slot `json:"a_slots"`
	BSlots []Slot `json:"b_slots"`
	CSlots []Slot `json:"c_slots"`
}

// JSON structure representing a slot in a SlotLayout
type Slot struct {
	Family          string     `json:"mod_family"`
	Volume          int        `json:"volume"`
	TexturePosition [2]float64 `json:"hp_pos"`
}

// Converts from a SlotLayout to JSON
func (a SlotLayout) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a SlotLayout
func (a *SlotLayout) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Finds and returns a ship template by its id
func (s ShipTemplateService) GetShipTemplateByID(shipTemplateID uuid.UUID) (*ShipTemplate, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find ship with this id
	t := ShipTemplate{}

	sqlStatement :=
		`
			SELECT id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, 
				   baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, 
				   baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid,
				   canundock, wrecktexture, explosiontexture
			FROM public.shiptemplates
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, shipTemplateID)

	switch err := row.Scan(&t.ID, &t.Created, &t.ShipTemplateName, &t.Texture, &t.Radius, &t.BaseAccel, &t.BaseMass, &t.BaseTurn,
		&t.BaseShield, &t.BaseShieldRegen, &t.BaseArmor, &t.BaseHull, &t.BaseFuel, &t.BaseHeatCap, &t.BaseHeatSink,
		&t.BaseEnergy, &t.BaseEnergyRegen, &t.ShipTypeID, &t.SlotLayout, &t.BaseCargoBayVolume, &t.ItemTypeID,
		&t.CanUndock, &t.WreckTexture, &t.ExplosionTexture); err {
	case sql.ErrNoRows:
		return nil, errors.New("ship template not found")
	case nil:
		return &t, nil
	default:
		return nil, err
	}
}
