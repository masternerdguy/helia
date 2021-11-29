package universe

import (
	"helia/listener/models"
	"helia/physics"
	"math"

	"github.com/google/uuid"
)

// Structure representing an in-flight missile
type Missile struct {
	ID             uuid.UUID
	FiredByID      uuid.UUID
	TargetID       uuid.UUID
	TargetType     int
	PosX           float64
	PosY           float64
	Texture        string
	Radius         float64
	Module         *FittedSlot
	TicksRemaining int
	MaxVelocity    float64
}

// Processes the missile for a tick
func (s *Missile) PeriodicUpdate() {
	if s.TicksRemaining <= 0 {
		return
	}

	// get target type registry
	tgtTypeReg := models.NewTargetTypeRegistry()

	// get target position
	sB := physics.Dummy{}

	if s.TargetType == tgtTypeReg.Ship {
		// get ship
		sh := s.Module.shipMountedOn.CurrentSystem.ships[s.TargetID.String()]

		if sh != nil {
			// store coordinates
			sB = sh.ToPhysicsDummy()
		}
	} else if s.TargetType == tgtTypeReg.Station {
		// get station
		st := s.Module.shipMountedOn.CurrentSystem.stations[s.TargetID.String()]

		if st != nil {
			// store coordinates
			sB = st.ToPhysicsDummy()
		}
	}

	// get distance to target
	mA := s.ToPhysicsDummy()
	d := physics.Distance(mA, sB)

	// get angle to target
	dX := sB.PosX - mA.PosX
	dY := sB.PosY - mA.PosY
	t := math.Atan2(dY, dX)

	// get velocity to apply
	dP := math.Min(d, s.MaxVelocity/(1000/Heartbeat))

	// apply to position
	s.PosX += math.Cos(t) * dP
	s.PosY += math.Sin(t) * dP

	// decrement lifespan
	s.TicksRemaining--
}

// Returns a new physics dummy structure representing this missile
func (s *Missile) ToPhysicsDummy() physics.Dummy {
	return physics.Dummy{
		VelX: 0,
		VelY: 0,
		PosX: s.PosX,
		PosY: s.PosY,
		Mass: 0,
	}
}
