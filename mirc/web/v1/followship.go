package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Followship 关注者模式 服务
type Followship struct {
	Schema `mir:"v1,chain"`

	// FollowUser 关注用户
	FollowUser func(Post, web.FollowUserReq) `mir:"user/follow"`

	// UnfollowUser  取消关注用户
	UnfollowUser func(Post, web.UnfollowUserReq) `mir:"user/unfollow"`

	// ListFollows 获取用户的关注列表
	ListFollows func(Get, web.ListFollowsReq) web.ListFollowsResp `mir:"user/follows"`

	// ListFollowings 获取用户的追随者列表
	ListFollowings func(Get, web.ListFollowingsReq) web.ListFollowingsResp `mir:"user/followings"`
}
