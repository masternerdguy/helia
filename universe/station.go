package universe

import (
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
	Lock sync.Mutex
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
