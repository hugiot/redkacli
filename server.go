package redkacli

func (c *client) Echo(message string) (string, error) {
	return message, nil
}

func (c *client) LOLWUT() (string, error) {
	// todo something
	return "version", nil
}

func (c *client) Ping() (string, error) {
	return "PONG", nil
}

func (c *client) Select(index int) error {
	return nil
}
