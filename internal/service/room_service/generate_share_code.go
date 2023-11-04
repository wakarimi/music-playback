package room_service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
)

func (s Service) GenerateShareCode(tx *sqlx.Tx, roomId int, accountID int) (err error) {
	log.Debug().Int("roomId", roomId).Msg("Generating share code")

	roomExists, err := s.RoomRepo.IsExists(tx, roomId)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to check room existence")
		return err
	}
	if !roomExists {
		err = errors.NotFound{Resource: fmt.Sprintf("room with id=%d", roomId)}
		log.Error().Err(err).Int("roomId", roomId).Msg("Room not found")
		return err
	}

	room, err := s.RoomRepo.Read(tx, roomId)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to read room")
		return err
	}
	if room.OwnerId != accountID {
		err = errors.Forbidden{Message: fmt.Sprintf("room %d is not owned by account %d", roomId, accountID)}
		log.Error().Err(err).Int("roomId", roomId).Int("accountID", accountID).Msg("Room is not owned by account")
		return err
	}

	var shareCode string
	alreadyUsed := true
	for alreadyUsed {
		shareCode, err = generateShareCode()
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate share code")
			return err
		}

		alreadyUsed, err = s.RoomRepo.IsShareCodeUsed(tx, shareCode)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check share code usage")
			return err
		}
	}

	err = s.RoomRepo.UpdateShareCode(tx, roomId, shareCode)
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to update share code")
		return err
	}

	log.Debug().Int("roomId", roomId).Msg("Share code generated and set")
	return nil
}

func generateShareCode() (shareCode string, err error) {
	bytes := make([]byte, 16)
	_, err = rand.Read(bytes)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read bytes")
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
