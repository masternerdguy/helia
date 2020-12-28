package engine

import (
	"errors"
	"fmt"
	"helia/physics"
	"helia/sql"
	"time"

	"github.com/google/uuid"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//CreateNoobShipForPlayer Creates a noob (starter) ship for a player and put them in it
func CreateNoobShipForPlayer(start *sql.Start, uid uuid.UUID) (*sql.User, error) {
	const moduleCreationReason = "Module for new noob ship for player"

	//safety check
	if start == nil {
		return nil, errors.New("no start provided")
	}

	//get services
	userSvc := sql.GetUserService()
	itemSvc := sql.GetItemService()
	shipSvc := sql.GetShipService()
	shipTmpSvc := sql.GetShipTemplateService()
	containerSvc := sql.GetContainerService()

	//get user by id
	u, err := userSvc.GetUserByID(uid)

	if err != nil {
		return u, err
	}

	//get ship template from start
	temp, err := shipTmpSvc.GetShipTemplateByID(start.ShipTemplateID)

	if err != nil {
		return u, err
	}

	//create container for trashed items
	tb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	//create container for cargo bay
	cb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	//create container for fitting bay
	fb, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		return u, err
	}

	//initialize fitting
	fitting := sql.Fitting{
		ARack: []sql.FittedSlot{},
		BRack: []sql.FittedSlot{},
		CRack: []sql.FittedSlot{},
	}

	//create rack a modules
	for _, l := range start.ShipFitting.ARack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
		})

		if err != nil {
			return u, err
		}

		//link item to slot
		fitting.ARack = append(fitting.ARack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create rack b modules
	for _, l := range start.ShipFitting.BRack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
		})

		if err != nil {
			return u, err
		}

		//link item to slot
		fitting.BRack = append(fitting.BRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create rack c modules
	for _, l := range start.ShipFitting.CRack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
			ContainerID:   fb.ID,
			Quantity:      1,
		})

		if err != nil {
			return u, err
		}

		//link item to slot
		fitting.CRack = append(fitting.CRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create starter ship
	t := sql.Ship{
		SystemID:              start.SystemID,
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
		PosX:                  float64(physics.RandInRange(-50000, 50000)),
		PosY:                  float64(physics.RandInRange(-50000, 50000)),
		Fitting:               fitting,
		Destroyed:             false,
		CargoBayContainerID:   cb.ID,
		FittingBayContainerID: fb.ID,
		ReMaxDirty:            true,
		TrashContainerID:      tb.ID,
	}

	starterShip, err := shipSvc.NewShip(t)

	if err != nil {
		return u, err
	}

	//put user in starter ship
	err = userSvc.SetCurrentShipID(u.ID, &starterShip.ID)
	u.CurrentShipID = &starterShip.ID

	if err != nil {
		return u, err
	}

	//success!
	return u, nil
}
