package shared

import (
	"encoding/json"
	"helia/listener/models"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//GameClient Structure representing a game client connected to the server
type GameClient struct {
	SID  *uuid.UUID
	UID  *uuid.UUID
	Conn *websocket.Conn
	lock sync.Mutex

	// local event queue for player's current ship
	shipEventQueue *eventQueue

	// keys for quick lookups
	CurrentShipID   uuid.UUID
	CurrentSystemID uuid.UUID
	StartID         uuid.UUID
}

//Initialize Initializes the internals of a GameClient
func (c *GameClient) Initialize() {
	c.lock.Lock()
	defer c.lock.Unlock()

	//initialize empty event queue
	c.shipEventQueue = &eventQueue{
		Events: make([]Event, 0),
	}
}

//WriteMessage Writes a message to a client
func (c *GameClient) WriteMessage(msg *models.GameMessage) {
	c.lock.Lock()
	defer c.lock.Unlock()

	//package message as json
	json, err := json.Marshal(msg)

	if err == nil {
		//send message
		c.Conn.WriteMessage(1, json)
	} else {
		//dump error message to console
		log.Println(err)
	}
}

//PushShipEvent Adds an event to the ship event queue
func (c *GameClient) PushShipEvent(evt interface{}, eventType int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.shipEventQueue.Events = append(c.shipEventQueue.Events, Event{
		Type: eventType,
		Body: evt,
	})
}

//PopShipEvent Gets the latest event for the player's current ship
func (c *GameClient) PopShipEvent() *Event {
	c.lock.Lock()
	defer c.lock.Unlock()

	if len(c.shipEventQueue.Events) > 0 {
		//get top element
		x := c.shipEventQueue.Events[0]

		//pop top element
		c.shipEventQueue.Events = c.shipEventQueue.Events[1:]

		//return event
		return &x
	}

	return nil
}
