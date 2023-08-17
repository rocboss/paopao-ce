package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Rank = (*rankSrv)(nil)
)

type rankSrv struct {
	api.UnimplementedRankServant
	*base.DaoServant
}

func (s *rankSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *rankSrv) GetHighQuality() ([]*web.GetHighQualityRankingResp, mir.Error) {
	if s.Ds == nil {
		logrus.Errorf("GetKeyDetail err: %s", web.ErrDsNil)
		return nil, web.ErrDsNil
	} else {
		logrus.Info("GetKeyDetail success")
	}
	//调用数据源的方式查询所有的rank信息
	ranks, err := s.Ds.GetHighQualityRanking()
	if err != nil {
		return nil, web.ErrGetHighQualityRankingFailed
	}

	//将查询到的rank信息转换为GetHighQualityRankingResp结构体
	var rankInfos []*web.GetHighQualityRankingResp
	for _, rank := range ranks {
		rankInfos = append(rankInfos, &web.GetHighQualityRankingResp{
			UserName:           rank.UserName,
			Avatar:             rank.Avatar,
			PostCount:          rank.PostCount,
			LikesCount:         rank.LikesCount,
			ComprehensiveScore: rank.ComprehensiveScore,
		})
	}
	return rankInfos, nil
}

func NewRankServant(ds *base.DaoServant) api.Rank {
	return &rankSrv{
		DaoServant: ds,
	}
}
