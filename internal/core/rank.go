package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	GetHighQualityRankingResp = dbr.GetHighQualityRankingResp
)

type RankService interface {
	GetHighQualityRanking() ([]*GetHighQualityRankingResp, error)
}
