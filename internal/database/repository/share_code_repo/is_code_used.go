package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsCodeUsed(tx *sqlx.Tx, code string) (used bool, err error) {
	log.Debug().Msg("Checking share code usage")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM share_codes
			WHERE code = :code
		)
	`
	args := map[string]interface{}{
		"code": code,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Msg("Failed to prepare query")
		return false, err
	}
	err = stmt.Get(&used, args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check share code usage")
		return false, err
	}

	log.Debug().Bool("used", used).Msg("Share code usage checked")
	return used, nil
}
