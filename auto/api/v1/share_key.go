package v1

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"net/http"
)

type KeyQuery interface {
	_default_

	Chain() gin.HandlersChain

	GetKeyDetail(*web.GetUserKeysReq) (*web.GetUserKeysResp, mir.Error)
	mustEmbedUnimplementedShareKeyServant()
	DeleteKeyDetail(req *web.DeleteKeyReq) (*web.DeleteKeyResp, mir.Error)
}

// RegisterKeyQueryServant 使用路由组的方式注册路由 /v1/key
func RegisterKeyQueryServant(e *gin.Engine, s KeyQuery) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// 注册路由
	router.Handle("GET", "/user/keys", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.GetUserKeysReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.GetKeyDetail(req)
		s.Render(c, resp, err)
	})

	//逻辑删除，将status置为1
	router.Handle("POST", "/user/keys", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.DeleteKeyReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.DeleteKeyDetail(req)
		s.Render(c, resp, err)
	})
}

type UnimplementedShareKeyServant struct{}

// GetKeyDetail 用于提供接口的默认行为
func (UnimplementedShareKeyServant) GetKeyDetail(req *web.GetUserKeysReq) (*web.GetUserKeysResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedShareKeyServant) mustEmbedUnimplementedShareKeyServant() {}
