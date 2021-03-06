// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/efectn/go-orm-benchmarks/benchs/ent/migrate"

	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Model is the client for interacting with the Model builders.
	Model *ModelClient
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
	c.Model = NewModelClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Model:  NewModelClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:    ctx,
		config: cfg,
		Model:  NewModelClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Model.
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
	c.Model.Use(hooks...)
}

// ModelClient is a client for the Model schema.
type ModelClient struct {
	config
}

// NewModelClient returns a client for the Model from the given config.
func NewModelClient(c config) *ModelClient {
	return &ModelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `model.Hooks(f(g(h())))`.
func (c *ModelClient) Use(hooks ...Hook) {
	c.hooks.Model = append(c.hooks.Model, hooks...)
}

// Create returns a create builder for Model.
func (c *ModelClient) Create() *ModelCreate {
	mutation := newModelMutation(c.config, OpCreate)
	return &ModelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Model entities.
func (c *ModelClient) CreateBulk(builders ...*ModelCreate) *ModelCreateBulk {
	return &ModelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Model.
func (c *ModelClient) Update() *ModelUpdate {
	mutation := newModelMutation(c.config, OpUpdate)
	return &ModelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ModelClient) UpdateOne(m *Model) *ModelUpdateOne {
	mutation := newModelMutation(c.config, OpUpdateOne, withModel(m))
	return &ModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ModelClient) UpdateOneID(id int) *ModelUpdateOne {
	mutation := newModelMutation(c.config, OpUpdateOne, withModelID(id))
	return &ModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Model.
func (c *ModelClient) Delete() *ModelDelete {
	mutation := newModelMutation(c.config, OpDelete)
	return &ModelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ModelClient) DeleteOne(m *Model) *ModelDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ModelClient) DeleteOneID(id int) *ModelDeleteOne {
	builder := c.Delete().Where(model.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ModelDeleteOne{builder}
}

// Query returns a query builder for Model.
func (c *ModelClient) Query() *ModelQuery {
	return &ModelQuery{
		config: c.config,
	}
}

// Get returns a Model entity by its id.
func (c *ModelClient) Get(ctx context.Context, id int) (*Model, error) {
	return c.Query().Where(model.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ModelClient) GetX(ctx context.Context, id int) *Model {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ModelClient) Hooks() []Hook {
	return c.hooks.Model
}
