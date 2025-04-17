package redkacli

import (
	"errors"
	"github.com/nalgeon/redka"
	"time"
)

func (c *client) DBSize() (int, error) {
	return c.db.Key().Len()
}

func (c *client) Del(keys ...string) (int, error) {
	return c.db.Key().Delete(keys...)
}

func (c *client) Exists(key string) (bool, error) {
	return c.db.Key().Exists(key)
}

func (c *client) Expire(key string, expiration time.Duration) error {
	return c.db.Key().Expire(key, expiration)
}

func (c *client) ExpireAt(key string, expiration time.Time) error {
	return c.db.Key().ExpireAt(key, expiration)
}

func (c *client) FlushAll() error {
	return c.db.Key().DeleteAll()
}

func (c *client) FlushDB() error {
	return c.db.Key().DeleteAll()
}

func (c *client) Keys(pattern string) ([]string, error) {
	keys, err := c.db.Key().Keys(pattern)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key.Key)
	}
	return result, nil
}

func (c *client) Persist(key string) error {
	return c.db.Key().Persist(key)
}

func (c *client) PExpire(key string, expiration time.Duration) error {
	return c.db.Key().Expire(key, expiration)
}

func (c *client) PExpireAt(key string, expiration time.Time) error {
	return c.db.Key().ExpireAt(key, expiration)
}

func (c *client) RandomKey() (string, error) {
	key, err := c.db.Key().Random()
	if err != nil {
		return "", err
	}
	return key.Key, nil
}

func (c *client) Rename(key, newKey string) error {
	return c.db.Key().Rename(key, newKey)
}

func (c *client) RenameNX(key, newKey string) (bool, error) {
	return c.db.Key().RenameNotExists(key, newKey)
}

func (c *client) Scan(cursor int, match string, count int) ([]string, int, error) {
	list, err := c.db.Key().Scan(cursor, match, 0, count)
	if err != nil {
		return nil, 0, err
	}
	result := make([]string, 0, len(list.Keys))
	for _, item := range list.Keys {
		result = append(result, item.Key)
	}
	return result, list.Cursor, nil
}

func (c *client) TTL(key string) (int64, error) {
	k, err := c.db.Key().Get(key)
	if err != nil {
		if errors.Is(err, redka.ErrNotFound) {
			return -2, nil
		}
		return 0, err
	}
	if k.ETime == nil {
		return -1, nil
	}
	return *k.ETime / 1000, nil
}

func (c *client) Type(key string) (string, error) {
	k, err := c.db.Key().Get(key)
	if err != nil {
		return "", err
	}
	return k.TypeName(), nil
}
