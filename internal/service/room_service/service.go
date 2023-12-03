package room_service

import (
	"music-playback/internal/database/repository/room_repo"
	"music-playback/internal/service/roommate_service"
)

type Service struct {
	RoomRepo        room_repo.Repo
	RoommateService roommate_service.Service
}

func NewService(roomRepo room_repo.Repo,
	roommateService roommate_service.Service) (s *Service) {

	s = &Service{
		RoomRepo:        roomRepo,
		RoommateService: roommateService,
	}

	return s
}
