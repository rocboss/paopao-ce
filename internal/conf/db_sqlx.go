// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"sync"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/pkg/naming"
	"github.com/sirupsen/logrus"
)

var (
	_sqlxdb   *sqlx.DB
	_onceSqlx sync.Once
)

func MustSqlxDB() *sqlx.DB {
	_onceSqlx.Do(func() {
		driver, db, err := newSqlDB()
		if err != nil {
			logrus.Fatalf("new sqlx db failed: %s", err)
		}
		_sqlxdb = sqlx.NewDb(db, driver)
		_sqlxdb.MapperFunc(naming.NewSnakeNamingStrategy().Naming)
	})
	return _sqlxdb
}
