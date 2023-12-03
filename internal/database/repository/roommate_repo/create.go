package roommate_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) Create(tx *sqlx.Tx, roommate model.Roommate) (roommateID int, err error) {
	log.Debug().Msg("Creating roommate")

	query := `
		INSERT INTO roommates(account_id, room_id)
		VALUES (:account_id, :room_id)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, roommate)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create roommate")
		return 0, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	if rows.Next() {
		if err := rows.Scan(&roommateID); err != nil {
			log.Error().Err(err).Msg("Failed to scan id into filed")
			return 0, err
		}
	} else {
		err := fmt.Errorf("no id returned after roommate insert")
		log.Error().Err(err).Msg("No id returned after roommate insert")
		return 0, err
	}

	log.Debug().Int("roommateId", roommateID).Msg("roommate created")
	return roommateID, nil
}
