package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/types"
)

func RegisterAppRoutes(e *echo.Echo, appState *types.AppState) {
	PaymentRoutes(e, appState)
}
