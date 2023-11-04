package model

type Room struct {
	Id                 int               `db:"id"`
	OwnerId            int               `db:"owner_id"`
	CurrentQueueItemId *int              `db:"current_queue_item_id"`
	Name               string            `db:"name"`
	ShareCode          *string           `db:"share_code"`
	PlaybackOrderType  PlaybackOrderType `db:"playback_order_type"`
}
