package grpc

import (
	"context"
	"log"
	"net"

	"github.com/milo1150/cart-demo-payment/internal/repositories"
	"github.com/milo1150/cart-demo-payment/internal/types"
	"google.golang.org/grpc"

	pb "github.com/milo1150/cart-demo-proto/pkg/payment"
)

type PaymentGRPCServer struct {
	pb.UnimplementedPaymentServiceServer
	AppState *types.AppState
}

func StartPaymentGRPCServer(appState *types.AppState) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &PaymentGRPCServer{AppState: appState})

	log.Println("gRPC server is running on port 50051")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve Payment gRPC server: %v", err)
	}

}

func (p *PaymentGRPCServer) GetPayment(_ context.Context, payload *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	pr := repositories.PaymentOrder{DB: p.AppState.DB}
	paymentOrder, err := pr.FindPaymentOrderByCheckoutId(uint(payload.PaymentOrderId))
	if err != nil {
		return nil, err
	}

	res := pb.GetPaymentResponse{
		Id:     uint64(paymentOrder.ID),
		Total:  float32(paymentOrder.Total),
		Status: paymentOrder.Status.ToString(),
	}

	return &res, nil
}
