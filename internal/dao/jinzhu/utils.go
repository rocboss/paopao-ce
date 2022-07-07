package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

func createTag(db *gorm.DB, tag *model.Tag) (*model.Tag, error) {
	t, err := tag.Get(db)
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

func deleteTag(db *gorm.DB, tag *model.Tag) error {
	tag, err := tag.Get(db)
	if err != nil {
		return err
	}
	tag.QuoteNum--
	return tag.Update(db)
}

// 根据IDs获取用户列表
func getUsersByIDs(db *gorm.DB, ids []int64) ([]*model.User, error) {
	user := &model.User{}
	return user.List(db, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}
