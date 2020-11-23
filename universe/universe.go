package universe

import "github.com/google/uuid"

//Heartbeat Sleep interval between solar system periodic updates in ms
const Heartbeat = 20

//SpaceDrag Space drag coefficient :)
const SpaceDrag float64 = 0.025

//Universe Structure representing the current game universe
type Universe struct {
	Starts  map[string]*Start
	Regions map[string]*Region
}

//FindShip Finds the ship with the specified ID in the running game simulation
func (u *Universe) FindShip(shipID uuid.UUID) *Ship {
	//iterate over all systems in all regions
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//lock system
			s.Lock.Lock()
			defer s.Lock.Unlock()

			//look for ship in system
			sh := s.ships[shipID.String()]

			if sh != nil {
				return sh
			}
		}
	}

	return nil
}

//FindStart Finds the start with the specified ID in the start cache
func (u *Universe) FindStart(startID uuid.UUID) *Start {
	x, f := u.Starts[startID.String()]

	if f {
		return x
	}

	return nil
}
