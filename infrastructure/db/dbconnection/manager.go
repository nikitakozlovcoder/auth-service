package dbconnection

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"service/auth/infrastructure/db/dbexecutor"

	_ "github.com/lib/pq"
)

var (
	ErrTxAlreadyStarted = errors.New("transaction already started")
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

func (m *Manager) BeginTx(ctx context.Context) (Transaction, error) {
	_, hasTx := ctx.Value(key).(*sql.Tx)
	if hasTx {
		return Transaction{}, ErrTxAlreadyStarted
	}

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return Transaction{}, err
	}

	txctx := context.WithValue(ctx, key, tx)
	return Transaction{txctx}, nil
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

type Transaction struct {
	Ctx context.Context
}

func (t *Transaction) Commit() error {
	tx := t.Ctx.Value(key).(*sql.Tx)
	return tx.Commit()
}

func (t *Transaction) Rollback() error {
	tx := t.Ctx.Value(key).(*sql.Tx)
	return tx.Rollback()
}
