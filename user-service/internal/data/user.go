package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/jackycsl/catalog/pkg/util/helper"
	"github.com/jackycsl/catalog/user-service/internal/biz"
	"github.com/jackycsl/catalog/user-service/internal/data/ent"
	"github.com/jackycsl/catalog/user-service/internal/data/ent/user"
	"github.com/jackycsl/catalog/user-service/internal/pkg/util"
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
	target, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, helper.ErrUserNotFound
		default:
			return nil, err
		}
	}

	return &biz.User{
		Id:       target.ID,
		Username: target.Username,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       po.ID,
		Username: po.Username,
	}, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, helper.ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &biz.User{
		Id:       po.ID,
		Username: po.Username,
	}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return false, helper.ErrUserNotFound
		default:
			return false, err
		}
	}
	return util.CheckPasswordHash(u.Password, po.PasswordHash), nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		UpdateOneID(u.Id).
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, helper.ErrUserNotFound
		default:
			return nil, err
		}
	}

	return &biz.User{
		Id:       po.ID,
		Username: po.Username,
	}, nil
}
