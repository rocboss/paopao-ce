package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"gorm.io/gorm"
)

type ShareKeyService struct {
	db *gorm.DB
}

func (s ShareKeyService) DeleteUserKey(UserName string, key string) (core.DeleteKeySuccessOrFail, error) {
	var deleteKey core.DeleteKeySuccessOrFail
	deleteKey.SuccessOrFail = false
	//根据username和key删除相应的key
	err := s.db.Table("p_share_key").Where("user_name = ? AND share_key = ?", UserName, key).Update("status", 1).Error
	if err != nil {
		return deleteKey, err
	}
	deleteKey.SuccessOrFail = true
	return deleteKey, nil
}

func (s ShareKeyService) GetUserKeys(UserName string) ([]*core.ShareKey, error) {
	//根据username查询相应的keys
	var keys []*core.ShareKey
	err := s.db.Table("p_share_key").Where("user_name = ? and status = 0", UserName).Find(&keys).Error
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func NewShareKeyService(db *gorm.DB) *ShareKeyService {
	return &ShareKeyService{
		db: db,
	}
}
