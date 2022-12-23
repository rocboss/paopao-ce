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
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// 	RequestingFriend 请求添加朋友
	RequestingFriend func(Post) `mir:"/friend/requesting"`

	// AddFriend 同意添加好友
	AddFriend func(Post) `mir:"/friend/add"`

	// RejectFriend 拒绝添加好友
	RejectFriend func(Post) `mir:"/friend/reject"`

	// 	DeleteFriend 删除好友
	DeleteFriend func(Post) `mir:"/friend/delete"`

	// GetContacts 获取好友列表
	GetContacts func(Get, web.GetContactsReq) web.GetContactsResp `mir:"/user/contacts"`
}
