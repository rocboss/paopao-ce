package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebFriendship))
}

// // routeFriendship register Friendship feature releated routes
// func routeFriendship(authApi gin.IRoutes) {
// 	// 请求添加朋友
// 	authApi.POST("/friend/requesting", api.RequestingFriend)

// 	// 同意添加好友
// 	authApi.POST("/friend/add", api.AddFriend)

// 	// 拒绝添加好友
// 	authApi.POST("/friend/reject", api.RejectFriend)

// 	// 删除好友
// 	authApi.POST("/friend/delete", api.DeleteFriend)

// 	// 获取好友列表
// 	authApi.GET("/user/contacts", api.GetContacts)
// }

type WebFriendship struct {
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
