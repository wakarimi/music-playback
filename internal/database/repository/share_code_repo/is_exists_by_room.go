package share_code_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExistsByRoom(tx *sqlx.Tx, roomID int) (exists bool, err error) {
	log.Debug().Int("roomID", roomID).Msg("Checking share code existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM share_codes
			WHERE room_id = :room_id
		)
	`
	args := map[string]interface{}{
		"room_id": roomID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomID", roomID).Msg("Failed to prepare query")
		return false, err
	}
	err = stmt.Get(&exists, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check share code existence")
		return false, err
	}

	log.Debug().Int("roomID", roomID).Bool("exists", exists).Msg("Share code existence checked")
	return exists, nil
}
