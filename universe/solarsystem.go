package universe

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/physics"
	"helia/shared"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Structure representing a solar system
type SolarSystem struct {
	ID                    uuid.UUID
	SystemName            string
	RegionID              uuid.UUID
	RegionName            string
	HoldingFactionID      uuid.UUID
	HoldingFactionName    string
	PosX                  float64
	PosY                  float64
	Universe              *Universe
	ships                 map[string]*Ship
	stars                 map[string]*Star
	planets               map[string]*Planet
	jumpholes             map[string]*Jumphole
	stations              map[string]*Station
	asteroids             map[string]*Asteroid
	clients               map[string]*shared.GameClient       // clients in this system
	missiles              map[string]*Missile                 // missiles in flight in this system
	wrecks                map[string]*Wreck                   // wrecks of destroyed ships in this system
	pushModuleEffects     []models.GlobalPushModuleEffectBody // module visual effect aggregation for tick
	pushPointEffects      []models.GlobalPushPointEffectBody  // non-module point visual effect aggregation for tick
	tickCounter           int                                 // counter that is used to control frequency of certain global updates
	newSystemChatMessages []models.ServerSystemChatBody
	globalAckToken        int // counter for number of ticks this system has gone through since server start (daily restart is assumed)
	Lock                  sync.Mutex
	// event escalations to engine core
	PlayerNeedRespawn    map[string]*shared.GameClient     // clients in need of a respawn by core
	NPCNeedRespawn       map[string]*Ship                  // NPCs in need of a respawn by core
	DeadShips            map[string]*Ship                  // dead ships in need of cleanup by core
	MovedItems           map[string]*Item                  // items moved to a new container in need of saving by core
	PackagedItems        map[string]*Item                  // items packaged in need of saving by core
	UnpackagedItems      map[string]*Item                  // items unpackaged in need of saving by core
	ChangedQuantityItems map[string]*Item                  // stacks of items that have changed quantity and need saving by core
	NewItems             map[string]*Item                  // stacks of items that are newly created and need to be saved by core
	NewSellOrders        map[string]*SellOrder             // new sell orders in need of saving by core
	BoughtSellOrders     map[string]*SellOrder             // sell orders that have been fulfilled in need of saving by core
	NewShipTickets       map[string]*NewShipTicket         // newly purchased/delivered ships that need to be generated and saved by core
	ShipSwitches         map[string]*ShipSwitch            // approved requests to switch a client to another ship in need of saving by core
	SetNoLoad            map[string]*ShipNoLoadSet         // updates to the no load flag in need of saving by core
	UsedShipPurchases    map[string]*UsedShipPurchase      // purchased used ships that need to be hooked in and saved by core
	ShipRenames          map[string]*ShipRename            // renamed ships that need to be saved by core
	SchematicRunViews    map[string]*shared.GameClient     // requests for a schematic run summary
	NewSchematicRuns     map[string]*NewSchematicRunTicket // newly invoked schematics that need to be started
	NewFactions          map[string]*NewFactionTicket      // partially approved requests to create a new faction and automatically join it
	LeaveFactions        map[string]*LeaveFactionTicket    // approved requests to leave a faction and rejoin the starter faction
	JoinFactions         map[string]*JoinFactionTicket     // partially approved requests to join a player into a player faction
	ViewMembers          map[string]*ViewMembersTicket     // approved requests to view player faction member list
	KickMembers          map[string]*KickMemberTicket      // partially approved requests to kick a member from a player faction
	ChangedMetaItems     map[string]*Item                  // items with changed metadata in need of saving
}

// Initializes internal aspects of SolarSystem
func (s *SolarSystem) Initialize() {
	// obtain lock
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// initialize maps
	s.clients = make(map[string]*shared.GameClient)
	s.missiles = make(map[string]*Missile)
	s.wrecks = make(map[string]*Wreck)
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
	s.NewShipTickets = make(map[string]*NewShipTicket)
	s.ShipSwitches = make(map[string]*ShipSwitch)
	s.SetNoLoad = make(map[string]*ShipNoLoadSet)
	s.UsedShipPurchases = make(map[string]*UsedShipPurchase)
	s.ShipRenames = make(map[string]*ShipRename)
	s.SchematicRunViews = make(map[string]*shared.GameClient)
	s.NewSchematicRuns = make(map[string]*NewSchematicRunTicket)
	s.NewFactions = make(map[string]*NewFactionTicket)
	s.LeaveFactions = make(map[string]*LeaveFactionTicket)
	s.JoinFactions = make(map[string]*JoinFactionTicket)
	s.ViewMembers = make(map[string]*ViewMembersTicket)
	s.KickMembers = make(map[string]*KickMemberTicket)
	s.ChangedMetaItems = make(map[string]*Item)

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
		// close underlying connection
		c.Conn.Close()

		// mark as dead
		c.Dead = true
	}
}

// Processes the solar system for a tick
func (s *SolarSystem) PeriodicUpdate() {
	// obtain lock
	s.Lock.Lock()

	// skip if nobody in system
	if len(s.ships) == 0 {
		// unlock
		s.Lock.Unlock()

		// sleep and return
		time.Sleep(1 * time.Second)
		return
	}

	// defer unlock
	defer s.Lock.Unlock()

	// process player current ship event queues
	s.processClientEventQueues()

	// update ships (both player + npc)
	s.updateShips()

	// update npc stations
	if s.tickCounter%32 == 0 {
		s.updateStations()
	}

	// update in-flight missiles
	s.updateMissiles()

	// ship collision testing
	s.shipCollisionTesting()

	// brief sleep
	time.Sleep(250 * time.Microsecond)

	// send client updates
	if s.tickCounter%2 == 0 && len(s.clients) > 0 {
		s.sendClientUpdates()
	}

	// increment tick counter
	s.tickCounter++
}

// processes the next message from each client in the system, should only be called from PeriodicUpdate
func (s *SolarSystem) processClientEventQueues() {
	// get message registry
	msgRegistry := models.SharedMessageRegistry

	for _, c := range s.clients {
		// skip if connection dead
		if c.Dead {
			continue
		}

		// pop latest event from client
		evt, lastMeaningfulActionAt := c.PopShipEvent()

		// disconnect client if more than 5 minutes since last meaningful interaction
		dMI := time.Since(lastMeaningfulActionAt)
		dMIAsMinutes := dMI.Minutes()

		if dMIAsMinutes >= 4.5 && dMIAsMinutes <= 5 {
			// calculate time remaining
			sr := (5 - dMIAsMinutes) * 60

			if s.tickCounter%42 == 0 {
				// inform client of imminent disconnection
				c.WriteErrorMessage(fmt.Sprintf("you will be disconnected due to inactivity in %v seconds", int(sr)))
			}
		} else if dMIAsMinutes > 5 {
			// inform client of disconnection
			c.WriteErrorMessage("you are being disconnected due to inactivity")

			// disconnect client
			s.RemoveClient(c, false)

			// move to next one
			continue
		}

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
		if evt.Type == msgRegistry.GlobalAck {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientGlobalAckBody)

				// verify that this ack is for this system
				if data.SolarSystemID != s.ID {
					continue
				}

				// verify the client has a token
				ct := c.GetLastGlobalAckToken()

				if ct == -1 {
					continue
				}

				// verify token isn't from the future
				if data.Token > s.globalAckToken {
					continue
				}

				// verify token isn't too old
				if data.Token <= ct {
					continue
				}

				// update token
				c.SetLastGlobalAckToken(s.globalAckToken)
			}
		} else if evt.Type == msgRegistry.NavClick {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientNavClickBody)

				// apply effect to player's current ship
				sh.CmdManualNav(data.ScreenTheta, data.ScreenMagnitude, false)
			}
		} else if evt.Type == msgRegistry.Goto {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientGotoBody)

				// apply effect to player's current ship
				sh.CmdGoto(data.TargetID, data.Type, false)
			}
		} else if evt.Type == msgRegistry.Orbit {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientOrbitBody)

				// apply effect to player's current ship
				sh.CmdOrbit(data.TargetID, data.Type, false)
			}
		} else if evt.Type == msgRegistry.Dock {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientDockBody)

				// get registry
				targetTypeReg := models.SharedTargetTypeRegistry

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
		} else if evt.Type == msgRegistry.Undock {
			if sh != nil {
				// extract data (currently nothing to process)
				// data := evt.Body.(models.ClientUndockBody)

				// make sure cargo isn't overloaded
				usedBay := sh.TotalCargoBayVolumeUsed(false)
				maxBay := sh.GetRealCargoBayVolume()

				if usedBay > maxBay {
					c.WriteErrorMessage("cargo bay overloaded")
					continue
				}

				// make sure this isn't a station warehouse
				if !sh.TemplateData.CanUndock {
					c.WriteErrorMessage("station workshops and warehouses cannot undock")
					continue
				}

				// apply effect to player's current ship
				sh.CmdUndock(false)
			}
		} else if evt.Type == msgRegistry.ActivateModule {
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
		} else if evt.Type == msgRegistry.DeactivateModule {
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
		} else if evt.Type == msgRegistry.ViewCargoBay {
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

					// skip if running schematic
					if i.SchematicInUse {
						continue
					}

					// skip if no quantity
					if i.Quantity <= 0 {
						continue
					}

					// convert for transmission
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

					if i.ItemFamilyID == "schematic" && i.Process != nil {
						pc := *i.Process

						// include process data
						sv := models.ServerSchematicViewBody{
							ID:   pc.ID,
							Time: pc.Time,
						}

						for _, pi := range pc.Inputs {
							sv.Inputs = append(sv.Inputs, models.ServerSchematicFactorViewBody{
								ItemTypeID:   pi.ItemTypeID,
								ItemTypeName: pi.ItemTypeName,
								Quantity:     pi.Quantity,
							})
						}

						for _, po := range pc.Outputs {
							sv.Outputs = append(sv.Outputs, models.ServerSchematicFactorViewBody{
								ItemTypeID:   po.ItemTypeID,
								ItemTypeName: po.ItemTypeName,
								Quantity:     po.Quantity,
							})
						}

						r.Schematic = &sv
					}

					// add to message
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
		} else if evt.Type == msgRegistry.UnfitModule {
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
		} else if evt.Type == msgRegistry.TrashItem {
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
		} else if evt.Type == msgRegistry.PackageItem {
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
		} else if evt.Type == msgRegistry.UnpackageItem {
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
		} else if evt.Type == msgRegistry.StackItem {
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
		} else if evt.Type == msgRegistry.SplitItem {
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
		} else if evt.Type == msgRegistry.FitModule {
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
		} else if evt.Type == msgRegistry.SellAsOrder {
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
		} else if evt.Type == msgRegistry.DeactivateModule {
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
		} else if evt.Type == msgRegistry.ViewOpenSellOrders {
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
		} else if evt.Type == msgRegistry.BuySellOrder {
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
		} else if evt.Type == msgRegistry.ViewIndustrialOrders {
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
		} else if evt.Type == msgRegistry.BuyFromSilo {
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
		} else if evt.Type == msgRegistry.SellToSilo {
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
		} else if evt.Type == msgRegistry.ViewStarMap {
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
		} else if evt.Type == msgRegistry.ConsumeFuel {
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
		} else if evt.Type == msgRegistry.SelfDestruct {
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
		} else if evt.Type == msgRegistry.ConsumeRepairKit {
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
		} else if evt.Type == msgRegistry.ViewProperty {
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
		} else if evt.Type == msgRegistry.Board {
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

				// verify that the integer balance of the target ship is >= 0
				if (toBoard.Wallet + 1) < 0 {
					// get defecit
					defecit := int64(math.Abs(toBoard.Wallet))

					c.WriteErrorMessage(fmt.Sprintf("you must transfer at least %v CBN to this ship before boarding to settle a station debt", defecit))
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
		} else if evt.Type == msgRegistry.TransferCredits {
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
		} else if evt.Type == msgRegistry.SellShipAsOrder {
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

				// verify no running schematics in ship
				noSchematicRuns := true

				for _, i := range toSell.CargoBay.Items {
					if i.ItemFamilyID == "schematic" {
						if i.SchematicInUse {
							noSchematicRuns = false
						}
					}
				}

				if !noSchematicRuns {
					c.WriteErrorMessage("a schematic is currently running on this ship")
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
		} else if evt.Type == msgRegistry.TrashShip {
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

				// verify no running schematics in ship
				noSchematicRuns := true

				for _, i := range toTrash.CargoBay.Items {
					if i.ItemFamilyID == "schematic" {
						if i.SchematicInUse {
							noSchematicRuns = false
						}
					}
				}

				if !noSchematicRuns {
					c.WriteErrorMessage("a schematic is currently running on this ship")
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
		} else if evt.Type == msgRegistry.RenameShip {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientRenameShipBody)

				// verify player is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to rename a ship")
					continue
				}

				// get ship to sell and verify it is owned by the player
				toRename := s.ships[data.ShipID.String()]

				if toRename == nil {
					c.WriteErrorMessage("ship not available to rename")
					continue
				}

				if toRename.UserID != sh.UserID {
					c.WriteErrorMessage("ship not available to rename")
					continue
				}

				// verify it is docked at the same station as the player
				if toRename.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked at the same station as the ship being renamed")
					continue
				}

				if toRename.DockedAtStation.ID != sh.DockedAtStation.ID {
					c.WriteErrorMessage("you must be docked at the same station as the ship being renamed")
					continue
				}

				// verify length constraint on new name
				if len(data.Name) > 32 {
					c.WriteErrorMessage("ship names must be 32 characters or less")
					continue
				}

				// update name in memory
				toRename.ShipName = data.Name

				// escalate rename save request
				rn := ShipRename{
					ShipID: toRename.ID,
					Name:   data.Name,
				}

				s.ShipRenames[toRename.ID.String()] = &rn

				// update renamed ship in property cache (so it changes immediately instead of as part of the periodic rebuild)
				pc := c.GetPropertyCache()
				no := make([]shared.ShipPropertyCacheEntry, 0)

				for _, e := range pc.ShipCaches {
					if e.ShipID == toRename.ID {
						e.Name = data.Name
					}

					no = append(no, e)
				}

				pc.ShipCaches = no
				c.SetPropertyCache(pc)
			}
		} else if evt.Type == msgRegistry.PostSystemChatMessage {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientPostSystemChatMessageBody)

				// throttle posting rate
				lp := c.LastChatPostedAt
				ct := time.Now()

				dt := ct.Sub(lp)

				if dt.Seconds() < 2 {
					c.WriteErrorMessage("you need to wait to post again")

					// reset timestamp to deter spam attempts at posting
					c.LastChatPostedAt = time.Now()
					continue
				}

				// verify message constraints
				if len(data.Message) > 57 {
					c.WriteErrorMessage("chat messages must be 57 characters or less")
					continue
				}

				if len(data.Message) == 0 {
					c.WriteErrorMessage("message must have content")
					continue
				}

				// store message
				s.newSystemChatMessages = append(s.newSystemChatMessages, models.ServerSystemChatBody{
					SenderID:   sh.UserID,
					SenderName: sh.CharacterName,
					Message:    data.Message,
				})

				// update timestamp
				c.LastChatPostedAt = time.Now()
			}
		} else if evt.Type == msgRegistry.TransferItem {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientTransferItemBody)

				// find item in source
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// find receiver
					receiver := s.ships[data.ReceiverID.String()]

					if receiver == nil {
						// do nothing
						continue
					} else {
						// verify both ships are docked at the same station
						if sh.DockedAtStationID == nil || receiver.DockedAtStationID == nil {
							c.WriteErrorMessage("both ships must be docked at the same station to transfer an item")
							continue
						}

						if *sh.DockedAtStationID != *receiver.DockedAtStationID {
							c.WriteErrorMessage("both ships must be docked at the same station to transfer an item")
							continue
						}

						// verify both ships are owned by the same player
						if sh.UserID != receiver.UserID {
							c.WriteErrorMessage("destination not available for item transfer")
							continue
						}

						// make sure this isn't the same ship
						if sh.ID == receiver.ID {
							c.WriteErrorMessage("item already there")
							continue
						}

						// make sure receiver isn't in debt
						if receiver.Wallet < 0 {
							c.WriteErrorMessage("you cannot transfer items to a ship in debt with the station manager")
							continue
						}

						// pull item from source ship
						item, err := sh.RemoveItemFromCargo(data.ItemID, false)

						if item == nil || err != nil {
							c.WriteErrorMessage("unable to complete transfer")
							continue
						}

						// push item to receiver ship
						err = receiver.PutItemInCargo(item, false)

						if err != nil {
							// put item back in source ship
							sh.PutItemInCargo(item, false)

							// write error to client
							c.WriteErrorMessage("unable to complete transfer")
							continue
						}

						// escalate to core for saving
						item.ContainerID = receiver.CargoBayContainerID
						item.CoreDirty = true
						s.MovedItems[item.ID.String()] = item
					}
				}
			}
		} else if evt.Type == msgRegistry.ViewExperience {
			if sh != nil {
				// extract data
				// data := evt.Body.(models.ClientViewExperienceBody)

				// build experience update
				u := c.ExperienceSheet.CopyAsUpdate()

				// package message
				b, _ := json.Marshal(&u)

				z := models.GameMessage{
					MessageType: msgRegistry.ExperienceUpdate,
					MessageBody: string(b),
				}

				// write response to client
				c.WriteMessage(&z)
			}
		} else if evt.Type == msgRegistry.ViewSchematicRuns {
			if sh != nil {
				// extract data
				// data := evt.Body.(models.ClientViewSchematicRunsBody)

				// escalate to core
				s.SchematicRunViews[c.UID.String()] = c
			}
		} else if evt.Type == msgRegistry.RunSchematic {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientRunSchematicBody)

				// find item in cargo
				item := sh.FindItemInCargo(data.ItemID)

				// make sure we found something
				if item == nil {
					// do nothing
					continue
				} else {
					// verify ship is docked
					if sh.DockedAtStation == nil {
						c.WriteErrorMessage("you must be docked to run a schematic")
						continue
					}

					// verify ship is a station warehouse or workshop
					if sh.TemplateData.CanUndock {
						c.WriteErrorMessage("schematics can only be run from a station workshop or warehouse")
						continue
					}

					// verify item is a schematic
					if item.ItemFamilyID != "schematic" {
						c.WriteErrorMessage("this item is not a schematic")
						continue
					}

					// verify schematic is available
					if item.SchematicInUse {
						c.WriteErrorMessage("this schematic is already running")
						continue
					}

					// verify schematic is clean
					if item.CoreDirty {
						c.WriteErrorMessage("schematic is dirty")
						continue
					}

					// verify process is linked
					if item.Process == nil {
						c.WriteErrorMessage("schematic is improperly initialized")
						continue
					}

					// verify that all input requirements are met
					inputsMet := true
					inputsUsed := make([]*Item, 0)
					inputsSize := make([]int, 0)

					for _, i := range item.Process.Inputs {
						// look for sufficient stack
						so := sh.FindFirstAvailablePackagedStackOfSizeInCargo(i.ItemTypeID, i.Quantity)

						if so == nil {
							inputsMet = false
						} else {
							// store pointer and size
							inputsUsed = append(inputsUsed, so)
							inputsSize = append(inputsSize, i.Quantity)
						}
					}

					if !inputsMet {
						c.WriteErrorMessage("not all required resources were provided")
						continue
					}

					// verify there is sufficient space to store outputs
					outputSize := 0.0

					for _, o := range item.Process.Outputs {
						// get stack volume
						v, _ := o.ItemTypeMeta.GetFloat64("volume")
						sv := float64(o.Quantity) * v

						// accumulate
						outputSize += sv
					}

					bayUsed := sh.TotalCargoBayVolumeUsed(false)
					bayMax := sh.GetRealCargoBayVolume()

					if bayUsed+outputSize > bayMax {
						c.WriteErrorMessage("insufficient space to store outputs")
						continue
					}

					// consume input resources
					for x, i := range inputsUsed {
						// decrement quantity
						i.Quantity -= inputsSize[x]

						// escalate to core for saving
						s.ChangedQuantityItems[i.ID.String()] = i
						i.CoreDirty = true
					}

					// start job
					t := NewSchematicRunTicket{
						Client:        c,
						Ship:          sh,
						SchematicItem: item,
					}

					s.NewSchematicRuns[item.ID.String()] = &t
					item.SchematicInUse = true
				}
			}
		} else if evt.Type == msgRegistry.CreateNewFaction {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientCreateNewFactionBody)

				// verify ship is docked
				if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
					c.WriteErrorMessage("you must be docked to create a faction")
					continue
				}

				if sh.Faction != nil {
					// verify player is in an NPC faction
					if !sh.Faction.IsNPC {
						c.WriteErrorMessage("you must be in an NPC faction to create a faction")
						continue
					}

					// required fields
					if len(data.Name) == 0 {
						c.WriteErrorMessage("you must enter a faction name")
						continue
					}

					if len(data.Description) == 0 {
						c.WriteErrorMessage("you must enter a faction description")
						continue
					}

					if len(data.Ticker) == 0 {
						c.WriteErrorMessage("you must enter a faction ticker")
						continue
					}

					// enforce string length limits
					if len(data.Name) > 24 {
						c.WriteErrorMessage("faction name must be 24 characters or less")
						continue
					}

					if len(data.Description) > 8192 {
						c.WriteErrorMessage("please enter a shorter description")
						continue
					}

					if len(data.Ticker) > 3 {
						c.WriteErrorMessage("ticker must be 3 characters or less")
						continue
					}

					// escalate to core for creation
					t := NewFactionTicket{
						Name:          data.Name,
						Description:   data.Description,
						Ticker:        data.Ticker,
						Client:        c,
						HomeStationID: *sh.DockedAtStationID,
					}

					s.NewFactions[c.UID.String()] = &t
				}
			}
		} else if evt.Type == msgRegistry.LeaveFaction {
			if sh != nil {
				// extract data
				//data := evt.Body.(models.ClientLeaveFactionBody)

				// verify ship is docked
				if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
					c.WriteErrorMessage("you must be docked to leave a faction")
					continue
				}

				if sh.Faction != nil {
					// check if armed
					if !sh.LeaveFactionArmed {
						// arm
						sh.LeaveFactionArmed = true

						c.WriteErrorMessage("issue leave faction command again to return to your starter faction")
						continue
					}

					// escalate to core for leave
					t := LeaveFactionTicket{
						Client: c,
					}

					s.LeaveFactions[c.UID.String()] = &t
				}
			}
		} else if evt.Type == msgRegistry.ApplyToFaction {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientApplyToFactionBody)

				// null checks
				if sh.Faction == nil {
					continue
				}

				if c.UID == nil {
					continue
				}

				// verify ship is docked
				if sh.DockedAtStation == nil || sh.DockedAtStationID == nil {
					c.WriteErrorMessage("you must be docked to apply to join a faction")
					continue
				}

				// verify player is in an NPC faction
				if !sh.Faction.IsNPC {
					c.WriteErrorMessage("you must be in an NPC faction to apply to join a faction")
					continue
				}

				// get faction being applied to
				f := s.Universe.Factions[data.FactionID.String()]

				if f == nil {
					c.WriteErrorMessage("faction does not exist")
					continue
				}

				// verify this is a player faction
				if f.IsNPC {
					c.WriteErrorMessage("this is an NPC faction")
					continue
				}

				// verify it is accepting applications
				if !f.IsJoinable {
					c.WriteErrorMessage("faction is closed to applications")
					continue
				}

				// add application on separate goroutine
				cf := sh.Faction

				go func(f *Faction, cf *Faction, c *shared.GameClient) {
					// obtain lock
					f.Lock.Lock()
					defer f.Lock.Unlock()

					// add ticket
					f.Applications[c.UID.String()] = FactionApplicationTicket{
						UserID:         *c.UID,
						CurrentFaction: cf,
						CharacterName:  sh.CharacterName,
					}

					// notify client
					c.WriteInfoMessage(fmt.Sprintf("Application to join %v submitted!", f.Name))
				}(f, cf, c)
			}
		} else if evt.Type == msgRegistry.ViewApplications {
			// extract data
			//data := evt.Body.(models.ClientViewApplicationsBody)

			// null check
			if sh.Faction == nil {
				continue
			}

			// verify faction has an owner
			if sh.Faction.OwnerID == nil {
				continue
			}

			// verify faction is player controlled
			if sh.Faction.IsNPC {
				continue
			}

			// verify client is faction owner
			oID := *sh.Faction.OwnerID
			cID := *c.UID

			if oID != cID {
				c.WriteErrorMessage("you are not the owner of this faction")
				continue
			}

			// build and send update on separate goroutine
			go func(c *shared.GameClient, f *Faction) {
				// obtain lock
				f.Lock.Lock()
				defer f.Lock.Unlock()

				// build update
				um := models.ServerApplicationsUpdateBody{
					Applications: make([]models.ServerApplicationEntry, 0),
				}

				for _, v := range f.Applications {
					um.Applications = append(um.Applications, models.ServerApplicationEntry{
						UserID:        v.UserID,
						CharacterName: v.CharacterName,
						FactionName:   v.CurrentFaction.Name,
						FactionTicker: v.CurrentFaction.Ticker,
					})
				}

				// serialize update
				b, _ := json.Marshal(&um)

				msg := models.GameMessage{
					MessageType: msgRegistry.ApplicationsUpdate,
					MessageBody: string(b),
				}

				// write update to client
				c.WriteMessage(&msg)
			}(c, sh.Faction)
		} else if evt.Type == msgRegistry.ApproveApplication {
			// extract data
			data := evt.Body.(models.ClientApproveApplicationBody)

			// null check
			if sh.Faction == nil {
				continue
			}

			// verify faction has an owner
			if sh.Faction.OwnerID == nil {
				continue
			}

			// verify faction is player controlled
			if sh.Faction.IsNPC {
				continue
			}

			// verify client is faction owner
			oID := *sh.Faction.OwnerID
			cID := *c.UID

			if oID != cID {
				c.WriteErrorMessage("you are not the owner of this faction")
				continue
			}

			// escalate to core for final approval and handling
			s.JoinFactions[cID.String()] = &JoinFactionTicket{
				UserID:      data.UserID,
				FactionID:   sh.FactionID,
				OwnerClient: c,
			}

			// remove application on separate goroutine
			go func(f *Faction, userID uuid.UUID) {
				// obtain lock
				f.Lock.Lock()
				defer f.Lock.Unlock()

				// remove entry
				delete(f.Applications, userID.String())
			}(sh.Faction, data.UserID)
		} else if evt.Type == msgRegistry.RejectApplication {
			// extract data
			data := evt.Body.(models.ClientRejectApplicationBody)

			// null check
			if sh.Faction == nil {
				continue
			}

			// verify faction has an owner
			if sh.Faction.OwnerID == nil {
				continue
			}

			// verify faction is player controlled
			if sh.Faction.IsNPC {
				continue
			}

			// verify client is faction owner
			oID := *sh.Faction.OwnerID
			cID := *c.UID

			if oID != cID {
				c.WriteErrorMessage("you are not the owner of this faction")
				continue
			}

			// remove application on separate goroutine
			go func(f *Faction, userID uuid.UUID) {
				// obtain lock
				f.Lock.Lock()
				defer f.Lock.Unlock()

				// remove entry
				delete(f.Applications, userID.String())
			}(sh.Faction, data.UserID)
		} else if evt.Type == msgRegistry.ViewMembers {
			// extract data
			//data := evt.Body.(models.ClientViewMembersBody)

			// null check
			if sh.Faction == nil {
				continue
			}

			// verify faction has an owner
			if sh.Faction.OwnerID == nil {
				continue
			}

			// verify faction is player controlled
			if sh.Faction.IsNPC {
				continue
			}

			// verify client is faction owner
			oID := *sh.Faction.OwnerID
			cID := *c.UID

			if oID != cID {
				c.WriteErrorMessage("you are not the owner of this faction")
				continue
			}

			// escalate to core
			s.ViewMembers[oID.String()] = &ViewMembersTicket{
				FactionID:   sh.FactionID,
				OwnerClient: c,
			}
		} else if evt.Type == msgRegistry.KickMember {
			// extract data
			data := evt.Body.(models.ClientKickMemberBody)

			// null check
			if sh.Faction == nil {
				continue
			}

			// verify faction has an owner
			if sh.Faction.OwnerID == nil {
				continue
			}

			// verify faction is player controlled
			if sh.Faction.IsNPC {
				continue
			}

			// verify client is faction owner
			oID := *sh.Faction.OwnerID
			cID := *c.UID

			if oID != cID {
				c.WriteErrorMessage("you are not the owner of this faction")
				continue
			}

			// verify client is not the one being kicked
			if data.UserID == cID {
				c.WriteErrorMessage("you can't kick yourself")
				continue
			}

			// escalate to core for final validation and handling
			s.KickMembers[cID.String()] = &KickMemberTicket{
				UserID:      data.UserID,
				FactionID:   sh.FactionID,
				OwnerClient: c,
			}
		} else if evt.Type == msgRegistry.UseModKit {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientUseModKitBody)

				// verify ship is docked
				if sh.DockedAtStation == nil {
					c.WriteErrorMessage("you must be docked to apply a mod kit")
					continue
				}

				// find mod kit
				modKit := sh.FindItemInCargo(data.ModKitItemID)

				if modKit == nil || modKit.ItemFamilyID != "mod_kit" {
					c.WriteErrorMessage("unable to find mod kit")
					continue
				}

				// verify mod kit is unpackaged
				if modKit.IsPackaged || modKit.Quantity != 1 {
					c.WriteErrorMessage("you must unpackage a mod kit before use")
					continue
				}

				// find module to mod
				module := sh.FindItemInCargo(data.ModuleItemID)

				if module == nil {
					c.WriteErrorMessage("unable to find module")
					continue
				}

				// verify its a module
				modFamily := getModuleFamily(module.ItemFamilyID)

				if modFamily == "" {
					c.WriteErrorMessage("target is not a module")
					continue
				}

				// verify the module is unpackaged
				if module.IsPackaged || module.Quantity != 1 {
					c.WriteErrorMessage("modules must be unpackaged before applying a mod kit")
					continue
				}

				// consume mod kit
				modKit.Quantity = 0
				modKit.CoreDirty = true

				// escalate to core for saving
				s.ChangedQuantityItems[modKit.ID.String()] = modKit

				// get mod kit attributes
				damageChance, _ := modKit.Meta.GetFloat64("damage_chance")
				successChance, _ := modKit.Meta.GetFloat64("success_chance")
				maxAttributesAffected, _ := modKit.Meta.GetInt("max_attributes_affected")
				maxMutation, _ := modKit.Meta.GetFloat64("max_mutation")

				// iterate over item attributes
				attributesChanged := 0
				changedAttributes := make(map[string]float64)
				damage := 0.0

				for k := range module.Meta {
					// check if limit reached
					if attributesChanged >= maxAttributesAffected {
						break
					}

					// check if this is mutable
					mutable := itemMetaIsMutable(k)

					if !mutable {
						continue
					}

					// get float64 value (all mutable attributes must be of this type)
					attrVal, f := module.Meta.GetFloat64(k)

					if !f {
						continue
					}

					// do success roll
					successRoll := rand.Float64()

					if successRoll <= successChance {
						// roll for mutation amount
						mutationRoll := rand.Float64() * maxMutation

						// roll for mutation direction
						directionRoll := rand.Float64()
						direction := 0

						if directionRoll <= 0.5 {
							direction = 1
						} else {
							direction = -1
						}

						mutationRoll *= float64(direction)

						// apply to attribute
						scaleFactor := 1.0 + mutationRoll
						attrVal *= scaleFactor

						// store as changed
						changedAttributes[k] = attrVal

						// increment counter
						attributesChanged++
					}

					// do damage roll
					damageRoll := rand.Float64()

					if damageRoll <= damageChance {
						// roll for damage amount
						mutationRoll := rand.Float64() * maxMutation
						damage += mutationRoll
					}
				}

				// apply damage
				if damage > 0 {
					// get hp
					hp, _ := module.Meta.GetFloat64("hp")

					// apply damage
					hp *= 1.0 - damage

					// update damage
					module.Meta["hp"] = int(hp)

					// check if broken
					if int(hp) <= 1 {
						// destroy module
						module.CoreDirty = true
						module.Quantity = 0

						s.ChangedQuantityItems[module.ID.String()] = module
					}
				}

				// save new attributes
				if attributesChanged > 0 {
					// update metadata
					for k, v := range changedAttributes {
						module.Meta[k] = v
					}

					// increment modifications counter
					mf, _ := module.Meta.GetInt("customization_factor")
					mf += attributesChanged

					module.Meta["customization_factor"] = int(mf)

					// flag as modified
					module.Meta["**MODIFIED**"] = true

					// save module
					module.CoreDirty = true
					s.ChangedMetaItems[module.ID.String()] = module
				}
			}
		} else if evt.Type == msgRegistry.ViewActionReportsPage {
			if sh != nil {
				// extract data
				data := evt.Body.(models.ClientViewActionReportsPageBody)

				// todo
				shared.TeeLog(fmt.Sprintf("todo: view action reports page %v", data))
			}
		}
	}
}

// Updates the state of ships in the solar system, should only be called from PeriodicUpdate
func (s *SolarSystem) updateShips() {
	// update ships
	for _, e := range s.ships {
		// check if dead
		if e.Destroyed {
			// skip
			continue
		}

		if e.DockedAtStationID == nil {
			// update player flight experience
			if e.BeingFlownByPlayer && e.ExperienceSheet != nil {
				// make sure they still have a connected client
				c := s.clients[e.UserID.String()]

				if c != nil {
					// get experience entry for this ship template
					x := e.ExperienceSheet.GetShipExperienceEntry(e.TemplateData.ID)

					// stash level
					xl := math.Trunc(x.GetExperience())

					// update entry
					x.SecondsOfExperience += (Heartbeat / 1000.0)
					x.ShipTemplateName = e.TemplateData.ShipTemplateName
					e.ExperienceSheet.SetShipExperienceEntry(x)

					// compare new and old levels
					nl := math.Trunc(x.GetExperience())

					if nl > xl {
						// notify player of the level up!
						c.WriteInfoMessage(
							fmt.Sprintf(
								"you have advanced to %v level %v! be sure to fuel and repair them to take advantage of their increased abilities!",
								x.ShipTemplateName,
								nl,
							),
						)
					}
				}
			}
		}

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
				GfxEffect: e.TemplateData.ExplosionTexture,
				PosX:      e.PosX,
				PosY:      e.PosY,
				Radius:    e.TemplateData.Radius * 1.5,
			}

			s.pushPointEffects = append(s.pushPointEffects, exp)

			// create wreck for ship
			wr := Wreck{
				ID:            uuid.New(),
				SystemID:      s.ID,
				PosX:          e.PosX,
				PosY:          e.PosY,
				Theta:         e.Theta,
				Radius:        e.TemplateData.Radius,
				Texture:       e.TemplateData.WreckTexture,
				WreckName:     fmt.Sprintf("Wreck of %v", e.ShipName),
				DeadShipItems: make([]*Item, 0),
				DeadShip:      e,
			}

			for _, it := range e.CargoBay.Items {
				if it.Quantity <= 0 {
					continue
				}

				// drop chance roll
				roll := rand.Float64()

				if roll < 0.5 {
					continue
				}

				wr.DeadShipItems = append(wr.DeadShipItems, it)
			}

			for _, it := range e.FittingBay.Items {
				if it.Quantity <= 0 {
					continue
				}

				// drop chance roll
				roll := rand.Float64()

				if roll < 0.75 {
					continue
				}

				wr.DeadShipItems = append(wr.DeadShipItems, it)
			}

			s.wrecks[wr.ID.String()] = &wr

			// log destruction to console
			bm := 0

			if e.BehaviourMode != nil {
				bm = *e.BehaviourMode
			}

			shared.TeeLog(
				fmt.Sprintf(
					"[%v] %v was destroyed (%v::%v>>%v)",
					s.SystemName,
					e.CharacterName,
					e.Texture,
					bm,
					e.PlayerAggressors,
				),
			)
		} else {
			// update ship
			e.PeriodicUpdate()
		}
	}
}

// updates the state of station in the system, should only be called from PeriodicUpdate
func (s *SolarSystem) updateStations() {
	for _, e := range s.stations {
		e.PeriodicUpdate()
	}
}

// updates the state of missiles in the system, should only be called from PeriodicUpdate
func (s *SolarSystem) updateMissiles() {
	// get target type registry
	tgtTypeReg := models.SharedTargetTypeRegistry

	// missile collission testing
	dropMissiles := make([]string, 0)

	for _, mA := range s.missiles {
		if mA.TargetType == tgtTypeReg.Ship {
			// get target ship
			sB := s.ships[mA.TargetID.String()]

			if sB != nil {
				// get physics dummies
				dummyA := mA.ToPhysicsDummy()
				dummyB := sB.ToPhysicsDummy()

				// get distance between ships
				d := physics.Distance(dummyA, dummyB)

				// check for radius intersection
				if d <= sB.TemplateData.Radius {
					m := mA.Module

					// get damage values
					shieldDmg, _ := m.ItemMeta.GetFloat64("shield_damage")
					armorDmg, _ := m.ItemMeta.GetFloat64("armor_damage")
					hullDmg, _ := m.ItemMeta.GetFloat64("hull_damage")

					// apply damage to ship
					sB.DealDamage(
						shieldDmg,
						armorDmg,
						hullDmg,
						m.shipMountedOn.ReputationSheet,
						m,
					)

					// schedule missile removal
					dropMissiles = append(dropMissiles, mA.ID.String())
				}
			}
		} else if mA.TargetType == tgtTypeReg.Station {
			// get target station
			sB := s.stations[mA.TargetID.String()]

			if sB != nil {
				// get physics dummies
				dummyA := mA.ToPhysicsDummy()
				dummyB := sB.ToPhysicsDummy()

				// get distance between ships
				d := physics.Distance(dummyA, dummyB)

				// check for half radius intersection
				if d <= sB.Radius/2 {
					m := mA.Module

					// get damage values
					shieldDmg, _ := m.ItemMeta.GetFloat64("shield_damage")
					armorDmg, _ := m.ItemMeta.GetFloat64("armor_damage")
					hullDmg, _ := m.ItemMeta.GetFloat64("hull_damage")

					// apply damage to station
					sB.DealDamage(shieldDmg, armorDmg, hullDmg)

					// schedule missile removal
					dropMissiles = append(dropMissiles, mA.ID.String())
				}
			}
		}
	}

	// update missiles
	for _, m := range s.missiles {
		if m.TicksRemaining <= 0 {
			// schedule missile removal
			dropMissiles = append(dropMissiles, m.ID.String())
		} else {
			m.PeriodicUpdate()
		}
	}

	// remove dropped missiles
	for _, k := range dropMissiles {
		// get missile and backing module
		mA := s.missiles[k]
		m := mA.Module

		// drop explosion for missile
		expEffect, _ := m.ItemMeta.GetString("missile_explosion_effect")
		expRad, _ := m.ItemMeta.GetFloat64("missile_explosion_radius")

		exp := models.GlobalPushPointEffectBody{
			GfxEffect: expEffect,
			PosX:      mA.PosX,
			PosY:      mA.PosY,
			Radius:    expRad,
		}

		s.pushPointEffects = append(s.pushPointEffects, exp)

		// remove from map
		delete(s.missiles, k)
	}
}

// tests for and applies the effects of collisions in the system, should only be called from PeriodicUpdate
func (s *SolarSystem) shipCollisionTesting() {
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
				sysRad := (sA.TemplateData.Radius + sB.TemplateData.Radius)

				if d <= sysRad {
					// calculate collission results
					mTa, mTb := physics.ElasticCollide(&dummyA, &dummyB, sysRad)

					// update ships with results
					sA.ApplyPhysicsDummy(dummyA)
					sB.ApplyPhysicsDummy(dummyB)

					// determine angle sign
					mS := 1.0

					if rand.Float64() <= 0.5 {
						mS = -1.0
					}

					// apply mixing angles
					sA.Theta += mTa * mS
					sB.Theta -= mTb * mS
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

							// defer to avoid deadlocking either system
							defer func(s *SolarSystem) {
								go func(s *SolarSystem) {
									defer gjB.OutSystem.AddClient(gc, true)
									defer s.RemoveClient(gc, true)
								}(s)
							}(s)
						}
					}

					// kill ship autopilot
					defer gsA.CmdAbort(false)

					// place ship on the opposite side of the hole
					riX := gjB.PosX - gsA.PosX
					riY := gjB.PosY - gsA.PosY

					gsA.PosX = (gjB.OutJumphole.PosX + (riX * 1.25))
					gsA.PosY = (gjB.OutJumphole.PosY + (riY * 1.25))

					// in the destination system
					gsA.SystemID = gjB.OutSystemID

					// perform move operation
					defer func(s *SolarSystem) {
						go func(s *SolarSystem) {
							defer gjB.OutSystem.AddShip(gsA, true)
						}(s)
					}(s)

					defer s.RemoveShip(gsA, false)
				}(sA, jB, c)

				break
			}
		}
	}
}

// sends routine updates to clients in the system, should only be called from PeriodicUpdate
func (s *SolarSystem) sendClientUpdates() {
	// initialize or decay tokens
	for _, c := range s.clients {
		lt := c.GetLastGlobalAckToken()

		if lt == -1 {
			c.SetLastGlobalAckToken(s.globalAckToken)
		} else {
			c.SetLastGlobalAckToken(lt - 1)
		}
	}

	// increment global update counter
	s.globalAckToken++

	// check tick counter to determine whether to send static world data
	sendStatic := s.tickCounter > 200

	// check tick counter to determine whether to send secret updates
	sendSecret := s.tickCounter%12 == 0

	// check tick counter to determine whether to send player rep sheets
	sendPlayerRepSheets := s.tickCounter%128 == 0

	if sendStatic {
		// reset tick counter
		s.tickCounter = 0
	}

	// get message registry
	msgRegistry := models.SharedMessageRegistry

	// build global update of non-secret info for clients
	gu := models.ServerGlobalUpdateBody{
		SystemChat: s.newSystemChatMessages,
		Token:      s.globalAckToken,
	}

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

		// skip cloaked ships
		if d.IsCloaked {
			continue
		}

		// build ship info and append
		gu.Ships = append(gu.Ships, models.GlobalShipInfo{
			ID:            d.ID,
			UserID:        d.UserID,
			Created:       d.Created,
			ShipName:      d.ShipName,
			CharacterName: d.CharacterName,
			PosX:          d.PosX,
			PosY:          d.PosY,
			SystemID:      d.SystemID,
			Texture:       d.Texture,
			Theta:         d.Theta,
			VelX:          d.VelX,
			VelY:          d.VelY,
			Mass:          d.GetRealMass(),
			Radius:        d.TemplateData.Radius,
			ShieldP:       ((d.Shield / d.GetRealMaxShield()) * 100) + Epsilon,
			ArmorP:        ((d.Armor / d.GetRealMaxArmor()) * 100) + Epsilon,
			HullP:         ((d.Hull / d.GetRealMaxHull()) * 100) + Epsilon,
			FactionID:     d.FactionID,
		})
	}

	for _, d := range s.missiles {
		if d.TicksRemaining <= 0 {
			continue
		}

		gu.Missiles = append(gu.Missiles, models.GlobalMissileBody{
			ID:      d.ID,
			PosX:    d.PosX,
			PosY:    d.PosY,
			Radius:  d.Radius,
			Texture: d.Texture,
		})
	}

	emptyWreckIDs := make([]string, 0)

	for k, r := range s.wrecks {
		// skip empty wrecks and mark for deletion
		if len(r.DeadShipItems) == 0 {
			emptyWreckIDs = append(emptyWreckIDs, k)
			continue
		}

		gu.Wrecks = append(gu.Wrecks, models.GlobalWreckInfo{
			ID:        r.ID,
			SystemID:  r.SystemID,
			WreckName: r.WreckName,
			PosX:      r.PosX,
			PosY:      r.PosY,
			Texture:   r.Texture,
			Radius:    r.Radius,
			Theta:     r.Theta,
		})
	}

	// delete empty wrecks
	for _, i := range emptyWreckIDs {
		delete(s.wrecks, i)
	}

	if sendStatic {
		/*
		 * Performance note: This data is very static and can be sent rarely. As long as
		 * a client is guaranteed to get it within a few seconds of entering a system, it
		 * should be fine. Sending this frequently wastes an enormous amount of bandwidth.
		 * Clients will only be sent static data if they need it, which would be when they
		 * enter the system.
		 */

		// stars
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

		// planets
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

		// asteroids
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

		// jumpholes
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

		// stations
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

	// map to store clients that received global update
	receivedGlobal := make(map[string]*shared.GameClient)

	// write global update to clients
	for _, c := range s.clients {
		/*
		 * Performance note: An attempt to send this message to the client
		 * must be made every tick, or gameplay will be choppy from their
		 * perspective.
		 *
		 * Note that if a client is not acknowledging receipt of global
		 * updates promptly that the rate at which those updates will be
		 * sent will be throttled. This is to save bandwidth and avoid
		 * saturating clients with more updates than they can handle.
		 */

		// skip if connection dead
		if c.Dead {
			s.RemoveClient(c, false)
			continue
		}

		// skip if already has static and sending static data
		if sendStatic && c.HasStatic {
			continue
		}

		// get token
		ct := c.GetLastGlobalAckToken()

		// get difference between tokens
		dt := s.globalAckToken - ct

		// verify client is reasonably up to date
		up := (float64(dt) / 50.0) * 100
		ur := physics.RandInRange(0, 100)

		if ur >= int(up) {
			// send global update
			c.WriteMessage(&msg)

			if sendStatic {
				go func(c *shared.GameClient) {
					// obtain lock
					c.Lock.Lock()
					defer c.Lock.Unlock()

					// mark as having static
					c.HasStatic = true
				}(c)
			}

			// allow additional updates
			receivedGlobal[c.UID.String()] = c
		} else if up > 100 {
			// bring token partially up to date
			c.SetLastGlobalAckToken(s.globalAckToken - dt/2)
		}
	}

	// write secret current ship updates to individual clients
	if sendSecret {
		/*
		 * Performance note: This message is sent less often than the global update because the data
		 * contained doesn't need to be as fresh for the game to feel smooth and responsive, and
		 * sending it just a little less often significantly reduced bandwidth usage.
		 */

		for _, c := range receivedGlobal {
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
				module.HardpointPosition = slot.TexturePosition

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

		for _, c := range receivedGlobal {
			// skip if connection dead
			if c.Dead {
				continue
			}

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

	// reset system chat messages for next tick
	s.newSystemChatMessages = make([]models.ServerSystemChatBody, 0)
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

	// store docked station pointer if docked
	if c.DockedAtStationID != nil {
		c.DockedAtStation = s.stations[c.DockedAtStationID.String()]

		if c.DockedAtStation != nil {
			c.IsDocked = true
		} else {
			c.DockedAtStationID = nil
			c.IsDocked = false
		}
	} else {
		c.IsDocked = false
	}

	// disarm self destruct
	c.DestructArmed = false

	// disarm leave faction
	c.LeaveFactionArmed = false

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

	// get target type registry
	tgtRegistry := models.SharedTargetTypeRegistry

	// remove missiles tracking or fired by ship
	dropMissiles := make([]string, 0)

	for k, m := range s.missiles {
		if m.TargetType == tgtRegistry.Ship && m.TargetID == c.ID {
			dropMissiles = append(dropMissiles, k)
		} else if m.FiredByID == c.ID {
			dropMissiles = append(dropMissiles, k)
		}
	}

	for _, k := range dropMissiles {
		delete(s.missiles, k)
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

	// clear token
	c.ClearLastGlobalAckToken()

	// ensure client gets new static data
	go func(c *shared.GameClient) {
		for i := 0; i < 3; i++ {
			// obtain lock
			c.Lock.Lock()

			// mark as needing static data
			c.HasStatic = false

			// unlock
			c.Lock.Unlock()

			// short sleep
			time.Sleep(100 * time.Millisecond)
		}
	}(c)

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

	// mark as needing static data
	c.HasStatic = false

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
				Lock:           sync.Mutex{},
			},
			CurrentShipID:     c.CurrentShipID,
			CurrentSystemID:   c.CurrentSystemID,
			StartID:           c.StartID,
			EscrowContainerID: c.EscrowContainerID,
			HasStatic:         c.HasStatic,
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
		c := v.CopyShip(lock)
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

// Returns a new map containing pointers to the clients in the system - use with care!
func (s *SolarSystem) MirrorClientMap(lock bool) map[string]*shared.GameClient {

	if lock {
		// obtain lock
		s.Lock.Lock()
		defer s.Lock.Unlock()
	}

	// copy pointers into new map
	m := make(map[string]*shared.GameClient)

	for k, v := range s.clients {
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
