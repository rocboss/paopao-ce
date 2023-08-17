package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/pkg/convert"
)

// GetHighQualityResp 构造相应参数，包含一个GetHighQualityRankingResp数组
type GetHighQualityResp struct {
	List any `json:"list"`
}

type GetHighQualityRankingResp struct {
	UserName           string `json:"name"`
	Avatar             string `json:"avatar"`
	PostCount          int64  `json:"post_count"`
	LikesCount         int64  `json:"likes"`
	ComprehensiveScore int64  `json:"comprehensive_score"`
}

type GetRankListReq struct {
	ListType string `json:"list_type"`
}

type GetRankListResp struct {
	UserName           string `json:"name"`
	Avatar             string `json:"avatar"`
	PostCount          int64  `json:"post_count"`
	LikesCount         int64  `json:"likes"`
	ComprehensiveScore int64  `json:"comprehensive_score"`
}

func (g *GetRankListReq) Bind(c *gin.Context) mir.Error {
	g.ListType = convert.StrTo(c.Query("list_type")).String()
	return nil
}
