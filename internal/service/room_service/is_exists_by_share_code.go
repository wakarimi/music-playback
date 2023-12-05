package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExistsByShareCode(tx *sqlx.Tx, shareCode string) (exists bool, err error) {
	log.Debug().Msg("Checking room existence by share code")

	exists, err = s.RoomRepo.IsExistsByShareCode(tx, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check room existence by share code")
		return false, err
	}

	log.Debug().Bool("exists", exists).Msg("Room existence by share code checked")
	return exists, nil
}
