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
		log.Print("ws protocol update error:", err)
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

	client.Initialize()
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
			log.Println("ws read error:", err)
			break
		}

		//handle message based on type
		if m.MessageType == msgRegistry.Join {
			//decode body as ClientJoinBody
			b := models.ClientJoinBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			//handle message
			l.handleClientJoin(&client, &b)
		} else if m.MessageType == msgRegistry.NavClick {
			//decode body as ClientNavBody
			b := models.ClientNavClickBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			//handle message
			l.handleClientNavClick(&client, &b)
		}
	}
}

func (l *SocketListener) handleClientJoin(client *shared.GameClient, body *models.ClientJoinBody) {
	//safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	//debug out
	log.Println(fmt.Sprintf("player join attempt: %v", &body.SessionID))

	//initialize services
	sessionSvc := sql.GetSessionService()
	userSvc := sql.GetUserService()
	msgRegistry := models.NewMessageRegistry()
	shipSvc := sql.GetShipService()
	shipTmpSvc := sql.GetShipTemplateService()

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

		//lookup user in database
		u, _ := userSvc.GetUserByID(session.UserID)

		//lookup current ship in memory
		currShip := l.Engine.Universe.FindShip(*u.CurrentShipID)

		if currShip == nil {
			// they must have registered today - get their ship from the db
			dbShip, _ := shipSvc.GetShipByID(*u.CurrentShipID)

			if dbShip == nil {
				return
			}

			dbTemp, _ := shipTmpSvc.GetShipTemplateByID(dbShip.ShipTemplateID)

			if dbTemp == nil {
				return
			}

			// build in-memory ship
			currShip = &universe.Ship{
				ID:       dbShip.ID,
				UserID:   dbShip.UserID,
				Created:  dbShip.Created,
				ShipName: dbShip.ShipName,
				PosX:     dbShip.PosX,
				PosY:     dbShip.PosY,
				SystemID: dbShip.SystemID,
				Texture:  dbShip.Texture,
				Theta:    dbShip.Theta,
				VelX:     dbShip.VelX,
				VelY:     dbShip.VelY,
				Shield:   dbShip.Shield,
				Armor:    dbShip.Armor,
				Hull:     dbShip.Hull,
				Fuel:     dbShip.Fuel,
				Heat:     dbShip.Heat,
				Energy:   dbShip.Energy,
				TemplateData: universe.ShipTemplate{
					ID:               dbTemp.ID,
					Created:          dbTemp.Created,
					ShipTemplateName: dbTemp.ShipTemplateName,
					Texture:          dbTemp.Texture,
					Radius:           dbTemp.Radius,
					BaseAccel:        dbTemp.BaseAccel,
					BaseMass:         dbTemp.BaseMass,
					BaseTurn:         dbTemp.BaseTurn,
					BaseShield:       dbTemp.BaseShield,
					BaseShieldRegen:  dbTemp.BaseShieldRegen,
					BaseArmor:        dbTemp.BaseArmor,
					BaseHull:         dbTemp.BaseHull,
					BaseFuel:         dbTemp.BaseFuel,
					BaseHeatCap:      dbTemp.BaseHeatCap,
					BaseHeatSink:     dbTemp.BaseHeatSink,
					BaseEnergy:       dbTemp.BaseEnergy,
					BaseEnergyRegen:  dbTemp.BaseEnergyRegen,
					ShipTypeID:       dbTemp.ShipTypeID,
				},
			}
		}

		//obtain ship lock
		currShip.Lock.Lock()
		defer currShip.Lock.Unlock()

		//build current ship info data for welcome message
		shipInfo := models.CurrentShipInfo{
			//global stuff
			ID:       currShip.ID,
			UserID:   currShip.UserID,
			Created:  currShip.Created,
			ShipName: currShip.ShipName,
			PosX:     currShip.PosX,
			PosY:     currShip.PosY,
			SystemID: currShip.SystemID,
			Texture:  currShip.Texture,
			Theta:    currShip.Theta,
			VelX:     currShip.VelX,
			VelY:     currShip.VelY,
			Accel:    currShip.GetRealAccel(),
			Mass:     currShip.GetRealMass(),
			Radius:   currShip.TemplateData.Radius,
			Turn:     currShip.GetRealTurn(),
			ShieldP:  (currShip.Shield / currShip.GetRealMaxShield()) * 100,
			ArmorP:   (currShip.Armor / currShip.GetRealMaxArmor()) * 100,
			HullP:    (currShip.Hull / currShip.GetRealMaxHull()) * 100,
			//secret stuff
			EnergyP: (currShip.Energy / currShip.GetRealMaxEnergy()) * 100,
			HeatP:   (currShip.Heat / currShip.GetRealMaxHeat()) * 100,
			FuelP:   (currShip.Fuel / currShip.GetRealMaxFuel()) * 100,
		}

		w.CurrentShipInfo = shipInfo

		//place player into system
		for _, r := range l.Engine.Universe.Regions {
			//lookup system in region
			s := r.Systems[currShip.SystemID.String()]

			if s == nil {
				continue
			}

			s.AddClient(client, true)
			s.AddShip(currShip, true)

			//build current system info for welcome message
			w.CurrentSystemInfo = models.CurrentSystemInfo{}
			w.CurrentSystemInfo.ID = s.ID
			w.CurrentSystemInfo.SystemName = s.SystemName

			//stash current ship and system ids for quick reference
			client.CurrentShipID = currShip.ID
			client.CurrentSystemID = currShip.SystemID

			break
		}

		//package message
		b, _ := json.Marshal(&w)

		msg := models.GameMessage{
			MessageType: msgRegistry.Join,
			MessageBody: string(b),
		}

		//send welcome message to client
		client.WriteMessage(&msg)

		//debug out
		log.Println(fmt.Sprintf("player joined: %v", &body.SessionID))
	} else {
		//dump error to console
		log.Println(fmt.Sprintf("player join error: %v", err))
	}
}

func (l *SocketListener) handleClientNavClick(client *shared.GameClient, body *models.ClientNavClickBody) {
	//safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	//verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientNavClick: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		//initialize services
		msgRegistry := models.NewMessageRegistry()

		//push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.NavClick)
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
