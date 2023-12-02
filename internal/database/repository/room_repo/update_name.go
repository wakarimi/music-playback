package room_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) UpdateName(tx *sqlx.Tx, roomID int, name string) (err error) {
	log.Debug().Int("roomID", roomID).Msg("Updating room name")

	query := `
		UPDATE rooms
		SET name = :name
		WHERE id = :id
	`
	args := map[string]interface{}{
		"name": name,
		"id":   roomID,
	}

	result, err := tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to update room name")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get rows affected after room update")
		return err
	}
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows affected while updating room name")
		log.Error().Err(err).Int("roomID", roomID).Msg("No rows affected while updating room name")
		return err
	}

	log.Debug().Int("roomID", roomID).Msg("Room name updated")
	return err
}
