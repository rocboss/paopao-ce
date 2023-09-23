// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.CommentService       = (*commentSrv)(nil)
	_ core.CommentManageService = (*commentManageSrv)(nil)
)

type commentSrv struct {
	*pgxSrv
}

type commentManageSrv struct {
	*pgxSrv
}

func (s *commentSrv) GetComments(tweetId int64, style cs.StyleCommentType, limit int, offset int) ([]*ms.Comment, int64, error) {
	// TODO
	return nil, 0, cs.ErrNotImplemented
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

func (s *commentSrv) GetCommentContentsByIDs(ids []int64) ([]*ms.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentRepliesByID(ids []int64) ([]*ms.CommentReplyFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error) {
	// TODO
	return nil, nil, debug.ErrNotImplemented
}

func (s *commentManageSrv) DeleteComment(comment *ms.Comment) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageSrv) HighlightComment(userId, commentId int64) (res int8, err error) {
	// TODO
	return 0, cs.ErrNotImplemented
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

func newCommentService(db *pgx.Conn) core.CommentService {
	return &commentSrv{
		pgxSrv: newPgxSrv(db),
	}
}

func newCommentManageService(db *pgx.Conn) core.CommentManageService {
	return &commentManageSrv{
		pgxSrv: newPgxSrv(db),
	}
}
