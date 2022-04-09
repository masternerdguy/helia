package shared

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Structure representing a game client connected to the server
type GameClient struct {
	SID    *uuid.UUID
	UID    *uuid.UUID
	Conn   *websocket.Conn
	Lock   LabeledMutex
	Joined bool

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
	// label mutex
	c.Lock.Structure = "GameClient"
	c.Lock.UID = fmt.Sprintf("%v :: %v :: %v", c.UID, time.Now(), rand.Float64())

	// obtain lock
	c.Lock.Lock("gameclient.Initialize")
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
	// dispatch to a separate goroutine
	go func(c *GameClient) {
		// return if connection closed
		if c.Dead {
			return
		}

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

			// obfuscate (must use same key as in socketlistener.go)
			utc := time.Now().UTC()
			o = obfuscate(o, fmt.Sprintf("%v^%v|%v*%v", utc.Minute(), utc.Hour(), utc.Day(), utc.Year()))

			// obtain lock (doing so way down here as an optimization)
			c.Lock.Lock("gameclient.WriteMessage")
			defer c.Lock.Unlock()

			// send message
			c.Conn.WriteMessage(1, []byte(o))
		} else {
			// dump error message to console
			TeeLog(err.Error())
		}
	}(c)
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

// Send an informational string to the client to be displayed
func (c *GameClient) WriteInfoMessage(msg string) {
	// get message registry
	msgRegistry := models.NewMessageRegistry()

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
	c.Lock.Lock("gameclient.PushShipEvent")
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
	c.Lock.Lock("gameclient.PopShipEvent")
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
	c.Lock.Lock("gameclient.GetPropertyCache")
	defer c.Lock.Unlock()

	return c.propertyCache
}

func (c *GameClient) SetPropertyCache(x PropertyCache) {
	c.Lock.Lock("gameclient.SetPropertyCache")
	defer c.Lock.Unlock()

	c.propertyCache = x
}

// Sets the value of a client's global update counter
func (c *GameClient) SetLastGlobalAckToken(x int) {
	c.Lock.Lock("gameclient.SetLastGlobalAckToken")
	defer c.Lock.Unlock()

	c.lastGlobalAckToken = x
}

// Returns a client's global update counter
func (c *GameClient) GetLastGlobalAckToken() int {
	c.Lock.Lock("gameclient.GetLastGlobalAckToken")
	defer c.Lock.Unlock()

	return c.lastGlobalAckToken
}

// Resets a client's global update counter to -1
func (c *GameClient) ClearLastGlobalAckToken() {
	c.Lock.Lock("gameclient.ClearLastGlobalAckToken")
	defer c.Lock.Unlock()

	c.lastGlobalAckToken = -1
}

// performs a XOR on a string to obfuscate / deobfuscate it
func obfuscate(input string, key string) (output string) {
	// this must be the same logic as deobfuscateHelper in gamemessage.go!
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}
