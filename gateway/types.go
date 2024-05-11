package main

import pb "github.com/mcfe91/commons/api"

// TODO: should this be 'response'
type CreateOrderRequest struct {
	Order         *pb.Order `"json":order`
	RedirectToUrl string    `"json":redirectToUrl`
}
