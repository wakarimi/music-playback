package model

type Session struct {
	Id       int `db:"id"`
	DeviceId int `db:"device_id"`
	RoomId   int `db:"room_id"`
}
