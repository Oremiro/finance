package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

const (
	MaxPoolSize        = 1
	ConnectionAttempts = 2
	ConnectionTimeout  = time.Second
)

type ConnectionString struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	NetLoc   string `json:"netLoc,omitempty"`
	Port     string `json:"port,omitempty"`
	DbName   string `json:"dbName,omitempty"`
}

func (c *ConnectionString) String() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.User, c.Password, c.NetLoc, c.Port, c.DbName)
}

type Option func(*Context)

func WithMaxPoolSize(maxPoolSize int) Option {
	return func(context *Context) {
		context.maxPoolSize = maxPoolSize
	}
}

func WithConnectionAttempts(connectionAttempts int) Option {
	return func(context *Context) {
		context.connectionAttempts = connectionAttempts
	}
}

func WithConnectionTimeout(connectionTimeout time.Duration) Option {
	return func(context *Context) {
		context.connectionTimeout = connectionTimeout
	}
}

type IDisposable interface {
	Dispose()
}

type Context struct {
	maxPoolSize        int
	connectionAttempts int
	connectionTimeout  time.Duration
	Builder            *squirrel.StatementBuilderType
	Pool               *pgxpool.Pool
}

func NewContext(connectionString *ConnectionString, options ...Option) (*Context, error) {
	pg := &Context{
		maxPoolSize:        MaxPoolSize,
		connectionAttempts: ConnectionAttempts,
		connectionTimeout:  ConnectionTimeout,
	}

	for _, opt := range options {
		opt(pg)
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pg.Builder = &builder

	config, err := pgxpool.ParseConfig(connectionString.String())
	config.MaxConns = int32(pg.maxPoolSize)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for pg.connectionAttempts > 0 {
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), config)
		if err == nil {
			break
		}
		time.Sleep(pg.connectionTimeout)
		pg.connectionAttempts--
	}

	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *Context) Dispose() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
