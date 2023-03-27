// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"context"
	"database/sql"
	_ "embed"
	"strings"

	"github.com/alimy/yesql"
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

//go:embed yesql.sql
var yesqlBytes []byte

var (
	_db *sqlx.DB
)

type sqlxSrv struct {
	db *sqlx.DB
}

func (s *sqlxSrv) with(handle func(tx *sqlx.Tx) error) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sqlxSrv) withTx(ctx context.Context, opts *sql.TxOptions, handle func(*sqlx.Tx) error) error {
	tx, err := s.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sqlxSrv) in(query string, args ...any) (string, []any, error) {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return "", nil, err
	}
	return s.db.Rebind(q), params, nil
}

func (s *sqlxSrv) inExec(execer sqlx.Execer, query string, args ...any) (sql.Result, error) {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	return execer.Exec(s.db.Rebind(q), params...)
}

func (s *sqlxSrv) inSelect(queryer sqlx.Queryer, dest any, query string, args ...any) error {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	return sqlx.Select(queryer, dest, s.db.Rebind(q), params...)
}

func (s *sqlxSrv) inGet(queryer sqlx.Queryer, dest any, query string, args ...any) error {
	q, params, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	return sqlx.Get(queryer, dest, s.db.Rebind(q), params...)
}

func newSqlxSrv(db *sqlx.DB) *sqlxSrv {
	return &sqlxSrv{
		db: db,
	}
}

func r(query string) string {
	return _db.Rebind(t(query))
}

func c(query string) *sqlx.Stmt {
	query = _db.Rebind(t(query))
	stmt, err := _db.Preparex(query)
	if err != nil {
		logrus.Fatalf("prepare query(%s) error: %s", query, err)
	}
	return stmt
}

func n(query string) *sqlx.NamedStmt {
	query = t(query)
	stmt, err := _db.PrepareNamed(query)
	if err != nil {
		logrus.Fatalf("prepare named query(%s) error: %s", query, err)
	}
	return stmt
}

// t repace table prefix for query
func t(query string) string {
	return strings.Replace(query, "@", conf.DatabaseSetting.TablePrefix, -1)
}

func initSqlxDB() {
	_db = conf.MustSqlxDB()
	yesql.UseSqlx(_db)
	yesql.SetDefaultQueryHooks(func(query *yesql.Query) (*yesql.Query, error) {
		qstr := strings.TrimRight(query.Query, ";")
		// table name fixed
		qstr = strings.Replace(qstr, "@", conf.DatabaseSetting.TablePrefix, -1)
		// rebind query
		if clause, exist := query.Tags["clause"]; !exist || clause != "in" {
			qstr = _db.Rebind(qstr)
		}
		query.Query = qstr
		return query, nil
	})
}
