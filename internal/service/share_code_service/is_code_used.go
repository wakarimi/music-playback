package share_code_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsCodeUsed(tx *sqlx.Tx, code string) (used bool, err error) {
	log.Debug().Msg("Checking share code usage")

	used, err = s.ShareCodeRepo.IsCodeUsed(tx, code)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check share code usage")
		return false, err
	}

	log.Debug().Bool("used", used).Msg("Share code usage checked")
	return used, nil
}
