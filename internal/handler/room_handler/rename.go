package room_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"music-playback/internal/handler/response"
	"music-playback/internal/model"
	"net/http"
	"strconv"
)

// renameRequest represents the request format for renaming a room
type renameRequest struct {
	// Desired room name
	Name string `json:"name" validate:"required"`
}

// renameResponse represents the response format for a room creation request
type renameResponse struct {
	// ID of the created room
	ID int `json:"id"`
	// Room owner
	OwnerID int `json:"ownerId"`
	// Current queue item ID
	CurrentQueueItemID *int `json:"currentQueueItemId"`
	// Name of the created room
	Name string `json:"name"`
	// Playback order in the created room
	PlaybackOrderType model.PlaybackOrderType `json:"playbackOrderType"`
}

// Rename a room
// @Summary Renames a room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param Produce-Language 	header 	string 			false 	"Language preference" default(en-US)
// @Param X-Account-ID 		header 	int 			true 	"Account ID"
// @Param roomID			path	int				true	"Room ID"
// @Param request			body	renameRequest	true	"Room data"
// @Success 201 {object} renameResponse
// @Failure 400 {object} response.Error "Trying to rename someone else's room; Failed to encode request; Validation failed for request"
// @Failure 403 {object} response.Error "Invalid X-Account-ID header format"
// @Failure 404 {object} response.Error "The room does not exist"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms/{roomID}/rename [patch]
func (h *Handler) Rename(c *gin.Context) {
	log.Debug().Msg("Renaming a room")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	accountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("accountIDHeader", accountIDHeader).Msg("Invalid X-Account-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat"}),
			Reason: err.Error(),
		})
		return
	}

	roomIDStr := c.Param("roomID")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Error().Err(err).Str("roomIDStr", roomIDStr).Msg("Invalid roomID format")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidUrlParameterFormat"}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Int("roomID", roomID).Msg("Url parameter read successfully")

	var request renameRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "FailedToEncodeRequest"}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Str("roomName", request.Name).Msg("Request encoded successfully")

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation failed for request")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ValidationFailedForRequest"}),
			Reason: err.Error(),
		})
		return
	}

	var room model.Room
	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.RoomService.Rename(tx, roomID, request.Name, accountID)
		if err != nil {
			return err
		}
		room, err = h.RoomService.Get(tx, roomID, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create room")
		c.JSON(http.StatusInternalServerError, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID:    "FailedToRenameRoom",
				TemplateData: map[string]interface{}{"RoomName": request.Name}}),
			Reason: err.Error(),
		})
		return
	}

	log.Debug().Msg("Room created")
	c.JSON(http.StatusOK, createResponse{
		ID:                room.ID,
		OwnerID:           room.OwnerID,
		Name:              room.Name,
		PlaybackOrderType: room.PlaybackOrderType,
	})
}
