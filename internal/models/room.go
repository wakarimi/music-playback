package models

type PlaybackOrderType string

const (
	InOrder PlaybackOrderType = "IN_ORDER"
	Replay  PlaybackOrderType = "REPLAY"
	Random  PlaybackOrderType = "RANDOM"
)

type Room struct {
	Id                 int               `db:"id"`
	OwnerId            int               `db:"owner_id"`
	CurrentQueueItemId *int              `db:"current_queue_item_id"`
	Name               string            `db:"name"`
	ConnectionString   *string           `db:"connection_string"`
	PlaybackOrderType  PlaybackOrderType `db:"playback_order_type"`
}
