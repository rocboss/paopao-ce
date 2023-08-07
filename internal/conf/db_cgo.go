// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build cgo
// +build cgo

package conf

import (
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	sqlite3InCgoEnabled = true
)

func OpenSqlite3() (string, *sql.DB, error) {
	db, err := sql.Open("sqlite3", Sqlite3Setting.Dsn("sqlite3"))
	return "sqlite3", db, err
}

func gormOpenSqlite3(opts ...gorm.Option) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(Sqlite3Setting.Dsn("sqlite3")), opts...)
}
