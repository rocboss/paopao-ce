// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

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
	tweetCommentsExpire      int64
	prefixUserTweets         string
	prefixIdxTweetsNewest    string
	prefixIdxTweetsHots      string
	prefixIdxTweetsFollowing string
	prefixTweetComment       string
}

func (s *looseSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JwtLoose()}
}

func (s *looseSrv) Timeline(req *web.TimelineReq) (*web.TimelineResp, error) {
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
	userId := int64(-1)
	if req.User != nil {
		userId = req.User.ID
	}
	if err := s.PrepareTweets(userId, posts); err != nil {
		logrus.Errorf("timeline occurs error[2]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(posts, req.Page, req.PageSize, res.Total)
	return &web.TimelineResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) getIndexTweets(req *web.TimelineReq, limit int, offset int) (res *web.TimelineResp, err error) {
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
			// return nil, web.ErrGetPostsNilUser
			// 宽松处理，前端退出登录后马上获取动态列表，可能错误走到这里
			posts, total, xerr = s.Ds.ListIndexNewestTweets(limit, offset)
		}
	case web.StyleTweetsNewest:
		posts, total, xerr = s.Ds.ListIndexNewestTweets(limit, offset)
	case web.StyleTweetsHots:
		posts, total, xerr = s.Ds.ListIndexHotsTweets(limit, offset)
	default:
		return nil, web.ErrGetPostsUnknowStyle
	}
	if xerr != nil {
		logrus.Errorf("getIndexTweets occurs error[1]: %s", xerr)
		return nil, web.ErrGetPostFailed
	}
	postsFormated, verr := s.Ds.MergePosts(posts)
	if verr != nil {
		logrus.Errorf("getIndexTweets in merge posts occurs error: %s", verr)
		return nil, web.ErrGetPostFailed
	}
	userId := int64(-1)
	if req.User != nil {
		userId = req.User.ID
	}
	if err := s.PrepareTweets(userId, postsFormated); err != nil {
		logrus.Errorf("getIndexTweets occurs error[2]: %s", err)
		return nil, web.ErrGetPostsFailed
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
	username := "_"
	if req.User != nil {
		username = req.User.Username
	}
	switch req.Style {
	case web.StyleTweetsFollowing:
		key = fmt.Sprintf("%s%s:%d:%d", s.prefixIdxTweetsFollowing, username, offset, limit)
	case web.StyleTweetsNewest:
		key = fmt.Sprintf("%s%s:%d:%d", s.prefixIdxTweetsNewest, username, offset, limit)
	case web.StyleTweetsHots:
		key = fmt.Sprintf("%s%s:%d:%d", s.prefixIdxTweetsHots, username, offset, limit)
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

func (s *looseSrv) tweetCommentsFromCache(req *web.TweetCommentsReq, limit int, offset int) (res *web.TweetCommentsResp, key string, ok bool) {
	key = fmt.Sprintf("%s%d:%s:%d:%d", s.prefixTweetComment, req.TweetId, req.Style, limit, offset)
	if data, err := s.ac.Get(key); err == nil {
		ok, res = true, &web.TweetCommentsResp{
			CachePageResp: joint.CachePageResp{
				JsonResp: data,
			},
		}
	}
	return
}

func (s *looseSrv) GetUserTweets(req *web.GetUserTweetsReq) (res *web.GetUserTweetsResp, err error) {
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
		key = fmt.Sprintf("%s%d:%s:%s:%d:%d", s.prefixUserTweets, user.UserId, req.Style, user.RelTyp, req.Page, req.PageSize)
	default:
		meName := "_"
		if user.RelTyp != cs.RelationGuest {
			meName = req.User.Username
		}
		key = fmt.Sprintf("%s%d:%s:%s:%d:%d", s.prefixUserTweets, user.UserId, req.Style, meName, req.Page, req.PageSize)
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

func (s *looseSrv) getUserStarTweets(req *web.GetUserTweetsReq, user *cs.VistUser) (*web.GetUserTweetsResp, error) {
	stars, totalRows, err := s.Ds.ListUserStarTweets(user, req.PageSize, (req.Page-1)*req.PageSize)
	if err != nil {
		logrus.Errorf("getUserStarTweets err[1]: %s", err)
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
	userId := int64(-1)
	if req.User != nil {
		userId = req.User.ID
	}
	if err := s.PrepareTweets(userId, postsFormated); err != nil {
		logrus.Errorf("getUserStarTweets err[2]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, totalRows)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) listUserTweets(req *web.GetUserTweetsReq, user *cs.VistUser) (*web.GetUserTweetsResp, error) {
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
		logrus.Errorf("s.listUserTweets unknow style[1]: %s", req.Style)
		return nil, web.ErrGetPostsFailed
	}
	if err != nil {
		logrus.Errorf("s.listUserTweets err[2]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	postsFormated, err := s.Ds.MergePosts(tweets)
	if err != nil {
		logrus.Errorf("s.listUserTweets err[3]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	userId := int64(-1)
	if req.User != nil {
		userId = req.User.ID
	}
	if err := s.PrepareTweets(userId, postsFormated); err != nil {
		logrus.Errorf("s.listUserTweets err[4]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, total)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) getUserPostTweets(req *web.GetUserTweetsReq, user *cs.VistUser, isHighlight bool) (*web.GetUserTweetsResp, error) {
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
	userId := int64(-1)
	if req.User != nil {
		userId = req.User.ID
	}
	if err := s.PrepareTweets(userId, postsFormated); err != nil {
		logrus.Errorf("s.GetTweetList error[3]: %s", err)
		return nil, web.ErrGetPostsFailed
	}
	resp := joint.PageRespFrom(postsFormated, req.Page, req.PageSize, total)
	return &web.GetUserTweetsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) GetUserProfile(req *web.GetUserProfileReq) (*web.GetUserProfileResp, error) {
	he, err := s.Ds.UserProfileByName(req.Username)
	if err != nil {
		logrus.Errorf("looseSrv.GetUserProfile occurs error[1]: %s", err)
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
		TweetsCount: he.TweetsCount,
	}, nil
}

func (s *looseSrv) TopicList(req *web.TopicListReq) (*web.TopicListResp, error) {
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
		tags, err = s.Ds.GetFollowTags(req.Uid, false, num, 0)
	case web.TagTypePin:
		tags, err = s.Ds.GetFollowTags(req.Uid, true, num, 0)
	case web.TagTypeHotExtral:
		extralNum := req.ExtralNum
		if extralNum <= 0 {
			extralNum = num
		}
		tags, err = s.Ds.GetHotTags(req.Uid, num, 0)
		if err == nil {
			extralTags, err = s.Ds.GetFollowTags(req.Uid, false, extralNum, 0)
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

func (s *looseSrv) TweetComments(req *web.TweetCommentsReq) (res *web.TweetCommentsResp, err error) {
	limit, offset := req.PageSize, (req.Page-1)*req.PageSize
	// 尝试直接从缓存中获取数据
	key, ok := "", false
	if res, key, ok = s.tweetCommentsFromCache(req, limit, offset); ok {
		logrus.Debugf("looseSrv.TweetComments from cache key:%s", key)
		return
	}

	comments, totalRows, xerr := s.Ds.GetComments(req.TweetId, req.Style.ToInnerValue(), limit, offset)
	if xerr != nil {
		logrus.Errorf("looseSrv.TweetComments occurs error[1]: %s", xerr)
		return nil, web.ErrGetCommentsFailed
	}

	userIDs := []int64{}
	commentIDs := []int64{}
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
		commentIDs = append(commentIDs, comment.ID)
	}

	users, xerr := s.Ds.GetUsersByIDs(userIDs)
	if xerr != nil {
		logrus.Errorf("looseSrv.TweetComments occurs error[2]: %s", xerr)
		return nil, web.ErrGetCommentsFailed
	}

	contents, xerr := s.Ds.GetCommentContentsByIDs(commentIDs)
	if xerr != nil {
		logrus.Errorf("looseSrv.TweetComments occurs error[3]: %s", xerr)
		return nil, web.ErrGetCommentsFailed
	}

	replies, xerr := s.Ds.GetCommentRepliesByID(commentIDs)
	if xerr != nil {
		logrus.Errorf("looseSrv.TweetComments occurs error[4]: %s", xerr)
		return nil, web.ErrGetCommentsFailed
	}

	var commentThumbs, replyThumbs cs.CommentThumbsMap
	if req.Uid > 0 {
		commentThumbs, replyThumbs, xerr = s.Ds.GetCommentThumbsMap(req.Uid, req.TweetId)
		if xerr != nil {
			logrus.Errorf("looseSrv.TweetComments occurs error[5]: %s", xerr)
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
	resp := joint.PageRespFrom(commentsFormated, req.Page, req.PageSize, totalRows)
	// 缓存处理
	base.OnCacheRespEvent(s.ac, key, resp, s.tweetCommentsExpire)
	return &web.TweetCommentsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *looseSrv) TweetDetail(req *web.TweetDetailReq) (*web.TweetDetailResp, error) {
	post, err := s.Ds.GetPostByID(req.TweetId)
	if err != nil {
		return nil, web.ErrGetPostFailed
	}
	postContents, err := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, web.ErrGetPostFailed
	}
	users, err := s.Ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, web.ErrGetPostFailed
	}
	// 数据整合
	postFormated := post.Format()
	for _, user := range users {
		postFormated.User = user.Format()
	}
	for _, content := range postContents {
		if content.PostID == post.ID {
			postFormated.Contents = append(postFormated.Contents, content.Format())
		}
	}
	if err = s.PrepareTweet(req.User, postFormated); err != nil {
		return nil, web.ErrGetPostFailed
	}
	// 检测访问权限
	// TODO: 提到最前面去检测
	switch {
	case req.User != nil && (req.User.ID == postFormated.User.ID || req.User.IsAdmin):
		// read by self of super admin
		break
	case post.Visibility == core.PostVisitPublic:
		break
	case post.Visibility == core.PostVisitFriend && postFormated.User.IsFriend:
		break
	case post.Visibility == core.PostVisitFollowing && postFormated.User.IsFollowing:
		break
	default:
		return nil, web.ErrNoPermission
	}
	return (*web.TweetDetailResp)(postFormated), nil
}

func newLooseSrv(s *base.DaoServant, ac core.AppCache) api.Loose {
	cs := conf.CacheSetting
	return &looseSrv{
		DaoServant:               s,
		ac:                       ac,
		userTweetsExpire:         cs.UserTweetsExpire,
		idxTweetsExpire:          cs.IndexTweetsExpire,
		tweetCommentsExpire:      cs.TweetCommentsExpire,
		prefixUserTweets:         conf.PrefixUserTweets,
		prefixIdxTweetsNewest:    conf.PrefixIdxTweetsNewest,
		prefixIdxTweetsHots:      conf.PrefixIdxTweetsHots,
		prefixIdxTweetsFollowing: conf.PrefixIdxTweetsFollowing,
		prefixTweetComment:       conf.PrefixTweetComment,
	}
}
