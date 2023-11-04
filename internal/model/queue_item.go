package model

type QueueItem struct {
	Id         int  `db:"id"`
	SongId     int  `db:"song_id"`
	PrevItemId *int `db:"prev_item_id"`
	NextItemId *int `db:"next_item_id"`
}
