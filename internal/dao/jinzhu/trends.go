// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"fmt"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"gorm.io/gorm"
)

type trendsSrvA struct {
	db *gorm.DB
}

func (s *trendsSrvA) GetIndexTrends(userId int64, limit int, offset int) (res []*cs.TrendsItem, total int64, err error) {
	db := s.db.Table(_user_).
		Joins(fmt.Sprintf("JOIN %s r ON r.he_uid=%s.id", _userRelation_, _user_)).
		Joins(fmt.Sprintf("JOIN %s m ON r.he_uid=m.user_id", _userMetric_)).
		Where("r.user_id=? AND m.tweets_count>0 AND m.is_del=0", userId)
	if err = db.Count(&total).Error; err != nil || total == 0 {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if err = db.Order("r.style ASC").Select("username", "nickname", "avatar").Find(&res).Error; err == nil {
		res = cs.DistinctTrends(res)
	}
	return
}

func newTrendsManageServentA(db *gorm.DB) core.TrendsManageServantA {
	return &trendsSrvA{
		db: db,
	}
}
