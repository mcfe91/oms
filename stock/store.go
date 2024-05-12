package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type store struct {
}

func NewStore() *store {
	return &store{}
}

func (s *store) GetItem(ctx context.Context, id string) (*pb.Item, error) {
	return nil, nil
}

func (s *store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	return nil, nil
}
