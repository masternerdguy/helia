package universe

import (
	"encoding/json"
	"helia/listener/models"
	"helia/shared"
	"sync"

	"github.com/google/uuid"
)

//SolarSystem Structure representing a solar system
type SolarSystem struct {
	ID         uuid.UUID
	SystemName string
	RegionID   uuid.UUID
	Ships      []*Ship
	Clients    []*shared.GameClient //clients in this system
	Lock       sync.Mutex
}

//PeriodicUpdate Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//get message registry
	msgRegistry := models.NewMessageRegistry()

	//todo: need to process the next event on each client's event queue if it is relevant to this system

	//update ships
	for _, e := range s.Ships {
		e.PeriodicUpdate()
	}

	//build global update of non-secret info for clients
	gu := models.ServerGlobalUpdateBody{}
	gu.CurrentSystemInfo = models.CurrentSystemInfo{
		ID:         s.ID,
		SystemName: s.SystemName,
	}

	for _, d := range s.Ships {
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
		})
	}

	//serialize global update
	b, _ := json.Marshal(&gu)

	msg := models.GameMessage{
		MessageType: msgRegistry.GlobalUpdate,
		MessageBody: string(b),
	}

	//write global update to clients
	for _, c := range s.Clients {
		c.WriteMessage(&msg)
	}
}

//AddShip Adds a ship to the system
func (s *SolarSystem) AddShip(c *Ship) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//make sure we aren't adding a duplicate
	for _, s := range s.Ships {
		if s.ID == c.ID {
			return
		}
	}

	//add ship
	s.Ships = append(s.Ships, c)
}

//RemoveShip Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//find the ship to remove
	e := -1
	for i, s := range s.Ships {
		if s == c {
			e = i
			break
		}
	}

	//remove ship
	if e > -1 {
		t := len(s.Ships)

		x := s.Ships[t-1]
		s.Ships[e] = x

		s.Ships = s.Ships[:t-1]
	}
}

//AddClient Adds a client to the server
func (s *SolarSystem) AddClient(c *shared.GameClient) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Clients = append(s.Clients, c)
}

//RemoveClient Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//find the client to remove
	e := -1
	for i, s := range s.Clients {
		if s == c {
			e = i
			break
		}
	}

	//remove client
	if e > -1 {
		t := len(s.Clients)

		x := s.Clients[t-1]
		s.Clients[e] = x

		s.Clients = s.Clients[:t-1]
	}
}
