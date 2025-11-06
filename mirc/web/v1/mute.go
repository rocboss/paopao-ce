package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Muteship 屏蔽模式 服务
type Muteship struct {
	Schema `mir:"v1"`

	// MuteUser 屏蔽用户
	MuteUser func(Post, web.MuteUserReq) `mir:"user/mute"`

	// UnmuteUser 取消屏蔽用户
	UnmuteUser func(Post, web.UnmuteUserReq) `mir:"user/unmute"`

	// ListMutes 获取用户的屏蔽列表
	ListMutes func(Get, web.ListMutesReq) web.ListMutesResp `mir:"user/mutes"`
}
