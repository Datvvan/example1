package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var appLog *log.Logger

func InitLogger() {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	appLog = logger
}

func GetLogger() *log.Logger {
	return appLog
}

func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := GetLogger()
		start := time.Now()
		res := next(c)

		logger.WithFields(log.Fields{
			"method":     c.Request().Method,
			"path":       c.Path(),
			"status":     c.Response().Status,
			"latency_ns": time.Since(start).Nanoseconds(),
		}).Info("request details")
		return res
	}
}
