package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	origin := os.Getenv("ORIGIN_URL")

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{origin},
		AllowMethods: []string{
			echo.PATCH,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	})
}
