package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	Entry[Trends]()
}

// Trends 动态相关 服务
type Trends struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// GetIndexTrends 获取广场页面动态条栏的索引item
	GetIndexTrends func(Get, web.GetIndexTrendsReq) web.GetIndexTrendsResp `mir:"/trends/index"`
}
