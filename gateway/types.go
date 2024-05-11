package main

import pb "github.com/mcfe91/commons/api"

// TODO: should this be 'response'
type CreateOrderResponse struct {
	Order         *pb.Order `"json":order`
	RedirectToUrl string    `"json":redirectToUrl`
}
