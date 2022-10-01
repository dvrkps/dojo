package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type Connection interface {
	Channel() (Channel, error)
	Close() error
}

type connection struct {
	conn *amqp091.Connection
}

func (c *connection) Channel() (Channel, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}
	return &channel{channel: ch}, nil
}

func (c *connection) Close() error {
	return c.conn.Close()
}

func dial(url string) (*connection, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	return &connection{conn: conn}, nil
}
