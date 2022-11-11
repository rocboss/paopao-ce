package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model"
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
