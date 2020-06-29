package universe

import "github.com/google/uuid"

//Region Structure representing a region of the game universe
type Region struct {
	ID         uuid.UUID
	RegionName string
	Systems    map[string]*SolarSystem
}
