package room_repo

import (
	"github.com/jmoiron/sqlx"
	"music-playback/internal/model"
)

type Repo interface {
	Create(tx *sqlx.Tx, room model.Room) (roomID int, err error)
	Read(tx *sqlx.Tx, roomID int) (room model.Room, err error)
	ReadAllByAccount(tx *sqlx.Tx, accountID int) (rooms []model.Room, err error)
	UpdateName(tx *sqlx.Tx, roomID int, name string) (err error)
	Delete(tx *sqlx.Tx, roomID int) (err error)
	IsExists(tx *sqlx.Tx, roomID int) (exists bool, err error)
	ReadByShareCode(tx *sqlx.Tx, shareCode string) (room model.Room, err error)
	IsExistsByShareCode(tx *sqlx.Tx, shareCode string) (exists bool, err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
