package engine

import (
	"errors"
	"fmt"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"log"

	"github.com/google/uuid"
)

// Loads the state of the universe from the database
func LoadUniverse() (*universe.Universe, error) {
	// get services
	regionSvc := sql.GetRegionService()
	systemSvc := sql.GetSolarSystemService()
	shipSvc := sql.GetShipService()
	starSvc := sql.GetStarService()
	planetSvc := sql.GetPlanetService()
	stationSvc := sql.GetStationService()
	stationProcessSvc := sql.GetStationProcessService()
	jumpholeSvc := sql.GetJumpholeService()
	asteroidSvc := sql.GetAsteroidService()
	itemTypeSvc := sql.GetItemTypeService()
	itemFamilySvc := sql.GetItemFamilyService()
	sellOrderSvc := sql.GetSellOrderService()
	itemSvc := sql.GetItemService()
	factionSvc := sql.GetFactionService()

	// empty universe to fill
	u := universe.Universe{}

	// load factions
	dfs, err := factionSvc.GetAllFactions()

	if err != nil {
		return nil, err
	}

	factions := make(map[string]*universe.Faction)

	for _, f := range dfs {
		uf := FactionFromSQL(&f)

		factions[f.ID.String()] = uf
	}

	u.Factions = factions

	// for linking jumpholes later
	jhMap := make(map[string]*universe.Jumphole)
	sMap := make(map[string]*universe.SolarSystem)

	// load regions
	rs, err := regionSvc.GetAllRegions()

	if err != nil {
		return nil, err
	}

	regions := make(map[string]*universe.Region)
	for _, e := range rs {
		// load basic region information
		r := universe.Region{
			ID:         e.ID,
			RegionName: e.RegionName,
			PosX:       e.PosX,
			PosY:       e.PosY,
		}

		// load systems in region
		ss, err := systemSvc.GetSolarSystemsByRegion(e.ID)

		if err != nil {
			return nil, err
		}

		systems := make(map[string]*universe.SolarSystem)

		for _, f := range ss {
			s := universe.SolarSystem{
				ID:               f.ID,
				SystemName:       f.SystemName,
				RegionID:         f.RegionID,
				HoldingFactionID: f.HoldingFactionID,
				PosX:             f.PosX,
				PosY:             f.PosY,
			}

			// initialize and store system
			s.Initialize()
			systems[s.ID.String()] = &s

			// link universe into system
			s.Universe = &u

			// for jumphole linking later
			sMap[s.ID.String()] = &s

			// load ships
			ships, err := shipSvc.GetShipsBySolarSystem(s.ID, false)

			if err != nil {
				return nil, err
			}

			for _, sh := range ships {
				es, err := LoadShip(&sh)

				if err != nil {
					return nil, err
				}

				s.AddShip(es, true)
			}

			// load stars
			stars, err := starSvc.GetStarsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, st := range stars {
				star := universe.Star{
					ID:       st.ID,
					SystemID: st.SystemID,
					PosX:     st.PosX,
					PosY:     st.PosY,
					Texture:  st.Texture,
					Radius:   st.Radius,
					Mass:     st.Mass,
					Theta:    st.Theta,
				}

				s.AddStar(&star)
			}

			// load planets
			planets, err := planetSvc.GetPlanetsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, p := range planets {
				planet := universe.Planet{
					ID:         p.ID,
					SystemID:   p.SystemID,
					PlanetName: p.PlanetName,
					PosX:       p.PosX,
					PosY:       p.PosY,
					Texture:    p.Texture,
					Radius:     p.Radius,
					Mass:       p.Mass,
					Theta:      p.Theta,
				}

				s.AddPlanet(&planet)
			}

			// load asteroids
			asteroids, err := asteroidSvc.GetAsteroidsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, p := range asteroids {
				// get ore item type
				oi, err := itemTypeSvc.GetItemTypeByID(p.ItemTypeID)

				if err != nil {
					return nil, err
				}

				// get ore item family
				of, err := itemFamilySvc.GetItemFamilyByID(oi.Family)

				if err != nil {
					return nil, err
				}

				// build asteroid
				asteroid := universe.Asteroid{
					ID:             p.ID,
					SystemID:       p.SystemID,
					Name:           p.Name,
					Texture:        p.Texture,
					Radius:         p.Radius,
					Theta:          p.Theta,
					PosX:           p.PosX,
					PosY:           p.PosY,
					Yield:          p.Yield,
					Mass:           p.Mass,
					ItemTypeName:   oi.Name,
					ItemTypeID:     oi.ID,
					ItemFamilyName: of.FriendlyName,
					ItemFamilyID:   oi.Family,
					ItemTypeMeta:   universe.Meta(oi.Meta),
				}

				// store in universe
				s.AddAsteroid(&asteroid)
			}

			// load jumpholes
			jumpholes, err := jumpholeSvc.GetJumpholesBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, j := range jumpholes {
				jumphole := universe.Jumphole{
					ID:           j.ID,
					SystemID:     j.SystemID,
					OutSystemID:  j.OutSystemID,
					JumpholeName: j.JumpholeName,
					PosX:         j.PosX,
					PosY:         j.PosY,
					Texture:      j.Texture,
					Radius:       j.Radius,
					Mass:         j.Mass,
					Theta:        j.Theta,
				}

				s.AddJumphole(&jumphole)

				// for jumphole linking later
				jhMap[j.ID.String()] = &jumphole
			}

			// load npc stations
			stations, err := stationSvc.GetStationsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, currStation := range stations {
				// load station processes
				sqlProcesses, err := stationProcessSvc.GetStationProcessesByStation(currStation.ID)

				if err != nil {
					return nil, err
				}

				processes := make(map[string]*universe.StationProcess)

				for _, sp := range sqlProcesses {
					spx, err := LoadStationProcess(&sp)

					if err != nil {
						return nil, err
					}

					processes[spx.ID.String()] = spx
				}

				// build station
				station := universe.Station{
					ID:          currStation.ID,
					SystemID:    currStation.SystemID,
					StationName: currStation.StationName,
					PosX:        currStation.PosX,
					PosY:        currStation.PosY,
					Texture:     currStation.Texture,
					Radius:      currStation.Radius,
					Mass:        currStation.Mass,
					Theta:       currStation.Theta,
					FactionID:   currStation.FactionID,
					// link solar system into station
					CurrentSystem: &s,
					Processes:     processes,
					Faction:       *u.Factions[currStation.FactionID.String()],
				}

				// initialize station
				station.Initialize()

				// load open sell orders
				sos, err := sellOrderSvc.GetOpenSellOrdersByStation(station.ID)

				if err != nil {
					return nil, err
				}

				for _, o := range sos {
					// convert to engine type
					so := SellOrderFromSQL(&o)

					if so == nil {
						return nil, errors.New("unable to load sell order")
					}

					// load item for sale
					it, err := itemSvc.GetItemByID(so.ItemID)

					if err != nil {
						return nil, err
					}

					item, err := LoadItem(it)

					if err != nil {
						return nil, err
					}

					// link item into order
					so.Item = item

					// store on station
					station.OpenSellOrders[so.ID.String()] = so
				}

				// add to solar system
				s.AddStation(&station)
			}
		}

		// store and append region
		r.Systems = systems
		regions[r.ID.String()] = &r
	}

	// link jumpholes
	for _, j := range jhMap {
		// get out system
		o := sMap[j.OutSystemID.String()]

		// copy jumpholes
		jhs := o.CopyJumpholes()

		// find and link destination jumphole into jumphole
		for _, k := range jhs {
			if k.OutSystemID == j.SystemID {
				// get real jumphole pointer from map
				j.OutJumphole = jhMap[k.ID.String()]

				// link destination system into jumphole
				j.OutSystem = o
				break
			}
		}
	}

	// link regions into universe
	u.Regions = regions

	// return universe
	return &u, nil
}

// Saves the current state of dynamic entities in the simulation to the database
func saveUniverse(u *universe.Universe) {
	// get services
	stationSvc := sql.GetStationService()
	stationProcessSvc := sql.GetStationProcessService()

	// iterate over systems
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// get ships in system
			ships := s.CopyShips()

			// save ships to database
			for _, ship := range ships {
				saveShip(ship)
			}

			// get npc stations in system
			stations := s.CopyStations()

			// save npc stations to database
			for _, station := range stations {
				// obtain lock on copy
				station.Lock.Lock()
				defer station.Lock.Unlock()

				dbStation := sql.Station{
					ID:          station.ID,
					StationName: station.StationName,
					PosX:        station.PosX,
					PosY:        station.PosY,
					SystemID:    station.SystemID,
					Texture:     station.Texture,
					Theta:       station.Theta,
					Mass:        station.Mass,
					Radius:      station.Radius,
					FactionID:   station.FactionID,
				}

				err := stationSvc.UpdateStation(dbStation)

				if err != nil {
					log.Println(fmt.Sprintf("Error saving station: %v | %v", dbStation, err))
				}

				// save manufacturing processes
				for _, p := range station.Processes {
					// convert internal state to db model
					dbState := sql.StationProcessInternalState{
						IsRunning: p.InternalState.IsRunning,
					}

					dbState.Inputs = make(map[string]sql.StationProcessInternalStateFactor)
					dbState.Outputs = make(map[string]sql.StationProcessInternalStateFactor)

					for key := range p.InternalState.Inputs {
						v := p.InternalState.Inputs[key]
						dbState.Inputs[key] = sql.StationProcessInternalStateFactor{
							Quantity: v.Quantity,
							Price:    v.Price,
						}
					}

					for key := range p.InternalState.Outputs {
						v := p.InternalState.Outputs[key]
						dbState.Outputs[key] = sql.StationProcessInternalStateFactor{
							Quantity: v.Quantity,
							Price:    v.Price,
						}
					}

					// convert station process to db model
					dbProcess := sql.StationProcess{
						ID:            p.ID,
						StationID:     p.StationID,
						ProcessID:     p.ProcessID,
						Progress:      p.Progress,
						Installed:     p.Installed,
						InternalState: dbState,
						Meta:          sql.Meta(p.Meta),
					}

					err := stationProcessSvc.UpdateStationProcess(dbProcess)

					if err != nil {
						log.Println(fmt.Sprintf("Error saving station process: %v | %v", dbProcess, err))
					}
				}
			}
		}
	}
}

// Converts a SlotLayout from the SQL type to the engine type
func SlotLayoutFromSQL(value *sql.SlotLayout) universe.SlotLayout {
	// set up empty layout
	layout := universe.SlotLayout{}

	// null check
	if value == nil {
		// return empty layout
		return layout
	}

	// copy slot data into layout
	for _, v := range value.ASlots {
		slot := SlotFromSQL(&v)
		layout.ASlots = append(layout.ASlots, slot)
	}

	for _, v := range value.BSlots {
		slot := SlotFromSQL(&v)
		layout.BSlots = append(layout.BSlots, slot)
	}

	for _, v := range value.CSlots {
		slot := SlotFromSQL(&v)
		layout.CSlots = append(layout.CSlots, slot)
	}

	// return filled layout
	return layout
}

// Converts a Slot from the SQL type to the engine type
func SlotFromSQL(value *sql.Slot) universe.Slot {
	// set up empty slot
	slot := universe.Slot{}

	// null check
	if value == nil {
		// return empty slot
		return slot
	}

	// copy slot data
	slot.Family = value.Family
	slot.Volume = value.Volume
	slot.TexturePosition = value.TexturePosition

	// return filled slot
	return slot
}

// Converts a Fitting from the SQL type to the engine type
func FittingFromSQL(value *sql.Fitting) (*universe.Fitting, error) {
	// set up empty layout
	fitting := universe.Fitting{}

	// null check
	if value == nil {
		// return empty layout
		return &fitting, nil
	}

	// copy slot data into layout
	for _, v := range value.ARack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		slot.Rack = "A"
		fitting.ARack = append(fitting.ARack, *slot)
	}

	for _, v := range value.BRack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		slot.Rack = "B"
		fitting.BRack = append(fitting.BRack, *slot)
	}

	for _, v := range value.CRack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		slot.Rack = "C"
		fitting.CRack = append(fitting.CRack, *slot)
	}

	// return filled layout
	return &fitting, nil
}

// Converts a FittedSlot from the SQL type to the engine type
func FittedSlotFromSQL(value *sql.FittedSlot) (*universe.FittedSlot, error) {
	// get default uuid
	emptyUUID := uuid.UUID{}
	defaultUUID := emptyUUID.String()

	// get services
	itemSvc := sql.GetItemService()
	itemTypeSvc := sql.GetItemTypeService()
	itemFamilySvc := sql.GetItemFamilyService()

	// set up empty slot
	slot := universe.FittedSlot{}

	// null check
	if value == nil {
		// return empty slot
		return &slot, nil
	}

	// copy slot data
	slot.ItemID = value.ItemID
	slot.ItemTypeID = value.ItemTypeID

	// check if this slot is empty
	if slot.ItemID.String() == defaultUUID ||
		slot.ItemTypeID.String() == defaultUUID {
		// return empty slot
		return &slot, nil
	}

	// load item data
	item, err := itemSvc.GetItemByID(slot.ItemID)

	if err != nil {
		return nil, err
	}

	// load item type data
	itemType, err := itemTypeSvc.GetItemTypeByID(item.ItemTypeID)

	if err != nil {
		return nil, err
	}

	// load item family data
	itemFamily, err := itemFamilySvc.GetItemFamilyByID(itemType.Family)

	if err != nil {
		return nil, err
	}

	// store on slot
	slot.ItemMeta = MetaFromSQL(&item.Meta)
	slot.ItemTypeMeta = MetaFromSQL(&itemType.Meta)
	slot.ItemTypeFamily = itemType.Family
	slot.ItemTypeFamilyName = itemFamily.FriendlyName
	slot.ItemTypeName = itemType.Name

	// return filled slot
	return &slot, nil
}

// Converts a Fitting from the engine type to the SQL type
func SQLFromFitting(value *universe.Fitting) sql.Fitting {
	// set up empty layout
	fitting := sql.Fitting{}

	// null check
	if value == nil {
		// return empty layout
		return fitting
	}

	// copy slot data into layout
	for _, v := range value.ARack {
		slot := SQLFromFittedSlot(&v)
		fitting.ARack = append(fitting.ARack, slot)
	}

	for _, v := range value.BRack {
		slot := SQLFromFittedSlot(&v)
		fitting.BRack = append(fitting.BRack, slot)
	}

	for _, v := range value.CRack {
		slot := SQLFromFittedSlot(&v)
		fitting.CRack = append(fitting.CRack, slot)
	}

	// return filled layout
	return fitting
}

// Converts a FittedSlot from the engine type to the SQL type
func SQLFromFittedSlot(value *universe.FittedSlot) sql.FittedSlot {
	// set up empty slot
	slot := sql.FittedSlot{}

	// null check
	if value == nil {
		// return empty slot
		return slot
	}

	// copy slot data
	slot.ItemID = value.ItemID
	slot.ItemTypeID = value.ItemTypeID

	// return filled slot
	return slot
}

// Converts generic metadata from the SQL type to the engine type
func MetaFromSQL(value *sql.Meta) universe.Meta {
	// set up empty metadata
	meta := universe.Meta{}

	// null check
	if value == nil {
		// return empty slot
		return meta
	}

	// copy metadata
	for k, v := range *value {
		meta[k] = v
	}

	// return filled meta
	return meta
}

// Converts generic metadata from the SQL type to the engine type
func SQLFromMeta(value *universe.Meta) sql.Meta {
	// set up empty metadata
	meta := sql.Meta{}

	// null check
	if value == nil {
		// return empty slot
		return meta
	}

	// copy metadata
	for k, v := range *value {
		meta[k] = v
	}

	// return filled meta
	return meta
}

// Converts a StartFitting from the SQL type to the engine type
func StartFittingFromSQL(value *sql.StartFitting) universe.StartFitting {
	// set up empty layout
	layout := universe.StartFitting{}

	// null check
	if value == nil {
		// return empty layout
		return layout
	}

	// copy slot data into layout
	for _, v := range value.ARack {
		slot := StartFittedSlotFromSQL(&v)
		layout.ARack = append(layout.ARack, slot)
	}

	for _, v := range value.BRack {
		slot := StartFittedSlotFromSQL(&v)
		layout.BRack = append(layout.BRack, slot)
	}

	for _, v := range value.CRack {
		slot := StartFittedSlotFromSQL(&v)
		layout.CRack = append(layout.CRack, slot)
	}

	// return filled layout
	return layout
}

// Converts a StartFittedSlot from the SQL type to the engine type
func StartFittedSlotFromSQL(value *sql.StartFittedSlot) universe.StartFittedSlot {
	// set up empty slot
	slot := universe.StartFittedSlot{}

	// null check
	if value == nil {
		// return empty slot
		return slot
	}

	// copy slot data
	slot.ItemTypeID = value.ItemTypeID

	// return filled slot
	return slot
}

// Converts a Container from the SQL type to the engine type
func ContainerFromSQL(value *sql.Container) *universe.Container {
	// set up empty container
	container := universe.Container{}

	// null check
	if value == nil {
		// return nil
		return nil
	}

	// copy container data
	container.ID = value.ID
	container.Created = value.Created
	container.Meta = MetaFromSQL(&value.Meta)

	// return filled container
	return &container
}

// Converts a Faction from the SQL type to the engine type
func FactionFromSQL(value *sql.Faction) *universe.Faction {
	// set up empty faction
	faction := universe.Faction{}

	// null check
	if value == nil {
		// return nil
		return nil
	}

	// copy basic data
	faction.ID = value.ID
	faction.Name = value.Name
	faction.Description = value.Description
	faction.IsNPC = value.IsNPC
	faction.IsJoinable = value.IsJoinable
	faction.IsClosed = value.IsClosed
	faction.CanHoldSov = value.CanHoldSov
	faction.Meta = universe.Meta(value.Meta)
	faction.Ticker = value.Ticker

	// copy reputation sheet entries
	faction.ReputationSheet = shared.FactionReputationSheet{
		Entries:        make(map[string]shared.FactionReputationSheetEntry),
		HostFactionIDs: make([]uuid.UUID, 0),
		WorldPercent:   0,
	}

	if value.ReputationSheet.Entries != nil {
		for k, v := range value.ReputationSheet.Entries {
			faction.ReputationSheet.Entries[k] = shared.FactionReputationSheetEntry{
				SourceFactionID:  v.SourceFactionID,
				TargetFactionID:  v.TargetFactionID,
				StandingValue:    v.StandingValue,
				AreOpenlyHostile: v.AreOpenlyHostile,
			}
		}
	}

	// copy world data
	if value.ReputationSheet.HostFactionIDs != nil {
		faction.ReputationSheet.HostFactionIDs = append(faction.ReputationSheet.HostFactionIDs, value.ReputationSheet.HostFactionIDs...)
	}

	faction.ReputationSheet.WorldPercent = value.ReputationSheet.WorldPercent

	// return filled faction
	return &faction
}

// Converts a SellOrder from the SQL type to the engine type
func SellOrderFromSQL(value *sql.SellOrder) *universe.SellOrder {
	// set up empty sell order
	order := universe.SellOrder{}

	// null check
	if value == nil {
		// return nil
		return nil
	}

	// copy order data
	order.ID = value.ID
	order.StationID = value.StationID
	order.ItemID = value.ItemID
	order.SellerUserID = value.SellerUserID
	order.AskPrice = value.AskPrice
	order.Created = value.Created
	order.Bought = value.Bought
	order.BuyerUserID = value.BuyerUserID

	// return filled order
	return &order
}

// Converts a SellOrder from the engine type to the SQL type
func SQLFromSellOrder(value *universe.SellOrder) *sql.SellOrder {
	// set up empty sell order
	order := sql.SellOrder{}

	// null check
	if value == nil {
		// return nil
		return nil
	}

	// copy order data
	order.ID = value.ID
	order.StationID = value.StationID
	order.ItemID = value.ItemID
	order.SellerUserID = value.SellerUserID
	order.AskPrice = value.AskPrice
	order.Created = value.Created
	order.Bought = value.Bought
	order.BuyerUserID = value.BuyerUserID

	// return filled order
	return &order
}

// Converts an Item from the SQL type to the engine type
func ItemFromSQL(value *sql.Item) *universe.Item {
	// set up empty item
	item := universe.Item{}

	// null check
	if value == nil {
		// return nil
		return nil
	}

	// copy item data
	item.ID = value.ID
	item.ItemTypeID = value.ItemTypeID
	item.ContainerID = value.ContainerID
	item.Created = value.Created
	item.CreatedBy = value.CreatedBy
	item.CreatedReason = value.CreatedReason
	item.Meta = MetaFromSQL(&value.Meta)
	item.Quantity = value.Quantity
	item.IsPackaged = value.IsPackaged

	// return filled item
	return &item
}

// Converts an item from the engine type to the SQL type
func SQLFromItem(value *universe.Item) *sql.Item {
	// set up empty item
	item := sql.Item{}

	// copy item data
	item.ID = value.ID
	item.ItemTypeID = value.ItemTypeID
	item.ContainerID = value.ContainerID
	item.Created = value.Created
	item.CreatedBy = value.CreatedBy
	item.CreatedReason = value.CreatedReason
	item.Meta = SQLFromMeta(&value.Meta)
	item.Quantity = value.Quantity
	item.IsPackaged = value.IsPackaged

	// return filled item
	return &item
}

// Loads an item with some type and family data for use in the simulation.
func LoadItem(i *sql.Item) (*universe.Item, error) {
	itemTypeSvc := sql.GetItemTypeService()
	itemFamilySvc := sql.GetItemFamilyService()

	// load base item
	ei := ItemFromSQL(i)

	// null check
	if ei == nil {
		return nil, errors.New("item from SQL failed")
	}

	// load item type
	it, err := itemTypeSvc.GetItemTypeByID(ei.ItemTypeID)

	if err != nil {
		return nil, err
	}

	// include item type data
	ei.ItemTypeName = it.Name
	ei.ItemFamilyID = it.Family
	ei.ItemTypeMeta = MetaFromSQL(&it.Meta)

	// load item family
	fm, err := itemFamilySvc.GetItemFamilyByID(it.Family)

	if err != nil {
		return nil, err
	}

	// include item family data
	ei.ItemFamilyName = fm.FriendlyName

	// return filled item
	return ei, nil
}

// Takes a SQL Container and converts it, and items it loads, into the engine type ready for insertion into the universe.
func LoadContainer(c *sql.Container) (*universe.Container, error) {
	// load base container
	container := ContainerFromSQL(c)

	// load items
	itemSvc := sql.GetItemService()

	s, err := itemSvc.GetItemsByContainer(container.ID)

	if err != nil {
		return nil, err
	}

	// load items and push into container
	for _, i := range s {
		m, err := LoadItem(&i)

		if err != nil {
			return nil, err
		}

		// null check
		if m == nil {
			return nil, errors.New("item argument was nil")
		}

		// push into container
		container.Items = append(container.Items, m)
	}

	// return full container
	return container, nil
}

// Takes a SQL Process and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadProcess(p *sql.Process) (*universe.Process, error) {
	inputSvc := sql.GetProcessInputService()
	outputSvc := sql.GetProcessOutputService()
	itemTypeSvc := sql.GetItemTypeService()
	itemFamilySvc := sql.GetItemFamilyService()

	// get inputs
	inputs, err := inputSvc.GetProcessInputsByProcess(p.ID)

	if err != nil {
		return nil, err
	}

	// get outputs
	outputs, err := outputSvc.GetProcessOutputsByProcess(p.ID)

	if err != nil {
		return nil, err
	}

	// build in-memory process
	process := universe.Process{
		ID:   p.ID,
		Name: p.Name,
		Meta: universe.Meta(p.Meta),
		Time: p.Time,
	}

	i := make(map[string]universe.ProcessInput)
	for _, e := range inputs {
		// get item type and family
		itemType, err := itemTypeSvc.GetItemTypeByID(e.ItemTypeID)

		if err != nil {
			return nil, err
		}

		itemFamily, err := itemFamilySvc.GetItemFamilyByID(itemType.Family)

		if err != nil {
			return nil, err
		}

		// build model
		i[itemType.ID.String()] = universe.ProcessInput{
			ID:             e.ID,
			ItemTypeID:     e.ItemTypeID,
			Quantity:       e.Quantity,
			Meta:           universe.Meta(e.Meta),
			ProcessID:      e.ProcessID,
			ItemTypeName:   itemType.Name,
			ItemFamilyName: itemFamily.FriendlyName,
			ItemFamilyID:   itemType.Family,
			ItemTypeMeta:   universe.Meta(itemType.Meta),
		}
	}

	o := make(map[string]universe.ProcessOutput)
	for _, e := range outputs {
		// get item type and family
		itemType, err := itemTypeSvc.GetItemTypeByID(e.ItemTypeID)

		if err != nil {
			return nil, err
		}

		itemFamily, err := itemFamilySvc.GetItemFamilyByID(itemType.Family)

		if err != nil {
			return nil, err
		}

		// build model
		o[itemType.ID.String()] = universe.ProcessOutput{
			ID:             e.ID,
			ItemTypeID:     e.ItemTypeID,
			Quantity:       e.Quantity,
			Meta:           universe.Meta(e.Meta),
			ProcessID:      e.ProcessID,
			ItemTypeName:   itemType.Name,
			ItemFamilyName: itemFamily.FriendlyName,
			ItemFamilyID:   itemType.Family,
			ItemTypeMeta:   universe.Meta(itemType.Meta),
		}
	}

	process.Inputs = i
	process.Outputs = o

	return &process, nil
}

// Takes a SQL Station Process and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadStationProcess(sp *sql.StationProcess) (*universe.StationProcess, error) {
	procesSvc := sql.GetProcessService()

	// load process
	sqlP, err := procesSvc.GetProcessByID(sp.ProcessID)

	if err != nil {
		return nil, err
	}

	process, err := LoadProcess(sqlP)

	if err != nil {
		return nil, err
	}

	// convert internal state
	internalState := universe.StationProcessInternalState{
		IsRunning: sp.InternalState.IsRunning,
	}

	internalState.Inputs = make(map[string]*universe.StationProcessInternalStateFactor)
	internalState.Outputs = make(map[string]*universe.StationProcessInternalStateFactor)

	for key := range sp.InternalState.Inputs {
		v := sp.InternalState.Inputs[key]
		internalState.Inputs[key] = &universe.StationProcessInternalStateFactor{
			Quantity: v.Quantity,
			Price:    v.Price,
		}
	}

	for key := range sp.InternalState.Outputs {
		v := sp.InternalState.Outputs[key]
		internalState.Outputs[key] = &universe.StationProcessInternalStateFactor{
			Quantity: v.Quantity,
			Price:    v.Price,
		}
	}

	// build station process
	o := universe.StationProcess{
		ID:            sp.ID,
		StationID:     sp.StationID,
		ProcessID:     sp.ProcessID,
		Progress:      sp.Progress,
		Installed:     sp.Installed,
		InternalState: internalState,
		Meta:          universe.Meta(sp.Meta),
		Process:       *process,
	}

	return &o, nil
}

// Takes a SQL Ship and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadShip(sh *sql.Ship) (*universe.Ship, error) {
	shipTmpSvc := sql.GetShipTemplateService()
	userSvc := sql.GetUserService()
	containerSvc := sql.GetContainerService()

	// get template
	temp, err := shipTmpSvc.GetShipTemplateByID(sh.ShipTemplateID)

	if err != nil {
		return nil, err
	}

	// get owner info
	owner, err := userSvc.GetUserByID((sh.UserID))

	if err != nil {
		return nil, err
	}

	// get fitting
	fitting, err := FittingFromSQL(&sh.Fitting)

	if err != nil {
		return nil, err
	}

	// load cargo bay
	cb, err := containerSvc.GetContainerByID(sh.CargoBayContainerID)

	if err != nil {
		return nil, err
	}

	cargoBay, err := LoadContainer(cb)

	if err != nil {
		return nil, err
	}

	// load fitting bay
	fb, err := containerSvc.GetContainerByID(sh.FittingBayContainerID)

	if err != nil {
		return nil, err
	}

	fittingBay, err := LoadContainer(fb)

	if err != nil {
		return nil, err
	}

	// build in-memory ship
	es := universe.Ship{
		ID:                    sh.ID,
		UserID:                sh.UserID,
		Created:               sh.Created,
		ShipName:              sh.ShipName,
		OwnerName:             owner.Username,
		PosX:                  sh.PosX,
		PosY:                  sh.PosY,
		SystemID:              sh.SystemID,
		Texture:               sh.Texture,
		Theta:                 sh.Theta,
		VelX:                  sh.VelX,
		VelY:                  sh.VelY,
		Shield:                sh.Shield,
		Armor:                 sh.Armor,
		Hull:                  sh.Hull,
		Fuel:                  sh.Fuel,
		Heat:                  sh.Heat,
		Energy:                sh.Energy,
		DockedAtStationID:     sh.DockedAtStationID,
		Fitting:               *fitting,
		Destroyed:             sh.Destroyed,
		DestroyedAt:           sh.DestroyedAt,
		CargoBayContainerID:   sh.CargoBayContainerID,
		CargoBay:              cargoBay,
		FittingBayContainerID: sh.FittingBayContainerID,
		FittingBay:            fittingBay,
		ReMaxDirty:            sh.ReMaxDirty,
		TrashContainerID:      sh.TrashContainerID,
		Wallet:                sh.Wallet,
		BeingFlownByPlayer:    owner.CurrentShipID != nil && (sh.ID == *owner.CurrentShipID),
		TemplateData: universe.ShipTemplate{
			ID:                 temp.ID,
			Created:            temp.Created,
			ShipTemplateName:   temp.ShipTemplateName,
			Texture:            temp.Texture,
			Radius:             temp.Radius,
			BaseAccel:          temp.BaseAccel,
			BaseMass:           temp.BaseMass,
			BaseTurn:           temp.BaseTurn,
			BaseShield:         temp.BaseShield,
			BaseShieldRegen:    temp.BaseShieldRegen,
			BaseArmor:          temp.BaseArmor,
			BaseHull:           temp.BaseHull,
			BaseFuel:           temp.BaseFuel,
			BaseHeatCap:        temp.BaseHeatCap,
			BaseHeatSink:       temp.BaseHeatSink,
			BaseEnergy:         temp.BaseEnergy,
			BaseEnergyRegen:    temp.BaseEnergyRegen,
			ShipTypeID:         temp.ShipTypeID,
			SlotLayout:         SlotLayoutFromSQL(&temp.SlotLayout),
			BaseCargoBayVolume: temp.BaseCargoBayVolume,
		},
		FactionID: owner.CurrentFactionID,
	}

	// get pointer to ship
	sp := &es

	// link ship into fitting
	for i := range sp.Fitting.ARack {
		m := &sp.Fitting.ARack[i]
		m.LinkShip(sp)
	}

	for i := range sp.Fitting.BRack {
		m := &sp.Fitting.BRack[i]
		m.LinkShip(sp)
	}

	for i := range sp.Fitting.CRack {
		m := &sp.Fitting.CRack[i]
		m.LinkShip(sp)
	}

	// return pointer to ship
	return sp, nil
}

// Updates a ship in the database
func saveShip(ship *universe.Ship) error {
	// get ship service
	shipSvc := sql.GetShipService()

	// obtain lock on copy
	ship.Lock.Lock()
	defer ship.Lock.Unlock()

	dbShip := sql.Ship{
		ID:                    ship.ID,
		UserID:                ship.UserID,
		Created:               ship.Created,
		ShipName:              ship.ShipName,
		PosX:                  ship.PosX,
		PosY:                  ship.PosY,
		SystemID:              ship.SystemID,
		Texture:               ship.Texture,
		Theta:                 ship.Theta,
		VelX:                  ship.VelX,
		VelY:                  ship.VelY,
		Shield:                ship.Shield,
		Armor:                 ship.Armor,
		Hull:                  ship.Hull,
		Fuel:                  ship.Fuel,
		Heat:                  ship.Heat,
		Energy:                ship.Energy,
		ShipTemplateID:        ship.TemplateData.ID,
		DockedAtStationID:     ship.DockedAtStationID,
		Fitting:               SQLFromFitting(&ship.Fitting),
		Destroyed:             ship.Destroyed,
		DestroyedAt:           ship.DestroyedAt,
		CargoBayContainerID:   ship.CargoBayContainerID,
		FittingBayContainerID: ship.FittingBayContainerID,
		TrashContainerID:      ship.TrashContainerID,
		ReMaxDirty:            ship.ReMaxDirty,
		Wallet:                ship.Wallet,
	}

	err := shipSvc.UpdateShip(dbShip)

	if err != nil {
		log.Println(fmt.Sprintf("Error saving ship: %v | %v", dbShip, err))
	}

	return err
}

// Moves an item to a different container in the database
func saveItemLocation(itemID uuid.UUID, containerID uuid.UUID) error {
	itemSvc := sql.GetItemService()
	return itemSvc.SetContainerID(itemID, containerID)
}

// Marks an item as packaged in the database
func packageItem(itemID uuid.UUID) error {
	itemSvc := sql.GetItemService()
	return itemSvc.PackageItem(itemID)
}

// Marks an item as unpackaged in the database
func unpackageItem(itemID uuid.UUID, meta universe.Meta) error {
	itemSvc := sql.GetItemService()
	return itemSvc.UnpackageItem(itemID, SQLFromMeta(&meta))
}

// Changes the quantity of an item stack in the database
func changeQuantity(itemID uuid.UUID, quantity int) error {
	itemSvc := sql.GetItemService()
	return itemSvc.ChangeQuantity(itemID, quantity)
}

// Saves a new item to the database
func newItem(item *universe.Item) (*uuid.UUID, error) {
	itemSvc := sql.GetItemService()

	// convert to sql type
	sql := SQLFromItem(item)

	if sql == nil {
		return nil, errors.New("error converting item to SQL type")
	}

	// save item
	o, err := itemSvc.NewItem(*sql)

	return &o.ID, err
}

// Saves a new sell order to the database
func newSellOrder(sellOrder *universe.SellOrder) (*uuid.UUID, error) {
	sellOrderSvc := sql.GetSellOrderService()

	// convert to sql type
	sql := SQLFromSellOrder(sellOrder)

	if sql == nil {
		return nil, errors.New("error converting sell order to SQL type")
	}

	// save sell order
	o, err := sellOrderSvc.NewSellOrder(*sql)

	return &o.ID, err
}

// Saves a new sell order to the database
func markSellOrderAsBought(sellOrder *universe.SellOrder) error {
	sellOrderSvc := sql.GetSellOrderService()

	// convert to sql type
	sql := SQLFromSellOrder(sellOrder)

	if sql == nil {
		return errors.New("error converting sell order to SQL type")
	}

	// save sell order
	err := sellOrderSvc.MarkSellOrderAsBought(*sql)

	return err
}
