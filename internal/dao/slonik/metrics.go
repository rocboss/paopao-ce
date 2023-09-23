// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
)

type tweetMetricSrvA struct {
	*pgxSrv
}

type commentMetricSrvA struct {
	*pgxSrv
}

type userMetricSrvA struct {
	*pgxSrv
}

func (s *tweetMetricSrvA) UpdateTweetMetric(metric *cs.TweetMetric) error {
	// TODO
	return cs.ErrNotImplemented
}

func (s *tweetMetricSrvA) AddTweetMetric(postId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func (s *tweetMetricSrvA) DeleteTweetMetric(postId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func (s *commentMetricSrvA) UpdateCommentMetric(metric *cs.CommentMetric) error {
	// TODO
	return cs.ErrNotImplemented
}

func (s *commentMetricSrvA) AddCommentMetric(commentId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func (s *commentMetricSrvA) DeleteCommentMetric(commentId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func (s *userMetricSrvA) UpdateUserMetric(userId int64, action uint8) error {
	// TODO
	return cs.ErrNotImplemented
}

func (s *userMetricSrvA) AddUserMetric(userId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func (s *userMetricSrvA) DeleteUserMetric(userId int64) (err error) {
	// TODO
	return cs.ErrNotImplemented
}

func newTweetMetricServentA(db *pgx.Conn) core.TweetMetricServantA {
	return &tweetMetricSrvA{
		pgxSrv: newPgxSrv(db),
	}
}

func newCommentMetricServentA(db *pgx.Conn) core.CommentMetricServantA {
	return &commentMetricSrvA{
		pgxSrv: newPgxSrv(db),
	}
}

func newUserMetricServentA(db *pgx.Conn) core.UserMetricServantA {
	return &userMetricSrvA{
		pgxSrv: newPgxSrv(db),
	}
}
