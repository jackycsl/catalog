package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.FindByUsername(ctx, username)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}

func (uc *UserUseCase) Save(ctx context.Context, u *User) (*User, error) {
	_, err := uc.repo.GetUser(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	p, err := uc.repo.UpdateUser(ctx, u)
	if p.Id == 0 {
		return nil, err
	}
	return p, nil
}
