package models

import (
	"time"

	"github.com/google/uuid"
)

// Registry of game message types
type MessageRegistry struct {
	PushError              int
	Join                   int
	GlobalUpdate           int
	NavClick               int
	CurrentShipUpdate      int
	Goto                   int
	Orbit                  int
	Dock                   int
	Undock                 int
	ActivateModule         int
	DeactivateModule       int
	ViewCargoBay           int
	CargoBayUpdate         int
	UnfitModule            int
	TrashItem              int
	PackageItem            int
	UnpackageItem          int
	StackItem              int
	SplitItem              int
	FitModule              int
	SellAsOrder            int
	ViewOpenSellOrders     int
	OpenSellOrdersUpdate   int
	BuySellOrder           int
	ViewIndustrialOrders   int
	IndustrialOrdersUpdate int
	BuyFromSilo            int
	SellToSilo             int
	FactionUpdate          int
	ViewStarMap            int
	StarMapUpdate          int
	ConsumeFuel            int
}

// Registry of target types
type TargetTypeRegistry struct {
	Ship     int
	Station  int
	Star     int
	Planet   int
	Jumphole int
	Asteroid int
}

// Returns a MessageRegistry with correct enum values
func NewMessageRegistry() *MessageRegistry {
	return &MessageRegistry{
		PushError:              -1,
		Join:                   0,
		GlobalUpdate:           1,
		NavClick:               2,
		CurrentShipUpdate:      3,
		Goto:                   4,
		Orbit:                  5,
		Dock:                   6,
		Undock:                 7,
		ActivateModule:         8,
		DeactivateModule:       9,
		ViewCargoBay:           10,
		CargoBayUpdate:         11,
		UnfitModule:            12,
		TrashItem:              13,
		PackageItem:            14,
		UnpackageItem:          15,
		StackItem:              16,
		SplitItem:              17,
		FitModule:              18,
		SellAsOrder:            19,
		ViewOpenSellOrders:     20,
		OpenSellOrdersUpdate:   21,
		BuySellOrder:           22,
		ViewIndustrialOrders:   23,
		IndustrialOrdersUpdate: 24,
		BuyFromSilo:            25,
		SellToSilo:             26,
		FactionUpdate:          27,
		ViewStarMap:            28,
		StarMapUpdate:          29,
		ConsumeFuel:            30,
	}
}

// Returns a TargetTypeRegistry with correct enum values
func NewTargetTypeRegistry() *TargetTypeRegistry {
	return &TargetTypeRegistry{
		Ship:     1,
		Station:  2,
		Star:     3,
		Planet:   4,
		Jumphole: 5,
		Asteroid: 6,
	}
}

// Information about the user's current ship
type CurrentShipInfo struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"uid"`
	Created   time.Time `json:"createdAt"`
	ShipName  string    `json:"shipName"`
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
	FactionID uuid.UUID `json:"factionId"`
	// secrets that should not be globally known
	Accel             float64                 `json:"accel"`
	Turn              float64                 `json:"turn"`
	EnergyP           float64                 `json:"energyP"`
	HeatP             float64                 `json:"heatP"`
	FuelP             float64                 `json:"fuelP"`
	FitStatus         ServerFittingStatusBody `json:"fitStatus"`
	DockedAtStationID *uuid.UUID              `json:"dockedAtStationID"`
	CargoP            float64                 `json:"cargoP"`
	Wallet            float64                 `json:"wallet"`
}

// Structure for passing non-secret information about a ship
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
	FactionID uuid.UUID `json:"factionId"`
}

// Structure for passing non-secret information about a star
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

// Structure for passing non-secret information about a planet
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

// Structure for passing non-secret information about an asteroid
type GlobalAsteroidInfo struct {
	ID       uuid.UUID `json:"id"`
	SystemID uuid.UUID `json:"systemId"`
	Name     string    `json:"name"`
	PosX     float64   `json:"x"`
	PosY     float64   `json:"y"`
	Texture  string    `json:"texture"`
	Radius   float64   `json:"radius"`
	Mass     float64   `json:"mass"`
	Theta    float64   `json:"theta"`
}

// Structure for passing non-secret information about a jumphole
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

// Structure for passing non-secret information about an NPC station
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
	FactionID   uuid.UUID `json:"factionId"`
}

// CurrentSystemInfo Information about the user's current location
type CurrentSystemInfo struct {
	ID         uuid.UUID `json:"id"`
	SystemName string    `json:"systemName"`
	FactionID  uuid.UUID `json:"factionId"`
}

// Message container exchanged between client and server
type GameMessage struct {
	MessageType int    `json:"type"`
	MessageBody string `json:"body"`
}

// Body for a server join request from the client
type ClientJoinBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body for the response to a ClientJoinBody request from the client
type ServerJoinBody struct {
	UserID            uuid.UUID         `json:"uid"`
	CurrentShipInfo   CurrentShipInfo   `json:"currentShipInfo"`
	CurrentSystemInfo CurrentSystemInfo `json:"currentSystemInfo"`
}

// Body for periodically updating clients with globally-known (non-secret) system info
type ServerGlobalUpdateBody struct {
	CurrentSystemInfo CurrentSystemInfo            `json:"currentSystemInfo"`
	Ships             []GlobalShipInfo             `json:"ships"`
	Stars             []GlobalStarInfo             `json:"stars"`
	Planets           []GlobalPlanetInfo           `json:"planets"`
	Jumpholes         []GlobalJumpholeInfo         `json:"jumpholes"`
	Stations          []GlobalStationInfo          `json:"stations"`
	Asteroids         []GlobalAsteroidInfo         `json:"asteroids"`
	NewModuleEffects  []GlobalPushModuleEffectBody `json:"newModuleEffects"`
	NewPointEffects   []GlobalPushPointEffectBody  `json:"newPointEffects"`
}

// Body containing a click-in-space move event from the client
type ClientNavClickBody struct {
	SessionID       uuid.UUID `json:"sid"`
	ScreenTheta     float64   `json:"dT"`
	ScreenMagnitude float64   `json:"m"`
}

// Body containing information about the player's current ship including secrets
type ServerCurrentShipUpdate struct {
	CurrentShipInfo CurrentShipInfo `json:"currentShipInfo"`
}

// Body containing a go-to move order
type ClientGotoBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

// Body containing an orbit move order
type ClientOrbitBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

// Body containing a dock move order
type ClientDockBody struct {
	SessionID uuid.UUID `json:"sid"`
	TargetID  uuid.UUID `json:"targetId"`
	Type      int       `json:"type"`
}

// Body containing an undock move order
type ClientUndockBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing information about a ship's current fitting
type ServerFittingStatusBody struct {
	ARack ServerRackStatusBody `json:"aRack"`
	BRack ServerRackStatusBody `json:"bRack"`
	CRack ServerRackStatusBody `json:"cRack"`
}

// Body containing information about a module fitted to a ship
type ServerModuleStatusBody struct {
	ItemID       uuid.UUID `json:"itemID"`
	ItemTypeID   uuid.UUID `json:"itemTypeID"`
	Family       string    `json:"family"`
	FamilyName   string    `json:"familyName"`
	Type         string    `json:"type"`
	IsCycling    bool      `json:"isCycling"`
	WillRepeat   bool      `json:"willRepeat"`
	CyclePercent int       `json:"cyclePercent"`
	Meta         Meta      `json:"meta"`
	// hardpoint details
	HardpointFamily string `json:"hpFamily"`
	HardpointVolume int    `json:"hpVolume"`
}

// Body containing information about a ship's rack
type ServerRackStatusBody struct {
	Modules []ServerModuleStatusBody `json:"modules"`
}

// Body containing an order to activate a module
type ClientActivateModuleBody struct {
	SessionID  uuid.UUID  `json:"sid"`
	Rack       string     `json:"rack"`
	ItemID     uuid.UUID  `json:"itemID"`
	TargetID   *uuid.UUID `json:"targetId"`
	TargetType *int       `json:"targetType"`
}

// Body containing an order to deactivate a module
type ClientDeactivateModuleBody struct {
	SessionID uuid.UUID `json:"sid"`
	Rack      string    `json:"rack"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a module visual effect to be rendered by the client
type GlobalPushModuleEffectBody struct {
	GfxEffect    string     `json:"gfxEffect"`
	ObjStartID   uuid.UUID  `json:"objStartID"`
	ObjStartType int        `json:"objStartType"`
	ObjEndID     *uuid.UUID `json:"objEndID"`
	ObjEndType   *int       `json:"objEndType"`
}

// Body containing a non-module visual effect to be rendered at a point in space by the client
type GlobalPushPointEffectBody struct {
	GfxEffect string  `json:"gfxEffect"`
	PosX      float64 `json:"x"`
	PosY      float64 `json:"y"`
	Radius    float64 `json:"r"`
}

// Body containing a request for the contents of the ship's cargo bay
type ClientViewCargoBayBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Structure representing a view of an item in a container from the server
type ServerItemViewBody struct {
	ID             uuid.UUID `json:"id"`
	ItemTypeID     uuid.UUID `json:"itemTypeID"`
	ItemTypeName   string    `json:"itemTypeName"`
	ItemFamilyID   string    `json:"itemFamilyID"`
	ItemFamilyName string    `json:"itemFamilyName"`
	Quantity       int       `json:"quantity"`
	IsPackaged     bool      `json:"isPackaged"`
	Meta           Meta      `json:"meta"`
	ItemTypeMeta   Meta      `json:"itemTypeMeta"`
}

// Type representing metadata to be sent between the client and server
type Meta map[string]interface{}

// Generic body for returning container views requested by the client (ex: cargo bay)
type ServerContainerViewBody struct {
	ContainerID uuid.UUID            `json:"id"`
	Items       []ServerItemViewBody `json:"items"`
}

// Body containing a request to unfit a module from a rack and move it into the cargo bay
type ClientUnfitModuleBody struct {
	SessionID uuid.UUID `json:"sid"`
	Rack      string    `json:"rack"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a message string to be displayed to the player from the server
type ServerPushErrorMessage struct {
	Message string `json:"message"`
}

// Body containing a request to trash an item in the current ship's cargo hold
type ClientTrashItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to package an item in the current ship's cargo hold
type ClientPackageItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to unpackage an item in the current ship's cargo hold
type ClientUnpackageItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to stack an item in the current ship's cargo hold
type ClientStackItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to split an item stack in the current ship's cargo hold
type ClientSplitItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
	Size      int       `json:"size"`
}

// Body containing a request to fit a module from the cargo bay to its rack
type ClientFitModuleBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to sell an item stack in the current ship's cargo hold on the stations order market
type ClientSellAsOrderBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
	Price     int       `json:"price"`
}

// Body containing a request for open sell orders at the currently docked-at station
type ClientViewOpenSellOrdersBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Structure containing information about a sell order for return to the client
type ServerSellOrderBody struct {
	ID           uuid.UUID          `json:"id"`
	StationID    uuid.UUID          `json:"stationId"`
	ItemID       uuid.UUID          `json:"itemId"`
	SellerUserID uuid.UUID          `json:"sellerId"`
	AskPrice     float64            `json:"ask"`
	Created      time.Time          `json:"createdAt"`
	Bought       *time.Time         `json:"boughtAt"`
	BuyerUserID  *uuid.UUID         `json:"buyerId"`
	Item         ServerItemViewBody `json:"item"`
}

// Body containing a response to a request to view open sell orders at their station from a client
type ServerOpenSellOrdersUpdateBody struct {
	Orders []ServerSellOrderBody `json:"orders"`
}

// Body containing a request to buy an item listed for sale as a sell order at their station
type ClientBuySellOrderBody struct {
	SessionID uuid.UUID `json:"sid"`
	OrderID   uuid.UUID `json:"orderID"`
}

// Body containing information about an industrial process silo at a station
type ServerIndustrialSiloBody struct {
	StationID        string `json:"stationId"`
	StationProcessID string `json:"stationProcessId"`
	ItemTypeID       string `json:"itemTypeID"`
	ItemTypeName     string `json:"itemTypeName"`
	ItemFamilyID     string `json:"itemFamilyID"`
	ItemFamilyName   string `json:"itemFamilyName"`
	Price            int    `json:"price"`
	Available        int    `json:"available"`
	Meta             Meta   `json:"meta"`
	ItemTypeMeta     Meta   `json:"itemTypeMeta"`
	IsSelling        bool   `json:"isSelling"`
}

// Body containing a request from the client for information about the industrial silos at their current station
type ClientViewIndustrialOrdersBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing the current public state of the industrial silos at a station
type ServerIndustrialOrdersUpdateBody struct {
	OutSilos []ServerIndustrialSiloBody `json:"outSilos"`
	InSilos  []ServerIndustrialSiloBody `json:"inSilos"`
}

// Body containing a request to buy an item for sale from a silo at their station
type ClientBuyFromSiloBody struct {
	SessionID  uuid.UUID `json:"sid"`
	SiloID     uuid.UUID `json:"siloId"`
	ItemTypeID uuid.UUID `json:"itemTypeId"`
	Quantity   int       `json:"quantity"`
}

// Body containing a request to sell an item in their cargo bay to a silo at their station
type ClientSellToSiloBody struct {
	SessionID uuid.UUID `json:"sid"`
	SiloID    uuid.UUID `json:"siloId"`
	ItemID    uuid.UUID `json:"itemId"`
	Quantity  int       `json:"quantity"`
}

// Body containing information about a faction for the client
type ServerFactionBody struct {
	ID            uuid.UUID                   `json:"id"`
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	IsNPC         bool                        `json:"isNPC"`
	IsJoinable    bool                        `json:"isJoinable"`
	IsClosed      bool                        `json:"isClosed"`
	CanHoldSov    bool                        `json:"canHoldSov"`
	Relationships []ServerFactionRelationship `json:"relationships"`
	Ticker        string                      `json:"ticker"`
}

// Structure representing a relationship between two factions for the client
type ServerFactionRelationship struct {
	FactionID        uuid.UUID `json:"factionId"`
	AreOpenlyHostile bool      `json:"openlyHostile"`
	StandingValue    float64   `json:"standingValue"`
}

// Body containing an update on some or all of the unviverse's current factions for the client
type ServerFactionUpdateBody struct {
	Factions []ServerFactionBody `json:"factions"`
}

// Body containing a request from the client for the starmap
type ClientViewStarMapBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing the starmap for the client
type ServerStarMapUpdateBody struct {
	CachedMapData string `json:"cachedMapData"`
}

// Body containing a request to consume a fuel pellet and convert it into fuel for the current ship
type ClientConsumeFuelBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemId"`
}
