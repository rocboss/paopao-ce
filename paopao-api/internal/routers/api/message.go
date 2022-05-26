package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/internal/service"
	"github.com/rocboss/paopao-api/pkg/app"
	"github.com/rocboss/paopao-api/pkg/errcode"
)

func GetUnreadMsgCount(c *gin.Context) {
	response := app.NewResponse(c)

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}
	svc := service.New(c)

	count, _ := svc.GetUnreadCount(user.ID)

	response.ToResponse(gin.H{
		"count": count,
	})
}

func GetMessages(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	svc := service.New(c)
	messages, totalRows, err := svc.GetMessages(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		global.Logger.Errorf("svc.GetMessages err: %v\n", err)
		response.ToErrorResponse(errcode.GetMessagesFailed)
		return
	}

	response.ToResponseList(messages, totalRows)
}

func ReadMessage(c *gin.Context) {
	param := service.ReadMessageReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, _ := c.Get("UID")
	svc := service.New(c)
	err := svc.ReadMessage(param.ID, userID.(int64))
	if err != nil {
		global.Logger.Errorf("svc.ReadMessage err: %v\n", err)
		response.ToErrorResponse(errcode.ReadMessageFailed)
		return
	}

	response.ToResponse(nil)
}
