// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	"time"
	"unicode/utf8"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

const (
	// _MaxWhisperNumDaily 当日单用户私信总数限制（TODO 配置化、积分兑换等）
	_MaxWhisperNumDaily = 20
	_MaxCaptchaTimes    = 2
)

var (
	_ api.Core        = (*coreSrv)(nil)
	_ api.CoreBinding = (*coreBinding)(nil)
	_ api.CoreRender  = (*coreRender)(nil)
)

type coreSrv struct {
	api.UnimplementedCoreServant
	*base.DaoServant

	oss core.ObjectStorageService
}

type coreBinding struct {
	*api.UnimplementedCoreBinding
}

type coreRender struct {
	*api.UnimplementedCoreRender
}

func (b *coreBinding) BindGetUserInfo(c *gin.Context) (*web.UserInfoReq, mir.Error) {
	username, exist := base.UserNameFrom(c)
	if !exist {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	return &web.UserInfoReq{
		Username: username,
	}, nil
}

func (b *coreBinding) BindGetMessages(c *gin.Context) (*web.GetMessagesReq, mir.Error) {
	v, err := web.BasePageReqFrom(c)
	return (*web.GetMessagesReq)(v), err
}

func (b *coreBinding) BindGetCollections(c *gin.Context) (*web.GetCollectionsReq, mir.Error) {
	v, err := web.BasePageReqFrom(c)
	return (*web.GetCollectionsReq)(v), err
}

func (b *coreBinding) BindGetStars(c *gin.Context) (*web.GetStarsReq, mir.Error) {
	v, err := web.BasePageReqFrom(c)
	return (*web.GetStarsReq)(v), err
}

func (b *coreBinding) BindSuggestTags(c *gin.Context) (*web.SuggestTagsReq, mir.Error) {
	return &web.SuggestTagsReq{
		Keyword: c.Query("k"),
	}, nil
}

func (b *coreBinding) BindSuggestUsers(c *gin.Context) (*web.SuggestUsersReq, mir.Error) {
	return &web.SuggestUsersReq{
		Keyword: c.Query("k"),
	}, nil
}

func (b *coreBinding) BindTweetCollectionStatus(c *gin.Context) (*web.TweetCollectionStatusReq, mir.Error) {
	UserId, exist := base.UserIdFrom(c)
	if !exist {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	return &web.TweetCollectionStatusReq{
		SimpleInfo: web.SimpleInfo{
			Uid: UserId,
		},
		TweetId: convert.StrTo(c.Query("id")).MustInt64(),
	}, nil
}

func (b *coreBinding) BindTweetStarStatus(c *gin.Context) (*web.TweetStarStatusReq, mir.Error) {
	UserId, exist := base.UserIdFrom(c)
	if !exist {
		return nil, xerror.UnauthorizedAuthNotExist
	}
	return &web.TweetStarStatusReq{
		SimpleInfo: web.SimpleInfo{
			Uid: UserId,
		},
		TweetId: convert.StrTo(c.Query("id")).MustInt64(),
	}, nil
}

func (s *coreSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *coreSrv) SyncSearchIndex(req *web.SyncSearchIndexReq) mir.Error {
	if req.User != nil && req.User.IsAdmin {
		go s.PushPostsToSearch(context.Background())
	} else {
		logrus.Warnf("sync search index need admin permision user: %#v", req.User)
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

func (s *coreSrv) GetMessages(req *web.GetMessagesReq) (*web.GetMessagesResp, mir.Error) {
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
	resp := base.PageRespFrom(messages, req.Page, req.PageSize, totalRows)
	return (*web.GetMessagesResp)(resp), nil
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

func (s *coreSrv) SendUserWhisper(req *web.SendWhisperReq) mir.Error {
	// 不允许发送私信给自己
	if req.Uid == req.UserID {
		return _errNoWhisperToSelf
	}

	// 今日频次限制
	ctx := context.Background()
	if count, _ := s.Redis.GetCountWhisper(ctx, req.Uid); count >= _MaxWhisperNumDaily {
		return _errTooManyWhisperNum
	}

	// 创建私信
	_, err := s.Ds.CreateMessage(&core.Message{
		SenderUserID:   req.Uid,
		ReceiverUserID: req.UserID,
		Type:           core.MsgTypeWhisper,
		Brief:          "给你发送新私信了",
		Content:        req.Content,
	})
	if err != nil {
		logrus.Errorf("Ds.CreateWhisper err: %s", err)
		return _errSendWhisperFailed
	}

	// 写入当日（自然日）计数缓存
	s.Redis.IncrCountWhisper(ctx, req.Uid)

	return nil
}

func (s *coreSrv) GetCollections(req *web.GetCollectionsReq) (*web.GetCollectionsResp, mir.Error) {
	collections, err := s.Ds.GetUserPostCollections(req.UserId, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("Ds.GetUserPostCollections err: %s", err)
		return nil, _errGetCollectionsFailed
	}
	totalRows, err := s.Ds.GetUserPostCollectionCount(req.UserId)
	if err != nil {
		logrus.Errorf("Ds.GetUserPostCollectionCount err: %s", err)
		return nil, _errGetCollectionsFailed
	}

	var posts []*core.Post
	for _, collection := range collections {
		posts = append(posts, collection.Post)
	}
	postsFormated, err := s.Ds.MergePosts(posts)
	if err != nil {
		logrus.Errorf("Ds.MergePosts err: %s", err)
		return nil, _errGetCollectionsFailed
	}
	resp := base.PageRespFrom(postsFormated, req.Page, req.PageSize, totalRows)

	return (*web.GetCollectionsResp)(resp), nil
}

func (s *coreSrv) UserPhoneBind(req *web.UserPhoneBindReq) mir.Error {
	// 手机重复性检查
	u, err := s.Ds.GetUserByPhone(req.Phone)
	if err == nil && u.Model != nil && u.ID != 0 && u.ID != req.User.ID {
		return _errExistedUserPhone
	}

	// 如果禁止phone verify 则允许通过任意验证码
	if _enablePhoneVerify {
		c, err := s.Ds.GetLatestPhoneCaptcha(req.Phone)
		if err != nil {
			return _errErrorPhoneCaptcha
		}
		if c.Captcha != req.Captcha {
			return _errErrorPhoneCaptcha
		}
		if c.ExpiredOn < time.Now().Unix() {
			return _errErrorPhoneCaptcha
		}
		if c.UseTimes >= _MaxCaptchaTimes {
			return _errMaxPhoneCaptchaUseTimes
		}
		// 更新检测次数
		s.Ds.UsePhoneCaptcha(c)
	}

	// 执行绑定
	user := req.User
	user.Phone = req.Phone
	if err := s.Ds.UpdateUser(user); err != nil {
		// TODO: 优化错误处理逻辑，失败后上面的逻辑也应该回退
		logrus.Errorf("Ds.UpdateUser err: %s", err)
		return xerror.ServerError
	}
	return nil
}

func (s *coreSrv) GetStars(req *web.GetStarsReq) (*web.GetStarsResp, mir.Error) {
	stars, err := s.Ds.GetUserPostStars(req.UserId, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("Ds.GetUserPostStars err: %s", err)
		return nil, _errGetStarsFailed
	}
	totalRows, err := s.Ds.GetUserPostStarCount(req.UserId)
	if err != nil {
		logrus.Errorf("Ds.GetUserPostStars err: %s", err)
		return nil, _errGetStarsFailed
	}

	var posts []*core.Post
	for _, star := range stars {
		posts = append(posts, star.Post)
	}
	postsFormated, err := s.Ds.MergePosts(posts)
	if err != nil {
		logrus.Errorf("Ds.MergePosts err: %s", err)
		return nil, _errGetStarsFailed
	}
	resp := base.PageRespFrom(postsFormated, req.Page, req.PageSize, totalRows)

	return (*web.GetStarsResp)(resp), nil
}

func (s *coreSrv) ChangePassword(req *web.ChangePasswordReq) mir.Error {
	// 密码检查
	if err := checkPassword(req.Password); err != nil {
		return err
	}
	// 旧密码校验
	user := req.User
	if !validPassword(user.Password, req.OldPassword, req.User.Salt) {
		return _errErrorOldPassword
	}
	// 更新入库
	user.Password, user.Salt = encryptPasswordAndSalt(req.Password)
	if err := s.Ds.UpdateUser(user); err != nil {
		logrus.Errorf("Ds.UpdateUser err: %s", err)
		return xerror.ServerError
	}
	return nil
}

func (s *coreSrv) SuggestTags(req *web.SuggestTagsReq) (*web.SuggestTagsResp, mir.Error) {
	tags, err := s.Ds.GetTagsByKeyword(req.Keyword)
	if err != nil {
		logrus.Errorf("Ds.GetTagsByKeyword err: %s", err)
		return nil, xerror.ServerError
	}
	resp := &web.SuggestTagsResp{}
	for _, t := range tags {
		resp.Suggests = append(resp.Suggests, t.Tag)
	}
	return resp, nil
}

func (s *coreSrv) SuggestUsers(req *web.SuggestUsersReq) (*web.SuggestUsersResp, mir.Error) {
	users, err := s.Ds.GetUsersByKeyword(req.Keyword)
	if err != nil {
		logrus.Errorf("Ds.GetUsersByKeyword err: %s", err)
		return nil, xerror.ServerError
	}
	resp := &web.SuggestUsersResp{}
	for _, user := range users {
		resp.Suggests = append(resp.Suggests, user.Username)
	}
	return resp, nil
}

func (s *coreSrv) ChangeNickname(req *web.ChangeNicknameReq) mir.Error {
	if utf8.RuneCountInString(req.Nickname) < 2 || utf8.RuneCountInString(req.Nickname) > 12 {
		return _errNicknameLengthLimit
	}
	user := req.User
	user.Nickname = req.Nickname
	if err := s.Ds.UpdateUser(user); err != nil {
		logrus.Errorf("Ds.UpdateUser err: %s", err)
		return xerror.ServerError
	}
	return nil
}

func (s *coreSrv) ChangeAvatar(req *web.ChangeAvatarReq) (xerr mir.Error) {
	defer func() {
		if xerr != nil {
			deleteOssObjects(s.oss, []string{req.Avatar})
		}
	}()

	if err := s.Ds.CheckAttachment(req.Avatar); err != nil {
		logrus.Errorf("Ds.CheckAttachment failed: %s", err)
		return xerror.InvalidParams
	}
	if err := s.oss.PersistObject(s.oss.ObjectKey(req.Avatar)); err != nil {
		logrus.Errorf("Ds.ChangeUserAvatar persist object failed: %s", err)
		return xerror.ServerError
	}
	user := req.User
	user.Avatar = req.Avatar
	if err := s.Ds.UpdateUser(user); err != nil {
		logrus.Errorf("Ds.UpdateUser failed: %s", err)
		return xerror.ServerError
	}
	return nil
}

func (s *coreSrv) TweetCollectionStatus(req *web.TweetCollectionStatusReq) (*web.TweetCollectionStatusResp, mir.Error) {
	resp := &web.TweetCollectionStatusResp{
		Status: true,
	}
	if _, err := s.Ds.GetUserPostCollection(req.TweetId, req.Uid); err != nil {
		resp.Status = false
		return resp, nil
	}
	return resp, nil
}

func (s *coreSrv) TweetStarStatus(req *web.TweetStarStatusReq) (*web.TweetStarStatusResp, mir.Error) {
	resp := &web.TweetStarStatusResp{
		Status: true,
	}
	if _, err := s.Ds.GetUserPostStar(req.TweetId, req.Uid); err != nil {
		resp.Status = false
		return resp, nil
	}
	return resp, nil
}

func newCoreSrv(s *base.DaoServant, oss core.ObjectStorageService) api.Core {
	return &coreSrv{
		DaoServant: s,
		oss:        oss,
	}
}

func newCoreBinding() api.CoreBinding {
	return &coreBinding{
		UnimplementedCoreBinding: &api.UnimplementedCoreBinding{
			BindAny: base.NewBindAnyFn(),
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
