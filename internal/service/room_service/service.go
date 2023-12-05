package room_service

import (
	"music-playback/internal/database/repository/room_repo"
)

type Service struct {
	RoomRepo room_repo.Repo
}

func NewService(roomRepo room_repo.Repo) (s *Service) {

	s = &Service{
		RoomRepo: roomRepo,
	}

	return s
}
