package universe

import (
	"helia/shared"
	"time"

	"github.com/google/uuid"
)

// Structure representing a sell order in the game universe
type SellOrder struct {
	ID           uuid.UUID
	StationID    uuid.UUID
	ItemID       uuid.UUID
	SellerUserID uuid.UUID
	AskPrice     float64
	Created      time.Time
	Bought       *time.Time
	BuyerUserID  *uuid.UUID
	// in-memory only
	Lock              shared.LabeledMutex
	Item              *Item
	CoreDirty         bool
	CoreWait          int
	GetItemIDFromItem bool
}
