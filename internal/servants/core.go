package servants

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model"
)

type baseServant struct {
	// TODO
}

type baseBinding struct {
	// TODO
}

type baseRender struct {
	// TODO
}

func (baseServant) userFrom(c *gin.Context) (*model.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*model.User)
		return user, ok
	}
	return nil, false
}
