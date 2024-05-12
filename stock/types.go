package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type StockService interface {
	CheckIfItemsAreInStock(context.Context, *pb.CheckIfItemsAreInStockRequest) (*pb.CheckIfItemsAreInStockResponse, error)
	GetItems(context.Context, *pb.GetItemsRequest) (*pb.GetItemsResponse, error)
}

type StockStore interface {
}
