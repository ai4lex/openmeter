// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent1/db/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent1/db/example1"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Example1 is the client for interacting with the Example1 builders.
	Example1 *Example1Client
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Example1 = NewExample1Client(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("db: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("db: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Example1: NewExample1Client(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Example1: NewExample1Client(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Example1.
//		Query().
//		Count(ctx)
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
	c.Example1.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Example1.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *Example1Mutation:
		return c.Example1.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("db: unknown mutation type %T", m)
	}
}

// Example1Client is a client for the Example1 schema.
type Example1Client struct {
	config
}

// NewExample1Client returns a client for the Example1 from the given config.
func NewExample1Client(c config) *Example1Client {
	return &Example1Client{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `example1.Hooks(f(g(h())))`.
func (c *Example1Client) Use(hooks ...Hook) {
	c.hooks.Example1 = append(c.hooks.Example1, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `example1.Intercept(f(g(h())))`.
func (c *Example1Client) Intercept(interceptors ...Interceptor) {
	c.inters.Example1 = append(c.inters.Example1, interceptors...)
}

// Create returns a builder for creating a Example1 entity.
func (c *Example1Client) Create() *Example1Create {
	mutation := newExample1Mutation(c.config, OpCreate)
	return &Example1Create{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Example1 entities.
func (c *Example1Client) CreateBulk(builders ...*Example1Create) *Example1CreateBulk {
	return &Example1CreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *Example1Client) MapCreateBulk(slice any, setFunc func(*Example1Create, int)) *Example1CreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &Example1CreateBulk{err: fmt.Errorf("calling to Example1Client.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*Example1Create, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &Example1CreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Example1.
func (c *Example1Client) Update() *Example1Update {
	mutation := newExample1Mutation(c.config, OpUpdate)
	return &Example1Update{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *Example1Client) UpdateOne(e *Example1) *Example1UpdateOne {
	mutation := newExample1Mutation(c.config, OpUpdateOne, withExample1(e))
	return &Example1UpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *Example1Client) UpdateOneID(id string) *Example1UpdateOne {
	mutation := newExample1Mutation(c.config, OpUpdateOne, withExample1ID(id))
	return &Example1UpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Example1.
func (c *Example1Client) Delete() *Example1Delete {
	mutation := newExample1Mutation(c.config, OpDelete)
	return &Example1Delete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *Example1Client) DeleteOne(e *Example1) *Example1DeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *Example1Client) DeleteOneID(id string) *Example1DeleteOne {
	builder := c.Delete().Where(example1.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &Example1DeleteOne{builder}
}

// Query returns a query builder for Example1.
func (c *Example1Client) Query() *Example1Query {
	return &Example1Query{
		config: c.config,
		ctx:    &QueryContext{Type: TypeExample1},
		inters: c.Interceptors(),
	}
}

// Get returns a Example1 entity by its id.
func (c *Example1Client) Get(ctx context.Context, id string) (*Example1, error) {
	return c.Query().Where(example1.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *Example1Client) GetX(ctx context.Context, id string) *Example1 {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *Example1Client) Hooks() []Hook {
	return c.hooks.Example1
}

// Interceptors returns the client interceptors.
func (c *Example1Client) Interceptors() []Interceptor {
	return c.inters.Example1
}

func (c *Example1Client) mutate(ctx context.Context, m *Example1Mutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&Example1Create{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&Example1Update{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&Example1UpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&Example1Delete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("db: unknown Example1 mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Example1 []ent.Hook
	}
	inters struct {
		Example1 []ent.Interceptor
	}
)