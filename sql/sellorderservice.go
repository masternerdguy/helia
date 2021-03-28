package sql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Facility for interacting with the SellOrders table in the database
type SellOrderService struct{}

// Returns a SellOrder service for interacting with SellOrders in the database
func GetSellOrderService() SellOrderService {
	return SellOrderService{}
}

// Structure representing a row in the SellOrders table
type SellOrder struct {
	ID           uuid.UUID
	StationID    uuid.UUID
	ItemID       uuid.UUID
	SellerUserID uuid.UUID
	AskPrice     float64
	Created      time.Time
	Bought       *time.Time
	BuyerUserID  *uuid.UUID
}

// Finds and returns a sell order by its id
func (s SellOrderService) GetSellOrderByID(SellOrderID uuid.UUID) (*SellOrder, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//find sell order with this id
	sellOrder := SellOrder{}

	sqlStatement :=
		`
			SELECT id, universe_stationid, itemid, seller_userid, askprice, created, bought, buyer_userid
			FROM public.SellOrders
			WHERE id = $1
		`

	row := db.QueryRow(sqlStatement, SellOrderID)

	switch err := row.Scan(&sellOrder.ID, &sellOrder.StationID, &sellOrder.ItemID, &sellOrder.SellerUserID,
		&sellOrder.AskPrice, &sellOrder.Created, &sellOrder.Bought, &sellOrder.BuyerUserID); err {
	case sql.ErrNoRows:
		return nil, errors.New("SellOrder not found")
	case nil:
		return &sellOrder, nil
	default:
		return nil, err
	}
}

// Retrieves all sell orders at a given station that have not been bought yet
func (s SellOrderService) GetOpenSellOrdersByStation(containerID uuid.UUID) ([]SellOrder, error) {
	sellOrders := make([]SellOrder, 0)

	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//load sell orders
	sql :=
		`
			SELECT id, universe_stationid, itemid, seller_userid, askprice, created, bought, buyer_userid
			FROM public.SellOrders
			WHERE universe_stationid = $1
			AND bought is null
		`

	rows, err := db.Query(sql, containerID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := SellOrder{}

		//scan into sell order structure
		rows.Scan(&i.ID, &i.StationID, &i.ItemID, &i.SellerUserID,
			&i.AskPrice, &i.Created, &i.Bought, &i.BuyerUserID)

		//append to slice
		sellOrders = append(sellOrders, i)
	}

	return sellOrders, err
}

// Creates a new sell order
func (s SellOrderService) NewSellOrder(e SellOrder) (*SellOrder, error) {
	//get db handle
	db, err := connect()

	if err != nil {
		return nil, err
	}

	//insert sell order
	sql := `
			INSERT INTO public.sellorders(
				id, universe_stationid, itemid, seller_userid, askprice, created, bought, buyer_userid)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		   `

	uid := uuid.New()
	createdAt := time.Now()

	q, err := db.Query(sql, uid, e.StationID, e.ItemID, e.SellerUserID, e.AskPrice, e.Created, e.Bought, e.BuyerUserID)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	//update id in model
	e.ID = uid
	e.Created = createdAt

	//return pointer to inserted sell order model
	return &e, nil
}

// Updates the buyer and buyer_userid on a sell order
func (s SellOrderService) MarkSellOrderAsBought(e SellOrder) error {
	//get db handle
	db, err := connect()

	if err != nil {
		return err
	}

	//insert sell order
	sql := `
			UPDATE public.sellorders
			SET bought = $2, buyer_userid = $3
			WHERE id = $1
		   `

	q, err := db.Query(sql, e.ID, e.Bought, e.BuyerUserID)

	if err != nil {
		return err
	}

	defer q.Close()

	//success
	return nil
}
