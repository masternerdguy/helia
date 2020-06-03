package listener

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		//todo: lock this down when frontend domains are known
		return true
	},
} // use default options

//SocketListener Listener for handling and dispatching incoming websocket messages
type SocketListener struct {
	clients []*GameClient
	lock    sync.Mutex
}

//GameClient Structure representing a game client connected to the server
type GameClient struct {
	sid  *uuid.UUID
	conn *websocket.Conn
	lock sync.Mutex
}

//WriteMessage Writes a message to a client
func (c *GameClient) WriteMessage(author uuid.UUID, msg *models.GameMessage) {
	c.lock.Lock()
	defer c.lock.Unlock()

	//package message as json
	json, err := json.Marshal(msg)

	if err == nil {
		//send message
		c.conn.WriteMessage(0, json)
	} else {
		//dump error message to console
		log.Println(err)
	}
}

//HandleConnect Handles a client joining the server
func (l *SocketListener) HandleConnect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	//upgrade to websocket connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)

	//return if protocol change failed
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	//defer cleanup of connection
	defer c.Close()

	//return if uuid generation failed
	if err != nil {
		log.Print("uuid:", err)
		return
	}

	//add client to pool
	client := GameClient{
		sid:  nil,
		conn: c,
	}

	l.addClient(&client)

	//defer cleanup of client when they disconnect
	defer l.removeClient(&client)

	//get message type registry
	msgRegistry := models.NewMessageRegistry()

	//socket listener loop
	for {
		//read message from client
		_, r, err := c.ReadMessage()

		m := models.GameMessage{}
		json.Unmarshal(r, &m)

		//exit if read failed
		if err != nil {
			log.Println("read:", err)
			break
		}

		//handle message based on type
		if m.MessageType == msgRegistry.Join {
			//decode body as ClientJoinBody
			b := models.ClientJoinBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			//handle message
			l.handleClientJoin(&client, &b)
		}
	}
}

func (l *SocketListener) handleClientJoin(client *GameClient, body *models.ClientJoinBody) {
	//debug out
	log.Println(fmt.Sprintf("client joined %v", &body.SessionID))

	//store sid
	client.sid = &body.SessionID
}

//addClient Adds a client to the server
func (l *SocketListener) addClient(c *GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.clients = append(l.clients, c)
}

//removeClient Removes a client from the server
func (l *SocketListener) removeClient(c *GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	//find the client to remove
	e := -1
	for i, s := range l.clients {
		if s == c {
			e = i
			break
		}
	}

	//remove client
	if e > -1 {
		t := len(l.clients)

		x := l.clients[t-1]
		l.clients[e] = x

		l.clients = l.clients[:t-1]
	}
}
