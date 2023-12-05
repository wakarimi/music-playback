package roommate_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Leave(tx *sqlx.Tx, roomID int, accountID int) (roommateID int, err error) {
	log.Debug().Int("roomId", roomID).Int("accountId", accountID).Msg("Creating roommate")

	room, err := s.RoomService.Get(tx, roomID, accountID)
	if err != nil {
		return 0, err
	}

	if room.OwnerID == accountID {
		err = s.RoommateRepo.DeleteAllByRoom(tx, roomID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete roommates")
			return 0, err
		}
		err = s.RoomService.Delete(tx, roomID, accountID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete room")
			return 0, err
		}
	} else {
		err = s.RoommateRepo.Delete(tx, roomID, accountID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete roommate")
			return 0, err
		}
	}

	log.Debug().Int("roomId", roomID).Int("accountId", accountID).Msg("Roommate deleted")
	return roommateID, nil
}
