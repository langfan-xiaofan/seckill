package model

import "time"

type Stock struct {
	SkuID          uint64    `json:"sku_id" gorm:"primary_key;autoIncrement:false"`
	TotalStock     int       `json:"total_stock" gorm:"size:32;not null"`
	AvailableStock int       `json:"available_stock" gorm:"size:32;not null;check:chk_available,available_stock >= 0;comment:可售库存，只能通过条件更新扣减"`
	SoldStock      int       `json:"sold_stock" gorm:"size:32;not null;default:0"`
	Version        int       `json:"version" gorm:"size:32;not null;default:0;comment:乐观锁版本（备用方案）"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
