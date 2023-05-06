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
	Chain `mir:"-"`
	Group `mir:"v1"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"/posts"`

	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get, web.GetUserTweetsReq) web.GetUserTweetsResp `mir:"/user/posts"`

	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get, web.GetUserProfileReq) web.GetUserProfileResp `mir:"/user/profile"`

	// TopicList 获取话题列表
	TopicList func(Get, web.TopicListReq) web.TopicListResp `mir:"/tags"`

	// TweetComments 获取动态评论
	TweetComments func(Get, web.TweetCommentsReq) web.TweetCommentsResp `mir:"/post/comments"`
}
