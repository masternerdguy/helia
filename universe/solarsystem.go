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
	ships      map[string]*Ship
	clients    map[string]*shared.GameClient //clients in this system
	Lock       sync.Mutex
}

//Initialize Initializes internal aspects of SolarSystem
func (s *SolarSystem) Initialize() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//initialize maps
	s.clients = make(map[string]*shared.GameClient)
	s.ships = make(map[string]*Ship)
}

//PeriodicUpdate Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
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
	if c == nil {
		return
	}

	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add ship
	s.ships[c.ID.String()] = c
}

//RemoveShip Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship) {
	if c == nil {
		return
	}

	s.Lock.Lock()
	defer s.Lock.Unlock()

	//remove ship
	delete(s.ships, c.ID.String())
}

//AddClient Adds a client to the server
func (s *SolarSystem) AddClient(c *shared.GameClient) {
	if c == nil {
		return
	}

	s.Lock.Lock()
	defer s.Lock.Unlock()

	//add client
	s.clients[(*c.UID).String()] = c
}

//RemoveClient Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient) {
	if c == nil {
		return
	}

	s.Lock.Lock()
	defer s.Lock.Unlock()

	//remove client
	delete(s.clients, (*c.UID).String())
}
