package service

import (
	"context"
	"seckill/internal/dto"
	"seckill/internal/model"
	"seckill/internal/pkg/jwt"
	"seckill/internal/pkg/password"
	"seckill/internal/repo/user"

	"gorm.io/gorm"
)

type UserService struct {
	repo *user.Repo
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: user.NewRepo(db),
	}
}

func (svc *UserService) Register(ctx context.Context, username string, pwd string) error {
	var u model.User
	u.Username = username
	passwordHash, err := password.HashPassword(pwd)
	if err != nil {
		return err
	}
	u.PasswordHash = passwordHash
	err = svc.repo.CreateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) Login(ctx context.Context, username string, pwd string) (dto.LoginResp, error) {
	var u model.User
	u, err := svc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return dto.LoginResp{}, err
	}
	if !password.CheckPasswordHash(pwd, u.PasswordHash) {
		return dto.LoginResp{}, nil
	}
	token, err := jwt.GenerateToken(u.Username, u.ID)
	if err != nil {
		return dto.LoginResp{}, err
	}
	return dto.LoginResp{
		Token: token,
	}, nil
}
