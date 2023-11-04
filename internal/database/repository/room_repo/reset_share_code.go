package room_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) ResetShareCode(tx *sqlx.Tx, roomID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Resetting share code")

	query := `
		UPDATE rooms
		SET share_code = NULL
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": roomID,
	}

	result, err := tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to reset share code")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get rows affected after share code reset")
		return err
	}
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows affected while updating share code")
		log.Error().Err(err).Int("roomID", roomID).Msg("No rows affected while resetting share code")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code reset")
	return err
}
