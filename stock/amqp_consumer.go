package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/mcfe91/commons/api"
	"github.com/mcfe91/commons/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type consumer struct {
	service StockService
}

func NewConsumer(service StockService) *consumer {
	return &consumer{service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		q.Name,                // queue name
		"",                    // routing key
		broker.OrderPaidEvent, // exchange
		false,                 // no-wait
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {

			ctx := broker.ExtractAMQPHeader(context.Background(), d.Headers)

			tr := otel.Tracer("amqp")
			_, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - consumer - %s", q.Name))

			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", err)
				continue
			}
			log.Printf("received message: %s", d.Body)

			orderID := string(d.Body)

			log.Printf("order received: %s", orderID)

			messageSpan.End()

			// TODO: do something with message
			// _, err := c.service.UpdateOrder(context.Background(), o)
			// if err != nil {
			// 	log.Printf("failed to update order: %v", err)

			// 	if err := broker.HandleRetry(ch, &d); err != nil {
			// 		log.Printf("error handling retry: %v", err)
			// 	}

			// 	continue
			// }

			// messageSpan.AddEvent("order.updated")
			// messageSpan.End()

			// log.Printf("order has been updated from AMQP")
			// d.Ack(false)
		}
	}()

	log.Println("AMQP listening...")
	<-forever
}
