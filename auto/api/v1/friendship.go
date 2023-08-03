// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

type Friendship interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	GetContacts(*web.GetContactsReq) (*web.GetContactsResp, mir.Error)
	DeleteFriend(*web.DeleteFriendReq) mir.Error
	RejectFriend(*web.RejectFriendReq) mir.Error
	AddFriend(*web.AddFriendReq) mir.Error
	RequestingFriend(*web.RequestingFriendReq) mir.Error

	mustEmbedUnimplementedFriendshipServant()
}

// RegisterFriendshipServant register Friendship servant to gin
func RegisterFriendshipServant(e *gin.Engine, s Friendship) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/user/contacts", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.GetContactsReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.GetContacts(req)
		s.Render(c, resp, err)
	})
	router.Handle("POST", "/friend/delete", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.DeleteFriendReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.DeleteFriend(req))
	})
	router.Handle("POST", "/friend/reject", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.RejectFriendReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.RejectFriend(req))
	})
	router.Handle("POST", "/friend/add", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.AddFriendReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.AddFriend(req))
	})
	router.Handle("POST", "/friend/requesting", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.RequestingFriendReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.RequestingFriend(req))
	})
}

// UnimplementedFriendshipServant can be embedded to have forward compatible implementations.
type UnimplementedFriendshipServant struct{}

func (UnimplementedFriendshipServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedFriendshipServant) GetContacts(req *web.GetContactsReq) (*web.GetContactsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFriendshipServant) DeleteFriend(req *web.DeleteFriendReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFriendshipServant) RejectFriend(req *web.RejectFriendReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFriendshipServant) AddFriend(req *web.AddFriendReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFriendshipServant) RequestingFriend(req *web.RequestingFriendReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFriendshipServant) mustEmbedUnimplementedFriendshipServant() {}
