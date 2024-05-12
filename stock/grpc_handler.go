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
	inStock, items, err := h.service.CheckIfItemsAreInStock(ctx, p.Items)
	if err != nil {
		return nil, err
	}

	return &pb.CheckIfItemsAreInStockResponse{
		InStock: inStock,
		Items:   items,
	}, nil
}

func (h *grpcHandler) GetItems(ctx context.Context, p *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	items, err := h.service.GetItems(ctx, p.ItemIDs)
	if err != nil {
		return nil, err
	}

	return &pb.GetItemsResponse{
		Items: items,
	}, nil
}
