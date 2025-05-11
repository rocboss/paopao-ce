package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Trends 动态相关 服务
type Trends struct {
	Schema `mir:"v1,chain"`

	// GetIndexTrends 获取广场页面动态条栏的索引item
	GetIndexTrends func(Get, web.GetIndexTrendsReq) web.GetIndexTrendsResp `mir:"trends/index"`
}
