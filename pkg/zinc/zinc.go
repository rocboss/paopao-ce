package zinc

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/json"
)

type ZincClient struct {
	*ZincClientConfig
}

type ZincClientConfig struct {
	ZincHost     string
	ZincUser     string
	ZincPassword string
}

type ZincIndex struct {
	Name        string             `json:"name"`
	StorageType string             `json:"storage_type"`
	Mappings    *ZincIndexMappings `json:"mappings"`
}

type ZincIndexMappings struct {
	Properties *ZincIndexProperty `json:"properties"`
}

type ZincIndexProperty map[string]*ZincIndexPropertyT

type ZincIndexPropertyT struct {
	Type           string `json:"type"`
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Aggregatable   bool   `json:"aggregatable"`
	Highlightable  bool   `json:"highlightable"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
	Format         string `json:"format"`
}

type QueryResultT struct {
	Took     int          `json:"took"`
	TimedOut bool         `json:"timed_out"`
	Hits     *HitsResultT `json:"hits"`
}

type HitsResultT struct {
	Total    *HitsResultTotalT `json:"total"`
	MaxScore float64           `json:"max_score"`
	Hits     []*HitItem        `json:"hits"`
}

type HitsResultTotalT struct {
	Value int64 `json:"value"`
}

type HitItem struct {
	Index     string      `json:"_index"`
	Type      string      `json:"_type"`
	ID        string      `json:"_id"`
	Score     float64     `json:"_score"`
	Timestamp time.Time   `json:"@timestamp"`
	Source    interface{} `json:"_source"`
}

// NewClient 获取ZincClient新实例
func NewClient(conf *conf.ZincSettingS) *ZincClient {
	return &ZincClient{
		ZincClientConfig: &ZincClientConfig{
			ZincHost:     conf.Host,
			ZincUser:     conf.User,
			ZincPassword: conf.Password,
		},
	}
}

// 创建索引
func (c *ZincClient) CreateIndex(name string, p *ZincIndexProperty) bool {
	data := &ZincIndex{
		Name:        name,
		StorageType: "disk",
		Mappings: &ZincIndexMappings{
			Properties: p,
		},
	}
	client := resty.New() // 创建一个restry客户端
	client.DisableWarn = true
	resp, err := client.R().SetBody(data).SetBasicAuth(c.ZincUser, c.ZincPassword).Put(c.ZincHost + "/api/index")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false
	}

	return true
}

// 检查索引是否存在
func (c *ZincClient) ExistIndex(name string) bool {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBasicAuth(c.ZincUser, c.ZincPassword).Get(c.ZincHost + "/api/index")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false
	}

	retData := &map[string]interface{}{}
	err = json.Unmarshal([]byte(resp.String()), retData)
	if err != nil {
		return false
	}

	if _, ok := (*retData)[name]; ok {
		return true
	}

	return false
}

// 新增/更新文档
func (c *ZincClient) PutDoc(name string, id int64, doc interface{}) (bool, error) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBody(doc).SetBasicAuth(c.ZincUser, c.ZincPassword).Put(fmt.Sprintf("%s/api/%s/_doc/%d", c.ZincHost, name, id))

	if err != nil {
		return false, err
	}

	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

// 批量新增文档
func (c *ZincClient) BulkPushDoc(docs []map[string]interface{}) (bool, error) {
	dataStr := ""
	for _, doc := range docs {
		str, err := json.Marshal(doc)
		if err == nil {
			dataStr = dataStr + string(str) + "\n"
		}
	}

	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBody(dataStr).SetBasicAuth(c.ZincUser, c.ZincPassword).Post(fmt.Sprintf("%s/api/_bulk", c.ZincHost))
	if err != nil {
		return false, err
	}

	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func (c *ZincClient) EsQuery(indexName string, q interface{}) (*QueryResultT, error) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBody(q).SetBasicAuth(c.ZincUser, c.ZincPassword).Post(fmt.Sprintf("%s/es/%s/_search", c.ZincHost, indexName))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZincClient) ApiQuery(indexName string, q interface{}) (*QueryResultT, error) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBody(q).SetBasicAuth(c.ZincUser, c.ZincPassword).Post(fmt.Sprintf("%s/api/%s/_search", c.ZincHost, indexName))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZincClient) DelDoc(indexName, id string) error {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().SetBasicAuth(c.ZincUser, c.ZincPassword).Delete(fmt.Sprintf("%s/api/%s/_doc/%s", c.ZincHost, indexName, id))
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	return nil
}
