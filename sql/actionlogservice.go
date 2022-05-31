package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the ActionReports table in the database
type ActionReportService struct{}

// Returns a ActionReport service for interacting with ActionReports in the database
func GetActionReportService() ActionReportService {
	return ActionReportService{}
}

// Structure representing a row in the ActionReports table
type ActionReport struct {
	ID              uuid.UUID
	Report          KillLog
	Timestamp       time.Time
	InvolvedParties Meta
}

// Structure representing a row in the ActionReports table
type ActionReportSummary struct {
	ID                     uuid.UUID
	VictimIsNPC            bool
	VictimName             string
	VictimShipTemplateName string
	VictimTicker           *string
	SolarSystemName        string
	RegionName             string
	Parties                int
	Timestamp              time.Time
	SearchUserID           uuid.UUID
}

// Structure representing involved parties in an action report
type InvolvedPartiesMeta struct {
	InvolvedUserIDs []uuid.UUID
}

// Converts from a InvolvedPartiesMeta to JSON
func (a InvolvedPartiesMeta) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a InvolvedPartiesMeta
func (a *InvolvedPartiesMeta) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Finds and returns an ActionReport by its id
func (s ActionReportService) GetActionReportByID(id uuid.UUID) (*ActionReport, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// find ActionReport with this id
	actionReport := ActionReport{}

	sqlStatement :=
		`
			SELECT id, "timestamp", actionreport, involvedparties
			FROM public.actionreports
			WHERE id=$1;
		`

	row := db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&actionReport.ID, &actionReport.Timestamp,
		&actionReport.Report, &actionReport.InvolvedParties); err {
	case sql.ErrNoRows:
		return nil, errors.New("actionReport not found")
	case nil:
		return &actionReport, nil
	default:
		return nil, err
	}
}

// Creates a new action report
func (s ActionReportService) NewActionReport(e ActionReport) (*ActionReport, error) {
	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// insert action log
	sql :=
		`
			INSERT INTO public.actionreports(
				id, "timestamp", actionreport, involvedparties)
			VALUES ($1, $2, $3, $4);
		`

	uid := uuid.New()

	q, err := db.Query(sql, uid, e.Timestamp,
		e.Report, e.InvolvedParties)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	// update id in model
	e.ID = uid

	// return pointer to inserted action report model
	return &e, nil
}

// Retrieves all action reports involving a given user from the database
func (s ActionReportService) GetActionReportSummariesByUserID(userID uuid.UUID, skip int, take int) ([]ActionReportSummary, error) {
	reports := make([]ActionReportSummary, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load action reports
	sql := `
				SELECT 
					victim_isnpc, victim_name, victim_shiptemplatename, victim_ticker, 
					solarsystemname, regionname, parties, "timestamp", search_userid, id
				FROM public.vw_actionreports_summary
				WHERE search_userid = $1
				ORDER BY "timestamp" desc
				OFFSET $2
				LIMIT $3;
			`

	rows, err := db.Query(sql, userID, skip, take)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := ActionReportSummary{}

		// scan into report structure
		rows.Scan(
			&r.VictimIsNPC, &r.VictimName, &r.VictimShipTemplateName,
			&r.VictimTicker, &r.SolarSystemName, &r.RegionName,
			&r.Parties, &r.Timestamp, &r.SearchUserID, &r.ID,
		)

		// append to report slice
		reports = append(reports, r)
	}

	return reports, err
}

// Structure represening a copy-pastable report of the death of a ship
type KillLog struct {
	Header          KillLogHeader          `json:"header"`
	Fitting         KillLogFitting         `json:"fitting"`
	Cargo           []KillLogCargoItem     `json:"cargo"`
	InvolvedParties []KillLogInvolvedParty `json:"involvedParties"`
	Wallet          int64                  `json:"wallet"`
}

// Converts from a KillLog to JSON
func (a KillLog) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLog
func (a *KillLog) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing a summary header for a kill log
type KillLogHeader struct {
	VictimID               uuid.UUID `json:"victimID"`
	VictimName             string    `json:"victimName"`
	VictimFactionID        uuid.UUID `json:"victimFactionID"`
	VictimFactionName      string    `json:"victimFactionName"`
	VictimShipTemplateID   uuid.UUID `json:"victimShipTemplateID"`
	VictimShipTemplateName string    `json:"victimShipTemplateName"`
	VictimShipID           uuid.UUID `json:"victimShipID"`
	VictimShipName         string    `json:"victimShipName"`
	Timestamp              time.Time `json:"timestamp"`
	SolarSystemID          uuid.UUID `json:"solarSystemID"`
	SolarSystemName        string    `json:"solarSystemName"`
	RegionID               uuid.UUID `json:"regionID"`
	RegionName             string    `json:"regionName"`
	HoldingFactionID       uuid.UUID `json:"holdingFactionID"`
	HoldingFactionName     string    `json:"holdingFactionName"`
	InvolvedParties        int       `json:"involvedParties"`
	IsNPC                  bool      `json:"isNPC"`
	PosX                   float64   `json:"posX"`
	PosY                   float64   `json:"posY"`
}

// Converts from a KillLogHeader to JSON
func (a KillLogHeader) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogHeader
func (a *KillLogHeader) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing any combat between two ships
type KillLogInvolvedParty struct {
	// aggressor info
	UserID        uuid.UUID `json:"userID"`
	FactionID     uuid.UUID `json:"factionID"`
	CharacterName string    `json:"characterName"`
	FactionName   string    `json:"factionName"`
	IsNPC         bool      `json:"isNPC"`
	LastAggressed time.Time `json:"lastAggressed"`
	// aggressor ship info
	ShipID           uuid.UUID `json:"shipID"`
	ShipName         string    `json:"shipName"`
	ShipTemplateID   uuid.UUID `json:"shipTemplateID"`
	ShipTemplateName string    `json:"shipTemplateName"`
	// location info
	LastSolarSystemID   uuid.UUID `json:"lastSolarSystemID"`
	LastSolarSystemName string    `json:"lastSolarSystemName"`
	LastRegionID        uuid.UUID `json:"lastRegionID"`
	LastRegionName      string    `json:"lastRegionName"`
	LastPosX            float64   `json:"lastPosX"`
	LastPosY            float64   `json:"lastPosY"`
	// weapons used against victim
	WeaponUse map[string]*KillLogWeaponUse `json:"weaponUse"`
}

// Converts from a KillLogInvolvedParty to JSON
func (a KillLogHeader) KillLogInvolvedParty() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogInvolvedParty
func (a *KillLogInvolvedParty) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing a weapon used in combat between two ships
type KillLogWeaponUse struct {
	ItemID          uuid.UUID `json:"itemID"`
	ItemTypeID      uuid.UUID `json:"itemTypeID"`
	ItemFamilyID    string    `json:"itemFamilyID"`
	ItemFamilyName  string    `json:"itemFamilyName"`
	ItemTypeName    string    `json:"itemTypeName"`
	LastUsed        time.Time `json:"lastUsed"`
	DamageInflicted float64   `json:"damageInflicted"`
}

// Converts from a KillLogWeaponUse to JSON
func (a KillLogHeader) KillLogWeaponUse() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogWeaponUse
func (a *KillLogWeaponUse) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing a ship fitting for a kill log
type KillLogFitting struct {
	ARack []KillLogSlot `json:"rackA"`
	BRack []KillLogSlot `json:"rackB"`
	CRack []KillLogSlot `json:"rackC"`
}

// Converts from a KillLogFitting to JSON
func (a KillLogHeader) KillLogFitting() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogFitting
func (a *KillLogFitting) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing a slot in a kill log fitting
type KillLogSlot struct {
	ItemID              uuid.UUID `json:"itemID"`
	ItemTypeID          uuid.UUID `json:"itemTypeID"`
	ItemFamilyID        string    `json:"itemFamilyID"`
	ItemTypeName        string    `json:"itemTypeName"`
	ItemFamilyName      string    `json:"itemFamilyName"`
	IsModified          bool      `json:"isModified"`
	CustomizationFactor int       `json:"customizationFactor"`
}

// Converts from a KillLogSlot to JSON
func (a KillLogHeader) KillLogSlot() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogSlot
func (a *KillLogSlot) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Structure representing an item in the dead ship's cargo bay
type KillLogCargoItem struct {
	ItemID         uuid.UUID `json:"itemID"`
	ItemTypeID     uuid.UUID `json:"itemTypeID"`
	ItemFamilyID   string    `json:"itemFamilyID"`
	ItemTypeName   string    `json:"ItemTypeName"`
	ItemFamilyName string    `json:"itemFamilyName"`
	Quantity       int       `json:"quantity"`
	IsPackaged     bool      `json:"isPackaged"`
}

// Converts from a KillLogCargoItem to JSON
func (a KillLogHeader) KillLogCargoItem() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts from JSON to a KillLogCargoItem
func (a *KillLogCargoItem) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
