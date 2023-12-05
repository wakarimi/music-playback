package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) ReadByShareCode(tx *sqlx.Tx, shareCode string) (room model.Room, err error) {
	log.Debug().Str("shareCode", shareCode).Msg("Reading room by share code")

	query := `
        SELECT r.id, r.owner_id, r.current_queue_item_id, r.name, r.playback_order_type
        FROM rooms r
        INNER JOIN share_codes sc ON r.id = sc.room_id
        WHERE sc.code = $1
    `
	err = tx.Get(&room, query, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read room by share code")
		return model.Room{}, err
	}

	log.Debug().Msg("Room read by share code")
	return room, nil
}
