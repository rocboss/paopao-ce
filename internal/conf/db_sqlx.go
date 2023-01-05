// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"database/sql"
	"sync"

	"github.com/alimy/cfg"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	_sqlxdb   *sqlx.DB
	_onceSqlx sync.Once
)

func MustSqlxDB() *sqlx.DB {
	_onceSqlx.Do(func() {
		var err error
		if _sqlxdb, err = newSqlxDB(); err != nil {
			logrus.Fatalf("new sqlx db failed: %s", err)
		}
	})
	return _sqlxdb
}

func newSqlxDB() (db *sqlx.DB, err error) {
	if cfg.If("MySQL") {
		db, err = sqlx.Open("mysql", MysqlSetting.Dsn())
	} else if cfg.If("PostgreSQL") || cfg.If("Postgres") {
		db, err = sqlx.Open("postgres", PostgresSetting.Dsn())
	} else if cfg.If("Sqlite3") {
		var (
			driver string
			sqldb  *sql.DB
		)
		driver, sqldb, err = OpenSqlite3()
		db = sqlx.NewDb(sqldb, driver)
	} else {
		db, err = sqlx.Open("mysql", MysqlSetting.Dsn())
	}
	return
}
