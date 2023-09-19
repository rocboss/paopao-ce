// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx/db"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.CommentManageService = (*pgcCommentManageSrv)(nil)
)

type pgcCommentManageSrv struct {
	*commentManageSrv
	p *pgc.CommentManage
}

func (s *pgcCommentManageSrv) CreateComment(r *ms.Comment) (*ms.Comment, error) {
	return db.Get[ms.Comment](s.p.CreateComment, r.PostID, r.UserID, r.IP, r.IPLoc, time.Now().Unix())
}

func (s *pgcCommentManageSrv) CreateCommentReply(r *ms.CommentReply) (*ms.CommentReply, error) {
	return db.Get[ms.CommentReply](s.p.CreateCommentReply, r.CommentID,
		r.UserID, r.Content, r.AtUserID,
		r.IP, r.IPLoc, time.Now().Unix())
}

func (s *pgcCommentManageSrv) CreateCommentContent(r *ms.CommentContent) (*ms.CommentContent, error) {
	return db.Get[ms.CommentContent](s.p.CreateCommentContent,
		r.CommentID, r.UserID, r.Content,
		r.Type, r.Sort, time.Now().Unix())

}
