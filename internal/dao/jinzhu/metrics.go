// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type tweetMetricSrvA struct {
	db *gorm.DB
}

func (s *tweetMetricSrvA) UpdateRankScore(metric *cs.TweetMetric) error {
	return s.db.Transaction(func(tx *gorm.DB) (err error) {
		postMetric := &dbr.PostMetric{}
		db := s.db.Model(postMetric).Where("post_id=?", metric.PostId)
		if err = db.First(postMetric).Error; err != nil {
			return
		}
		postMetric.RankScore = metric.RankScore(postMetric.MotivationFactor)
		err = db.Save(postMetric).Error
		return
	})

}

func (s *tweetMetricSrvA) AddTweetMetric(postId int64) (err error) {
	_, err = (&dbr.PostMetric{}).Create(s.db)
	return
}

func (s *tweetMetricSrvA) DeleteTweetMetric(postId int64) (err error) {
	return (&dbr.PostMetric{PostId: postId}).Delete(s.db)
}

func NewTweetMetricServentA(db *gorm.DB) core.TweetMetricServantA {
	return &tweetMetricSrvA{
		db: db,
	}
}
