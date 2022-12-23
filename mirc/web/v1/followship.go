package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebFollowship))
}

type WebFollowship struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// AddFollowing 添加关注
	AddFollowing func(Post) `mir:"/following/add"`

	// DeleteFollowing  取消关注
	DeleteFollowing func(Post) `mir:"/following/delete"`

	// ListFollowings 获取用户的关注列表
	ListFollowings func(Get) `mir:"/following/list"`

	// ListFollowers 获取用户的追随者列表
	ListFollowers func(Get, web.ListFollowersReq) web.ListFollowersResp `mir:"/follower/list"`
}