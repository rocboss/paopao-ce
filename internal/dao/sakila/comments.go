// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.CommentService       = (*commentSrv)(nil)
	_ core.CommentManageService = (*commentManageSrv)(nil)
)

type commentSrv struct {
	*sqlxSrv
	stmtGetComments *sqlx.Stmt
	stmtGetReply    *sqlx.Stmt
}

type commentManageSrv struct {
	*sqlxSrv
	stmtDelComments *sqlx.Stmt
	stmtAddComents  *sqlx.Stmt
}

func (s *commentSrv) GetComments(conditions *core.ConditionsT, offset, limit int) ([]*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentByID(id int64) (*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentReplyByID(id int64) (*core.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *commentSrv) GetCommentContentsByIDs(ids []int64) ([]*core.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentSrv) GetCommentRepliesByID(ids []int64) ([]*core.CommentReplyFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) DeleteComment(comment *core.Comment) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageSrv) CreateComment(comment *core.Comment) (*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) CreateCommentReply(reply *core.CommentReply) (*core.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageSrv) DeleteCommentReply(reply *core.CommentReply) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageSrv) CreateCommentContent(content *core.CommentContent) (*core.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func newCommentService(db *sqlx.DB) core.CommentService {
	return &commentSrv{
		sqlxSrv:         newSqlxSrv(db),
		stmtGetComments: c(`SELECT * FROM @user WHERE username=?`),
		stmtGetReply:    c(`SELECT * FROM @user WHERE username=?`),
	}
}

func newCommentManageService(db *sqlx.DB) core.CommentManageService {
	return &commentManageSrv{
		sqlxSrv:         newSqlxSrv(db),
		stmtAddComents:  c(`SELECT * FROM @user WHERE username=?`),
		stmtDelComments: c(`SELECT * FROM @user WHERE username=?`),
	}
}
