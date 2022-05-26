package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	Redis    *redis.Client
)
