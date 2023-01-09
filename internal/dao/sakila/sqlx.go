// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

var (
	_db   *sqlx.DB
	_once sync.Once
)

type sqlxServant struct {
	db *sqlx.DB
}

func (s *sqlxServant) with(handle func(*sqlx.Tx) error) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	if err = handle(tx); err == nil {
		return tx.Commit()
	}
	return tx.Rollback()
}

func (s *sqlxServant) withTx(ctx context.Context, opts *sql.TxOptions, handle func(*sqlx.Tx) error) error {
	tx, err := s.db.BeginTxx(ctx, opts)
	if err != nil {
		return err

	}
	if err = handle(tx); err == nil {
		return tx.Commit()
	}
	return tx.Rollback()
}

func newSqlxServant(db *sqlx.DB) *sqlxServant {
	return &sqlxServant{
		db: db,
	}
}

func sqlxDB() *sqlx.DB {
	_once.Do(func() {
		_db = conf.MustSqlxDB()
	})
	return _db
}

func c(query string) *sqlx.Stmt {
	db := sqlxDB()
	stmt, err := db.Preparex(db.Rebind(t(query)))
	if err != nil {
		logrus.Fatalf("prepare query(%s) error: %s", query, err)
	}
	return stmt
}

func n(query string) *sqlx.NamedStmt {
	db := sqlxDB()
	stmt, err := db.PrepareNamed(t(query))
	if err != nil {
		logrus.Fatalf("prepare named query(%s) error: %s", query, err)
	}
	return stmt
}

// t repace table prefix for query
func t(query string) string {
	return strings.Replace(query, "@", conf.DatabaseSetting.TablePrefix, -1)
}
