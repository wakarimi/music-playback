package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) GetAllByAccount(tx *sqlx.Tx, accountID int) (rooms []model.Room, err error) {
	log.Debug().Int("accountId", accountID).Msg("Getting account's rooms")

	rooms, err = s.RoomRepo.ReadAllByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to get room")
		return make([]model.Room, 0), err
	}

	log.Debug().Int("accountId", accountID).Int("count", len(rooms)).Msg("Rooms got")
	return rooms, nil
}
