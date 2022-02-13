package service

import (
	"context"

	v1 "github.com/jackycsl/catalog/api/bff/interface/v1"
)

func (s *ShopInterface) ListGame(ctx context.Context, req *v1.ListGameReq) (*v1.ListGameReply, error) {
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
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}

func (s *ShopInterface) GetGame(ctx context.Context, req *v1.GetGameReq) (*v1.GetGameReply, error) {
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
