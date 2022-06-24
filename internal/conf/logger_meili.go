package conf

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/sirupsen/logrus"
)

type meiliLogData []map[string]interface{}

type meiliLogHook struct {
	config    meilisearch.ClientConfig
	idxName   string
	addDocsCh chan *meiliLogData
}

func (h *meiliLogHook) Fire(entry *logrus.Entry) error {
	data := meiliLogData{{
		"id":      entry.Time.Unix(),
		"time":    entry.Time,
		"level":   entry.Level,
		"message": entry.Message,
		"data":    entry.Data,
	}}

	// 先尝试进log缓存，否则直接新开goroutine加文档
	select {
	case h.addDocsCh <- &data:
	default:
		go func(index *meilisearch.Index, item meiliLogData) {
			index.AddDocuments(item)
		}(h.index(), data)
	}

	return nil
}

func (h *meiliLogHook) handleAddDocs() {
	index := h.index()
	for item := range h.addDocsCh {
		index.AddDocuments(item)
	}
}

func (h *meiliLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *meiliLogHook) index() *meilisearch.Index {
	return meilisearch.NewClient(h.config).Index(h.idxName)
}

func newMeiliLogHook() *meiliLogHook {
	hook := &meiliLogHook{
		config: meilisearch.ClientConfig{
			Host:   loggerMeiliSetting.Endpoint(),
			APIKey: loggerMeiliSetting.ApiKey,
		},
		idxName: loggerMeiliSetting.Index,
	}

	client := meilisearch.NewClient(hook.config)
	index := client.Index(hook.idxName)
	if _, err := index.FetchInfo(); err != nil {
		client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        hook.idxName,
			PrimaryKey: "id",
		})
		sortableAttributes := []string{
			"time",
		}
		index.UpdateSortableAttributes(&sortableAttributes)
	}

	// 初始化addDocsCh
	hook.addDocsCh = make(chan *meiliLogData, loggerMeiliSetting.maxLogBuffer())

	// 启动后台log工作者
	for minWork := loggerMeiliSetting.minWork(); minWork > 0; minWork-- {
		go hook.handleAddDocs()
	}

	return hook
}
