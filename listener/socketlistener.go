package listener

import (
	"encoding/json"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo: lock this down when frontend domains are known
		return true
	},
} // use default options

// Listener for handling and dispatching incoming websocket messages
type SocketListener struct {
	Engine  *engine.HeliaEngine
	clients []*shared.GameClient
	lock    sync.Mutex
}

// Handles a client joining the server
func (l *SocketListener) HandleConnect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	// upgrade to websocket connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)

	// return if protocol change failed
	if err != nil {
		log.Print("ws protocol update error:", err)
		return
	}

	// defer cleanup of connection
	defer c.Close()

	// add client to pool
	client := shared.GameClient{
		SID:  nil,
		UID:  nil,
		Conn: c,
	}

	client.Initialize()
	l.addClient(&client)

	// defer cleanup of client when they disconnect
	defer l.removeClient(&client)

	// get message type registry
	msgRegistry := models.NewMessageRegistry()

	// socket listener loop
	for {
		// read message from client
		_, r, err := c.ReadMessage()

		m := models.GameMessage{}
		json.Unmarshal(r, &m)

		// exit if read failed
		if err != nil {
			log.Println("ws read error:", err)
			break
		}

		// handle message based on type
		if m.MessageType == msgRegistry.Join {
			// decode body as ClientJoinBody
			b := models.ClientJoinBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientJoin(&client, &b)
		} else if m.MessageType == msgRegistry.NavClick {
			// decode body as ClientNavBody
			b := models.ClientNavClickBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientNavClick(&client, &b)
		} else if m.MessageType == msgRegistry.Goto {
			// decode body as ClientGotoBody
			b := models.ClientGotoBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientGoto(&client, &b)
		} else if m.MessageType == msgRegistry.Orbit {
			// decode body as ClientOrbitBody
			b := models.ClientOrbitBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientOrbit(&client, &b)
		} else if m.MessageType == msgRegistry.Dock {
			// decode body as ClientDockBody
			b := models.ClientDockBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientDock(&client, &b)
		} else if m.MessageType == msgRegistry.Undock {
			// decode body as ClientUndockBody
			b := models.ClientUndockBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientUndock(&client, &b)
		} else if m.MessageType == msgRegistry.ActivateModule {
			// decode body as ClientActivateModuleBody
			b := models.ClientActivateModuleBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientActivateModule(&client, &b)
		} else if m.MessageType == msgRegistry.DeactivateModule {
			// decode body as ClientDeactivateModuleBody
			b := models.ClientDeactivateModuleBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientDeactivateModule(&client, &b)
		} else if m.MessageType == msgRegistry.ViewCargoBay {
			// decode body as ClientViewCargoBayBody
			b := models.ClientViewCargoBayBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewCargoBay(&client, &b)
		} else if m.MessageType == msgRegistry.UnfitModule {
			// decode body as ClientUnfitModuleBody
			b := models.ClientUnfitModuleBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientUnfitModule(&client, &b)
		} else if m.MessageType == msgRegistry.TrashItem {
			// decode body as ClientTrashItemBody
			b := models.ClientTrashItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientTrashItem(&client, &b)
		} else if m.MessageType == msgRegistry.PackageItem {
			// decode body as ClientPackageItemBody
			b := models.ClientPackageItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientPackageItem(&client, &b)
		} else if m.MessageType == msgRegistry.UnpackageItem {
			// decode body as ClientUnpackageItemBody
			b := models.ClientUnpackageItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientUnpackageItem(&client, &b)
		} else if m.MessageType == msgRegistry.StackItem {
			// decode body as ClientStackItemBody
			b := models.ClientStackItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientStackItem(&client, &b)
		} else if m.MessageType == msgRegistry.SplitItem {
			// decode body as ClientSplitItemBody
			b := models.ClientSplitItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientSplitItem(&client, &b)
		} else if m.MessageType == msgRegistry.FitModule {
			// decode body as ClientFitModuleBody
			b := models.ClientFitModuleBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientFitModule(&client, &b)
		} else if m.MessageType == msgRegistry.SellAsOrder {
			// decode body as ClientSellAsOrderBody
			b := models.ClientSellAsOrderBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientSellAsOrder(&client, &b)
		} else if m.MessageType == msgRegistry.ViewOpenSellOrders {
			// decode body as ClientViewOpenSellOrdersBody
			b := models.ClientViewOpenSellOrdersBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewOpenSellOrders(&client, &b)
		} else if m.MessageType == msgRegistry.BuySellOrder {
			// decode body as ClientBuySellOrderBody
			b := models.ClientBuySellOrderBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientBuySellOrder(&client, &b)
		} else if m.MessageType == msgRegistry.ViewIndustrialOrders {
			// decode body as ClientViewIndustrialOrdersBody
			b := models.ClientViewIndustrialOrdersBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewIndustrialOrders(&client, &b)
		} else if m.MessageType == msgRegistry.BuyFromSilo {
			// decode body as ClientBuyFromSiloBody
			b := models.ClientBuyFromSiloBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientBuyFromSilo(&client, &b)
		} else if m.MessageType == msgRegistry.SellToSilo {
			// decode body as ClientSellToSiloBody
			b := models.ClientSellToSiloBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientSellToSilo(&client, &b)
		} else if m.MessageType == msgRegistry.ViewStarMap {
			// decode body as ClientViewStarMapBody
			b := models.ClientViewStarMapBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewStarMap(&client, &b)
		} else if m.MessageType == msgRegistry.ConsumeFuel {
			// decode body as ClientConsumeFuelBody
			b := models.ClientConsumeFuelBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientConsumeFuel(&client, &b)
		}
	}
}

func (l *SocketListener) handleClientJoin(client *shared.GameClient, body *models.ClientJoinBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// debug out
	log.Println(fmt.Sprintf("player join attempt: %v", &body.SessionID))

	// initialize services
	sessionSvc := sql.GetSessionService()
	userSvc := sql.GetUserService()
	msgRegistry := models.NewMessageRegistry()
	shipSvc := sql.GetShipService()
	startSvc := sql.GetStartService()

	// store sid on server
	client.SID = &body.SessionID

	// prepare welcome message to client
	w := models.ServerJoinBody{}

	// lookup user session
	session, err := sessionSvc.GetSessionByID(body.SessionID)

	if err == nil {
		// store userid
		client.UID = &session.UserID
		w.UserID = session.UserID

		// lookup user in database
		u, _ := userSvc.GetUserByID(session.UserID)

		// store start
		client.StartID = u.StartID

		// store escrow container
		client.EscrowContainerID = u.EscrowContainerID

		// lookup current ship in memory
		currShip := l.Engine.Universe.FindShip(*u.CurrentShipID)

		if currShip == nil {
			// they must have registered today - get their ship from the db
			dbShip, _ := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if dbShip == nil {
				// report error
				log.Println(fmt.Sprintf("player join error: unable to find ship %v for %v", u.CurrentShipID, u.ID))

				// try to recover by creating a noob ship for player
				nStart, err := startSvc.GetStartByID(u.StartID)

				if err != nil {
					// dump error to console
					log.Println(fmt.Sprintf("player join recovery - unable to get start for player: %v", err))
					return
				}

				u, err = engine.CreateNoobShipForPlayer(nStart, u.ID, false)

				if err != nil {
					// dump error to console
					log.Println(fmt.Sprintf("player join recovery - unable to create noob ship for player: %v", err))
					return
				}

				// get their new noob ship from the db
				dbShip, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

				if dbShip == nil || err != nil {
					log.Println(fmt.Sprintf("player join recovery: unable to find recovery noob ship %v for %v", u.CurrentShipID, u.ID))
					return
				}
			}

			// build in-memory ship
			currShip, err = engine.LoadShip(dbShip)

			if dbShip == nil {
				log.Println(fmt.Sprintf("player join error: %v", err))
				return
			}
		}

		// obtain ship lock
		currShip.Lock.Lock()
		defer currShip.Lock.Unlock()

		// build current ship info data for welcome message
		shipInfo := models.CurrentShipInfo{
			// global stuff
			ID:        currShip.ID,
			UserID:    currShip.UserID,
			Created:   currShip.Created,
			ShipName:  currShip.ShipName,
			PosX:      currShip.PosX,
			PosY:      currShip.PosY,
			SystemID:  currShip.SystemID,
			Texture:   currShip.Texture,
			Theta:     currShip.Theta,
			VelX:      currShip.VelX,
			VelY:      currShip.VelY,
			Accel:     currShip.GetRealAccel(),
			Mass:      currShip.GetRealMass(),
			Radius:    currShip.TemplateData.Radius,
			Turn:      currShip.GetRealTurn(),
			ShieldP:   (currShip.Shield / currShip.GetRealMaxShield()) * 100,
			ArmorP:    (currShip.Armor / currShip.GetRealMaxArmor()) * 100,
			HullP:     (currShip.Hull / currShip.GetRealMaxHull()) * 100,
			FactionID: u.CurrentFactionID,
			// secret stuff
			EnergyP: (currShip.Energy / currShip.GetRealMaxEnergy()) * 100,
			HeatP:   (currShip.Heat / currShip.GetRealMaxHeat()) * 100,
			FuelP:   (currShip.Fuel / currShip.GetRealMaxFuel()) * 100,
		}

		w.CurrentShipInfo = shipInfo

		// place player into system
		for _, r := range l.Engine.Universe.Regions {
			// lookup system in region
			s := r.Systems[currShip.SystemID.String()]

			if s == nil {
				continue
			}

			s.AddClient(client, true)
			s.AddShip(currShip, true)

			// build current system info for welcome message
			w.CurrentSystemInfo = models.CurrentSystemInfo{}
			w.CurrentSystemInfo.ID = s.ID
			w.CurrentSystemInfo.SystemName = s.SystemName

			// stash current ship and system ids for quick reference
			client.CurrentShipID = currShip.ID
			client.CurrentSystemID = currShip.SystemID

			break
		}

		// package message
		b, _ := json.Marshal(&w)

		msg := models.GameMessage{
			MessageType: msgRegistry.Join,
			MessageBody: string(b),
		}

		// send welcome message to client
		client.WriteMessage(&msg)

		// prepare initial faction info for client
		af := models.ServerFactionUpdateBody{
			Factions: make([]models.ServerFactionBody, 0),
		}

		factions := l.Engine.Universe.Factions

		for _, f := range factions {
			af.Factions = append(af.Factions, models.ServerFactionBody{
				ID:          f.ID,
				Name:        f.Name,
				Description: f.Description,
				IsNPC:       f.IsNPC,
				IsJoinable:  f.IsJoinable,
				IsClosed:    f.IsClosed,
				CanHoldSov:  f.CanHoldSov,
				Ticker:      f.Ticker,
			})
		}

		// package faction data
		b, _ = json.Marshal(&af)

		msg = models.GameMessage{
			MessageType: msgRegistry.FactionUpdate,
			MessageBody: string(b),
		}

		// send all faction update message to client
		client.WriteMessage(&msg)

		// debug out
		log.Println(fmt.Sprintf("player joined: %v", &body.SessionID))
	} else {
		// dump error to console
		log.Println(fmt.Sprintf("player join error: %v", err))
	}
}

func (l *SocketListener) handleClientNavClick(client *shared.GameClient, body *models.ClientNavClickBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientNavClick: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.NavClick)
	}
}

func (l *SocketListener) handleClientGoto(client *shared.GameClient, body *models.ClientGotoBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientGoto: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Goto)
	}
}

func (l *SocketListener) handleClientOrbit(client *shared.GameClient, body *models.ClientOrbitBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientOrbit: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Orbit)
	}
}

func (l *SocketListener) handleClientDock(client *shared.GameClient, body *models.ClientDockBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientDock: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Dock)
	}
}

func (l *SocketListener) handleClientUndock(client *shared.GameClient, body *models.ClientUndockBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientUndock: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Undock)
	}
}

func (l *SocketListener) handleClientActivateModule(client *shared.GameClient, body *models.ClientActivateModuleBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientActivateModule: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ActivateModule)
	}
}

func (l *SocketListener) handleClientDeactivateModule(client *shared.GameClient, body *models.ClientDeactivateModuleBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientDeactivateModule: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.DeactivateModule)
	}
}

func (l *SocketListener) handleClientViewCargoBay(client *shared.GameClient, body *models.ClientViewCargoBayBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientViewCargoBay: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewCargoBay)
	}
}

func (l *SocketListener) handleClientUnfitModule(client *shared.GameClient, body *models.ClientUnfitModuleBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientUnfitModule: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.UnfitModule)
	}
}

func (l *SocketListener) handleClientTrashItem(client *shared.GameClient, body *models.ClientTrashItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientTrashItem: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.TrashItem)
	}
}

func (l *SocketListener) handleClientPackageItem(client *shared.GameClient, body *models.ClientPackageItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientPackageItem: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.PackageItem)
	}
}

func (l *SocketListener) handleClientUnpackageItem(client *shared.GameClient, body *models.ClientUnpackageItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientUnpackageItem: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.UnpackageItem)
	}
}

func (l *SocketListener) handleClientStackItem(client *shared.GameClient, body *models.ClientStackItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientStackItem: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.StackItem)
	}
}

func (l *SocketListener) handleClientSplitItem(client *shared.GameClient, body *models.ClientSplitItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientSplitItem: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SplitItem)
	}
}

func (l *SocketListener) handleClientFitModule(client *shared.GameClient, body *models.ClientFitModuleBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientFitModule: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.FitModule)
	}
}

func (l *SocketListener) handleClientSellAsOrder(client *shared.GameClient, body *models.ClientSellAsOrderBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientSellAsOrder: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SellAsOrder)
	}
}

func (l *SocketListener) handleClientViewOpenSellOrders(client *shared.GameClient, body *models.ClientViewOpenSellOrdersBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientViewOpenSellOrders: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewOpenSellOrders)
	}
}

func (l *SocketListener) handleClientBuySellOrder(client *shared.GameClient, body *models.ClientBuySellOrderBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientBuySellOrder: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.BuySellOrder)
	}
}

func (l *SocketListener) handleClientViewIndustrialOrders(client *shared.GameClient, body *models.ClientViewIndustrialOrdersBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientViewIndustrialOrders: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewIndustrialOrders)
	}
}

func (l *SocketListener) handleClientBuyFromSilo(client *shared.GameClient, body *models.ClientBuyFromSiloBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientBuyFromSilo: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.BuyFromSilo)
	}
}

func (l *SocketListener) handleClientSellToSilo(client *shared.GameClient, body *models.ClientSellToSiloBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientSellToSilo: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SellToSilo)
	}
}

func (l *SocketListener) handleClientViewStarMap(client *shared.GameClient, body *models.ClientViewStarMapBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientViewStarMap: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewStarMap)
	}
}

func (l *SocketListener) handleClientConsumeFuel(client *shared.GameClient, body *models.ClientConsumeFuelBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		log.Println(fmt.Sprintf("handleClientConsumeFuel: id spoof attempt: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.NewMessageRegistry()

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ConsumeFuel)
	}
}

// Adds a client to the server
func (l *SocketListener) addClient(c *shared.GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.clients = append(l.clients, c)
}

// Removes a client from the server
func (l *SocketListener) removeClient(c *shared.GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	// find the client to remove
	e := -1
	for i, s := range l.clients {
		if s == c {
			e = i
			break
		}
	}

	// remove client
	if e > -1 {
		t := len(l.clients)

		x := l.clients[t-1]
		l.clients[e] = x

		l.clients = l.clients[:t-1]
	}
}
