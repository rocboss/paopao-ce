package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Blockship 拉黑模式 服务
type Blockship struct {
	Schema `mir:"v1"`

	// BlockUser 拉黑用户
	BlockUser func(Post, web.BlockUserReq) `mir:"user/block"`

	// UnblockUser 取消拉黑用户
	UnblockUser func(Post, web.UnblockUserReq) `mir:"user/unblock"`

	// ListBlocks 获取用户的拉黑列表
	ListBlocks func(Get, web.ListBlocksReq) web.ListBlocksResp `mir:"user/blocks"`
}
