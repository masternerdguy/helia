package engine

import (
	"helia/sql"
	"helia/universe"
)

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
