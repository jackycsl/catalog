package service

import (
	"context"
	"errors"

	v1 "github.com/jackycsl/catalog/api/user/service/v1"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"github.com/jackycsl/catalog/user-service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, v1.ErrorUnknownError(err.Error())
	}
	return &v1.CreateUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrUserNotFound):
			return nil, v1.ErrorUserNotFound("User %d not found", req.Id)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrUserNotFound):
			return nil, v1.ErrorUserNotFound("User %d not found", req.Username)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	if !rv {
		return nil, v1.ErrorLoginFailed("Incorrect Password")
	}
	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	rv, err := s.uc.GetUserByUsername(ctx, req.Username)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrUserNotFound):
			return nil, v1.ErrorUserNotFound("User %s not found", req.Username)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.GetUserByUsernameReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}

func (s *UserService) Save(ctx context.Context, req *v1.SaveUserReq) (*v1.SaveUserReply, error) {
	rv, err := s.uc.Save(ctx, &biz.User{Id: req.Id, Username: req.Username, Password: req.Password})
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrUserNotFound):
			return nil, v1.ErrorUserNotFound("User %d not found", req.Username)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.SaveUserReply{
		Id: rv.Id,
	}, nil
}
