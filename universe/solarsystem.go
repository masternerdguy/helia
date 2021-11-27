package universe

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/physics"
	"helia/shared"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Structure representing a solar system
type SolarSystem struct {
	ID                uuid.UUID
	SystemName        string
	RegionID          uuid.UUID
	HoldingFactionID  uuid.UUID
	PosX              float64
	PosY              float64
	Universe          *Universe
	ships             map[string]*Ship
	stars             map[string]*Star
	planets           map[string]*Planet
	jumpholes         map[string]*Jumphole
	stations          map[string]*Station
	asteroids         map[string]*Asteroid
	clients           map[string]*shared.GameClient       // clients in this system
	pushModuleEffects []models.GlobalPushModuleEffectBody // module visual effect aggregation for tick
	pushPointEffects  []models.GlobalPushPointEffectBody  // non-module point visual effect aggregation for tick
	tickCounter       int                                 // counter that is used to control frequency of certain global updates
	Lock              sync.Mutex
	// event escalations to engine core
	PlayerNeedRespawn    map[string]*shared.GameClient // clients in need of a respawn by core
	NPCNeedRespawn       map[string]*Ship              // NPCs in need of a respawn by core
	DeadShips            map[string]*Ship              // dead ships in need of cleanup by core
	MovedItems           map[string]*Item              // items moved to a new container in need of saving by core
	PackagedItems        map[string]*Item              // items packaged in need of saving by core
	UnpackagedItems      map[string]*Item              // items unpackaged in need of saving by core
	ChangedQuantityItems map[string]*Item              // stacks of items that have changed quantity and need saving by core
	NewItems             map[string]*Item              // stacks of items that are newly created and need to be saved by core
	NewSellOrders        map[string]*SellOrder         // new sell orders in need of saving by core
	BoughtSellOrders     map[string]*SellOrder         // sell orders that have been fulfilled in need of saving by core
	NewShipPurchases     map[string]*NewShipPurchase   // newly purchased ships that need to be generated and saved by core
	ShipSwitches         map[string]*ShipSwitch        // approved requests to switch a client to another ship in need of saving by core
	SetNoLoad            map[string]*ShipNoLoadSet     // updates to the no load flag in need of saving by core
	UsedShipPurchases    map[string]*UsedShipPurchase  // purchased used ships that need to be hooked in and saved by core
}

// Initializes internal aspects of SolarSystem
func (s *SolarSystem) Initialize() {
	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// initialize maps
	s.clients = make(map[string]*shared.GameClient)
	s.ships = make(map[string]*Ship)
	s.stars = make(map[string]*Star)
	s.planets = make(map[string]*Planet)
	s.jumpholes = make(map[string]*Jumphole)
	s.stations = make(map[string]*Station)
	s.asteroids = make(map[string]*Asteroid)
	s.DeadShips = make(map[string]*Ship)
	s.PlayerNeedRespawn = make(map[string]*shared.GameClient)
	s.NPCNeedRespawn = make(map[string]*Ship)
	s.MovedItems = make(map[string]*Item)
	s.PackagedItems = make(map[string]*Item)
	s.UnpackagedItems = make(map[string]*Item)
	s.ChangedQuantityItems = make(map[string]*Item)
	s.NewItems = make(map[string]*Item)
	s.NewSellOrders = make(map[string]*SellOrder)
	s.BoughtSellOrders = make(map[string]*SellOrder)
	s.NewShipPurchases = make(map[string]*NewShipPurchase)
	s.ShipSwitches = make(map[string]*ShipSwitch)
	s.SetNoLoad = make(map[string]*ShipNoLoadSet)
	s.UsedShipPurchases = make(map[string]*UsedShipPurchase)

	// initialize slices
	s.pushModuleEffects = make([]models.GlobalPushModuleEffectBody, 0)
	s.pushPointEffects = make([]models.GlobalPushPointEffectBody, 0)
}

// Performs one-time shutdown actions when the server is stopping
func (s *SolarSystem) HandleShutdownSignal() {
	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// propagate shutdown to connected clients
	for _, c := range s.clients {
		c.Dead = true
	}
}

// Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// skip if nobody in system
	if len(s.ships) == 0 {
		return
	}

	// check tick counter to determine whether to send static world data
	sendStatic := s.tickCounter > 50

	// check tick counter to determine whether to send secret updates
	sendSecret := s.tickCounter%4 == 0

	// check tick counter to determine whether to send player rep sheets
	sendPlayerRepSheets := s.tickCounter%8 == 0

	if sendStatic {
		// reset tick counter
		s.tickCounter = 0
	}

	// get message registry
	msgRegistry := models.NewMessageRegistry()

	// process player current ship event queue
	for _, c := range s.clients {
		evt := c.PopShipEvent()

		// skip if nothing to do
		if evt == nil {
			continue
		}

		// find player ship
		sh := s.ships[c.CurrentShipID.String()]

		// null check
		if sh == nil {
			continue
		}

		// associate escrow container id
		sh.EscrowContainerID = &c.EscrowContainerID

		// process event
		if evt.Type == models.NewMessageRegistry().NavClick {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientNavClickBody)

				// apply effect to player's current ship
				sh.CmdManualNav(data.ScreenTheta, data.ScreenMagnitude, false)
			}
		} else if evt.Type == models.NewMessageRegistry().Goto {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientGotoBody)

				// apply effect to player's current ship
				sh.CmdGoto(data.TargetID, data.Type, false)
			}
		} else if evt.Type == models.NewMessageRegistry().Orbit {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientOrbitBody)

				// apply effect to player's current ship
				sh.CmdOrbit(data.TargetID, data.Type, false)
			}
		} else if evt.Type == models.NewMessageRegistry().Dock {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientDockBody)

				// get registry
				targetTypeReg := models.NewTargetTypeRegistry()

				if data.Type == targetTypeReg.Station {
					// find station
					station := s.stations[string(data.TargetID.String())]

					if station == nil {
						c.WriteErrorMessage("docking denied - station not found")
						return
					}

					// check standings
					v, oh := sh.CheckStandings(station.FactionID)

					if oh {
						c.WriteErrorMessage("docking denied - openly hostile")
						return
					}

					if v < shared.MIN_DOCK_STANDING {
						c.WriteErrorMessage("docking denied - reputation too low")
						return
					}
				}

				// apply effect to player's current ship
				sh.CmdDock(data.TargetID, data.Type, false)
			}
		} else if evt.Type == models.NewMessageRegistry().Undock {
			if sh != nil {
				// extract data (currently nothing to process)
				// data := evt.Body.(models.ClientUndockBody)

				// apply effect to player's current ship
				sh.CmdUndock(false)
			}
		} else if evt.Type == models.NewMessageRegistry().ActivateModule {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientActivateModuleBody)

				// skip if rack c (passive modules)
				if data.Rack == "C" {
					continue
				} else {
					// find module
					mod := sh.FindModule(data.ItemID, data.Rack)

					// make sure we found something
					if mod == nil {
						// do nothing
						continue
					} else {
						if !mod.WillRepeat {
							// set repeat to true
							mod.WillRepeat = true

							// store target
							mod.TargetID = data.TargetID
							mod.TargetType = data.TargetType
						}
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().DeactivateModule {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientDeactivateModuleBody)

				// skip if rack c (passive modules)
				if data.Rack == "C" {
					continue
				} else {
					// find module
					mod := sh.FindModule(data.ItemID, data.Rack)

					// make sure we found something
					if mod == nil {
						// do nothing
						continue
					} else {
						if mod.WillRepeat {
							// set repeat to false
							mod.WillRepeat = false

							// clear target
							mod.TargetID = nil
							mod.TargetType = nil
						}
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ViewCargoBay {
			if sh != nil {
				// extract data (currently nothing to process)
				// data := evt.Body.(models.ClientViewCargoBayBody)

				// convert cargo bay to an update message for the client
				vw := models.ServerContainerViewBody{
					ContainerID: sh.CargoBayContainerID,
				}

				for _, i := range sh.CargoBay.Items {
					// skip if dirty
					if i.CoreDirty {
						continue
					}

					// skip if no quantity
					if i.Quantity <= 0 {
						continue
					}

					// add to message
					r := models.ServerItemViewBody{
						ID:             i.ID,
						ItemTypeID:     i.ItemTypeID,
						ItemTypeName:   i.ItemTypeName,
						ItemFamilyID:   i.ItemFamilyID,
						ItemFamilyName: i.ItemFamilyName,
						Quantity:       i.Quantity,
						IsPackaged:     i.IsPackaged,
						Meta:           models.Meta(i.Meta),
						ItemTypeMeta:   models.Meta(i.ItemTypeMeta),
					}

					vw.Items = append(vw.Items, r)
				}

				// package message
				b, _ := json.Marshal(&vw)

				cu := models.GameMessage{
					MessageType: msgRegistry.CargoBayUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&cu)
			}
		} else if evt.Type == models.NewMessageRegistry().UnfitModule {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientUnfitModuleBody)

				// find module
				mod := sh.FindModule(data.ItemID, data.Rack)

				// make sure we found something
				if mod == nil {
					// do nothing
					continue
				} else {
					// unfit module
					err := sh.UnfitModule(mod, false)

					// there are lots of reasons this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().TrashItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientTrashItemBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// trash item
					err := sh.TrashItemInCargo(item.ID, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().PackageItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientPackageItemBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// package item
					err := sh.PackageItemInCargo(item.ID, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().UnpackageItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientUnpackageItemBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// unpackage item
					err := sh.UnpackageItemInCargo(item.ID, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().StackItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientStackItemBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// stack item
					err := sh.StackItemInCargo(item.ID, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().SplitItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientSplitItemBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// split item
					err := sh.SplitItemInCargo(item.ID, data.Size, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().FitModule {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientFitModuleBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// fit module
					err := sh.FitModule(item.ID, false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().SellAsOrder {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientSellAsOrderBody)

				// find item
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// sell item stack as order
					err := sh.SellItemAsOrder(item.ID, float64(data.Price), false)

					// there is a reason this could fail the player will need to know about
					if err != nil {
						// send error message to client
						c.WriteErrorMessage(err.Error())
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().DeactivateModule {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientDeactivateModuleBody)

				// skip if rack c (passive modules)
				if data.Rack == "C" {
					continue
				} else {
					// find module
					mod := sh.FindModule(data.ItemID, data.Rack)

					// make sure we found something
					if mod == nil {
						// do nothing
						continue
					} else {
						if mod.WillRepeat {
							// set repeat to false
							mod.WillRepeat = false

							// clear target
							mod.TargetID = nil
							mod.TargetType = nil
						}
					}
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ViewOpenSellOrders {
			if sh != nil {
				// extract data (currently nothing to process)
				// data := evt.Body.(models.ClientViewOpenSellOrdersBody)

				// make sure the ship is docked
				if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
					continue
				}

				// convert station's open sell orders to an update message for the client
				vw := models.ServerOpenSellOrdersUpdateBody{}

				for _, i := range sh.DockedAtStation.OpenSellOrders {
					// skip if dirty
					if i.CoreDirty {
						continue
					}

					// skip if closed
					if i.BuyerUserID != nil || i.Bought != nil {
						continue
					}

					// make sure item is present
					if i.Item == nil {
						continue
					}

					// add to message
					item := models.ServerItemViewBody{
						ID:             i.Item.ID,
						ItemTypeID:     i.Item.ItemTypeID,
						ItemTypeName:   i.Item.ItemTypeName,
						ItemFamilyID:   i.Item.ItemFamilyID,
						ItemFamilyName: i.Item.ItemFamilyName,
						Quantity:       i.Item.Quantity,
						IsPackaged:     i.Item.IsPackaged,
						Meta:           models.Meta(i.Item.Meta),
						ItemTypeMeta:   models.Meta(i.Item.ItemTypeMeta),
					}

					order := models.ServerSellOrderBody{
						ID:           i.ID,
						StationID:    i.StationID,
						ItemID:       i.ItemID,
						SellerUserID: i.SellerUserID,
						AskPrice:     i.AskPrice,
						Created:      i.Created,
						Bought:       i.Bought,
						BuyerUserID:  i.BuyerUserID,
						Item:         item,
					}

					vw.Orders = append(vw.Orders, order)
				}

				// package message
				b, _ := json.Marshal(&vw)

				cu := models.GameMessage{
					MessageType: msgRegistry.OpenSellOrdersUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&cu)
			}
		} else if evt.Type == models.NewMessageRegistry().BuySellOrder {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientBuySellOrderBody)

				// sell item stack as order
				err := sh.BuyItemFromOrder(data.OrderID, false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ViewIndustrialOrders {
			if sh != nil {
				// extract data (currently nothing to process)
				// data := evt.Body.(models.ViewIndustrialOrders)

				// make sure the ship is docked
				if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
					continue
				}

				// convert station's open industrial orders to an update message for the client
				vw := models.ServerIndustrialOrdersUpdateBody{}

				// fill message
				for _, p := range sh.DockedAtStation.Processes {
					// copy I/O silos
					for k := range p.InternalState.Inputs {
						s := p.InternalState.Inputs[k]
						t := p.Process.Inputs[k]

						vw.InSilos = append(vw.InSilos, models.ServerIndustrialSiloBody{
							StationID:        p.StationID.String(),
							StationProcessID: p.ID.String(),
							ItemTypeID:       t.ItemTypeID.String(),
							ItemTypeName:     t.ItemTypeName,
							ItemFamilyID:     t.ItemFamilyID,
							ItemFamilyName:   t.ItemFamilyName,
							Price:            s.Price,
							Available:        s.Quantity,
							Meta:             models.Meta(t.Meta),
							ItemTypeMeta:     models.Meta(t.ItemTypeMeta),
							IsSelling:        false,
						})
					}

					for k := range p.InternalState.Outputs {
						s := p.InternalState.Outputs[k]
						t := p.Process.Outputs[k]

						vw.OutSilos = append(vw.OutSilos, models.ServerIndustrialSiloBody{
							StationID:        p.StationID.String(),
							StationProcessID: p.ID.String(),
							ItemTypeID:       t.ItemTypeID.String(),
							ItemTypeName:     t.ItemTypeName,
							ItemFamilyID:     t.ItemFamilyID,
							ItemFamilyName:   t.ItemFamilyName,
							Price:            s.Price,
							Available:        s.Quantity,
							Meta:             models.Meta(t.Meta),
							ItemTypeMeta:     models.Meta(t.ItemTypeMeta),
							IsSelling:        true,
						})
					}
				}

				// package message
				b, _ := json.Marshal(&vw)

				cu := models.GameMessage{
					MessageType: msgRegistry.IndustrialOrdersUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&cu)
			}
		} else if evt.Type == models.NewMessageRegistry().BuyFromSilo {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientBuyFromSiloBody)

				// buy product from silo
				err := sh.BuyItemFromSilo(data.SiloID, data.ItemTypeID, data.Quantity, false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().SellToSilo {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientSellToSiloBody)

				// sell item to silo
				err := sh.SellItemToSilo(data.SiloID, data.ItemID, data.Quantity, false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ViewStarMap {
			if sh != nil {
				// extract data
				// data := evt.Body.(models.ClientViewStarMapBody)

				// build message with cached map data
				m := models.ServerStarMapUpdateBody{
					CachedMapData: s.Universe.CachedMapData,
				}

				b, _ := json.Marshal(&m)

				cu := models.GameMessage{
					MessageType: msgRegistry.StarMapUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&cu)
			}
		} else if evt.Type == models.NewMessageRegistry().ConsumeFuel {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientConsumeFuelBody)

				// consume pellet
				err := sh.ConsumeFuelFromCargo(data.ItemID, false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().SelfDestruct {
			if sh != nil {
				// extract data
				//data := evt.Body.(models.ClientSelfDestructBody)

				// self destruct
				err := sh.SelfDestruct(false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ConsumeRepairKit {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientConsumeRepairKitBody)

				// consume repair kit
				err := sh.ConsumeRepairKitFromCargo(data.ItemID, false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				}
			}
		} else if evt.Type == models.NewMessageRegistry().ViewProperty {
			if sh != nil {
				// extract data
				// data := evt.Body.(models.ClientViewPropertyBody)

				// get property summary from client
				ps := c.GetPropertyCache()

				// build update for client
				cu := models.ServerPropertyUpdateBody{}

				for _, x := range ps.ShipCaches {
					cu.Ships = append(cu.Ships, models.ServerShipPropertyCacheEntry{
						Name:                x.Name,
						Texture:             x.Texture,
						ShipID:              x.ShipID,
						SolarSystemID:       x.SolarSystemID,
						SolarSystemName:     x.SolarSystemName,
						DockedAtStationID:   x.DockedAtStationID,
						DockedAtStationName: x.DockedAtStationName,
						Wallet:              x.Wallet,
					})
				}

				// package message
				b, _ := json.Marshal(&cu)

				z := models.GameMessage{
					MessageType: msgRegistry.PropertyUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&z)
			}
		} else if evt.Type == models.NewMessageRegistry().Board {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientBoardBody)

				// verify player is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to switch ships")
					continue
				}

				// get ship to board and verify it is owned by the player
				toBoard := s.ships[data.ShipID.String()]

				if toBoard == nil {
					c.WriteErrorMessage("ship not available to board")
					continue
				}

				if toBoard.UserID != sh.UserID {
					c.WriteErrorMessage("ship not available to board")
					continue
				}

				// verify it is docked at the same station as the player
				if toBoard.DockedAtStation == nil {
					c.WriteErrorMessage("both ships must be docked at the same station to switch ships")
					continue
				}

				if toBoard.DockedAtStation.ID != sh.DockedAtStation.ID {
					c.WriteErrorMessage("both ships must be docked at the same station to switch ships")
					continue
				}

				// verify it isn't the same ship
				if toBoard.ID == sh.ID {
					c.WriteErrorMessage("you are already flying this ship")
					continue
				}

				// escalate ship switch request to core
				key := fmt.Sprintf("%v>>%v", sh.ID, toBoard.ID)

				s.ShipSwitches[key] = &ShipSwitch{
					Client: c,
					Source: sh,
					Target: toBoard,
				}
			}
		} else if evt.Type == models.NewMessageRegistry().TransferCredits {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientTransferCreditsBody)

				// verify player is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to transfer money")
					continue
				}

				// get ship to exchange with and verify it is owned by the player
				toExchange := s.ships[data.ShipID.String()]

				if toExchange == nil {
					c.WriteErrorMessage("ship not available to exchange money with")
					continue
				}

				if toExchange.UserID != sh.UserID {
					c.WriteErrorMessage("ship not available to exchange money with")
					continue
				}

				// verify it is docked at the same station as the player
				if toExchange.DockedAtStation == nil {
					c.WriteErrorMessage("both ships must be docked at the same station to exchange money")
					continue
				}

				if toExchange.DockedAtStation.ID != sh.DockedAtStation.ID {
					c.WriteErrorMessage("both ships must be docked at the same station to exchange money")
					continue
				}

				// verify it isn't the same ship
				if toExchange.ID == sh.ID {
					c.WriteErrorMessage("you are currently flying this ship")
					continue
				}

				// verify this will not put either ship's balance below zero
				newSrcBalance := sh.Wallet - float64(data.Amount)
				newTgtBalance := toExchange.Wallet + float64(data.Amount)

				if newSrcBalance < 0 || newTgtBalance < 0 {
					c.WriteErrorMessage("insufficient funds")
					continue
				}

				// update balances
				sh.Wallet = newSrcBalance
				toExchange.Wallet = newTgtBalance

				// update property cache with new amounts (so it shows up immediately instead of as part of the periodic rebuild)
				pc := c.GetPropertyCache()

				for i, e := range pc.ShipCaches {
					if e.ShipID == sh.ID {
						pc.ShipCaches[i].Wallet = newSrcBalance
					} else if e.ShipID == toExchange.ID {
						pc.ShipCaches[i].Wallet = newTgtBalance
					}
				}

				c.SetPropertyCache(pc)
			}
		} else if evt.Type == models.NewMessageRegistry().SellShipAsOrder {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientSellShipAsOrderBody)

				// verify player is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to sell a ship")
					continue
				}

				// get ship to sell and verify it is owned by the player
				toSell := s.ships[data.ShipID.String()]

				if toSell == nil {
					c.WriteErrorMessage("ship not available to sell")
					continue
				}

				if toSell.UserID != sh.UserID {
					c.WriteErrorMessage("ship not available to sell")
					continue
				}

				// verify it is docked at the same station as the player
				if toSell.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked at the same station as the ship being sold")
					continue
				}

				if toSell.DockedAtStation.ID != sh.DockedAtStation.ID {
					c.WriteErrorMessage("you must be docked at the same station as the ship being sold")
					continue
				}

				// verify it isn't the same ship
				if toSell.ID == sh.ID {
					c.WriteErrorMessage("you are currently flying this ship")
					continue
				}

				// associate escrow container id with ship being sold
				toSell.EscrowContainerID = &c.EscrowContainerID

				// list ship on orders market
				err := toSell.SellSelfAsOrder(float64(data.Price), false)

				// there is a reason this could fail the player will need to know about
				if err != nil {
					// send error message to client
					c.WriteErrorMessage(err.Error())
				} else {
					// remove sold ship from property cache (so it goes away immediately instead of as part of the periodic rebuild)
					pc := c.GetPropertyCache()
					no := make([]shared.ShipPropertyCacheEntry, 0)

					for _, e := range pc.ShipCaches {
						if e.ShipID == toSell.ID {
							continue
						}

						no = append(no, e)
					}

					pc.ShipCaches = no
					c.SetPropertyCache(pc)
				}
			}
		} else if evt.Type == models.NewMessageRegistry().TrashShip {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientTrashShipBody)

				// verify player is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to trash a ship")
					continue
				}

				// get ship to sell and verify it is owned by the player
				toTrash := s.ships[data.ShipID.String()]

				if toTrash == nil {
					c.WriteErrorMessage("ship not available to trash")
					continue
				}

				if toTrash.UserID != sh.UserID {
					c.WriteErrorMessage("ship not available to trash")
					continue
				}

				// verify it is docked at the same station as the player
				if toTrash.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked at the same station as the ship being trashed")
					continue
				}

				if toTrash.DockedAtStation.ID != sh.DockedAtStation.ID {
					c.WriteErrorMessage("you must be docked at the same station as the ship being trashed")
					continue
				}

				// verify it isn't the same ship
				if toTrash.ID == sh.ID {
					c.WriteErrorMessage("you are currently flying this ship")
					continue
				}

				// unhook ship from system
				s.RemoveShip(toTrash, false)

				// escalate for saving in core
				nl := ShipNoLoadSet{
					ID:   toTrash.ID,
					Flag: true,
				}

				s.SetNoLoad[toTrash.ID.String()] = &nl

				// remove trashed ship from property cache (so it goes away immediately instead of as part of the periodic rebuild)
				pc := c.GetPropertyCache()
				no := make([]shared.ShipPropertyCacheEntry, 0)

				for _, e := range pc.ShipCaches {
					if e.ShipID == toTrash.ID {
						continue
					}

					no = append(no, e)
				}

				pc.ShipCaches = no
				c.SetPropertyCache(pc)
			}
		}
	}

	// update ships
	for _, e := range s.ships {
		// is hull at or below 0?
		if e.Hull <= 0 {
			now := time.Now()

			// mark as dead
			e.Destroyed = true
			e.DestroyedAt = &now

			// was this a ship actively being flown by a player?
			c := s.clients[e.UserID.String()]

			if c != nil {
				if c.CurrentShipID == e.ID {
					// escalate respawn request to core
					s.PlayerNeedRespawn[e.UserID.String()] = c
				}
			} else {
				// check if an NPC
				if e.IsNPC {
					// escalate NPC respawn request to core
					s.NPCNeedRespawn[e.UserID.String()] = e
				}
			}

			// escalate ship cleanup request to core
			s.DeadShips[e.ID.String()] = e

			// drop explosion for ship
			exp := models.GlobalPushPointEffectBody{
				GfxEffect: "basic_explosion",
				PosX:      e.PosX,
				PosY:      e.PosY,
				Radius:    e.TemplateData.Radius * 1.5,
			}

			s.pushPointEffects = append(s.pushPointEffects, exp)
		} else {
			// update ship
			e.PeriodicUpdate()
		}
	}

	// update npc stations
	for _, e := range s.stations {
		e.PeriodicUpdate()
	}

	// ship collission testing
	for _, sA := range s.ships {
		// skip dead ships
		if sA.Destroyed || sA.DestroyedAt != nil {
			continue
		}

		// skip docked ships
		if sA.DockedAtStationID != nil {
			continue
		}

		// with other ships
		for _, sB := range s.ships {
			// skip dead ships
			if sB.Destroyed || sB.DestroyedAt != nil {
				continue
			}

			// skip docked ships
			if sB.DockedAtStationID != nil {
				continue
			}

			if sA.ID != sB.ID {
				// get physics dummies
				dummyA := sA.ToPhysicsDummy()
				dummyB := sB.ToPhysicsDummy()

				// get distance between ships
				d := physics.Distance(dummyA, dummyB)

				// check for radius intersection
				if d <= (sA.TemplateData.Radius + sB.TemplateData.Radius) {
					// calculate collission results
					physics.ElasticCollide(&dummyA, &dummyB)

					// update ships with results
					sA.ApplyPhysicsDummy(dummyA)
					sB.ApplyPhysicsDummy(dummyB)
				}
			}
		}

		// with jumpholes
		for _, jB := range s.jumpholes {
			// get physics dummies
			dummyA := sA.ToPhysicsDummy()
			dummyB := jB.ToPhysicsDummy()

			// get distance between ship and jumphole
			d := physics.Distance(dummyA, dummyB)

			// check for deep radius intersection
			if d <= ((sA.TemplateData.Radius + jB.Radius) * 0.75) {
				// find client
				c := s.clients[sA.UserID.String()]

				// perform move at end of update cycle
				defer func(gsA *Ship, gjB *Jumphole, gc *shared.GameClient) {
					if gc != nil {
						// check if this was the current ship of a player
						if gsA.ID == gc.CurrentShipID {
							// move player to destination system
							gc.CurrentSystemID = gjB.OutSystemID

							defer gjB.OutSystem.AddClient(gc, true)
							defer s.RemoveClient(gc, false)
						}
					}

					// kill ship autopilot
					defer gsA.CmdAbort()

					// place ship on the opposite side of the hole
					riX := gjB.PosX - gsA.PosX
					riY := gjB.PosY - gsA.PosY

					gsA.PosX = (gjB.OutJumphole.PosX + (riX * 1.25))
					gsA.PosY = (gjB.OutJumphole.PosY + (riY * 1.25))

					// in the destination system
					gsA.SystemID = gjB.OutSystemID

					// perform move operation
					defer gjB.OutSystem.AddShip(gsA, true)
					defer s.RemoveShip(gsA, false)
				}(sA, jB, c)

				break
			}
		}
	}

	// build global update of non-secret info for clients
	gu := models.ServerGlobalUpdateBody{}
	gu.CurrentSystemInfo = models.CurrentSystemInfo{
		ID:         s.ID,
		SystemName: s.SystemName,
		FactionID:  s.HoldingFactionID,
	}

	for _, d := range s.ships {
		// skip dead ships
		if d.Destroyed || d.DestroyedAt != nil {
			continue
		}

		// skip docked ships
		if d.DockedAtStationID != nil {
			continue
		}

		// build ship info and append
		gu.Ships = append(gu.Ships, models.GlobalShipInfo{
			ID:        d.ID,
			UserID:    d.UserID,
			Created:   d.Created,
			ShipName:  d.ShipName,
			OwnerName: d.OwnerName,
			PosX:      d.PosX,
			PosY:      d.PosY,
			SystemID:  d.SystemID,
			Texture:   d.Texture,
			Theta:     d.Theta,
			VelX:      d.VelX,
			VelY:      d.VelY,
			Mass:      d.GetRealMass(),
			Radius:    d.TemplateData.Radius,
			ShieldP:   ((d.Shield / d.GetRealMaxShield()) * 100) + Epsilon,
			ArmorP:    ((d.Armor / d.GetRealMaxArmor()) * 100) + Epsilon,
			HullP:     ((d.Hull / d.GetRealMaxHull()) * 100) + Epsilon,
			FactionID: d.FactionID,
		})
	}

	if sendStatic {
		/*
		 * Performance note: This data is very static and can be sent rarely. Once every second is fine,
		 * and saves a significant amount of bandwidth. Note it needs to be sent often enough for someone
		 * jumping into the system to receive data about the celestial objects it contains within a
		 * reasonable amount of time (which I feel is ~1 second).
		 */

		for _, d := range s.stars {
			gu.Stars = append(gu.Stars, models.GlobalStarInfo{
				ID:       d.ID,
				SystemID: d.SystemID,
				PosX:     d.PosX,
				PosY:     d.PosY,
				Texture:  d.Texture,
				Radius:   d.Radius,
				Mass:     d.Mass,
				Theta:    d.Theta,
			})
		}

		for _, d := range s.planets {
			gu.Planets = append(gu.Planets, models.GlobalPlanetInfo{
				ID:         d.ID,
				SystemID:   d.SystemID,
				PlanetName: d.PlanetName,
				PosX:       d.PosX,
				PosY:       d.PosY,
				Texture:    d.Texture,
				Radius:     d.Radius,
				Mass:       d.Mass,
				Theta:      d.Theta,
			})
		}

		for _, d := range s.asteroids {
			gu.Asteroids = append(gu.Asteroids, models.GlobalAsteroidInfo{
				ID:       d.ID,
				SystemID: d.SystemID,
				Name:     d.Name,
				PosX:     d.PosX,
				PosY:     d.PosY,
				Texture:  d.Texture,
				Radius:   d.Radius,
				Mass:     d.Mass,
				Theta:    d.Theta,
			})
		}

		for _, d := range s.jumpholes {
			gu.Jumpholes = append(gu.Jumpholes, models.GlobalJumpholeInfo{
				ID:           d.ID,
				SystemID:     d.SystemID,
				OutSystemID:  d.OutSystemID,
				JumpholeName: d.JumpholeName,
				PosX:         d.PosX,
				PosY:         d.PosY,
				Texture:      d.Texture,
				Radius:       d.Radius,
				Mass:         d.Mass,
				Theta:        d.Theta,
			})
		}

		for _, d := range s.stations {
			gu.Stations = append(gu.Stations, models.GlobalStationInfo{
				ID:          d.ID,
				SystemID:    d.SystemID,
				StationName: d.StationName,
				PosX:        d.PosX,
				PosY:        d.PosY,
				Texture:     d.Texture,
				Radius:      d.Radius,
				Mass:        d.Mass,
				Theta:       d.Theta,
				FactionID:   d.FactionID,
			})
		}
	}

	gu.NewModuleEffects = append(gu.NewModuleEffects, s.pushModuleEffects...)
	gu.NewPointEffects = append(gu.NewPointEffects, s.pushPointEffects...)

	// serialize global update
	b, _ := json.Marshal(&gu)

	msg := models.GameMessage{
		MessageType: msgRegistry.GlobalUpdate,
		MessageBody: string(b),
	}

	// write global update to clients
	for _, c := range s.clients {
		/*
		 * Performance note: This message must be sent to every client in the
		 * system in every tick. Otherwise, movement will be choppy and the
		 * game will feel laggy and unresponsive.
		 */

		c.WriteMessage(&msg)
	}

	// write secret current ship updates to individual clients
	if sendSecret {
		/*
		 * Performance note: This message is sent less often than the global update because the data
		 * contained doesn't need to be as fresh for the game to feel smooth and responsive, and
		 * sending it just a little less often significantly reduced bandwidth usage.
		 */

		for _, c := range s.clients {
			// find current ship
			d := s.ships[c.CurrentShipID.String()]

			if d == nil {
				continue
			}

			// build fitting status info
			fs := models.ServerFittingStatusBody{}

			// rack a
			rackA := models.ServerRackStatusBody{}

			for idx, v := range d.Fitting.ARack {
				// build module statusinfo
				module := copyModuleInfo(v)

				// include slot info
				slot := d.TemplateData.SlotLayout.ASlots[idx]
				module.HardpointFamily = slot.Family
				module.HardpointVolume = slot.Volume

				// store on message
				rackA.Modules = append(rackA.Modules, module)
			}

			// rack b
			rackB := models.ServerRackStatusBody{}

			for idx, v := range d.Fitting.BRack {
				// build module statusinfo
				module := copyModuleInfo(v)

				// include slot info
				slot := d.TemplateData.SlotLayout.BSlots[idx]
				module.HardpointFamily = slot.Family
				module.HardpointVolume = slot.Volume

				// store on message
				rackB.Modules = append(rackB.Modules, module)
			}

			// rack c
			rackC := models.ServerRackStatusBody{}

			for idx, v := range d.Fitting.CRack {
				// build module statusinfo
				module := copyModuleInfo(v)

				// include slot info
				slot := d.TemplateData.SlotLayout.CSlots[idx]
				module.HardpointFamily = slot.Family
				module.HardpointVolume = slot.Volume

				// store on message
				rackC.Modules = append(rackC.Modules, module)
			}

			// store rack info on fitting info
			fs.ARack = rackA
			fs.BRack = rackB
			fs.CRack = rackC

			// build current ship info message
			si := models.ServerCurrentShipUpdate{
				CurrentShipInfo: models.CurrentShipInfo{
					// global stuff
					ID:        d.ID,
					UserID:    d.UserID,
					Created:   d.Created,
					ShipName:  d.ShipName,
					PosX:      d.PosX,
					PosY:      d.PosY,
					SystemID:  d.SystemID,
					Texture:   d.Texture,
					Theta:     d.Theta,
					VelX:      d.VelX,
					VelY:      d.VelY,
					Mass:      d.GetRealMass(),
					Radius:    d.TemplateData.Radius,
					ShieldP:   ((d.Shield / d.GetRealMaxShield()) * 100) + Epsilon,
					ArmorP:    ((d.Armor / d.GetRealMaxArmor()) * 100) + Epsilon,
					HullP:     ((d.Hull / d.GetRealMaxHull()) * 100) + Epsilon,
					FactionID: d.FactionID,
					// secret stuff
					EnergyP:           ((d.Energy / d.GetRealMaxEnergy()) * 100) + Epsilon,
					HeatP:             ((d.Heat / d.GetRealMaxHeat()) * 100) + Epsilon,
					FuelP:             ((d.Fuel / d.GetRealMaxFuel()) * 100) + Epsilon,
					FitStatus:         fs,
					DockedAtStationID: d.DockedAtStationID,
					CargoP:            ((d.TotalCargoBayVolumeUsed(false) / d.GetRealCargoBayVolume()) * 100) + Epsilon,
					Wallet:            d.Wallet,
				},
			}

			// serialize secret current ship update
			b, _ := json.Marshal(&si)

			sct := models.GameMessage{
				MessageType: msgRegistry.CurrentShipUpdate,
				MessageBody: string(b),
			}

			// write message to client
			c.WriteMessage(&sct)
		}
	}

	// write secret player rep sheet updates to individual clients
	if sendPlayerRepSheets {
		/*
		 * Performance note: This message is sent less often than the global update because the data
		 * contained doesn't need to be as fresh for the game to feel smooth and responsive, and
		 * sending it just a little less often significantly reduced bandwidth usage.
		 */

		for _, c := range s.clients {
			// find current ship
			d := s.ships[c.CurrentShipID.String()]

			if d == nil {
				continue
			}

			// build relationship message
			rm := models.ServerPlayerFactionUpdateBody{}
			rels := make([]models.ServerPlayerFactionRelationship, 0)

			for _, v := range c.ReputationSheet.FactionEntries {
				rels = append(rels, models.ServerPlayerFactionRelationship{
					FactionID:        v.FactionID,
					AreOpenlyHostile: v.AreOpenlyHostile,
					StandingValue:    v.StandingValue,
					IsMember:         v.FactionID == d.FactionID,
				})
			}

			rm.Factions = rels

			// serialize player-faction rep sheet update
			b, _ := json.Marshal(&rm)

			sct := models.GameMessage{
				MessageType: msgRegistry.PlayerFactionUpdate,
				MessageBody: string(b),
			}

			// write message to client
			c.WriteMessage(&sct)
		}
	}

	// reset new visual effects for next tick
	s.pushModuleEffects = make([]models.GlobalPushModuleEffectBody, 0)
	s.pushPointEffects = make([]models.GlobalPushPointEffectBody, 0)

	// increment tick counter
	s.tickCounter++
}

// Adds a ship to the system
func (s *SolarSystem) AddShip(c *Ship, lock bool) {
	// safety check
	if c == nil {
		return
	}

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// store pointer to system
	c.CurrentSystem = s

	// disarm self destruct
	c.DestructArmed = false

	// add ship
	s.ships[c.ID.String()] = c
}

// Removes a ship from the system
func (s *SolarSystem) RemoveShip(c *Ship, lock bool) {
	// safety check
	if c == nil {
		return
	}

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// remove ship
	delete(s.ships, c.ID.String())
}

// Adds a star to the system
func (s *SolarSystem) AddStar(c *Star) {
	// safety check
	if c == nil {
		return
	}

	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// add star
	s.stars[c.ID.String()] = c
}

// Adds a planet to the system
func (s *SolarSystem) AddPlanet(c *Planet) {
	// safety check
	if c == nil {
		return
	}

	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// add planet
	s.planets[c.ID.String()] = c
}

// Adds an asteroid to the system
func (s *SolarSystem) AddAsteroid(c *Asteroid) {
	// safety check
	if c == nil {
		return
	}

	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// add asteroid
	s.asteroids[c.ID.String()] = c
}

// Adds a jumphole to the system
func (s *SolarSystem) AddJumphole(c *Jumphole) {
	// safety check
	if c == nil {
		return
	}

	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// add jumphole
	s.jumpholes[c.ID.String()] = c
}

// Adds an NPC station to the system
func (s *SolarSystem) AddStation(c *Station) {
	// safety check
	if c == nil {
		return
	}

	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// add planet
	s.stations[c.ID.String()] = c
}

// Adds a client to the system
func (s *SolarSystem) AddClient(c *shared.GameClient, lock bool) {
	// safety check
	if c == nil {
		return
	}

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// add client
	s.clients[(*c.UID).String()] = c
}

// Removes a client from the server
func (s *SolarSystem) RemoveClient(c *shared.GameClient, lock bool) {
	// safety check
	if c == nil {
		return
	}

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// remove client
	delete(s.clients, (*c.UID).String())
}

// Returns a copy of clients in the system minus their connection / event queue
func (s *SolarSystem) CopyClients(lock bool) []*shared.GameClient {
	o := make([]*shared.GameClient, 0)

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// copy clients minus connection fields
	for _, c := range s.clients {
		sid := *c.SID
		uid := *c.UID

		v := shared.GameClient{
			SID:  &sid,
			UID:  &uid,
			Conn: nil,
			ReputationSheet: shared.PlayerReputationSheet{
				FactionEntries: make(map[string]*shared.PlayerReputationSheetFactionEntry),
			},
			CurrentShipID:     c.CurrentShipID,
			CurrentSystemID:   c.CurrentSystemID,
			StartID:           c.StartID,
			EscrowContainerID: c.EscrowContainerID,
		}

		for k, f := range c.ReputationSheet.FactionEntries {
			if f != nil {
				v.ReputationSheet.FactionEntries[k] = &shared.PlayerReputationSheetFactionEntry{
					FactionID:        f.FactionID,
					StandingValue:    f.StandingValue,
					AreOpenlyHostile: f.AreOpenlyHostile,
				}
			}
		}

		o = append(o, &v)
	}

	return o
}

// Returns a copy of the ships in the system
func (s *SolarSystem) CopyShips(lock bool) map[string]*Ship {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Ship)

	// copy ships into copy map
	for k, v := range s.ships {
		c := v.CopyShip()
		copy[k] = c
	}

	// return copy map
	return copy
}

// Returns a new map containing pointers to the ships in the system - use with care!
func (s *SolarSystem) MirrorShipMap(lock bool) map[string]*Ship {

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// copy pointers into new map
	m := make(map[string]*Ship)

	for k, v := range s.ships {
		m[k] = v
	}

	// return new map
	return m
}

// Returns a copy of the stations in the system
func (s *SolarSystem) CopyStations(lock bool) map[string]*Station {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Station)

	// copy stations into copy map
	for k, v := range s.stations {
		c := v.CopyStation()
		copy[k] = &c
	}

	// return copy map
	return copy
}

// Returns a new map containing pointers to the stations in the system - use with care!
func (s *SolarSystem) MirrorStationMap(lock bool) map[string]*Station {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// copy pointers into new map
	m := make(map[string]*Station)

	for k, v := range s.stations {
		m[k] = v
	}

	// return new map
	return m
}

// Returns a copy of the jumpholes in the system
func (s *SolarSystem) CopyJumpholes(lock bool) map[string]*Jumphole {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Jumphole)

	// copy jumpholes into copy map
	for k, v := range s.jumpholes {
		c := v.CopyJumphole()
		copy[k] = &c
	}

	// return copy map
	return copy
}

// Returns a copy of the asteroids in the system
func (s *SolarSystem) CopyAsteroids(lock bool) map[string]*Asteroid {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Asteroid)

	// copy asteroids into copy map
	for k, v := range s.asteroids {
		c := v.CopyAsteroid()
		copy[k] = &c
	}

	// return copy map
	return copy
}

// Returns a copy of the stars in the system
func (s *SolarSystem) CopyStars(lock bool) map[string]*Star {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Star)

	// copy star into copy map
	for k, v := range s.stars {
		c := v.CopyStar()
		copy[k] = &c
	}

	// return copy map
	return copy
}

// Returns a copy of the planets in the system
func (s *SolarSystem) CopyPlanets(lock bool) map[string]*Planet {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// make map for copies
	copy := make(map[string]*Planet)

	// copy plant into copy map
	for k, v := range s.planets {
		c := v.CopyPlanet()
		copy[k] = &c
	}

	// return copy map
	return copy
}

// Stores an open sell order on a station
func (s *SolarSystem) StoreOpenSellOrder(order *SellOrder, lock bool) {
	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// add to map
	s.stations[order.StationID.String()].OpenSellOrders[order.ID.String()] = order
}

// Attempt to lock every entity in the system, and the system itself - never call from within the system!
func (s *SolarSystem) TestLocks() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	for _, e := range s.asteroids {
		e.Lock.Lock()
		defer e.Lock.Unlock()
	}

	for _, e := range s.jumpholes {
		e.Lock.Lock()
		defer e.Lock.Unlock()
	}

	for _, e := range s.stations {
		e.Lock.Lock()
		defer e.Lock.Unlock()
	}

	for _, e := range s.ships {
		e.Lock.Lock()
		defer e.Lock.Unlock()
	}
}

func copyModuleInfo(v FittedSlot) models.ServerModuleStatusBody {
	module := models.ServerModuleStatusBody{}

	module.Family = v.ItemTypeFamily
	module.FamilyName = v.ItemTypeFamilyName
	module.Type = v.ItemTypeName
	module.IsCycling = v.IsCycling
	module.ItemID = v.ItemID
	module.ItemTypeID = v.ItemTypeID
	module.WillRepeat = v.WillRepeat
	module.CyclePercent = v.CyclePercent
	module.Meta = models.Meta(v.ItemMeta)

	return module
}
