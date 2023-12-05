package roommate_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"music-playback/internal/errors"
	"music-playback/internal/handler/response"
	"net/http"
	"strconv"
)

// joinRequest represents the request format for joining to the room
type joinRequest struct {
	// Desired room name
	ShareCode string `json:"shareCode" validate:"required"`
}

// Join to the room
// @Summary Join to the room
// @Tags Roommates
// @Accept json
// @Produce json
// @Param Produce-Language 	header 	string 			false 	"Language preference" default(en-US)
// @Param X-Account-ID 		header 	int 			true 	"Account ID"
// @Param request			body	joinRequest 	true	"ShareCode"
// @Success 201
// @Failure 400 {object} response.Error "Failed to encode request; Validation failed for request"
// @Failure 403 {object} response.Error "Invalid X-Account-ID header format"
// @Failure 404 {object} response.Error "Share code not found"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /rooms [post]
func (h *Handler) Join(c *gin.Context) {
	log.Debug().Msg("Joining to the room")

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

	var request joinRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "FailedToEncodeRequest"}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Msg("Request encoded successfully")

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

	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		_, err = h.RoommateService.JoinByShareCode(tx, accountID, request.ShareCode)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to join to the room room")
		if _, ok := err.(errors.NotFound); ok {
			c.JSON(http.StatusNotFound, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "NotFound"}),
				Reason: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "FailedToJoinToTheRoom"}),
				Reason: err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Rommate created")
	c.Status(http.StatusCreated)
}
