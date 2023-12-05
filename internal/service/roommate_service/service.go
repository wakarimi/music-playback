package roommate_service

import (
	"music-playback/internal/database/repository/roommate_repo"
	"music-playback/internal/service/room_service"
)

type Service struct {
	RoommateRepo roommate_repo.Repo
	RoomService  room_service.Service
}

func NewService(roommateRepo roommate_repo.Repo,
	roomService room_service.Service) (s *Service) {

	s = &Service{
		RoommateRepo: roommateRepo,
		RoomService:  roomService,
	}

	return s
}
