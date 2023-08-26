package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	GetHighQualityRankingResp = dbr.GetHighQualityRankingResp
	GetDownloadRankListResp   = dbr.GetDownloadRankListResp
)

type RankService interface {
	GetHighQualityRanking() ([]*GetHighQualityRankingResp, error)
	GetDownloadRankList(listType int) ([]*GetDownloadRankListResp, error)
}
