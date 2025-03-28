package types

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AppState struct {
	DB                        *gorm.DB
	NATS                      *nats.Conn
	Log                       *zap.Logger
	GrpcShopProductClientConn *grpc.ClientConn
}
