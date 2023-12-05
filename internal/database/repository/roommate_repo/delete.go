package roommate_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Delete(tx *sqlx.Tx, roomId int, accountId int) (err error) {
	log.Debug().Msg("Deleting room and related entities")

	_, err = tx.Exec(`DELETE FROM roommates WHERE room_id = $1 AND account_id = $2`, roomId, accountId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete from roommates")
		return err
	}

	_, err = tx.Exec(`DELETE FROM sessions WHERE room_id = $1`, roomId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete from sessions")
		return err
	}

	log.Debug().Msg("Room and related entities deleted successfully")
	return nil
}
