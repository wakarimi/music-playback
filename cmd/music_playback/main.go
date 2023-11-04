package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"music-playback/api"
	"music-playback/internal/config"
	"music-playback/internal/context"
	"music-playback/internal/database"
	"os"

	_ "music-playback/docs"
)

// @title Wakarimi Music Playback API
// @version 0.4.2

// @contact.name Dmitry Kolesnikov (Zalimannard)
// @contact.email zalimannard@mail.ru

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8024
// @BasePath /api
func main() {
	cfg := loadConfiguration()

	initializeLogger(cfg.Logger.Level)

	db := initializeDatabase(cfg.Database.ConnectionString)
	defer closeDatabase(db)
	initializeMigrations(db)

	ctx := context.AppContext{
		Config: cfg,
		Db:     db,
	}

	server := initializeServer(&ctx)
	runServer(server, cfg.HttpServer.Port)
}

func loadConfiguration() *config.Configuration {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to load configuration")
	}
	log.Info().Msg("Configuration loaded")
	return cfg
}

func initializeLogger(level zerolog.Level) (logger *zerolog.Logger) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().Caller().Logger().
		With().Str("service", "music_playback").Logger().
		Level(level)
	log.Debug().Msg("Logger initialized")
	return &log.Logger
}

func initializeDatabase(connectionString string) (db *sqlx.DB) {
	db, err := database.ConnectDb(connectionString)
	if err != nil {
		log.Panic().Msg("Failed to initialize database")
	}
	log.Debug().Msg("Database initialized")
	return db
}

func closeDatabase(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close database connection")
	}
	log.Debug().Msg("Database connection closed")
}

func initializeMigrations(db *sqlx.DB) {
	if err := database.RunMigrations(db, "./internal/database/migrations"); err != nil {
		log.Panic().Err(err).Msg("Failed to apply migrations")
	}
	log.Debug().Msg("Data schema actualized")
}

func initializeServer(ctx *context.AppContext) (r *gin.Engine) {
	r = api.SetupRouter(ctx)
	log.Debug().Msg("Router initialized")
	return r
}

func runServer(server *gin.Engine, port string) {
	if err := server.Run(":" + port); err != nil {
		log.Panic().Err(err).Msg("Failed to start server")
	}
}
