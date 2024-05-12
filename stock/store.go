package main

import (
	"context"
	"fmt"

	pb "github.com/mcfe91/commons/api"
)

type store struct {
	stock map[string]*pb.Item
}

func NewStore() *store {
	return &store{
		stock: map[string]*pb.Item{
			"2": {
				ID:       "2",
				Name:     "Potato Chips",
				PriceID:  "price_1PFel2CsyUNgRNLnfGo8VtJz",
				Quantity: 10,
			},
			"1": {
				ID:       "1",
				Name:     "Cheese Burger",
				PriceID:  "price_1PF100CsyUNgRNLnHgZ6933X",
				Quantity: 20,
			},
		},
	}
}

func (s *store) GetItem(ctx context.Context, id string) (*pb.Item, error) {
	for _, item := range s.stock {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}

func (s *store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	var res []*pb.Item
	for _, id := range ids {
		if i, ok := s.stock[id]; ok {
			res = append(res, i)
		}
	}

	return res, nil
}
