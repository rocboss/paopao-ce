// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"bytes"
	"context"

	"github.com/alimy/yesql"
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/conf"
	yc "github.com/rocboss/paopao-ce/internal/dao/sakila/auto"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/ac"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
	"github.com/sirupsen/logrus"
)

var (
	_db          *sqlx.DB
	_tablePrefix string
)

type sqlxSrv struct {
	db *sqlx.DB
	y  *yc.Yesql
}

func newSqlxSrv(db *sqlx.DB) *sqlxSrv {
	return &sqlxSrv{
		db: db,
		y:  mustBuild(db, yc.BuildYesql),
	}
}

//lint:ignore U1000 r
func r(query string) string {
	return _db.Rebind(t(query))
}

//lint:ignore U1000 c
func c(query string) *sqlx.Stmt {
	query = _db.Rebind(t(query))
	stmt, err := _db.Preparex(query)
	if err != nil {
		logrus.Fatalf("prepare query(%s) error: %s", query, err)
	}
	return stmt
}

//lint:ignore U1000 n
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
//
//lint:ignore U1000 yesqlScan
func yesqlScan[T any](query yesql.SQLQuery, obj T) T {
	if err := yesql.Scan(obj, query); err != nil {
		logrus.Fatal(err)
	}
	return obj
}

func mustBuild[T any](db *sqlx.DB, fn func(yc.PreparexBuilder, ...context.Context) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func mustBuildFn[T any](db *sqlx.DB, fn func(yc.PreparexBuilder) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func acBuild[T any](db *sqlx.DB, fn func(ac.PreparexBuilder, ...context.Context) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func acBuildFn[T any](db *sqlx.DB, fn func(ac.PreparexBuilder) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func ccBuild[T any](db *sqlx.DB, fn func(cc.PreparexBuilder, ...context.Context) (T, error)) T {
	p := yesql.NewPreparexBuilder(db, t)
	obj, err := fn(p)
	if err != nil {
		logrus.Fatalf("build object failure: %s", err)
	}
	return obj
}

func ccBuildFn[T any](db *sqlx.DB, fn func(cc.PreparexBuilder) (T, error)) T {
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
