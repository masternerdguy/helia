package engine

import (
	"errors"
	"fmt"
	"helia/sql"
	"time"

	"github.com/google/uuid"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Creates a noob (starter) ship for a player and put them in it
func CreatePurchasedShipForPlayer(ownerID uuid.UUID, templateID uuid.UUID, stationID uuid.UUID, systemID uuid.UUID) (*sql.Ship, error) {
	// get services
	userSvc := sql.GetUserService()
	shipSvc := sql.GetShipService()
	shipTmpSvc := sql.GetShipTemplateService()
	containerSvc := sql.GetContainerService()

	// get user by id
	u, err := userSvc.GetUserByID(ownerID)

	if err != nil {
		return nil, err
	}

	// get ship template from start
	temp, err := shipTmpSvc.GetShipTemplateByID(templateID)

	if err != nil {
		return nil, err
	}

	// create container for trashed items
	tb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return nil, err
	}

	// create container for cargo bay
	cb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return nil, err
	}

	// create container for fitting bay
	fb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return nil, err
	}

	// initialize empty fitting
	fitting := sql.Fitting{
		ARack: []sql.FittedSlot{},
		BRack: []sql.FittedSlot{},
		CRack: []sql.FittedSlot{},
	}

	for range temp.SlotLayout.ASlots {
		fitting.ARack = append(fitting.ARack, sql.FittedSlot{})
	}

	for range temp.SlotLayout.BSlots {
		fitting.BRack = append(fitting.BRack, sql.FittedSlot{})
	}

	for range temp.SlotLayout.CSlots {
		fitting.CRack = append(fitting.CRack, sql.FittedSlot{})
	}

	// create ship
	t := sql.Ship{
		SystemID:              systemID,
		DockedAtStationID:     &stationID,
		UserID:                u.ID,
		ShipName:              fmt.Sprintf("%s's %s", u.Username, temp.Texture),
		Texture:               temp.Texture,
		Theta:                 0,
		Shield:                temp.BaseShield,
		Armor:                 temp.BaseArmor,
		Hull:                  temp.BaseHull,
		Fuel:                  temp.BaseFuel,
		Heat:                  0,
		Energy:                temp.BaseEnergy,
		ShipTemplateID:        temp.ID,
		Fitting:               fitting,
		Destroyed:             false,
		CargoBayContainerID:   cb.ID,
		FittingBayContainerID: fb.ID,
		ReMaxDirty:            true,
		TrashContainerID:      tb.ID,
		Wallet:                0,
	}

	newShip, err := shipSvc.NewShip(t)

	if err != nil {
		return nil, err
	}

	// success!
	return newShip, nil
}

// Creates a noob (starter) ship for a player and put them in it
func CreateNoobShipForPlayer(start *sql.Start, uid uuid.UUID) (*sql.User, error) {
	const moduleCreationReason = "Module for new noob ship for player"

	// get default uuid
	emptyUUID := uuid.UUID{}
	defaultUUID := emptyUUID.String()

	// safety check
	if start == nil {
		return nil, errors.New("no start provided")
	}

	// get services
	userSvc := sql.GetUserService()
	itemSvc := sql.GetItemService()
	itemTypeSvc := sql.GetItemTypeService()
	shipSvc := sql.GetShipService()
	shipTmpSvc := sql.GetShipTemplateService()
	containerSvc := sql.GetContainerService()

	// get user by id
	u, err := userSvc.GetUserByID(uid)

	if err != nil {
		return u, err
	}

	// get ship template from start
	temp, err := shipTmpSvc.GetShipTemplateByID(start.ShipTemplateID)

	if err != nil {
		return u, err
	}

	// create container for trashed items
	tb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	// create container for cargo bay
	cb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	// create container for fitting bay
	fb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	// initialize fitting
	fitting := sql.Fitting{
		ARack: []sql.FittedSlot{},
		BRack: []sql.FittedSlot{},
		CRack: []sql.FittedSlot{},
	}

	// create rack a modules
	for _, l := range start.ShipFitting.ARack {
		// skip if empty slot
		if l.ItemTypeID.String() == defaultUUID {
			// link empty slot and move on
			fitting.ARack = append(fitting.ARack, sql.FittedSlot{})
			continue
		}

		// load item type data
		itemType, err := itemTypeSvc.GetItemTypeByID(l.ItemTypeID)

		if err != nil {
			return nil, err
		}

		// create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          itemType.Meta,
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
			IsPackaged:    false,
		})

		if err != nil {
			return u, err
		}

		// link item to slot
		fitting.ARack = append(fitting.ARack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	// create rack b modules
	for _, l := range start.ShipFitting.BRack {
		// skip if empty slot
		if l.ItemTypeID.String() == defaultUUID {
			// link empty slot and move on
			fitting.BRack = append(fitting.BRack, sql.FittedSlot{})
			continue
		}

		// load item type data
		itemType, err := itemTypeSvc.GetItemTypeByID(l.ItemTypeID)

		if err != nil {
			return nil, err
		}

		// create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          itemType.Meta,
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
			IsPackaged:    false,
		})

		if err != nil {
			return u, err
		}

		// link item to slot
		fitting.BRack = append(fitting.BRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	// create rack c modules
	for _, l := range start.ShipFitting.CRack {
		// skip if empty slot
		if l.ItemTypeID.String() == defaultUUID {
			// link empty slot and move on
			fitting.CRack = append(fitting.CRack, sql.FittedSlot{})
			continue
		}

		// load item type data
		itemType, err := itemTypeSvc.GetItemTypeByID(l.ItemTypeID)

		if err != nil {
			return nil, err
		}

		// create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          itemType.Meta,
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
			IsPackaged:    false,
		})

		if err != nil {
			return u, err
		}

		// link item to slot
		fitting.CRack = append(fitting.CRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	// create starter ship
	t := sql.Ship{
		SystemID:              start.SystemID,
		DockedAtStationID:     &start.HomeStationID,
		UserID:                u.ID,
		ShipName:              fmt.Sprintf("%s's Starter Ship", u.Username),
		Texture:               temp.Texture,
		Theta:                 0,
		Shield:                temp.BaseShield,
		Armor:                 temp.BaseArmor,
		Hull:                  temp.BaseHull,
		Fuel:                  temp.BaseFuel,
		Heat:                  0,
		Energy:                temp.BaseEnergy,
		ShipTemplateID:        temp.ID,
		Fitting:               fitting,
		Destroyed:             false,
		CargoBayContainerID:   cb.ID,
		FittingBayContainerID: fb.ID,
		ReMaxDirty:            true,
		TrashContainerID:      tb.ID,
		Wallet:                start.Wallet,
	}

	starterShip, err := shipSvc.NewShip(t)

	if err != nil {
		return u, err
	}

	// put user in starter ship
	err = userSvc.SetCurrentShipID(u.ID, &starterShip.ID)
	u.CurrentShipID = &starterShip.ID

	if err != nil {
		return u, err
	}

	// success!
	return u, nil
}
