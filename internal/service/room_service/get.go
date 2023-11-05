package room_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/model"
)

func (s Service) Get(tx *sqlx.Tx, roomID int, accountID int) (room model.Room, err error) {
	log.Debug().Int("roomID", roomID).Msg("Getting room")

	exists, err := s.IsExists(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to check room existence")
		return model.Room{}, err
	}
	if !exists {
		err = errors.NotFound{Resource: fmt.Sprintf("room with id=%d", roomID)}
		log.Error().Err(err).Int("roomID", roomID).Msg("Room not found")
		return model.Room{}, err
	}

	room, err = s.RoomRepo.Read(tx, roomID)
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to get room")
		return model.Room{}, err
	}

	if room.OwnerID != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("account %d does not have enough rights to room %d", accountID, roomID)}
		log.Error().Err(err).Int("accountID", accountID).Int("roomID", roomID).Msg("The account does not have enough permissions for the room")
		return model.Room{}, err
	}

	log.Debug().Int("roomID", roomID).Msg("Room got")
	return room, nil
}
