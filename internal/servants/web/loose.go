// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Loose        = (*looseSrv)(nil)
	_ api.LooseBinding = (*looseBinding)(nil)
	_ api.LooseRender  = (*looseRender)(nil)
)

type looseSrv struct {
	api.UnimplementedLooseServant
	*base.DaoServant
}

type looseBinding struct {
	*api.UnimplementedLooseBinding
}

type looseRender struct {
	*api.UnimplementedLooseRender
}

func (s *looseBinding) BindTimeline(c *gin.Context) (*web.TimelineReq, mir.Error) {
	user, _ := base.UserFrom(c)
	page, pageSize := app.GetPageInfo(c)
	v := &web.TimelineReq{
		BaseInfo: web.BaseInfo{
			User: user,
		},
		Query:    c.Query("query"),
		Type:     "search",
		Page:     page,
		PageSize: pageSize,
	}
	if c.Query("type") == "tag" {
		v.Type = "tag"
	}
	return v, nil
}

func (s *looseSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func (s *looseSrv) Timeline(req *web.TimelineReq) (*web.TimelineResp, mir.Error) {
	var resp *base.PageResp
	offset, limit := (req.Page-1)*req.PageSize, req.PageSize
	if req.Query == "" && req.Type == "search" {
		res, err := s.Ds.IndexPosts(req.User, offset, limit)
		if err != nil {
			logrus.Errorf("Ds.IndexPosts err: %s", err)
			return nil, _errGetPostsFailed
		}
		resp = base.PageRespFrom(res.Tweets, req.Page, req.PageSize, res.Total)
	} else {
		q := &core.QueryReq{
			Query: req.Query,
			Type:  core.SearchType(req.Type),
		}
		res, err := s.Ts.Search(req.User, q, offset, limit)
		if err != nil {
			logrus.Errorf("Ts.Search err: %s", err)
			return nil, _errGetPostsFailed
		}
		posts, err := s.Ds.RevampPosts(res.Items)
		if err != nil {
			logrus.Errorf("Ds.RevampPosts err: %s", err)
			return nil, _errGetPostsFailed
		}
		resp = base.PageRespFrom(posts, req.Page, req.PageSize, res.Total)
	}
	return (*web.TimelineResp)(resp), nil
}

func (s *looseSrv) GetUserTweets(req *web.GetUserTweetsReq) (*web.GetUserTweetsResp, mir.Error) {
	other, xerr := s.GetUserProfile(&web.GetUserProfileReq{
		BaseInfo: req.BaseInfo,
		Username: req.Username,
	})
	if xerr != nil {
		return nil, xerr
	}

	visibilities := []core.PostVisibleT{core.PostVisitPublic}
	if req.User != nil {
		if req.User.ID == other.ID || req.User.IsAdmin {
			visibilities = append(visibilities, core.PostVisitPrivate, core.PostVisitFriend)
		} else if other.IsFriend {
			visibilities = append(visibilities, core.PostVisitFriend)
		}
	}
	conditions := &core.ConditionsT{
		"user_id":         other.ID,
		"visibility IN ?": visibilities,
		"ORDER":           "latest_replied_on DESC",
	}
	_, posts, err := s.GetTweetList(conditions, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("s.GetTweetList err: %s", err)
		return nil, _errGetPostsFailed
	}
	totalRows, err := s.Ds.GetPostCount(conditions)
	if err != nil {
		logrus.Errorf("s.GetPostCount err: %s", err)
		return nil, _errGetPostsFailed
	}

	resp := base.PageRespFrom(posts, req.Page, req.PageSize, totalRows)
	return (*web.GetUserTweetsResp)(resp), nil
}

func (s *looseSrv) GetUserProfile(req *web.GetUserProfileReq) (*web.GetUserProfileResp, mir.Error) {
	he, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err: %s", err)
		return nil, _errNoExistUsername
	}
	if he.Model == nil && he.ID <= 0 {
		return nil, _errNoExistUsername
	}
	// 设定自己不是自己的朋友
	isFriend := !(req.User == nil || req.User.ID == he.ID)
	if req.User != nil && req.User.ID != he.ID {
		isFriend = s.Ds.IsFriend(req.User.ID, he.ID)
	}
	return &web.GetUserProfileResp{
		ID:       he.ID,
		Nickname: he.Nickname,
		Username: he.Username,
		Status:   he.Status,
		Avatar:   he.Avatar,
		IsAdmin:  he.IsAdmin,
		IsFriend: isFriend,
	}, nil
}

func (s *looseSrv) TopicList(req *web.TopicListReq) (*web.TopicListResp, mir.Error) {
	var (
		tags, extralTags []*core.TagFormated
		err              error
	)
	num := req.Num
	switch req.Type {
	case web.TagTypeHot:
		tags, err = s.Ds.GetHotTags(req.Uid, num, 0)
	case web.TagTypeNew:
		tags, err = s.Ds.GetNewestTags(req.Uid, num, 0)
	case web.TagTypeFollow:
		tags, err = s.Ds.GetFollowTags(req.Uid, num, 0)
	case web.TagTypeHotExtral:
		extralNum := req.ExtralNum
		if extralNum <= 0 {
			extralNum = num
		}
		tags, err = s.Ds.GetHotTags(req.Uid, num, 0)
		if err == nil {
			extralTags, err = s.Ds.GetFollowTags(req.Uid, extralNum, 0)
		}
	default:
		// TODO: return good error
		err = _errGetPostTagsFailed
	}
	if err != nil {
		return nil, _errGetPostTagsFailed
	}
	return &web.TopicListResp{
		Topics:       tags,
		ExtralTopics: extralTags,
	}, nil
}

func (s *looseSrv) TweetComments(req *web.TweetCommentsReq) (*web.TweetCommentsResp, mir.Error) {
	sort := "id ASC"
	if req.SortStrategy == "newest" {
		sort = "id DESC"
	}
	conditions := &core.ConditionsT{
		"post_id": req.TweetId,
		"ORDER":   sort,
	}

	comments, err := s.Ds.GetComments(conditions, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		return nil, _errGetCommentsFailed
	}

	userIDs := []int64{}
	commentIDs := []int64{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
		commentIDs = append(commentIDs, comment.ID)
	}

	users, err := s.Ds.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, _errGetCommentsFailed
	}

	contents, err := s.Ds.GetCommentContentsByIDs(commentIDs)
	if err != nil {
		return nil, _errGetCommentsFailed
	}

	replies, err := s.Ds.GetCommentRepliesByID(commentIDs)
	if err != nil {
		return nil, _errGetCommentsFailed
	}

	var commentThumbs, replyThumbs cs.CommentThumbsMap
	if req.Uid > 0 {
		commentThumbs, replyThumbs, err = s.Ds.GetCommentThumbsMap(req.Uid, req.TweetId)
		if err != nil {
			return nil, _errGetCommentsFailed
		}
	}

	replyMap := make(map[int64][]*dbr.CommentReplyFormated)
	if len(replyThumbs) > 0 {
		for _, reply := range replies {
			if thumbs, exist := replyThumbs[reply.ID]; exist {
				reply.IsThumbsUp, reply.IsThumbsDown = thumbs.IsThumbsUp, thumbs.IsThumbsDown
			}
			replyMap[reply.CommentID] = append(replyMap[reply.CommentID], reply)
		}
	} else {
		for _, reply := range replies {
			replyMap[reply.CommentID] = append(replyMap[reply.CommentID], reply)
		}
	}

	commentsFormated := []*core.CommentFormated{}
	for _, comment := range comments {
		commentFormated := comment.Format()
		if thumbs, exist := commentThumbs[comment.ID]; exist {
			commentFormated.IsThumbsUp, commentFormated.IsThumbsDown = thumbs.IsThumbsUp, thumbs.IsThumbsDown
		}
		for _, content := range contents {
			if content.CommentID == comment.ID {
				commentFormated.Contents = append(commentFormated.Contents, content)
			}
		}
		if replySlice, exist := replyMap[commentFormated.ID]; exist {
			commentFormated.Replies = replySlice
		}
		for _, user := range users {
			if user.ID == comment.UserID {
				commentFormated.User = user.Format()
			}
		}
		commentsFormated = append(commentsFormated, commentFormated)
	}

	// 获取总量
	totalRows, _ := s.Ds.GetCommentCount(conditions)
	resp := base.PageRespFrom(commentsFormated, req.Page, req.PageSize, totalRows)
	return (*web.TweetCommentsResp)(resp), nil
}

func newLooseSrv(s *base.DaoServant) api.Loose {
	return &looseSrv{
		DaoServant: s,
	}
}

func newLooseBinding() api.LooseBinding {
	return &looseBinding{
		UnimplementedLooseBinding: &api.UnimplementedLooseBinding{
			BindAny: base.NewBindAnyFn(),
		},
	}
}

func newLooseRender() api.LooseRender {
	return &looseRender{
		UnimplementedLooseRender: &api.UnimplementedLooseRender{
			RenderAny: base.RenderAny,
		},
	}
}
