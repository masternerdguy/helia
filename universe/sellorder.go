package universe

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

//SellOrder Structure representing a sell order in the game universe
type SellOrder struct {
	ID           uuid.UUID
	StationID    uuid.UUID
	ItemID       uuid.UUID
	SellerUserID uuid.UUID
	AskPrice     float64
	Created      time.Time
	Bought       *time.Time
	BuyerUserID  *uuid.UUID
	//in-memory only
	Lock      sync.Mutex
	Item      *Item
	CoreDirty bool
}
