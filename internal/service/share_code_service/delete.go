package share_code_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) Delete(tx *sqlx.Tx, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Deleting share code")

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

	exists, err := s.IsExistsByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check share code existence")
		return err
	}
	if !exists {
		err = errors.NotFound{Resource: fmt.Sprintf("share code for room with id=%d", roomID)}
		log.Error().Err(err).Int("roomID", roomID).Msg("Share code for room not found")
		return err
	}

	err = s.ShareCodeRepo.DeleteByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to delete share code")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code deleted")
	return nil
}
