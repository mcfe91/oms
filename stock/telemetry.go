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

func (t *TelemetryMiddleware) CheckIfItemsAreInStock(ctx context.Context, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error) {
	return t.next.CheckIfItemsAreInStock(ctx, items)
}

func (t *TelemetryMiddleware) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	return t.next.GetItems(ctx, ids)
}
