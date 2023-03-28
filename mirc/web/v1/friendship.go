package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Friendship))
}

// Friendship 好友模式 服务
type Friendship struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// 	RequestingFriend 请求添加朋友
	RequestingFriend func(Post, web.RequestingFriendReq) `mir:"/friend/requesting"`

	// AddFriend 同意添加好友
	AddFriend func(Post, web.AddFriendReq) `mir:"/friend/add"`

	// RejectFriend 拒绝添加好友
	RejectFriend func(Post, web.RejectFriendReq) `mir:"/friend/reject"`

	// 	DeleteFriend 删除好友
	DeleteFriend func(Post, web.DeleteFriendReq) `mir:"/friend/delete"`

	// GetContacts 获取好友列表
	GetContacts func(Get, web.GetContactsReq) web.GetContactsResp `mir:"/user/contacts"`
}
