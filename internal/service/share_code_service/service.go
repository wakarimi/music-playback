package share_code_service

import (
	"music-playback/internal/database/repository/share_code_repo"
	"music-playback/internal/service/room_service"
)

type Service struct {
	ShareCodeRepo share_code_repo.Repo
	RoomService   room_service.Service
}

func NewService(shareCodeRepo share_code_repo.Repo,
	roomService room_service.Service) (s *Service) {

	s = &Service{
		ShareCodeRepo: shareCodeRepo,
		RoomService:   roomService,
	}

	return s
}
