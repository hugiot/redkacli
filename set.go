package redkacli

func (c *Client) SAdd(key string, members ...interface{}) (int, error) {
	return c.db.Set().Add(key, members...)
}

func (c *Client) SCard(key string) (int, error) {
	return c.db.Set().Len(key)
}

func (c *Client) SDiff(keys ...string) ([]string, error) {
	values, err := c.db.Set().Diff(keys...)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.String())
	}
	return result, nil
}

func (c *Client) SDiffStore(destination string, keys ...string) (int, error) {
	return c.db.Set().DiffStore(destination, keys...)
}

func (c *Client) SInter(keys ...string) ([]string, error) {
	values, err := c.db.Set().Inter(keys...)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.String())
	}
	return result, nil
}

func (c *Client) SInterStore(destination string, keys ...string) (int, error) {
	return c.db.Set().InterStore(destination, keys...)
}

func (c *Client) SIsMember(key string, member interface{}) (bool, error) {
	return c.db.Set().Exists(key, member)
}

func (c *Client) SMembers(key string) ([]string, error) {
	values, err := c.db.Set().Items(key)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.String())
	}
	return result, nil
}

func (c *Client) SMove(source, destination string, member interface{}) error {
	return c.db.Set().Move(source, destination, member)
}

func (c *Client) SPop(key string) (string, error) {
	value, err := c.db.Set().Pop(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) SRandMember(key string) (string, error) {
	value, err := c.db.Set().Random(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) SRem(key string, members ...interface{}) (int, error) {
	return c.db.Set().Delete(key, members...)
}

func (c *Client) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	values, err := c.db.Set().Scan(key, int(cursor), match, int(count))
	if err != nil {
		return nil, 0, err
	}
	result := make([]string, 0, len(values.Items))
	for _, v := range values.Items {
		result = append(result, v.String())
	}
	return result, uint64(values.Cursor), nil
}

func (c *Client) SUnion(keys ...string) ([]string, error) {
	values, err := c.db.Set().Union(keys...)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, v.String())
	}
	return result, nil
}

func (c *Client) SUnionStore(destination string, keys ...string) (int, error) {
	return c.db.Set().UnionStore(destination, keys...)
}
