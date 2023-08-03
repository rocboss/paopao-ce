package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// GetUserKeysReq 获取一个用户的所有key
type GetUserKeysReq struct {
	UserId   int64
	UserName string
}

func (g *GetUserKeysReq) Bind(c *gin.Context) mir.Error {
	uid, ok := base.UserFrom(c)
	if !ok {
		return xerror.UnauthorizedTokenError
	}
	g.UserId = uid.ID
	g.UserName = uid.Username
	return nil
}

type KeyInfo struct {
	ShareKey    string `json:"share_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetUserKeysResp struct {
	ShareKeys []KeyInfo `json:"shareKeys"`
}

// 逻辑删除服务
type DeleteKeyReq struct {
	ShareKey string `json:"share_key"`
	UserId   int64
	UserName string
}

func (d *DeleteKeyReq) Bind(c *gin.Context) mir.Error {
	uid, ok := base.UserFrom(c)
	if !ok {
		return xerror.UnauthorizedTokenError
	}
	d.UserId = uid.ID
	d.UserName = uid.Username
	d.ShareKey = convert.StrTo(c.Query("share_key")).String()
	return nil
}

type DeleteKeyResp struct {
	Status string
}
