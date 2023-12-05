package roommate_handler

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"music-playback/internal/service"
	"music-playback/internal/service/room_service"
	"music-playback/internal/service/roommate_service"
)

type Handler struct {
	RoomService        room_service.Service
	RoommateService    roommate_service.Service
	TransactionManager service.TransactionManager
	Bundle             *i18n.Bundle
}

func NewHandler(roomService room_service.Service,
	roommateService roommate_service.Service,
	transactionManager service.TransactionManager,
	bundle *i18n.Bundle) (h *Handler) {

	h = &Handler{
		RoomService:        roomService,
		RoommateService:    roommateService,
		TransactionManager: transactionManager,
		Bundle:             bundle,
	}

	return h
}
