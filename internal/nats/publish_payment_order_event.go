package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func PublishCreatePaymentOrderEvent(js jetstream.JetStream) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:      "PAYMENT_ORDER",
		Subjects:  []string{"payment_order.*"},
		Retention: jetstream.WorkQueuePolicy,
	})
}
