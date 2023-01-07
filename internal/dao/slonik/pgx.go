// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/conf"
	dbr "github.com/rocboss/paopao-ce/internal/dao/slonik/ce/postgres"
)

var (
	_pgxDB   *pgx.Conn
	_oncePgx sync.Once
)

type pgxServant struct {
	db *pgx.Conn
	q  dbr.Querier
}

func (s *pgxServant) begin(ctx context.Context) (pgx.Tx, dbr.Querier, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}
	return tx, dbr.New(tx), nil
}

func (s *pgxServant) beingTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, dbr.Querier, error) {
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, nil, err
	}
	return tx, dbr.New(tx), nil
}

func (s *pgxServant) with(handle func(dbr.Querier) error) error {
	ctx := context.Background()
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	if err = handle(dbr.New(tx)); err == nil {
		return tx.Commit(ctx)
	}
	return tx.Rollback(ctx)
}

func (s *pgxServant) withTx(txOptions pgx.TxOptions, handle func(dbr.Querier) error) error {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}
	if err = handle(dbr.New(tx)); err == nil {
		return tx.Commit(ctx)
	}
	return tx.Rollback(ctx)
}

func (s *pgxServant) withCtx(ctx context.Context, handle func(dbr.Querier) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	if err = handle(dbr.New(tx)); err == nil {
		return tx.Commit(ctx)
	}
	return tx.Rollback(ctx)
}

func (s *pgxServant) withTxCtx(ctx context.Context, txOptions pgx.TxOptions, handle func(dbr.Querier) error) error {
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}
	if err = handle(dbr.New(tx)); err == nil {
		return tx.Commit(ctx)
	}
	return tx.Rollback(ctx)
}

func newPgxServant(db *pgx.Conn) *pgxServant {
	return &pgxServant{
		db: db,
		q:  dbr.New(db),
	}
}

func pgxDB() *pgx.Conn {
	_oncePgx.Do(func() {
		_pgxDB = conf.MustPgxDB()
	})
	return _pgxDB
}
