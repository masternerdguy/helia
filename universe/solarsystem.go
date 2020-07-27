package universe

import (
	"encoding/json"
	"helia/listener/models"
	"helia/physics"
	"helia/shared"
	"sync"

	"github.com/google/uuid"
)

//SolarSystem Structure representing a solar system
type SolarSystem struct {
	ID         uuid.UUID
	SystemName string
	RegionID   uuid.UUID
	ships      map[string]*Ship
	stars      map[string]*Star
	planets    map[string]*Planet
	jumpholes  map[string]*Jumphole
	stations   map[string]*Station
	clients    map[string]*shared.GameClient //clients in this system
	Lock       sync.Mutex
}

//Initialize Initializes internal aspects of SolarSystem
func (s *SolarSystem) Initialize() {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//initialize maps
	s.clients = make(map[string]*shared.GameClient)
	s.ships = make(map[string]*Ship)
	s.stars = make(map[string]*Star)
	s.planets = make(map[string]*Planet)
	s.jumpholes = make(map[string]*Jumphole)
	s.stations = make(map[string]*Station)
}

//PeriodicUpdate Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//get message registry
	msgRegistry := models.NewMessageRegistry()

	//process player current ship event queue
	for _, c := range s.clients {
		evt := c.PopShipEvent()

		//skip if nothing to do
		if evt == nil {
			continue
		}

		//process event
		if evt.Type == models.NewMessageRegistry().NavClick {
			//find player ship
			for _, sh := range s.ships {
				if sh.ID == c.CurrentShipID {
					//extract data
					data := evt.Body.(models.ClientNavClickBody)

					//apply effect to player's current ship
					sh.ManualTurn(data.ScreenTheta, data.ScreenMagnitude)

					//next client
					break
				}
			}
		}
	}

	//update ships
	for _, e := range s.ships {
		e.PeriodicUpdate()
	}

	//update npc stations
	for _, e := range s.stations {
		e.PeriodicUpdate()
	}

	//ship collission testing
	for _, sA := range s.ships {
		//with other ships
		for _, sB := range s.ships {
			if sA.ID != sB.ID {
				//get physics dummies
				dummyA := sA.ToPhysicsDummy()
				dummyB := sB.ToPhysicsDummy()

				//get distance between ships
				d := physics.Distance(dummyA, dummyB)

				//check for radius intersection
				if d <= (sA.TemplateData.Radius + sB.TemplateData.Radius) {
					//calculate collission results
					physics.ElasticCollide(&dummyA, &dummyB, TimeModifier)

					//update ships with results
					sA.ApplyPhysicsDummy(dummyA)
					sB.ApplyPhysicsDummy(dummyB)
				}
			}
		}

		//with jumpholes
		for _, jB := range s.jumpholes {
			//get physics dummies
			dummyA := sA.ToPhysicsDummy()
			dummyB := jB.ToPhysicsDummy()

			//get distance between ship and jumphole
			d := physics.Distance(dummyA, dummyB)

			//check for deep radius intersection
			if d <= ((sA.TemplateData.Radius + jB.Radius) * 0.75) {
				//find client
				c := s.clients[sA.UserID.String()]

				if c != nil {
					//check if this was the current ship of a player
					if sA.ID == c.CurrentShipID {
						//move player to destination system
						c.CurrentSystemID = jB.OutSystemID

						jB.OutSystem.AddClient(c, true)
						defer s.RemoveClient(c, false)
					}
				}

				//move ship to destination system
				sA.SystemID = jB.OutSystemID
				jB.OutSystem.AddShip(sA, true)
				defer s.RemoveShip(sA, false)

				//on the opposite side of the hole
				riX := jB.PosX - sA.PosX
				riY := jB.PosY - sA.PosY

				sA.PosX = (jB.OutJumphole.PosX + (riX * 1.25))
				sA.PosY = (jB.OutJumphole.PosY + (riY * 1.25))

				break
			}
		}
	}

	//build global update of non-secret info for clients
	gu := models.ServerGlobalUpdateBody{}
	gu.CurrentSystemInfo = models.CurrentSystemInfo{
		ID:         s.ID,
		SystemName: s.SystemName,
	}

	for _, d := range s.ships {
		gu.Ships = append(gu.Ships, models.GlobalShipInfo{
			ID:       d.ID,
			UserID:   d.UserID,
			Created:  d.Created,
			ShipName: d.ShipName,
			PosX:     d.PosX,
			PosY:     d.PosY,
			SystemID: d.SystemID,
			Texture:  d.Texture,
			Theta:    d.Theta,
			VelX:     d.VelX,
			VelY:     d.VelY,
			Mass:     d.GetRealMass(),
			Radius:   d.TemplateData.Radius,
			ShieldP:  (d.Shield / d.GetRealMaxShield()) * 100,
			ArmorP:   (d.Armor / d.GetRealMaxArmor()) * 100,
			HullP:    (d.Hull / d.GetRealMaxHull()) * 100,
		})
	}

	for _, d := range s.stars {
		gu.Stars = append(gu.Stars, models.GlobalStarInfo{
			ID:       d.ID,
			SystemID: d.SystemID,
			PosX:     d.PosX,
			PosY:     d.PosY,
			Texture:  d.Texture,
			Radius:   d.Radius,
			Mass:     d.Mass,
			Theta:    d.Theta,
		})
	}

	for _, d := range s.planets {
		gu.Planets = append(gu.Planets, models.GlobalPlanetInfo{
			ID:         d.ID,
			SystemID:   d.SystemID,
			PlanetName: d.PlanetName,
			PosX:       d.PosX,
			PosY:       d.PosY,
			Texture:    d.Texture,
			Radius:     d.Radius,
			Mass:       d.Mass,
			Theta:      d.Theta,
		})
	}

	for _, d := range s.jumpholes {
		gu.Jumpholes = append(gu.Jumpholes, models.GlobalJumpholeInfo{
			ID:           d.ID,
			SystemID:     d.SystemID,
			OutSystemID:  d.OutSystemID,
			JumpholeName: d.JumpholeName,
			PosX:         d.PosX,
			PosY:         d.PosY,
			Texture:      d.Texture,
			Radius:       d.Radius,
			Mass:         d.Mass,
			Theta:        d.Theta,
		})
	}

	for _, d := range s.stations {
		gu.Stations = append(gu.Stations, models.GlobalStationInfo{
			ID:          d.ID,
			SystemID:    d.SystemID,
			StationName: d.StationName,
			PosX:        d.PosX,
			PosY:        d.PosY,
			Texture:     d.Texture,
			Radius:      d.Radius,
			Mass:        d.Mass,
			Theta:       d.Theta,
		})
	}

	//serialize global update
	b, _ := json.Marshal(&gu)

	msg := models.GameMessage{
		MessageType: msgRegistry.GlobalUpdate,
		MessageBody: string(b),
	}

	//write global update to clients
	for _, c := range s.clients {
		c.WriteMessage(&msg)
	}

	//write secret current ship updates to individual clients
	for _, c := range s.clients {
		//find current ship
		d := s.ships[c.CurrentShipID.String()]

		if d == nil {
			continue
		}

		//build current ship info message
		si := models.ServerCurrentShipUpdate{
			CurrentShipInfo: models.CurrentShipInfo{
				//global stuff
				ID:       d.ID,
				UserID:   d.UserID,
				Created:  d.Created,
				ShipName: d.ShipName,
				PosX:     d.PosX,
				PosY:     d.PosY,
				SystemID: d.SystemID,
				Texture:  d.Texture,
				Theta:    d.Theta,
				VelX:     d.VelX,
				VelY:     d.VelY,
				Mass:     d.GetRealMass(),
				Radius:   d.TemplateData.Radius,
				ShieldP:  (d.Shield / d.GetRealMaxShield()) * 100,
				ArmorP:   (d.Armor / d.GetRealMaxArmor()) * 100,
				HullP:    (d.Hull / d.GetRealMaxHull()) * 100,
				//secret stuff
				EnergyP: (d.Energy / d.GetRealMaxEnergy()) * 100,
				HeatP:   (d.Heat / d.GetRealMaxHeat()) * 100,
				FuelP:   (d.Fuel / d.GetRealMaxFuel()) * 100,
			},
		}

		//serialize secret current ship update
		b, _ := json.Marshal(&si)

		sct := models.GameMessage{
			MessageType: msgRegistry.CurrentShipUpdate,
			MessageBody: string(b),
		}

		//write message to client
		c.WriteMessage(&sct)
	}
}

//AddShip Adds a ship to the system
func (s *SolarSystem) AddShip(c *Ship, lock bool) {
	//safety check
	if c == nil {
		return
	}

	if lock {
		//obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//add ship
	s.ships[c.ID.String()] = c
}

//RemoveShip Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship, lock bool) {
	//safety check
	if c == nil {
		return
	}

	if lock {
		//obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//remove ship
	delete(s.ships, c.ID.String())
}

//AddStar Adds a star to the system
func (s *SolarSystem) AddStar(c *Star) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add star
	s.stars[c.ID.String()] = c
}

//AddPlanet Adds a planet to the system
func (s *SolarSystem) AddPlanet(c *Planet) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add planet
	s.planets[c.ID.String()] = c
}

//AddJumphole Adds a jumphole to the system
func (s *SolarSystem) AddJumphole(c *Jumphole) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add jumphole
	s.jumpholes[c.ID.String()] = c
}

//AddStation Adds an NPC station to the system
func (s *SolarSystem) AddStation(c *Station) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add planet
	s.stations[c.ID.String()] = c
}

//AddClient Adds a client to the system
func (s *SolarSystem) AddClient(c *shared.GameClient, lock bool) {
	//safety check
	if c == nil {
		return
	}

	if lock {
		//obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//add client
	s.clients[(*c.UID).String()] = c
}

//RemoveClient Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient, lock bool) {
	//safety check
	if c == nil {
		return
	}

	if lock {
		//obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	//remove client
	delete(s.clients, (*c.UID).String())
}

//CopyShips Returns a copy of the ships in the system
func (s *SolarSystem) CopyShips() map[string]*Ship {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//make map for copies
	copy := make(map[string]*Ship)

	//copy ships into copy map
	for k, v := range s.ships {
		c := v.CopyShip()
		copy[k] = &c
	}

	//return copy map
	return copy
}

//CopyStations Returns a copy of the stations in the system
func (s *SolarSystem) CopyStations() map[string]*Station {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//make map for copies
	copy := make(map[string]*Station)

	//copy stations into copy map
	for k, v := range s.stations {
		c := v.CopyStation()
		copy[k] = &c
	}

	//return copy map
	return copy
}

//CopyJumpholes Returns a copy of the jumpholes in the system
func (s *SolarSystem) CopyJumpholes() map[string]*Jumphole {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//make map for copies
	copy := make(map[string]*Jumphole)

	//copy jumpholes into copy map
	for k, v := range s.jumpholes {
		c := v.CopyJumphole()
		copy[k] = &c
	}

	//return copy map
	return copy
}
