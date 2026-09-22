package model

import "time"

type Activity struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title     string    `json:"title" gorm:"size:128;not null"`
	StartAt   time.Time `json:"start_at" gorm:"not null;index:idx_status_start,priority:2"`
	EndAt     time.Time `json:"end_at" gorm:"not null"`
	Status    int       `json:"status" gorm:"size:8;not null;default:0;index:idx_status_start,priority:1;comment:0未开始 1进行中 2已结束 3已下线"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
