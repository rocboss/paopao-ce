package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/middleware"
	"github.com/rocboss/paopao-ce/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

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

	// 支付宝回调
	r.POST("/alipay/notify", api.AlipayNotify)

	// 无鉴权路由组
	noAuthApi := r.Group("/")
	{
		// 获取广场流
		noAuthApi.GET("/posts", api.GetPostList)

		// 获取动态详情
		noAuthApi.GET("/post", api.GetPost)

		// 获取动态评论
		noAuthApi.GET("/post/comments", api.GetPostComments)

		// 获取话题列表
		noAuthApi.GET("/tags", api.GetPostTags)

		// 获取用户基本信息
		noAuthApi.GET("/user/profile", api.GetUserProfile)

		// 获取用户动态列表
		noAuthApi.GET("/user/posts", api.GetUserPosts)
	}

	// 鉴权路由组
	authApi := r.Group("/").Use(middleware.JWT())
	privApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Priv())
	{
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

		// 用户充值
		authApi.POST("/user/recharge", api.GetUserRechargeLink)

		// 用户充值
		authApi.GET("/user/recharge", api.GetUserRechargeResult)

		// 获取用户账单
		authApi.GET("/user/wallet/bills", api.GetUserWalletBills)

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

		// 发布动态评论
		privApi.POST("/post/comment", api.CreatePostComment)

		// 删除动态评论
		privApi.DELETE("/post/comment", api.DeletePostComment)

		// 发布评论回复
		privApi.POST("/post/comment/reply", api.CreatePostCommentReply)

		// 删除评论回复
		privApi.DELETE("/post/comment/reply", api.DeletePostCommentReply)

	}
	// 默认404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})
	// 默认405
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})
	return r
}
