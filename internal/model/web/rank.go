package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/pkg/convert"
)

// GetRankListResp 构造相应参数，包含一个GetHighQualityRankingResp数组
type GetRankListResp struct {
	List any `json:"list"`
}

type GetHighQualityRankingResp struct {
	UserName           string `json:"name"`
	NickName           string `json:"nickname"`
	Avatar             string `json:"avatar"`
	PostCount          int64  `json:"post_count"`
	LikesCount         int64  `json:"likes"`
	ComprehensiveScore int64  `json:"comprehensive_score"`
}

type GetDownloadRankListReq struct {
	ListType int `json:"list_type"`
}

type GetDownloadRankListResp struct {
	UserName string `json:"name"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Download int64  `json:"download"`
}

func (g *GetDownloadRankListReq) Bind(c *gin.Context) mir.Error {
	g.ListType = convert.StrTo(c.Query("list_type")).MustInt()
	return nil
}
