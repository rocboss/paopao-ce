package conf

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	DBEngine *gorm.DB
	Redis    *redis.Client
)

func newDBEngine() (*gorm.DB, error) {
	newLogger := logger.New(
		Logger, // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,           // 慢 SQL 阈值
			LogLevel:                  mySQLSetting.LogLevel, // 日志级别
			IgnoreRecordNotFoundError: true,                  // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                 // 禁用彩色打印
		},
	)

	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		mySQLSetting.UserName,
		mySQLSetting.Password,
		mySQLSetting.Host,
		mySQLSetting.DBName,
		mySQLSetting.Charset,
		mySQLSetting.ParseTime,
	)), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mySQLSetting.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(mySQLSetting.MaxIdleConns).
		SetMaxOpenConns(mySQLSetting.MaxOpenConns))

	return db, nil
}

// setupDBEngine 暂时只支持MySQL
func setupDBEngine() error {
	var err error
	DBEngine, err = newDBEngine()
	if err != nil {
		return err
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     redisSetting.Host,
		Password: redisSetting.Password,
		DB:       redisSetting.DB,
	})

	return nil
}
