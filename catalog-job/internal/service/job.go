package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CatalogJob) ConsumerJob(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.oc.Consume(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
