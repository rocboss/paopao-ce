package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/rocboss/paopao-ce/internal/model/v2/web"
)

func init() {
	Entry[Loose]()
}

// Loose 宽松授权的服务
type Loose struct {
	Chain `mir:"-"`
	Group `mir:"v2"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"posts" binding:"query"`

	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get, web.GetUserTweetsReq) web.GetUserTweetsResp `mir:"user/posts" binding:"query"`

	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get, web.GetUserProfileReq) web.GetUserProfileResp `mir:"user/profile" binding:"query"`

	// TopicList 获取话题列表
	TopicList func(Get, web.TopicListReq) web.TopicListResp `mir:"tags" binding:"query"`

	// TweetComments 获取动态评论
	TweetComments func(Get, web.TweetCommentsReq) web.TweetCommentsResp `mir:"post/comments"  binding:"query"`

	// TweetDetail 获取动态详情
	TweetDetail func(Get, web.TweetDetailReq) web.TweetDetailResp `mir:"post"  binding:"query"`
}
