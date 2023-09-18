// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
	// IndexPostsService

	// 推文服务
	TweetService
	TweetManageService
	TweetHelpService

	// 推文指标服务
	TweetMetricServantA

	// 评论服务
	CommentService
	CommentManageService

	// 用户服务
	UserManageService
	ContactManageService
	FollowingManageService

	// 安全服务
	SecurityService
	AttachmentCheckService
}

// WebDataServantA Web数据服务集成(版本A)
type WebDataServantA interface {
	// 话题服务
	TopicServantA

	// 推文服务
	TweetServantA
	TweetManageServantA
	TweetHelpServantA
}
