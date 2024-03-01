// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"sync"
	"time"

	"github.com/alimy/tryst/cfg"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	_gormdb   *gorm.DB
	_onceGorm sync.Once
)

func MustGormDB() *gorm.DB {
	_onceGorm.Do(func() {
		var err error
		if _gormdb, err = newGormDB(); err != nil {
			logrus.Fatalf("new gorm db failed: %s", err)
		}
	})
	return _gormdb
}

func newGormDB() (db *gorm.DB, err error) {
	newLogger := logger.New(
		logrus.StandardLogger(), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,                // 慢 SQL 阈值
			LogLevel:                  DatabaseSetting.logLevel(), // 日志级别
			IgnoreRecordNotFoundError: true,                       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                      // 禁用彩色打印
		},
	)

	config := &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
	}

	plugin := dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(MysqlSetting.MaxIdleConns).
		SetMaxOpenConns(MysqlSetting.MaxOpenConns)

	if cfg.If("MySQL") {
		logrus.Debugln("use MySQL as db")
		db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config)
	} else if cfg.If("Postgres") {
		logrus.Debugln("use PostgreSQL as db")
		db, err = gorm.Open(postgres.Open(PostgresSetting.Dsn()), config)
	} else if cfg.If("Sqlite3") {
		logrus.Debugf("use Sqlite3 as db path:%s sqlite3InCgoEnabled:%t", Sqlite3Setting.Path, sqlite3InCgoEnabled)
		db, err = gormOpenSqlite3(config)
	} else {
		logrus.Debugln("use default of MySQL as db")
		db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config)
	}
	if err == nil {
		err = db.Use(plugin)
	}
	return
}
