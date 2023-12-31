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

// createRequest represents the request format for creating a room
type createRequest struct {
	// Desired room name
	Name string `json:"name" validate:"required"`
}

// createResponse represents the response format for a room creation request
type createResponse struct {
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

// Create a room
// @Summary Creates a room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param Produce-Language 	header 	string 			false 	"Language preference" default(en-US)
// @Param X-Account-ID 		header 	int 			true 	"Account ID"
// @Param request			body	createRequest	true	"Room data"
// @Success 201 {object} createResponse
// @Failure 400 {object} response.Error "Failed to encode request; Validation failed for request"
// @Failure 403 {object} response.Error "Invalid X-Account-ID header format"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms [post]
func (h *Handler) Create(c *gin.Context) {
	log.Debug().Msg("Creating a room")

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

	var request createRequest
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
		roomToCreate := model.Room{
			Name: request.Name,
		}
		roomID, err := h.RoomService.Create(tx, roomToCreate, accountID)
		if err != nil {
			return err
		}
		room, err = h.RoomService.Get(tx, roomID, accountID)
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create room")
		c.JSON(http.StatusInternalServerError, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID:    "FailedToCreateRoom",
				TemplateData: map[string]interface{}{"RoomName": request.Name}}),
			Reason: err.Error(),
		})
		return
	}

	log.Debug().Msg("Room created")
	c.JSON(http.StatusCreated, createResponse{
		ID:                room.ID,
		OwnerID:           room.OwnerID,
		Name:              room.Name,
		PlaybackOrderType: room.PlaybackOrderType,
	})
}
