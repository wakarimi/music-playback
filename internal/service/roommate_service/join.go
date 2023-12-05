package roommate_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) Join(tx *sqlx.Tx, roomID int, accountID int) (roommateID int, err error) {
	log.Debug().Int("roomId", roomID).Int("accountId", accountID).Msg("Creating roommate")

	roommate := model.Roommate{
		AccountID: accountID,
		RoomID:    roomID,
	}

	roommateID, err = s.RoommateRepo.Create(tx, roommate)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomID).Int("accountId", accountID).Msg("Failed to create roommate")
		return 0, err
	}

	log.Debug().Int("roomId", roomID).Int("accountId", accountID).Msg("Roommate created")
	return roommateID, nil
}
