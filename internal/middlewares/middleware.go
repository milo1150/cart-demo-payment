package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(logger())
	e.Use(middleware.Recover())
	e.Use(CORS())
}
