// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/slonik/sqlc/auto/pgc"
)

var (
	_ core.CommentService       = (*commentSrv)(nil)
	_ core.CommentManageService = (*commentManageSrv)(nil)
)

type commentSrv struct {
	*pgcSrv
}

type commentManageSrv struct {
	*pgcSrv
}

func (s *commentSrv) GetComments(tweetId int64, style cs.StyleCommentType, limit int, offset int) (res []*ms.Comment, total int64, err error) {
	c := context.Background()
	if total, err = s.q.GetCommentCount(c, tweetId); err != nil && total == 0 {
		return
	}
	var (
		comments []*pgc.PComment
	)
	switch style {
	case cs.StyleCommentHots:
		comments, err = s.q.GetHotsComments(c, &pgc.GetHotsCommentsParams{
			PostID: tweetId,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case cs.StyleCommentNewest:
		comments, err = s.q.GetNewestComments(c, &pgc.GetNewestCommentsParams{
			PostID: tweetId,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	case cs.StyleCommentDefault:
		fallthrough
	default:
		comments, err = s.q.GetDefaultComments(c, &pgc.GetDefaultCommentsParams{
			PostID: tweetId,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	}
	for _, comment := range comments {
		res = append(res, &ms.Comment{
			Model: &ms.Model{
				ID:         comment.ID,
				CreatedOn:  comment.CreatedOn,
				ModifiedOn: comment.ModifiedOn,
			},
			PostID:          comment.PostID,
			UserID:          comment.UserID,
			IPLoc:           comment.IpLoc,
			IsEssence:       int8(comment.IsEssence),
			ReplyCount:      comment.ReplyCount,
			ThumbsUpCount:   comment.ThumbsUpCount,
			ThumbsDownCount: comment.ThumbsDownCount,
		})
	}
	return
}

func (s *commentSrv) GetCommentByID(id int64) (*ms.Comment, error) {
	comment, err := s.q.GetCommentById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &ms.Comment{
		Model: &ms.Model{
			ID:        comment.ID,
			CreatedOn: comment.CreatedOn,
		},
		PostID:          comment.PostID,
		UserID:          comment.UserID,
		IPLoc:           comment.IpLoc,
		IsEssence:       int8(comment.IsEssence),
		ReplyCount:      comment.ReplyCount,
		ThumbsUpCount:   comment.ThumbsUpCount,
		ThumbsDownCount: comment.ThumbsDownCount,
	}, nil
}

func (s *commentSrv) GetCommentReplyByID(id int64) (*ms.CommentReply, error) {
	reply, err := s.q.GetCommentReplyById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &ms.CommentReply{
		Model: &ms.Model{
			ID:        reply.ID,
			CreatedOn: reply.CreatedOn,
		},
		CommentID:       reply.CommentID,
		UserID:          reply.UserID,
		AtUserID:        reply.AtUserID,
		Content:         reply.Content,
		IP:              reply.Ip,
		IPLoc:           reply.IpLoc,
		ThumbsUpCount:   reply.ThumbsUpCount,
		ThumbsDownCount: reply.ThumbsDownCount,
	}, nil
}

func (s *commentSrv) GetCommentContentsByIDs(ids []int64) (res []*ms.CommentContent, _ error) {
	if len(ids) == 0 {
		return []*ms.CommentContent{}, nil
	}
	contents, err := s.q.GetCommentContentsByIds(context.Background(), ids)
	if err != nil {
		return nil, err
	}
	for _, content := range contents {
		res = append(res, &ms.CommentContent{
			Model: &ms.Model{
				ID:        content.ID,
				CreatedOn: content.CreatedOn,
			},
			CommentID: content.CommentID,
			UserID:    content.UserID,
			Content:   content.Content,
			Type:      ms.PostContentT(content.Type),
			Sort:      content.Sort,
		})
	}
	return nil, nil
}

func (s *commentSrv) GetCommentRepliesByID(ids []int64) ([]*ms.CommentReplyFormated, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	ctx := context.Background()
	replies, err := s.q.GetCommentRepliesByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	userIds := []int64{}
	for _, reply := range replies {
		userIds = append(userIds, reply.UserID, reply.AtUserID)
	}
	var users []*pgc.PUser
	if len(userIds) > 0 {
		users, err = s.q.GetUsersByIds(ctx, userIds)
		if err != nil {
			return nil, err
		}
	}
	repliesFormated := []*ms.CommentReplyFormated{}
	for _, reply := range replies {
		replyFormated := &ms.CommentReplyFormated{
			ID:            reply.ID,
			CommentID:     reply.CommentID,
			UserID:        reply.UserID,
			AtUserID:      reply.AtUserID,
			Content:       reply.Content,
			IPLoc:         reply.IpLoc,
			ThumbsUpCount: reply.ThumbsUpCount,
			CreatedOn:     reply.CreatedOn,
			ModifiedOn:    reply.ModifiedOn,
		}
		for _, user := range users {
			formatedUser := &ms.UserFormated{
				ID:       user.ID,
				Nickname: user.Nickname,
				Status:   int(user.Status),
				Avatar:   user.Avatar,
				IsAdmin:  user.IsAdmin,
			}
			if reply.UserID == user.ID {
				replyFormated.User = formatedUser
			}
			if reply.AtUserID == user.ID {
				replyFormated.AtUser = formatedUser
			}
		}
		repliesFormated = append(repliesFormated, replyFormated)
	}
	return repliesFormated, nil
}

func (s *commentSrv) GetCommentThumbsMap(userId int64, tweetId int64) (cs.CommentThumbsMap, cs.CommentThumbsMap, error) {
	if userId < 0 {
		return nil, nil, nil
	}
	ctx := context.Background()
	commentThumbsList, err := s.q.GetCommentThumbs(ctx, &pgc.GetCommentThumbsParams{
		UserID:  userId,
		TweetID: tweetId,
	})
	if err != nil {
		return nil, nil, err
	}
	commentThumbs, replyThumbs := make(cs.CommentThumbsMap), make(cs.CommentThumbsMap)
	for _, thumb := range commentThumbsList {
		commentThumb := &cs.CommentThumbs{
			UserID:       thumb.UserID,
			TweetID:      thumb.TweetID,
			CommentID:    thumb.CommentID,
			ReplyID:      thumb.ReplyID.Int64,
			CommentType:  int8(thumb.CommentType),
			IsThumbsUp:   int8(thumb.IsThumbsUp),
			IsThumbsDown: int8(thumb.IsThumbsDown),
		}
		if thumb.CommentType == 0 {
			commentThumbs[thumb.CommentID] = commentThumb
		} else {
			replyThumbs[thumb.ReplyID.Int64] = commentThumb
		}
	}
	return commentThumbs, replyThumbs, nil
}

func (s *commentManageSrv) DeleteComment(comment *ms.Comment) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		now := time.Now().Unix()
		err := q.DeleteComment(c, &pgc.DeleteCommentParams{
			DeletedOn: now,
			ID:        comment.ID,
		})
		if err != nil {
			return err
		}
		return q.DeleteCommentThumbs(c, &pgc.DeleteCommentThumbsParams{
			DeletedOn: now,
			UserID:    comment.UserID,
			TweetID:   comment.PostID,
			CommentID: comment.ID,
		})
	})
}

func (s *commentManageSrv) HighlightComment(userId, commentId int64) (int8, error) {
	res, err := s.q.HighlightComment(context.Background(), &pgc.HighlightCommentParams{
		UserID:     userId,
		ID:         commentId,
		ModifiedOn: time.Now().Unix(),
	})
	return int8(res), err
}

func (s *commentManageSrv) CreateComment(comment *ms.Comment) (res *ms.Comment, err error) {
	res, comment.Model = comment, &ms.Model{CreatedOn: time.Now().Unix()}
	res.ID, err = s.q.CreateComment(context.Background(), &pgc.CreateCommentParams{
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Ip:        comment.IP,
		IpLoc:     comment.IPLoc,
		CreatedOn: res.CreatedOn,
	})
	return
}

func (s *commentManageSrv) CreateCommentReply(reply *ms.CommentReply) (res *ms.CommentReply, err error) {
	res, reply.Model = reply, &ms.Model{CreatedOn: time.Now().Unix()}
	res.ID, err = s.q.CreateCommentReply(context.Background(), &pgc.CreateCommentReplyParams{
		CommentID: reply.CommentID,
		UserID:    reply.UserID,
		Content:   reply.Content,
		AtUserID:  reply.AtUserID,
		Ip:        reply.IP,
		IpLoc:     reply.IPLoc,
		CreatedOn: reply.CreatedOn,
	})
	return
}

func (s *commentManageSrv) DeleteCommentReply(reply *ms.CommentReply) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		now := time.Now().Unix()
		err := q.DeleteCommentReply(c, &pgc.DeleteCommentReplyParams{
			DeletedOn: now,
			ID:        reply.ID,
		})
		if err != nil {
			return err
		}
		return q.DeleteReplyThumbs(c, &pgc.DeleteReplyThumbsParams{
			DeletedOn: now,
			UserID:    reply.UserID,
			CommentID: reply.CommentID,
			ReplyID:   pgtype.Int8{Int64: reply.ID, Valid: true},
		})
	})
}

func (s *commentManageSrv) CreateCommentContent(content *ms.CommentContent) (res *ms.CommentContent, err error) {
	res, content.Model = content, &ms.Model{CreatedOn: time.Now().Unix()}
	res.ID, err = s.q.CreateCommentContent(context.Background(), &pgc.CreateCommentContentParams{
		CommentID: content.ID,
		UserID:    content.UserID,
		Content:   content.Content,
		Type:      int16(content.Type),
		Sort:      content.Sort,
		CreatedOn: res.CreatedOn,
	})
	return
}

func (s *commentManageSrv) ThumbsUpComment(userId int64, tweetId, commentId int64) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		var (
			thumbsUpCount   int32 = 0
			thumbsDownCount int32 = 0
		)
		now := time.Now().Unix()
		// 检查thumbs状态
		commentThumbs, err := q.GetTweetCommentThumb(c, &pgc.GetTweetCommentThumbParams{
			UserID:    userId,
			TweetID:   tweetId,
			CommentID: commentId,
		})
		yes, no := int16(1), int16(0)
		if err == nil {
			switch {
			case commentThumbs.IsThumbsUp == yes && commentThumbs.IsThumbsDown == no:
				thumbsUpCount, thumbsDownCount = -1, 0
			case commentThumbs.IsThumbsUp == no && commentThumbs.IsThumbsDown == no:
				thumbsUpCount, thumbsDownCount = 1, 0
			default:
				thumbsUpCount, thumbsDownCount = 1, -1
				commentThumbs.IsThumbsDown = no
			}
			commentThumbs.IsThumbsUp = 1 - commentThumbs.IsThumbsUp
			if err = q.UpdateThumbsUpDownComment(c, &pgc.UpdateThumbsUpDownCommentParams{
				IsThumbsUp:   commentThumbs.IsThumbsUp,
				IsThumbsDown: commentThumbs.IsThumbsDown,
				ModifiedOn:   now,
				ID:           commentThumbs.ID,
			}); err != nil {
				return err
			}
		} else {
			thumbsUpCount, thumbsDownCount = 1, 0
			if _, err = q.CreateThumbsUpDownComment(c, &pgc.CreateThumbsUpDownCommentParams{
				UserID:       userId,
				TweetID:      tweetId,
				CommentID:    commentId,
				ReplyID:      pgtype.Int8{Int64: 0, Valid: true},
				IsThumbsUp:   yes,
				IsThumbsDown: no,
				CommentType:  0,
				CreatedOn:    now,
			}); err != nil {
				return err
			}
		}
		// 更新comment thumbsUpCount
		return q.UpdateCommentThumbsCount(c, &pgc.UpdateCommentThumbsCountParams{
			ThumbsUpCount:   thumbsUpCount,
			ThumbsDownCount: thumbsDownCount,
			ModifiedOn:      now,
			ID:              commentId,
		})
	})
}

func (s *commentManageSrv) ThumbsDownComment(userId int64, tweetId, commentId int64) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		var (
			thumbsUpCount   int32 = 0
			thumbsDownCount int32 = 0
		)
		now := time.Now().Unix()
		// 检查thumbs状态
		commentThumbs, err := q.GetTweetCommentThumb(c, &pgc.GetTweetCommentThumbParams{
			UserID:    userId,
			TweetID:   tweetId,
			CommentID: commentId,
		})
		yes, no := int16(1), int16(0)
		if err == nil {
			switch {
			case commentThumbs.IsThumbsDown == yes:
				thumbsUpCount, thumbsDownCount = 0, -1
			case commentThumbs.IsThumbsDown == no && commentThumbs.IsThumbsUp == yes:
				thumbsUpCount, thumbsDownCount = 0, 1
			default:
				thumbsUpCount, thumbsDownCount = -1, 1
				commentThumbs.IsThumbsUp = no
			}
			commentThumbs.IsThumbsDown = 1 - commentThumbs.IsThumbsDown
			if err = q.UpdateThumbsUpDownComment(c, &pgc.UpdateThumbsUpDownCommentParams{
				IsThumbsUp:   commentThumbs.IsThumbsUp,
				IsThumbsDown: commentThumbs.IsThumbsDown,
				ModifiedOn:   now,
				ID:           commentThumbs.ID,
			}); err != nil {
				return err
			}
		} else {
			thumbsUpCount, thumbsDownCount = 0, 1
			if _, err = q.CreateThumbsUpDownComment(c, &pgc.CreateThumbsUpDownCommentParams{
				UserID:       userId,
				TweetID:      tweetId,
				CommentID:    commentId,
				ReplyID:      pgtype.Int8{Int64: 0, Valid: true},
				IsThumbsUp:   no,
				IsThumbsDown: yes,
				CommentType:  0,
				CreatedOn:    now,
			}); err != nil {
				return err
			}
		}
		// 更新comment thumbsUpCount
		return q.UpdateCommentThumbsCount(c, &pgc.UpdateCommentThumbsCountParams{
			ThumbsUpCount:   thumbsUpCount,
			ThumbsDownCount: thumbsDownCount,
			ModifiedOn:      now,
			ID:              commentId,
		})
	})
}

func (s *commentManageSrv) ThumbsUpReply(userId int64, tweetId, commentId, replyId int64) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		var (
			thumbsUpCount   int32 = 0
			thumbsDownCount int32 = 0
		)
		now := time.Now().Unix()
		// 检查thumbs状态
		commentThumbs, err := q.GetCommentReplyThumb(c, &pgc.GetCommentReplyThumbParams{
			UserID:    userId,
			TweetID:   tweetId,
			CommentID: commentId,
			ReplyID:   pgtype.Int8{Int64: replyId, Valid: true},
		})
		yes, no := int16(1), int16(0)
		if err == nil {
			switch {
			case commentThumbs.IsThumbsUp == yes:
				thumbsUpCount, thumbsDownCount = -1, 0
			case commentThumbs.IsThumbsUp == no && commentThumbs.IsThumbsDown == no:
				thumbsUpCount, thumbsDownCount = 1, 0
			default:
				thumbsUpCount, thumbsDownCount = 1, -1
				commentThumbs.IsThumbsDown = no
			}
			commentThumbs.IsThumbsUp = 1 - commentThumbs.IsThumbsUp
			if err = q.UpdateThumbsUpDownComment(c, &pgc.UpdateThumbsUpDownCommentParams{
				IsThumbsUp:   commentThumbs.IsThumbsUp,
				IsThumbsDown: commentThumbs.IsThumbsDown,
				ModifiedOn:   now,
				ID:           commentThumbs.ID,
			}); err != nil {
				return err
			}
		} else {
			thumbsUpCount, thumbsDownCount = 1, 0
			if _, err = q.CreateThumbsUpDownComment(c, &pgc.CreateThumbsUpDownCommentParams{
				UserID:       userId,
				TweetID:      tweetId,
				CommentID:    commentId,
				ReplyID:      pgtype.Int8{Int64: replyId, Valid: true},
				IsThumbsUp:   yes,
				IsThumbsDown: no,
				CommentType:  1,
				CreatedOn:    now,
			}); err != nil {
				return err
			}
		}
		// 更新comment_reply thumbsUpCount
		return q.UpdateReplyThumbsCount(c, &pgc.UpdateReplyThumbsCountParams{
			ThumbsUpCount:   thumbsUpCount,
			ThumbsDownCount: thumbsDownCount,
			ModifiedOn:      now,
			ID:              replyId,
		})
	})
}

func (s *commentManageSrv) ThumbsDownReply(userId int64, tweetId, commentId, replyId int64) error {
	return with(s.db, s.q.WithTx, func(c context.Context, q *pgc.Queries) error {
		var (
			thumbsUpCount   int32 = 0
			thumbsDownCount int32 = 0
		)
		now := time.Now().Unix()
		// 检查thumbs状态
		commentThumbs, err := q.GetCommentReplyThumb(c, &pgc.GetCommentReplyThumbParams{
			UserID:    userId,
			TweetID:   tweetId,
			CommentID: commentId,
			ReplyID:   pgtype.Int8{Int64: replyId, Valid: true},
		})
		yes, no := int16(1), int16(0)
		if err == nil {
			switch {
			case commentThumbs.IsThumbsDown == yes:
				thumbsUpCount, thumbsDownCount = 0, -1
			case commentThumbs.IsThumbsUp == no && commentThumbs.IsThumbsDown == no:
				thumbsUpCount, thumbsDownCount = 0, 1
			default:
				thumbsUpCount, thumbsDownCount = -1, 1
				commentThumbs.IsThumbsUp = no
			}
			commentThumbs.IsThumbsDown = 1 - commentThumbs.IsThumbsDown
			if err = q.UpdateThumbsUpDownComment(c, &pgc.UpdateThumbsUpDownCommentParams{
				IsThumbsUp:   commentThumbs.IsThumbsUp,
				IsThumbsDown: commentThumbs.IsThumbsDown,
				ModifiedOn:   now,
				ID:           commentThumbs.ID,
			}); err != nil {
				return err
			}
		} else {
			thumbsUpCount, thumbsDownCount = 0, 1
			if _, err = q.CreateThumbsUpDownComment(c, &pgc.CreateThumbsUpDownCommentParams{
				UserID:       userId,
				TweetID:      tweetId,
				CommentID:    commentId,
				ReplyID:      pgtype.Int8{Int64: replyId, Valid: true},
				IsThumbsUp:   no,
				IsThumbsDown: yes,
				CommentType:  1,
				CreatedOn:    now,
			}); err != nil {
				return err
			}
		}
		// 更新comment_reply thumbsUpCount
		return q.UpdateReplyThumbsCount(c, &pgc.UpdateReplyThumbsCountParams{
			ThumbsUpCount:   thumbsUpCount,
			ThumbsDownCount: thumbsDownCount,
			ModifiedOn:      now,
			ID:              replyId,
		})
	})
}

func newCommentService(db *pgx.Conn) core.CommentService {
	return &commentSrv{
		pgcSrv: newPgcSrv(db),
	}
}

func newCommentManageService(db *pgx.Conn) core.CommentManageService {
	return &commentManageSrv{
		pgcSrv: newPgcSrv(db),
	}
}
