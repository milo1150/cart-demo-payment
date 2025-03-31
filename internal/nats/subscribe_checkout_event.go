package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/milo1150/cart-demo-payment/internal/schemas"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SubscribeCheckoutEvent(js jetstream.JetStream, log *zap.Logger, db *gorm.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cons, err := js.CreateOrUpdateConsumer(ctx, "CHECKOUT", jetstream.ConsumerConfig{
		Durable:       "CREATED_CHECKOUT_CONS",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "checkout.created",
		AckWait:       5 * time.Second,
	})
	if err != nil {
		log.Error("Failed to create CreateConsumer", zap.Error(err))
	}

	cons.Consume(func(msg jetstream.Msg) {
		err := func() error {
			payload := schemas.CreateCheckoutEventPayload{}
			if err := json.Unmarshal(msg.Data(), &payload); err != nil {
				log.Error("Failed to parse checkout.created payload", zap.Error(err))
				return err
			}

			if err := SubscribeCheckoutHandler(log, db, payload, msg); err != nil {
				log.Error("Failed to create payment order", zap.Error(err))
				return err
			}

			fmt.Println("created", payload)

			return nil
		}()

		if err == nil {
			msg.Ack()
		}
	})
}
