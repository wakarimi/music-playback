package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) ResetShareCode(tx *sqlx.Tx, roomId int, accountID int) (err error) {
	log.Debug().Int("roomId", roomId).Msg("Resetting share code")

	roomExists, err := s.RoomRepo.IsExists(tx, roomId)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to check room existence")
		return err
	}
	if !roomExists {
		err = errors.NotFound{Resource: fmt.Sprintf("room with id=%d", roomId)}
		log.Error().Err(err).Int("roomId", roomId).Msg("Room not found")
		return err
	}

	room, err := s.RoomRepo.Read(tx, roomId)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to read room")
		return err
	}
	if room.OwnerId != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("room %d is not owned by account %d", roomId, accountID)}
		log.Error().Err(err).Int("roomId", roomId).Int("accountID", accountID).Msg("Room is not owned by account")
		return err
	}

	err = s.RoomRepo.ResetShareCode(tx, roomId)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to reset share code")
		return err
	}

	log.Debug().Int("roomId", roomId).Msg("Share code reset")
	return nil
}
