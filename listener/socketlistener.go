package listener

import (
	"encoding/json"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var _ = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
} // use default options

// Listener for handling and dispatching incoming websocket messages
type SocketListener struct {
	Engine  *engine.HeliaEngine
	clients []*shared.GameClient
	lock    sync.Mutex
}

// Initializes the internals of the SocketListener
func (s *SocketListener) Initialize() {}

// Handles a client joining the server
func (l *SocketListener) HandleConnect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	// upgrade to websocket connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)

	// short wait
	time.Sleep(1000 * time.Millisecond)

	// return if protocol change failed
	if err != nil {
		shared.TeeLog(fmt.Sprintf("ws protocol update error: %v", err))
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

	// defer saves on disconnect
	defer func(client *shared.GameClient) {
		if client == nil {
			return
		}

		if !client.Joined {
			return
		}

		// get services
		userSvc := sql.GetUserService()

		// save reputation sheet
		client.ReputationSheet.Lock.Lock()
		defer client.ReputationSheet.Lock.Unlock()

		srs := engine.SQLFromPlayerReputationSheet(&client.ReputationSheet)
		err := userSvc.UpdateReputationSheet(*client.UID, srs)

		if err != nil {
			shared.TeeLog(fmt.Sprintf("! unable to save reputation sheet for %v on disconnect! %v", client.UID, err))
		}

		// save experience sheet
		client.ExperienceSheet.Lock.Lock()
		defer client.ExperienceSheet.Lock.Unlock()

		ers := engine.SQLFromPlayerExperienceSheet(&client.ExperienceSheet)
		err = userSvc.UpdateExperienceSheet(*client.UID, ers)

		if err != nil {
			shared.TeeLog(fmt.Sprintf("! unable to save experience sheet for %v on disconnect! %v", client.UID, err))
		}
	}(&client)

	// get message type registry
	msgRegistry := models.SharedMessageRegistry

	// socket listener loop
	for {
		// read message from client
		_, r, err := c.ReadMessage()
		m := models.GameMessage{}

		// exit if read failed
		if err != nil {
			shared.TeeLog(fmt.Sprintf("ws read error: %v", err.Error()))
			break
		}

		// unmarshall
		json.Unmarshal(r, &m)

		// handle message based on type
		if m.MessageType == msgRegistry.GlobalAck {
			// decode body as ClientGlobalAckBody
			b := models.ClientGlobalAckBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientGlobalAck(&client, &b)
		} else if m.MessageType == msgRegistry.Join {
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
		} else if m.MessageType == msgRegistry.SelfDestruct {
			// decode body as ClientSelfDestructBody
			b := models.ClientSelfDestructBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientSelfDestruct(&client, &b)
		} else if m.MessageType == msgRegistry.ConsumeRepairKit {
			// decode body as ClientConsumeRepairKitBody
			b := models.ClientConsumeRepairKitBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientConsumeRepairKit(&client, &b)
		} else if m.MessageType == msgRegistry.ViewProperty {
			// decode body as ClientViewPropertyBody
			b := models.ClientViewPropertyBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewProperty(&client, &b)
		} else if m.MessageType == msgRegistry.Board {
			// decode body as ClientBoardBody
			b := models.ClientBoardBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientBoard(&client, &b)
		} else if m.MessageType == msgRegistry.TransferCredits {
			// decode body as ClientTransferCreditsBody
			b := models.ClientTransferCreditsBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientTransferCredits(&client, &b)
		} else if m.MessageType == msgRegistry.SellShipAsOrder {
			// decode body as ClientSellShipAsOrderBody
			b := models.ClientSellShipAsOrderBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientSellShipAsOrder(&client, &b)
		} else if m.MessageType == msgRegistry.TrashShip {
			// decode body as ClientTrashShipBody
			b := models.ClientTrashShipBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientTrashShip(&client, &b)
		} else if m.MessageType == msgRegistry.RenameShip {
			// decode body as ClientRenameShipBody
			b := models.ClientRenameShipBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientRenameShip(&client, &b)
		} else if m.MessageType == msgRegistry.PostSystemChatMessage {
			// decode body as ClientRenameShipBody
			b := models.ClientPostSystemChatMessageBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientPostSystemChatMessage(&client, &b)
		} else if m.MessageType == msgRegistry.TransferItem {
			// decode body as ClientTransferItemBody
			b := models.ClientTransferItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientTransferItem(&client, &b)
		} else if m.MessageType == msgRegistry.ViewExperience {
			// decode body as ClientViewExperienceBody
			b := models.ClientViewExperienceBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewExperience(&client, &b)
		} else if m.MessageType == msgRegistry.ViewSchematicRuns {
			// decode body as ClientViewSchematicRunsBody
			b := models.ClientViewSchematicRunsBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewSchematicRuns(&client, &b)
		} else if m.MessageType == msgRegistry.RunSchematic {
			// decode body as ClientRunSchematicBody
			b := models.ClientRunSchematicBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientRunSchematic(&client, &b)
		} else if m.MessageType == msgRegistry.CreateNewFaction {
			// decode body as ClientCreateNewFactionBody
			b := models.ClientCreateNewFactionBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientCreateNewFaction(&client, &b)
		} else if m.MessageType == msgRegistry.LeaveFaction {
			// decode body as ClientLeaveFactionBody
			b := models.ClientLeaveFactionBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientLeaveFaction(&client, &b)
		} else if m.MessageType == msgRegistry.ApplyToFaction {
			// decode body as ClientApplyToFactionBody
			b := models.ClientApplyToFactionBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientApplyToFaction(&client, &b)
		} else if m.MessageType == msgRegistry.ViewApplications {
			// decode body as ClientViewApplicationsBody
			b := models.ClientViewApplicationsBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewApplications(&client, &b)
		} else if m.MessageType == msgRegistry.ApproveApplication {
			// decode body as ClientApproveApplicationBody
			b := models.ClientApproveApplicationBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientApproveApplication(&client, &b)
		} else if m.MessageType == msgRegistry.RejectApplication {
			// decode body as ClientRejectApplicationBody
			b := models.ClientRejectApplicationBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientRejectApplication(&client, &b)
		} else if m.MessageType == msgRegistry.ViewMembers {
			// decode body as ClientViewMembersBody
			b := models.ClientViewMembersBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewMembers(&client, &b)
		} else if m.MessageType == msgRegistry.KickMember {
			// decode body as ClientKickMemberBody
			b := models.ClientKickMemberBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientKickMember(&client, &b)
		} else if m.MessageType == msgRegistry.UseModKit {
			// decode body as ClientUseModKitBody
			b := models.ClientUseModKitBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientUseModKit(&client, &b)
		} else if m.MessageType == msgRegistry.ViewActionReportsPage {
			// decode body as ClientViewActionReportsPageBody
			b := models.ClientViewActionReportsPageBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewActionReportsPage(&client, &b)
		} else if m.MessageType == msgRegistry.ViewActionReportDetail {
			// decode body as ClientViewActionReportDetailBody
			b := models.ClientViewActionReportDetailBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewActionReportDetail(&client, &b)
		} else if m.MessageType == msgRegistry.GiveItem {
			// decode body as ClientGiveItemBody
			b := models.ClientGiveItemBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientGiveItem(&client, &b)
		} else if m.MessageType == msgRegistry.ViewDockedUsers {
			// decode body as ClientViewDockedUsersBody
			b := models.ClientViewDockedUsersBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientViewDockedUsers(&client, &b)
		} else if m.MessageType == msgRegistry.ConsumeOutpostKit {
			// decode body as ClientConsumeOutpostKitBody
			b := models.ClientConsumeOutpostKitBody{}
			json.Unmarshal([]byte(m.MessageBody), &b)

			// handle message
			l.handleClientConsumeOutpostKit(&client, &b)
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
	shared.TeeLog(fmt.Sprintf("player join attempt: %v", &body.SessionID))

	// initialize services
	sessionSvc := sql.GetSessionService()
	userSvc := sql.GetUserService()
	msgRegistry := models.SharedMessageRegistry
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

		// store developer status
		client.IsDev = u.IsDev

		// obtain factions read lock
		l.Engine.Universe.FactionsLock.RLock()
		defer l.Engine.Universe.FactionsLock.RUnlock()

		// lookup current ship in memory
		currShip := l.Engine.Universe.FindShip(*u.CurrentShipID, nil)

		if currShip == nil {
			// they must have registered today - get their ship from the db
			dbShip, _ := shipSvc.GetShipByID(*u.CurrentShipID, false)

			if dbShip == nil {
				// report error
				shared.TeeLog(fmt.Sprintf("player join error: unable to find ship %v for %v", u.CurrentShipID, u.ID))

				// try to recover by creating a noob ship for player
				nStart, err := startSvc.GetStartByID(u.StartID)

				if err != nil {
					// dump error to console
					shared.TeeLog(fmt.Sprintf("player join recovery - unable to get start for player: %v", err))
					return
				}

				// check if their home station is overriden by faction membership
				uf := l.Engine.Universe.Factions[u.CurrentFactionID.String()]

				if uf == nil {
					shared.TeeLog(fmt.Sprintf("player join recovery - no faction: %v", u.CurrentFactionID))
					return
				}

				if !uf.IsNPC && uf.HomeStationID != nil {
					// override home station on start in-memory only
					nStart.HomeStationID = *uf.HomeStationID

					// find their home station
					home := l.Engine.Universe.FindStation(nStart.HomeStationID, nil)

					if home == nil {
						shared.TeeLog(fmt.Sprintf("player join recovery - no home station: %v", nStart.HomeStationID))
						return
					}

					// make sure correct systemid is stored in-memory
					if !uf.IsNPC && uf.HomeStationID != nil {
						nStart.SystemID = home.SystemID
					}
				}

				u, err = engine.CreateNoobShipForPlayer(nStart, u.ID)

				if err != nil {
					// dump error to console
					shared.TeeLog(fmt.Sprintf("player join recovery - unable to create noob ship for player: %v", err))
					return
				}

				// get their new noob ship from the db
				dbShip, err = shipSvc.GetShipByID(*u.CurrentShipID, false)

				if dbShip == nil || err != nil {
					shared.TeeLog(fmt.Sprintf("player join recovery: unable to find recovery noob ship %v for %v", u.CurrentShipID, u.ID))
					return
				}
			}

			// build in-memory ship
			currShip, err = engine.LoadShip(dbShip, l.Engine.Universe)

			if dbShip == nil {
				shared.TeeLog(fmt.Sprintf("player join error: %v", err))
				return
			}
		}

		// obtain ship lock
		currShip.Lock.Lock()
		defer currShip.Lock.Unlock()

		// load reputation sheet
		client.ReputationSheet = shared.PlayerReputationSheet{
			FactionEntries: make(map[string]*shared.PlayerReputationSheetFactionEntry),
			UserID:         u.ID,
			CharacterName:  u.CharacterName,
		}

		for k, v := range u.ReputationSheet.FactionEntries {
			client.ReputationSheet.FactionEntries[k] = &shared.PlayerReputationSheetFactionEntry{
				FactionID:        v.FactionID,
				StandingValue:    v.StandingValue,
				AreOpenlyHostile: v.AreOpenlyHostile,
			}
		}

		// inject reputation sheet
		currShip.ReputationSheet = &client.ReputationSheet

		// load experience sheet
		client.ExperienceSheet = shared.PlayerExperienceSheet{
			ShipExperience:   make(map[string]*shared.ShipExperienceEntry),
			ModuleExperience: make(map[string]*shared.ModuleExperienceEntry),
		}

		for k, v := range u.ExperienceSheet.ShipExperience {
			client.ExperienceSheet.ShipExperience[k] = &shared.ShipExperienceEntry{
				SecondsOfExperience: v.SecondsOfExperience,
				ShipTemplateID:      v.ShipTemplateID,
				ShipTemplateName:    v.ShipTemplateName,
			}
		}

		for k, v := range u.ExperienceSheet.ModuleExperience {
			client.ExperienceSheet.ModuleExperience[k] = &shared.ModuleExperienceEntry{
				SecondsOfExperience: v.SecondsOfExperience,
				ItemTypeID:          v.ItemTypeID,
				ItemTypeName:        v.ItemTypeName,
			}
		}

		// inject experience sheet
		currShip.ExperienceSheet = &client.ExperienceSheet

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
			Mass:      currShip.GetRealMass(),
			Radius:    currShip.TemplateData.Radius,
			ShieldP:   (currShip.Shield / currShip.GetRealMaxShield(true)) * 100,
			ArmorP:    (currShip.Armor / currShip.GetRealMaxArmor(true)) * 100,
			HullP:     (currShip.Hull / currShip.GetRealMaxHull(true)) * 100,
			FactionID: u.CurrentFactionID,
			// secret stuff
			EnergyP: (currShip.Energy / currShip.GetRealMaxEnergy(true)) * 100,
			HeatP:   (currShip.Heat / currShip.GetRealMaxHeat(true)) * 100,
			FuelP:   (currShip.Fuel / currShip.GetRealMaxFuel(true)) * 100,
			Accel:   currShip.GetRealAccel(true),
			Turn:    currShip.GetRealTurn(true),
			// secret caches
			CachedHeatSink:       currShip.GetRealHeatSink(true),
			CachedMaxHeat:        currShip.GetRealMaxHeat(true),
			CachedRealSpaceDrag:  currShip.GetRealSpaceDrag(true),
			CachedMaxFuel:        currShip.GetRealMaxFuel(true),
			CachedMaxEnergy:      currShip.GetRealMaxEnergy(true),
			CachedMaxShield:      currShip.GetRealMaxShield(true),
			CachedMaxArmor:       currShip.GetRealMaxArmor(true),
			CachedMaxHull:        currShip.GetRealMaxHull(true),
			CachedEnergyRegen:    currShip.GetRealEnergyRegen(true),
			CachedShieldRegen:    currShip.GetRealShieldRegen(true),
			CachedCargoBayVolume: currShip.GetRealCargoBayVolume(true),
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

		// short sleep
		time.Sleep(1000 * time.Millisecond)

		// prepare initial faction info for client
		af := models.ServerFactionUpdateBody{
			Factions: make([]models.ServerFactionBody, 0),
		}

		factions := l.Engine.Universe.Factions

		for _, f := range factions {
			// include relationship data
			rels := make([]models.ServerFactionRelationship, 0)

			for _, rel := range f.ReputationSheet.Entries {
				rels = append(rels, models.ServerFactionRelationship{
					FactionID:        rel.TargetFactionID,
					AreOpenlyHostile: rel.AreOpenlyHostile,
					StandingValue:    rel.StandingValue,
				})
			}

			ax := models.ServerFactionBody{
				ID:            f.ID,
				Name:          f.Name,
				Description:   f.Description,
				IsNPC:         f.IsNPC,
				IsJoinable:    f.IsJoinable,
				CanHoldSov:    f.CanHoldSov,
				Ticker:        f.Ticker,
				Relationships: rels,
			}

			if f.OwnerID != nil {
				ax.OwnerID = f.OwnerID.String()
			}

			af.Factions = append(af.Factions, ax)

			// fill neutral entry into rep sheet if missing
			r := client.ReputationSheet.FactionEntries[f.ID.String()]

			if r == nil {
				client.ReputationSheet.FactionEntries[f.ID.String()] = &shared.PlayerReputationSheetFactionEntry{
					FactionID:        f.ID,
					StandingValue:    0,
					AreOpenlyHostile: false,
				}
			}
		}

		// package faction data
		b, _ = json.Marshal(&af)

		msg = models.GameMessage{
			MessageType: msgRegistry.FactionUpdate,
			MessageBody: string(b),
		}

		// send all faction update message to client
		client.WriteMessage(&msg)

		// short sleep
		time.Sleep(1000 * time.Millisecond)

		// debug out
		shared.TeeLog(fmt.Sprintf("player joined: %v", &body.SessionID))

		// mark as successfully joined
		client.Joined = true
	} else {
		// dump error to console
		shared.TeeLog(fmt.Sprintf("player join error: %v", err))
	}
}

func (l *SocketListener) handleClientGlobalAck(client *shared.GameClient, body *models.ClientGlobalAckBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientGlobalAck: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.GlobalAck, false)
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
		shared.TeeLog(fmt.Sprintf("handleClientNavClick: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.NavClick, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientGoto: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Goto, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientOrbit: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Orbit, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientDock: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Dock, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientUndock: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Undock, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientActivateModule: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ActivateModule, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientDeactivateModule: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.DeactivateModule, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientViewCargoBay: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewCargoBay, false)
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
		shared.TeeLog(fmt.Sprintf("handleClientUnfitModule: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.UnfitModule, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientTrashItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.TrashItem, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientPackageItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.PackageItem, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientUnpackageItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.UnpackageItem, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientStackItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.StackItem, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientSplitItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SplitItem, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientFitModule: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.FitModule, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientSellAsOrder: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SellAsOrder, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientViewOpenSellOrders: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewOpenSellOrders, false)
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
		shared.TeeLog(fmt.Sprintf("handleClientBuySellOrder: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.BuySellOrder, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientViewIndustrialOrders: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewIndustrialOrders, false)
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
		shared.TeeLog(fmt.Sprintf("handleClientBuyFromSilo: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.BuyFromSilo, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientSellToSilo: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SellToSilo, true)
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
		shared.TeeLog(fmt.Sprintf("handleClientViewStarMap: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewStarMap, false)
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
		shared.TeeLog(fmt.Sprintf("handleClientConsumeFuel: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ConsumeFuel, true)
	}
}

func (l *SocketListener) handleClientSelfDestruct(client *shared.GameClient, body *models.ClientSelfDestructBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientSelfDestruct: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SelfDestruct, true)
	}
}

func (l *SocketListener) handleClientConsumeRepairKit(client *shared.GameClient, body *models.ClientConsumeRepairKitBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientConsumeRepairKit: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ConsumeRepairKit, true)
	}
}

func (l *SocketListener) handleClientViewProperty(client *shared.GameClient, body *models.ClientViewPropertyBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewProperty: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewProperty, false)
	}
}

func (l *SocketListener) handleClientBoard(client *shared.GameClient, body *models.ClientBoardBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientBoard: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.Board, true)
	}
}

func (l *SocketListener) handleClientTransferCredits(client *shared.GameClient, body *models.ClientTransferCreditsBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientTransferCredits: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.TransferCredits, true)
	}
}

func (l *SocketListener) handleClientSellShipAsOrder(client *shared.GameClient, body *models.ClientSellShipAsOrderBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientSellShipAsOrder: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.SellShipAsOrder, true)
	}
}

func (l *SocketListener) handleClientTrashShip(client *shared.GameClient, body *models.ClientTrashShipBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientTrashShip: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.TrashShip, true)
	}
}

func (l *SocketListener) handleClientRenameShip(client *shared.GameClient, body *models.ClientRenameShipBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientRenameShip: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.RenameShip, true)
	}
}

func (l *SocketListener) handleClientPostSystemChatMessage(client *shared.GameClient, body *models.ClientPostSystemChatMessageBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientPostSystemChatMessageBody: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.PostSystemChatMessage, true)
	}
}

func (l *SocketListener) handleClientTransferItem(client *shared.GameClient, body *models.ClientTransferItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientTransferItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.TransferItem, true)
	}
}

func (l *SocketListener) handleClientViewExperience(client *shared.GameClient, body *models.ClientViewExperienceBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewExperience: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewExperience, false)
	}
}

func (l *SocketListener) handleClientViewSchematicRuns(client *shared.GameClient, body *models.ClientViewSchematicRunsBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewSchematicRuns: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewSchematicRuns, false)
	}
}

func (l *SocketListener) handleClientRunSchematic(client *shared.GameClient, body *models.ClientRunSchematicBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientRunSchematic: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.RunSchematic, true)
	}
}

func (l *SocketListener) handleClientCreateNewFaction(client *shared.GameClient, body *models.ClientCreateNewFactionBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientCreateNewFaction: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.CreateNewFaction, true)
	}
}

func (l *SocketListener) handleClientLeaveFaction(client *shared.GameClient, body *models.ClientLeaveFactionBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientLeaveFaction: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.LeaveFaction, true)
	}
}

func (l *SocketListener) handleClientApplyToFaction(client *shared.GameClient, body *models.ClientApplyToFactionBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientApplyToFaction: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ApplyToFaction, true)
	}
}

func (l *SocketListener) handleClientViewApplications(client *shared.GameClient, body *models.ClientViewApplicationsBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewApplications: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewApplications, false)
	}
}

func (l *SocketListener) handleClientApproveApplication(client *shared.GameClient, body *models.ClientApproveApplicationBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientApproveApplication: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ApproveApplication, true)
	}
}

func (l *SocketListener) handleClientRejectApplication(client *shared.GameClient, body *models.ClientRejectApplicationBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientRejectApplication: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.RejectApplication, true)
	}
}

func (l *SocketListener) handleClientViewMembers(client *shared.GameClient, body *models.ClientViewMembersBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewMembers: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewMembers, false)
	}
}

func (l *SocketListener) handleClientKickMember(client *shared.GameClient, body *models.ClientKickMemberBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientKickMember: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.KickMember, true)
	}
}

func (l *SocketListener) handleClientUseModKit(client *shared.GameClient, body *models.ClientUseModKitBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientUseModKit: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.UseModKit, true)
	}
}

func (l *SocketListener) handleClientViewActionReportsPage(client *shared.GameClient, body *models.ClientViewActionReportsPageBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewActionReportsPage: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewActionReportsPage, false)
	}
}

func (l *SocketListener) handleClientViewActionReportDetail(client *shared.GameClient, body *models.ClientViewActionReportDetailBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewActionReportDetail: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewActionReportDetail, false)
	}
}

func (l *SocketListener) handleClientGiveItem(client *shared.GameClient, body *models.ClientGiveItemBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientGiveItem: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.GiveItem, true)
	}
}

func (l *SocketListener) handleClientViewDockedUsers(client *shared.GameClient, body *models.ClientViewDockedUsersBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientViewDockedUsers: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ViewDockedUsers, false)
	}
}

func (l *SocketListener) handleClientConsumeOutpostKit(client *shared.GameClient, body *models.ClientConsumeOutpostKitBody) {
	// safety returns
	if body == nil {
		return
	}

	if client == nil {
		return
	}

	// verify session id
	if body.SessionID != *client.SID {
		shared.TeeLog(fmt.Sprintf("handleClientConsumeOutpostKit: unexpected id: %v vs %v", &body.SessionID, &client.SID))
	} else {
		// initialize services
		msgRegistry := models.SharedMessageRegistry

		// push event onto player's ship queue
		data := *body
		client.PushShipEvent(data, msgRegistry.ConsumeOutpostKit, false)
	}
}

// Adds a client to the server
func (l *SocketListener) addClient(c *shared.GameClient) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.clients = append(l.clients, c)

	// start client cache update goroutine
	go func(c *shared.GameClient) {
		for !c.Dead {
			if c.UID == nil {
				continue
			}

			// lookup all ships belonging to this player
			ownedShips := l.Engine.Universe.FindShipsByUserID(*c.UID, nil)

			// build property cache
			pc := shared.PropertyCache{}

			for _, os := range ownedShips {
				// copy entry
				osc := os.CopyShip(true)

				// copy guaranteed fields
				z := shared.ShipPropertyCacheEntry{
					Name:    osc.ShipName,
					Texture: osc.Texture,
					ShipID:  osc.ID,
					Wallet:  osc.Wallet,
				}

				z.SolarSystemID = osc.SystemID
				z.SolarSystemName = osc.SystemName

				// copy possibly null fields
				if osc.DockedAtStationID != nil {
					if osc.DockedAtStation != nil {
						n := osc.DockedAtStation.StationName

						z.DockedAtStationID = osc.DockedAtStationID
						z.DockedAtStationName = &n
					}
				}

				pc.ShipCaches = append(pc.ShipCaches, z)
			}

			// update property cache
			c.SetPropertyCache(pc)

			// wait 5 seconds
			time.Sleep(time.Second * 5)
		}
	}(c)
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

	// mark as dead
	c.Dead = true
}
