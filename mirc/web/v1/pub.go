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
	Group `mir:"v1"`

	// Version 获取后台版本信息
	Version func(Get) web.VersionResp `mir:"/"`

	// Login 用户登录
	Login func(Post, web.LoginReq) web.LoginResp `mir:"/auth/login"`

	// Register 用户注册
	Register func(Post, web.RegisterReq) web.RegisterResp `mir:"/auth/register"`

	// GetCaptcha 获取验证码
	GetCaptcha func(Get) web.GetCaptchaResp `mir:"/captcha"`

	// SendCaptcha 发送验证码
	SendCaptcha func(Post, web.SendCaptchaReq) `mir:"/captcha"`

	// TweetDetail 获取动态详情
	TweetDetail func(Get, web.TweetDetailReq) web.TweetDetailResp `mir:"/post"`
}
