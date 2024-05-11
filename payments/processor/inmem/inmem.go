package inmem

import pb "github.com/mcfe91/commons/api"

type Inmem struct{}

func NewInmem() *Inmem {
	return &Inmem{}
}

func (i *Inmem) CreatePaymentLink(o *pb.Order) (string, error) {
	return "dummy link", nil
}
