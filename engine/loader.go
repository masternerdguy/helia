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
	starSvc := sql.GetStarService()
	planetSvc := sql.GetPlanetService()
	stationSvc := sql.GetStationService()
	jumpholeSvc := sql.GetJumpholeService()

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

			for _, sh := range ships {
				es := universe.Ship{
					ID:       sh.ID,
					UserID:   sh.UserID,
					Created:  sh.Created,
					ShipName: sh.ShipName,
					PosX:     sh.PosX,
					PosY:     sh.PosY,
					SystemID: sh.SystemID,
					Texture:  sh.Texture,
					Theta:    sh.Theta,
					VelX:     sh.VelX,
					VelY:     sh.VelY,
					Accel:    sh.Accel,
					Mass:     sh.Mass,
					Radius:   sh.Radius,
					Turn:     sh.Turn,
				}

				s.AddShip(&es)
			}

			//load stars
			stars, err := starSvc.GetStarsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, st := range stars {
				star := universe.Star{
					ID:       st.ID,
					SystemID: st.SystemID,
					PosX:     st.PosX,
					PosY:     st.PosY,
					Texture:  st.Texture,
					Radius:   st.Radius,
					Mass:     st.Mass,
					Theta:    st.Theta,
				}

				s.AddStar(&star)
			}

			//load planets
			planets, err := planetSvc.GetPlanetsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, p := range planets {
				planet := universe.Planet{
					ID:         p.ID,
					SystemID:   p.SystemID,
					PlanetName: p.PlanetName,
					PosX:       p.PosX,
					PosY:       p.PosY,
					Texture:    p.Texture,
					Radius:     p.Radius,
					Mass:       p.Mass,
					Theta:      p.Theta,
				}

				s.AddPlanet(&planet)
			}

			//load jumpholes
			jumpholes, err := jumpholeSvc.GetJumpholesBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, j := range jumpholes {
				jumphole := universe.Jumphole{
					ID:           j.ID,
					SystemID:     j.SystemID,
					OutSystemID:  j.OutSystemID,
					JumpholeName: j.JumpholeName,
					PosX:         j.PosX,
					PosY:         j.PosY,
					Texture:      j.Texture,
					Radius:       j.Radius,
					Mass:         j.Mass,
					Theta:        j.Theta,
				}

				s.AddJumphole(&jumphole)
			}

			//load npc stations
			stations, err := stationSvc.GetStationsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, currStation := range stations {
				station := universe.Station{
					ID:          currStation.ID,
					SystemID:    currStation.SystemID,
					StationName: currStation.StationName,
					PosX:        currStation.PosX,
					PosY:        currStation.PosY,
					Texture:     currStation.Texture,
					Radius:      currStation.Radius,
					Mass:        currStation.Mass,
					Theta:       currStation.Theta,
				}

				s.AddStation(&station)
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
	stationSvc := sql.GetStationService()

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

			//get npc stations in system
			stations := s.CopyStations()

			//save npc stations to database
			for _, station := range stations {
				//obtain lock on copy
				station.Lock.Lock()
				defer station.Lock.Unlock()

				dbStation := sql.Station{
					ID:          station.ID,
					StationName: station.StationName,
					PosX:        station.PosX,
					PosY:        station.PosY,
					SystemID:    station.SystemID,
					Texture:     station.Texture,
					Theta:       station.Theta,
					Mass:        station.Mass,
					Radius:      station.Radius,
				}

				err := stationSvc.UpdateStation(dbStation)

				if err != nil {
					log.Println(fmt.Sprintf("Error saving station: %v | %v", dbStation, err))
				}
			}
		}
	}
}
