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

func (s *service) CheckIfItemsAreInStock(ctx context.Context, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error) {
	return false, nil, nil
}

func (s *service) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	return nil, nil
}
