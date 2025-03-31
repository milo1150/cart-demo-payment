package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func ConnectNATS() *nats.Conn {
	url := os.Getenv("NATS_URL")
	token := os.Getenv("NATS_TOKEN")

	nc, err := nats.Connect(url, nats.Token(token))
	if err != nil {
		log.Fatalf("Failed to connect NATS server")
	}

	return nc
}

func ConnectJetStream(nc *nats.Conn) jetstream.JetStream {
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("Failed to connect JetStream")
	}
	return js
}
