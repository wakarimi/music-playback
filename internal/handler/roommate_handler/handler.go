package roommate_handler

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"music-playback/internal/service"
	"music-playback/internal/service/roommate_service"
)

type Handler struct {
	RoommateService    roommate_service.Service
	TransactionManager service.TransactionManager
	Bundle             *i18n.Bundle
}

func NewHandler(roommateService roommate_service.Service,
	transactionManager service.TransactionManager,
	bundle *i18n.Bundle) (h *Handler) {

	h = &Handler{
		RoommateService:    roommateService,
		TransactionManager: transactionManager,
		Bundle:             bundle,
	}

	return h
}
