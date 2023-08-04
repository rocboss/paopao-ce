package core

import "github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"

type (
	ShareKey               = dbr.ShareKey
	DeleteKeySuccessOrFail = dbr.DeleteKeySuccessOrFail
)

// share_key服务
type ShareKeyService interface {
	GetUserKeys(UserName string, offset, limit int) ([]*ShareKey, error)
	DeleteUserKey(UserName string, key string) (DeleteKeySuccessOrFail, error)
	GetUserShareKeyCount(name string) (int64, error)
}
