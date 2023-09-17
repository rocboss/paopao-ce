// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

	"github.com/alimy/mir/v4"
	"github.com/alimy/tryst/lets"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Loose = (*looseSrv)(nil)
)

type looseSrv struct {
	api.UnimplementedLooseServant
	*base.DaoServant
	ac                       core.AppCache
	userTweetsExpire         int64
	idxTweetsExpire          int64
	prefixUserTweets         string
	prefixIdxTweetsNewest    string
	prefixIdxTweetsHots      string
	prefixIdxTweetsFollowing string
}

func (s *looseSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func (s *looseSrv) Timeline(req *web.TimelineReq) (*web.TimelineResp, mir.Error) {
	limit, offset := req.PageSize, (req.Page-1)*req.PageSize
	if req.Query == "" && req.Type == "search" {
		return s.getIndexTweets(req, limit, offset)
	}
	q := &core.QueryReq{
		Query: req.Query,
		Type:  core.SearchType(req.Type),
	}
	res, err := s.Ts.Search(req.User, q, offset, limit)
	if err != nil {
		logrus.Errorf("Ts.Search err: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	posts, err := s.Ds.RevampPosts(res.Items)
	if err != nil {
		logrus.Errorf("Ds.RevampPosts err: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(posts, req.Page, req.PageSize, res.Total)
	return &web.TimelineResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) getIndexTweets(req *web.TimelineReq, limit int, offset int) (res *web.TimelineResp, err mir.Error) {
	// 尝试直接从缓存中获取数据
	key, ok := "", false
	if res, key, ok = s.indexTweetsFromCache(req, limit, offset); ok {
		// logrus.Debugf("getIndexTweets from cache key:%s", key)
		return
	}
	var (
		posts []*ms.Post
		total int64
		xerr  error
	)
	switch req.Style {
	case web.StyleTweetsFollowing:
		if req.User != nil {
			posts, total, xerr = s.Ds.ListFollowingTweets(req.User.ID, limit, offset)
		} else {
			return nil, web.ErrGetPostsNilUser
		}
	case web.StyleTweetsNewest:
		posts, total, xerr = s.Ds.ListIndexTweets(cs.StyleIndexTweetsNewest, limit, offset)
	case web.StyleTweetsHots:
		posts, total, xerr = s.Ds.ListIndexTweets(cs.StyleIndexTweetsHots, limit, offset)
	default:
		return nil, web.ErrGetPostsUnknowStyle
	}
	if xerr != nil {
		logrus.Errorf("getIndexTweets occurs error: %s", xerr)
		return nil, web.ErrGetPostFailed
	}
	postsFormated, verr := s.Ds.MergePosts(posts)
	if verr != nil {
		logrus.Errorf("getIndexTweets in merge posts occurs error: %s", verr)
		return nil, web.ErrGetPostFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, total)
	// 缓存处理
	base.OnCacheRespEvent(s.ac, key, resp, s.idxTweetsExpire)
	return &web.TimelineResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) indexTweetsFromCache(req *web.TimelineReq, limit int, offset int) (res *web.TimelineResp, key string, ok bool) {
	switch req.Style {
	case web.StyleTweetsFollowing:
		username := lets.If(req.User != nil, req.User.Username, "_")
		key = fmt.Sprintf("%s%s:%d:%d", s.prefixIdxTweetsFollowing, username, offset, limit)
	case web.StyleTweetsNewest:
		key = fmt.Sprintf("%s%d:%d", s.prefixIdxTweetsNewest, offset, limit)
	case web.StyleTweetsHots:
		key = fmt.Sprintf("%s%d:%d", s.prefixIdxTweetsHots, offset, limit)
	default:
		return
	}
	if data, err := s.ac.Get(key); err == nil {
		ok, res = true, &web.TimelineResp{
			CachePageResp: joint.CachePageResp{
				JsonResp: data,
			},
		}
	}
	return
}

func (s *looseSrv) GetUserTweets(req *web.GetUserTweetsReq) (res *web.GetUserTweetsResp, err mir.Error) {
	user, xerr := s.RelationTypFrom(req.User, req.Username)
	if xerr != nil {
		return nil, err
	}
	// 尝试直接从缓存中获取数据
	key, ok := "", false
	if res, key, ok = s.userTweetsFromCache(req, user); ok {
		// logrus.Debugf("GetUserTweets from cache key:%s", key)
		return
	}
	// 缓存获取未成功，只能查库了
	switch req.Style {
	case web.UserPostsStyleComment, web.UserPostsStyleMedia:
		res, err = s.listUserTweets(req, user)
	case web.UserPostsStyleHighlight:
		res, err = s.getUserPostTweets(req, user, true)
	case web.UserPostsStyleStar:
		res, err = s.getUserStarTweets(req, user)
	case web.UserPostsStylePost:
		fallthrough
	default:
		res, err = s.getUserPostTweets(req, user, false)
	}
	// 缓存处理
	if err == nil {
		base.OnCacheRespEvent(s.ac, key, res.Data, s.userTweetsExpire)
	}
	return
}

func (s *looseSrv) userTweetsFromCache(req *web.GetUserTweetsReq, user *cs.VistUser) (res *web.GetUserTweetsResp, key string, ok bool) {
	switch req.Style {
	case web.UserPostsStylePost, web.UserPostsStyleHighlight, web.UserPostsStyleMedia:
		key = fmt.Sprintf("%s%s:%s:%s:%d:%d", s.prefixUserTweets, req.Username, req.Style, user.RelTyp, req.Page, req.PageSize)
	default:
		visitUserName := lets.If(user.RelTyp != cs.RelationGuest, user.Username, "_")
		key = fmt.Sprintf("%s%s:%s:%s:%d:%d", s.prefixUserTweets, req.Username, req.Style, visitUserName, req.Page, req.PageSize)
	}
	if data, err := s.ac.Get(key); err == nil {
		ok, res = true, &web.GetUserTweetsResp{
			CachePageResp: joint.CachePageResp{
				JsonResp: data,
			},
		}
	}
	return
}

func (s *looseSrv) getUserStarTweets(req *web.GetUserTweetsReq, user *cs.VistUser) (*web.GetUserTweetsResp, mir.Error) {
	stars, totalRows, err := s.Ds.ListUserStarTweets(user, req.PageSize, (req.Page-1)*req.PageSize)
	if err != nil {
		logrus.Errorf("Ds.GetUserPostStars err: %s", err)
		return nil, web.ErrGetStarsFailed
	}
	var posts []*ms.Post
	for _, star := range stars {
		if star.Post != nil {
			posts = append(posts, star.Post)
		}
	}
	postsFormated, err := s.Ds.MergePosts(posts)
	if err != nil {
		logrus.Errorf("Ds.MergePosts err: %s", err)
		return nil, web.ErrGetStarsFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, totalRows)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) listUserTweets(req *web.GetUserTweetsReq, user *cs.VistUser) (*web.GetUserTweetsResp, mir.Error) {
	var (
		tweets []*ms.Post
		total  int64
		err    error
	)
	if req.Style == web.UserPostsStyleComment {
		tweets, total, err = s.Ds.ListUserCommentTweets(user, req.PageSize, (req.Page-1)*req.PageSize)
	} else if req.Style == web.UserPostsStyleMedia {
		tweets, total, err = s.Ds.ListUserMediaTweets(user, req.PageSize, (req.Page-1)*req.PageSize)
	} else {
		logrus.Errorf("s.listUserTweets unknow style: %s", req.Style)
		return nil, web.ErrGetPostsFailed
	}
	if err != nil {
		logrus.Errorf("s.listUserTweets err: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	postFormated, err := s.Ds.MergePosts(tweets)
	if err != nil {
		logrus.Errorf("s.listUserTweets err: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(postFormated, req.Page, req.PageSize, total)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) getUserPostTweets(req *web.GetUserTweetsReq, user *cs.VistUser, isHighlight bool) (*web.GetUserTweetsResp, mir.Error) {
	style := cs.StyleUserTweetsGuest
	switch user.RelTyp {
	case cs.RelationAdmin:
		style = cs.StyleUserTweetsAdmin
	case cs.RelationSelf:
		style = cs.StyleUserTweetsSelf
	case cs.RelationFriend:
		style = cs.StyleUserTweetsFriend
	case cs.RelationFollowing:
		style = cs.StyleUserTweetsFollowing
	case cs.RelationGuest:
		fallthrough
	default:
		// nothing
	}
	posts, total, err := s.Ds.ListUserTweets(user.UserId, style, isHighlight, req.PageSize, (req.Page-1)*req.PageSize)
	if err != nil {
		logrus.Errorf("s.GetTweetList error[1]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	postsFormated, xerr := s.Ds.MergePosts(posts)
	if xerr != nil {
		logrus.Errorf("s.GetTweetList error[2]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, total)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) GetUserProfile(req *web.GetUserProfileReq) (*web.GetUserProfileResp, mir.Error) {
	he, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err: %s", err)
		return nil, web.ErrNoExistUsername
	}
	if he.Model == nil && he.ID <= 0 {
		return nil, web.ErrNoExistUsername
	}
	// 设定自己不是自己的朋友
	isFriend := !(req.User == nil || req.User.ID == he.ID)
	if req.User != nil && req.User.ID != he.ID {
		isFriend = s.Ds.IsFriend(req.User.ID, he.ID)
	}
	isFollowing := false
	if req.User != nil {
		isFollowing = s.Ds.IsFollow(req.User.ID, he.ID)
	}
	follows, followings, err := s.Ds.GetFollowCount(he.ID)
	if err != nil {
		return nil, web.ErrGetPostsFailed
	}
	return &web.GetUserProfileResp{
		ID:          he.ID,
		Nickname:    he.Nickname,
		Username:    he.Username,
		Status:      he.Status,
		Avatar:      he.Avatar,
		IsAdmin:     he.IsAdmin,
		IsFriend:    isFriend,
		IsFollowing: isFollowing,
		CreatedOn:   he.CreatedOn,
		Follows:     follows,
		Followings:  followings,
	}, nil
}

func (s *looseSrv) TopicList(req *web.TopicListReq) (*web.TopicListResp, mir.Error) {
	var (
		tags, extralTags cs.TagList
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
		err = web.ErrGetPostTagsFailed
	}
	if err != nil {
		return nil, web.ErrGetPostTagsFailed
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
	conditions := &ms.ConditionsT{
		"post_id": req.TweetId,
		"ORDER":   sort,
	}

	comments, err := s.Ds.GetComments(conditions, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		return nil, web.ErrGetCommentsFailed
	}

	userIDs := []int64{}
	commentIDs := []int64{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
		commentIDs = append(commentIDs, comment.ID)
	}

	users, err := s.Ds.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, web.ErrGetCommentsFailed
	}

	contents, err := s.Ds.GetCommentContentsByIDs(commentIDs)
	if err != nil {
		return nil, web.ErrGetCommentsFailed
	}

	replies, err := s.Ds.GetCommentRepliesByID(commentIDs)
	if err != nil {
		return nil, web.ErrGetCommentsFailed
	}

	var commentThumbs, replyThumbs cs.CommentThumbsMap
	if req.Uid > 0 {
		commentThumbs, replyThumbs, err = s.Ds.GetCommentThumbsMap(req.Uid, req.TweetId)
		if err != nil {
			return nil, web.ErrGetCommentsFailed
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

	commentsFormated := []*ms.CommentFormated{}
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

func newLooseSrv(s *base.DaoServant, ac core.AppCache) api.Loose {
	return &looseSrv{
		DaoServant:       s,
		ac:               ac,
		userTweetsExpire: conf.CacheSetting.UserTweetsExpire,
		prefixUserTweets: conf.PrefixUserTweets,
	}
}
