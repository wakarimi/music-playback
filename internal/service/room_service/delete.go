package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) Delete(tx *sqlx.Tx, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Getting share code")

	roomExists, err := s.RoomRepo.IsExists(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check room existence")
		return err
	}
	if !roomExists {
		err = errors.NotFound{Resource: fmt.Sprintf("room with id=%d", roomID)}
		log.Error().Err(err).Int("roomID", roomID).Msg("Room not found")
		return err
	}

	room, err := s.RoomRepo.Read(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to read room")
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
