package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) ReadByRoom(tx *sqlx.Tx, roomID int) (shareCode model.ShareCode, err error) {
	log.Debug().Int("roomID", roomID).Msg("Reading share code")

	query := `
		SELECT *
		FROM share_codes
		WHERE room_id = :room_id
	`
	args := map[string]interface{}{
		"room_id": roomID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomID", roomID).Msg("Failed to prepare query")
		return model.ShareCode{}, err
	}
	err = stmt.Get(&shareCode, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to read share code")
		return model.ShareCode{}, err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code read successfully")
	return shareCode, nil
}
