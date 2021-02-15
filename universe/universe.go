package universe

import (
	"github.com/google/uuid"
)

//Heartbeat Sleep interval between solar system periodic updates in ms
const Heartbeat = 20

//SpaceDrag Space drag coefficient :)
const SpaceDrag float64 = 0.025

//Universe Structure representing the current game universe
type Universe struct {
	Regions map[string]*Region
}

//FindShip Finds the ship with the specified ID in the running game simulation
func (u *Universe) FindShip(shipID uuid.UUID) *Ship {
	//iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//look for ship in system
			sh := s.ships[shipID.String()]

			if sh != nil {
				return sh
			}
		}
	}

	return nil
}

//FindCurrentPlayerShip Finds the ship currently being flown by the specified player in the running game simulation
func (u *Universe) FindCurrentPlayerShip(userID uuid.UUID) *Ship {
	//iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//look for ship in system
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

//FindStation Finds the station with the specified ID in the running game simulation
func (u *Universe) FindStation(stationID uuid.UUID) *Station {
	//iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//look for station in system
			sh := s.stations[stationID.String()]

			if sh != nil {
				return sh
			}
		}
	}

	return nil
}
