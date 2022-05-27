package dao

import (
	"github.com/rocboss/paopao-api/pkg/zinc"
)

type SearchType string

const SearchTypeDefault SearchType = "search"
const SearchTypeTag SearchType = "tag"

type QueryT struct {
	Query string
	Type  SearchType
}

func (d *Dao) CreateSearchIndex(indexName string) {
	// 不存在则创建索引
	d.zinc.CreateIndex(indexName, &zinc.ZincIndexProperty{
		"id": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Store:    true,
			Sortable: true,
		},
		"user_id": &zinc.ZincIndexPropertyT{
			Type:  "numeric",
			Index: true,
			Store: true,
		},
		"comment_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"collection_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"upvote_count": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"is_top": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"is_essence": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"content": &zinc.ZincIndexPropertyT{
			Type:           "text",
			Index:          true,
			Store:          true,
			Aggregatable:   true,
			Highlightable:  true,
			Analyzer:       "gse_search",
			SearchAnalyzer: "gse_standard",
		},
		"tags": &zinc.ZincIndexPropertyT{
			Type:  "keyword",
			Index: true,
			Store: true,
		},
		"ip_loc": &zinc.ZincIndexPropertyT{
			Type:  "keyword",
			Index: true,
			Store: true,
		},
		"latest_replied_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"attachment_price": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Sortable: true,
			Store:    true,
		},
		"created_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
		"modified_on": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Sortable: true,
			Store:    true,
		},
	})

}

func (d *Dao) BulkPushDoc(data []map[string]interface{}) (bool, error) {
	return d.zinc.BulkPushDoc(data)
}

func (d *Dao) DelDoc(indexName, id string) error {
	return d.zinc.DelDoc(indexName, id)
}

func (d *Dao) QueryAll(q *QueryT, indexName string, offset, limit int) (*zinc.QueryResultT, error) {
	// 普通搜索
	if q.Type == SearchTypeDefault && q.Query != "" {
		return d.QuerySearch(indexName, q.Query, offset, limit)
	}
	// Tag分类
	if q.Type == SearchTypeTag && q.Query != "" {
		return d.QueryTagSearch(indexName, q.Query, offset, limit)
	}

	queryMap := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]string{},
		},
		"sort": []string{"-is_top", "-latest_replied_on"},
		"from": offset,
		"size": limit,
	}
	rsp, err := d.zinc.EsQuery(indexName, queryMap)
	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (d *Dao) QuerySearch(indexName, query string, offset, limit int) (*zinc.QueryResultT, error) {
	rsp, err := d.zinc.EsQuery(indexName, map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"content": query,
			},
		},
		"sort": []string{"-is_top", "-latest_replied_on"},
		"from": offset,
		"size": limit,
	})
	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (d *Dao) QueryTagSearch(indexName, query string, offset, limit int) (*zinc.QueryResultT, error) {
	rsp, err := d.zinc.ApiQuery(indexName, map[string]interface{}{
		"search_type": "querystring",
		"query": map[string]interface{}{
			"term": "tags." + query + ":1",
		},
		"sort_fields": []string{"-is_top", "-latest_replied_on"},
		"from":        offset,
		"max_results": limit,
	})
	if err != nil {
		return nil, err
	}

	return rsp, err
}
