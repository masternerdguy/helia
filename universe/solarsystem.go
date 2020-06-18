package universe

import (
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
	lock       sync.Mutex
}

//PeriodicUpdate Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
	s.lock.Lock()
	defer s.lock.Unlock()

	//update ships
	for _, e := range s.Ships {
		e.PeriodicUpdate()
	}
}

//AddShip Adds a ship to the system
func (s *SolarSystem) AddShip(c *Ship) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Ships = append(s.Ships, c)
}

//RemoveShip Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship) {
	s.lock.Lock()
	defer s.lock.Unlock()

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
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Clients = append(s.Clients, c)
}

//RemoveClient Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

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
