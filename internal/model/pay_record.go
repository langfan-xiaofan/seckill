package model

import (
	"encoding/json"
	"time"
)

type PayRecord struct {
	ID          uint64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OrderNo     string          `json:"order_no" gorm:"type:char(20);not null;index:idx_order_no"`
	Channel     string          `json:"channel" gorm:"size:16;not null;default:mock;uniqueIndex:uk_channel,priority:1"`
	ChannelNo   string          `json:"channel_no" gorm:"size:64;not null;uniqueIndex:uk_channel,priority:2;comment:第三方支付单号，回调幂等的依据"`
	AmountCents int64           `json:"amount_cents" gorm:"not null"`
	Status      int             `json:"status" gorm:"size:8;not null;default:1"`
	RawPayload  json.RawMessage `json:"raw_payload" gorm:"type:json"`
	CreatedAt   time.Time       `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
