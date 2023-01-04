// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	_sqlxdb   *sqlx.DB
	_onceSqlx sync.Once
)

func MustSqlxDB() *sqlx.DB {
	_onceSqlx.Do(func() {
		// TODO: init sqlx.DB
	})
	return _sqlxdb
}
