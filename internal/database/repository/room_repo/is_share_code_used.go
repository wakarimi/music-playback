package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsShareCodeUsed(tx *sqlx.Tx, shareCode string) (used bool, err error) {
	log.Debug().Msg("Checking share code usage")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM rooms
			WHERE share_code = :share_code
		)
	`
	args := map[string]interface{}{
		"share_code": shareCode,
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
