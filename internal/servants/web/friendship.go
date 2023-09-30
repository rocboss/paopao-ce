// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Friendship = (*friendshipSrv)(nil)
)

type friendshipSrv struct {
	api.UnimplementedFriendshipServant
	*base.DaoServant
}

func (s *friendshipSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *friendshipSrv) GetContacts(req *web.GetContactsReq) (*web.GetContactsResp, mir.Error) {
	if req.User == nil {
		return nil, xerror.ServerError
	}
	res, err := s.Ds.GetContacts(req.User.ID, (req.Page-1)*req.PageSize, req.PageSize)
	if err != nil {
		logrus.Errorf("service.GetContacts err: %s", err)
		return nil, web.ErrGetContactsFailed
	}
	resp := base.PageRespFrom(res.Contacts, req.Page, req.PageSize, res.Total)
	return (*web.GetContactsResp)(resp), nil
}

func (s *friendshipSrv) DeleteFriend(req *web.DeleteFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User != nil && req.User.ID == req.UserId {
		return web.ErrNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return web.ErrNotExistFriendId
	}
	if err := s.Ds.DeleteFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.DeleteFriend err: %s", err)
		return web.ErrDeleteFriendFailed
	}
	// 触发用户关系缓存更新事件
	// cache.OnCacheMyFriendIdsEvent(s.Ds, req.User.ID, req.UserId)
	onTrendsActionEvent(_trendsActionDeleteFriend, req.User.ID, req.UserId)
	return nil
}

func (s *friendshipSrv) RejectFriend(req *web.RejectFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return web.ErrNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return web.ErrNotExistFriendId
	}
	if err := s.Ds.RejectFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.RejectFriend err: %s", err)
		return web.ErrRejectFriendFailed
	}
	return nil
}

func (s *friendshipSrv) AddFriend(req *web.AddFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return web.ErrNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return web.ErrNotExistFriendId
	}
	if err := s.Ds.AddFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.AddFriend err: %s", err)
		return web.ErrAddFriendFailed
	}
	// 触发用户关系缓存更新事件
	// cache.OnCacheMyFriendIdsEvent(s.Ds, req.User.ID, req.UserId)
	onTrendsActionEvent(_trendsActionAddFriend, req.User.ID, req.UserId)
	return nil
}

func (s *friendshipSrv) RequestingFriend(req *web.RequestingFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return web.ErrNoRequestingFriendToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return web.ErrNotExistFriendId
	}
	if err := s.Ds.RequestingFriend(req.User.ID, req.UserId, req.Greetings); err != nil {
		logrus.Errorf("Ds.RequestingFriend err: %s", err)
		return web.ErrSendRequestingFriendFailed
	}
	return nil
}

func newFriendshipSrv(s *base.DaoServant) api.Friendship {
	return &friendshipSrv{
		DaoServant: s,
	}
}
