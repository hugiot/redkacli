package redkacli

import (
	"errors"
	"github.com/nalgeon/redka"
	"strings"
)

func (c *Client) LIndex(key string, index int) (string, error) {
	value, err := c.db.List().Get(key, index)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) LInsert(key string, op string, pivot string, value string) (int64, error) {
	switch strings.ToLower(op) {
	case "before":
		i, err := c.db.List().InsertBefore(key, pivot, value)
		return int64(i), err
	case "after":
		i, err := c.db.List().InsertAfter(key, pivot, value)
		return int64(i), err
	}
	return 0, errors.New("unknown op")
}

func (c *Client) LLen(key string) (int64, error) {
	i, err := c.db.List().Len(key)
	return int64(i), err
}

func (c *Client) LPop(key string) (string, error) {
	value, err := c.db.List().PopFront(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) LPush(key string, elements ...interface{}) (int64, error) {
	var i int
	var err error
	err = c.db.Update(func(tx *redka.Tx) error {
		for _, element := range elements {
			i, err = tx.List().PushFront(key, element)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return int64(i), err
}

func (c *Client) LRange(key string, start int64, stop int64) ([]string, error) {
	values, err := c.db.List().Range(key, int(start), int(stop))
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.String())
	}
	return result, nil
}

func (c *Client) LRem(key string, count int, element string) (int64, error) {
	i, err := c.db.List().DeleteFront(key, element, count)
	return int64(i), err
}

func (c *Client) LSet(key string, index int, element string) error {
	return c.db.List().Set(key, index, element)
}

func (c *Client) LTrim(key string, start, stop int64) (int64, error) {
	i, err := c.db.List().Trim(key, int(start), int(stop))
	return int64(i), err
}

func (c *Client) RPop(key string) (string, error) {
	value, err := c.db.List().PopBack(key)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) RPopLPush(source string, destination string) (string, error) {
	value, err := c.db.List().PopBackPushFront(source, destination)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (c *Client) RPush(key string, elements ...interface{}) (int64, error) {
	var i int
	var err error
	err = c.db.Update(func(tx *redka.Tx) error {
		for _, element := range elements {
			i, err = tx.List().PushBack(key, element)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return i, err
}
