package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Admin))
}

// Admin 运维相关服务
type Admin struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// ChangeUserStatus 管理·禁言/解封用户
	ChangeUserStatus func(Post, web.ChangeUserStatusReq) `mir:"/admin/user/status"`
}
