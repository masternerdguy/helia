package sql

import "github.com/google/uuid"

// Facility for interacting with the universe_stations table
type StationService struct{}

// Gets a station service for interacting with stations in the database
func GetStationService() *StationService {
	return &StationService{}
}

// Structure representing a row in the universe_stations table
type Station struct {
	ID          uuid.UUID
	SystemID    uuid.UUID
	StationName string
	PosX        float64
	PosY        float64
	Texture     string
	Radius      float64
	Mass        float64
	Theta       float64
	FactionID   uuid.UUID
}

// Retrieves all stations from the database
func (s StationService) GetAllStations() ([]Station, error) {
	stations := make([]Station, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load stations
	sql := `
				SELECT id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta,
				factionid
				FROM public.universe_stations
				WHERE isoutpostshim = 'f';
			`

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Station{}

		// scan into station structure
		rows.Scan(&r.ID, &r.SystemID, &r.StationName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta,
			&r.FactionID)

		// append to station slice
		stations = append(stations, r)
	}

	return stations, err
}

// Retrieves all stations in a given solar system from the database
func (s StationService) GetStationsBySolarSystem(systemID uuid.UUID) ([]Station, error) {
	stations := make([]Station, 0)

	// get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	// load stations
	sql := `
				SELECT id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta,
				factionid
				FROM public.universe_stations
				WHERE universe_systemid = $1 and isoutpostshim = 'f';
			`

	rows, err := db.Query(sql, systemID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := Station{}

		// scan into station structure
		rows.Scan(&r.ID, &r.SystemID, &r.StationName, &r.PosX, &r.PosY, &r.Texture, &r.Radius, &r.Mass, &r.Theta,
			&r.FactionID)

		// append to station slice
		stations = append(stations, r)
	}

	return stations, err
}

// Updates an NPC station in the database
func (s StationService) UpdateStation(station Station) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// update station in database
	sqlStatement :=
		`
			UPDATE public.universe_stations
			SET universe_systemid=$2, stationname=$3, pos_x=$4, pos_y=$5, texture=$6, 
				radius=$7, mass=$8, theta=$9, factionid=$10
			WHERE id=$1;
		`

	_, err = db.Exec(sqlStatement, station.ID, station.SystemID, station.StationName, station.PosX, station.PosY, station.Texture,
		station.Radius, station.Mass, station.Theta, station.FactionID)

	return err
}

// Creates a new shim station for an outpost in the database
func (s StationService) NewStationForOutpost(r *Station) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert station
	sql := `
				INSERT INTO public.universe_stations
				(
					id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta, factionid, isoutpostshim
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 't');
			`

	q, err := db.Query(sql, r.ID, r.SystemID, r.StationName, r.PosX, r.PosY, r.Texture, r.Radius, r.Mass, r.Theta, r.FactionID)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}

// Creates a new station in the database (for worldfiller)
func (s StationService) NewStationWorldFiller(r *Station) error {
	// get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	// insert station
	sql := `
				INSERT INTO public.universe_stations
				(
					id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta, factionid
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
			`

	q, err := db.Query(sql, r.ID, r.SystemID, r.StationName, r.PosX, r.PosY, r.Texture, r.Radius, r.Mass, r.Theta, r.FactionID)

	if err != nil {
		return err
	}

	defer q.Close()

	return nil
}
