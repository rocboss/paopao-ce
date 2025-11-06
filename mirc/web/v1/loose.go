package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Loose 宽松授权的服务
type Loose struct {
	Schema `mir:"v1,chain"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"posts"`

	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get, web.GetUserTweetsReq) web.GetUserTweetsResp `mir:"user/posts"`

	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get, web.GetUserProfileReq) web.GetUserProfileResp `mir:"user/profile"`

	// TopicList 获取话题列表
	TopicList func(Get, web.TopicListReq) web.TopicListResp `mir:"tags"`

	// TweetComments 获取动态评论
	TweetComments func(Get, web.TweetCommentsReq) web.TweetCommentsResp `mir:"post/comments"`

	// TweetDetail 获取动态详情
	TweetDetail func(Get, web.TweetDetailReq) web.TweetDetailResp `mir:"post"`

	// GetArticle 获取长文章详情
	GetArticle func(Get, web.GetArticleReq) web.GetArticleResp `mir:"articles/detail"`

	// ListArticles 获取长文章列表
	ListArticles func(Get, web.ListArticlesReq) web.ListArticlesResp `mir:"articles"`

	// ArticleComments 获取长文章评论
	ArticleComments func(Get, web.ArticleCommentsReq) web.ArticleCommentsResp `mir:"articles/comments"`

	// GetArticleSeries 获取长文章系列详情
	GetArticleSeries func(Get, web.GetArticleSeriesReq) web.GetArticleSeriesResp `mir:"articles/series/detail"`

	// ListArticleSeries 获取长文章系列列表
	ListArticleSeries func(Get, web.ListArticleSeriesReq) web.ListArticleSeriesResp `mir:"articles/series"`

	// GetUserArticles 获取用户长文章列表
	GetUserArticles func(Get, web.GetUserArticlesReq) web.GetUserArticlesResp `mir:"user/articles"`

	// ArticleStarStatus 获取长文章点赞状态
	ArticleStarStatus func(Get, web.ArticleStarStatusReq) web.ArticleStarStatusResp `mir:"articles/star"`

	// ArticleCollectionStatus 获取长文章收藏状态
	ArticleCollectionStatus func(Get, web.ArticleCollectionStatusReq) web.ArticleCollectionStatusResp `mir:"articles/collection"`
}
