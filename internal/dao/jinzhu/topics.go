package jinzhu

import (
	"strings"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

var (
	_ core.TopicService = (*topicServant)(nil)
)

type topicServant struct {
	db *gorm.DB
}

func newTopicService(db *gorm.DB) core.TopicService {
	return &topicServant{
		db: db,
	}
}

func (s *topicServant) CreateTag(tag *model.Tag) (*model.Tag, error) {
	return createTag(s.db, tag)
}

func (s *topicServant) DeleteTag(tag *model.Tag) error {
	return deleteTag(s.db, tag)
}

func (s *topicServant) GetTags(conditions *model.ConditionsT, offset, limit int) ([]*model.Tag, error) {
	return (&model.Tag{}).List(s.db, conditions, offset, limit)
}

func (s *topicServant) GetTagsByKeyword(keyword string) ([]*model.Tag, error) {
	tag := &model.Tag{}

	keyword = "%" + strings.Trim(keyword, " ") + "%"
	if keyword == "%%" {
		return tag.List(s.db, &model.ConditionsT{
			"ORDER": "quote_num DESC",
		}, 0, 6)
	} else {
		return tag.List(s.db, &model.ConditionsT{
			"tag LIKE ?": keyword,
			"ORDER":      "quote_num DESC",
		}, 0, 6)
	}
}
