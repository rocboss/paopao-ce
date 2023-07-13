// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"

	"github.com/alimy/yesql"
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

var (
	_db          *sqlx.DB
	_tablePrefix string
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
	buf := bytes.NewBuffer(make([]byte, 0, len(query)))
	qr := make([]rune, 0, len(query))
	for _, c := range query {
		qr = append(qr, c)
	}
	isPrevAt := false
	size := len(qr)
	for i := 0; i < size; i++ {
		if qr[i] == '@' {
			if next := i + 1; next == size || (!isPrevAt && qr[next] != '@') {
				buf.WriteString(_tablePrefix)
			} else {
				buf.WriteRune('@')
			}
			isPrevAt = true
		} else {
			buf.WriteRune(qr[i])
			isPrevAt = false
		}
	}
	return buf.String()
}

// yesqlScan yesql.Scan help function
func yesqlScan[T any](query yesql.SQLQuery, obj T) T {
	if err := yesql.Scan(obj, query); err != nil {
		logrus.Fatal(err)
	}
	return obj
}

func mustBuild[T any](db *sqlx.DB, fn func(yesql.PreparexBuilder, ...context.Context) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func mustBuildFn[T any](db *sqlx.DB, fn func(yesql.PreparexBuilder) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func initSqlxDB() {
	_db = conf.MustSqlxDB()
	_tablePrefix = conf.DatabaseSetting.TablePrefix
}

// FnTest_t just for test t(...) function not use in out package
func FnTest_t(query string) string {
	return t(query)
}
