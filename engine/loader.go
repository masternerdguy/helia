package engine

import (
	"errors"
	"fmt"
	"helia/sql"
	"helia/universe"
	"log"

	"github.com/google/uuid"
)

//loadUniverse Loads the state of the universe from the database
func loadUniverse() (*universe.Universe, error) {
	//get services
	regionSvc := sql.GetRegionService()
	systemSvc := sql.GetSolarSystemService()
	shipSvc := sql.GetShipService()
	starSvc := sql.GetStarService()
	planetSvc := sql.GetPlanetService()
	stationSvc := sql.GetStationService()
	jumpholeSvc := sql.GetJumpholeService()

	u := universe.Universe{}

	//for linking jumpholes later
	jhMap := make(map[string]*universe.Jumphole)
	sMap := make(map[string]*universe.SolarSystem)

	//load regions
	rs, err := regionSvc.GetAllRegions()

	if err != nil {
		return nil, err
	}

	regions := make(map[string]*universe.Region, 0)
	for _, e := range rs {
		//load basic region information
		r := universe.Region{
			ID:         e.ID,
			RegionName: e.RegionName,
		}

		//load systems in region
		ss, err := systemSvc.GetSolarSystemsByRegion(e.ID)

		if err != nil {
			return nil, err
		}

		systems := make(map[string]*universe.SolarSystem, 0)

		for _, f := range ss {
			s := universe.SolarSystem{
				ID:         f.ID,
				SystemName: f.SystemName,
				RegionID:   f.RegionID,
			}

			//initialize and store system
			s.Initialize()
			systems[s.ID.String()] = &s

			//link universe into system
			s.Universe = &u

			//for jumphole linking later
			sMap[s.ID.String()] = &s

			//load ships
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

			//load stars
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

			//load planets
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

			//load jumpholes
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

				//for jumphole linking later
				jhMap[j.ID.String()] = &jumphole
			}

			//load npc stations
			stations, err := stationSvc.GetStationsBySolarSystem(s.ID)

			if err != nil {
				return nil, err
			}

			for _, currStation := range stations {
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
					//link solar system into station
					CurrentSystem: &s,
				}

				s.AddStation(&station)
			}
		}

		//store and append region
		r.Systems = systems
		regions[r.ID.String()] = &r
	}

	//link jumpholes
	for _, j := range jhMap {
		//get out system
		o := sMap[j.OutSystemID.String()]

		//copy jumpholes
		jhs := o.CopyJumpholes()

		//find and link destination jumphole into jumphole
		for _, k := range jhs {
			if k.OutSystemID == j.SystemID {
				//get real jumphole pointer from map
				j.OutJumphole = jhMap[k.ID.String()]

				//link destination system into jumphole
				j.OutSystem = o
				break
			}
		}
	}

	//link regions into universe
	u.Regions = regions

	return &u, nil
}

//saveUniverse Saves the current state of dynamic entities in the simulation to the database
func saveUniverse(u *universe.Universe) {
	//get services
	stationSvc := sql.GetStationService()

	//iterate over systems
	for _, r := range u.Regions {
		for _, s := range r.Systems {
			//get ships in system
			ships := s.CopyShips()

			//save ships to database
			for _, ship := range ships {
				saveShip(ship)
			}

			//get npc stations in system
			stations := s.CopyStations()

			//save npc stations to database
			for _, station := range stations {
				//obtain lock on copy
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
				}

				err := stationSvc.UpdateStation(dbStation)

				if err != nil {
					log.Println(fmt.Sprintf("Error saving station: %v | %v", dbStation, err))
				}
			}
		}
	}
}

//SlotLayoutFromSQL Converts a SlotLayout from the SQL type to the engine type
func SlotLayoutFromSQL(value *sql.SlotLayout) universe.SlotLayout {
	//set up empty layout
	layout := universe.SlotLayout{}

	// null check
	if value == nil {
		// return empty layout
		return layout
	}

	//copy slot data into layout
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

	//return filled layout
	return layout
}

//SlotFromSQL Converts a Slot from the SQL type to the engine type
func SlotFromSQL(value *sql.Slot) universe.Slot {
	//set up empty slot
	slot := universe.Slot{}

	//null check
	if value == nil {
		// return empty slot
		return slot
	}

	//copy slot data
	slot.Family = value.Family
	slot.Volume = value.Volume
	slot.TexturePosition = value.TexturePosition

	//return filled slot
	return slot
}

//FittingFromSQL Converts a Fitting from the SQL type to the engine type
func FittingFromSQL(value *sql.Fitting) (*universe.Fitting, error) {
	//set up empty layout
	fitting := universe.Fitting{}

	// null check
	if value == nil {
		// return empty layout
		return &fitting, nil
	}

	//copy slot data into layout
	for _, v := range value.ARack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		fitting.ARack = append(fitting.ARack, *slot)
	}

	for _, v := range value.BRack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		fitting.BRack = append(fitting.BRack, *slot)
	}

	for _, v := range value.CRack {
		slot, err := FittedSlotFromSQL(&v)

		if err != nil {
			return nil, err
		}

		fitting.CRack = append(fitting.CRack, *slot)
	}

	//return filled layout
	return &fitting, nil
}

//FittedSlotFromSQL Converts a FittedSlot from the SQL type to the engine type
func FittedSlotFromSQL(value *sql.FittedSlot) (*universe.FittedSlot, error) {
	// get services
	itemSvc := sql.GetItemService()
	itemTypeSvc := sql.GetItemTypeService()

	//set up empty slot
	slot := universe.FittedSlot{}

	//null check
	if value == nil {
		// return empty slot
		return &slot, nil
	}

	//copy slot data
	slot.ItemID = value.ItemID
	slot.ItemTypeID = value.ItemTypeID

	//load item data
	item, err := itemSvc.GetItemByID(slot.ItemID)

	if err != nil {
		return nil, err
	}

	//load item type data
	itemType, err := itemTypeSvc.GetItemTypeByID(item.ItemTypeID)

	if err != nil {
		return nil, err
	}

	//store on slot
	slot.ItemMeta = MetaFromSQL(&item.Meta)
	slot.ItemTypeMeta = MetaFromSQL(&itemType.Meta)
	slot.ItemTypeFamily = itemType.Family
	slot.ItemTypeName = itemType.Name

	//return filled slot
	return &slot, nil
}

//SQLFromFitting Converts a Fitting from the engine type to the SQL type
func SQLFromFitting(value *universe.Fitting) sql.Fitting {
	//set up empty layout
	fitting := sql.Fitting{}

	// null check
	if value == nil {
		// return empty layout
		return fitting
	}

	//copy slot data into layout
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

	//return filled layout
	return fitting
}

//SQLFromFittedSlot Converts a FittedSlot from the engine type to the SQL type
func SQLFromFittedSlot(value *universe.FittedSlot) sql.FittedSlot {
	//set up empty slot
	slot := sql.FittedSlot{}

	//null check
	if value == nil {
		// return empty slot
		return slot
	}

	//copy slot data
	slot.ItemID = value.ItemID
	slot.ItemTypeID = value.ItemTypeID

	//return filled slot
	return slot
}

//MetaFromSQL Converts generic metadata from the SQL type to the engine type
func MetaFromSQL(value *sql.Meta) universe.Meta {
	//set up empty metadata
	meta := universe.Meta{}

	//null check
	if value == nil {
		// return empty slot
		return meta
	}

	//copy metadata
	for k, v := range *value {
		meta[k] = v
	}

	//return filled meta
	return meta
}

//SQLFromMeta Converts generic metadata from the SQL type to the engine type
func SQLFromMeta(value *universe.Meta) sql.Meta {
	//set up empty metadata
	meta := sql.Meta{}

	//null check
	if value == nil {
		// return empty slot
		return meta
	}

	//copy metadata
	for k, v := range *value {
		meta[k] = v
	}

	//return filled meta
	return meta
}

//StartFittingFromSQL Converts a StartFitting from the SQL type to the engine type
func StartFittingFromSQL(value *sql.StartFitting) universe.StartFitting {
	//set up empty layout
	layout := universe.StartFitting{}

	// null check
	if value == nil {
		// return empty layout
		return layout
	}

	//copy slot data into layout
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

	//return filled layout
	return layout
}

//StartFittedSlotFromSQL Converts a StartFittedSlot from the SQL type to the engine type
func StartFittedSlotFromSQL(value *sql.StartFittedSlot) universe.StartFittedSlot {
	//set up empty slot
	slot := universe.StartFittedSlot{}

	//null check
	if value == nil {
		// return empty slot
		return slot
	}

	//copy slot data
	slot.ItemTypeID = value.ItemTypeID

	//return filled slot
	return slot
}

//ContainerFromSQL Converts a Container from the SQL type to the engine type
func ContainerFromSQL(value *sql.Container) *universe.Container {
	//set up empty container
	container := universe.Container{}

	//null check
	if value == nil {
		//return nil
		return nil
	}

	//copy container data
	container.ID = value.ID
	container.Created = value.Created
	container.Meta = MetaFromSQL(&value.Meta)

	//return filled container
	return &container
}

//ItemFromSQL Converts an Item from the SQL type to the engine type
func ItemFromSQL(value *sql.Item) *universe.Item {
	//set up empty item
	item := universe.Item{}

	//null check
	if value == nil {
		//return nil
		return nil
	}

	//copy item data
	item.ID = value.ID
	item.ItemTypeID = value.ItemTypeID
	item.ContainerID = value.ContainerID
	item.Created = value.Created
	item.CreatedBy = value.CreatedBy
	item.CreatedReason = value.CreatedReason
	item.Meta = MetaFromSQL(&value.Meta)
	item.Quantity = value.Quantity
	item.IsPackaged = value.IsPackaged

	//return filled item
	return &item
}

//LoadItem Loads an item with some type and family data for use in the simulation.
func LoadItem(i *sql.Item) (*universe.Item, error) {
	itemTypeSvc := sql.GetItemTypeService()
	itemFamilySvc := sql.GetItemFamilyService()

	//load base item
	ei := ItemFromSQL(i)

	//null check
	if ei == nil {
		return nil, errors.New("Item from SQL failed")
	}

	//load item type
	it, err := itemTypeSvc.GetItemTypeByID(ei.ItemTypeID)

	if err != nil {
		return nil, err
	}

	//include item type data
	ei.ItemTypeName = it.Name
	ei.ItemFamilyID = it.Family

	//load item family
	fm, err := itemFamilySvc.GetItemFamilyByID(it.Family)

	if err != nil {
		return nil, err
	}

	//include item family data
	ei.ItemFamilyName = fm.FriendlyName

	//return filled item
	return ei, nil
}

//LoadContainer Takes a SQL Container and converts it, and items it loads, into the engine type ready for insertion into the universe.
func LoadContainer(c *sql.Container) (*universe.Container, error) {
	//load base container
	container := ContainerFromSQL(c)

	//load items
	itemSvc := sql.GetItemService()

	s, err := itemSvc.GetItemsByContainer(container.ID)

	if err != nil {
		return nil, err
	}

	//load items and push into container
	for _, i := range s {
		m, err := LoadItem(&i)

		if err != nil {
			return nil, err
		}

		//null check
		if m == nil {
			return nil, errors.New("Item argument was nil")
		}

		//push into container
		container.Items = append(container.Items, m)
	}

	//return full container
	return container, nil
}

//LoadShip Takes a SQL Ship and converts it, along with additional loaded data from the database, into the engine type ready for insertion into the universe.
func LoadShip(sh *sql.Ship) (*universe.Ship, error) {
	shipTmpSvc := sql.GetShipTemplateService()
	userSvc := sql.GetUserService()
	containerSvc := sql.GetContainerService()

	//get template
	temp, err := shipTmpSvc.GetShipTemplateByID(sh.ShipTemplateID)

	if err != nil {
		return nil, err
	}

	//get owner info
	owner, err := userSvc.GetUserByID((sh.UserID))

	if err != nil {
		return nil, err
	}

	//get fitting
	fitting, err := FittingFromSQL(&sh.Fitting)

	if err != nil {
		return nil, err
	}

	//load cargo bay
	cb, err := containerSvc.GetContainerByID(sh.CargoBayContainerID)

	if err != nil {
		return nil, err
	}

	cargoBay, err := LoadContainer(cb)

	if err != nil {
		return nil, err
	}

	//load fitting bay
	fb, err := containerSvc.GetContainerByID(sh.FittingBayContainerID)

	if err != nil {
		return nil, err
	}

	fittingBay, err := LoadContainer(fb)

	if err != nil {
		return nil, err
	}

	//build in-memory ship
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
	}

	//get pointer to ship
	sp := &es

	//link ship into fitting
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

	//return pointer to ship
	return sp, nil
}

//saveShip Updates a ship in the database
func saveShip(ship *universe.Ship) error {
	//get ship service
	shipSvc := sql.GetShipService()

	//obtain lock on copy
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
	}

	err := shipSvc.UpdateShip(dbShip)

	if err != nil {
		log.Println(fmt.Sprintf("Error saving ship: %v | %v", dbShip, err))
	}

	return err
}

//saveItemLocation Moves an item to a different container in the database
func saveItemLocation(itemID uuid.UUID, containerID uuid.UUID) error {
	itemSvc := sql.GetItemService()
	return itemSvc.SetContainerID(itemID, containerID)
}
