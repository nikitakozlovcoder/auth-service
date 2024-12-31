package dbconnection

import (
	"context"
	"database/sql"
	"log"
	"service/auth/infrastructure/db/dbexecutor"

	_ "github.com/lib/pq"
)

type txKey string

const key txKey = "tx"

type Manager struct {
	db *sql.DB
}

func NewManager(connString string) *Manager {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Manager{db: db}
}

func (m *Manager) BeginTx(ctx context.Context) (context.Context, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	txctx := context.WithValue(ctx, key, tx)
	return txctx, nil
}

func (m *Manager) StopTx(ctx context.Context) (context.Context, error) {
	txctx := context.WithValue(ctx, key, nil)
	return txctx, nil
}

func (m *Manager) Executor(ctx context.Context) dbexecutor.Executor {
	tx, ok := ctx.Value(key).(*sql.Tx)
	if ok && tx != nil {
		return tx
	}

	return m.db
}

func (m *Manager) Close() error {
	return m.db.Close()
}
