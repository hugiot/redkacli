package redkacli

func (c *Client) Echo(message string) (string, error) {
	return message, nil
}

func (c *Client) LOLWUT() (string, error) {
	// todo something
	return "version", nil
}

func (c *Client) Ping() (string, error) {
	return "PONG", nil
}

func (c *Client) Select(index int) error {
	return nil
}
