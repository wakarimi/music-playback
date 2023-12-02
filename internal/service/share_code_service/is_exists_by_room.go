package share_code_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExistsByRoom(tx *sqlx.Tx, roomID int) (exists bool, err error) {
	log.Debug().Int("roomID", roomID).Msg("Checking share code usage")

	exists, err = s.ShareCodeRepo.IsExistsByRoom(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check share code usage")
		return false, err
	}

	log.Debug().Int("roomID", roomID).Bool("exists", exists).Msg("Share code usage checked")
	return exists, nil
}
