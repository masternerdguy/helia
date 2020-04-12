package universe

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

//SolarSystem Structure representing a solar system
type SolarSystem struct {
	ID         uuid.UUID
	SystemName string
	RegionID   uuid.UUID
}

//PeriodicUpdate Processes the solar system for a current tick
func (s SolarSystem) PeriodicUpdate() {
	log.Println(fmt.Sprintf("updating %v", s))
}
