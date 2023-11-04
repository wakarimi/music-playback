package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExists(tx *sqlx.Tx, roomId int) (exists bool, err error) {
	log.Debug().Int("roomId", roomId).Msg("Checking room existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM rooms
			WHERE id = :id
		)
	`
	args := map[string]interface{}{
		"id": roomId,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomId", roomId).Msg("Failed to prepare query")
		return false, err
	}
	err = stmt.Get(&exists, args)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to check room existence")
		return false, err
	}

	log.Debug().Int("roomId", roomId).Bool("exists", exists).Msg("Room existence checked")
	return exists, nil
}
