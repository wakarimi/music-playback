package roommate_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) JoinByShareCode(tx *sqlx.Tx, accountID int, shareCode string) (roommateID int, err error) {
	log.Debug().Int("accountId", accountID).Msg("Creating roommate")

	room, err := s.RoomService.GetByShareCode(tx, shareCode)
	if err != nil {
		return 0, err
	}

	roommateID, err = s.Join(tx, room.ID, accountID)
	if err != nil {
		return 0, err
	}

	log.Debug().Int("roomId", room.ID).Int("accountId", accountID).Msg("Roommate created")
	return roommateID, nil
}
