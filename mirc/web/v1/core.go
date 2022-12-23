package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebCore))
}

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