// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ms

// ArticleContentT 长文章内容类型
type ArticleContentT int

const (
	ArticleContentTypeTitle ArticleContentT = 1 + iota
	ArticleContentTypeText
	ArticleContentTypeImage
	ArticleContentTypeVideo
	ArticleContentTypeAudio
	ArticleContentTypeLink
	ArticleContentTypeAttachment
	ArticleContentTypePaidAttachment
)

// ArticleVisibleT 长文章可见性类型
type ArticleVisibleT int

const (
	ArticleVisiblePrivate ArticleVisibleT = 0 + iota // 私密
	ArticleVisiblePaid                                // 付费可见
	ArticleVisibleReserved1                            // 保留
	ArticleVisibleReserved2                            // 保留
	ArticleVisibleFriend                               // 好友可见
	ArticleVisibleFollowing                            // 关注可见
	ArticleVisibleReserved3                            // 保留
	ArticleVisibleReserved4                            // 保留
	ArticleVisiblePublic                               // 公开
)

// Article 文章别名
type Article = interface{}

// ArticleContent 文章内容别名
type ArticleContent = interface{}

// ArticleSeries 文章系列别名
type ArticleSeries = interface{}

// ArticleSeriesItem 系列文章关联别名
type ArticleSeriesItem = interface{}

// ArticleCollection 文章收藏别名
type ArticleCollection = interface{}

// ArticleStar 文章点赞别名
type ArticleStar = interface{}

// ArticleComment 文章评论别名
type ArticleComment = interface{}

// ArticleCommentReply 文章评论回复别名
type ArticleCommentReply = interface{}
