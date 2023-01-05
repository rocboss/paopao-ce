// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"database/sql"
	"sync"

	"github.com/alimy/cfg"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	_sqldb               *sql.DB
	_redisClient         *redis.Client
	_onceSql, _onceRedis sync.Once
)

func MustSqlDB() *sql.DB {
	_onceSql.Do(func() {
		var err error
		if _, _sqldb, err = newSqlDB(); err != nil {
			logrus.Fatalf("new sql db failed: %s", err)
		}
	})
	return _sqldb
}

func MustRedis() *redis.Client {
	_onceRedis.Do(func() {
		_redisClient = redis.NewClient(&redis.Options{
			Addr:     redisSetting.Host,
			Password: redisSetting.Password,
			DB:       redisSetting.DB,
		})
	})
	return _redisClient
}

func newSqlDB() (driver string, db *sql.DB, err error) {
	if cfg.If("MySQL") {
		driver = "mysql"
		db, err = sql.Open(driver, MysqlSetting.Dsn())
	} else if cfg.If("PostgreSQL") || cfg.If("Postgres") {
		driver = "pgx"
		db, err = sql.Open(driver, PostgresSetting.Dsn())
	} else if cfg.If("Sqlite3") {
		driver, db, err = OpenSqlite3()
	} else {
		driver = "mysql"
		db, err = sql.Open(driver, MysqlSetting.Dsn())
	}
	return
}
