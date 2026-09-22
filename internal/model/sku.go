package model

import "time"

type Sku struct {
	ID           uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ActivityID   uint64    `json:"activity_id" gorm:"not null;index:idx_activity"`
	Name         string    `json:"name" gorm:"size:128;not null"`
	PriceCents   int64     `json:"price_cents" gorm:"not null;comment:单价（分）"`
	LimitPerUser int       `json:"limit_per_user" gorm:"size:32;not null;default:1;comment:每人限购件数"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
