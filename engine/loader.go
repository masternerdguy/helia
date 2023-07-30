package engine

import (
	"errors"
	"fmt"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"sync"

	"github.com/google/uuid"
)

// caches to reduce trips to database
var processCache = make(map[string]sql.Process)
var processInputCache = make(map[string][]sql.ProcessInput)
var processOutputCache = make(map[string][]sql.ProcessOutput)
var itemTypeCache = make(map[string]sql.ItemType)
var itemFamilyCache = make(map[string]sql.ItemFamily)

// Loads the state of the universe from the database
func LoadUniverse() (*universe.Universe, error) {
	// caches that need to be garbage collected once loading is complete
	stationProcessCache := make(map[string][]sql.StationProcess)

	// preload station processes
	allStationProcesses, err := stationProcessSvc.GetAllStationProcesses()

	if err != nil {
		return nil, err
	}

	for _, asp := range allStationProcesses {
		k := asp.StationID.String()

		// check if key exists
		v := stationProcessCache[k]

		if v == nil {
			// initialize entry
			stationProcessCache[k] = make([]sql.StationProcess, 0)
			v = stationProcessCache[k]
		}

		// append to entry
		v = append(v, asp)

		// update cache
		stationProcessCache[k] = v
	}

	// preload processes
	allProcesses, err := processSvc.GetAllProcesses()

	if err != nil {
		return nil, err
	}

	for _, ap := range allProcesses {
		processCache[ap.ID.String()] = ap
	}

	// preload item types
	allItemTypes, err := itemTypeSvc.GetAllItemTypes()

	if err != nil {
		return nil, err
	}

	for _, ap := range allItemTypes {
		itemTypeCache[ap.ID.String()] = ap
	}

	// preload process inputs
	allProcessInputs, err := processInputSvc.GetAllProcessInputs()

	if err != nil {
		return nil, err
	}

	for _, asp := range allProcessInputs {
		k := asp.ProcessID.String()

		// check if key exists
		v := processInputCache[k]

		if v == nil {
			// initialize entry
			processInputCache[k] = make([]sql.ProcessInput, 0)
			v = processInputCache[k]
		}

		// append to entry
		v = append(v, asp)

		// update cache
		processInputCache[k] = v
	}

	// preload process outputs
	allProcessOutputs, err := processOutputSvc.GetAllProcessOutputs()

	if err != nil {
		return nil, err
	}

	for _, asp := range allProcessOutputs {
		k := asp.ProcessID.String()

		// check if key exists
		v := processOutputCache[k]

		if v == nil {
			// initialize entry
			processOutputCache[k] = make([]sql.ProcessOutput, 0)
			v = processOutputCache[k]
		}

		// append to entry
		v = append(v, asp)

		// update cache
		processOutputCache[k] = v
	}

	// preload item families
	allItemFamilies, err := itemFamilySvc.GetAllItemFamilies()

	if err != nil {
		return nil, err
	}

	for _, ap := range allItemFamilies {
		itemFamilyCache[ap.ID] = ap
	}

	// empty universe to fill
	u := universe.Universe{}

	// load schematic runs (without pointers)
	srs, err := schematicRunSvc.GetUndeliveredSchematicRuns()

	if err != nil {
		return nil, err
	}

	// hook into runner
	initializeSchematicsWatcher()

	for _, sr := range srs {
		usr := universe.SchematicRun{
			ID:              sr.ID,
			Created:         sr.Created,
			ProcessID:       sr.ProcessID,
			StatusID:        sr.StatusID,
			Progress:        sr.Progress,
			SchematicItemID: sr.SchematicItemID,
			UserID:          sr.UserID,
			Lock:            sync.Mutex{},
		}

		// don't load stuck ones
		if usr.StatusID == "error" {
			shared.TeeLog(fmt.Sprintf("Not loading schematic run %v (error)", usr.ID))
			continue
		} else if usr.StatusID == "deliverypending" {
			shared.TeeLog(fmt.Sprintf("Not loading schematic run %v (deliverypending)", usr.ID))
			continue
		} else if usr.StatusID == "delivered" {
			shared.TeeLog(fmt.Sprintf("Not loading schematic run %v (delivered)", usr.ID))
			continue
		}

		addSchematicRunForUser(sr.UserID, &usr)
	}

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
	for rIDX, e := range rs {
		// progress output
		shared.TeeLog(fmt.Sprintf("loading region %v of %v", rIDX+1, len(rs)))

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
			hf := factions[f.HoldingFactionID.String()]
			hfn := ""

			if hf != nil {
				hfn = hf.Name
			}

			s := universe.SolarSystem{
				ID:                 f.ID,
				SystemName:         f.SystemName,
				RegionID:           f.RegionID,
				RegionName:         e.RegionName,
				HoldingFactionID:   f.HoldingFactionID,
				HoldingFactionName: hfn,
				PosX:               f.PosX,
				PosY:               f.PosY,
				Lock:               sync.Mutex{},
			}

			// initialize and store system
			s.Initialize()
			systems[s.ID.String()] = &s

			// link universe into system
			s.Universe = &u

			// for jumphole linking later
			sMap[s.ID.String()] = &s

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
				oi := itemTypeCache[p.ItemTypeID.String()]

				// get ore item family
				of := itemFamilyCache[oi.Family]

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
					Lock:           sync.Mutex{},
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
					Lock:         sync.Mutex{},
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
				// get station processes for this station from cache
				sqlProcesses := stationProcessCache[currStation.ID.String()]

				// process collected station processes
				processes := make(map[string]*universe.StationProcess)

				for _, sp := range sqlProcesses {
					// get process for this station process
					spc := processCache[sp.ProcessID.String()]

					// finish loading station process
					spx, err := LoadStationProcess(&sp, &spc)

					if err != nil {
						return nil, err
					}

					spx.StationName = currStation.StationName
					spx.SolarSystemName = s.SystemName
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
					Lock:        sync.Mutex{},
					// link solar system into station
					CurrentSystem: &s,
					Processes:     processes,
					Faction:       u.Factions[currStation.FactionID.String()],
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

			// load player outposts
			outposts, err := outpostSvc.GetOutpostsBySolarSystem(s.ID, false)

			if err != nil {
				return nil, err
			}

			for _, currOutpost := range outposts {
				// load outpost
				outpost, err := LoadOutpost(&currOutpost, &u)

				if err != nil {
					return nil, err
				}

				// add to solar system
				osm := s.AddOutpost(outpost, false)

				// load open sell orders
				sos, err := sellOrderSvc.GetOpenSellOrdersByStation(osm.ID)

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
					osm.OpenSellOrders[so.ID.String()] = so
				}
			}

			// load ships
			ships, err := shipSvc.GetShipsBySolarSystem(s.ID, false)

			if err != nil {
				return nil, err
			}

			for _, sh := range ships {
				es, err := LoadShip(&sh, &u)

				if err != nil {
					return nil, err
				}

				s.AddShip(es, true)
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
		jhs := o.CopyJumpholes(true)

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

	// convert item type cache to universe type
	universeItemTypes := make(map[string]*universe.ItemTypeRaw)

	for _, t := range allItemTypes {
		raw := universe.ItemTypeRaw{
			ID:     t.ID,
			Family: t.Family,
			Name:   t.Name,
			Meta:   universe.Meta(t.Meta),
		}

		universeItemTypes[t.ID.String()] = &raw
	}

	// convert item family cache to universe type
	universeItemFamilies := make(map[string]*universe.ItemFamilyRaw)

	for _, t := range allItemFamilies {
		raw := universe.ItemFamilyRaw{
			ID:           t.ID,
			FriendlyName: t.FriendlyName,
			Meta:         universe.Meta(t.Meta),
		}

		universeItemFamilies[t.ID] = &raw
	}

	// store item type and family caches in universe
	u.CachedItemTypes = universeItemTypes
	u.CachedItemFamilies = universeItemFamilies

	// return universe
	return &u, nil
}

// Saves the current state of dynamic entities in the simulation to the database
func saveUniverse(u *universe.Universe) {
	// iterate over systems
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// get ships in system
			ships := s.CopyShips(false)

			// save ships to database
			for _, ship := range ships {
				err := saveShip(ship)

				if err != nil {
					shared.TeeLog(fmt.Sprintf("Error saving ship: %v | %v", ship, err))
				}
			}

			// get npc stations in system
			stations := s.CopyStations(false)

			// save npc stations to database
			for _, station := range stations {
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
					shared.TeeLog(fmt.Sprintf("Error saving station: %v | %v", dbStation, err))
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
						shared.TeeLog(fmt.Sprintf("Error saving station process: %v | %v", dbProcess, err))
					}
				}
			}
		}
	}

	for _, sre := range schematicRunMap {
		for _, sr := range sre {
			// skip if never initialized
			if !sr.Initialized {
				continue
			}

			// convert to sql type
			usr := sql.SchematicRun{
				ID:              sr.ID,
				Created:         sr.Created,
				ProcessID:       sr.ProcessID,
				StatusID:        sr.StatusID,
				Progress:        sr.Progress,
				SchematicItemID: sr.SchematicItemID,
				UserID:          sr.UserID,
			}

			// save changes
			err := schematicRunSvc.UpdateSchematicRun(usr)

			if err != nil {
				shared.TeeLog(fmt.Sprintf("Error saving schematic run: %v | %v", usr, err))
			}
		}
	}
}

// Converts a reputation sheet to its SQL representation for saving to the db
func SQLFromPlayerReputationSheet(value *shared.PlayerReputationSheet) sql.PlayerReputationSheet {
	sheet := sql.PlayerReputationSheet{
		FactionEntries: make(map[string]sql.PlayerReputationSheetFactionEntry),
	}

	// null check
	if value == nil {
		// return empty sheet
		return sheet
	}

	// copy into sheet
	for k, v := range value.FactionEntries {
		sheet.FactionEntries[k] = sql.PlayerReputationSheetFactionEntry{
			FactionID:        v.FactionID,
			StandingValue:    v.StandingValue,
			AreOpenlyHostile: v.AreOpenlyHostile,
		}
	}

	return sheet
}

// Converts an experience sheet to its SQL representation for saving to the db
func SQLFromPlayerExperienceSheet(value *shared.PlayerExperienceSheet) sql.PlayerExperienceSheet {
	sheet := sql.PlayerExperienceSheet{
		ShipExperience:   make(map[string]sql.PlayerShipExperienceEntry),
		ModuleExperience: make(map[string]sql.PlayerModuleExperienceEntry),
	}

	// null check
	if value == nil {
		// return empty sheet
		return sheet
	}

	// copy into sheet
	for k, v := range value.ShipExperience {
		sheet.ShipExperience[k] = sql.PlayerShipExperienceEntry{
			SecondsOfExperience: v.SecondsOfExperience,
			ShipTemplateID:      v.ShipTemplateID,
			ShipTemplateName:    v.ShipTemplateName,
		}
	}

	for k, v := range value.ModuleExperience {
		sheet.ModuleExperience[k] = sql.PlayerModuleExperienceEntry{
			SecondsOfExperience: v.SecondsOfExperience,
			ItemTypeID:          v.ItemTypeID,
			ItemTypeName:        v.ItemTypeName,
		}
	}

	return sheet
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
	slot.SlotIndex = value.SlotIndex

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
	itemType := itemTypeCache[item.ItemTypeID.String()]

	// load item family data
	itemFamily := itemFamilyCache[itemType.Family]

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
	slot.SlotIndex = value.SlotIndex

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
	container := universe.Container{
		Lock: sync.Mutex{},
	}

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
	faction := universe.Faction{
		Applications: make(map[string]universe.FactionApplicationTicket),
	}

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
	faction.CanHoldSov = value.CanHoldSov
	faction.Meta = universe.Meta(value.Meta)
	faction.Ticker = value.Ticker
	faction.OwnerID = value.OwnerID
	faction.HomeStationID = value.HomeStationID

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

	// create mutex
	faction.Lock = sync.Mutex{}

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
	item := universe.Item{
		Lock: sync.Mutex{},
	}

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
	// load base item
	ei := ItemFromSQL(i)

	// null check
	if ei == nil {
		return nil, errors.New("item from SQL failed")
	}

	// load item type
	it := itemTypeCache[ei.ItemTypeID.String()]

	// include item type data
	ei.ItemTypeName = it.Name
	ei.ItemFamilyID = it.Family
	ei.ItemTypeMeta = MetaFromSQL(&it.Meta)

	// load item family
	fm := itemFamilyCache[it.Family]

	// include item family data
	ei.ItemFamilyName = fm.FriendlyName

	if ei.ItemFamilyID == "schematic" {
		// load associated process from metadata
		l, f := ei.ItemTypeMeta.GetMap("industrialmarket")

		if f {
			processidStr, idf := l.GetString("process_id")

			if idf {
				pid := uuid.MustParse(processidStr)
				pc, err := processSvc.GetProcessByID(pid)

				if err != nil {
					return nil, err
				}

				p, err := LoadProcess(pc)

				if err != nil {
					return nil, err
				}

				ei.Process = p
			}
		}
	}

	// return filled item
	return ei, nil
}

// Takes a SQL Container and converts it, and items it loads, into the engine type ready for insertion into the universe.
func LoadContainer(c *sql.Container) (*universe.Container, error) {
	// load base container
	container := ContainerFromSQL(c)

	// load items
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
	// get inputs
	inputs := processInputCache[p.ID.String()]

	// get outputs
	outputs := processOutputCache[p.ID.String()]

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
		itemType := itemTypeCache[e.ItemTypeID.String()]
		itemFamily := itemFamilyCache[itemType.Family]

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
		itemType := itemTypeCache[e.ItemTypeID.String()]
		itemFamily := itemFamilyCache[itemType.Family]

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
func LoadStationProcess(sp *sql.StationProcess, sqlP *sql.Process) (*universe.StationProcess, error) {
	// load process
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
		Lock:          sync.Mutex{},
	}

	return &o, nil
}

// Takes a SQL User and extracts its reputation sheet for use in the game engine
func LoadReputationSheet(u *sql.User) *shared.PlayerReputationSheet {
	repSheet := shared.PlayerReputationSheet{
		FactionEntries: make(map[string]*shared.PlayerReputationSheetFactionEntry),
		UserID:         u.ID,
		CharacterName:  u.CharacterName,
	}

	for k, v := range u.ReputationSheet.FactionEntries {
		repSheet.FactionEntries[k] = &shared.PlayerReputationSheetFactionEntry{
			FactionID:        v.FactionID,
			StandingValue:    v.StandingValue,
			AreOpenlyHostile: v.AreOpenlyHostile,
		}
	}

	return &repSheet
}

// Takes a SQL User and extracts its experience sheet for use in the game engine
func LoadExperienceSheet(u *sql.User) *shared.PlayerExperienceSheet {
	expSheet := shared.PlayerExperienceSheet{
		ShipExperience:   make(map[string]*shared.ShipExperienceEntry),
		ModuleExperience: make(map[string]*shared.ModuleExperienceEntry),
	}

	for k, v := range u.ExperienceSheet.ShipExperience {
		expSheet.ShipExperience[k] = &shared.ShipExperienceEntry{
			SecondsOfExperience: v.SecondsOfExperience,
			ShipTemplateID:      v.ShipTemplateID,
			ShipTemplateName:    v.ShipTemplateName,
		}
	}

	for k, v := range u.ExperienceSheet.ModuleExperience {
		expSheet.ModuleExperience[k] = &shared.ModuleExperienceEntry{
			SecondsOfExperience: v.SecondsOfExperience,
			ItemTypeID:          v.ItemTypeID,
			ItemTypeName:        v.ItemTypeName,
		}
	}

	return &expSheet
}

// Takes a SQL Ship and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadShip(sh *sql.Ship, u *universe.Universe) (*universe.Ship, error) {
	// get template
	temp, err := shipTemplateSvc.GetShipTemplateByID(sh.ShipTemplateID)

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
		ID:                       sh.ID,
		UserID:                   sh.UserID,
		Created:                  sh.Created,
		ShipName:                 sh.ShipName,
		CharacterName:            owner.CharacterName,
		PosX:                     sh.PosX,
		PosY:                     sh.PosY,
		SystemID:                 sh.SystemID,
		Texture:                  sh.Texture,
		Theta:                    sh.Theta,
		VelX:                     sh.VelX,
		VelY:                     sh.VelY,
		Shield:                   sh.Shield,
		Armor:                    sh.Armor,
		Hull:                     sh.Hull,
		Fuel:                     sh.Fuel,
		Heat:                     sh.Heat,
		Energy:                   sh.Energy,
		DockedAtStationID:        sh.DockedAtStationID,
		Fitting:                  *fitting,
		Destroyed:                sh.Destroyed,
		DestroyedAt:              sh.DestroyedAt,
		CargoBayContainerID:      sh.CargoBayContainerID,
		CargoBay:                 cargoBay,
		FittingBayContainerID:    sh.FittingBayContainerID,
		FittingBay:               fittingBay,
		ReMaxDirty:               sh.ReMaxDirty,
		TrashContainerID:         sh.TrashContainerID,
		Wallet:                   sh.Wallet,
		BeingFlownByPlayer:       owner.CurrentShipID != nil && (sh.ID == *owner.CurrentShipID),
		FlightExperienceModifier: sh.FlightExperienceModifier,
		TemplateData: universe.ShipTemplate{
			ID:                 temp.ID,
			Created:            temp.Created,
			ShipTemplateName:   temp.ShipTemplateName,
			WreckTexture:       temp.WreckTexture,
			ExplosionTexture:   temp.ExplosionTexture,
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
			ItemTypeID:         temp.ItemTypeID,
			CanUndock:          temp.CanUndock,
		},
		FactionID:     owner.CurrentFactionID,
		IsNPC:         owner.IsNPC,
		BehaviourMode: owner.BehaviourMode,
		Aggressors:    make(map[string]*shared.PlayerReputationSheet),
		AggressionLog: make(map[string]*shared.AggressionLog),
		Lock:          sync.Mutex{},
	}

	// obtain factions read lock
	u.FactionsLock.RLock()
	defer u.FactionsLock.RUnlock()

	// load and inject reputation sheet
	repSheet := LoadReputationSheet(owner)
	es.ReputationSheet = repSheet

	// inject faction
	es.Faction = u.Factions[owner.CurrentFactionID.String()]

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

	// associate escrow container with ship
	sp.EscrowContainerID = &owner.EscrowContainerID

	// hook cargo bay schematics into running jobs
	hookSchematics(sp)

	// calculate initial stat caches
	universe.RecalcAllStatCaches(sp)

	// return pointer to ship
	return sp, nil
}

// Takes a SQL Outpost and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadOutpost(sh *sql.Outpost, u *universe.Universe) (*universe.Outpost, error) {
	// get template
	temp, err := outpostTemplateSvc.GetOutpostTemplateByID(sh.OutpostTemplateId)

	if err != nil {
		return nil, err
	}

	// get owner info
	owner, err := userSvc.GetUserByID((sh.UserID))

	if err != nil {
		return nil, err
	}

	// build in-memory outpost
	es := universe.Outpost{
		ID:            sh.ID,
		UserID:        sh.UserID,
		Created:       sh.Created,
		OutpostName:   sh.OutpostName,
		CharacterName: owner.CharacterName,
		PosX:          sh.PosX,
		PosY:          sh.PosY,
		SystemID:      sh.SystemID,
		Theta:         sh.Theta,
		Shield:        sh.Shield,
		Armor:         sh.Armor,
		Hull:          sh.Hull,
		Destroyed:     sh.Destroyed,
		DestroyedAt:   sh.DestroyedAt,
		Wallet:        sh.Wallet,
		TemplateData: universe.OutpostTemplate{
			ID:                  temp.ID,
			Created:             temp.Created,
			OutpostTemplateName: temp.OutpostTemplateName,
			WreckTexture:        temp.WreckTexture,
			ExplosionTexture:    temp.ExplosionTexture,
			Texture:             temp.Texture,
			Radius:              temp.Radius,
			BaseMass:            temp.BaseMass,
			BaseShield:          temp.BaseShield,
			BaseShieldRegen:     temp.BaseShieldRegen,
			BaseArmor:           temp.BaseArmor,
			BaseHull:            temp.BaseHull,
			ItemTypeID:          temp.ItemTypeID,
		},
		FactionID: owner.CurrentFactionID,
		Lock:      sync.Mutex{},
	}

	// obtain factions read lock
	u.FactionsLock.RLock()
	defer u.FactionsLock.RUnlock()

	// inject faction
	es.Faction = u.Factions[owner.CurrentFactionID.String()]

	// get pointer to outpost
	sp := &es

	// return pointer to ship
	return sp, nil
}

// Hooks pointers needed if a schematic is currently running
func hookSchematics(sp *universe.Ship) {
	// get schematic runs for ship owner
	runs := getSchematicRunsByUser(sp.UserID)

	for _, r := range runs {
		// obtain lock
		r.Lock.Lock()
		defer r.Lock.Unlock()

		// skip if initialized
		if r.Initialized {
			continue
		}

		// hook references
		for _, ci := range sp.CargoBay.Items {
			if ci.ID == r.SchematicItemID {
				r.SchematicItem = ci
				r.Process = ci.Process
				r.Ship = sp

				break
			}
		}

		// check if all hooked
		if r.SchematicItem != nil && r.Process != nil && r.Ship != nil {
			// mark as initialized
			r.Initialized = true

			if r.StatusID != "delivered" {
				// mark schematic item as in use
				r.SchematicItem.SchematicInUse = true
			}
		}
	}
}

// Generates a kill log for a dead ship
func generateKillLog(ship *universe.Ship) *sql.KillLog {
	// null check
	if ship == nil || ship.Faction == nil || ship.CurrentSystem == nil || ship.DestroyedAt == nil {
		return nil
	}

	// fill base structure
	log := sql.KillLog{
		Header: sql.KillLogHeader{
			VictimID:               ship.UserID,
			VictimName:             ship.CharacterName,
			VictimFactionID:        ship.FactionID,
			VictimFactionName:      ship.Faction.Name,
			VictimShipTemplateID:   ship.TemplateData.ID,
			VictimShipTemplateName: ship.TemplateData.ShipTemplateName,
			VictimShipID:           ship.ID,
			VictimShipName:         ship.ShipName,
			Timestamp:              *ship.DestroyedAt,
			SolarSystemID:          ship.CurrentSystem.ID,
			SolarSystemName:        ship.CurrentSystem.SystemName,
			RegionID:               ship.CurrentSystem.RegionID,
			RegionName:             ship.CurrentSystem.RegionName,
			InvolvedParties:        len(ship.AggressionLog),
			IsNPC:                  ship.IsNPC,
			HoldingFactionID:       ship.CurrentSystem.HoldingFactionID,
			HoldingFactionName:     ship.CurrentSystem.HoldingFactionName,
			PosX:                   ship.PosX,
			PosY:                   ship.PosY,
		},
		Fitting: sql.KillLogFitting{
			ARack: []sql.KillLogSlot{},
			BRack: []sql.KillLogSlot{},
			CRack: []sql.KillLogSlot{},
		},
		Cargo:           []sql.KillLogCargoItem{},
		InvolvedParties: []sql.KillLogInvolvedParty{},
		Wallet:          int64(ship.Wallet),
	}

	// fill slots
	for _, e := range ship.Fitting.ARack {
		m, _ := e.ItemMeta.GetBool("**MODIFIED**")
		cf, _ := e.ItemMeta.GetInt("customization_factor")

		log.Fitting.ARack = append(log.Fitting.ARack, sql.KillLogSlot{
			ItemID:              e.ItemID,
			ItemTypeID:          e.ItemTypeID,
			ItemFamilyID:        e.ItemTypeFamily,
			ItemTypeName:        e.ItemTypeName,
			ItemFamilyName:      e.ItemTypeFamilyName,
			IsModified:          m,
			CustomizationFactor: cf,
		})
	}

	for _, e := range ship.Fitting.BRack {
		m, _ := e.ItemMeta.GetBool("**MODIFIED**")
		cf, _ := e.ItemMeta.GetInt("customization_factor")

		log.Fitting.BRack = append(log.Fitting.BRack, sql.KillLogSlot{
			ItemID:              e.ItemID,
			ItemTypeID:          e.ItemTypeID,
			ItemFamilyID:        e.ItemTypeFamily,
			ItemTypeName:        e.ItemTypeName,
			ItemFamilyName:      e.ItemTypeFamilyName,
			IsModified:          m,
			CustomizationFactor: cf,
		})
	}

	for _, e := range ship.Fitting.CRack {
		m, _ := e.ItemMeta.GetBool("**MODIFIED**")
		cf, _ := e.ItemMeta.GetInt("customization_factor")

		log.Fitting.CRack = append(log.Fitting.CRack, sql.KillLogSlot{
			ItemID:              e.ItemID,
			ItemTypeID:          e.ItemTypeID,
			ItemFamilyID:        e.ItemTypeFamily,
			ItemTypeName:        e.ItemTypeName,
			ItemFamilyName:      e.ItemTypeFamilyName,
			IsModified:          m,
			CustomizationFactor: cf,
		})
	}

	// fill cargo
	for _, e := range ship.CargoBay.Items {
		log.Cargo = append(log.Cargo, sql.KillLogCargoItem{
			ItemID:         e.ID,
			ItemTypeID:     e.ItemTypeID,
			ItemFamilyID:   e.ItemFamilyID,
			ItemTypeName:   e.ItemTypeName,
			ItemFamilyName: e.ItemFamilyName,
			Quantity:       e.Quantity,
			IsPackaged:     e.IsPackaged,
		})
	}

	// fill involved parties
	for _, p := range ship.AggressionLog {
		// core of structure
		pt := sql.KillLogInvolvedParty{
			UserID:              p.UserID,
			FactionID:           p.FactionID,
			CharacterName:       p.CharacterName,
			FactionName:         p.FactionName,
			IsNPC:               p.IsNPC,
			LastAggressed:       p.LastAggressed,
			ShipID:              p.ShipID,
			ShipName:            p.ShipName,
			ShipTemplateID:      p.ShipTemplateID,
			ShipTemplateName:    p.ShipTemplateName,
			LastSolarSystemID:   p.LastSolarSystemID,
			LastSolarSystemName: p.LastSolarSystemName,
			LastRegionID:        p.LastRegionID,
			LastRegionName:      p.LastRegionName,
			LastPosX:            p.LastPosX,
			LastPosY:            p.LastPosY,
			WeaponUse:           map[string]*sql.KillLogWeaponUse{},
		}

		// fill weapon usages
		for _, w := range p.WeaponUse {
			pt.WeaponUse[w.ItemID.String()] = &sql.KillLogWeaponUse{
				ItemID:          w.ItemID,
				ItemTypeID:      w.ItemTypeID,
				ItemFamilyID:    w.ItemFamilyID,
				ItemFamilyName:  w.ItemFamilyName,
				ItemTypeName:    w.ItemTypeName,
				LastUsed:        w.LastUsed,
				DamageInflicted: w.DamageInflicted,
			}
		}

		// store in log
		log.InvolvedParties = append(log.InvolvedParties, pt)
	}

	// return kill log
	return &log
}

// Updates a ship in the database
func saveShip(ship *universe.Ship) error {
	// obtain lock on copy
	ship.Lock.Lock()
	defer ship.Lock.Unlock()

	dbShip := sql.Ship{
		ID:                       ship.ID,
		UserID:                   ship.UserID,
		Created:                  ship.Created,
		ShipName:                 ship.ShipName,
		PosX:                     ship.PosX,
		PosY:                     ship.PosY,
		SystemID:                 ship.SystemID,
		Texture:                  ship.Texture,
		Theta:                    ship.Theta,
		VelX:                     ship.VelX,
		VelY:                     ship.VelY,
		Shield:                   ship.Shield,
		Armor:                    ship.Armor,
		Hull:                     ship.Hull,
		Fuel:                     ship.Fuel,
		Heat:                     ship.Heat,
		Energy:                   ship.Energy,
		ShipTemplateID:           ship.TemplateData.ID,
		DockedAtStationID:        ship.DockedAtStationID,
		Fitting:                  SQLFromFitting(&ship.Fitting),
		Destroyed:                ship.Destroyed,
		DestroyedAt:              ship.DestroyedAt,
		CargoBayContainerID:      ship.CargoBayContainerID,
		FittingBayContainerID:    ship.FittingBayContainerID,
		TrashContainerID:         ship.TrashContainerID,
		ReMaxDirty:               ship.ReMaxDirty,
		Wallet:                   ship.Wallet,
		FlightExperienceModifier: ship.FlightExperienceModifier,
	}

	err := shipSvc.UpdateShip(dbShip)

	if err != nil {
		shared.TeeLog(fmt.Sprintf("Error saving ship: %v | %v", dbShip, err))
	}

	return err
}

// Moves an item to a different container in the database
func saveItemLocation(itemID uuid.UUID, containerID uuid.UUID) error {
	return itemSvc.SetContainerID(itemID, containerID)
}

// Marks an item as packaged in the database
func packageItem(itemID uuid.UUID) error {
	return itemSvc.PackageItem(itemID)
}

// Marks an item as unpackaged in the database
func unpackageItem(itemID uuid.UUID, meta universe.Meta) error {
	return itemSvc.UnpackageItem(itemID, SQLFromMeta(&meta))
}

// Changes the quantity of an item stack in the database
func changeQuantity(itemID uuid.UUID, quantity int) error {
	return itemSvc.ChangeQuantity(itemID, quantity)
}

// Changes the metadata of an item in the database
func changeMeta(itemID uuid.UUID, meta universe.Meta) error {
	return itemSvc.ChangeMeta(itemID, sql.Meta(meta))
}

// Saves a new item to the database
func newItem(item *universe.Item) (*uuid.UUID, error) {
	// convert to sql type
	sql := SQLFromItem(item)

	if sql == nil {
		return nil, errors.New("error converting item to SQL type")
	}

	// save item
	o, err := itemSvc.NewItem(*sql)

	if err != nil || o == nil {
		return nil, err
	}

	return &o.ID, err
}

// Saves a new sell order to the database
func newSellOrder(sellOrder *universe.SellOrder) (*uuid.UUID, error) {
	// convert to sql type
	sql := SQLFromSellOrder(sellOrder)

	if sql == nil {
		return nil, errors.New("error converting sell order to SQL type")
	}

	// save sell order
	o, err := sellOrderSvc.NewSellOrder(*sql)

	if err != nil || o == nil {
		return nil, err
	}

	return &o.ID, err
}

// Saves a new sell order to the database
func markSellOrderAsBought(sellOrder *universe.SellOrder) error {
	// convert to sql type
	sql := SQLFromSellOrder(sellOrder)

	if sql == nil {
		return errors.New("error converting sell order to SQL type")
	}

	// save sell order
	err := sellOrderSvc.MarkSellOrderAsBought(*sql)

	return err
}
