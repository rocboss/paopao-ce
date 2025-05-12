package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Friendship 好友模式 服务
type Friendship struct {
	Schema `mir:"v1,chain"`

	// 	RequestingFriend 请求添加朋友
	RequestingFriend func(Post, web.RequestingFriendReq) `mir:"friend/requesting" binding:"json"`

	// AddFriend 同意添加好友
	AddFriend func(Post, web.AddFriendReq) `mir:"friend/add" binding:"json"`

	// RejectFriend 拒绝添加好友
	RejectFriend func(Post, web.RejectFriendReq) `mir:"friend/reject" binding:"json"`

	// 	DeleteFriend 删除好友
	DeleteFriend func(Post, web.DeleteFriendReq) `mir:"friend/delete" binding:"json"`

	// GetContacts 获取好友列表
	GetContacts func(Get, web.GetContactsReq) web.GetContactsResp `mir:"user/contacts"`
}
