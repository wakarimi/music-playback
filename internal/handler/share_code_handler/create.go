package share_code_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/handler/response"
	"music-playback/internal/model"
	"net/http"
	"strconv"
)

// createResponse is the data format returned when generating the share code
type createResponse struct {
	// Room ID for code
	RoomID int `json:"roomID"`
	// Code to connect to the room
	Code string `json:"code"`
}

// Create creates and sets or replaces share code
// @Summary Creates and sets or replaces share code
// @Tags ShareCode
// @Accept 	json
// @Produce json
// @Param Produce-Language	header	string	false	"Language preference" default(en-US)
// @Param X-Account-ID		header	int		true	"Account ID"
// @Param roomID			path	int		true	"Room ID"
// @Failure 400 {object} response.Error "Invalid roomID parameter"
// @Failure 403 {object} response.Error "Trying to generate a code for someone else's room; Invalid X-Account-ID header format"
// @Failure 404 {object} response.Error "The room does not exist"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms/{roomID}/share-code [post]
func (h *Handler) Create(c *gin.Context) {
	log.Debug().Msg("Creating share code")

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

	var shareCode model.ShareCode
	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.ShareCodeService.CreateOrUpdate(tx, roomID, accountID)
		if err != nil {
			return err
		}
		shareCode, err = h.ShareCodeService.Get(tx, roomID, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Int("roomID", roomID).Msg("Failed to create share code")
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
					MessageID: "NotFound"}),
				Reason: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "FailedToGenerateShareCode"}),
				Reason: err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Share code generated")
	c.JSON(http.StatusOK, createResponse{
		RoomID: shareCode.RoomID,
		Code:   shareCode.Code,
	})
}
