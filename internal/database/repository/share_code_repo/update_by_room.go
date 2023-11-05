package share_code_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) UpdateByRoom(tx *sqlx.Tx, shareCode model.ShareCode, roomID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Updating share code")

	query := `
		UPDATE share_codes
		SET code = :code
		WHERE room_id = :room_id
	`
	args := map[string]interface{}{
		"code":    shareCode.Code,
		"room_id": roomID,
	}

	result, err := tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to update share code")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get rows affected after share code update")
		return err
	}
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows affected while updating share code")
		log.Error().Err(err).Int("roomID", roomID).Msg("No rows affected while updating share code")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code updated")
	return err
}
