package roommate_repo

import (
	"github.com/jmoiron/sqlx"
	"music-playback/internal/model"
)

type Repo interface {
	Create(tx *sqlx.Tx, roommate model.Roommate) (roommateID int, err error)
	Delete(tx *sqlx.Tx, roomId int, accountId int) (err error)
	DeleteAllByRoom(tx *sqlx.Tx, roomId int) (err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
