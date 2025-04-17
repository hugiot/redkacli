package redkacli

func (c *client) ZAdd(key string, members map[interface{}]float64) (int64, error) {
	i, err := c.db.ZSet().AddMany(key, members)
	return int64(i), err
}

func (c *client) ZCard(key string) (int64, error) {
	i, err := c.db.ZSet().Len(key)
	return int64(i), err
}

func (c *client) ZCount(key string, min, max float64) (int64, error) {
	i, err := c.db.ZSet().Count(key, min, max)
	return int64(i), err
}

func (c *client) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return c.db.ZSet().Incr(key, member, increment)
}

func (c *client) ZInter(keys ...string) ([]string, error) {
	members, err := c.db.ZSet().InterWith(keys...).Run()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZInterStore(destination string, keys ...string) (int64, error) {
	i, err := c.db.ZSet().InterWith(keys...).Dest(destination).Store()
	return int64(i), err
}

func (c *client) ZRange(key string, start, stop int64) ([]string, error) {
	members, err := c.db.ZSet().Range(key, int(start), int(stop))
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZRangeByScore(key string, min, max float64) ([]string, error) {
	members, err := c.db.ZSet().RangeWith(key).ByScore(min, max).Run()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZRank(key, member string) (int64, error) {
	rank, _, err := c.db.ZSet().GetRank(key, member)
	if err != nil {
		return 0, err
	}
	return int64(rank), nil
}

func (c *client) ZRankWithScore(key, member string) (int64, float64, error) {
	rank, score, err := c.db.ZSet().GetRank(key, member)
	if err != nil {
		return 0, 0, err
	}
	return int64(rank), score, nil
}

func (c *client) ZRem(key string, members ...interface{}) (int64, error) {
	i, err := c.db.ZSet().Delete(key, members...)
	return int64(i), err
}

func (c *client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	i, err := c.db.ZSet().DeleteWith(key).ByRank(int(start), int(stop)).Run()
	return int64(i), err
}

func (c *client) ZRemRangeByScore(key string, min, max float64) (int64, error) {
	i, err := c.db.ZSet().DeleteWith(key).ByScore(min, max).Run()
	return int64(i), err
}

func (c *client) ZRevRange(key string, start, stop int64) ([]string, error) {
	members, err := c.db.ZSet().RangeWith(key).ByRank(int(start), int(stop)).Desc().Run()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZRevRangeByScore(key string, min, max float64) ([]string, error) {
	members, err := c.db.ZSet().RangeWith(key).ByScore(min, max).Desc().Run()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZRevRank(key, member string) (int64, error) {
	rank, _, err := c.db.ZSet().GetRankRev(key, member)
	if err != nil {
		return 0, err
	}
	return int64(rank), nil
}

func (c *client) ZRevRankWithScore(key, member string) (int64, float64, error) {
	rank, score, err := c.db.ZSet().GetRankRev(key, member)
	if err != nil {
		return 0, 0, err
	}
	return int64(rank), score, nil
}

func (c *client) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	result, err := c.db.ZSet().Scan(key, int(cursor), match, int(count))
	if err != nil {
		return nil, 0, err
	}
	members := make([]string, 0, len(result.Items))
	for _, member := range result.Items {
		members = append(members, member.Elem.String())
	}
	return members, uint64(result.Cursor), nil
}

func (c *client) ZScore(key, member string) (float64, error) {
	return c.db.ZSet().GetScore(key, member)
}

func (c *client) ZUnion(keys ...string) ([]string, error) {
	members, err := c.db.ZSet().UnionWith(keys...).Run()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(members))
	for _, member := range members {
		result = append(result, member.Elem.String())
	}
	return result, nil
}

func (c *client) ZUnionWithScore(keys ...string) (map[string]float64, error) {
	members, err := c.db.ZSet().UnionWith(keys...).Run()
	if err != nil {
		return nil, err
	}
	result := make(map[string]float64, len(members))
	for _, member := range members {
		result[member.Elem.String()] = member.Score
	}
	return result, nil
}

func (c *client) ZUnionStore(destination string, keys ...string) (int64, error) {
	i, err := c.db.ZSet().UnionWith(keys...).Dest(destination).Store()
	return int64(i), err
}
