// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
)

type trendsSrvA struct {
	*sqlxSrv
	q *cc.TrendsManager
}

func (s *trendsSrvA) GetIndexTrends(userId int64, limit int, offset int) (res []*cs.TrendsItem, total int64, err error) {
	if err = s.q.CountIndexTrends.Get(&total, userId); err == nil && total > 0 {
		err = s.q.GetIndexTrends.Select(&res, userId, limit, offset)
	}
	return
}

func newTrendsManageServentA(db *sqlx.DB) core.TrendsManageServantA {
	return &trendsSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildTrendsManager),
	}
}
