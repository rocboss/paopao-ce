package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	Entry[Relax]()
}

// Relax 放宽授权的服务
type Relax struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// GetUnreadMsgCount 获取当前用户未读消息数量
	GetUnreadMsgCount func(Get, Chain, web.GetUnreadMsgCountReq) web.GetUnreadMsgCountResp `mir:"/user/msgcount/unread"`
}
