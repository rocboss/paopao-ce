// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/bitbus/sqlx/db"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
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

type pgcCommentSrv struct {
	*commentSrv
	p *pgc.Comment
}

func (s *pgcCommentManageSrv) CreateComment(r *ms.Comment) (*ms.Comment, error) {
	return db.Get[ms.Comment](s.p.CreateComment, r.PostID, r.UserID, r.IP, r.IPLoc, time.Now().Unix())
}

func (s *pgcCommentManageSrv) CreateCommentReply(r *ms.CommentReply) (res *ms.CommentReply, err error) {
	now := time.Now().Unix()
	if res, err = db.Get[ms.CommentReply](s.p.CreateCommentReply, r.CommentID,
		r.UserID, r.Content, r.AtUserID,
		r.IP, r.IPLoc, now); err == nil {
		// 宽松处理错误
		s.q.IncrCommentReplyCount.Exec(now, r.CommentID)
	}
	return
}

func (s *pgcCommentManageSrv) CreateCommentContent(r *ms.CommentContent) (*ms.CommentContent, error) {
	return db.Get[ms.CommentContent](s.p.CreateCommentContent,
		r.CommentID, r.UserID, r.Content,
		r.Type, r.Sort, time.Now().Unix())
}

func (s *pgcCommentSrv) GetComments(tweetId int64, style cs.StyleCommentType, limit int, offset int) (res []*ms.Comment, total int64, err error) {
	switch style {
	case cs.StyleCommentHots:
		if err = s.p.GetHotsComments.Select(&res, tweetId, limit, offset); err == nil {
			err = s.q.GetCommentCount.Get(&total, tweetId)
		}
	case cs.StyleCommentNewest:
		if err = s.q.GetNewestComments.Select(&res, tweetId, limit, offset); err == nil {
			err = s.q.GetCommentCount.Get(&total, tweetId)
		}
	case cs.StyleCommentDefault:
		fallthrough
	default:
		if err = s.q.GetDefaultComments.Select(&res, tweetId, limit, offset); err == nil {
			err = s.q.GetCommentCount.Get(&total, tweetId)
		}
	}
	return
}
