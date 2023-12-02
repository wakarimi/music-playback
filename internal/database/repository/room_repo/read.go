package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) Read(tx *sqlx.Tx, roomID int) (room model.Room, err error) {
	log.Debug().Int("roomID", roomID).Msg("Reading room")

	query := `
		SELECT *
		FROM rooms
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": roomID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomID", roomID).Msg("Failed to prepare query")
		return model.Room{}, err
	}
	err = stmt.Get(&room, args)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to read room")
		return model.Room{}, err
	}

	log.Debug().Int("roomID", roomID).Msg("Room read successfully")
	return room, nil
}
