package share_code_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/model"
)

func (s Service) Update(tx *sqlx.Tx, shareCode model.ShareCode, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Updating share code")

	room, err := s.RoomService.Get(tx, roomID, accountID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to read room")
		return err
	}

	if room.OwnerID != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("room %d is not owned by account %d", roomID, accountID)}
		log.Error().Err(err).Int("roomID", roomID).Int("accountID", accountID).Msg("Room is not owned by account")
		return err
	}

	err = s.ShareCodeRepo.UpdateByRoom(tx, shareCode, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to create share code")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code created")
	return nil
}
