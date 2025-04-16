package redkacli

import "time"

func (c *client) Incr(key string) (int, error) {
	return c.db.Str().Incr(key, 1)
}

func (c *client) Decr(key string) (int, error) {
	return c.db.Str().Incr(key, -1)
}

func (c *client) IncrBy(key string, value int) (int, error) {
	return c.db.Str().Incr(key, value)
}

func (c *client) DecrBy(key string, value int) (int, error) {
	return c.db.Str().Incr(key, -value)
}

func (c *client) IncrByFloat(key string, value float64) (float64, error) {
	return c.db.Str().IncrFloat(key, value)
}

func (c *client) DecrByFloat(key string, value float64) (float64, error) {
	return c.db.Str().IncrFloat(key, -value)
}

func (c *client) Get(key string) (Value, error) {
	return c.db.Str().Get(key)
}

func (c *client) GetSet(key string, value interface{}) (old Value, err error) {
	output, err := c.db.Str().SetWith(key, value).Run()
	if err != nil {
		return nil, err
	}
	return output.Prev, nil
}

func (c *client) MGet(keys ...string) (map[string]Value, error) {
	result := make(map[string]Value)
	kv, err := c.db.Str().GetMany(keys...)
	if err != nil {
		return nil, err
	}
	for k, v := range kv {
		result[k] = v
	}
	return result, nil
}

func (c *client) PSetEx(key string, value interface{}, milliseconds int64) error {
	expiration := time.Duration(milliseconds) * time.Millisecond
	return c.db.Str().SetExpires(key, value, expiration)
}

func (c *client) Set(key string, value interface{}) error {
	return c.db.Str().Set(key, value)
}

func (c *client) MSet(values map[string]interface{}) error {
	return c.db.Str().SetMany(values)
}

func (c *client) SetEx(key string, value interface{}, expiration time.Duration) error {
	return c.db.Str().SetExpires(key, value, expiration)
}

func (c *client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	result, err := c.db.Str().SetWith(key, value).IfNotExists().Run()
	if err != nil {
		return false, err
	}
	return result.Created, nil
}

func (c *client) StrLen(key string) (int, error) {
	result, err := c.db.Str().Get(key)
	if err != nil {
		return 0, err
	}
	return len(result.String()), nil
}
