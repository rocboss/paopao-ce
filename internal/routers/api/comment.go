package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

func GetPostComments(c *gin.Context) {

	postID := convert.StrTo(c.Query("id")).MustInt64()
	response := app.NewResponse(c)

	svc := service.New(c)
	contents, totalRows, err := svc.GetPostComments(postID, "id ASC", 0, 0)

	if err != nil {
		global.Logger.Errorf("svc.GetPostComments err: %v\n", err)
		response.ToErrorResponse(errcode.GetCommentsFailed)
		return
	}

	response.ToResponseList(contents, totalRows)
}

func CreatePostComment(c *gin.Context) {
	param := service.CommentCreationReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")
	svc := service.New(c)
	comment, err := svc.CreatePostComment(userID.(int64), param)

	if err != nil {
		if err == errcode.MaxCommentCount {
			response.ToErrorResponse(errcode.MaxCommentCount)
		} else {
			global.Logger.Errorf("svc.CreatePostComment err: %v\n", err)
			response.ToErrorResponse(errcode.CreateCommentFailed)
		}
		return
	}

	response.ToResponse(comment)
}

func DeletePostComment(c *gin.Context) {
	param := service.CommentDelReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	user, _ := c.Get("USER")
	svc := service.New(c)

	comment, err := svc.GetPostComment(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.GetPostComment err: %v\n", err)
		response.ToErrorResponse(errcode.GetCommentFailed)
		return
	}

	if user.(*model.User).ID != comment.UserID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	// 执行删除
	err = svc.DeletePostComment(comment)
	if err != nil {
		global.Logger.Errorf("svc.DeletePostComment err: %v\n", err)
		response.ToErrorResponse(errcode.DeleteCommentFailed)
		return
	}

	response.ToResponse(nil)
}

func CreatePostCommentReply(c *gin.Context) {
	param := service.CommentReplyCreationReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	user, _ := c.Get("USER")
	svc := service.New(c)

	comment, err := svc.CreatePostCommentReply(param.CommentID, param.Content, user.(*model.User).ID, param.AtUserID)
	if err != nil {
		global.Logger.Errorf("svc.CreatePostCommentReply err: %v\n", err)
		response.ToErrorResponse(errcode.CreateReplyFailed)
		return
	}

	response.ToResponse(comment)
}

func DeletePostCommentReply(c *gin.Context) {
	param := service.ReplyDelReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("USER")
	svc := service.New(c)

	reply, err := svc.GetPostCommentReply(param.ID)
	if err != nil {
		global.Logger.Errorf("svc.GetPostCommentReply err: %v\n", err)
		response.ToErrorResponse(errcode.GetReplyFailed)
		return
	}

	if user.(*model.User).ID != reply.UserID && !user.(*model.User).IsAdmin {
		response.ToErrorResponse(errcode.NoPermission)
		return
	}

	// 执行删除
	err = svc.DeletePostCommentReply(reply)
	if err != nil {
		global.Logger.Errorf("svc.DeletePostCommentReply err: %v\n", err)
		response.ToErrorResponse(errcode.DeleteCommentFailed)
		return
	}

	response.ToResponse(nil)
}
