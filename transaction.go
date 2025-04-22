package redkacli

import (
	"context"
	"github.com/nalgeon/redka"
)

func (c *Client) Transaction(f func(tx *redka.Tx) error) error {
	return c.db.Update(f)
}

func (c *Client) TransactionContext(ctx context.Context, f func(tx *redka.Tx) error) error {
	return c.db.UpdateContext(ctx, f)
}
