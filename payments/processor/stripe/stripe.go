package stripe

import (
	"fmt"
	"log"

	common "github.com/mcfe91/commons"
	pb "github.com/mcfe91/commons/api"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

var (
	gatewayAddr = common.EnvString("GATEWAY_HTTP_ADDRESS", "http://localhost:8080")
)

type Stripe struct {
}

func NewProcessor() *Stripe {
	return &Stripe{}
}

func (s *Stripe) CreatePaymentLink(o *pb.Order) (string, error) {
	log.Printf("creating payment link for order %v", o)

	gatewaySuccessURL := fmt.Sprintf("%s/success.html?customerID=%s&orderID=%s", gatewayAddr, o.CustomerID, o.ID)
	gatewayCancelURL := fmt.Sprintf("%s/cancel.html", gatewayAddr)

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range o.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:  items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessURL),
		CancelURL:  stripe.String(gatewayCancelURL),
	}
	result, err := session.New(params)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
