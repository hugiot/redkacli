package redkacli

import "time"

func (c *Client) Incr(key string) (int64, error) {
	i, err := c.db.Str().Incr(key, 1)
	return int64(i), err
}

func (c *Client) Decr(key string) (int64, error) {
	i, err := c.db.Str().Incr(key, -1)
	return int64(i), err
}

func (c *Client) IncrBy(key string, increment int64) (int64, error) {
	i, err := c.db.Str().Incr(key, int(increment))
	return int64(i), err
}

func (c *Client) DecrBy(key string, increment int64) (int64, error) {
	i, err := c.db.Str().Incr(key, -int(increment))
	return int64(i), err
}

func (c *Client) IncrByFloat(key string, increment float64) (float64, error) {
	return c.db.Str().IncrFloat(key, increment)
}

func (c *Client) DecrByFloat(key string, increment float64) (float64, error) {
	return c.db.Str().IncrFloat(key, -increment)
}

func (c *Client) Get(key string) (string, error) {
	v, err := c.db.Str().Get(key)
	return v.String(), err
}

func (c *Client) GetSet(key string, value interface{}) (old string, err error) {
	out, err := c.db.Str().SetWith(key, value).Run()
	if err != nil {
		return "", err
	}
	return out.Prev.String(), nil
}

func (c *Client) MGet(keys ...string) (map[string]string, error) {
	result := make(map[string]string)
	kv, err := c.db.Str().GetMany(keys...)
	if err != nil {
		return nil, err
	}
	for k, v := range kv {
		result[k] = v.String()
	}
	return result, nil
}

func (c *Client) PSetEx(key string, value interface{}, milliseconds int64) error {
	expiration := time.Duration(milliseconds) * time.Millisecond
	return c.db.Str().SetExpires(key, value, expiration)
}

func (c *Client) Set(key string, value interface{}) error {
	return c.db.Str().Set(key, value)
}

func (c *Client) MSet(kv map[string]interface{}) error {
	return c.db.Str().SetMany(kv)
}

func (c *Client) SetEx(key string, value interface{}, seconds int64) error {
	expiration := time.Duration(seconds) * time.Second
	return c.db.Str().SetExpires(key, value, expiration)
}

func (c *Client) SetNX(key string, value interface{}, seconds int64) (bool, error) {
	cmd := c.db.Str().SetWith(key, value).IfNotExists()
	if seconds != 0 {
		expiration := time.Duration(seconds) * time.Second
		cmd = cmd.TTL(expiration)
	}
	result, err := cmd.Run()
	if err != nil {
		return false, err
	}
	return result.Created, nil
}

func (c *Client) StrLen(key string) (int64, error) {
	result, err := c.db.Str().Get(key)
	if err != nil {
		return 0, err
	}
	return int64(len(result.String())), nil
}
