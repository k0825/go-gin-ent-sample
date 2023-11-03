package datasource

import (
	"context"
	"fmt"

	"github.com/k0825/go-gin-ent-sample/config"
	"github.com/k0825/go-gin-ent-sample/ent"
	_ "github.com/lib/pq"
)

type RDBConnection struct {
	client *ent.Client
}

type key int

const (
	txCtxKey key = iota
)

type RDBConnectionInterface interface {
	GetClient() *ent.Client
	GetTx(ctx context.Context) *ent.Client
	Begin(ctx context.Context) (*ent.Tx, error)
	Rollback(tx *ent.Tx) error
	Commit(tx *ent.Tx) error
	RunInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error)
}

func NewRDBConnection(conf *config.Config) (*RDBConnection, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			conf.RDB.Host,
			conf.RDB.Port,
			conf.RDB.UserName,
			conf.RDB.Database,
			conf.RDB.Password))

	if err != nil {
		return nil, err
	}

	return &RDBConnection{client}, nil
}

func (conn *RDBConnection) GetClient() *ent.Client {
	return conn.client
}

func (conn *RDBConnection) GetTx(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(txCtxKey).(*ent.Client)

	if !ok {
		return nil
	}

	return tx
}

func (conn *RDBConnection) Begin(ctx context.Context) (*ent.Tx, error) {
	tx, err := conn.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (conn *RDBConnection) Rollback(tx *ent.Tx) error {
	return tx.Rollback()
}

func (conn *RDBConnection) Commit(tx *ent.Tx) error {
	return tx.Commit()
}

func (conn *RDBConnection) RunInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	txClient := tx.Client()

	ctx = context.WithValue(ctx, txCtxKey, txClient)

	v, err := f(ctx)
	if err != nil {
		if err := conn.Rollback(tx); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := conn.Commit(tx); err != nil {
		return nil, err
	}

	return v, nil
}
