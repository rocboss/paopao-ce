// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

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

func (s *pgcCommentManageSrv) CreateComment(r *ms.Comment) (res *ms.Comment, err error) {
	err = stmtGet(s.p.CreateComment, &res,
		r.PostID, r.UserID, r.IP,
		r.IPLoc, time.Now().Unix())
	return
}

func (s *pgcCommentManageSrv) CreateCommentReply(r *ms.CommentReply) (res *ms.CommentReply, err error) {
	err = stmtGet(s.p.CreateCommentReply, &res,
		r.CommentID, r.UserID, r.Content,
		r.AtUserID, r.IP, r.IPLoc, time.Now().Unix())
	return
}

func (s *pgcCommentManageSrv) CreateCommentContent(r *ms.CommentContent) (res *ms.CommentContent, err error) {
	err = stmtGet(s.p.CreateCommentContent, &res,
		r.CommentID, r.UserID, r.Content,
		r.Type, r.Sort, time.Now().Unix())
	return
}
