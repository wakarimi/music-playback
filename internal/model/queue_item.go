package model

type QueueItem struct {
	ID         int  `db:"id"`
	SongID     int  `db:"song_id"`
	PrevItemID *int `db:"prev_item_id"`
	NextItemID *int `db:"next_item_id"`
}
