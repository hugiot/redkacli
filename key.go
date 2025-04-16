package redkacli

import "time"

func (c *client) Expire(key string, expiration time.Duration) error {
	return c.db.Key().Expire(key, expiration)
}
