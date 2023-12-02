package model

type ShareCode struct {
	ID     int    `db:"id"`
	RoomID int    `db:"room_id"`
	Code   string `db:"code"`
}
