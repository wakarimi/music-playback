package share_code_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) CreateOrUpdate(tx *sqlx.Tx, roomID int, accountID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Setting share code")

	exists, err := s.IsExistsByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check share code existence")
		return err
	}

	shareCodeStr, err := s.GenerateShareCode(tx)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to generate share code")
		return err
	}

	shareCode := model.ShareCode{
		Code: shareCodeStr,
	}

	if exists {
		err = s.Update(tx, shareCode, roomID, accountID)
		if err != nil {
			log.Error().Err(err).Int("roomID", roomID).Msg("Failed to update share code")
			return err
		}
	} else {
		err = s.Create(tx, shareCode, roomID, accountID)
		if err != nil {
			log.Error().Err(err).Int("roomID", roomID).Msg("Failed to create share code")
			return err
		}
	}

	log.Debug().Int("roomID", roomID).Msg("Share code generated and set")
	return nil
}
