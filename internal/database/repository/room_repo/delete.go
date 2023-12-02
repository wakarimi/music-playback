package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Delete(tx *sqlx.Tx, roomID int) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Deleting room")

	query := `
		DELETE FROM rooms
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": roomID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to delete room")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Room deleted")
	return nil
}
