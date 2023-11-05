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
	"music-playback/internal/database/repository/share_code_repo"
	"music-playback/internal/handlers/room_handler"
	"music-playback/internal/handlers/share_code_handler"
	"music-playback/internal/middleware"
	"music-playback/internal/service"
	"music-playback/internal/service/room_service"
	"music-playback/internal/service/share_code_service"
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
	shareCodeRepo := share_code_repo.NewRepository()

	txManager := service.NewTransactionManager(*ac.Db)

	roomService := room_service.NewService(roomRepo)
	shareCodeService := share_code_service.NewService(shareCodeRepo, *roomService)

	roomHandler := room_handler.NewHandler(*roomService, txManager, bundle)
	shareCodeHandler := share_code_handler.NewHandler(*shareCodeService, txManager, bundle)

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		rooms := api.Group("/rooms")
		{
			rooms.GET("")                      // Получение комнат, в которых состоит пользователь
			rooms.POST("", roomHandler.Create) // Создание комнаты
			room := rooms.Group("/:roomID")
			{
				room.POST("/join")                  // Присоединение к комнате
				room.PATCH("/rename")               // Переименование комнаты
				room.DELETE("", roomHandler.Delete) // Удаление комнаты

				playback := room.Group("/playback")
				{
					playback.GET("")        // Информация о текущем воспроизведении
					playback.POST("/play")  // Отпаузить
					playback.POST("/pause") // Пауза
					playback.POST("/next")  // Переход к следующему треку
					playback.POST("/jump")  // Перейти к конкретному треку в очереди
					playback.POST("/order") // Изменение порядка воспроизведения

				}
				queue := room.Group("/queue")
				{
					queue.GET("")       // Получение текущей очереди комнаты
					queue.POST("")      // Добавление трека в очередь
					queue.POST("/move") // Переместить трек в очереди
					queue.DELETE("")    // Очищение очереди
					queueItem := queue.Group("/:queueItemID")
					{
						queueItem.DELETE("") // Удаление трека из очереди
					}
				}
				roommates := room.Group("/roommates")
				{
					room.GET("") // Получение всех членов комнаты
					roommate := roommates.Group("/:roommateID")
					{
						roommate.DELETE("") // Удаление члена группы/Выход из группы
					}
				}
				shareCode := room.Group("/share-code")
				{
					shareCode.GET("", shareCodeHandler.Get)       // Запрос кода комнаты
					shareCode.POST("", shareCodeHandler.Create)   // Генерация/Перегенерация кода
					shareCode.DELETE("", shareCodeHandler.Delete) // Запрет входа в комнату
				}
			}
		}
	}

	log.Debug().Msg("Router setup successfully")
	return r
}
