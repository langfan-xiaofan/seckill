package user

import (
	"context"
	"seckill/internal/model"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
func (r *Repo) CreateUser(ctx context.Context, user model.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error
	return user, err
}
