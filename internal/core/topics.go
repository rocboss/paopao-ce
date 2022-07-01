package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// TopicService 话题服务
type TopicService interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
	DeleteTag(tag *model.Tag) error
	GetTags(conditions *model.ConditionsT, offset, limit int) ([]*model.Tag, error)
	GetTagsByKeyword(keyword string) ([]*model.Tag, error)
}
