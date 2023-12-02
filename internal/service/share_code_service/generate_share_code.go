package share_code_service

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) GenerateShareCode(tx *sqlx.Tx) (shareCode string, err error) {
	var shareCodeStr string
	alreadyUsed := true
	for alreadyUsed {
		shareCodeStr, err = generateShareCode()
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate share code")
			return "", err
		}

		alreadyUsed, err = s.IsCodeUsed(tx, shareCodeStr)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check share code usage")
			return "", err
		}
	}

	bytes := make([]byte, 16)
	_, err = rand.Read(bytes)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read bytes")
		return "", err
	}
	return hex.EncodeToString(bytes), nil
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
