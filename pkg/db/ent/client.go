// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/migrate"

	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/pubsubmessage"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// OpLog is the client for interacting with the OpLog builders.
	OpLog *OpLogClient
	// PubsubMessage is the client for interacting with the PubsubMessage builders.
	PubsubMessage *PubsubMessageClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.OpLog = NewOpLogClient(c.config)
	c.PubsubMessage = NewPubsubMessageClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		OpLog:         NewOpLogClient(cfg),
		PubsubMessage: NewPubsubMessageClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		OpLog:         NewOpLogClient(cfg),
		PubsubMessage: NewPubsubMessageClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		OpLog.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.OpLog.Use(hooks...)
	c.PubsubMessage.Use(hooks...)
}

// OpLogClient is a client for the OpLog schema.
type OpLogClient struct {
	config
}

// NewOpLogClient returns a client for the OpLog from the given config.
func NewOpLogClient(c config) *OpLogClient {
	return &OpLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `oplog.Hooks(f(g(h())))`.
func (c *OpLogClient) Use(hooks ...Hook) {
	c.hooks.OpLog = append(c.hooks.OpLog, hooks...)
}

// Create returns a builder for creating a OpLog entity.
func (c *OpLogClient) Create() *OpLogCreate {
	mutation := newOpLogMutation(c.config, OpCreate)
	return &OpLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OpLog entities.
func (c *OpLogClient) CreateBulk(builders ...*OpLogCreate) *OpLogCreateBulk {
	return &OpLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OpLog.
func (c *OpLogClient) Update() *OpLogUpdate {
	mutation := newOpLogMutation(c.config, OpUpdate)
	return &OpLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OpLogClient) UpdateOne(ol *OpLog) *OpLogUpdateOne {
	mutation := newOpLogMutation(c.config, OpUpdateOne, withOpLog(ol))
	return &OpLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OpLogClient) UpdateOneID(id uint32) *OpLogUpdateOne {
	mutation := newOpLogMutation(c.config, OpUpdateOne, withOpLogID(id))
	return &OpLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OpLog.
func (c *OpLogClient) Delete() *OpLogDelete {
	mutation := newOpLogMutation(c.config, OpDelete)
	return &OpLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OpLogClient) DeleteOne(ol *OpLog) *OpLogDeleteOne {
	return c.DeleteOneID(ol.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *OpLogClient) DeleteOneID(id uint32) *OpLogDeleteOne {
	builder := c.Delete().Where(oplog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OpLogDeleteOne{builder}
}

// Query returns a query builder for OpLog.
func (c *OpLogClient) Query() *OpLogQuery {
	return &OpLogQuery{
		config: c.config,
	}
}

// Get returns a OpLog entity by its id.
func (c *OpLogClient) Get(ctx context.Context, id uint32) (*OpLog, error) {
	return c.Query().Where(oplog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OpLogClient) GetX(ctx context.Context, id uint32) *OpLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OpLogClient) Hooks() []Hook {
	hooks := c.hooks.OpLog
	return append(hooks[:len(hooks):len(hooks)], oplog.Hooks[:]...)
}

// PubsubMessageClient is a client for the PubsubMessage schema.
type PubsubMessageClient struct {
	config
}

// NewPubsubMessageClient returns a client for the PubsubMessage from the given config.
func NewPubsubMessageClient(c config) *PubsubMessageClient {
	return &PubsubMessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pubsubmessage.Hooks(f(g(h())))`.
func (c *PubsubMessageClient) Use(hooks ...Hook) {
	c.hooks.PubsubMessage = append(c.hooks.PubsubMessage, hooks...)
}

// Create returns a builder for creating a PubsubMessage entity.
func (c *PubsubMessageClient) Create() *PubsubMessageCreate {
	mutation := newPubsubMessageMutation(c.config, OpCreate)
	return &PubsubMessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PubsubMessage entities.
func (c *PubsubMessageClient) CreateBulk(builders ...*PubsubMessageCreate) *PubsubMessageCreateBulk {
	return &PubsubMessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PubsubMessage.
func (c *PubsubMessageClient) Update() *PubsubMessageUpdate {
	mutation := newPubsubMessageMutation(c.config, OpUpdate)
	return &PubsubMessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PubsubMessageClient) UpdateOne(pm *PubsubMessage) *PubsubMessageUpdateOne {
	mutation := newPubsubMessageMutation(c.config, OpUpdateOne, withPubsubMessage(pm))
	return &PubsubMessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PubsubMessageClient) UpdateOneID(id uint32) *PubsubMessageUpdateOne {
	mutation := newPubsubMessageMutation(c.config, OpUpdateOne, withPubsubMessageID(id))
	return &PubsubMessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PubsubMessage.
func (c *PubsubMessageClient) Delete() *PubsubMessageDelete {
	mutation := newPubsubMessageMutation(c.config, OpDelete)
	return &PubsubMessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PubsubMessageClient) DeleteOne(pm *PubsubMessage) *PubsubMessageDeleteOne {
	return c.DeleteOneID(pm.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PubsubMessageClient) DeleteOneID(id uint32) *PubsubMessageDeleteOne {
	builder := c.Delete().Where(pubsubmessage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PubsubMessageDeleteOne{builder}
}

// Query returns a query builder for PubsubMessage.
func (c *PubsubMessageClient) Query() *PubsubMessageQuery {
	return &PubsubMessageQuery{
		config: c.config,
	}
}

// Get returns a PubsubMessage entity by its id.
func (c *PubsubMessageClient) Get(ctx context.Context, id uint32) (*PubsubMessage, error) {
	return c.Query().Where(pubsubmessage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PubsubMessageClient) GetX(ctx context.Context, id uint32) *PubsubMessage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PubsubMessageClient) Hooks() []Hook {
	hooks := c.hooks.PubsubMessage
	return append(hooks[:len(hooks):len(hooks)], pubsubmessage.Hooks[:]...)
}
