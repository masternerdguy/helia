package models

import (
	"time"

	"github.com/google/uuid"
)

//MessageRegistry Registry of game message types
type MessageRegistry struct {
	PushError         int
	Join              int
	GlobalUpdate      int
	NavClick          int
	CurrentShipUpdate int
	Goto              int
	Orbit             int
	Dock              int
	Undock            int
	ActivateModule    int
	DeactivateModule  int
	ViewCargoBay      int
	CargoBayUpdate    int
	UnfitModule       int
	TrashItem         int
	PackageItem       int
	UnpackageItem     int
}

//TargetTypeRegistry Registry of target types
type TargetTypeRegistry struct {
	Ship     int
	Station  int
	Star     int
	Planet   int
	Jumphole int
}

//NewMessageRegistry Returns a MessageRegistry with correct enum values
func NewMessageRegistry() *MessageRegistry {
	return &MessageRegistry{
		PushError:         -1,
		Join:              0,
		GlobalUpdate:      1,
		NavClick:          2,
		CurrentShipUpdate: 3,
		Goto:              4,
		Orbit:             5,
		Dock:              6,
		Undock:            7,
		ActivateModule:    8,
		DeactivateModule:  9,
		ViewCargoBay:      10,
		CargoBayUpdate:    11,
		UnfitModule:       12,
		TrashItem:         13,
		PackageItem:       14,
		UnpackageItem:     15,
	}
}

//NewTargetTypeRegistry Returns a TargetTypeRegistry with correct enum values
func NewTargetTypeRegistry() *TargetTypeRegistry {
	return &TargetTypeRegistry{
		Ship:     1,
		Station:  2,
		Star:     3,
		Planet:   4,
		Jumphole: 5,
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
	ShieldP  float64   `json:"shieldP"`
	ArmorP   float64   `json:"armorP"`
	HullP    float64   `json:"hullP"`
	// secrets that should not be globally known
	Accel             float64                 `json:"accel"`
	Turn              float64                 `json:"turn"`
	EnergyP           float64                 `json:"energyP"`
	HeatP             float64                 `json:"heatP"`
	FuelP             float64                 `json:"fuelP"`
	FitStatus         ServerFittingStatusBody `json:"fitStatus"`
	DockedAtStationID *uuid.UUID              `json:"dockedAtStationID"`
}

//GlobalShipInfo Structure for passing non-secret information about a ship
type GlobalShipInfo struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"uid"`
	Created   time.Time `json:"createdAt"`
	ShipName  string    `json:"shipName"`
	OwnerName string    `json:"ownerName"`
	PosX      float64   `json:"x"`
	PosY      float64   `json:"y"`
	SystemID  uuid.UUID `json:"systemId"`
	Texture   string    `json:"texture"`
	Theta     float64   `json:"theta"`
	VelX      float64   `json:"velX"`
	VelY      float64   `json:"velY"`
	Mass      float64   `json:"mass"`
	Radius    float64   `json:"radius"`
	ShieldP   float64   `json:"shieldP"`
	ArmorP    float64   `json:"armorP"`
	HullP     float64   `json:"hullP"`
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
	CurrentSystemInfo CurrentSystemInfo            `json:"currentSystemInfo"`
	Ships             []GlobalShipInfo             `json:"ships"`
	Stars             []GlobalStarInfo             `json:"stars"`
	Planets           []GlobalPlanetInfo           `json:"planets"`
	Jumpholes         []GlobalJumpholeInfo         `json:"jumpholes"`
	Stations          []GlobalStationInfo          `json:"stations"`
	NewModuleEffects  []GlobalPushModuleEffectBody `json:"newModuleEffects"`
	NewPointEffects   []GlobalPushPointEffectBody  `json:"newPointEffects"`
}

//ClientNavClickBody Body containing a click-in-space move event from the client
type ClientNavClickBody struct {
	SessionID       uuid.UUID `json:"sid"`
	ScreenTheta     float64   `json:"dT"`
	ScreenMagnitude float64   `json:"m"`
}

//ServerCurrentShipUpdate Body containing information about the player's current ship including secrets
type ServerCurrentShipUpdate struct {
	CurrentShipInfo CurrentShipInfo `json:"currentShipInfo"`
}

//ClientGotoBody Body containing a go-to move order
type ClientGotoBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

//ClientOrbitBody Body containing an orbit move order
type ClientOrbitBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

//ClientDockBody Body containing a dock move order
type ClientDockBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

//ClientUndockBody Body containing an undock move order
type ClientUndockBody struct {
	SessionID uuid.UUID `json:"sid"`
}

//ServerFittingStatusBody Body containing information about a ship's current fitting
type ServerFittingStatusBody struct {
	ARack ServerRackStatusBody `json:"aRack"`
	BRack ServerRackStatusBody `json:"bRack"`
	CRack ServerRackStatusBody `json:"cRack"`
}

//ServerModuleStatusBody Body containing information about a module fitted to a ship
type ServerModuleStatusBody struct {
	ItemID       uuid.UUID `json:"itemID"`
	ItemTypeID   uuid.UUID `json:"itemTypeID"`
	Family       string    `json:"family"`
	Type         string    `json:"type"`
	IsCycling    bool      `json:"isCycling"`
	WillRepeat   bool      `json:"willRepeat"`
	CyclePercent int       `json:"cyclePercent"`
}

//ServerRackStatusBody Body containing information about a ship's rack
type ServerRackStatusBody struct {
	Modules []ServerModuleStatusBody `json:"modules"`
}

//ClientActivateModuleBody Body containing an order to activate a module
type ClientActivateModuleBody struct {
	SessionID  uuid.UUID  `json:"sid"`
	Rack       string     `json:"rack"`
	ItemID     uuid.UUID  `json:"itemID"`
	TargetID   *uuid.UUID `json:"targetId"`
	TargetType *int       `json:"targetType"`
}

//ClientDeactivateModuleBody Body containing an order to deactivate a module
type ClientDeactivateModuleBody struct {
	SessionID uuid.UUID `json:"sid"`
	Rack      string    `json:"rack"`
	ItemID    uuid.UUID `json:"itemID"`
}

//GlobalPushModuleEffectBody Body containing a module visual effect to be rendered by the client
type GlobalPushModuleEffectBody struct {
	GfxEffect    string     `json:"gfxEffect"`
	ObjStartID   uuid.UUID  `json:"objStartID"`
	ObjStartType int        `json:"objStartType"`
	ObjEndID     *uuid.UUID `json:"objEndID"`
	ObjEndType   *int       `json:"objEndType"`
}

//GlobalPushPointEffectBody Body containing a non-module visual effect to be rendered at a point in space by the client
type GlobalPushPointEffectBody struct {
	GfxEffect string  `json:"gfxEffect"`
	PosX      float64 `json:"x"`
	PosY      float64 `json:"y"`
	Radius    float64 `json:"r"`
}

//ClientViewCargoBayBody Body containing a request for the contents of the ship's cargo bay
type ClientViewCargoBayBody struct {
	SessionID uuid.UUID `json:"sid"`
}

//ServerItemViewBody Structure representing a view of an item in a container from the server
type ServerItemViewBody struct {
	ID             uuid.UUID `json:"id"`
	ItemTypeID     uuid.UUID `json:"itemTypeID"`
	ItemTypeName   string    `json:"itemTypeName"`
	ItemFamilyID   string    `json:"itemFamilyID"`
	ItemFamilyName string    `json:"itemFamilyName"`
	Quantity       int       `json:"quantity"`
	IsPackaged     bool      `json:"isPackaged"`
}

//ServerContainerViewBody Generic body for returning container views requested by the client (ex: cargo bay)
type ServerContainerViewBody struct {
	ContainerID uuid.UUID            `json:"id"`
	Items       []ServerItemViewBody `json:"items"`
}

//ClientUnfitModuleBody Body containing a request to unfit a module from a rack and move it into the cargo bay
type ClientUnfitModuleBody struct {
	SessionID uuid.UUID `json:"sid"`
	Rack      string    `json:"rack"`
	ItemID    uuid.UUID `json:"itemID"`
}

//ServerPushErrorMessage Body containing a message string to be displayed to the player from the server
type ServerPushErrorMessage struct {
	Message string `json:"message"`
}

//ClientTrashItemBody Body containing a request to trash an item in the current ship's cargo hold
type ClientTrashItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

//ClientPackageItemBody Body containing a request to package an item in the current ship's cargo hold
type ClientPackageItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

//ClientUnpackageItemBody Body containing a request to unpackage an item in the current ship's cargo hold
type ClientUnpackageItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}
