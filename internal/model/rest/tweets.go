package rest

import "github.com/rocboss/paopao-ce/internal/model"

type IndexTweetsResp struct {
	Tweets []*model.PostFormated
	Total  int64
}
