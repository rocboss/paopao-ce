// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type tweetMetricSrvA struct {
	db *gorm.DB
}

type commentMetricSrvA struct {
	db *gorm.DB
}

type userMetricSrvA struct {
	db *gorm.DB
}

func (s *tweetMetricSrvA) UpdateTweetMetric(metric *cs.TweetMetric) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		postMetric := &dbr.PostMetric{PostId: metric.PostId}
		tx.Model(postMetric).Where("post_id=?", metric.PostId).First(postMetric)
		postMetric.RankScore = metric.RankScore(postMetric.MotivationFactor)
		return tx.Save(postMetric).Error
	})
}

func (s *tweetMetricSrvA) AddTweetMetric(postId int64) (err error) {
	_, err = (&dbr.PostMetric{PostId: postId}).Create(s.db)
	return
}

func (s *tweetMetricSrvA) DeleteTweetMetric(postId int64) (err error) {
	return (&dbr.PostMetric{PostId: postId}).Delete(s.db)
}

func (s *commentMetricSrvA) UpdateCommentMetric(metric *cs.CommentMetric) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		commentMetric := &dbr.CommentMetric{CommentId: metric.CommentId}
		tx.Model(commentMetric).Where("comment_id=?", metric.CommentId).First(commentMetric)
		commentMetric.RankScore = metric.RankScore(commentMetric.MotivationFactor)
		return tx.Save(commentMetric).Error
	})
}

func (s *commentMetricSrvA) AddCommentMetric(commentId int64) (err error) {
	_, err = (&dbr.CommentMetric{CommentId: commentId}).Create(s.db)
	return
}

func (s *commentMetricSrvA) DeleteCommentMetric(commentId int64) (err error) {
	return (&dbr.CommentMetric{CommentId: commentId}).Delete(s.db)
}

func (s *userMetricSrvA) UpdateUserMetric(userId int64, action uint8) error {
	metric := &dbr.UserMetric{UserId: userId}
	s.db.Model(metric).Where("user_id=?", userId).First(metric)
	metric.LatestTrendsOn = time.Now().Unix()
	switch action {
	case cs.MetricActionCreateTweet:
		metric.TweetsCount++
	case cs.MetricActionDeleteTweet:
		if metric.TweetsCount > 0 {
			metric.TweetsCount--
		}
	}
	return s.db.Save(metric).Error
}

func (s *userMetricSrvA) AddUserMetric(userId int64) (err error) {
	_, err = (&dbr.UserMetric{UserId: userId}).Create(s.db)
	return
}

func (s *userMetricSrvA) DeleteUserMetric(userId int64) (err error) {
	return (&dbr.UserMetric{UserId: userId}).Delete(s.db)
}

func newTweetMetricServentA(db *gorm.DB) core.TweetMetricServantA {
	return &tweetMetricSrvA{
		db: db,
	}
}

func newCommentMetricServentA(db *gorm.DB) core.CommentMetricServantA {
	return &commentMetricSrvA{
		db: db,
	}
}

func newUserMetricServentA(db *gorm.DB) core.UserMetricServantA {
	return &userMetricSrvA{
		db: db,
	}
}
