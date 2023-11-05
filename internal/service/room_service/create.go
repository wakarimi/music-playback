package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) Create(tx *sqlx.Tx, roomToCreate model.Room, accountID int) (createdRoom model.Room, err error) {
	log.Debug().Str("roomName", roomToCreate.Name).Int("accountID", accountID).Msg("Creating room")

	room := model.Room{
		Name:              roomToCreate.Name,
		OwnerID:           accountID,
		PlaybackOrderType: model.PlaybackInOrder,
	}

	createdRoomID, err := s.RoomRepo.Create(tx, room)
	if err != nil {
		log.Error().Err(err).Str("roomName", room.Name).Int("account", accountID).Msg("Failed to create room")
		return model.Room{}, err
	}

	createdRoom, err = s.Get(tx, createdRoomID, accountID)
	if err != nil {
		log.Error().Err(err).Int("createdRoomID", createdRoomID).Msg("Failed to read created room")
		return model.Room{}, err
	}

	log.Debug().Int("roomID", createdRoom.ID).Str("roomName", createdRoom.Name).Msg("Room created")
	return createdRoom, nil
}
