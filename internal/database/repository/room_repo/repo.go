package room_repo

import (
	"github.com/jmoiron/sqlx"
	"music-playback/internal/model"
)

type Repo interface {
	Create(tx *sqlx.Tx, room model.Room) (roomId int, err error)
	Read(tx *sqlx.Tx, roomId int) (room model.Room, err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
