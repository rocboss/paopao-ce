package dao

import "github.com/rocboss/paopao-ce/internal/model"

func (d *dataServant) CreateTag(tag *model.Tag) (*model.Tag, error) {
	t, err := tag.Get(d.engine)
	if err != nil {
		tag.QuoteNum = 1
		return tag.Create(d.engine)
	}

	// 更新
	t.QuoteNum++
	err = t.Update(d.engine)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (d *dataServant) DeleteTag(tag *model.Tag) error {
	tag, err := tag.Get(d.engine)
	if err != nil {
		return err
	}

	tag.QuoteNum--
	return tag.Update(d.engine)
}

func (d *dataServant) GetTags(conditions *model.ConditionsT, offset, limit int) ([]*model.Tag, error) {
	return (&model.Tag{}).List(d.engine, conditions, offset, limit)
}
