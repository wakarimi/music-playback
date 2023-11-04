package room_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) UpdateShareCode(tx *sqlx.Tx, roomID int, shareCode string) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Updating share code")

	query := `
		UPDATE rooms
		SET share_code = :share_code
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id":         roomID,
		"share_code": shareCode,
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
