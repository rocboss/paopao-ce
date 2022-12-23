package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebAdmin))
}

type WebAdmin struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// ChangeUserStatus 管理·禁言/解封用户
	ChangeUserStatus func(Post, web.ChangeUserStatusReq) `mir:"/admin/user/status"`
}
