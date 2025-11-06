// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// 长文章相关请求和响应结构

type ArticleContentItem struct {
	Type    ms.ArticleContentT `json:"type" binding:"required"`
	Content string             `json:"content" binding:"required"`
	Sort    int64              `json:"sort"`
}

type CreateArticleReq struct {
	BaseInfo      `json:"-" binding:"-"`
	Title         string               `json:"title" binding:"required,max=200"`
	Summary       string               `json:"summary" binding:"max=500"`
	SeriesID      *int64               `json:"series_id"`
	Contents      []ArticleContentItem `json:"contents" binding:"required,dive"`
	Visibility    ms.ArticleVisibleT   `json:"visibility"`
	Price         int64                `json:"price"`
	Tags          []string             `json:"tags"`
}

type CreateArticleResp struct {
	ID int64 `json:"id"`
}

type UpdateArticleReq struct {
	BaseInfo      `json:"-" binding:"-"`
	ID            int64                `json:"id" binding:"required"`
	Title         string               `json:"title" binding:"required,max=200"`
	Summary       string               `json:"summary" binding:"max=500"`
	SeriesID      *int64               `json:"series_id"`
	Contents      []ArticleContentItem `json:"contents" binding:"required,dive"`
	Visibility    ms.ArticleVisibleT   `json:"visibility"`
	Price         int64                `json:"price"`
	Tags          []string             `json:"tags"`
}

type DeleteArticleReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type GetArticleReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `form:"id" binding:"required"`
}

type GetArticleResp struct {
	joint.CacheResp
}

type ListArticlesReq struct {
	BaseInfo   `json:"-" binding:"-"`
	Query      string               `form:"query"`
	Visibility []ms.ArticleVisibleT `form:"visibility"`
	Type       string               `form:"type"`
	Style      string               `form:"style"`
	Page       int                  `form:"-" binding:"-"`
	PageSize   int                  `form:"-" binding:"-"`
}

type ListArticlesResp struct {
	joint.CachePageResp
}

type ArticleStarReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type ArticleStarResp struct {
	Status bool `json:"status"`
}

type ArticleCollectionReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type ArticleCollectionResp struct {
	Status bool `json:"status"`
}

type CreateArticleCommentReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64  `json:"id" binding:"required"`
	Content  string `json:"content" binding:"required,max=1000"`
}

type CreateArticleCommentResp struct {
	ID int64 `json:"id"`
}

type DeleteArticleCommentReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type ArticleCommentsReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `form:"id" binding:"required"`
	Page     int   `form:"-" binding:"-"`
	PageSize int   `form:"-" binding:"-"`
}

type ArticleCommentsResp struct {
	joint.CachePageResp
}

type CreateArticleCommentReplyReq struct {
	BaseInfo  `json:"-" binding:"-"`
	CommentID int64  `json:"comment_id" binding:"required"`
	ReplyToID *int64 `json:"reply_to_id"`
	Content   string `json:"content" binding:"required,max=1000"`
}

type CreateArticleCommentReplyResp struct {
	ID int64 `json:"id"`
}

type DeleteArticleCommentReplyReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type CreateArticleSeriesReq struct {
	BaseInfo   `json:"-" binding:"-"`
	Title      string             `json:"title" binding:"required,max=200"`
	Description string            `json:"description" binding:"max=1000"`
	Cover      string             `json:"cover"`
	Visibility ms.ArticleVisibleT `json:"visibility"`
	Tags       []string           `json:"tags"`
}

type CreateArticleSeriesResp struct {
	ID int64 `json:"id"`
}

type UpdateArticleSeriesReq struct {
	BaseInfo    `json:"-" binding:"-"`
	ID          int64              `json:"id" binding:"required"`
	Title       string             `json:"title" binding:"required,max=200"`
	Description string             `json:"description" binding:"max=1000"`
	Cover       string             `json:"cover"`
	Visibility  ms.ArticleVisibleT `json:"visibility"`
	Tags        []string           `json:"tags"`
}

type DeleteArticleSeriesReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `json:"id" binding:"required"`
}

type GetArticleSeriesReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `form:"id" binding:"required"`
}

type GetArticleSeriesResp struct {
	joint.CacheResp
}

type ListArticleSeriesReq struct {
	BaseInfo   `json:"-" binding:"-"`
	Query      string               `form:"query"`
	Visibility []ms.ArticleVisibleT `form:"visibility"`
	Page       int                  `form:"-" binding:"-"`
	PageSize   int                  `form:"-" binding:"-"`
}

type ListArticleSeriesResp struct {
	joint.CachePageResp
}

type AddToSeriesReq struct {
	BaseInfo `json:"-" binding:"-"`
	SeriesID int64 `json:"series_id" binding:"required"`
	ArticleID int64 `json:"article_id" binding:"required"`
	Sort     int   `json:"sort"`
}

type RemoveFromSeriesReq struct {
	BaseInfo  `json:"-" binding:"-"`
	SeriesID  int64 `json:"series_id" binding:"required"`
	ArticleID int64 `json:"article_id" binding:"required"`
}

type GetUserArticlesReq struct {
	BaseInfo `json:"-" binding:"-"`
	Username string `form:"username" binding:"required"`
	Style    string `form:"style"`
	Page     int    `form:"-" binding:"-"`
	PageSize int    `form:"-" binding:"-"`
}

type GetUserArticlesResp struct {
	joint.CachePageResp
}

type GetUserArticleCollectionsReq BasePageReq
type GetUserArticleCollectionsResp base.PageResp

type GetUserArticleStarsReq BasePageReq
type GetUserArticleStarsResp base.PageResp

type ArticleStarStatusReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `form:"id" binding:"required"`
}

type ArticleStarStatusResp struct {
	Status bool `json:"status"`
}

type ArticleCollectionStatusReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64 `form:"id" binding:"required"`
}

type ArticleCollectionStatusResp struct {
	Status bool `json:"status"`
}

type AdminDeleteArticleReq struct {
	BaseInfo `json:"-" binding:"-"`
	ID       int64  `json:"id" binding:"required"`
	Reason   string `json:"reason" binding:"required"`
}

// Bind methods for request validation and binding

func (r *ListArticlesReq) Bind(c *gin.Context) error {
	r.Page, r.PageSize = app.GetPageInfo(c)
	return nil
}

func (r *ArticleCommentsReq) Bind(c *gin.Context) error {
	r.Page, r.PageSize = app.GetPageInfo(c)
	return nil
}

func (r *ListArticleSeriesReq) Bind(c *gin.Context) error {
	r.Page, r.PageSize = app.GetPageInfo(c)
	return nil
}

func (r *GetUserArticlesReq) Bind(c *gin.Context) error {
	r.Page, r.PageSize = app.GetPageInfo(c)
	return nil
}

func (r *GetUserArticleCollectionsReq) Bind(c *gin.Context) error {
	return (*BasePageReq)(r).Bind(c)
}

func (r *GetUserArticleStarsReq) Bind(c *gin.Context) error {
	return (*BasePageReq)(r).Bind(c)
}

func (r *ArticleStarStatusReq) Bind(c *gin.Context) error {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.BaseInfo = BaseInfo{
		User: &ms.User{Model: &ms.Model{ID: userId}},
	}
	r.ID = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

func (r *ArticleCollectionStatusReq) Bind(c *gin.Context) error {
	userId, exist := base.UserIdFrom(c)
	if !exist {
		return xerror.UnauthorizedAuthNotExist
	}
	r.BaseInfo = BaseInfo{
		User: &ms.User{Model: &ms.Model{ID: userId}},
	}
	r.ID = convert.StrTo(c.Query("id")).MustInt64()
	return nil
}

func (r *CreateArticleReq) ValidateContents() error {
	totalWords := 0
	for _, content := range r.Contents {
		if content.Type == ms.ArticleContentTypeText {
			// 计算字数，这里简化为字符数，实际应该用中文分词库
			totalWords += len([]rune(content.Content))
		}
	}
	if totalWords > 50000 {
		return xerror.NewError(400, "文章字数超过50000字限制")
	}
	return nil
}
