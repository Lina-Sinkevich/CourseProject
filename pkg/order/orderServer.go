package order

import (
	api "CourseProject/pkg/api"
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) MakeOrder(ctx context.Context, req *api.OrderInfo) (*api.OrderStatus, error) {
	return &api.OrderStatus{Status: req.GetFio() + req.GetGameName()}, nil
}
