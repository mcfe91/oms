package processor

import pb "github.com/mcfe91/commons/api"

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
