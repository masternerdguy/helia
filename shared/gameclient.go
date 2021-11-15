package shared

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"helia/listener/models"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Structure representing a game client connected to the server
type GameClient struct {
	SID  *uuid.UUID
	UID  *uuid.UUID
	Conn *websocket.Conn
	lock sync.Mutex

	// standing to be kept in sync with flown ship(s)
	ReputationSheet PlayerReputationSheet

	// local event queue for player's current ship
	shipEventQueue *eventQueue

	// keys for quick lookups
	CurrentShipID     uuid.UUID
	CurrentSystemID   uuid.UUID
	StartID           uuid.UUID
	EscrowContainerID uuid.UUID

	// property cache
	propertyCache PropertyCache

	// kill switch
	Dead bool
}

type PropertyCache struct {
	ShipCaches []ShipPropertyCacheEntry
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

// Initializes the internals of a GameClient
func (c *GameClient) Initialize() {
	c.lock.Lock()
	defer c.lock.Unlock()

	// initialize empty event queue
	c.shipEventQueue = &eventQueue{
		Events: make([]Event, 0),
	}
}

// Writes a message to a client
func (c *GameClient) WriteMessage(msg *models.GameMessage) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// package message as json
	json, err := json.Marshal(msg)

	if err == nil {
		// compress message
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)

		if _, err := gz.Write([]byte(json)); err != nil {
			// dump error message to console
			log.Println(err)
			return
		}

		if err := gz.Close(); err != nil {
			// dump error message to console
			log.Println(err)
			return
		}

		// convert to string
		o := base64.RawStdEncoding.EncodeToString(b.Bytes())

		// send message
		c.Conn.WriteMessage(1, []byte(o))
	} else {
		// dump error message to console
		log.Println(err)
	}
}

// Send an error string to the client to be displayed
func (c *GameClient) WriteErrorMessage(msg string) {
	// get message registry
	msgRegistry := models.NewMessageRegistry()

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

// Adds an event to the ship event queue
func (c *GameClient) PushShipEvent(evt interface{}, eventType int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.shipEventQueue.Events = append(c.shipEventQueue.Events, Event{
		Type: eventType,
		Body: evt,
	})
}

// Gets the latest event for the player's current ship
func (c *GameClient) PopShipEvent() *Event {
	c.lock.Lock()
	defer c.lock.Unlock()

	if len(c.shipEventQueue.Events) > 0 {
		// get top element
		x := c.shipEventQueue.Events[0]

		// pop top element
		c.shipEventQueue.Events = c.shipEventQueue.Events[1:]

		// return event
		return &x
	}

	return nil
}

func (c *GameClient) GetPropertyCache() PropertyCache {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.propertyCache
}

func (c *GameClient) SetPropertyCache(x PropertyCache) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.propertyCache = x
}
