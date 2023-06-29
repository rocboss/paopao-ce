// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

type BaseServant struct {
	bindAny func(c *gin.Context, obj any) mir.Error
}

type DaoServant struct {
	*BaseServant

	Dsa   core.WebDataServantA
	Ds    core.DataService
	Ts    core.TweetSearchService
	Redis core.RedisCache
}

type JsonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

type SentryHubSetter interface {
	SetSentryHub(hub *sentry.Hub)
}

type UserSetter interface {
	SetUser(*ms.User)
}

type UserIdSetter interface {
	SetUserId(int64)
}

type PageInfoSetter interface {
	SetPageInfo(page, pageSize int)
}

func UserFrom(c *gin.Context) (*ms.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*ms.User)
		return user, ok
	}
	return nil, false
}

func UserIdFrom(c *gin.Context) (int64, bool) {
	if uid, exists := c.Get("UID"); exists {
		v, ok := uid.(int64)
		return v, ok
	}
	return -1, false
}

func UserNameFrom(c *gin.Context) (string, bool) {
	if username, exists := c.Get("USERNAME"); exists {
		v, ok := username.(string)
		return v, ok
	}
	return "", false
}

func bindAny(c *gin.Context, obj any) mir.Error {
	var errs xerror.ValidErrors
	err := c.ShouldBind(obj)
	if err != nil {
		return mir.NewError(xerror.InvalidParams.StatusCode(), xerror.InvalidParams.WithDetails(errs.Error()))
	}
	// setup *core.User if needed
	if setter, ok := obj.(UserSetter); ok {
		user, _ := UserFrom(c)
		setter.SetUser(user)
	}
	// setup UserId if needed
	if setter, ok := obj.(UserIdSetter); ok {
		uid, _ := UserIdFrom(c)
		setter.SetUserId(uid)
	}
	// setup PageInfo if needed
	if setter, ok := obj.(PageInfoSetter); ok {
		page, pageSize := app.GetPageInfo(c)
		setter.SetPageInfo(page, pageSize)
	}
	return nil
}

func bindAnySentry(c *gin.Context, obj any) mir.Error {
	hub := sentrygin.GetHubFromContext(c)
	var errs xerror.ValidErrors
	err := c.ShouldBind(obj)
	if err != nil {
		xerr := mir.NewError(xerror.InvalidParams.StatusCode(), xerror.InvalidParams.WithDetails(errs.Error()))
		if hub != nil {
			hub.CaptureException(errors.Wrap(xerr, "bind object"))
		}
		return xerr
	}
	// setup sentry hub if needed
	if setter, ok := obj.(SentryHubSetter); ok && hub != nil {
		setter.SetSentryHub(hub)
	}
	// setup *core.User if needed
	if setter, ok := obj.(UserSetter); ok {
		user, _ := UserFrom(c)
		setter.SetUser(user)
	}
	// setup UserId if needed
	if setter, ok := obj.(UserIdSetter); ok {
		uid, _ := UserIdFrom(c)
		setter.SetUserId(uid)
	}
	// setup PageInfo if needed
	if setter, ok := obj.(PageInfoSetter); ok {
		page, pageSize := app.GetPageInfo(c)
		setter.SetPageInfo(page, pageSize)
	}
	return nil

}

func RenderAny(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, &JsonResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(xerror.HttpStatusCode(err), &JsonResp{
			Code: err.StatusCode(),
			Msg:  err.Error(),
		})
	}
}

func (s *BaseServant) Bind(c *gin.Context, obj any) mir.Error {
	return s.bindAny(c, obj)
}

func (s *BaseServant) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, &JsonResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(xerror.HttpStatusCode(err), &JsonResp{
			Code: err.StatusCode(),
			Msg:  err.Error(),
		})
	}
}

func (s *DaoServant) GetTweetBy(id int64) (*ms.PostFormated, error) {
	post, err := s.Ds.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	postContents, err := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, err
	}
	users, err := s.Ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, err
	}
	// 数据整合
	postFormated := post.Format()
	for _, user := range users {
		postFormated.User = user.Format()
	}
	for _, content := range postContents {
		if content.PostID == post.ID {
			postFormated.Contents = append(postFormated.Contents, content.Format())
		}
	}
	return postFormated, nil
}

func (s *DaoServant) PushPostsToSearch(c context.Context) {
	if err := s.Redis.SetPushToSearchJob(c); err == nil {
		defer s.Redis.DelPushToSearchJob(c)

		splitNum := 1000
		totalRows, _ := s.Ds.GetPostCount(&ms.ConditionsT{
			"visibility IN ?": []core.PostVisibleT{core.PostVisitPublic, core.PostVisitFriend},
		})
		pages := math.Ceil(float64(totalRows) / float64(splitNum))
		nums := int(pages)
		for i := 0; i < nums; i++ {
			posts, postsFormated, err := s.GetTweetList(&ms.ConditionsT{}, i*splitNum, splitNum)
			if err != nil || len(posts) != len(postsFormated) {
				continue
			}
			for i, pf := range postsFormated {
				contentFormated := ""
				for _, content := range pf.Contents {
					if content.Type == ms.ContentTypeText || content.Type == ms.ContentTypeTitle {
						contentFormated = contentFormated + content.Content + "\n"
					}
				}
				docs := []core.TsDocItem{{
					Post:    posts[i],
					Content: contentFormated,
				}}
				s.Ts.AddDocuments(docs, fmt.Sprintf("%d", posts[i].ID))
			}
		}
	} else {
		logrus.Errorf("redis: set JOB_PUSH_TO_SEARCH error: %s", err)
	}
}

func (s *DaoServant) PushPostToSearch(post *ms.Post) {
	postFormated := post.Format()
	postFormated.User = &ms.UserFormated{
		ID: post.UserID,
	}
	contents, _ := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	for _, content := range contents {
		postFormated.Contents = append(postFormated.Contents, content.Format())
	}

	contentFormated := ""
	for _, content := range postFormated.Contents {
		if content.Type == ms.ContentTypeText || content.Type == ms.ContentTypeTitle {
			contentFormated = contentFormated + content.Content + "\n"
		}
	}

	docs := []core.TsDocItem{{
		Post:    post,
		Content: contentFormated,
	}}
	s.Ts.AddDocuments(docs, fmt.Sprintf("%d", post.ID))
}

func (s *DaoServant) DeleteSearchPost(post *ms.Post) error {
	return s.Ts.DeleteDocuments([]string{fmt.Sprintf("%d", post.ID)})
}

func (s *DaoServant) GetTweetList(conditions *ms.ConditionsT, offset, limit int) ([]*ms.Post, []*ms.PostFormated, error) {
	posts, err := s.Ds.GetPosts(conditions, offset, limit)
	if err != nil {
		return nil, nil, err
	}
	postFormated, err := s.Ds.MergePosts(posts)
	return posts, postFormated, err
}

func NewBindAnyFn() func(c *gin.Context, obj any) mir.Error {
	if conf.UseSentryGin() {
		return bindAnySentry
	}
	return bindAny
}

func NewBaseServant() *BaseServant {
	return &BaseServant{
		bindAny: NewBindAnyFn(),
	}
}

func NewDaoServant() *DaoServant {
	return &DaoServant{
		BaseServant: NewBaseServant(),
		Redis:       cache.NewRedisCache(),
		Dsa:         dao.WebDataServantA(),
		Ds:          dao.DataService(),
		Ts:          dao.TweetSearchService(),
	}
}
