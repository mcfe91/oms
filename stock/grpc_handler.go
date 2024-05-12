package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedStockServiceServer

	service StockService
	channel *amqp.Channel
}

func NewGRPCHandler(grpcServer *grpc.Server, service StockService, channel *amqp.Channel) {
	handler := &grpcHandler{
		service: service,
		channel: channel,
	}

	pb.RegisterStockServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CheckIfItemsAreInStock(ctx context.Context, p *pb.CheckIfItemsAreInStockRequest) (*pb.CheckIfItemsAreInStockResponse, error) {
	return h.service.CheckIfItemsAreInStock(ctx, p)
}

func (h *grpcHandler) GetItems(ctx context.Context, p *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return h.service.GetItems(ctx, p)
}
