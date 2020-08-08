package api

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (api API) enableMiddlewares() {
	api.Echo.Use(middleware.Logger())
	api.Echo.Use(middleware.Recover())
	api.Echo.Use(api.traceRequest)
}

func (api API) traceRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("requestID", time.Now().String())
		return next(c)
	}
}
