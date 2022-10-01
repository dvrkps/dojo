package rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type Channel interface {
	Close() error
	Consume(p *ConsumeParams) (<-chan Delivery, error)
	PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg Publishing) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
}

type channel struct {
	channel *amqp091.Channel
}

func (c *channel) Close() error {
	return c.channel.Close()
}

type ConsumeParams struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      Table
}

func (c *channel) Consume(p *ConsumeParams) (<-chan Delivery, error) {
	return c.channel.Consume(p.Queue, p.Consumer, p.AutoAck, p.Exclusive, p.NoLocal, p.NoWait, p.Args)
}

func (c *channel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg Publishing) error {
	return c.channel.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}

func (c *channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error) {
	return c.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}
