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
	"github.com/rocboss/paopao-ce/internal/core/cs"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/types"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseServant struct {
	bindAny  func(c *gin.Context, obj any) mir.Error
	bindJson func(c *gin.Context, obj any) mir.Error
}

type DaoServant struct {
	*BaseServant

	Dsa   core.WebDataServantA
	Ds    core.DataService
	Ts    core.TweetSearchService
	Redis core.RedisCache
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
		c.JSON(http.StatusOK, &joint.JsonResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(xerror.HttpStatusCode(err), &joint.JsonResp{
			Code: err.StatusCode(),
			Msg:  err.Error(),
		})
	}
}

func (s *BaseServant) Bind(c *gin.Context, obj any) mir.Error {
	return s.bindAny(c, obj)
}

func (s *BaseServant) BindJson(c *gin.Context, obj any) mir.Error {
	return s.bindJson(c, obj)
}

func (s *BaseServant) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, &joint.JsonResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(xerror.HttpStatusCode(err), &joint.JsonResp{
			Code: err.StatusCode(),
			Msg:  err.Error(),
		})
	}
}

func (s *DaoServant) PrepareUser(userId int64, user *ms.UserFormated) error {
	// guest用户的userId<0
	if userId < 0 {
		return nil
	}
	// friendMap, err := s.Ds.IsMyFriend(userId, user.ID)
	// if err != nil {
	// 	return err
	// }
	followMap, err := s.Ds.IsMyFollow(userId, user.ID)
	if err != nil {
		return err
	}
	// user.IsFriend, user.IsFollowing = friendMap[user.ID], followMap[user.ID]
	user.IsFollowing = followMap[user.ID]
	return nil
}

func (s *DaoServant) PrepareMessages(userId int64, messages []*ms.MessageFormated) error {
	// guest用户的userId<0
	if userId < 0 {
		return nil
	}
	userIds := make([]int64, 0, len(messages))
	for _, msg := range messages {
		if msg.SenderUser != nil {
			userIds = append(userIds, msg.SenderUserID)
		}
		if msg.ReceiverUser != nil {
			userIds = append(userIds, msg.ReceiverUserID)
		}
	}
	// friendMap, err := s.Ds.IsMyFriend(userId, userIds...)
	// if err != nil {
	// 	return err
	// }
	followMap, err := s.Ds.IsMyFollow(userId, userIds...)
	if err != nil {
		return err
	}
	for _, msg := range messages {
		if msg.SenderUser != nil {
			//  msg.SenderUser.IsFriend, msg.SenderUser.IsFollowing = friendMap[msg.SenderUserID], followMap[msg.SenderUserID]
			msg.SenderUser.IsFollowing = followMap[msg.SenderUserID]
		}
		if msg.ReceiverUser != nil {
			//  msg.ReceiverUser.IsFriend, msg.ReceiverUser.IsFollowing = friendMap[msg.ReceiverUserID], followMap[msg.ReceiverUserID]
			msg.ReceiverUser.IsFollowing = followMap[msg.ReceiverUserID]
		}
	}
	return nil
}

func (s *DaoServant) PrepareTweet(userId int64, tweet *ms.PostFormated) error {
	// 转换一下可见性的值
	tweet.Visibility = ms.PostVisibleT(tweet.Visibility.ToOutValue())
	// guest用户的userId<0
	if userId < 0 {
		return nil
	}
	// friendMap, err := s.Ds.IsMyFriend(userId, userIds)
	// if err != nil {
	// 	return err
	// }
	followMap, err := s.Ds.IsMyFollow(userId, tweet.UserID)
	if err != nil {
		return err
	}
	// tweet.User.IsFriend, tweet.User.IsFollowing = friendMap[tweet.UserID], followMap[tweet.UserID]
	tweet.User.IsFollowing = followMap[tweet.UserID]
	return nil
}

func (s *DaoServant) PrepareTweets(userId int64, tweets []*ms.PostFormated) error {
	userIdSet := make(map[int64]types.Empty, len(tweets))
	for _, tweet := range tweets {
		userIdSet[tweet.UserID] = types.Empty{}
		// 顺便转换一下可见性的值
		tweet.Visibility = ms.PostVisibleT(tweet.Visibility.ToOutValue())
	}
	// guest用户的userId<0
	if userId < 0 {
		return nil
	}
	userIds := make([]int64, 0, len(userIdSet))
	for id := range userIdSet {
		userIds = append(userIds, id)
	}
	// friendMap, err := s.Ds.IsMyFriend(userId, userIds...)
	// if err != nil {
	// 	return err
	// }
	followMap, err := s.Ds.IsMyFollow(userId, userIds...)
	if err != nil {
		return err
	}
	for _, tweet := range tweets {
		// tweet.User.IsFriend, tweet.User.IsFollowing = friendMap[tweet.UserID], followMap[tweet.UserID]
		tweet.User.IsFollowing = followMap[tweet.UserID]
	}
	return nil
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

func (s *DaoServant) PushAllPostToSearch() {
	events.OnEvent(&pushAllPostToSearchEvent{
		fn: s.pushAllPostToSearch,
	})
}

func (s *DaoServant) pushAllPostToSearch() error {
	ctx := context.Background()
	if err := s.Redis.SetPushToSearchJob(ctx); err == nil {
		defer s.Redis.DelPushToSearchJob(ctx)
		splitNum := 1000
		posts, totalRows, err := s.Ds.ListSyncSearchTweets(splitNum, 0)
		if err != nil {
			return fmt.Errorf("get first page tweets push to search failed: %s", err)
		}
		i, nums := 0, int(math.Ceil(float64(totalRows)/float64(splitNum)))
		for {
			postsFormated, xerr := s.Ds.MergePosts(posts)
			if xerr != nil || len(posts) != len(postsFormated) {
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
			if i++; i >= nums {
				break
			}
			if posts, _, err = s.Ds.ListSyncSearchTweets(splitNum, i*splitNum); err != nil {
				return fmt.Errorf("get tweets push to search failed: %s, limit[%d] offset[%d]", err, splitNum, i*splitNum)
			}
		}
	} else {
		return fmt.Errorf("redis: set JOB_PUSH_TO_SEARCH error: %w", err)
	}
	return nil
}

func (s *DaoServant) PushPostToSearch(post *ms.Post) {
	events.OnEvent(&pushPostToSearchEvent{
		fn:   s.pushPostToSearch,
		post: post,
	})
}

func (s *DaoServant) pushPostToSearch(post *ms.Post) {
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

func (s *DaoServant) RelationTypFrom(me *ms.User, username string) (res *cs.VistUser, err error) {
	res = &cs.VistUser{
		RelTyp:   cs.RelationSelf,
		Username: username,
	}
	// visit by self
	if me != nil && me.Username == username {
		res.UserId = me.ID
		return
	}
	he, xerr := s.Ds.GetUserByUsername(username)
	if xerr != nil || (he.Model != nil && he.ID <= 0) {
		return nil, errors.New("get user failed with username: " + username)
	}
	res.UserId = he.ID
	// visit by guest
	if me == nil {
		res.RelTyp = cs.RelationGuest
		return
	}
	// visit by admin/friend/other
	if me.IsAdmin {
		res.RelTyp = cs.RelationAdmin
	} else if s.Ds.IsFriend(me.ID, he.ID) {
		res.RelTyp = cs.RelationFriend
	} else {
		res.RelTyp = cs.RelationGuest
	}
	return
}

func NewBindAnyFn() func(c *gin.Context, obj any) mir.Error {
	if conf.UseSentryGin() {
		return bindAnySentry
	}
	return bindAny
}

func NewBindJsonFn() func(c *gin.Context, obj any) mir.Error {
	if conf.UseSentryGin() {
		return bindAnySentry
	}
	return bindAny
}

func NewBaseServant() *BaseServant {
	return &BaseServant{
		bindAny:  NewBindAnyFn(),
		bindJson: NewBindJsonFn(),
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
