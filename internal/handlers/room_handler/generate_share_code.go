package room_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/handlers/response"
	"net/http"
	"strconv"
)

// generateShareCodeResponse is the data format returned when generating the share code
type generateShareCodeResponse struct {
	// Code to connect to the room
	ShareCode *string `json:"shareCode"`
}

// GenerateShareCode creates or recreates a code to connect to a room
// @Summary Creates or recreates a code to connect to a room
// @Tags Rooms
// @Accept 	json
// @Produce json
// @Param Produce-Language 	header 	string 	false 	"Language preference" default(en-US)
// @Param X-Account-Id 		header 	int 	true 	"Account ID"
// @Param roomId 			path 	int 	true 	"Room ID"
// @Success 200 {object} generateShareCodeResponse
// @Failure 403 {object} response.Error "Trying to generate a code for someone else's room; Invalid X-Account-Id header format"
// @Failure 404 {object} response.Error "The room does not exist"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms/{roomId}/share-gen [patch]
func (h *Handler) GenerateShareCode(c *gin.Context) {
	log.Debug().Msg("Generating share code")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-Id")
	accountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("accountIDHeader", accountIDHeader).Msg("Invalid X-Account-Id format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID:    "InvalidHeaderFormat",
				TemplateData: map[string]interface{}{"Header": "X-Account-Id"}}),
			Reason: err.Error(),
		})
		return
	}

	roomIdStr := c.Param("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		log.Error().Err(err).Str("roomIdStr", roomIdStr).Msg("Invalid roomId format")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidUrlParameterFormat"}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Int("roomId", roomId).Msg("Url parameter read successfully")

	var shareCode *string
	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.RoomService.GenerateShareCode(tx, roomId, accountID)
		if err != nil {
			return err
		}
		shareCode, err = h.RoomService.GetShareCode(tx, roomId, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Int("roomId", roomId).Msg("Failed to generate share code")
		if _, ok := err.(errors.Forbidden); ok {
			c.JSON(http.StatusForbidden, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "NotEnoughRightsToTheRoom"}),
				Reason: err.Error(),
			})
			return
		} else if _, ok := err.(errors.NotFound); ok {
			c.JSON(http.StatusNotFound, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "RoomNotFound"}),
				Reason: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "InternalServerError"}),
				Reason: err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Share code generated")
	c.JSON(http.StatusOK, generateShareCodeResponse{
		ShareCode: shareCode,
	})
}
