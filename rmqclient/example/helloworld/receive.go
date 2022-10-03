package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dvrkps/dojo/rmqclient"
	"github.com/dvrkps/dojo/rmqclient/rabbitmq"
)

func mainReceive(ctx context.Context) error {
	c := rabbitmq.NewClient()

	conn, err := c.Dial(rmqclient.URL)
	if err != nil {
		return fmt.Errorf("dial: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("queue declare: %v", err)
	}

	msgs, err := ch.Consume(&rabbitmq.ConsumeParameters{
		QueueName: q.Name,
		Consumer:  "",
		AutoAck:   true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Table:     nil,
	})

	if err != nil {
		return fmt.Errorf("consume %v: %v", q.Name, err)
	}

	done := make(chan struct{})

	go func() {
		log.Println("start receiving")
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
		close(done)
	}()

	<-done

	return nil
}
