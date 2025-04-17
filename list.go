package redkacli

import (
	"errors"
	"github.com/nalgeon/redka"
	"strings"
)

func (c *client) LIndex(key string, index int) (string, error) {
	value, err := c.db.List().Get(key, index)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) LInsert(key string, op string, pivot string, value string) (int, error) {
	switch strings.ToLower(op) {
	case "before":
		return c.db.List().InsertBefore(key, pivot, value)
	case "after":
		return c.db.List().InsertAfter(key, pivot, value)
	}
	return 0, errors.New("unknown op")
}

func (c *client) LLen(key string) (int, error) {
	return c.db.List().Len(key)
}

func (c *client) LPop(key string) (string, error) {
	value, err := c.db.List().PopFront(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) LPush(key string, values ...string) (int, error) {
	var i int
	var err error
	err = c.db.Update(func(tx *redka.Tx) error {
		for _, value := range values {
			i, err = tx.List().PushFront(key, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return i, err
}

func (c *client) LRange(key string, start int, stop int) ([]string, error) {
	values, err := c.db.List().Range(key, start, stop)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.String())
	}
	return result, nil
}

func (c *client) LRem(key string, count int, value string) (int, error) {
	return c.db.List().DeleteFront(key, value, count)
}

func (c *client) LSet(key string, index int, value string) error {
	return c.db.List().Set(key, index, value)
}

func (c *client) LTrim(key string, start, stop int) (int, error) {
	return c.db.List().Trim(key, start, stop)
}

func (c *client) RPop(key string) (string, error) {
	value, err := c.db.List().PopBack(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) RPopLPush(source string, destination string) (string, error) {
	value, err := c.db.List().PopBackPushFront(source, destination)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *client) RPush(key string, values ...string) (int, error) {
	var i int
	var err error
	err = c.db.Update(func(tx *redka.Tx) error {
		for _, value := range values {
			i, err = tx.List().PushFront(key, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return i, err
}
