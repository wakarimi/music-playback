package room_handler

import (
	"music-playback/internal/handler/response"
	"music-playback/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// getMyResponseItem represents the item of response format for a account's rooms request
type getMyResponseItem struct {
	// ID of the room
	ID int `json:"id"`
	// Room owner
	OwnerID int `json:"ownerId"`
	// Current queue item ID
	CurrentQueueItemID *int `json:"currentQueueItemId"`
	// Name of the room
	Name string `json:"name"`
	// Playback order in the created room
	PlaybackOrderType model.PlaybackOrderType `json:"playbackOrderType"`
}

// getMyResponse represents the response format for a account's rooms request
type getMyResponse struct {
	// Account's rooms
	Rooms []getMyResponseItem `json:"rooms"`
}

// Get account's rooms
// @Summary Get account's rooms
// @Tags Rooms
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param X-Account-ID header int true "Account ID"
// @Param request body createRequest true "Room data"
// @Success 201 {object} createResponse
// @Failure 403 {object} response.Error "Invalid X-Account-ID header format"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms [post]
func (h *Handler) GetMy(c *gin.Context) {
	log.Debug().Msg("Getting account's rooms")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	accountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("accountIDHeader", accountIDHeader).Msg("Invalid X-Account-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat",
			}),
			Reason: err.Error(),
		})
		return
	}

	var rooms []model.Room
	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		rooms, err = h.RoomService.GetAllByAccount(tx, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create room")
		c.JSON(http.StatusInternalServerError, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "FailedToGetAccountRooms",
			}),
			Reason: err.Error(),
		})
		return
	}

	responseRooms := make([]getMyResponseItem, len(rooms))
	for i, room := range rooms {
		responseRooms[i] = getMyResponseItem{
			ID:                 room.ID,
			OwnerID:            room.OwnerID,
			CurrentQueueItemID: room.CurrentQueueItemID,
			Name:               room.Name,
			PlaybackOrderType:  room.PlaybackOrderType,
		}
	}

	log.Debug().Msg("Room created")
	c.JSON(http.StatusCreated, getMyResponse{
		Rooms: responseRooms,
	})
}
