package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModelSchema struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Uuid      uuid.UUID      `json:"uuid"`
}
