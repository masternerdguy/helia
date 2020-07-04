package engine

import (
	"fmt"
	"helia/sql"
	"helia/universe"
	"log"
)

//loadUniverse Loads the state of the universe from the database
func loadUniverse() (*universe.Universe, error) {
	//get services
	regionSvc := sql.GetRegionService()
	systemSvc := sql.GetSolarSystemService()
	shipSvc := sql.GetShipService()

	u := universe.Universe{}

	//load regions
	rs, err := regionSvc.GetAllRegions()

	if err != nil {
		return nil, err
	}

	regions := make(map[string]*universe.Region, 0)
	for _, e := range rs {
		//load basic region information
		r := universe.Region{
			ID:         e.ID,
			RegionName: e.RegionName,
		}

		//load systems in region
		ss, err := systemSvc.GetSolarSystemsByRegion(e.ID)

		if err != nil {
			return nil, err
		}

		systems := make(map[string]*universe.SolarSystem, 0)

		for _, f := range ss {
			s := universe.SolarSystem{
				ID:         f.ID,
				SystemName: f.SystemName,
				RegionID:   f.RegionID,
			}

			//initialize and store system
			s.Initialize()
			systems[s.ID.String()] = &s

			//load ships
			ships, err := shipSvc.GetShipsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, currShip := range ships {
				es := universe.Ship{
					ID:       currShip.ID,
					UserID:   currShip.UserID,
					Created:  currShip.Created,
					ShipName: currShip.ShipName,
					PosX:     currShip.PosX,
					PosY:     currShip.PosY,
					SystemID: currShip.SystemID,
					Texture:  currShip.Texture,
					Theta:    currShip.Theta,
					VelX:     currShip.VelX,
					VelY:     currShip.VelY,
					Accel:    currShip.Accel,
					Mass:     currShip.Mass,
					Radius:   currShip.Radius,
					Turn:     currShip.Turn,
				}

				s.AddShip(&es)
			}
		}

		//store and append region
		r.Systems = systems
		regions[r.ID.String()] = &r
	}

	//link regions into universe
	u.Regions = regions

	return &u, nil
}

//saveUniverse Saves the current state of dynamic entities in the simulation to the database
func saveUniverse(u *universe.Universe) {
	//get services
	shipSvc := sql.GetShipService()

	//iterate over systems
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//get ships in system
			ships := s.CopyShips()

			//save ships to database
			for _, ship := range ships {
				//obtain lock on copy
				ship.Lock.Lock()
				defer ship.Lock.Unlock()

				dbShip := sql.Ship{
					ID:       ship.ID,
					UserID:   ship.UserID,
					Created:  ship.Created,
					ShipName: ship.ShipName,
					PosX:     ship.PosX,
					PosY:     ship.PosY,
					SystemID: ship.SystemID,
					Texture:  ship.Texture,
					Theta:    ship.Theta,
					VelX:     ship.VelX,
					VelY:     ship.VelY,
					Accel:    ship.Accel,
					Mass:     ship.Mass,
					Radius:   ship.Radius,
					Turn:     ship.Turn,
				}

				err := shipSvc.UpdateShip(dbShip)

				if err != nil {
					log.Println(fmt.Sprintf("Error saving ship: %v | %v", dbShip, err))
				}
			}
		}
	}
}
