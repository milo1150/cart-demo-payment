package nats

import (
	"context"
	"encoding/json"

	ps "github.com/milo1150/cart-demo-payment/pkg/schemas"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func PublishCreatePaymentOrderHandler(js jetstream.JetStream, log *zap.Logger, ctx context.Context, payload ps.CreateCheckoutEventPayload) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Error("Failed to parse payment_order.created payload", zap.Error(err))
		return
	}

	_, err = js.Publish(ctx, "payment_order.created", data)
	if err != nil {
		log.Error("Failed to publish payment_order.created", zap.Error(err))
	}
}
