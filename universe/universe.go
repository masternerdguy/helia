package universe

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/physics"
	"helia/shared"
	"sync"

	"github.com/google/uuid"
)

// Sleep interval between solar system periodic updates in ms
const Heartbeat = 20

// Space drag coefficient :)
const SpaceDrag float64 = 0.025

// Minimum transient jumphole pairs at startup
const MinTransientEdges = 5

// Maximum transient jumphole pairs at startup
const MaxTransientEdges = 35

// Structure representing the current game universe
type Universe struct {
	Regions       map[string]*Region
	Factions      map[string]*Faction
	AllSystems    []*SolarSystem
	MapData       MapData
	CachedMapData string // cached MapData to avoid overhead of extracting and stringifying over and over again
	FactionsLock  sync.RWMutex
}

// Structure representing a region in a starmap
type MapDataRegion struct {
	ID      uuid.UUID       `json:"id"`
	PosX    float64         `json:"x"`
	PosY    float64         `json:"y"`
	Name    string          `json:"name"`
	Systems []MapDataSystem `json:"systems"`
}

// Structure representing a solar system in a starmap
type MapDataSystem struct {
	ID        uuid.UUID `json:"id"`
	PosX      float64   `json:"x"`
	PosY      float64   `json:"y"`
	Name      string    `json:"name"`
	FactionID uuid.UUID `json:"factionId"`
}

// Structure representing a jumphole connection in a starmap
type MapDataEdge struct {
	StartSystemId uuid.UUID `json:"aID"`
	EndSystemId   uuid.UUID `json:"bID"`
	Transient     bool      `json:"transient"`
}

// Structure representing a starmap
type MapData struct {
	Regions []MapDataRegion `json:"regions"`
	Edges   []MapDataEdge   `json:"edges"`
}

// Builds the starmap and the shared starmap cache string
func (u *Universe) BuildMapWithCache() error {
	// empty map
	data := MapData{}

	// iterate over all systems in all regions
	for _, r := range u.Regions {
		// map region
		reg := MapDataRegion{
			ID:   r.ID,
			PosX: r.PosX,
			PosY: r.PosY,
			Name: r.RegionName,
		}

		for _, s := range r.Systems {
			// copy jumpholes
			jhs := s.CopyJumpholes(true)

			// map system into region
			sys := MapDataSystem{
				ID:        s.ID,
				PosX:      s.PosX,
				PosY:      s.PosY,
				Name:      s.SystemName,
				FactionID: s.HoldingFactionID,
			}

			reg.Systems = append(reg.Systems, sys)

			// store edges
			for _, j := range jhs {
				edge := MapDataEdge{
					StartSystemId: j.SystemID,
					EndSystemId:   j.OutSystemID,
					Transient:     j.Transient,
				}

				data.Edges = append(data.Edges, edge)
			}
		}

		// store region
		data.Regions = append(data.Regions, reg)
	}

	// store map
	u.MapData = data

	// cache map data as json
	b, err := json.Marshal(data)
	u.CachedMapData = string(b)

	return err
}

// Generates transient objects that are not stored in the DB and will go away upon server restart
func (u *Universe) BuildTransientCelestials() {
	// capture region list
	allRegions := make([]*Region, 0)

	for _, e := range u.Regions {
		allRegions = append(allRegions, e)
	}

	// capture global system list
	allSystems := make([]*SolarSystem, 0)

	for _, r := range allRegions {
		for _, s := range r.Systems {
			allSystems = append(allSystems, s)
		}
	}

	// build transient jumpholes
	edgeCount := physics.RandInRange(MinTransientEdges, MaxTransientEdges)

	for i := 0; i < edgeCount; i++ {
		// pick random A system
		sysAIDX := physics.RandInRange(0, len(allSystems))
		sysA := allSystems[sysAIDX]

		// pick random B system
		sysBIDX := physics.RandInRange(0, len(allSystems))
		sysB := allSystems[sysBIDX]

		// make sure this isn't the same system
		if sysA.ID == sysB.ID {
			continue
		}

		// create transient jumpholes
		nidA := uuid.New()

		jhA := Jumphole{
			ID:           nidA,
			SystemID:     sysA.ID,
			OutSystemID:  sysB.ID,
			JumpholeName: fmt.Sprintf("⚠ %v Jumphole", sysB.SystemName),
			PosX:         float64(physics.RandInRange(-2500000, 2500000)),
			PosY:         float64(physics.RandInRange(-2500000, 2500000)),
			Texture:      "Jumphole-Transient",
			Radius:       float64(physics.RandInRange(50, 400)),
			Mass:         float64(physics.RandInRange(1000, 10000)),
			Theta:        float64(physics.RandInRange(0, 360)),
			Transient:    true,
			Lock:         sync.Mutex{},
		}

		nidB := uuid.New()

		jhB := Jumphole{
			ID:           nidB,
			SystemID:     sysB.ID,
			OutSystemID:  sysA.ID,
			JumpholeName: fmt.Sprintf("⚠ %v Jumphole", sysA.SystemName),
			PosX:         float64(physics.RandInRange(-2500000, 2500000)),
			PosY:         float64(physics.RandInRange(-2500000, 2500000)),
			Texture:      "Jumphole-Transient",
			Radius:       float64(physics.RandInRange(50, 400)),
			Mass:         float64(physics.RandInRange(1000, 10000)),
			Theta:        float64(physics.RandInRange(0, 360)),
			Transient:    true,
			Lock:         sync.Mutex{},
		}

		// link jumpholes
		jhA.OutJumphole = &jhB
		jhA.OutSystem = sysB

		jhB.OutJumphole = &jhA
		jhB.OutSystem = sysA

		// inject into universe
		sysA.jumpholes[jhA.ID.String()] = &jhA
		sysB.jumpholes[jhB.ID.String()] = &jhB
	}
}

// Finds the ship with the specified ID in the running game simulation
func (u *Universe) FindShip(shipID uuid.UUID, noLockSystemID *uuid.UUID) *Ship {
	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		// do not lock if system is exempted
		var lock = true

		if noLockSystemID != nil {
			lock = s.ID != *noLockSystemID
		}

		// get raw pointers to ships in system
		ships := s.MirrorShipMap(lock)

		// look for ship in system
		sh := ships[shipID.String()]

		if sh != nil {
			return sh
		}
	}

	return nil
}

// Finds the ship with the specified ID in the running game simulation
func (u *Universe) FindShipsByUserID(userID uuid.UUID, noLockSystemID *uuid.UUID) []*Ship {
	o := make([]*Ship, 0)

	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		// do lock not if system is exempted
		var lock = true

		if noLockSystemID != nil {
			lock = s.ID != *noLockSystemID
		}

		// get raw pointers to ships in system
		ships := s.MirrorShipMap(lock)

		// look for ships in system owned by this user
		for _, u := range ships {
			if u.UserID == userID {
				// store reference
				o = append(o, u)
			}
		}
	}

	return o
}

// Finds the ship currently being flown by the specified player in the running game simulation
func (u *Universe) FindCurrentPlayerShip(userID uuid.UUID, noLockSystemID *uuid.UUID) *Ship {
	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		// do not lock if system is exempted
		var lock = true

		if noLockSystemID != nil {
			lock = s.ID != *noLockSystemID
		}

		// get raw pointers to ships in system
		ships := s.MirrorShipMap(lock)

		// look for ship in system
		for _, shx := range ships {
			if shx != nil {
				// id and flown check
				if shx.BeingFlownByPlayer && shx.UserID == userID {
					return shx
				}
			}
		}
	}

	return nil
}

// Finds the client of a currently connected player
func (u *Universe) FindCurrentPlayerClient(userID uuid.UUID, noLockSystemID *uuid.UUID) *shared.GameClient {
	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		// do not lock if system is exempted
		var lock = true

		if noLockSystemID != nil {
			lock = s.ID != *noLockSystemID
		}

		// get raw pointers to clients in system
		clients := s.MirrorClientMap(lock)

		// look for client in system
		for _, cx := range clients {
			if cx != nil {
				// id check
				cID := *cx.UID

				if cID == userID {
					// match
					return cx
				}
			}
		}
	}

	return nil
}

// Finds the station with the specified ID in the running game simulation
func (u *Universe) FindStation(stationID uuid.UUID, noLockSystemID *uuid.UUID) *Station {
	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		// do not if system is exempted
		var lock = true

		if noLockSystemID != nil {
			lock = s.ID != *noLockSystemID
		}

		// get raw pointers to ships in system
		stations := s.MirrorStationMap(lock)

		// look for station in system
		sh := stations[stationID.String()]

		if sh != nil {
			return sh
		}
	}

	return nil
}

// writes a global update to all clients on separate goroutines (per system)
func (u *Universe) SendGlobalMessage(msg *models.GameMessage) {
	// iterate over all systems in all regions
	for _, s := range u.AllSystems {
		go func(s *SolarSystem) {
			// obtain lock
			s.Lock.Lock()
			defer s.Lock.Unlock()

			// send messages to clients
			for _, c := range s.clients {
				c.WriteMessage(msg)
			}
		}(s)
	}
}
