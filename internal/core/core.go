package core

// DataService 数据服务集成
type DataService interface {
	// 钱包服务
	WalletService

	// 消息服务
	MessageService

	// 话题服务
	TopicService

	// 广场泡泡服务
	IndexPostsService

	// 推文服务
	TweetService
	TweetManageService
	TweetHelpService

	// 附件检测服务
	AttachmentCheckService

	// 评论服务
	CommentService
	CommentManageService

	// 用户服务
	UserManageService

	// 安全服务
	SecurityService
}
