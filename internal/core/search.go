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

type DocItems []map[string]interface{}

// TweetSearchService tweet search service interface
type TweetSearchService interface {
	VersionInfo

	IndexName() string
	AddDocuments(documents DocItems, primaryKey ...string) (bool, error)
	DeleteDocuments(identifiers []string) error
	Search(q *QueryReq, offset, limit int) (*QueryResp, error)
}
