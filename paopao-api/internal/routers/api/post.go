package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/dao"
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/internal/service"
	"github.com/rocboss/paopao-api/pkg/app"
	"github.com/rocboss/paopao-api/pkg/convert"
	"github.com/rocboss/paopao-api/pkg/errcode"
)

func GetPostList(c *gin.Context) {
	response := app.NewResponse(c)

	q := &dao.QueryT{
		Query: c.Query("query"),
		Type:  "search",
	}
	if c.Query("type") == "tag" {
		q.Type = "tag"
	}

	svc := service.New(c)
	if q.Query == "" && q.Type == "search" {
		// 直接读库
		posts, err := svc.GetPostList(&service.PostListReq{
			Conditions: &model.ConditionsT{
				"ORDER": "is_top DESC, latest_replied_on DESC",
			},
			Offset: (app.GetPage(c) - 1) * app.GetPageSize(c),
			Limit:  app.GetPageSize(c),
		})
		if err != nil {
			global.Logger.Errorf("svc.GetPostList err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}
		totalRows, _ := svc.GetPostCount(&model.ConditionsT{
			"ORDER": "latest_replied_on DESC",
		})

		response.ToResponseList(posts, totalRows)
	} else {
		posts, totalRows, err := svc.GetPostListFromSearch(q, (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

		if err != nil {
			global.Logger.Errorf("svc.GetPostListFromSearch err: %v\n", err)
			response.ToErrorResponse(errcode.GetPostsFailed)
			return
		}
		response.ToResponseList(posts, totalRows)
	}
}

func GetPost(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	svc := service.New(c)
	postFormated, err := svc.GetPost(postID)

	if err != nil {
		global.Logger.Errorf("svc.GetPost err: %v\n", err)
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
	svc := service.New(c)
	post, err := svc.CreatePost(userID.(int64), param)

	if err != nil {
		global.Logger.Errorf("svc.CreatePost err: %v\n", err)
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
	svc := service.New(c)

	// 获取Post
	postFormated, err := svc.GetPost(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if postFormated.UserID != user.(*model.User).ID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	err = svc.DeletePost(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.DeletePost err: %v\n", err)
		response.ToErrorResponse(errcode.DeletePostFailed)
		return
	}

	response.ToResponse(nil)
}

func GetPostStar(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	svc := service.New(c)
	userID, _ := c.Get("UID")

	_, err := svc.GetPostStar(postID, userID.(int64))
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

	svc := service.New(c)
	userID, _ := c.Get("UID")

	status := false
	star, err := svc.GetPostStar(param.ID, userID.(int64))
	if err != nil {
		// 创建Star
		svc.CreatePostStar(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		svc.DeletePostStar(star)
	}

	response.ToResponse(gin.H{
		"status": status,
	})
}

func GetPostCollection(c *gin.Context) {
	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	svc := service.New(c)
	userID, _ := c.Get("UID")

	_, err := svc.GetPostCollection(postID, userID.(int64))
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

	svc := service.New(c)
	userID, _ := c.Get("UID")

	status := false
	collection, err := svc.GetPostCollection(param.ID, userID.(int64))
	if err != nil {
		// 创建collection
		svc.CreatePostCollection(param.ID, userID.(int64))
		status = true
	} else {
		// 取消Star
		svc.DeletePostCollection(collection)
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
	svc := service.New(c)

	// 获取Post
	postFormated, err := svc.GetPost(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.GetPost err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostFailed)
		return
	}

	if postFormated.UserID != user.(*model.User).ID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}
	err = svc.LockPost(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.LockPost err: %v\n", err)
		response.ToErrorResponse(errcode.LockPostFailed)
		return
	}

	response.ToResponse(gin.H{
		"lock_status": 1 - postFormated.IsLock,
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

	svc := service.New(c)

	tags, err := svc.GetPostTags(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetPostTags err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostTagsFailed)
		return

	}

	response.ToResponse(tags)
}
