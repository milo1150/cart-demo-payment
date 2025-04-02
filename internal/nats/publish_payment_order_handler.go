package nats

import (
	"context"
	"encoding/json"
	"fmt"

	ps "github.com/milo1150/cart-demo-payment/pkg/schemas"
	"github.com/nats-io/nats.go/jetstream"
)

func PublishCreatePaymentOrderHandler(js jetstream.JetStream, ctx context.Context, payload ps.PublishCreatedPaymentOrderPayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Failed to parse payment_order.created payload: %w", err)
	}

	_, err = js.Publish(ctx, "payment_order.created", data)
	if err != nil {
		return fmt.Errorf("Failed to publish payment_order.created: %w", err)
	}

	return nil
}
