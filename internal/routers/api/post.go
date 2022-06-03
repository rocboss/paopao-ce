package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

func GetPostList(c *gin.Context) {
	response := app.NewResponse(c)

	q := &core.QueryT{
		Query: c.Query("query"),
		Type:  "search",
	}
	if c.Query("type") == "tag" {
		q.Type = "tag"
	}

	if q.Query == "" && q.Type == "search" {
		// 直接读库
		posts, err := service.GetPostList(&service.PostListReq{
			Conditions: &model.ConditionsT{
				"ORDER": "is_top DESC, latest_replied_on DESC",
			},
			Offset: (app.GetPage(c) - 1) * app.GetPageSize(c),
			Limit:  app.GetPageSize(c),
		})
		if err != nil {
			global.Logger.Errorf("service.GetPostList err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}
		totalRows, _ := service.GetPostCount(&model.ConditionsT{
			"ORDER": "latest_replied_on DESC",
		})

		response.ToResponseList(posts, totalRows)
	} else {
		posts, totalRows, err := service.GetPostListFromSearch(q, (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

		if err != nil {
			global.Logger.Errorf("service.GetPostListFromSearch err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}
		response.ToResponseList(posts, totalRows)
	}
}

func GetPost(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	postFormated, err := service.GetPost(postID)

	if err != nil {
		global.Logger.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	response.ToResponse(postFormated)
}

func CreatePost(c *gin.Context) {
	param := service.PostCreationReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")
	post, err := service.CreatePost(c, userID.(int64), param)

	if err != nil {
		global.Logger.Errorf("service.CreatePost err: %v\n", err)
		response.ToErrorResponse(errcode.CreatePostFailed)
		return
	}

	response.ToResponse(post)
}

func DeletePost(c *gin.Context) {
	param := service.PostDelReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")

	// 获取Post
	postFormated, err := service.GetPost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if postFormated.UserID != user.(*model.User).ID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	err = service.DeletePost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.DeletePost err: %v\n", err)
		response.ToErrorResponse(errcode.DeletePostFailed)
		return
	}

	response.ToResponse(nil)
}

func GetPostStar(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")

	_, err := service.GetPostStar(postID, userID.(int64))
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
	param := service.PostStarReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")

	status := false
	star, err := service.GetPostStar(param.ID, userID.(int64))
	if err != nil {
		// 创建Star
		service.CreatePostStar(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		service.DeletePostStar(star)
	}

	response.ToResponse(gin.H{
		"status": status,
	})
}

func GetPostCollection(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")

	_, err := service.GetPostCollection(postID, userID.(int64))
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
	param := service.PostCollectionReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")

	status := false
	collection, err := service.GetPostCollection(param.ID, userID.(int64))
	if err != nil {
		// 创建collection
		service.CreatePostCollection(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		service.DeletePostCollection(collection)
	}

	response.ToResponse(gin.H{
		"status": status,
	})
}

func LockPost(c *gin.Context) {
	param := service.PostLockReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")

	// 获取Post
	postFormated, err := service.GetPost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if postFormated.UserID != user.(*model.User).ID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}
	err = service.LockPost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.LockPost err: %v\n", err)
		response.ToErrorResponse(errcode.LockPostFailed)
		return
	}

	response.ToResponse(gin.H{
		"lock_status": 1 - postFormated.IsLock,
	})
}

func StickPost(c *gin.Context) {
	param := service.PostStickReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")

	// 获取Post
	postFormated, err := service.GetPost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}
	err = service.StickPost(param.ID)
	if err != nil {
		global.Logger.Errorf("service.StickPost err: %v\n", err)
		response.ToErrorResponse(errcode.LockPostFailed)
		return
	}

	response.ToResponse(gin.H{
		"top_status": 1 - postFormated.IsTop,
	})
}

func GetPostTags(c *gin.Context) {
	param := service.PostTagsReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tags, err := service.GetPostTags(&param)
	if err != nil {
		global.Logger.Errorf("service.GetPostTags err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostTagsFailed)
		return

	}

	response.ToResponse(tags)
}
