package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type StockService interface {
	CheckIfItemsAreInStock(context.Context, []*pb.ItemsWithQuantity) (bool, []*pb.Item, error)
	GetItems(ctx context.Context, ids []string) ([]*pb.Item, error)
}

type StockStore interface {
	GetItem(ctx context.Context, id string) (*pb.Item, error)
	GetItems(ctx context.Context, ids []string) ([]*pb.Item, error)
}
