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

type commentMetricSrvA struct {
	db *gorm.DB
}

type userMetricSrvA struct {
	db *gorm.DB
}

func (s *tweetMetricSrvA) UpdateTweetMetric(metric *cs.TweetMetric) error {
	return s.db.Transaction(func(tx *gorm.DB) (err error) {
		postMetric := &dbr.PostMetric{PostId: metric.PostId}
		db := s.db.Model(postMetric).Where("post_id=?", metric.PostId)
		db.First(postMetric)
		postMetric.RankScore = metric.RankScore(postMetric.MotivationFactor)
		err = db.Save(postMetric).Error
		return
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
	return s.db.Transaction(func(tx *gorm.DB) (err error) {
		commentMetric := &dbr.CommentMetric{CommentId: metric.CommentId}
		db := s.db.Model(commentMetric).Where("comment_id=?", metric.CommentId)
		db.First(commentMetric)
		commentMetric.RankScore = metric.RankScore(commentMetric.MotivationFactor)
		err = db.Save(commentMetric).Error
		return
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
	// TODO
	return cs.ErrNotImplemented
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
