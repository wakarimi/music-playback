package roommate_service

import (
	"music-playback/internal/database/repository/roommate_repo"
)

type Service struct {
	RoommateRepo roommate_repo.Repo
}

func NewService(roommateRepo roommate_repo.Repo) (s *Service) {

	s = &Service{
		RoommateRepo: roommateRepo,
	}

	return s
}
