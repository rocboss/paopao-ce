package routers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/middleware"
	"github.com/rocboss/paopao-ce/internal/routers/api"
	"github.com/sirupsen/logrus"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// 按需注册 docs、静态资源、LocalOSS 路由
	{
		registerDocs(e)
		registerStatick(e)
		routeLocalOSS(e)
	}

	// v1 group api
	r := e.Group("/v1")

	// 获取version
	r.GET("/", api.Version)

	// 用户登录
	r.POST("/auth/login", api.Login)

	// 用户注册
	r.POST("/auth/register", api.Register)

	// 获取验证码
	r.GET("/captcha", api.GetCaptcha)

	// 发送验证码
	r.POST("/captcha", api.PostCaptcha)

	// 无鉴权路由组
	noAuthApi := r.Group("/")
	{
		// 获取动态详情
		noAuthApi.GET("/post", api.GetPost)

		// 获取动态评论
		noAuthApi.GET("/post/comments", api.GetPostComments)

		// 获取话题列表
		noAuthApi.GET("/tags", api.GetPostTags)

		// 获取用户基本信息
		noAuthApi.GET("/user/profile", api.GetUserProfile)
	}

	// 宽松鉴权路由组
	looseApi := r.Group("/").Use(middleware.JwtLoose())
	{
		// 获取广场流
		looseApi.GET("/posts", api.GetPostList)

		// 获取用户动态列表
		looseApi.GET("/user/posts", api.GetUserPosts)
	}

	// 鉴权路由组
	authApi := r.Group("/").Use(middleware.JWT())
	privApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Priv())
	adminApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Admin())

	// 核心路由注册
	routeCore(authApi, privApi, adminApi)

	// 支付宝路由注册
	routeAlipay(r, authApi)

	// Relationship相关路由注册
	routeRelationship(authApi)

	// 默认404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	return e
}

func routeCore(authApi gin.IRoutes, privApi gin.IRoutes, adminApi gin.IRoutes) {
	// 同步索引
	authApi.GET("/sync/index", api.SyncSearchIndex)

	// 获取当前用户信息
	authApi.GET("/user/info", api.GetUserInfo)

	// 获取当前用户未读消息数量
	authApi.GET("/user/msgcount/unread", api.GetUnreadMsgCount)

	// 获取消息列表
	authApi.GET("/user/messages", api.GetMessages)

	// 标记消息已读
	authApi.POST("/user/message/read", api.ReadMessage)

	// 发送用户私信
	authApi.POST("/user/whisper", api.SendUserWhisper)

	// 获取用户收藏列表
	authApi.GET("/user/collections", api.GetUserCollections)

	// 获取用户点赞列表
	authApi.GET("/user/stars", api.GetUserStars)

	// 绑定用户手机号
	authApi.POST("/user/phone", api.BindUserPhone)

	// 修改密码
	authApi.POST("/user/password", api.ChangeUserPassword)

	// 修改昵称
	authApi.POST("/user/nickname", api.ChangeNickname)

	// 修改头像
	authApi.POST("/user/avatar", api.ChangeAvatar)

	// 检索用户
	authApi.GET("/suggest/users", api.GetSuggestUsers)

	// 检索标签
	authApi.GET("/suggest/tags", api.GetSuggestTags)

	// 上传资源
	privApi.POST("/attachment", api.UploadAttachment)

	// 下载资源预检
	privApi.GET("/attachment/precheck", api.DownloadAttachmentPrecheck)

	// 下载资源
	privApi.GET("/attachment", api.DownloadAttachment)

	// 发布动态
	privApi.POST("/post", api.CreatePost)

	// 删除动态
	privApi.DELETE("/post", api.DeletePost)

	// 获取动态点赞状态
	authApi.GET("/post/star", api.GetPostStar)

	// 动态点赞操作
	privApi.POST("/post/star", api.PostStar)

	// 获取动态收藏状态
	authApi.GET("/post/collection", api.GetPostCollection)

	// 动态收藏操作
	privApi.POST("/post/collection", api.PostCollection)

	// 锁定动态
	privApi.POST("/post/lock", api.LockPost)

	// 置顶动态
	privApi.POST("/post/stick", api.StickPost)

	// 修改动态可见度
	privApi.POST("/post/visibility", api.VisiblePost)

	// 发布动态评论
	privApi.POST("/post/comment", api.CreatePostComment)

	// 删除动态评论
	privApi.DELETE("/post/comment", api.DeletePostComment)

	// 发布评论回复
	privApi.POST("/post/comment/reply", api.CreatePostCommentReply)

	// 删除评论回复
	privApi.DELETE("/post/comment/reply", api.DeletePostCommentReply)

	// 管理·禁言/解封用户
	adminApi.POST("/admin/user/status", api.ChangeUserStatus)
}

// routeLocalOSS register LocalOSS route if needed
func routeLocalOSS(e *gin.Engine) {
	if !conf.CfgIf("LocalOSS") {
		return
	}

	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("get localOSS save path err: %v", err)
	}
	e.Static("/oss", savePath)

	logrus.Infof("register LocalOSS route in /oss on save path: %s", savePath)
}

// routeAlipay register Alipay feature releated route if needed
func routeAlipay(public gin.IRoutes, authApi gin.IRoutes) {
	if !conf.CfgIf("Alipay") {
		return
	}

	// 支付宝回调
	public.POST("/alipay/notify", api.AlipayNotify)

	// 用户充值
	authApi.POST("/user/recharge", api.GetUserRechargeLink)

	// 获取钱包余额
	authApi.GET("/user/recharge", api.GetUserRechargeResult)

	// 获取用户账单
	authApi.GET("/user/wallet/bills", api.GetUserWalletBills)
}

// routeRelationship register Relationship releated routes
func routeRelationship(authApi gin.IRoutes) {
	if conf.CfgIf("Friendship") {
		routeFriendship(authApi)
	} else if conf.CfgIf("Followship") {
		routeFollowship(authApi)
	} else {
		// 暂时默认使用好友模式
		// TODO: 后期提供一种无关系模式(既不是好友模式也不是关注者模式)作为默认的关系模式
		routeFriendship(authApi)
	}

}

// routeFriendship register Friendship feature releated routes
func routeFriendship(authApi gin.IRoutes) {
	// 请求添加朋友
	authApi.POST("/friend/requesting", api.RequestingFriend)

	// 同意添加好友
	authApi.POST("/friend/add", api.AddFriend)

	// 拒绝添加好友
	authApi.POST("/friend/reject", api.RejectFriend)

	// 删除好友
	authApi.POST("/friend/delete", api.DeleteFriend)

	// 获取好友列表
	authApi.GET("/user/contacts", api.GetContacts)
}

// routeFollowship register Followship feature releated routes
func routeFollowship(authApi gin.IRoutes) {
	// TODO: register followship routes
}
