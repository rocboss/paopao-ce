// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.CommentService       = (*commentServant)(nil)
	_ core.CommentManageService = (*commentManageServant)(nil)
)

type commentServant struct {
	*pgxServant
}

type commentManageServant struct {
	*pgxServant
}

func (s *commentServant) GetComments(conditions *core.ConditionsT, offset, limit int) ([]*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentServant) GetCommentByID(id int64) (*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentServant) GetCommentReplyByID(id int64) (*core.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentServant) GetCommentCount(conditions *core.ConditionsT) (int64, error) {
	// TODO
	debug.NotImplemented()
	return 0, nil
}

func (s *commentServant) GetCommentContentsByIDs(ids []int64) ([]*core.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentServant) GetCommentRepliesByID(ids []int64) ([]*core.CommentReplyFormated, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageServant) DeleteComment(comment *core.Comment) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageServant) CreateComment(comment *core.Comment) (*core.Comment, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageServant) CreateCommentReply(reply *core.CommentReply) (*core.CommentReply, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func (s *commentManageServant) DeleteCommentReply(reply *core.CommentReply) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func (s *commentManageServant) CreateCommentContent(content *core.CommentContent) (*core.CommentContent, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

func newCommentService(db *pgx.Conn) core.CommentService {
	return &commentServant{
		pgxServant: newPgxServant(db),
	}
}

func newCommentManageService(db *pgx.Conn) core.CommentManageService {
	return &commentManageServant{
		pgxServant: newPgxServant(db),
	}
}
