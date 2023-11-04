package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) Create(tx *sqlx.Tx, roomToCreate model.Room, ownerId int) (createdRoom model.Room, err error) {
	log.Debug().Str("roomName", roomToCreate.Name).Int("ownerId", ownerId).Msg("Creating room")

	preparedRoom := model.Room{
		Name:              roomToCreate.Name,
		OwnerId:           ownerId,
		PlaybackOrderType: "IN_ORDER",
	}

	createdRoomId, err := s.RoomRepo.Create(tx, preparedRoom)
	if err != nil {
		log.Error().Err(err).Str("roomName", preparedRoom.Name).Int("owner", ownerId).Msg("Failed to create room")
		return model.Room{}, err
	}

	createdRoom, err = s.RoomRepo.Read(tx, createdRoomId)
	if err != nil {
		log.Error().Err(err).Int("createdRoomId", createdRoomId).Msg("Failed to read created room")
		return model.Room{}, err
	}

	log.Debug().Int("roomId", createdRoom.Id).Str("roomName", createdRoom.Name).Msg("Room created")
	return createdRoom, nil
}
