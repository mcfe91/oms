package gateway

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type OrdersGateway interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
}
