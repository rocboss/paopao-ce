package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	AddEntry(new(WebPriv))
}

type WebPriv struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`

	// UploadAttachment 上传资源
	UploadAttachment func(Post) `mir:"/attachment"`

	// DownloadAttachmentPrecheck 下载资源预检
	DownloadAttachmentPrecheck func(Get) `mir:"/attachment/precheck"`

	// DownloadAttachment 下载资源
	DownloadAttachment func(Get) `mir:"/attachment"`

	// CreateTweet 发布动态
	CreateTweet func(Post, web.CreateTweetReq) web.CreateTweetResp `mir:"/post"`

	// DeleteTweet 删除动态
	DeleteTweet func(Delete) `mir:"/post"`

	// StarTweet 动态点赞操作
	StarTweet func(Post) `mir:"/post/start"`

	// CollectionTweet 动态收藏操作
	CollectionTweet func(Post) `mir:"/post/collection"`

	// LockTweet 锁定动态
	LockTweet func(Post) `mir:"/post/lock"`

	// StickTweet 置顶动态
	StickTweet func(Post) `mir:"/post/stick"`

	// VisiblePost 修改动态可见度
	VisiblePost func(Post) `mir:"/post/visibility"`

	// CreateTweetComment 发布动态评论
	CreateComment func(Post) `mir:"/post/comment"`

	// DeletePostComment 删除动态评论
	DeleteComment func(Delete) `mir:"/post/comment"`

	// CreateCommentReply 发布评论回复
	CreateCommentReply func(Post) `mir:"/post/comment/reply"`

	// DeleteCommentReply 删除评论回复
	DeleteCommentReply func(Delete) `mir:"/post/comment/reply"`
}
