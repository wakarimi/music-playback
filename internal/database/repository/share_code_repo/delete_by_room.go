package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) DeleteByRoom(tx *sqlx.Tx, roomID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Deleting share code")

	query := `
		DELETE FROM share_codes
		WHERE room_id = :room_id
	`
	args := map[string]interface{}{
		"room_id": roomID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to delete share code")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Share code deleted")
	return nil
}
