// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Friendship        = (*friendshipSrv)(nil)
	_ api.FriendshipBinding = (*friendshipBinding)(nil)
	_ api.FriendshipRender  = (*friendshipRender)(nil)
)

type friendshipSrv struct {
	api.UnimplementedFriendshipServant
	*base.DaoServant
}

type friendshipBinding struct {
	*api.UnimplementedFriendshipBinding
}

type friendshipRender struct {
	*api.UnimplementedFriendshipRender
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
		return nil, _errGetContactsFailed
	}
	resp := base.PageRespFrom(res.Contacts, req.Page, req.PageSize, res.Total)
	return (*web.GetContactsResp)(resp), nil
}

func (s *friendshipSrv) DeleteFriend(req *web.DeleteFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User != nil && req.User.ID == req.UserId {
		return _errNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return _errNotExistFriendId
	}
	if err := s.Ds.DeleteFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.DeleteFriend err: %s", err)
		return _errDeleteFriendFailed
	}
	return nil
}

func (s *friendshipSrv) RejectFriend(req *web.RejectFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return _errNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return _errNotExistFriendId
	}
	if err := s.Ds.RejectFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.RejectFriend err: %s", err)
		return _errRejectFriendFailed
	}
	return nil
}

func (s *friendshipSrv) AddFriend(req *web.AddFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return _errNoActionToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return _errNotExistFriendId
	}
	if err := s.Ds.AddFriend(req.User.ID, req.UserId); err != nil {
		logrus.Errorf("Ds.AddFriend err: %s", err)
		return _errAddFriendFailed
	}
	return nil
}

func (s *friendshipSrv) RequestingFriend(req *web.RequestingFriendReq) mir.Error {
	if req.User == nil {
		return xerror.ServerError
	}
	if req.User.ID == req.UserId {
		return _errNoRequestingFriendToSelf
	}
	if _, err := s.Ds.GetUserByID(req.UserId); err != nil {
		return _errNotExistFriendId
	}
	if err := s.Ds.RequestingFriend(req.User.ID, req.UserId, req.Greetings); err != nil {
		logrus.Errorf("Ds.RequestingFriend err: %s", err)
		return _errSendRequestingFriendFailed
	}
	return nil
}

func newFriendshipSrv(s *base.DaoServant) api.Friendship {
	return &friendshipSrv{
		DaoServant: s,
	}
}

func newFriendshipBinding() api.FriendshipBinding {
	return &friendshipBinding{
		UnimplementedFriendshipBinding: &api.UnimplementedFriendshipBinding{
			BindAny: base.NewBindAnyFn(),
		},
	}
}

func newFriendshipRender() api.FriendshipRender {
	return &friendshipRender{
		UnimplementedFriendshipRender: &api.UnimplementedFriendshipRender{
			RenderAny: base.RenderAny,
		},
	}
}
