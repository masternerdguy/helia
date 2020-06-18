package universe

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

//Ship Structure representing a player ship in the game universe
type Ship struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Created  time.Time
	ShipName string
	PosX     float64
	PosY     float64
	SystemID uuid.UUID
}

//PeriodicUpdate Processes the ship for a tick
func (s Ship) PeriodicUpdate() {
	log.Println(fmt.Sprintf("updating %v", s))
}
