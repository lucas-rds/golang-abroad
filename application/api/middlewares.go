package api

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (api API) middlewares() {
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: &lumberjack.Logger{
			Filename:   "./output.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   false,
		},
	}))
	api.Use(middleware.Recover())
	api.Use(api.traceRequest)
}

func (api API) traceRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("requestID", time.Now().String())
		return next(c)
	}
}
