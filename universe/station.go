package universe

import (
	"helia/physics"
	"sync"

	"github.com/google/uuid"
)

//Station Structure representing an NPC space station
type Station struct {
	ID          uuid.UUID
	SystemID    uuid.UUID
	StationName string
	PosX        float64
	PosY        float64
	Texture     string
	Radius      float64
	Mass        float64
	Theta       float64
	//in-memory only
	Lock           sync.Mutex
	CurrentSystem  *SolarSystem
	OpenSellOrders map[string]*SellOrder
}

//Initialize Initializes internal aspects of Station
func (s *Station) Initialize() {
	//obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//initialize maps
	s.OpenSellOrders = make(map[string]*SellOrder)
}

//PeriodicUpdate Processes the ship for a tick
func (s *Station) PeriodicUpdate() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//todo
}

//CopyStation Returns a copy of the station
func (s *Station) CopyStation() Station {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	return Station{
		ID:          s.ID,
		StationName: s.StationName,
		PosX:        s.PosX,
		PosY:        s.PosY,
		SystemID:    s.SystemID,
		Texture:     s.Texture,
		Theta:       s.Theta,
		Radius:      s.Radius,
		Mass:        s.Mass,
		//in-memory only
		Lock: sync.Mutex{},
	}
}

//ToPhysicsDummy Returns a new physics dummy structure representing this station
func (s *Station) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: s.Mass,
	}
}

//DealDamage Deals damage to the station (not yet implemented!)
func (s *Station) DealDamage(shieldDmg float64, armorDmg float64, hullDmg float64) {
	// todo
}
