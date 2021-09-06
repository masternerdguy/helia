package universe

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Sleep interval between solar system periodic updates in ms
const Heartbeat = 20

// Space drag coefficient :)
const SpaceDrag float64 = 0.025

// Structure representing the current game universe
type Universe struct {
	Regions       map[string]*Region
	Factions      map[string]*Faction
	MapData       MapData
	CachedMapData string // cached MapData to avoid overhead of extracting and stringifying over and over again
}

// Structure representing a region in a starmap
type MapDataRegion struct {
	ID      uuid.UUID `json:"id"`
	PosX    float64   `json:"x"`
	PosY    float64   `json:"y"`
	Name    string    `json:"name"`
	Systems []MapDataSystem
}

// Structure representing a solar system in a starmap
type MapDataSystem struct {
	ID   uuid.UUID `json:"id"`
	PosX float64   `json:"x"`
	PosY float64   `json:"y"`
	Name string    `json:"name"`
}

// Structure representing a jumphole connection in a starmap
type MapDataEdge struct {
	StartSystemId uuid.UUID `json:"aID"`
	EndSystemId   uuid.UUID `json:"bID"`
}

// Structure representing a starmap
type MapData struct {
	Regions []MapDataRegion `json:"regions"`
	Edges   []MapDataEdge   `json:"edges"`
}

// Builds the starmap and the shared starmap cache string
func (u *Universe) BuildMapWithCache() error {
	// empty map
	data := MapData{}

	// iterate over all systems in all regions
	for _, r := range u.Regions {
		// map region
		reg := MapDataRegion{
			ID:   r.ID,
			PosX: r.PosX,
			PosY: r.PosY,
			Name: r.RegionName,
		}

		for _, s := range r.Systems {
			// copy jumpholes
			jhs := s.CopyJumpholes()

			// map system into region
			sys := MapDataSystem{
				ID:   s.ID,
				PosX: s.PosX,
				PosY: s.PosY,
				Name: s.SystemName,
			}

			reg.Systems = append(reg.Systems, sys)

			// store edges
			for _, j := range jhs {
				edge := MapDataEdge{
					StartSystemId: j.SystemID,
					EndSystemId:   j.OutSystemID,
				}

				data.Edges = append(data.Edges, edge)
			}
		}

		// store region
		data.Regions = append(data.Regions, reg)
	}

	// store map
	u.MapData = data

	// cache map data as json
	b, err := json.Marshal(data)
	u.CachedMapData = string(b)

	return err
}

// Finds the ship with the specified ID in the running game simulation
func (u *Universe) FindShip(shipID uuid.UUID) *Ship {
	// iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// look for ship in system
			sh := s.ships[shipID.String()]

			if sh != nil {
				return sh
			}
		}
	}

	return nil
}

// Finds the ship currently being flown by the specified player in the running game simulation
func (u *Universe) FindCurrentPlayerShip(userID uuid.UUID) *Ship {
	// iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// look for ship in system
			for _, shx := range s.ships {
				if shx != nil {
					// id and flown check
					if shx.BeingFlownByPlayer && shx.UserID == userID {
						return shx
					}
				}
			}
		}
	}

	return nil
}

// Finds the station with the specified ID in the running game simulation
func (u *Universe) FindStation(stationID uuid.UUID) *Station {
	// iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// look for station in system
			sh := s.stations[stationID.String()]

			if sh != nil {
				return sh
			}
		}
	}

	return nil
}
