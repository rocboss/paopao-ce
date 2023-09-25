// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

type trendsSrvA struct {
	*pgxSrv
}

func (s *trendsSrvA) GetIndexTrends(userId int64, limit int, offset int) (res []*cs.TrendsItem, total int64, err error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
}

func newTrendsManageServentA(db *pgx.Conn) core.TrendsManageServantA {
	return &trendsSrvA{
		pgxSrv: newPgxSrv(db),
	}
}
