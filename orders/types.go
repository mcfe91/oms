package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type OrderService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
}

type OrderStore interface {
	Create(context.Context) error
}
