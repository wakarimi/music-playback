package room_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) Create(tx *sqlx.Tx, room model.Room) (roomId int, err error) {
	log.Debug().Msg("Creating room")

	query := `
		INSERT INTO rooms(owner_id, current_queue_item_id, name, share_code, playback_order_type)
		VALUES (:owner_id, :current_queue_item_id, :name, :share_code, :playback_order_type)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, room)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create room")
		return 0, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	if rows.Next() {
		if err := rows.Scan(&roomId); err != nil {
			log.Error().Err(err).Msg("Failed to scan id into filed")
			return 0, err
		}
	} else {
		err := fmt.Errorf("no id returned after album insert")
		log.Error().Err(err).Msg("No id returned after album insert")
		return 0, err
	}

	log.Debug().Int("roomId", roomId).Msg("Room created")
	return roomId, nil
}
