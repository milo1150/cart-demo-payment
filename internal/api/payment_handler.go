package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/enums"
	"github.com/milo1150/cart-demo-payment/internal/repositories"
	"github.com/milo1150/cart-demo-payment/internal/types"
	"github.com/milo1150/cart-demo-pkg/pkg"
)

func ConfirmPaymentOrderHandler(c echo.Context, appState *types.AppState) error {
	checkoutId, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.GetSimpleErrorMessage("Invalid param"))
	}

	r := repositories.PaymentOrder{DB: appState.DB}

	// Prevent mutate already updated payment order
	findPaymentOrder, err := r.FindPaymentOrderByCheckoutId(uint(checkoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.GetSimpleErrorMessage(err.Error()))
	}
	if findPaymentOrder.Status != enums.PENDING {
		return c.JSON(http.StatusConflict, pkg.GetSimpleErrorMessage(fmt.Errorf("this payment already update order status").Error()))
	}

	// Update payment order status
	if err := r.ConfirmPaymentOrder(uint(checkoutId)); err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
