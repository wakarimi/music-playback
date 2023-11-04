package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) Read(tx *sqlx.Tx, roomId int) (room model.Room, err error) {
	log.Debug().Int("roomId", roomId).Msg("Reading room")

	query := `
		SELECT *
		FROM rooms
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": roomId,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomId", roomId).Msg("Failed to prepare query")
		return model.Room{}, err
	}

	err = stmt.Get(&room, args)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to read room")
		return model.Room{}, err
	}

	log.Debug().Int("roomId", roomId).Msg("Room read successfully")
	return room, nil
}
