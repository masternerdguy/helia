package universe

import (
	"helia/physics"
	"math/rand"

	"github.com/google/uuid"
)

// Structure representing a wreck
type Wreck struct {
	ID            uuid.UUID
	SystemID      uuid.UUID
	WreckName     string
	PosX          float64
	PosY          float64
	Texture       string
	Radius        float64
	Theta         float64
	DeadShipItems []*Item
	DeadShip      *Ship
}

// Returns a new physics dummy structure representing this wreck
func (s *Wreck) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: Epsilon,
	}
}

// Attempts to salvage from this wreck and returns a pointer to the retrieved item if successful along with the stack volume
func (s *Wreck) TrySalvage(prob float64, vol float64) (*Item, float64) {
	// make sure wreck is available
	if !s.DeadShip.WreckReady {
		return nil, 0.0
	}

	// iterate over items
	for _, i := range s.DeadShipItems {
		// check volume
		v := 0.0

		if i.IsPackaged {
			// get item type volume metadata
			volume, f := i.ItemTypeMeta.GetFloat64("volume")

			if f {
				vol = (volume * float64(i.Quantity))
			}
		} else {
			// get item volume metadata
			volume, f := i.Meta.GetFloat64("volume")

			if f {
				vol = (volume * float64(i.Quantity))
			}
		}

		if v > vol {
			continue
		}

		// try salvage
		roll := rand.Float64()

		if roll <= prob {
			// remove item from list
			l := make([]*Item, 0)

			for _, x := range s.DeadShipItems {
				if x.ID != i.ID {
					l = append(l, x)
				}
			}

			s.DeadShipItems = l

			// return pointer to item
			return i, v
		}
	}

	// unsuccessful
	return nil, 0
}
