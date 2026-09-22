package model

import "time"

type ReconcileDiff struct {
	ID             uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SkuID          uint64     `json:"sku_id" gorm:"not null;index:idx_sku_checked,priority:1"`
	RedisAvailable int        `json:"redis_available" gorm:"size:32;not null"`
	DBAvailable    int        `json:"db_available" gorm:"size:32;not null"`
	PendingOrders  int        `json:"pending_orders" gorm:"size:32;not null;comment:待支付订单数，解释差异的中间态"`
	Diff           int        `json:"diff" gorm:"size:32;not null;comment:redis - db - pending"`
	Status         int        `json:"status" gorm:"size:8;not null;default:0;index:idx_status;comment:0待处理 1已自动修复 2需人工"`
	Remark         string     `json:"remark" gorm:"size:255;not null;default:''"`
	CheckedAt      time.Time  `json:"checked_at" gorm:"not null;index:idx_sku_checked,priority:2"`
	FixedAt        *time.Time `json:"fixed_at"`
}
