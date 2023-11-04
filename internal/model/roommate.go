package model

type Roommate struct {
	ID        int `db:"id"`
	AccountID int `db:"account_id"`
	RoomID    int `db:"room_id"`
}
