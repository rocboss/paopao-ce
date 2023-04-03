// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

var (
	_pgxDB   *pgx.Conn
	_oncePgx sync.Once
)

func MustPgxDB() *pgx.Conn {
	_oncePgx.Do(func() {
		conn, err := pgx.Connect(context.Background(), PostgresSetting.Dsn())
		if err != nil {
			logrus.Fatalf("pgx.Connect occurs error: %s", err)
		}
		_pgxDB = conn
	})
	return _pgxDB
}
