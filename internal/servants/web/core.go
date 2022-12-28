// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Core        = (*coreSrv)(nil)
	_ api.CoreBinding = (*coreBinding)(nil)
	_ api.CoreRender  = (*coreRender)(nil)
)

type coreSrv struct {
	api.UnimplementedCoreServant
	*base.DaoServant

	ts  core.TweetSearchService
	oss core.ObjectStorageService
}

type coreBinding struct {
	*api.UnimplementedCoreBinding
}

type coreRender struct {
	*api.UnimplementedCoreRender
}

func (b *coreBinding) BindSyncSearchIndex(c *gin.Context) (*web.SyncSearchIndexReq, mir.Error) {
	user, _ := base.UserFrom(c)
	return &web.SyncSearchIndexReq{
		BaseInfo: &web.BaseInfo{
			User: user,
		},
		Ctx: c.Request.Context(),
	}, nil
}

func (b *coreBinding) BindGetUserInfo(c *gin.Context) (*web.UserInfoReq, mir.Error) {
	username, exists := base.UserNameFrom(c)
	if !exists {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	return &web.UserInfoReq{
		Username: username,
	}, nil
}

func (b *coreBinding) BindGetMessages(c *gin.Context) (*web.GetMessagesReq, mir.Error) {
	uid, ok := base.UserIdFrom(c)
	if !ok {
		return nil, xerror.UnauthorizedTokenError
	}
	page, pageSize := app.GetPageOffset(c)
	return &web.GetMessagesReq{
		UserId:   uid,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (s *coreSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *coreSrv) SyncSearchIndex(req *web.SyncSearchIndexReq) mir.Error {
	if req.User != nil && req.User.IsAdmin {
		go s.pushPostsToSearch(req.Ctx)
	}
	return nil
}

func (s *coreSrv) GetUserInfo(req *web.UserInfoReq) (*web.UserInfoResp, mir.Error) {
	user, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	if user.Model == nil || user.ID < 0 {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	resp := &web.UserInfoResp{
		Id:       user.ID,
		Nickname: user.Nickname,
		Username: user.Username,
		Status:   user.Status,
		Avatar:   user.Avatar,
		Balance:  user.Balance,
		IsAdmin:  user.IsAdmin,
	}
	if user.Phone != "" && len(user.Phone) == 11 {
		resp.Phone = user.Phone[0:3] + "****" + user.Phone[7:]
	}
	return resp, nil
}

func (s *coreSrv) GetUnreadMsgCount(req *web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, mir.Error) {
	count, err := s.Ds.GetUnreadCount(req.Uid)
	if err != nil {
		return nil, xerror.ServerError
	}
	return &web.GetUnreadMsgCountResp{
		Count: count,
	}, nil
}

func (s *coreSrv) GetMessages(req *web.GetMessagesReq) (*base.PageResp, mir.Error) {
	conditions := &core.ConditionsT{
		"receiver_user_id": req.UserId,
		"ORDER":            "id DESC",
	}
	messages, err := s.Ds.GetMessages(conditions, (req.Page-1)*req.PageSize, req.PageSize)
	for _, mf := range messages {
		if mf.SenderUserID > 0 {
			user, err := s.Ds.GetUserByID(mf.SenderUserID)
			if err == nil {
				mf.SenderUser = user.Format()
			}
		}
		// 好友申请消息不需要获取其他信息
		if mf.Type == core.MsgTypeRequestingFriend {
			continue
		}
		if mf.PostID > 0 {
			post, err := s.GetTweetBy(mf.PostID)
			if err == nil {
				mf.Post = post
				if mf.CommentID > 0 {
					comment, err := s.Ds.GetCommentByID(mf.CommentID)
					if err == nil {
						mf.Comment = comment
						if mf.ReplyID > 0 {
							reply, err := s.Ds.GetCommentReplyByID(mf.ReplyID)
							if err == nil {
								mf.Reply = reply
							}
						}
					}
				}
			}
		}
	}
	if err != nil {
		logrus.Errorf("Ds.GetMessages err: %v\n", err)
		return nil, _errGetMessagesFailed
	}
	totalRows, _ := s.Ds.GetMessageCount(conditions)
	return base.PageRespFrom(messages, req.Page, req.PageSize, totalRows), nil
}

func (s *coreSrv) ReadMessage(req *web.ReadMessageReq) mir.Error {
	message, err := s.Ds.GetMessageByID(req.ID)
	if err != nil {
		return _errReadMessageFailed
	}
	if message.ReceiverUserID != req.Uid {
		return _errNoPermission
	}
	if err = s.Ds.ReadMessage(message); err != nil {
		logrus.Errorf("Ds.ReadMessage err: %s", err)
		return _errReadMessageFailed
	}
	return nil
}

func (s *coreSrv) pushPostsToSearch(c context.Context) {
	if ok, _ := s.Redis.SetNX(c, "JOB_PUSH_TO_SEARCH", 1, time.Hour).Result(); ok {
		defer s.Redis.Del(c, "JOB_PUSH_TO_SEARCH")

		splitNum := 1000
		totalRows, _ := s.Ds.GetPostCount(&core.ConditionsT{
			"visibility IN ?": []core.PostVisibleT{core.PostVisitPublic, core.PostVisitFriend},
		})
		pages := math.Ceil(float64(totalRows) / float64(splitNum))
		nums := int(pages)
		for i := 0; i < nums; i++ {
			posts, postsFormated, err := s.getTweetList(&core.ConditionsT{}, i*splitNum, splitNum)
			if err != nil || len(posts) != len(postsFormated) {
				continue
			}
			for i, pf := range postsFormated {
				contentFormated := ""
				for _, content := range pf.Contents {
					if content.Type == core.ContentTypeText || content.Type == core.ContentTypeTitle {
						contentFormated = contentFormated + content.Content + "\n"
					}
				}
				docs := []core.TsDocItem{{
					Post:    posts[i],
					Content: contentFormated,
				}}
				s.ts.AddDocuments(docs, fmt.Sprintf("%d", posts[i].ID))
			}
		}
	}
}

func (s *coreSrv) deleteSearchPost(post *core.Post) error {
	return s.ts.DeleteDocuments([]string{fmt.Sprintf("%d", post.ID)})
}

func (s *coreSrv) getTweetList(conditions *core.ConditionsT, offset, limit int) ([]*core.Post, []*core.PostFormated, error) {
	posts, err := s.Ds.GetPosts(conditions, offset, limit)
	if err != nil {
		return nil, nil, err
	}
	postFormated, err := s.Ds.MergePosts(posts)
	return posts, postFormated, err
}

func newCoreSrv(s *base.DaoServant, ts core.TweetSearchService, oss core.ObjectStorageService) api.Core {
	return &coreSrv{
		DaoServant: s,
		ts:         ts,
		oss:        oss,
	}
}

func newCoreBinding() api.CoreBinding {
	return &coreBinding{
		UnimplementedCoreBinding: &api.UnimplementedCoreBinding{
			BindAny: base.BindAny,
		},
	}
}

func newCoreRender() api.CoreRender {
	return &coreRender{
		UnimplementedCoreRender: &api.UnimplementedCoreRender{
			RenderAny: base.RenderAny,
		},
	}
}
