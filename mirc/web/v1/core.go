package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Core))
}

// Core 核心服务，需要授权访问
type Core struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// SyncSearchIndex 同步索引
	SyncSearchIndex func(Get, web.SyncSearchIndexReq) `mir:"/sync/index"`

	// GetUserInfo 获取当前用户信息
	GetUserInfo func(Get, web.UserInfoReq) web.UserInfoResp `mir:"/user/info"`

	// GetUnreadMsgCount 获取当前用户未读消息数量
	GetUnreadMsgCount func(Get, web.GetUnreadMsgCountReq) web.GetUnreadMsgCountResp `mir:"/user/msgcount/unread"`

	// GetMessages 获取消息列表
	GetMessages func(Get, web.GetMessagesReq) web.GetMessagesResp `mir:"/user/messages"`

	// ReadMessage 标记消息已读
	ReadMessage func(Post, web.ReadMessageReq) `mir:"/user/message/read"`

	// SendUserWhisper 发送用户私信
	SendUserWhisper func(Post, web.SendWhisperReq) `mir:"/user/whisper"`

	// GetCollections 获取用户收藏列表
	GetCollections func(Get, web.GetCollectionsReq) web.GetCollectionsResp `mir:"/user/collections"`

	// GetStars 获取用户点赞列表
	GetStars func(Get, web.GetStarsReq) web.GetStarsResp `mir:"/user/stars"`

	// UserPhoneBind 绑定用户手机号
	UserPhoneBind func(Post, web.UserPhoneBindReq) `mir:"/user/phone"`

	// ChangePassword 修改密码
	ChangePassword func(Post, web.ChangePasswordReq) `mir:"/user/password"`

	// ChangeNickname 修改昵称
	ChangeNickname func(Post, web.ChangeNicknameReq) `mir:"/user/nickname"`

	// ChangeAvatar 修改头像
	ChangeAvatar func(Post, web.ChangeAvatarReq) `mir:"/user/avatar"`

	// SuggestUsers 检索用户
	SuggestUsers func(Get, web.SuggestUsersReq) web.SuggestUsersResp `mir:"/suggest/users"`

	// SuggestTags 检索标签
	SuggestTags func(Get, web.SuggestTagsReq) web.SuggestTagsResp `mir:"/suggest/tags"`

	// TweetStarStatus 获取动态点赞状态
	TweetStarStatus func(Get, web.TweetStarStatusReq) web.TweetStarStatusResp `mir:"/post/star"`

	// TweetCollectionStatus 获取动态收藏状态
	TweetCollectionStatus func(Get, web.TweetCollectionStatusReq) web.TweetCollectionStatusResp `mir:"/post/collection"`
}
