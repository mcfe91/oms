package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type TelemetryMiddleware struct {
	next StockService
}

func NewTelemetryMiddleware(next StockService) *TelemetryMiddleware {
	return &TelemetryMiddleware{next}
}

func (t *TelemetryMiddleware) CheckIfItemsAreInStock(ctx context.Context, p *pb.CheckIfItemsAreInStockRequest) (*pb.CheckIfItemsAreInStockResponse, error) {

	return t.next.CheckIfItemsAreInStock(ctx, p)
}

func (t *TelemetryMiddleware) GetItems(ctx context.Context, p *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return t.next.GetItems(ctx, p)
}
