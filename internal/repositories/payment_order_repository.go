package repositories

import (
	"github.com/milo1150/cart-demo-payment/internal/enums"
	"github.com/milo1150/cart-demo-payment/internal/models"
	"github.com/milo1150/cart-demo-payment/internal/schemas"
	"gorm.io/gorm"
)

type PaymentOrder struct {
	DB *gorm.DB
}

func (p *PaymentOrder) CreatePaymentOrder(payload schemas.CreateCheckoutEventPayload) (*models.PaymentOrder, error) {
	po := models.PaymentOrder{
		Total:      0,
		Status:     enums.PENDING,
		UserId:     payload.UserId,
		CheckoutId: payload.CheckoutId,
	}
	if err := p.DB.Create(&po).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *PaymentOrder) ExistsPaymentOrderByCheckoutId(checkoutId uint) bool {
	po := models.PaymentOrder{}
	if err := p.DB.Where("checkout_id = ?", checkoutId).Find(&po).Error; err != nil {
		return false
	}
	if po.ID != 0 {
		return true
	}
	return false
}
