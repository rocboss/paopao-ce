package v1

import (
	"fmt"
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"net/http"
)

type Rank interface {
	_default_

	GetHighQuality() ([]*web.GetHighQualityRankingResp, mir.Error)

	mustEmbedUnimplementedRankServant()
	GetDownloadRankList(req *web.GetDownloadRankListReq) ([]web.GetDownloadRankListResp, mir.Error)
	//Chain() gin.HandlersChain
}

// RegisterRankServant 注册路由
func RegisterRankServant(e *gin.Engine, s Rank) {
	router := e.Group("v1")
	// use chain for router
	//middlewares := s.Chain()
	//router.Use(middlewares...)

	//注册路由
	router.Handle("GET", "/rank/highquality", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		//获取GetHighQuality完成后的数据，然后将其放到GetHighQualityResp中，并返回
		resp := new(web.GetRankListResp)
		highQualityRanking, err := s.GetHighQuality()
		resp.List = highQualityRanking
		s.Render(c, resp, err)
	})

	//获取所有的rank
	router.Handle("GET", "/rank/list", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.GetDownloadRankListReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		fmt.Print(req.ListType)
		resp := new(web.GetRankListResp)
		downloadRankList, err := s.GetDownloadRankList(req)
		resp.List = downloadRankList
		//返回值
		s.Render(c, resp, err)
	})

}

// UnimplementedRankServant can be embedded to have forward compatible implementations.
type UnimplementedRankServant struct{}

// GetHighQuality 设置RegisterRankServant的默认实现
func (UnimplementedRankServant) GetHighQuality() (*web.GetHighQualityRankingResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedRankServant) mustEmbedUnimplementedRankServant() {}
