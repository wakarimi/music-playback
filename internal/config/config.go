package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Database
	HttpServer
	Logger
}

type Database struct {
	ConnectionString string
}

type HttpServer struct {
	Port string
	OtherHttpServers
}

type OtherHttpServers struct {
	MusicMetadataAddress string
}

type Logger struct {
	Level zerolog.Level
}

func LoadConfiguration() (config *Configuration, err error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config = &Configuration{
		Database{
			ConnectionString: viper.GetString("WAKARIMI_MUSIC_PLAYBACK_DB_STRING"),
		},
		HttpServer{
			Port: viper.GetString("HTTP_SERVER_PORT"),
			OtherHttpServers: OtherHttpServers{
				MusicMetadataAddress: viper.GetString("WAKARIMI_MUSIC_METADATA_ADDRESS"),
			},
		},
		Logger{
			Level: loadLoggingLevel(),
		},
	}

	return config, nil
}

func loadLoggingLevel() zerolog.Level {
	levelStr := viper.GetString("LOGGING_LEVEL")
	switch levelStr {
	case "TRACE":
		return zerolog.TraceLevel
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
