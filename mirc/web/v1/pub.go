package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Pub))
}

// Pub 不用授权的公开服务
type Pub struct {
	Group Group `mir:"v1"`

	// Version 获取后台版本信息
	Version func(Get) `mir:"/"`

	// Login 用户登录
	Login func(Post) `mir:"/auth/login"`

	// Register 用户注册
	Register func(Post) `mir:"/auth/register"`

	// GetCaptcha 获取验证码
	GetCaptcha func(Get) `mir:"/captcha"`

	// PostCaptcha 发送验证码
	SendCaptcha func(Post) `mir:"/captcha"`

	// TweetDetail 获取动态详情
	TweetDetail func(Get, web.TweetDetailReq) web.TweetDetailResp `mir:"/post"`

	// TweetComments 获取动态评论
	TweetComments func(Get) `mir:"/post/comments"`

	// TopicList 获取话题列表
	TopicList func(Get) `mir:"/tags"`
}
