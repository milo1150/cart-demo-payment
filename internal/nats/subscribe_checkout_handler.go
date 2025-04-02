package nats

import (
	"encoding/json"
	"fmt"

	"github.com/milo1150/cart-demo-payment/internal/repositories"
	ps "github.com/milo1150/cart-demo-payment/pkg/schemas"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SubscribeCheckoutHandler(log *zap.Logger, db *gorm.DB, msg jetstream.Msg) error {
	payload := ps.CreateCheckoutEventPayload{}
	if err := json.Unmarshal(msg.Data(), &payload); err != nil {
		log.Error("Failed to parse checkout.created payload", zap.Error(err))
		return err
	}

	paymentOrderRepository := repositories.PaymentOrder{DB: db}

	// Check if payment_order already exists
	existsPaymentOrder := paymentOrderRepository.ExistsPaymentOrderByCheckoutId(payload.CheckoutId)
	if existsPaymentOrder {
		msg.Ack() // acknowledge
		return fmt.Errorf("already create payment order with checkout_id = %v", payload.CheckoutId)
	}

	// Create payment_order
	_, err := paymentOrderRepository.CreatePaymentOrder(payload)
	if err != nil {
		return err
	}

	return nil
}
