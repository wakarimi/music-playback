package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/model"
)

func (s Service) GetByShareCode(tx *sqlx.Tx, shareCode string) (room model.Room, err error) {
	log.Debug().Msg("Getting room by share code")

	exists, err := s.IsExistsByShareCode(tx, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check room existence")
		return model.Room{}, err
	}
	if !exists {
		err = errors.NotFound{Resource: fmt.Sprintf("room by share code")}
		log.Error().Err(err).Msg("Room by share code not found")
		return model.Room{}, err
	}

	room, err = s.RoomRepo.ReadByShareCode(tx, shareCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get room")
		return model.Room{}, err
	}

	log.Debug().Int("roomId", room.ID).Msg("Room got")
	return room, nil
}
