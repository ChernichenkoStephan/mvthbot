package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NewConnection(db DB) *Connection {
	return &Connection{
		DB: db,
	}
}

func withConnection(ctx context.Context, db *sqlx.DB) (*Connection, context.Context) {
	if conn, ok := ctx.Value(connKey{}).(*Connection); ok {
		return conn, ctx
	}
	conn := NewConnection(db.MustBegin())
	return conn, context.WithValue(ctx, connKey{}, conn)
}

func getConnection(ctx context.Context, db *sqlx.DB) *Connection {
	if conn, ok := ctx.Value(connKey{}).(*Connection); ok {
		return conn
	}
	return NewConnection(db)
}

func (c *Connection) Bind() {
	c.mx.Lock()
	defer c.mx.Unlock()

	if _, ok := c.DB.(*sqlx.Tx); ok {
		c.refCount++
	}

}

func (c *Connection) Release() bool {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.refCount > 1 {
		c.refCount--
	} else {
		if tx, ok := c.DB.(*sqlx.Tx); ok {
			tx.Commit()
			return true
		}
	}

	return false
}
