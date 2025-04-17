package redkacli

func (c *client) SAdd(key string, members ...interface{}) (int, error) {
	return c.db.Set().Add(key, members...)
}

func (c *client) SCard(key string) (int, error) {
	return c.db.Set().Len(key)
}

func (c *client) SDiff(keys ...string) ([]string, error) {
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

func (c *client) SDiffStore(destination string, keys ...string) (int, error) {
	return c.db.Set().DiffStore(destination, keys...)
}

func (c *client) SInter(keys ...string) ([]string, error) {
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

func (c *client) SInterStore(destination string, keys ...string) (int, error) {
	return c.db.Set().InterStore(destination, keys...)
}

func (c *client) SIsMember(key string, member interface{}) (bool, error) {
	return c.db.Set().Exists(key, member)
}

func (c *client) SMembers(key string) ([]string, error) {
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

func (c *client) SMove(source, destination string, member interface{}) error {
	return c.db.Set().Move(source, destination, member)
}

func (c *client) SPop(key string) (string, error) {
	value, err := c.db.Set().Pop(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) SRandMember(key string) (string, error) {
	value, err := c.db.Set().Random(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) SRem(key string, members ...interface{}) (int, error) {
	return c.db.Set().Delete(key, members...)
}

func (c *client) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
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

func (c *client) SUnion(keys ...string) ([]string, error) {
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

func (c *client) SUnionStore(destination string, keys ...string) (int, error) {
	return c.db.Set().UnionStore(destination, keys...)
}
