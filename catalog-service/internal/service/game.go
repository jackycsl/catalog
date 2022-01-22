package service

import (
	"context"

	v1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"go.opentelemetry.io/otel"
)

func (s *CatalogService) CreateGame(ctx context.Context, req *v1.CreateGameReq) (*v1.CreateGameReply, error) {
	// b := &biz.Game{
	// 	Name:        req.Name,
	// 	Description: req.Description,
	// 	Count:       req.Count,
	// 	Images:      make([]biz.Image, 0),
	// }
	// for _, x := range req.Image {
	// 	b.Images = append(b.Images, biz.Image{URL: x.Url})
	// }
	// x, err := s.bc.Create(ctx, b)
	// img := make([]*v1.CreateGameReply_Image, 0)
	// for _, i := range x.Images {
	// 	img = append(img, &v1.CreateGameReply_Image{Url: i.URL})
	// }
	// return &v1.CreateGameReply{
	// 	Id:          x.Id,
	// 	Name:        x.Name,
	// 	Description: x.Description,
	// 	Count:       x.Count,
	// 	Image:       img,
	// }, err
	return nil, nil
}

func (s *CatalogService) GetGame(ctx context.Context, req *v1.GetGameReq) (*v1.GetGameReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGame")
	defer span.End()

	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*v1.GetGameReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetGameReply_Image{Url: i.URL})
	}
	return &v1.GetGameReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdateGame(ctx context.Context, req *v1.UpdateGameReq) (*v1.UpdateGameReply, error) {
	// b := &biz.Game{
	// 	Id:          req.Id,
	// 	Name:        req.Name,
	// 	Description: req.Description,
	// 	Count:       req.Count,
	// 	Images:      make([]biz.Image, 0),
	// }
	// for _, x := range req.Image {
	// 	b.Images = append(b.Images, biz.Image{URL: x.Url})
	// }
	// x, err := s.bc.Update(ctx, b)
	// img := make([]*v1.UpdateGameReply_Image, 0)
	// for _, i := range x.Images {
	// 	img = append(img, &v1.UpdateGameReply_Image{Url: i.URL})
	// }
	// return &v1.UpdateGameReply{
	// 	Id:          x.Id,
	// 	Name:        x.Name,
	// 	Description: x.Description,
	// 	Count:       x.Count,
	// 	Image:       img,
	// }, err
	return nil, nil
}

func (s *CatalogService) ListGame(ctx context.Context, req *v1.ListGameReq) (*v1.ListGameReply, error) {

	// rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	// rs := make([]*v1.ListGameReply_Game, 0)
	// for _, x := range rv {
	// 	img := make([]*v1.ListGameReply_Game_Image, 0)
	// 	for _, i := range x.Images {
	// 		img = append(img, &v1.ListGameReply_Game_Image{Url: i.URL})
	// 	}
	// 	rs = append(rs, &v1.ListGameReply_Game{
	// 		Id:          x.Id,
	// 		Name:        x.Name,
	// 		Description: x.Description,
	// 		Image:       img,
	// 	})
	// }

	// return &v1.ListGameReply{
	// 	Results: rs,
	// }, err
	return nil, nil
}

func (s *CatalogService) HealthCheck(ctx context.Context, req *v1.HealthReq) (*v1.HealthReply, error) {
	return &v1.HealthReply{Status: 1}, nil
}
