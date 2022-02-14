package service

import (
	"context"

	v1 "github.com/jackycsl/catalog/api/bff/admin/v1"
	"github.com/jackycsl/catalog/bff/admin/internal/biz"
)

func (s *ShopAdmin) ListGame(ctx context.Context, req *v1.ListGameReq) (*v1.ListGameReply, error) {
	rv, err := s.cc.ListGame(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListGameReply{
		Results: make([]*v1.ListGameReply_Game, 0),
	}
	for _, x := range rv {
		item := &v1.ListGameReply_Game{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}

func (s *ShopAdmin) GetGame(ctx context.Context, req *v1.GetGameReq) (*v1.GetGameReply, error) {
	x, err := s.cc.GetGame(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}
	return reply, nil
}

func (s *ShopAdmin) CreateGame(ctx context.Context, req *v1.CreateGameReq) (*v1.CreateGameReply, error) {
	x, err := s.cc.CreateGame(ctx, &biz.Game{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.CreateGameReply{
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}
	return reply, nil
}

func (s *ShopAdmin) UpdateGame(ctx context.Context, req *v1.UpdateGameReq) (*v1.UpdateGameReply, error) {
	x, err := s.cc.UpdateGame(ctx, &biz.Game{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.UpdateGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
	}
	return reply, nil
}
