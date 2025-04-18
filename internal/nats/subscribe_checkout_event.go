package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SubscribeCheckoutEvent(js jetstream.JetStream, log *zap.Logger, db *gorm.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tick := time.NewTicker(2 * time.Second)
	ok := make(chan bool)
	var consumer jetstream.Consumer

	go func() {
		defer tick.Stop()
		for range tick.C {
			cons, err := js.CreateOrUpdateConsumer(ctx, "CHECKOUT", jetstream.ConsumerConfig{
				Durable:       "CREATED_CHECKOUT_CONS",
				AckPolicy:     jetstream.AckExplicitPolicy,
				FilterSubject: "checkout.created",
				AckWait:       5 * time.Second,
			})

			if err != nil {
				log.Error("Failed to create checkout.created consumer", zap.Error(err))
				continue
			}

			// Create consumer
			log.Info("create CREATED_CHECKOUT_CONS consumer OK")
			consumer = cons
			ok <- true
			return
		}
	}()

	<-ok

	consumer.Consume(func(msg jetstream.Msg) {
		err := SubscribeCheckoutHandler(log, db, msg, js)
		if err != nil {
			log.Error("Failed to create payment order", zap.Error(err))
		}
		if err == nil {
			msg.Ack()
		}
	})
}
