package repo

import (
	"context"
	"database/sql"
)

type ITxRepo interface {
	Begin(ctx context.Context) (*sql.Tx, error)
}

type TxRepo struct {
	db *sql.DB
}

func NewTxRepo(db *sql.DB) *TxRepo {
	return &TxRepo{db: db}
}

func (r *TxRepo) Begin(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}
