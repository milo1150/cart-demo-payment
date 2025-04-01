package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/api"
	"github.com/milo1150/cart-demo-payment/internal/types"
)

func PaymentRoutes(e *echo.Echo, appState *types.AppState) {
	paymentGroup := e.Group("/order")

	paymentGroup.PATCH("/confirm/:checkout_id", func(c echo.Context) error {
		return api.ConfirmPaymentOrderHandler(c, appState)
	})
}
