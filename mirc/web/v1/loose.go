package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebLoose))
}

// // 宽松鉴权路由组
// looseApi := r.Group("/").Use(chain.JwtLoose())
// {
// 	// 获取广场流
// 	looseApi.GET("/posts", api.GetPostList)

// 	// 获取用户动态列表
// 	looseApi.GET("/user/posts", api.GetUserPosts)

// 	// 获取用户基本信息
// 	looseApi.GET("/user/profile", api.GetUserProfile)
// }

type WebLoose struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"/Posts"`
	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get) `mir:"/user/posts"`
	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get) `mir:"/user/profile"`
}
