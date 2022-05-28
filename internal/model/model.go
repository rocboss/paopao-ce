package model

import (
	"fmt"
	"time"

	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/soft_delete"
)

// 公共Model
type Model struct {
	ID         int64                 `gorm:"primary_key" json:"id"`
	CreatedOn  int64                 `json:"created_on"`
	ModifiedOn int64                 `json:"modified_on"`
	DeletedOn  int64                 `json:"deleted_on"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_del"`
}

type ConditionsT map[string]interface{}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	newLogger := logger.New(
		global.Logger, // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,              // 慢 SQL 阈值
			LogLevel:                  databaseSetting.LogLevel, // 日志级别
			IgnoreRecordNotFoundError: true,                     // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                    // 禁用彩色打印
		},
	)

	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(databaseSetting.MaxIdleConns).
		SetMaxOpenConns(databaseSetting.MaxOpenConns))

	return db, nil
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	nowTime := time.Now().Unix()

	tx.Statement.SetColumn("created_on", nowTime)
	tx.Statement.SetColumn("modified_on", nowTime)
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if !tx.Statement.Changed("modified_on") {
		tx.Statement.SetColumn("modified_on", time.Now().Unix())
	}

	return
}
