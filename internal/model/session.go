package model

type Session struct {
	ID       int `db:"id"`
	DeviceID int `db:"device_id"`
	RoomID   int `db:"room_id"`
}
