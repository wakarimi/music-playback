package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) Delete(tx *sqlx.Tx, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Deleting room")

	room, err := s.Get(tx, roomID, accountID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get room")
		return err
	}

	if room.OwnerID != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("room %d is not owned by account %d", roomID, accountID)}
		log.Error().Err(err).Int("roomID", roomID).Int("accountID", accountID).Msg("Room is not owned by account")
		return err
	}

	err = s.RoomRepo.Delete(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to delete room")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Room deleted")
	return nil
}
