package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebCore))
}

// // 同步索引
// authApi.GET("/sync/index", api.SyncSearchIndex)

// // 获取当前用户信息
// authApi.GET("/user/info", api.GetUserInfo)

// // 获取当前用户未读消息数量
// authApi.GET("/user/msgcount/unread", api.GetUnreadMsgCount)

// // 获取消息列表
// authApi.GET("/user/messages", api.GetMessages)

// // 标记消息已读
// authApi.POST("/user/message/read", api.ReadMessage)

// // 发送用户私信
// authApi.POST("/user/whisper", api.SendUserWhisper)

// // 获取用户收藏列表
// authApi.GET("/user/collections", api.GetUserCollections)

// // 获取用户点赞列表
// authApi.GET("/user/stars", api.GetUserStars)

// // 绑定用户手机号
// authApi.POST("/user/phone", api.BindUserPhone)

// // 修改密码
// authApi.POST("/user/password", api.ChangeUserPassword)

// // 修改昵称
// authApi.POST("/user/nickname", api.ChangeNickname)

// // 修改头像
// authApi.POST("/user/avatar", api.ChangeAvatar)

// // 检索用户
// authApi.GET("/suggest/users", api.GetSuggestUsers)

// // 检索标签
// authApi.GET("/suggest/tags", api.GetSuggestTags)

// // 获取动态点赞状态
// authApi.GET("/post/star", api.GetPostStar)

// // 获取动态收藏状态
// authApi.GET("/post/collection", api.GetPostCollection)

type WebCore struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// SyncSearchIndex 同步索引
	SyncSearchIndex func(Get) `mir:"/sync/index"`
	// GetUserInfo 获取当前用户信息
	GetUserInfo func(Get, web.UserInfoReq) web.UserInfoResp `mir:"/user/info"`
	// GetUnreadMsgCount 获取当前用户未读消息数量
	GetUnreadMsgCount func(Get) `mir:"/user/msgcount/unread"`
	// GetMessages 获取消息列表
	GetMessages func(Get) `mir:"/user/messages"`
	// ReadMessage 标记消息已读
	ReadMessage func(Post) `mir:"/user/message/read"`
	// SendUserWhisper 发送用户私信
	SendUserWhisper func(Post) `mir:"/user/whisper"`
	// GetCollections 获取用户收藏列表
	GetCollections func(Get) `mir:"/user/collections"`
	// GetStars 获取用户点赞列表
	GetStars func(Get) `mir:"/user/stars"`
	// UserPhoneBind 绑定用户手机号
	UserPhoneBind func(Post) `mir:"/user/phone"`
	// ChangePassword 修改密码
	ChangePassword func(Post) `mir:"/user/password"`
	// ChangeNickname 修改昵称
	ChangeNickname func(Post) `mir:"/user/nickname"`
	// ChangeAvatar 修改头像
	ChangeAvatar func(Post, web.ChangeAvatarReq) `mir:"/user/avatar"`
	// SuggestUsers 检索用户
	SuggestUsers func(Get) `mir:"/suggest/users"`
	// SuggestTags 检索标签
	SuggestTags func(Get) `mir:"/suggest/tags"`
	// TweetStarStatus 获取动态点赞状态
	TweetStarStatus func(Get) `mir:"/post/star"`
	// TweetCollectionStatus 获取动态收藏状态
	TweetCollectionStatus func(Get) `mir:"/post/collection"`
}
