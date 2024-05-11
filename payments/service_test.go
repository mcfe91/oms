package main

import (
	"context"
	"testing"

	"github.com/mcfe91/commons/api"
	inmemRegistry "github.com/mcfe91/commons/discovery/inmem"
	"github.com/mcfe91/oms-payments/gateway"
	"github.com/mcfe91/oms-payments/processor/inmem"
)

func TestService(t *testing.T) {
	processor := inmem.NewInmem()
	registry := inmemRegistry.NewRegsitry()

	gateway := gateway.NewGRPCGateway(registry)
	svc := NewService(processor, gateway)

	t.Run("should create a payment link", func(t *testing.T) {
		link, err := svc.CreatePayment(context.Background(), &api.Order{})
		if err != nil {
			t.Errorf("CreatePayment() error = %v, want nil", err)
		}
		if link == "" {
			t.Errorf("CreatePayment() link is empty")
		}
	})
}
