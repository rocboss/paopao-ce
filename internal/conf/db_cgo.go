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

func OpenSqlite3() (*sql.DB, error) {
	return sql.Open("sqlite3", Sqlite3Setting.Dsn("sqlite3"))
}

func gormOpenSqlite3(opts ...gorm.Option) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(Sqlite3Setting.Dsn("sqlite3")), opts...)
}
