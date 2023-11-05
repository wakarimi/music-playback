package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/text/language"
	"music-playback/internal/context"
	"music-playback/internal/database/repository/room_repo"
	"music-playback/internal/handlers/room_handler"
	"music-playback/internal/middleware"
	"music-playback/internal/service"
	"music-playback/internal/service/room_service"
)

func SetupRouter(ac *context.AppContext) (r *gin.Engine) {
	log.Debug().Msg("Router setup")
	gin.SetMode(gin.ReleaseMode)

	r = gin.New()
	r.Use(middleware.ZerologMiddleware(log.Logger))
	r.Use(middleware.ProduceLanguageMiddleware())

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.LoadMessageFile("internal/locales/en-US.json")
	bundle.LoadMessageFile("internal/locales/ru-RU.json")

	roomRepo := room_repo.NewRepository()

	txManager := service.NewTransactionManager(*ac.Db)

	roomService := room_service.NewService(roomRepo)

	roomHandler := room_handler.NewHandler(*roomService, txManager, bundle)

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		rooms := api.Group("/rooms")
		{
			rooms.POST("", roomHandler.Create)
			rooms.PATCH("/:roomID/share", roomHandler.GenerateShareCode)
			rooms.PATCH("/:roomID/share-reset", roomHandler.ResetShareCode)
			rooms.GET("/:roomID/share", roomHandler.GetShareCode)
			rooms.DELETE("/:roomID", roomHandler.Delete)
		}
	}

	log.Debug().Msg("Router setup successfully")
	return r
}
