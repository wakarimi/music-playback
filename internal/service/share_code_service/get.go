package share_code_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/model"
)

func (s Service) Get(tx *sqlx.Tx, roomID int, accountID int) (shareCode model.ShareCode, err error) {
	log.Debug().Int("roomID", roomID).Msg("Getting share code")

	room, err := s.RoomService.Get(tx, roomID, accountID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to read room")
		return model.ShareCode{}, err
	}

	if room.OwnerID != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("account %d does not have enough rights to room %d", accountID, roomID)}
		log.Error().Err(err).Int("accountID", accountID).Int("roomID", roomID).Msg("The account does not have enough permissions for the room")
		return model.ShareCode{}, err
	}

	exists, err := s.IsExistsByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check share code existence")
		return model.ShareCode{}, err
	}
	if !exists {
		err = errors.NotFound{Resource: fmt.Sprintf("share code for room with id=%d", roomID)}
		log.Error().Err(err).Int("roomID", roomID).Msg("Share code for room not found")
		return model.ShareCode{}, err
	}

	shareCode, err = s.ShareCodeRepo.ReadByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get share code")
		return model.ShareCode{}, err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code got")
	return shareCode, nil
}
