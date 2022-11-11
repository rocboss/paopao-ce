package core

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

type BaseServant struct {
	// TODO
}

type BaseBinding struct {
	// TODO
}

type BaseRender struct {
	// TODO
}

func (BaseServant) userFrom(c *gin.Context) (*model.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*model.User)
		return user, ok
	}
	return nil, false
}

func BindAny(c *gin.Context, obj any) error {
	var errs xerror.ValidErrors
	err := c.ShouldBind(obj)
	if err != nil {
		return fmt.Errorf("bind errs: %w", xerror.InvalidParams.WithDetails(errs.Error()))
	}
	return nil
}

func RenderAny(c *gin.Context, data any, err error) {
	if err == nil {
		hostname, _ := os.Hostname()
		if data == nil {
			data = gin.H{
				"code":      0,
				"msg":       "success",
				"tracehost": hostname,
			}
		} else {
			data = gin.H{
				"code":      0,
				"msg":       "success",
				"data":      data,
				"tracehost": hostname,
			}
		}
		c.JSON(http.StatusOK, data)
	} else {
		// TODO: error process logic
		resp := gin.H{"code": http.StatusInternalServerError, "msg": err.Error()}
		// xerr := &xerror.Error{}
		// if errors.As(err, xerr) {
		// 	resp["code"], resp["msg"] = xerr.Code(), xerr.Msg()
		// 	details := xerr.Details()
		// 	if len(details) > 0 {
		// 		resp["details"] = details
		// 	}
		// }
		c.JSON(http.StatusInternalServerError, resp)
	}
}
