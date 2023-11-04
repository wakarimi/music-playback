package model

type Roommate struct {
	Id        int `db:"id"`
	AccountId int `db:"account_id"`
	RoomId    int `db:"room_id"`
}
