package repositories

import (
	"fmt"

	"github.com/milo1150/cart-demo-payment/internal/enums"
	"github.com/milo1150/cart-demo-payment/internal/models"
	ps "github.com/milo1150/cart-demo-payment/pkg/schemas"
	"gorm.io/gorm"
)

type PaymentOrder struct {
	DB *gorm.DB
}

func (p *PaymentOrder) CreatePaymentOrder(payload ps.CreateCheckoutEventPayload) (*models.PaymentOrder, error) {
	po := models.PaymentOrder{
		Total:      0,
		Status:     enums.PENDING,
		UserId:     payload.UserId,
		CheckoutId: payload.CheckoutId,
	}
	if err := p.DB.Create(&po).Error; err != nil {
		return nil, err
	}
	return &po, nil
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

func (p *PaymentOrder) FindPaymentOrderByCheckoutId(checkoutId uint) (*models.PaymentOrder, error) {
	po := models.PaymentOrder{}
	query := p.DB.Where("checkout_id = ?", checkoutId).First(&po)
	if query.Error != nil {
		return nil, query.Error
	}
	return &po, nil
}

func (p *PaymentOrder) ConfirmPaymentOrder(checkoutId uint) error {
	result := p.DB.Model(&models.PaymentOrder{}).
		Where("checkout_id = ?", checkoutId).
		Update("status", enums.DONE)
	if result.Error != nil {
		return fmt.Errorf("failed to update payment order status: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no payment order found with checkout_id %d", checkoutId)
	}
	return nil
}
