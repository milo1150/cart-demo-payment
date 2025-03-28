package models

import (
	"cart-service/internal/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentOrder struct {
	ID        uint                `json:"id" gorm:"primarykey"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	Uuid      uuid.UUID           `json:"uuid" gorm:"not null;type:uuid;unique;index"`
	Total     float64             `json:"total"`
	Status    enums.PaymentStatus `json:"status" gorm:"not null"`

	// External relation
	UserId     uint `json:"user_id"`
	CheckoutId uint `json:"checkout_id"`
}

func (p *PaymentOrder) BeforeCreate(tx *gorm.DB) error {
	if p.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		p.Uuid = uuidV7
	}
	return nil
}
