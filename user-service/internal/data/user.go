package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/user-service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	return false, nil
}
