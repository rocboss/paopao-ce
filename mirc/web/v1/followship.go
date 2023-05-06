package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Followship))
}

// Followship 关注者模式 服务
type Followship struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// AddFollowing 添加关注
	AddFollowing func(Post, web.AddFollowingReq) `mir:"/following/add"`

	// DeleteFollowing  取消关注
	DeleteFollowing func(Post, web.DeleteFollowingReq) `mir:"/following/delete"`

	// ListFollowings 获取用户的关注列表
	ListFollowings func(Get, web.ListFollowingsReq) web.ListFollowingsResp `mir:"/following/list"`

	// ListFollowers 获取用户的追随者列表
	ListFollowers func(Get, web.ListFollowersReq) web.ListFollowersResp `mir:"/follower/list"`
}
