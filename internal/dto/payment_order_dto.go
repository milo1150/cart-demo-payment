package dto

import (
	"github.com/milo1150/cart-demo-payment/internal/models"
	pb "github.com/milo1150/cart-demo-proto/pkg/payment"
)

func TransformProtoPaymentOrder(model models.PaymentOrder) *pb.GetPaymentOrderResponse {
	data := &pb.GetPaymentOrderResponse{
		Id:     uint64(model.ID),
		Total:  float32(model.Total),
		Status: model.Status.ToString(),
	}
	return data
}

func TransformProtoPaymentOrderList(models []models.PaymentOrder) []*pb.GetPaymentOrderResponse {
	datas := []*pb.GetPaymentOrderResponse{}
	for _, model := range models {
		datas = append(datas, TransformProtoPaymentOrder(model))
	}
	return datas
}
