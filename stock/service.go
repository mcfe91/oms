package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type service struct {
	store StockStore
}

func NewService(store StockStore) *service {
	return &service{store}
}

func (s *service) CheckIfItemsAreInStock(ctx context.Context, p *pb.CheckIfItemsAreInStockRequest) (*pb.CheckIfItemsAreInStockResponse, error) {
	return nil, nil
}

func (s *service) GetItems(ctx context.Context, p *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return nil, nil
}
