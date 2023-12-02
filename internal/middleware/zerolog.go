package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"time"
)

func ZerologMiddleware(logger zerolog.Logger) (f gin.HandlerFunc) {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("duration", duration).
			Msg("Request handled")
	}
}
