package models

import (
	"time"

	"github.com/google/uuid"
)

// Shared registries - do not modify at runtime!
var SharedTargetTypeRegistry = NewTargetTypeRegistry()
var SharedMessageRegistry = NewMessageRegistry()

// Registry of game message types
type MessageRegistry struct {
	PushInfo               int
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
	PlayerFactionUpdate    int
	SelfDestruct           int
	ConsumeRepairKit       int
	ViewProperty           int
	PropertyUpdate         int
	Board                  int
	TransferCredits        int
	SellShipAsOrder        int
	TrashShip              int
	RenameShip             int
	PostSystemChatMessage  int
	TransferItem           int
	ViewExperience         int
	ExperienceUpdate       int
	GlobalAck              int
	ViewSchematicRuns      int
	SchematicRunsUpdate    int
	RunSchematic           int
	CreateNewFaction       int
	LeaveFaction           int
	ApplyToFaction         int
	ViewApplications       int
	ApplicationsUpdate     int
	ApproveApplication     int
	RejectApplication      int
	ViewMembers            int
	MembersUpdate          int
	KickMember             int
	UseModKit              int
	ViewActionReportsPage  int
	ActionReportsPage      int
	ViewActionReportDetail int
	ActionReportDetail     int
	GiveItem               int
	ViewDockedUsers        int
	DockedUsersUpdate      int
	ConsumeOutpostKit      int
}

// Registry of target types
type TargetTypeRegistry struct {
	Ship     int
	Station  int
	Star     int
	Planet   int
	Jumphole int
	Asteroid int
	Wreck    int
}

// Returns a MessageRegistry with correct enum values
func NewMessageRegistry() *MessageRegistry {
	return &MessageRegistry{
		PushInfo:               -2,
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
		PlayerFactionUpdate:    31,
		SelfDestruct:           32,
		ConsumeRepairKit:       33,
		ViewProperty:           34,
		PropertyUpdate:         35,
		Board:                  36,
		TransferCredits:        37,
		SellShipAsOrder:        38,
		TrashShip:              39,
		RenameShip:             40,
		PostSystemChatMessage:  41,
		TransferItem:           42,
		ViewExperience:         43,
		ExperienceUpdate:       44,
		GlobalAck:              45,
		ViewSchematicRuns:      46,
		SchematicRunsUpdate:    47,
		RunSchematic:           48,
		CreateNewFaction:       49,
		LeaveFaction:           50,
		ApplyToFaction:         51,
		ViewApplications:       52,
		ApplicationsUpdate:     53,
		ApproveApplication:     54,
		RejectApplication:      55,
		ViewMembers:            56,
		MembersUpdate:          57,
		KickMember:             58,
		UseModKit:              59,
		ViewActionReportsPage:  60,
		ActionReportsPage:      61,
		ViewActionReportDetail: 62,
		ActionReportDetail:     63,
		GiveItem:               64,
		ViewDockedUsers:        65,
		DockedUsersUpdate:      66,
		ConsumeOutpostKit:      67,
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
		Wreck:    7,
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
	// secret caches that should not be globally known
	CachedHeatSink       float64 `json:"cHeatSink"`
	CachedMaxHeat        float64 `json:"cMaxHeat"`
	CachedRealSpaceDrag  float64 `json:"cRealDrag"`
	CachedMaxFuel        float64 `json:"cMaxFuel"`
	CachedMaxEnergy      float64 `json:"cMaxEnergy"`
	CachedMaxShield      float64 `json:"cMaxShield"`
	CachedMaxArmor       float64 `json:"cMaxArmor"`
	CachedMaxHull        float64 `json:"cMaxHull"`
	CachedEnergyRegen    float64 `json:"cEnergyRegen"`
	CachedShieldRegen    float64 `json:"cShieldRegen"`
	CachedCargoBayVolume float64 `json:"cCargoBayVolume"`
	SumCloaking          float64 `json:"sumCloak"`
	SumVeiling           float64 `json:"sumVeil"`
}

// Structure for passing non-secret information about a ship
type GlobalShipInfo struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"uid"`
	Created       time.Time `json:"createdAt"`
	ShipName      string    `json:"shipName"`
	CharacterName string    `json:"ownerName"`
	PosX          float64   `json:"x"`
	PosY          float64   `json:"y"`
	SystemID      uuid.UUID `json:"systemId"`
	Texture       string    `json:"texture"`
	Theta         float64   `json:"theta"`
	VelX          float64   `json:"velX"`
	VelY          float64   `json:"velY"`
	Mass          float64   `json:"mass"`
	Radius        float64   `json:"radius"`
	ShieldP       float64   `json:"shieldP"`
	ArmorP        float64   `json:"armorP"`
	HullP         float64   `json:"hullP"`
	FactionID     uuid.UUID `json:"factionId"`
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

// Structure for passing non-secret information about a wreck
type GlobalWreckInfo struct {
	ID        uuid.UUID `json:"id"`
	SystemID  uuid.UUID `json:"systemId"`
	WreckName string    `json:"wreckName"`
	PosX      float64   `json:"x"`
	PosY      float64   `json:"y"`
	Texture   string    `json:"texture"`
	Radius    float64   `json:"radius"`
	Theta     float64   `json:"theta"`
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

// Structure for passing non-secret information about a player-owned station
type GlobalOutpostInfo struct {
	ID          uuid.UUID `json:"id"`
	SystemID    uuid.UUID `json:"systemId"`
	OutpostName string    `json:"outpostName"`
	PosX        float64   `json:"x"`
	PosY        float64   `json:"y"`
	ShieldP     float64   `json:"shieldP"`
	ArmorP      float64   `json:"armorP"`
	HullP       float64   `json:"hullP"`
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
	Outposts          []GlobalOutpostInfo          `json:"outposts"`
	Asteroids         []GlobalAsteroidInfo         `json:"asteroids"`
	NewModuleEffects  []GlobalPushModuleEffectBody `json:"newModuleEffects"`
	NewPointEffects   []GlobalPushPointEffectBody  `json:"newPointEffects"`
	Missiles          []GlobalMissileBody          `json:"missiles"`
	SystemChat        []ServerSystemChatBody       `json:"systemChat"`
	Wrecks            []GlobalWreckInfo            `json:"wrecks"`
	Token             int                          `json:"token"`
}

// Body containing a new chat message to push to clients in a system
type ServerSystemChatBody struct {
	SenderID   uuid.UUID `json:"senderId"`
	SenderName string    `json:"senderName"`
	Message    string    `json:"message"`
}

// Body containing a message posted to system chat from the client
type ClientPostSystemChatMessageBody struct {
	SessionID uuid.UUID `json:"sid"`
	Message   string    `json:"message"`
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
	HardpointFamily   string     `json:"hpFamily"`
	HardpointVolume   int        `json:"hpVolume"`
	HardpointPosition [2]float64 `json:"hpPos"`
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
	GfxEffect               string     `json:"gfxEffect"`
	ObjStartID              uuid.UUID  `json:"objStartID"`
	ObjStartType            int        `json:"objStartType"`
	ObjStartHardpointOffset [2]float64 `json:"objStartHPOffset"` // [radius, theta]
	ObjEndID                *uuid.UUID `json:"objEndID"`
	ObjEndType              *int       `json:"objEndType"`
}

// Body containing a non-module visual effect to be rendered at a point in space by the client
type GlobalPushPointEffectBody struct {
	GfxEffect string  `json:"gfxEffect"`
	PosX      float64 `json:"x"`
	PosY      float64 `json:"y"`
	Radius    float64 `json:"r"`
}

// Body containing a missile to be rendered by the client
type GlobalMissileBody struct {
	ID      uuid.UUID `json:"id"`
	PosX    float64   `json:"x"`
	PosY    float64   `json:"y"`
	Texture string    `json:"t"`
	Radius  float64   `json:"r"`
}

// Body containing a request for the contents of the ship's cargo bay
type ClientViewCargoBayBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Structure representing a view of an item in a container from the server
type ServerItemViewBody struct {
	ID             uuid.UUID                `json:"id"`
	ItemTypeID     uuid.UUID                `json:"itemTypeID"`
	ItemTypeName   string                   `json:"itemTypeName"`
	ItemFamilyID   string                   `json:"itemFamilyID"`
	ItemFamilyName string                   `json:"itemFamilyName"`
	Quantity       int                      `json:"quantity"`
	IsPackaged     bool                     `json:"isPackaged"`
	Meta           Meta                     `json:"meta"`
	ItemTypeMeta   Meta                     `json:"itemTypeMeta"`
	Schematic      *ServerSchematicViewBody `json:"schematic"`
}

// Structure representing a view of the factors of a schematic
type ServerSchematicViewBody struct {
	ID      uuid.UUID                       `json:"id"`
	Time    int                             `json:"time"`
	Inputs  []ServerSchematicFactorViewBody `json:"inputs"`
	Outputs []ServerSchematicFactorViewBody `json:"outputs"`
}

// Structure representing a view of a factor in a schematic
type ServerSchematicFactorViewBody struct {
	ItemTypeID   uuid.UUID `json:"itemTypeId"`
	ItemTypeName string    `json:"itemTypeName"`
	Quantity     int       `json:"quantity"`
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

// Body containing an error message string to be displayed to the player from the server
type ServerPushErrorMessage struct {
	Message string `json:"message"`
}

// Body containing an informational message string to be displayed to the player from the server
type ServerPushInfoMessage struct {
	Message string `json:"message"`
}

// Body containing a request to trash an item in the current ship's cargo hold
type ClientTrashItemBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemID"`
}

// Body containing a request to transfer an item in the current ship's cargo hold to another ship
type ClientTransferItemBody struct {
	SessionID  uuid.UUID `json:"sid"`
	ItemID     uuid.UUID `json:"itemID"`
	ReceiverID uuid.UUID `json:"receiverID"`
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
	ID               string `json:"id"`
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
	CanHoldSov    bool                        `json:"canHoldSov"`
	Relationships []ServerFactionRelationship `json:"relationships"`
	Ticker        string                      `json:"ticker"`
	OwnerID       string                      `json:"ownerId"`
}

// Structure representing a relationship between two factions for the client
type ServerFactionRelationship struct {
	FactionID        uuid.UUID `json:"factionId"`
	AreOpenlyHostile bool      `json:"openlyHostile"`
	StandingValue    float64   `json:"standingValue"`
}

// Body containing an update on some or all of the universe's current factions for the client
type ServerFactionUpdateBody struct {
	Factions []ServerFactionBody `json:"factions"`
}

// Structure representing a relationship between two factions for the client
type ServerPlayerFactionRelationship struct {
	FactionID        uuid.UUID `json:"factionId"`
	AreOpenlyHostile bool      `json:"openlyHostile"`
	StandingValue    float64   `json:"standingValue"`
	IsMember         bool      `json:"isMember"`
}

// Body containing an update on their relationship to some or all of the universe's current factions for the client
type ServerPlayerFactionUpdateBody struct {
	Factions []ServerPlayerFactionRelationship `json:"factions"`
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

// Body containing a request to self destruct the current ship
type ClientSelfDestructBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing a request to consume a repair kit and convert it into health for the current ship
type ClientConsumeRepairKitBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemId"`
}

// Body containing a request to consume an outpost kit and convert it into a new outpost
type ClientConsumeOutpostKitBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemId"`
}

// Body containing a request from the client for a summary of the player's owned property
type ClientViewPropertyBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing a summary of the property owned by a player
type ServerPropertyUpdateBody struct {
	Ships []ServerShipPropertyCacheEntry `json:"ships"`
}

type ServerShipPropertyCacheEntry struct {
	Name                string     `json:"name"`
	Texture             string     `json:"texture"`
	ShipID              uuid.UUID  `json:"id"`
	SolarSystemID       uuid.UUID  `json:"systemId"`
	SolarSystemName     string     `json:"systemName"`
	DockedAtStationID   *uuid.UUID `json:"dockedAtId"`
	DockedAtStationName *string    `json:"dockedAtName"`
	Wallet              float64    `json:"wallet"`
}

// Body containing a request to board a ship owned by the player in the player's current station
type ClientBoardBody struct {
	SessionID uuid.UUID `json:"sid"`
	ShipID    uuid.UUID `json:"shipId"`
}

// Body containing a request to transfer credits between another ship owned by the player in the player's current station
type ClientTransferCreditsBody struct {
	SessionID uuid.UUID `json:"sid"`
	ShipID    uuid.UUID `json:"shipId"`
	Amount    int       `json:"amount"`
}

// Body containing a request to sell a ship owned by a player docked at their current station on the orders market
type ClientSellShipAsOrderBody struct {
	SessionID uuid.UUID `json:"sid"`
	ShipID    uuid.UUID `json:"shipId"`
	Price     int       `json:"price"`
}

// Body containing a request to trash a ship owned by the player in the player's current station
type ClientTrashShipBody struct {
	SessionID uuid.UUID `json:"sid"`
	ShipID    uuid.UUID `json:"shipId"`
}

// Body containing a request to rename a ship owned by the player in the player's current station
type ClientRenameShipBody struct {
	SessionID uuid.UUID `json:"sid"`
	ShipID    uuid.UUID `json:"shipId"`
	Name      string    `json:"name"`
}

// Body containing a request from the client for a summary of the player's experience levels
type ClientViewExperienceBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing a response to the client containing the player's experience levels
type ServerExperienceUpdateBody struct {
	ShipEntries   []ServerExperienceUpdateShipEntryBody   `json:"ships"`
	ModuleEntries []ServerExperienceUpdateModuleEntryBody `json:"modules"`
}

// Body containing the player's experience level with a given ship template
type ServerExperienceUpdateShipEntryBody struct {
	ExperienceLevel  float64   `json:"experienceLevel"`
	ShipTemplateID   uuid.UUID `json:"shipTemplateID"`
	ShipTemplateName string    `json:"shipTemplateName"`
}

// Body containing the player's experience level with a given module
type ServerExperienceUpdateModuleEntryBody struct {
	ExperienceLevel float64   `json:"experienceLevel"`
	ItemTypeID      uuid.UUID `json:"itemTypeID"`
	ItemTypeName    string    `json:"itemTypeName"`
}

// Body indicating to the server that the last global update was received
type ClientGlobalAckBody struct {
	SessionID     uuid.UUID `json:"sid"`
	SolarSystemID uuid.UUID `json:"sysId"`
	Token         int       `json:"token"`
}

// Body containing a request for schematic runs
type ClientViewSchematicRunsBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing currently running schematics
type ServerSchematicRunsUpdateBody struct {
	Runs []ServerSchematicRunEntryBody `json:"runs"`
}

// Body containing a request to run a schematic in the player's station warehouse/workshop
type ClientRunSchematicBody struct {
	SessionID uuid.UUID `json:"sid"`
	ItemID    uuid.UUID `json:"itemId"`
}

// Body containing a currently running schematic
type ServerSchematicRunEntryBody struct {
	SchematicRunID     uuid.UUID `json:"id"`
	SchematicName      string    `json:"schematicName"`
	HostShipName       string    `json:"hostShipName"`
	HostStationName    string    `json:"hostStationName"`
	SolarSystemName    string    `json:"solarSystemName"`
	StatusID           string    `json:"statusId"`
	PercentageComplete float64   `json:"percentageComplete"`
}

// Body containing a request to create a new player faction
type ClientCreateNewFactionBody struct {
	SessionID   uuid.UUID `json:"sid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ticker      string    `json:"ticker"`
}

// Body containing a request to leave a player faction
type ClientLeaveFactionBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing a request to apply to join a player faction
type ClientApplyToFactionBody struct {
	SessionID uuid.UUID `json:"sid"`
	FactionID uuid.UUID `json:"factionId"`
}

// Body containing a request to view applications to the player's faction
type ClientViewApplicationsBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing current applicants to the player's faction
type ServerApplicationsUpdateBody struct {
	Applications []ServerApplicationEntry `json:"applications"`
}

// Structure representing an application to a player's faction
type ServerApplicationEntry struct {
	UserID        uuid.UUID `json:"id"`
	CharacterName string    `json:"name"`
	FactionName   string    `json:"faction"`
	FactionTicker string    `json:"ticker"`
}

// Body containing a request to approve an application to a player's faction
type ClientApproveApplicationBody struct {
	SessionID uuid.UUID `json:"sid"`
	UserID    uuid.UUID `json:"userId"`
}

// Body containing a request to reject an application to a player's faction
type ClientRejectApplicationBody struct {
	SessionID uuid.UUID `json:"sid"`
	UserID    uuid.UUID `json:"userId"`
}

// Body containing current members of the player's faction
type ServerMembersUpdateBody struct {
	Members []ServerMemberEntry `json:"members"`
}

// Structure representing a member of a player's faction
type ServerMemberEntry struct {
	UserID        uuid.UUID `json:"id"`
	CharacterName string    `json:"name"`
}

// Body containing a request to view members of the player's faction
type ClientViewMembersBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Body containing a request to kick a member from a player's faction
type ClientKickMemberBody struct {
	SessionID uuid.UUID `json:"sid"`
	UserID    uuid.UUID `json:"userId"`
}

// Body containing a request to apply a mod kit to a module in the player's cargo bay
type ClientUseModKitBody struct {
	SessionID    uuid.UUID `json:"sid"`
	ModuleItemID uuid.UUID `json:"moduleId"`
	ModKitItemID uuid.UUID `json:"kitId"`
}

// Body containing a request to view a page of action reports
type ClientViewActionReportsPageBody struct {
	SessionID uuid.UUID `json:"sid"`
	Page      int       `json:"page"`
	Count     int       `json:"count"`
}

// Body containing a low-overhead index of available action reports on a page
type ServerActionReportsPage struct {
	Page int                         `json:"page"`
	Logs []ServerActionReportSummary `json:"logs"`
}

// Body containing a request to view a full action report
type ClientViewActionReportDetailBody struct {
	SessionID uuid.UUID `json:"sid"`
	KillID    uuid.UUID `json:"killId"`
}

// Body containing a full action report requested by the client
type ServerActionReportDetail struct {
	ID            uuid.UUID     `json:"id"`
	ServerKillLog ServerKillLog `json:"log"`
}

// Structure containing a low-overhead action report row
type ServerActionReportSummary struct {
	ID                     uuid.UUID `json:"id"`
	VictimName             string    `json:"victim"`
	VictimShipTemplateName string    `json:"ship"`
	VictimTicker           *string   `json:"ticker"`
	Timestamp              time.Time `json:"timestamp"`
	SystemName             string    `json:"systemName"`
	RegionName             string    `json:"regionName"`
	Parties                int       `json:"parties"`
}

// Structure represening a copy-pastable report of the death of a ship
type ServerKillLog struct {
	Header          ServerKillLogHeader          `json:"header"`
	Fitting         ServerKillLogFitting         `json:"fitting"`
	Cargo           []ServerKillLogCargoItem     `json:"cargo"`
	InvolvedParties []ServerKillLogInvolvedParty `json:"involvedParties"`
	Wallet          int64                        `json:"wallet"`
}

// Structure representing a summary header for a kill log
type ServerKillLogHeader struct {
	VictimID               uuid.UUID `json:"victimID"`
	VictimName             string    `json:"victimName"`
	VictimFactionID        uuid.UUID `json:"victimFactionID"`
	VictimFactionName      string    `json:"victimFactionName"`
	VictimShipTemplateID   uuid.UUID `json:"victimShipTemplateID"`
	VictimShipTemplateName string    `json:"victimShipTemplateName"`
	VictimShipID           uuid.UUID `json:"victimShipID"`
	VictimShipName         string    `json:"victimShipName"`
	Timestamp              time.Time `json:"timestamp"`
	SolarSystemID          uuid.UUID `json:"solarSystemID"`
	SolarSystemName        string    `json:"solarSystemName"`
	RegionID               uuid.UUID `json:"regionID"`
	RegionName             string    `json:"regionName"`
	HoldingFactionID       uuid.UUID `json:"holdingFactionID"`
	HoldingFactionName     string    `json:"holdingFactionName"`
	InvolvedParties        int       `json:"involvedParties"`
	IsNPC                  bool      `json:"isNPC"`
	PosX                   float64   `json:"posX"`
	PosY                   float64   `json:"posY"`
}

// Structure representing any combat between two ships
type ServerKillLogInvolvedParty struct {
	// aggressor info
	UserID        uuid.UUID `json:"userID"`
	FactionID     uuid.UUID `json:"factionID"`
	CharacterName string    `json:"characterName"`
	FactionName   string    `json:"factionName"`
	IsNPC         bool      `json:"isNPC"`
	LastAggressed time.Time `json:"lastAggressed"`
	// aggressor ship info
	ShipID           uuid.UUID `json:"shipID"`
	ShipName         string    `json:"shipName"`
	ShipTemplateID   uuid.UUID `json:"shipTemplateID"`
	ShipTemplateName string    `json:"shipTemplateName"`
	// location info
	LastSolarSystemID   uuid.UUID `json:"lastSolarSystemID"`
	LastSolarSystemName string    `json:"lastSolarSystemName"`
	LastRegionID        uuid.UUID `json:"lastRegionID"`
	LastRegionName      string    `json:"lastRegionName"`
	LastPosX            float64   `json:"lastPosX"`
	LastPosY            float64   `json:"lastPosY"`
	// weapons used against victim
	WeaponUse map[string]*ServerKillLogWeaponUse `json:"weaponUse"`
}

// Structure representing a weapon used in combat between two ships
type ServerKillLogWeaponUse struct {
	ItemID          uuid.UUID `json:"itemID"`
	ItemTypeID      uuid.UUID `json:"itemTypeID"`
	ItemFamilyID    string    `json:"itemFamilyID"`
	ItemFamilyName  string    `json:"itemFamilyName"`
	ItemTypeName    string    `json:"itemTypeName"`
	LastUsed        time.Time `json:"lastUsed"`
	DamageInflicted float64   `json:"damageInflicted"`
}

// Structure representing a ship fitting for a kill log
type ServerKillLogFitting struct {
	ARack []ServerKillLogSlot `json:"rackA"`
	BRack []ServerKillLogSlot `json:"rackB"`
	CRack []ServerKillLogSlot `json:"rackC"`
}

// Structure representing a slot in a kill log fitting
type ServerKillLogSlot struct {
	ItemID              uuid.UUID `json:"itemID"`
	ItemTypeID          uuid.UUID `json:"itemTypeID"`
	ItemFamilyID        string    `json:"itemFamilyID"`
	ItemTypeName        string    `json:"itemTypeName"`
	ItemFamilyName      string    `json:"itemFamilyName"`
	IsModified          bool      `json:"isModified"`
	CustomizationFactor int       `json:"customizationFactor"`
}

// Structure representing an item in the dead ship's cargo bay
type ServerKillLogCargoItem struct {
	ItemID         uuid.UUID `json:"itemID"`
	ItemTypeID     uuid.UUID `json:"itemTypeID"`
	ItemFamilyID   string    `json:"itemFamilyID"`
	ItemTypeName   string    `json:"ItemTypeName"`
	ItemFamilyName string    `json:"itemFamilyName"`
	Quantity       int       `json:"quantity"`
	IsPackaged     bool      `json:"isPackaged"`
}

// Body containing a request to give an item in the current ship's cargo hold to another player
type ClientGiveItemBody struct {
	SessionID  uuid.UUID `json:"sid"`
	ItemID     uuid.UUID `json:"itemID"`
	ReceiverID uuid.UUID `json:"receiverID"`
}

// Body containing a request for a list of players docked at the same station
type ClientViewDockedUsersBody struct {
	SessionID uuid.UUID `json:"sid"`
}

// Response to a request for a list of players docked at the same station
type ServerDockedUsersUpdateBody struct {
	Users map[string]uuid.UUID `json:"users"`
}
