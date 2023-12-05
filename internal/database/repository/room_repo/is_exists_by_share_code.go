package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExistsByShareCode(tx *sqlx.Tx, shareCode string) (exists bool, err error) {
	log.Debug().Str("shareCode", shareCode).Msg("Checking if room exists by share code")

	query := `
        SELECT EXISTS (
            SELECT 1
            FROM share_codes
            WHERE code = $1
        )
    `
	err = tx.Get(&exists, query, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check if room exists by share code")
		return false, err
	}

	log.Debug().Msg("Checked existence of room by share code")
	return exists, nil
}
