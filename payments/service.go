package main

import (
	"context"

	pb "github.com/mcfe91/commons/api"
	"github.com/mcfe91/oms-payments/processor"
)

type service struct {
	processor processor.PaymentProcessor
}

func NewService(processor processor.PaymentProcessor) *service {
	return &service{processor}
}

func (s *service) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(o)
	if err != nil {
		return "", nil
	}

	// update order with the link

	return link, nil
}
