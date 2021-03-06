package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/jackycsl/catalog/api/user/service/v1"
	"github.com/jackycsl/catalog/user-service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase
	cc *biz.CardUseCase
	ac *biz.AddressUseCase

	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, cc *biz.CardUseCase, ac *biz.AddressUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		cc:  cc,
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/server-service"))}
}
