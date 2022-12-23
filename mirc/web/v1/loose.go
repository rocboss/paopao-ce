package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Loose))
}

// Loose 宽松授权的服务
type Loose struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"/posts"`

	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get) `mir:"/user/posts"`

	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get) `mir:"/user/profile"`
}
