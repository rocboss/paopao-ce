package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Relax 放宽授权的服务
type Relax struct {
	Schema `mir:"v1,chain"`

	// GetUnreadMsgCount 获取当前用户未读消息数量
	GetUnreadMsgCount func(Get, Chain, web.GetUnreadMsgCountReq) web.GetUnreadMsgCountResp `mir:"user/msgcount/unread"`
}
