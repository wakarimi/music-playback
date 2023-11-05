package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) Create(tx *sqlx.Tx, shareCode model.ShareCode, roomID int) (err error) {
	log.Debug().Msg("Creating share code")

	shareCode.RoomID = roomID
	query := `
		INSERT INTO share_codes(room_id, code)
		VALUES (:room_id, :code)
	`
	rows, err := tx.NamedQuery(query, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create share code")
		return err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	log.Debug().Msg("Share code created")
	return nil
}
