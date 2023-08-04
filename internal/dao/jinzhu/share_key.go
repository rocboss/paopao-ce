package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
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

func (s ShareKeyService) GetUserKeys(UserName string, offset, limit int) ([]*core.ShareKey, error) {
	statement := &dbr.ShareKey{
		UserName: UserName,
	}
	//fmt.Print(statement.List(s.db, &dbr.ConditionsT{}, offset, limit))
	return statement.List(s.db, &dbr.ConditionsT{}, offset, limit)
}

func (s ShareKeyService) GetUserShareKeyCount(name string) (int64, error) {
	var count int64
	err := s.db.Table("p_share_key").Where("user_name = ? and status = 0", name).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewShareKeyService(db *gorm.DB) *ShareKeyService {
	return &ShareKeyService{
		db: db,
	}
}
