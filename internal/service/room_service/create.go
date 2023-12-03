package room_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/model"
)

func (s Service) Create(tx *sqlx.Tx, roomToCreate model.Room, accountID int) (roomID int, err error) {
	log.Debug().Str("roomName", roomToCreate.Name).Int("accountID", accountID).Msg("Creating room")

	room := model.Room{
		Name:              roomToCreate.Name,
		OwnerID:           accountID,
		PlaybackOrderType: model.PlaybackInOrder,
	}

	createdRoomID, err := s.RoomRepo.Create(tx, room)
	if err != nil {
		log.Error().Err(err).Str("roomName", room.Name).Int("account", accountID).Msg("Failed to create room")
		return 0, err
	}

	_, err = s.RoommateService.Create(tx, createdRoomID, accountID)
	if err != nil {
		log.Error().Err(err).Int("roomId", createdRoomID).Int("accountId", accountID).Msg("Failed to create owner's roommate")
		return 0, err
	}

	log.Debug().Int("roomID", createdRoomID).Str("roomToCreate.Name", roomToCreate.Name).Msg("Room created")
	return createdRoomID, nil
}
