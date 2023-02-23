// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/servants/web/broker"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/sirupsen/logrus"
)

func GetPostList(c *gin.Context) {
	response := app.NewResponse(c)

	q := &core.QueryReq{
		Query: c.Query("query"),
		Type:  "search",
	}
	if c.Query("type") == "tag" {
		q.Type = "tag"
	}

	user, _ := userFrom(c)
	offset, limit := app.GetPageOffset(c)
	if q.Query == "" && q.Type == "search" {
		resp, err := broker.GetIndexPosts(user, offset, limit)
		if err != nil {
			logrus.Errorf("service.GetPostList err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}

		response.ToResponseList(resp.Tweets, resp.Total)
	} else {
		posts, totalRows, err := broker.GetPostListFromSearch(user, q, offset, limit)

		if err != nil {
			logrus.Errorf("service.GetPostListFromSearch err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}
		response.ToResponseList(posts, totalRows)
	}
}

func GetPost(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	postFormated, err := broker.GetPost(postID)

	if err != nil {
		logrus.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	response.ToResponse(postFormated)
}

func CreatePost(c *gin.Context) {
	param := broker.PostCreationReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")
	post, err := broker.CreatePost(c, userID.(int64), param)

	if err != nil {
		logrus.Errorf("service.CreatePost err: %v\n", err)
		response.ToErrorResponse(errcode.CreatePostFailed)
		return
	}

	response.ToResponse(post)
}

func DeletePost(c *gin.Context) {
	param := broker.PostDelReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, exist := userFrom(c)
	if !exist {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	err := broker.DeletePost(user, param.ID)
	if err != nil {
		logrus.Errorf("service.DeletePost err: %v\n", err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(nil)
}

func GetPostStar(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")

	_, err := broker.GetPostStar(postID, userID.(int64))
	if err != nil {
		response.ToResponse(gin.H{
			"status": false,
		})

		return
	}

	response.ToResponse(gin.H{
		"status": true,
	})
}

func PostStar(c *gin.Context) {
	param := broker.PostStarReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")

	status := false
	star, err := broker.GetPostStar(param.ID, userID.(int64))
	if err != nil {
		// 创建Star
		_, err = broker.CreatePostStar(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		err = broker.DeletePostStar(star)
	}

	if err != nil {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	response.ToResponse(gin.H{
		"status": status,
	})
}

func GetPostCollection(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")

	_, err := broker.GetPostCollection(postID, userID.(int64))
	if err != nil {
		response.ToResponse(gin.H{
			"status": false,
		})

		return
	}

	response.ToResponse(gin.H{
		"status": true,
	})
}

func PostCollection(c *gin.Context) {
	param := broker.PostCollectionReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")

	status := false
	collection, err := broker.GetPostCollection(param.ID, userID.(int64))
	if err != nil {
		// 创建collection
		_, err = broker.CreatePostCollection(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		err = broker.DeletePostCollection(collection)
	}

	if err != nil {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	response.ToResponse(gin.H{
		"status": status,
	})
}

func LockPost(c *gin.Context) {
	param := broker.PostLockReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")

	// 获取Post
	postFormated, err := broker.GetPost(param.ID)
	if err != nil {
		logrus.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if postFormated.UserID != user.(*core.User).ID && !user.(*core.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}
	err = broker.LockPost(param.ID)
	if err != nil {
		logrus.Errorf("service.LockPost err: %v\n", err)
		response.ToErrorResponse(errcode.LockPostFailed)
		return
	}

	response.ToResponse(gin.H{
		"lock_status": 1 - postFormated.IsLock,
	})
}

func StickPost(c *gin.Context) {
	param := broker.PostStickReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")

	// 获取Post
	postFormated, err := broker.GetPost(param.ID)
	if err != nil {
		logrus.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if !user.(*core.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}
	err = broker.StickPost(param.ID)
	if err != nil {
		logrus.Errorf("service.StickPost err: %v\n", err)
		response.ToErrorResponse(errcode.LockPostFailed)
		return
	}

	response.ToResponse(gin.H{
		"top_status": 1 - postFormated.IsTop,
	})
}

func VisiblePost(c *gin.Context) {
	param := broker.PostVisibilityReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := userFrom(c)
	if err := broker.VisiblePost(user, param.ID, param.Visibility); err != nil {
		logrus.Errorf("service.VisiblePost err: %v\n", err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{
		"visibility": param.Visibility,
	})
}

func GetPostTags(c *gin.Context) {
	param := broker.PostTagsReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tags, err := broker.GetPostTags(&param)
	if err != nil {
		logrus.Errorf("service.GetPostTags err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostTagsFailed)
		return

	}

	response.ToResponse(gin.H{
		"topics": tags,
	})
}
