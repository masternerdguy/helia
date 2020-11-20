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
	shipTmpSvc := sql.GetShipTemplateService()
	starSvc := sql.GetStarService()
	planetSvc := sql.GetPlanetService()
	stationSvc := sql.GetStationService()
	jumpholeSvc := sql.GetJumpholeService()
	userSvc := sql.GetUserService()

	u := universe.Universe{}

	//for linking jumpholes later
	jhMap := make(map[string]*universe.Jumphole)
	sMap := make(map[string]*universe.SolarSystem)

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

			//for jumphole linking later
			sMap[s.ID.String()] = &s

			//load ships
			ships, err := shipSvc.GetShipsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, sh := range ships {
				//get template
				temp, err := shipTmpSvc.GetShipTemplateByID(sh.ShipTemplateID)

				if err != nil {
					return nil, err
				}

				//get owner info
				owner, err := userSvc.GetUserByID((sh.UserID))

				if err != nil {
					return nil, err
				}

				//build in-memory ship
				es := universe.Ship{
					ID:                sh.ID,
					UserID:            sh.UserID,
					Created:           sh.Created,
					ShipName:          sh.ShipName,
					OwnerName:         owner.Username,
					PosX:              sh.PosX,
					PosY:              sh.PosY,
					SystemID:          sh.SystemID,
					Texture:           sh.Texture,
					Theta:             sh.Theta,
					VelX:              sh.VelX,
					VelY:              sh.VelY,
					Shield:            sh.Shield,
					Armor:             sh.Armor,
					Hull:              sh.Hull,
					Fuel:              sh.Fuel,
					Heat:              sh.Heat,
					Energy:            sh.Energy,
					DockedAtStationID: sh.DockedAtStationID,
					TemplateData: universe.ShipTemplate{
						ID:               temp.ID,
						Created:          temp.Created,
						ShipTemplateName: temp.ShipTemplateName,
						Texture:          temp.Texture,
						Radius:           temp.Radius,
						BaseAccel:        temp.BaseAccel,
						BaseMass:         temp.BaseMass,
						BaseTurn:         temp.BaseTurn,
						BaseShield:       temp.BaseShield,
						BaseShieldRegen:  temp.BaseShieldRegen,
						BaseArmor:        temp.BaseArmor,
						BaseHull:         temp.BaseHull,
						BaseFuel:         temp.BaseFuel,
						BaseHeatCap:      temp.BaseHeatCap,
						BaseHeatSink:     temp.BaseHeatSink,
						BaseEnergy:       temp.BaseEnergy,
						BaseEnergyRegen:  temp.BaseEnergyRegen,
						ShipTypeID:       temp.ShipTypeID,
						SlotLayout:       SlotLayoutFromSQL(&temp.SlotLayout),
					},
				}

				s.AddShip(&es, true)
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

				//for jumphole linking later
				jhMap[j.ID.String()] = &jumphole
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

	//link jumpholes
	for _, j := range jhMap {
		//get out system
		o := sMap[j.OutSystemID.String()]

		//copy jumpholes
		jhs := o.CopyJumpholes()

		//find and link destination jumphole into jumphole
		for _, k := range jhs {
			if k.OutSystemID == j.SystemID {
				//get real jumphole pointer from map
				j.OutJumphole = jhMap[k.ID.String()]

				//link destination system into jumphole
				j.OutSystem = o
				break
			}
		}
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
					ID:                ship.ID,
					UserID:            ship.UserID,
					Created:           ship.Created,
					ShipName:          ship.ShipName,
					PosX:              ship.PosX,
					PosY:              ship.PosY,
					SystemID:          ship.SystemID,
					Texture:           ship.Texture,
					Theta:             ship.Theta,
					VelX:              ship.VelX,
					VelY:              ship.VelY,
					Shield:            ship.Shield,
					Armor:             ship.Armor,
					Hull:              ship.Hull,
					Fuel:              ship.Fuel,
					Heat:              ship.Heat,
					Energy:            ship.Energy,
					ShipTemplateID:    ship.TemplateData.ID,
					DockedAtStationID: ship.DockedAtStationID,
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

//SlotLayoutFromSQL Converts a SlotLayout from the SQL type to the engine type
func SlotLayoutFromSQL(value *sql.SlotLayout) universe.SlotLayout {
	//set up empty layout
	layout := universe.SlotLayout{}

	// null check
	if value == nil {
		// return empty layout
		return layout
	}

	//copy slot data into layout
	for _, v := range value.ASlots {
		slot := SlotFromSQL(&v)
		layout.ASlots = append(layout.ASlots, slot)
	}

	for _, v := range value.BSlots {
		slot := SlotFromSQL(&v)
		layout.BSlots = append(layout.BSlots, slot)
	}

	for _, v := range value.CSlots {
		slot := SlotFromSQL(&v)
		layout.CSlots = append(layout.CSlots, slot)
	}

	log.Println(fmt.Sprintf("%v", layout))

	//return filled layout
	return layout
}

//SlotFromSQL Converts a Slot from the SQL type to the engine type
func SlotFromSQL(value *sql.Slot) universe.Slot {
	//set up empty slot
	slot := universe.Slot{}

	//null check
	if value == nil {
		// return empty slot
		return slot
	}

	//copy slot data
	slot.Family = value.Family
	slot.Volume = value.Volume
	slot.TexturePosition = value.TexturePosition

	//return filled slot
	return slot
}
