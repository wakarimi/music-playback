package room_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsRoommate(tx *sqlx.Tx, roomID int, accountID int) (isRoommate bool, err error) {
	log.Debug().Int("roomID", roomID).Int("accountID", accountID).Msg("Checking if user is a roommate")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM roommates
			WHERE room_id = :room_id
				AND account_id = :account_id
		)
	`
	args := map[string]interface{}{
		"room_id":    roomID,
		"account_id": accountID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("roomID", roomID).Msg("Failed to prepare query")
		return false, err
	}
	err = stmt.Get(&isRoommate, args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to execute query for IsRoommate")
		return false, err
	}

	log.Debug().Bool("isRoommate", isRoommate).Msg("Check complete")
	return isRoommate, nil
}
