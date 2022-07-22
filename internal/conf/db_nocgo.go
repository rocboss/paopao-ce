//go:build !cgo
// +build !cgo

package conf

import (
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

const (
	sqlite3InCgoEnabled = false
)

func OpenSqlite3() (*sql.DB, error) {
	return sql.Open("sqlite", Sqlite3Setting.Dsn("sqlite"))
}

func gormOpenSqlite3(opts ...gorm.Option) (*gorm.DB, error) {
	dialector := &sqlite.Dialector{DriverName: "sqlite", DSN: Sqlite3Setting.Dsn("sqlite")}
	return gorm.Open(dialector, opts...)
}
