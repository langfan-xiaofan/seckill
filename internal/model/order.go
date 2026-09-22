package model

import "time"

type Order struct {
	ID             uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;index:idx_user_id,priority:2"`
	OrderNo        string     `json:"order_no" gorm:"type:char(20);not null;uniqueIndex:uk_order_no;comment:对外订单号（雪花ID），不暴露自增ID"`
	UserID         uint64     `json:"user_id" gorm:"not null;uniqueIndex:uk_idem,priority:1;uniqueIndex:uk_user_sku,priority:1;index:idx_user_id,priority:1"`
	ActivityID     uint64     `json:"activity_id" gorm:"not null"`
	SkuID          uint64     `json:"sku_id" gorm:"not null;uniqueIndex:uk_user_sku,priority:2"`
	Quantity       int        `json:"quantity" gorm:"size:32;not null;default:1"`
	AmountCents    int64      `json:"amount_cents" gorm:"not null"`
	Status         int        `json:"status" gorm:"size:8;not null;default:0;index:idx_status_expire,priority:1;comment:0待支付 1已支付 2已取消 3超时关闭 4已退款"`
	IdempotencyKey string     `json:"idempotency_key" gorm:"size:64;not null;uniqueIndex:uk_idem,priority:2"`
	PayNo          string     `json:"pay_no" gorm:"size:64;not null;default:''"`
	ExpireAt       time.Time  `json:"expire_at" gorm:"not null;index:idx_status_expire,priority:2;comment:支付截止时间"`
	PaidAt         *time.Time `json:"paid_at"`
	CreatedAt      time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
