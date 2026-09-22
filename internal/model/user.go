package model

import "time"

type User struct {
	ID           uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username     string    `json:"username" gorm:"size:64;not null;uniqueIndex:uk_username"`
	PasswordHash string    `json:"-" gorm:"size:100;not null;comment:bcrypt"`
	Status       int       `json:"status" gorm:"size:8;not null;default:1;comment:1正常 0禁用"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP(3)"`
}
