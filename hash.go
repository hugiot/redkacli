package redkacli

func (c *client) HDel(key string, fields ...string) (int, error) {
	return c.db.Hash().Delete(key, fields...)
}

func (c *client) HExists(key string, field string) (bool, error) {
	return c.db.Hash().Exists(key, field)
}

func (c *client) HGet(key string, field string) (Value, error) {
	return c.db.Hash().Get(key, field)
}

func (c *client) HGetAll(key string) (map[string]Value, error) {
	result := make(map[string]Value)
	kv, err := c.db.Hash().Items(key)
	if err != nil {
		return nil, err
	}
	for k, v := range kv {
		result[k] = v
	}
	return result, nil
}

func (c *client) HIncrBy(key string, field string, incr int) (int, error) {
	return c.db.Hash().Incr(key, field, incr)
}

func (c *client) HIncrByFloat(key string, field string, incr float64) (float64, error) {
	return c.db.Hash().IncrFloat(key, field, incr)
}

func (c *client) HKeys(key string) ([]string, error) {
	return c.db.Hash().Fields(key)
}

func (c *client) HLen(key string) (int, error) {
	return c.db.Hash().Len(key)
}

func (c *client) HMGet(key string, fields ...string) (map[string]Value, error) {
	kv, err := c.db.Hash().GetMany(key, fields...)
	if err != nil {
		return nil, err
	}
	result := make(map[string]Value)
	for k, v := range kv {
		result[k] = v
	}
	return result, nil
}

func (c *client) HMSet(key string, values map[string]interface{}) (int, error) {
	return c.db.Hash().SetMany(key, values)
}

func (c *client) HScan(key string, cursor int, match string, count int) {

}
