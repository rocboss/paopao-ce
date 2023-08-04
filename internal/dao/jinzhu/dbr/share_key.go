package dbr

import "gorm.io/gorm"

type ShareKey struct {
	UserName    string `json:"user_name"`
	ShareKey    string `json:"share_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (s *ShareKey) Get(db *gorm.DB) ([]ShareKey, error) {
	var sks []ShareKey

	if s.UserName != "" {
		db = db.Where("user_name = ? AND status = 0", s.UserName)
	}
	//if s.ShareKey != "" {
	//	db = db.Where("share_key = ?", s.ShareKey)
	//}

	err := db.Table("p_share_key").Find(&sks).Error
	if err != nil {
		return sks, err
	}

	return sks, nil
}

type DeleteKeySuccessOrFail struct {
	SuccessOrFail bool
}

func (sf *ShareKey) POST(db *gorm.DB) (DeleteKeySuccessOrFail, error) {
	var sof DeleteKeySuccessOrFail

	return sof, nil
}

func (s *ShareKey) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*ShareKey, error) {
	var records []*ShareKey
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	db = db.Table("p_share_key").Where("user_name = ?", s.UserName)

	////如果没有传order，则执行过滤条件
	//for k, v := range *conditions {
	//	if k == "ORDER" {
	//		db = db.Order(v)
	//	} else {
	//		db = db.Where(k, v)
	//	}
	//}

	if err = db.Table("p_share_key").Where("status = ?", 0).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
