package dbr

type Rank struct {
	UserName              string `json:"user_name"`
	ShareKey              string `json:"share_key"`
	AllDownLoadCount      int    `json:"all_download_count"`
	DownloadCountPerWeek  int    `json:"download_count_per_week"`
	DownloadCountPerMonth int    `json:"download_count_per_month"`
}

type HighQualityRanking struct {
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
	Avatar   string `json:"avatar"`
}

// GetHighQualityRankingResp 优质榜单返回数据
type GetHighQualityRankingResp struct {
	UserName           string `json:"name"`
	NickName           string `json:"nickname"`
	Avatar             string `json:"avatar"`
	PostCount          int64  `json:"post_count"`
	LikesCount         int64  `json:"likes"`
	ComprehensiveScore int64  `json:"comprehensive_score"`
}

// GetDownloadRankListResp 下载榜单返回数据
type GetDownloadRankListResp struct {
	UserName string `json:"name"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Download int64  `json:"download"`
}
