package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"music-playback/internal/context"
	"music-playback/internal/middleware"
)

func SetupRouter(ac *context.AppContext) (r *gin.Engine) {
	log.Debug().Msg("Router setup")
	gin.SetMode(gin.ReleaseMode)

	r = gin.New()
	r.Use(middleware.ZerologMiddleware(log.Logger))

	log.Debug().Msg("Router setup successfully")
	return r
}
