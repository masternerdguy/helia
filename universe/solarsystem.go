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

	//collission test between ships
	for _, sA := range s.ships {
		for _, sB := range s.ships {
			if sA.ID != sB.ID {
				//get physics dummies
				dummyA := sA.ToPhysicsDummy()
				dummyB := sB.ToPhysicsDummy()

				//get distance between ships
				d := physics.Distance(dummyA, dummyB)

				//check for radius intersection
				if d <= sA.Radius || d <= sB.Radius {

					//calculate collission results
					physics.ElasticCollide(&dummyA, &dummyB, TimeModifier)

					//update ships with results
					sA.ApplyPhysicsDummy(dummyA)
					sB.ApplyPhysicsDummy(dummyB)
				}
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
			Mass:     d.Mass,
			Radius:   d.Radius,
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
}

//AddShip Adds a ship to the system
func (s *SolarSystem) AddShip(c *Ship) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add ship
	s.ships[c.ID.String()] = c
}

//RemoveShip Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

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

//AddClient Adds a client to the system
func (s *SolarSystem) AddClient(c *shared.GameClient) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add client
	s.clients[(*c.UID).String()] = c
}

//RemoveClient Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient) {
	//safety check
	if c == nil {
		return
	}

	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

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
