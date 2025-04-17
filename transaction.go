package redkacli

import (
	"context"
	"github.com/nalgeon/redka"
)

func (c *client) Transaction(f func(tx *redka.Tx) error) error {
	return c.db.Update(f)
}

func (c *client) TransactionContext(ctx context.Context, f func(tx *redka.Tx) error) error {
	return c.db.UpdateContext(ctx, f)
}
