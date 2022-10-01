package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dvrkps/dojo/rmqclient"
	"github.com/dvrkps/dojo/rmqclient/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func mainSend(ctx context.Context, delay time.Duration) error {
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

	var i int
	var done bool
	for {
		if done {
			break
		}

		select {
		case <-ctx.Done():
			done = true
		default:
			time.Sleep(delay)

			i++

			ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)

			body := "Hello World! " + strconv.Itoa(i)
			err = ch.PublishWithContext(ctxTimeout,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp091.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})

			cancel()

			if err != nil {
				return fmt.Errorf("publish %v: %v", q.Name, err)
			}

			log.Printf(" [x] Sent %s\n", body)
		}
	}

	return nil
}
