package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/zinc"
)

const (
	SearchTypeDefault SearchType = "search"
	SearchTypeTag     SearchType = "tag"
)

type SearchType string

type QueryT struct {
	Query      string
	Visibility []model.PostVisibleT
	Type       SearchType
}

// SearchService search service interface that implement base zinc
type SearchService interface {
	CreateSearchIndex(indexName string)
	BulkPushDoc(data []map[string]interface{}) (bool, error)
	DelDoc(indexName, id string) error
	QueryAll(q *QueryT, indexName string, offset, limit int) (*zinc.QueryResultT, error)
	QuerySearch(indexName, query string, offset, limit int) (*zinc.QueryResultT, error)
	QueryTagSearch(indexName, query string, offset, limit int) (*zinc.QueryResultT, error)
}
