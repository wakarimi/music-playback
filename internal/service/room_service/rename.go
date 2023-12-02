package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) Rename(tx *sqlx.Tx, roomID int, name string, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Renaming room")

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

	err = s.RoomRepo.UpdateName(tx, roomID, name)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to rename room")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Room renamed")
	return nil
}
