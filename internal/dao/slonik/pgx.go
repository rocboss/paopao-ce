// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/conf"
	pg "github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/auto"
)

var (
	_pgxDB   *pgx.Conn
	_oncePgx sync.Once
)

type pgxSrv struct {
	db *pgx.Conn
	p  pg.Querier
}

func (s *pgxSrv) with(fn func(c context.Context, tx pgx.Tx) error) error {
	ctx := context.Background()
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err = fn(ctx, tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *pgxSrv) withTx(txOptions pgx.TxOptions, fn func(ctx context.Context, tx pgx.Tx) error) error {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err = fn(ctx, tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *pgxSrv) withCtx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err = fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *pgxSrv) withTxCtx(ctx context.Context, txOptions pgx.TxOptions, fn func(tx pgx.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err = fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func newPgxSrv(db *pgx.Conn) *pgxSrv {
	return &pgxSrv{
		db: db,
		p:  pg.New(db),
	}
}

func pgxDB() *pgx.Conn {
	_oncePgx.Do(func() {
		_pgxDB = conf.MustPgxDB()
	})
	return _pgxDB
}
