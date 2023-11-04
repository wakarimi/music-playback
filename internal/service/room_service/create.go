package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) Create(tx *sqlx.Tx, roomToCreate model.Room, ownerID int) (createdRoom model.Room, err error) {
	log.Debug().Str("roomName", roomToCreate.Name).Int("ownerID", ownerID).Msg("Creating room")

	preparedRoom := model.Room{
		Name:              roomToCreate.Name,
		OwnerID:           ownerID,
		PlaybackOrderType: "IN_ORDER",
	}

	createdRoomID, err := s.RoomRepo.Create(tx, preparedRoom)
	if err != nil {
		log.Error().Err(err).Str("roomName", preparedRoom.Name).Int("owner", ownerID).Msg("Failed to create room")
		return model.Room{}, err
	}

	createdRoom, err = s.RoomRepo.Read(tx, createdRoomID)
	if err != nil {
		log.Error().Err(err).Int("createdRoomID", createdRoomID).Msg("Failed to read created room")
		return model.Room{}, err
	}

	log.Debug().Int("roomID", createdRoom.ID).Str("roomName", createdRoom.Name).Msg("Room created")
	return createdRoom, nil
}
