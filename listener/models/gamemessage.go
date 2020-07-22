package models

import (
	"time"

	"github.com/google/uuid"
)

//MessageRegistry Registry of game message types
type MessageRegistry struct {
	Join         int
	GlobalUpdate int
	NavClick     int
}

//NewMessageRegistry Returns a MessageRegistry with correct enum values
func NewMessageRegistry() *MessageRegistry {
	return &MessageRegistry{
		Join:         0,
		GlobalUpdate: 1,
		NavClick:     2,
	}
}

//CurrentShipInfo Information about the user's current ship
type CurrentShipInfo struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"uid"`
	Created  time.Time `json:"createdAt"`
	ShipName string    `json:"shipName"`
	PosX     float64   `json:"x"`
	PosY     float64   `json:"y"`
	SystemID uuid.UUID `json:"systemId"`
	Texture  string    `json:"texture"`
	Theta    float64   `json:"theta"`
	VelX     float64   `json:"velX"`
	VelY     float64   `json:"velY"`
	Mass     float64   `json:"mass"`
	Radius   float64   `json:"radius"`
	// secrets that should not be globally known
	Accel float64 `json:"accel"`
	Turn  float64 `json:"turn"`
}

//GlobalShipInfo Structure for passing non-secret information about a ship
type GlobalShipInfo struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"uid"`
	Created  time.Time `json:"createdAt"`
	ShipName string    `json:"shipName"`
	PosX     float64   `json:"x"`
	PosY     float64   `json:"y"`
	SystemID uuid.UUID `json:"systemId"`
	Texture  string    `json:"texture"`
	Theta    float64   `json:"theta"`
	VelX     float64   `json:"velX"`
	VelY     float64   `json:"velY"`
	Mass     float64   `json:"mass"`
	Radius   float64   `json:"radius"`
}

//GlobalStarInfo Structure for passing non-secret information about a star
type GlobalStarInfo struct {
	ID       uuid.UUID `json:"id"`
	SystemID uuid.UUID `json:"systemId"`
	PosX     float64   `json:"x"`
	PosY     float64   `json:"y"`
	Texture  string    `json:"texture"`
	Radius   float64   `json:"radius"`
	Mass     float64   `json:"mass"`
	Theta    float64   `json:"theta"`
}

//GlobalPlanetInfo Structure for passing non-secret information about a planet
type GlobalPlanetInfo struct {
	ID         uuid.UUID `json:"id"`
	SystemID   uuid.UUID `json:"systemId"`
	PlanetName string    `json:"planetName"`
	PosX       float64   `json:"x"`
	PosY       float64   `json:"y"`
	Texture    string    `json:"texture"`
	Radius     float64   `json:"radius"`
	Mass       float64   `json:"mass"`
	Theta      float64   `json:"theta"`
}

//GlobalJumpholeInfo Structure for passing non-secret information about a jumphole
type GlobalJumpholeInfo struct {
	ID           uuid.UUID `json:"id"`
	SystemID     uuid.UUID `json:"systemId"`
	OutSystemID  uuid.UUID `json:"outSystemId"`
	JumpholeName string    `json:"jumpholeName"`
	PosX         float64   `json:"x"`
	PosY         float64   `json:"y"`
	Texture      string    `json:"texture"`
	Radius       float64   `json:"radius"`
	Mass         float64   `json:"mass"`
	Theta        float64   `json:"theta"`
}

//GlobalStationInfo Structure for passing non-secret information about an NPC station
type GlobalStationInfo struct {
	ID          uuid.UUID `json:"id"`
	SystemID    uuid.UUID `json:"systemId"`
	StationName string    `json:"stationName"`
	PosX        float64   `json:"x"`
	PosY        float64   `json:"y"`
	Texture     string    `json:"texture"`
	Radius      float64   `json:"radius"`
	Mass        float64   `json:"mass"`
	Theta       float64   `json:"theta"`
}

//CurrentSystemInfo Information about the user's current location
type CurrentSystemInfo struct {
	ID         uuid.UUID `json:"id"`
	SystemName string    `json:"systemName"`
}

//GameMessage Message container exchanged between client and server
type GameMessage struct {
	MessageType int    `json:"type"`
	MessageBody string `json:"body"`
}

//ClientJoinBody Body for a server join request from the client
type ClientJoinBody struct {
	SessionID uuid.UUID `json:"sid"`
}

//ServerJoinBody Body for the response to a ClientJoinBody request from the client
type ServerJoinBody struct {
	UserID            uuid.UUID         `json:"uid"`
	CurrentShipInfo   CurrentShipInfo   `json:"currentShipInfo"`
	CurrentSystemInfo CurrentSystemInfo `json:"currentSystemInfo"`
}

//ServerGlobalUpdateBody Body for periodically updating clients with globally-known (non-secret) system info
type ServerGlobalUpdateBody struct {
	CurrentSystemInfo CurrentSystemInfo    `json:"currentSystemInfo"`
	Ships             []GlobalShipInfo     `json:"ships"`
	Stars             []GlobalStarInfo     `json:"stars"`
	Planets           []GlobalPlanetInfo   `json:"planets"`
	Jumpholes         []GlobalJumpholeInfo `json:"jumpholes"`
	Stations          []GlobalStationInfo  `json:"stations"`
}

//ClientNavClickBody Body containing a click-in-space move event from the client
type ClientNavClickBody struct {
	SessionID       uuid.UUID `json:"sid"`
	ScreenTheta     float64   `json:"dT"`
	ScreenMagnitude float64   `json:"m"`
}
