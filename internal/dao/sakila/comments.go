// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.CommentService       = (*commentSrv)(nil)
	_ core.CommentManageService = (*commentManageSrv)(nil)
)

type commentSrv struct {
	*sqlxSrv
	q *cc.Comment
}

type commentManageSrv struct {
	*sqlxSrv
	q *cc.CommentManage
}

func (s *commentSrv) GetComments(conditions *ms.ConditionsT, offset, limit int) ([]*ms.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentByID(id int64) (*ms.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentReplyByID(id int64) (*ms.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentCount(conditions *ms.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *commentSrv) GetCommentContentsByIDs(ids []int64) ([]*ms.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error) {
	// TODO
	return nil, nil, debug.ErrNotImplemented
}

func (s *commentSrv) GetCommentRepliesByID(ids []int64) ([]*ms.CommentReplyFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) DeleteComment(comment *ms.Comment) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageSrv) CreateComment(comment *ms.Comment) (*ms.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) CreateCommentReply(reply *ms.CommentReply) (*ms.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) DeleteCommentReply(reply *ms.CommentReply) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageSrv) CreateCommentContent(content *ms.CommentContent) (*ms.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) ThumbsUpComment(userId int64, tweetId, commentId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *commentManageSrv) ThumbsDownComment(userId int64, tweetId, commentId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *commentManageSrv) ThumbsUpReply(userId int64, tweetId, commentId, replyId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func (s *commentManageSrv) ThumbsDownReply(userId int64, tweetId, commentId, replyId int64) error {
	// TODO
	return debug.ErrNotImplemented
}

func newCommentService(db *sqlx.DB) core.CommentService {
	return &commentSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildComment),
	}
}

func newCommentManageService(db *sqlx.DB) core.CommentManageService {
	return &commentManageSrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildCommentManage),
	}
}
