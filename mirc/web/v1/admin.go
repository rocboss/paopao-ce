package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Admin 运维相关服务
type Admin struct {
	Schema `mir:"v1,chain"`

	// ChangeUserStatus 管理·禁言/解封用户
	ChangeUserStatus func(Post, web.ChangeUserStatusReq)         `mir:"admin/user/status"`
	SiteInfo         func(Get, web.SiteInfoReq) web.SiteInfoResp `mir:"admin/site/status"`
}
