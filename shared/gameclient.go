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
