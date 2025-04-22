package redkacli

func (c *Client) HDel(key string, fields ...string) (int64, error) {
	i, err := c.db.Hash().Delete(key, fields...)
	return int64(i), err
}

func (c *Client) HExists(key string, field string) (bool, error) {
	return c.db.Hash().Exists(key, field)
}

func (c *Client) HGet(key string, field string) (string, error) {
	v, err := c.db.Hash().Get(key, field)
	if err != nil {
		return "", err
	}
	return v.String(), err
}

func (c *Client) HGetAll(key string) (map[string]string, error) {
	result := make(map[string]string)
	kv, err := c.db.Hash().Items(key)
	if err != nil {
		return nil, err
	}
	for k, v := range kv {
		result[k] = v.String()
	}
	return result, nil
}

func (c *Client) HIncrBy(key string, field string, increment int) (int64, error) {
	v, err := c.db.Hash().Incr(key, field, increment)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func (c *Client) HIncrByFloat(key string, field string, increment float64) (float64, error) {
	return c.db.Hash().IncrFloat(key, field, increment)
}

func (c *Client) HKeys(key string) ([]string, error) {
	return c.db.Hash().Fields(key)
}

func (c *Client) HLen(key string) (int64, error) {
	v, err := c.db.Hash().Len(key)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func (c *Client) HMGet(key string, fields ...string) (map[string]string, error) {
	kv, err := c.db.Hash().GetMany(key, fields...)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for k, v := range kv {
		result[k] = v.String()
	}
	return result, nil
}

func (c *Client) HMSet(key string, values map[string]interface{}) (bool, error) {
	_, err := c.db.Hash().SetMany(key, values)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) HScan(key string, cursor int, match string, count int) (map[string]string, int, error) {
	fields, err := c.db.Hash().Scan(key, cursor, match, count)
	if err != nil {
		return nil, 0, err
	}
	result := make(map[string]string)
	for _, item := range fields.Items {
		result[item.Field] = item.Value.String()
	}
	return result, fields.Cursor, err
}

func (c *Client) HSet(key string, fields map[string]interface{}) (int64, error) {
	i, err := c.db.Hash().SetMany(key, fields)
	return int64(i), err
}

func (c *Client) HSetNX(key string, field string, value interface{}) (bool, error) {
	return c.db.Hash().SetNotExists(key, field, value)
}

func (c *Client) HVals(key string) ([]string, error) {
	values, err := c.db.Hash().Values(key)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.String())
	}
	return result, nil
}
