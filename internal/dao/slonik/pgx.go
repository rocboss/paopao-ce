// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/conf"
)

var (
	_pgxDB   *pgx.Conn
	_oncePgx sync.Once
)

func pgxDB() *pgx.Conn {
	_oncePgx.Do(func() {
		_pgxDB = conf.MustPgxDB()
	})
	return _pgxDB
}
