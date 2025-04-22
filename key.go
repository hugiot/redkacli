package redkacli

import (
	"errors"
	"github.com/nalgeon/redka"
	"time"
)

func (c *Client) DBSize() (int64, error) {
	i, err := c.db.Key().Len()
	return int64(i), err
}

func (c *Client) Del(keys ...string) (int64, error) {
	i, err := c.db.Key().Delete(keys...)
	return int64(i), err
}

func (c *Client) Exists(key string) (bool, error) {
	return c.db.Key().Exists(key)
}

func (c *Client) Expire(key string, seconds int64) error {
	expiration := time.Duration(seconds) * time.Second
	return c.db.Key().Expire(key, expiration)
}

func (c *Client) ExpireAt(key string, timestamp int64) error {
	expiration := time.Unix(timestamp, 0)
	return c.db.Key().ExpireAt(key, expiration)
}

func (c *Client) FlushAll() error {
	return c.db.Key().DeleteAll()
}

func (c *Client) FlushDB() error {
	return c.db.Key().DeleteAll()
}

func (c *Client) Keys(pattern string) ([]string, error) {
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

func (c *Client) Persist(key string) error {
	return c.db.Key().Persist(key)
}

func (c *Client) PExpire(key string, milliseconds int64) error {
	expiration := time.Duration(milliseconds) * time.Millisecond
	return c.db.Key().Expire(key, expiration)
}

func (c *Client) PExpireAt(key string, timestamp int64) error {
	expiration := time.Unix(timestamp, 0)
	return c.db.Key().ExpireAt(key, expiration)
}

func (c *Client) RandomKey() (string, error) {
	key, err := c.db.Key().Random()
	if err != nil {
		return "", err
	}
	return key.Key, nil
}

func (c *Client) Rename(key, newKey string) error {
	return c.db.Key().Rename(key, newKey)
}

func (c *Client) RenameNX(key, newKey string) (bool, error) {
	return c.db.Key().RenameNotExists(key, newKey)
}

func (c *Client) Scan(cursor int, match string, count int) ([]string, int, error) {
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

func (c *Client) TTL(key string) (int64, error) {
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

func (c *Client) Type(key string) (string, error) {
	k, err := c.db.Key().Get(key)
	if err != nil {
		return "", err
	}
	return k.TypeName(), nil
}
