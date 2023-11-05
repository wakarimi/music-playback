package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExists(tx *sqlx.Tx, roomID int) (exists bool, err error) {
	log.Debug().Int("roomID", roomID).Msg("Checking room existence")

	exists, err = s.RoomRepo.IsExists(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check room existence")
		return false, err
	}

	log.Debug().Bool("exists", exists).Msg("Room existence checked")
	return exists, nil
}
