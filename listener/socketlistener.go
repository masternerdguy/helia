package listener

import (
	"encoding/json"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"log"
	"net/http"
	"sync"

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
	Engine  *engine.HeliaEngine
	clients []*shared.GameClient
	lock    sync.Mutex
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

	//add client to pool
	client := shared.GameClient{
		SID:  nil,
		UID:  nil,
		Conn: c,
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

func (l *SocketListener) handleClientJoin(client *shared.GameClient, body *models.ClientJoinBody) {
	//debug out
	log.Println(fmt.Sprintf("join attempt: %v", &body.SessionID))

	//initialize services
	sessionSvc := sql.GetSessionService()
	shipSvc := sql.GetShipService()
	userSvc := sql.GetUserService()
	msgRegistry := models.NewMessageRegistry()

	//store sid on server
	client.SID = &body.SessionID

	//prepare welcome message to client
	w := models.ServerJoinBody{}

	//lookup user session
	session, err := sessionSvc.GetSessionByID(body.SessionID)

	if err == nil {
		//store userid
		client.UID = &session.UserID
		w.UserID = session.UserID

		//lookup user
		u, _ := userSvc.GetUserByID(session.UserID)

		//lookup current ship
		currShip, _ := shipSvc.GetShipByID(*u.CurrentShipID)

		shipInfo := models.CurrentShipInfo{
			ID:       currShip.ID,
			UserID:   currShip.UserID,
			Created:  currShip.Created,
			ShipName: currShip.ShipName,
			PosX:     currShip.PosX,
			PosY:     currShip.PosY,
			SystemID: currShip.SystemID,
		}

		w.CurrentShipInfo = shipInfo

		//place ship and client into system
		for _, r := range l.Engine.Universe.Regions {
			for _, s := range r.Systems {
				if s.ID == currShip.SystemID {
					s.AddClient(client)

					es := universe.Ship{
						ID:       currShip.ID,
						UserID:   currShip.UserID,
						Created:  currShip.Created,
						ShipName: currShip.ShipName,
						PosX:     currShip.PosX,
						PosY:     currShip.PosY,
						SystemID: currShip.SystemID,
					}

					s.AddShip(&es)
					goto exitLoop
				}
			}
		}

	exitLoop:

		//package message
		b, _ := json.Marshal(&w)

		msg := models.GameMessage{
			MessageType: msgRegistry.Join,
			MessageBody: string(b),
		}

		//send welcome message to client
		client.WriteMessage(&msg)

		//debug out
		log.Println(fmt.Sprintf("joined: %v", &body.SessionID))
	} else {
		//dump error to console
		log.Println(fmt.Sprintf("join error: %v", err))
	}
}

//addClient Adds a client to the server
func (l *SocketListener) addClient(c *shared.GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.clients = append(l.clients, c)
}

//removeClient Removes a client from the server
func (l *SocketListener) removeClient(c *shared.GameClient) {
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
