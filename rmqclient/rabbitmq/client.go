package rabbitmq

type Client struct {
	conn Connection
}

func NewClient() Client {
	return Client{}
}

func (c *Client) Dial(url string) (Connection, error) {
	conn, err := dial(url)
	if err != nil {
		return nil, err
	}
	c.conn = conn

	return c.conn, nil
}
