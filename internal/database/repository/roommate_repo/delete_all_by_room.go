package roommate_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) DeleteAllByRoom(tx *sqlx.Tx, roomId int) (err error) {
	log.Debug().Int("roomId", roomId).Msg("Deleting all entities associated with room")

	_, err = tx.Exec(`DELETE FROM sessions WHERE room_id = $1`, roomId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete from sessions")
		return err
	}

	_, err = tx.Exec(`DELETE FROM roommates WHERE room_id = $1`, roomId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete from roommates")
		return err
	}

	_, err = tx.Exec(`DELETE FROM share_codes WHERE room_id = $1`, roomId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete from share_codes")
		return err
	}

	log.Debug().Int("roomId", roomId).Msg("All entities associated with room deleted successfully")
	return nil
}
