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
