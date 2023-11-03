package models

type Roommate struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
	RoomId int `db:"room_id"`
}
