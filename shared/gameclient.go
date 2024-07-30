package shared

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"helia/listener/models"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Structure representing a game client connected to the server
type GameClient struct {
	SID       *uuid.UUID
	UID       *uuid.UUID
	Conn      *websocket.Conn
	Lock      sync.Mutex
	Joined    bool
	HasStatic bool
	IsDev     bool

	// standing to be kept in sync with flown ship(s)
	ReputationSheet PlayerReputationSheet

	// experience to be kept in sync with flown ship(s)
	ExperienceSheet PlayerExperienceSheet

	// local event queue for player's current ship
	shipEventQueue *eventQueue

	// keys for quick lookups
	CurrentShipID     uuid.UUID
	CurrentSystemID   uuid.UUID
	StartID           uuid.UUID
	EscrowContainerID uuid.UUID

	// to throttle incoming chat messages
	LastChatPostedAt time.Time

	// to disconnect if no meaningful user input is being received
	LastMeaningfulActionAt time.Time

	// property cache
	propertyCache PropertyCache

	// kill switch
	Dead bool

	// last global update message token
	lastGlobalAckToken int
}

type PropertyCache struct {
	ShipCaches    []ShipPropertyCacheEntry
	OutpostCaches []OutpostPropertyCacheEntry
}

type ShipPropertyCacheEntry struct {
	Name                string
	Texture             string
	ShipID              uuid.UUID
	SolarSystemID       uuid.UUID
	SolarSystemName     string
	DockedAtStationID   *uuid.UUID
	DockedAtStationName *string
	Wallet              float64
}

type OutpostPropertyCacheEntry struct {
	Name            string
	Texture         string
	OutpostID       uuid.UUID
	SolarSystemID   uuid.UUID
	SolarSystemName string
	Wallet          float64
}

// Initializes the internals of a GameClient
func (c *GameClient) Initialize() {
	// obtain lock
	c.Lock.Lock()
	defer c.Lock.Unlock()

	// initialize empty event queue
	c.shipEventQueue = &eventQueue{
		Events: make([]Event, 0),
	}

	// initialize last message stamp
	c.LastChatPostedAt = time.Now()

	// initialize last interaction stamp
	c.LastMeaningfulActionAt = time.Now()
}

// Writes a message to a client
func (c *GameClient) WriteMessage(msg *models.GameMessage) {
	// obtain lock
	c.Lock.Lock()
	defer c.Lock.Unlock()

	// return if connection closed
	if c.Dead {
		return
	}

	// set a deadline to write the message
	c.Conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 5000))

	// package message as json
	json, err := json.Marshal(msg)

	if err == nil {
		// compress message
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)

		if _, err := gz.Write([]byte(json)); err != nil {
			// dump error message to console
			TeeLog(err.Error())
			return
		}

		if err := gz.Close(); err != nil {
			// dump error message to console
			TeeLog(err.Error())
			return
		}

		// convert to string
		o := base64.RawStdEncoding.EncodeToString(b.Bytes())

		// send message
		err := c.Conn.WriteMessage(1, []byte(o))

		if err != nil {
			// dump error message to console
			TeeLog(err.Error())

			// close connection
			c.Conn.Close()
			c.Dead = true
		}
	} else {
		// dump error message to console
		TeeLog(err.Error())
	}
}

// Send an error string to the client to be displayed
func (c *GameClient) WriteErrorMessage(msg string) {
	// get message registry
	msgRegistry := models.SharedMessageRegistry

	// package message
	d := models.ServerPushErrorMessage{
		Message: msg,
	}

	b, _ := json.Marshal(&d)

	cu := models.GameMessage{
		MessageType: msgRegistry.PushError,
		MessageBody: string(b),
	}

	// write error to client
	c.WriteMessage(&cu)
}

// Send an informational string to the client to be displayed
func (c *GameClient) WriteInfoMessage(msg string) {
	// get message registry
	msgRegistry := models.SharedMessageRegistry

	// package message
	d := models.ServerPushInfoMessage{
		Message: msg,
	}

	b, _ := json.Marshal(&d)

	cu := models.GameMessage{
		MessageType: msgRegistry.PushInfo,
		MessageBody: string(b),
	}

	// write message to client
	c.WriteMessage(&cu)
}

// Adds an event to the ship event queue
func (c *GameClient) PushShipEvent(evt interface{}, eventType int, meaningful bool) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if meaningful {
		c.LastMeaningfulActionAt = time.Now()
	}

	c.shipEventQueue.Events = append(c.shipEventQueue.Events, Event{
		Type: eventType,
		Body: evt,
	})
}

// Gets the latest event for the player's current ship
func (c *GameClient) PopShipEvent() (*Event, time.Time) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if len(c.shipEventQueue.Events) > 0 {
		// get top element
		x := c.shipEventQueue.Events[0]

		// pop top element
		c.shipEventQueue.Events = c.shipEventQueue.Events[1:]

		// return event
		return &x, c.LastMeaningfulActionAt
	}

	return nil, c.LastMeaningfulActionAt
}

func (c *GameClient) GetPropertyCache() PropertyCache {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	return c.propertyCache
}

func (c *GameClient) SetPropertyCache(x PropertyCache) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.propertyCache = x
}

// Sets the value of a client's global update counter
func (c *GameClient) SetLastGlobalAckToken(x int) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.lastGlobalAckToken = x
}

// Returns a client's global update counter
func (c *GameClient) GetLastGlobalAckToken() int {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	return c.lastGlobalAckToken
}

// Resets a client's global update counter to -1
func (c *GameClient) ClearLastGlobalAckToken() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.lastGlobalAckToken = -1
}
