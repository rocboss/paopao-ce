package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

const (
	SearchTypeDefault SearchType = "search"
	SearchTypeTag     SearchType = "tag"
)

type SearchType string

type QueryReq struct {
	Query      string
	Visibility []model.PostVisibleT
	Type       SearchType
}

type QueryResp struct {
	Items []*model.PostFormated
	Total int64
}

type TsDocItem struct {
	Post    *model.Post
	Content string
}

// TweetSearchService tweet search service interface
type TweetSearchService interface {
	IndexName() string
	AddDocuments(data []TsDocItem, primaryKey ...string) (bool, error)
	DeleteDocuments(identifiers []string) error
	Search(user *model.User, q *QueryReq, offset, limit int) (*QueryResp, error)
}
