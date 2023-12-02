package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"music-playback/internal/model"
)

type Repo interface {
	Create(tx *sqlx.Tx, shareCode model.ShareCode, roomID int) (err error)
	ReadByRoom(tx *sqlx.Tx, roomID int) (shareCode model.ShareCode, err error)
	UpdateByRoom(tx *sqlx.Tx, shareCode model.ShareCode, roomID int) (err error)
	DeleteByRoom(tx *sqlx.Tx, roomID int) (err error)
	IsExistsByRoom(tx *sqlx.Tx, roomID int) (exists bool, err error)
	IsCodeUsed(tx *sqlx.Tx, code string) (used bool, err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
