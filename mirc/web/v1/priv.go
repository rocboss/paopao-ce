package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Priv 私有授权的服务
type Priv struct {
	Schema `mir:"v1,chain"`

	// UploadAttachment 上传资源
	UploadAttachment func(Post, web.UploadAttachmentReq) web.UploadAttachmentResp `mir:"attachment"`

	// DownloadAttachmentPrecheck 下载资源预检
	DownloadAttachmentPrecheck func(Get, web.DownloadAttachmentPrecheckReq) web.DownloadAttachmentPrecheckResp `mir:"attachment/precheck"`

	// DownloadAttachment 下载资源
	DownloadAttachment func(Get, web.DownloadAttachmentReq) web.DownloadAttachmentResp `mir:"attachment"`

	// CreateTweet 发布动态
	CreateTweet func(Post, Chain, web.CreateTweetReq) web.CreateTweetResp `mir:"post"`

	// DeleteTweet 删除动态
	DeleteTweet func(Delete, web.DeleteTweetReq) `mir:"post"`

	// StarTweet 动态点赞操作
	StarTweet func(Post, web.StarTweetReq) web.StarTweetResp `mir:"post/star"`

	// CollectionTweet 动态收藏操作
	CollectionTweet func(Post, web.CollectionTweetReq) web.CollectionTweetResp `mir:"post/collection"`

	// LockTweet 锁定动态
	LockTweet func(Post, web.LockTweetReq) web.LockTweetResp `mir:"post/lock"`

	// StickTweet 置顶动态
	StickTweet func(Post, web.StickTweetReq) web.StickTweetResp `mir:"post/stick"`

	// HighlightTweet 推文亮点设置
	HighlightTweet func(Post, web.HighlightTweetReq) web.HighlightTweetResp `mir:"post/highlight"`

	// VisibleTweet 修改动态可见度
	VisibleTweet func(Post, web.VisibleTweetReq) web.VisibleTweetResp `mir:"post/visibility"`

	// CreateTweetComment 发布动态评论
	CreateComment func(Post, web.CreateCommentReq) web.CreateCommentResp `mir:"post/comment"`

	// DeletePostComment 删除动态评论
	DeleteComment func(Delete, web.DeleteCommentReq) `mir:"post/comment"`

	// HighlightComment 精选动态评论
	HighlightComment func(Post, web.HighlightCommentReq) web.HighlightCommentResp `mir:"post/comment/highlight"`

	// CreateCommentReply 发布评论回复
	CreateCommentReply func(Post, web.CreateCommentReplyReq) web.CreateCommentReplyResp `mir:"post/comment/reply"`

	// DeleteCommentReply 删除评论回复
	DeleteCommentReply func(Delete, web.DeleteCommentReplyReq) `mir:"post/comment/reply"`

	// ThumbsUpTweetComment 点赞评论
	ThumbsUpTweetComment func(Post, web.TweetCommentThumbsReq) `mir:"tweet/comment/thumbsup"`

	// ThumbsDownTweetComment 点踩评论
	ThumbsDownTweetComment func(Post, web.TweetCommentThumbsReq) `mir:"tweet/comment/thumbsdown"`

	// ThumbsUpTweetReply 点赞评论回复·
	ThumbsUpTweetReply func(Post, web.TweetReplyThumbsReq) `mir:"tweet/reply/thumbsup"`

	// ThumbsDownTweetReply 点踩评论回复
	ThumbsDownTweetReply func(Post, web.TweetReplyThumbsReq) `mir:"tweet/reply/thumbsdown"`

	// StickTopic 置顶话题
	StickTopic func(Post, web.StickTopicReq) web.StickTopicResp `mir:"topic/stick"`

	// PinTopic 钉住话题
	PinTopic func(Post, web.PinTopicReq) web.PinTopicResp `mir:"topic/pin"`

	// FollowTopic 关注话题
	FollowTopic func(Post, web.FollowTopicReq) `mir:"topic/follow"`

	// UnfollowTopic 取消关注话题
	UnfollowTopic func(Post, web.UnfollowTopicReq) `mir:"topic/unfollow"`

	// 长文章相关API
	// CreateArticle 创建长文章
	CreateArticle func(Post, Chain, web.CreateArticleReq) web.CreateArticleResp `mir:"articles"`

	// UpdateArticle 更新长文章
	UpdateArticle func(Put, Chain, web.UpdateArticleReq) `mir:"articles"`

	// DeleteArticle 删除长文章
	DeleteArticle func(Delete, web.DeleteArticleReq) `mir:"articles"`

	// StarArticle 长文章点赞操作
	StarArticle func(Post, web.ArticleStarReq) web.ArticleStarResp `mir:"articles/star"`

	// CollectionArticle 长文章收藏操作
	CollectionArticle func(Post, web.ArticleCollectionReq) web.ArticleCollectionResp `mir:"articles/collection"`

	// CreateArticleComment 发布长文章评论
	CreateArticleComment func(Post, web.CreateArticleCommentReq) web.CreateArticleCommentResp `mir:"articles/comment"`

	// DeleteArticleComment 删除长文章评论
	DeleteArticleComment func(Delete, web.DeleteArticleCommentReq) `mir:"articles/comment"`

	// CreateArticleCommentReply 发布评论回复
	CreateArticleCommentReply func(Post, web.CreateArticleCommentReplyReq) web.CreateArticleCommentReplyResp `mir:"articles/comment/reply"`

	// DeleteArticleCommentReply 删除评论回复
	DeleteArticleCommentReply func(Delete, web.DeleteArticleCommentReplyReq) `mir:"articles/comment/reply"`

	// CreateArticleSeries 创建长文章系列
	CreateArticleSeries func(Post, web.CreateArticleSeriesReq) web.CreateArticleSeriesResp `mir:"articles/series"`

	// UpdateArticleSeries 更新长文章系列
	UpdateArticleSeries func(Put, web.UpdateArticleSeriesReq) `mir:"articles/series"`

	// DeleteArticleSeries 删除长文章系列
	DeleteArticleSeries func(Delete, web.DeleteArticleSeriesReq) `mir:"articles/series"`

	// AdminDeleteArticle 管理员删除长文章
	AdminDeleteArticle func(Delete, web.AdminDeleteArticleReq) `mir:"admin/articles"`
}
