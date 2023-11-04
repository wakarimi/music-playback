package model

type Room struct {
	ID                 int               `db:"id"`
	OwnerID            int               `db:"owner_id"`
	CurrentQueueItemID *int              `db:"current_queue_item_id"`
	Name               string            `db:"name"`
	ShareCode          *string           `db:"share_code"`
	PlaybackOrderType  PlaybackOrderType `db:"playback_order_type"`
}
