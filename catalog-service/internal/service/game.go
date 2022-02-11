package service

import (
	"context"
	"errors"

	v1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

func (s *CatalogService) CreateGame(ctx context.Context, req *v1.CreateGameReq) (*v1.CreateGameReply, error) {
	b := &biz.Game{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
	}
	x, err := s.bc.Create(ctx, b)
	if err != nil {
		return nil, v1.ErrorUnknownError(err.Error())
	}
	return &v1.CreateGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}, err
}

func (s *CatalogService) GetGame(ctx context.Context, req *v1.GetGameReq) (*v1.GetGameReply, error) {

	x, err := s.bc.Get(ctx, req.Id)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrRecordNotFound):
			return nil, v1.ErrorGameNotFound("Game %s not found", req.Id)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}

	return &v1.GetGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}, err
}

func (s *CatalogService) UpdateGame(ctx context.Context, req *v1.UpdateGameReq) (*v1.UpdateGameReply, error) {
	b := &biz.Game{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
	}
	if req.Id == 0 {
		return nil, v1.ErrorUnknownError("Invalid Input")
	}
	x, err := s.bc.Update(ctx, b)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrRecordNotFound):
			return nil, v1.ErrorGameNotFound("Game %s not found", req.Id)
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.UpdateGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}, err
}

func (s *CatalogService) ListGame(ctx context.Context, req *v1.ListGameReq) (*v1.ListGameReply, error) {

	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrRecordNotFound):
			return nil, v1.ErrorGameNotFound("No Games not found")
		default:
			return nil, v1.ErrorUnknownError(err.Error())
		}
	}
	rs := make([]*v1.ListGameReply_Game, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListGameReply_Game{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
		})
	}

	return &v1.ListGameReply{
		Results: rs,
	}, err
}

func (s *CatalogService) HealthCheck(ctx context.Context, req *v1.HealthReq) (*v1.HealthReply, error) {
	return &v1.HealthReply{Status: 1}, nil
}
