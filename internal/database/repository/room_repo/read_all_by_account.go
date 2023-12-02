package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (r Repository) ReadAllByAccount(tx *sqlx.Tx, accountID int) ([]model.Room, error) {
	log.Debug().Int("accountID", accountID).Msg("Reading rooms by account")

	query := `
		SELECT r.*
		FROM rooms r
		INNER JOIN roommates rm ON r.id = rm.room_id
		WHERE rm.account_id = :account_id
	`
	args := map[string]interface{}{
		"account_id": accountID,
	}

	var rooms []model.Room
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("accountID", accountID).Msg("Failed to prepare query")
		return nil, err
	}
	err = stmt.Select(&rooms, args)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Msg("Failed to read rooms")
		return nil, err
	}

	log.Debug().Int("accountID", accountID).Msg("Rooms read successfully")
	return rooms, nil
}
