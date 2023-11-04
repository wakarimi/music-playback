package room_repo

import (
	"github.com/jmoiron/sqlx"
	"music-playback/internal/model"
)

type Repo interface {
	Create(tx *sqlx.Tx, room model.Room) (roomId int, err error)
	Read(tx *sqlx.Tx, roomId int) (room model.Room, err error)
	UpdateShareCode(tx *sqlx.Tx, roomId int, shareCode string) (err error)
	IsExists(tx *sqlx.Tx, roomId int) (exists bool, err error)
	IsShareCodeUsed(tx *sqlx.Tx, shareCode string) (used bool, err error)
	ResetShareCode(tx *sqlx.Tx, roomId int) (err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
