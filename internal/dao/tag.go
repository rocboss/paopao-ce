package dao

import (
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

func (d *dataServant) CreateTag(tag *model.Tag) (*model.Tag, error) {
	return d.createTag(d.engine, tag)
}

func (d *dataServant) DeleteTag(tag *model.Tag) error {
	return d.deleteTag(d.engine, tag)
}

func (d *dataServant) GetTags(conditions *model.ConditionsT, offset, limit int) ([]*model.Tag, error) {
	return (&model.Tag{}).List(d.engine, conditions, offset, limit)
}

func (d *dataServant) createTag(db *gorm.DB, tag *model.Tag) (*model.Tag, error) {
	t, err := tag.Get(d.engine)
	if err != nil {
		tag.QuoteNum = 1
		return tag.Create(db)
	}

	// 更新
	t.QuoteNum++
	err = t.Update(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (d *dataServant) deleteTag(db *gorm.DB, tag *model.Tag) error {
	tag, err := tag.Get(db)
	if err != nil {
		return err
	}
	tag.QuoteNum--
	return tag.Update(db)
}
