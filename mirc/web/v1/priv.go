package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(Priv))
}

// Priv 私有授权的服务
type Priv struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// UploadAttachment 上传资源
	UploadAttachment func(Post, web.UploadAttachmentReq) web.UploadAttachmentResp `mir:"/attachment"`

	// DownloadAttachmentPrecheck 下载资源预检
	DownloadAttachmentPrecheck func(Get, web.DownloadAttachmentPrecheckReq) web.DownloadAttachmentPrecheckResp `mir:"/attachment/precheck"`

	// DownloadAttachment 下载资源
	DownloadAttachment func(Get, web.DownloadAttachmentReq) web.DownloadAttachmentResp `mir:"/attachment"`

	// CreateTweet 发布动态
	CreateTweet func(Post, web.CreateTweetReq) web.CreateTweetResp `mir:"/post"`

	// DeleteTweet 删除动态
	DeleteTweet func(Delete, web.DeleteTweetReq) `mir:"/post"`

	// StarTweet 动态点赞操作
	StarTweet func(Post, web.StarTweetReq) web.StarTweetResp `mir:"/post/star"`

	// CollectionTweet 动态收藏操作
	CollectionTweet func(Post, web.CollectionTweetReq) web.CollectionTweetResp `mir:"/post/collection"`

	// LockTweet 锁定动态
	LockTweet func(Post, web.LockTweetReq) web.LockTweetResp `mir:"/post/lock"`

	// StickTweet 置顶动态
	StickTweet func(Post, web.StickTweetReq) web.StickTweetResp `mir:"/post/stick"`

	// VisibleTweet 修改动态可见度
	VisibleTweet func(Post, web.VisibleTweetReq) web.VisibleTweetResp `mir:"/post/visibility"`

	// CreateTweetComment 发布动态评论
	CreateComment func(Post, web.CreateCommentReq) web.CreateCommentResp `mir:"/post/comment"`

	// DeletePostComment 删除动态评论
	DeleteComment func(Delete, web.DeleteCommentReq) `mir:"/post/comment"`

	// CreateCommentReply 发布评论回复
	CreateCommentReply func(Post, web.CreateCommentReplyReq) web.CreateCommentReplyResp `mir:"/post/comment/reply"`

	// DeleteCommentReply 删除评论回复
	DeleteCommentReply func(Delete, web.DeleteCommentReplyReq) `mir:"/post/comment/reply"`

	// ThumbsUpTweetComment 点赞评论
	ThumbsUpTweetComment func(Post, web.TweetCommentThumbsReq) `mir:"/tweet/comment/thumbsup"`

	// ThumbsDownTweetComment 点踩评论
	ThumbsDownTweetComment func(Post, web.TweetCommentThumbsReq) `mir:"/tweet/comment/thumbsdown"`

	// ThumbsUpTweetReply 点赞评论回复
	ThumbsUpTweetReply func(Post, web.TweetReplyThumbsReq) `mir:"/tweet/reply/thumbsup"`

	// ThumbsDownTweetReply 点踩评论回复
	ThumbsDownTweetReply func(Post, web.TweetReplyThumbsReq) `mir:"/tweet/reply/thumbsdown"`

	// StickTopic 置顶动态
	StickTopic func(Post, web.StickTopicReq) web.StickTopicResp `mir:"/topic/stick"`

	// FollowTopic 关注话题
	FollowTopic func(Post, web.FollowTopicReq) `mir:"/topic/follow"`

	// UnfollowTopic 取消关注话题
	UnfollowTopic func(Post, web.UnfollowTopicReq) `mir:"/topic/unfollow"`
}
