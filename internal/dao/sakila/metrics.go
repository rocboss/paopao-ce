// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
)

type tweetMetricSrvA struct {
	*sqlxSrv
	q *cc.TweetMetrics
}

type commentMetricSrvA struct {
	*sqlxSrv
	q *cc.CommentMetrics
}

type userMetricSrvA struct {
	*sqlxSrv
	q *cc.UserMetrics
}

func (s *tweetMetricSrvA) UpdateTweetMetric(metric *cs.TweetMetric) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		var motivationFactor int
		now := time.Now().Unix()
		err := tx.Stmtx(s.q.GetMotivationFactor).Get(&motivationFactor, metric.PostId)
		if err == nil {
			_, err = tx.Stmtx(s.q.UpdateRankScore).Exec(metric.RankScore(motivationFactor), now, metric.PostId)
		} else {
			_, err = tx.Stmtx(s.q.UpsertTweetMetric).Exec(metric.PostId, metric.RankScore(motivationFactor), now)
		}
		return err
	})
}

func (s *tweetMetricSrvA) AddTweetMetric(postId int64) error {
	_, err := s.q.AddTweetMetric.Exec(postId, time.Now().Unix())
	return err
}

func (s *tweetMetricSrvA) DeleteTweetMetric(postId int64) error {
	_, err := s.q.DeleteTweetMetric.Exec(time.Now().Unix(), postId)
	return err
}

func (s *commentMetricSrvA) UpdateCommentMetric(metric *cs.CommentMetric) error {
	return s.db.Withx(func(tx *sqlx.Tx) error {
		var motivationFactor int
		now := time.Now().Unix()
		err := tx.Stmtx(s.q.GetMotivationFactor).Get(&motivationFactor, metric.CommentId)
		if err == nil {
			_, err = tx.Stmtx(s.q.UpdateRankScore).Exec(metric.RankScore(motivationFactor), now, metric.CommentId)
		} else {
			_, err = tx.Stmtx(s.q.UpsertCommentMetric).Exec(metric.CommentId, metric.RankScore(motivationFactor), now)
		}
		return err
	})
}

func (s *commentMetricSrvA) AddCommentMetric(commentId int64) error {
	_, err := s.q.AddCommentMetric.Exec(commentId, time.Now().Unix())
	return err
}

func (s *commentMetricSrvA) DeleteCommentMetric(commentId int64) error {
	_, err := s.q.DeleteCommentMetric.Exec(time.Now().Unix(), commentId)
	return err
}

func (s *userMetricSrvA) UpdateUserMetric(userId int64, action uint8) (err error) {
	var tweets_count int
	now := time.Now().Unix()
	if err = s.q.GetTweetsCount.Get(&tweets_count, userId); err != nil {
		_, err = s.q.AddUserMetric.Exec(userId, now)
		return
	}
	switch action {
	case cs.MetricActionCreateTweet:
		tweets_count++
	case cs.MetricActionDeleteTweet:
		if tweets_count > 0 {
			tweets_count--
		}
	}
	_, err = s.q.UpdateUserMetric.Exec(userId, tweets_count)
	return
}

func (s *userMetricSrvA) AddUserMetric(userId int64) error {
	_, err := s.q.AddUserMetric.Exec(userId, time.Now().Unix())
	return err
}

func (s *userMetricSrvA) DeleteUserMetric(userId int64) error {
	_, err := s.q.DeleteUserMetric.Exec(time.Now().Unix(), userId)
	return err
}

func newTweetMetricServentA(db *sqlx.DB) core.TweetMetricServantA {
	return &tweetMetricSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildTweetMetrics),
	}
}

func newCommentMetricServentA(db *sqlx.DB) core.CommentMetricServantA {
	return &commentMetricSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildCommentMetrics),
	}
}

func newUserMetricServentA(db *sqlx.DB) core.UserMetricServantA {
	return &userMetricSrvA{
		sqlxSrv: newSqlxSrv(db),
		q:       ccBuild(db, cc.BuildUserMetrics),
	}
}
