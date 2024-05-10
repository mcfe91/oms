package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
)

type PaymentService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
