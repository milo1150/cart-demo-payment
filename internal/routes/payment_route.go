package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/types"
)

func PaymentRoutes(e *echo.Echo, appState *types.AppState) {
	paymentGroup := e.Group("/payment")

	paymentGroup.POST("/order/confirm", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "TODO: Confirm payment")
	})
}
