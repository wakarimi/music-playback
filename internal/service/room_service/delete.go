package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Delete(tx *sqlx.Tx, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Deleting room")

	err = s.RoomRepo.Delete(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to delete room")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Room deleted")
	return nil
}
