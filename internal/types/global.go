package types

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AppState struct {
	DB                        *gorm.DB
	NATS                      *nats.Conn
	JS                        jetstream.JetStream
	Log                       *zap.Logger
	GrpcShopProductClientConn *grpc.ClientConn
}
