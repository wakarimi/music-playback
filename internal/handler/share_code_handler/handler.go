package share_code_handler

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"music-playback/internal/service"
	"music-playback/internal/service/share_code_service"
)

type Handler struct {
	ShareCodeService   share_code_service.Service
	TransactionManager service.TransactionManager
	Bundle             *i18n.Bundle
}

func NewHandler(shareCodeService share_code_service.Service,
	transactionManager service.TransactionManager,
	bundle *i18n.Bundle) (h *Handler) {

	h = &Handler{
		ShareCodeService:   shareCodeService,
		TransactionManager: transactionManager,
		Bundle:             bundle,
	}

	return h
}
