package conf

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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
		logrus.StandardLogger(), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,                // 慢 SQL 阈值
			LogLevel:                  DatabaseSetting.logLevel(), // 日志级别
			IgnoreRecordNotFoundError: true,                       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                      // 禁用彩色打印
		},
	)

	config := &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
	}

	plugin := dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(MysqlSetting.MaxIdleConns).
		SetMaxOpenConns(MysqlSetting.MaxOpenConns)

	var (
		db  *gorm.DB
		err error
	)
	if CfgIf("MySQL") {
		logrus.Debugln("use MySQL as db")
		if db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config); err == nil {
			db.Use(plugin)
		}
	} else if CfgIf("Postgres") {
		logrus.Debugln("use PostgreSQL as db")
		db, err = gorm.Open(postgres.Open(PostgresSetting.Dsn()), config)
	} else if CfgIf("Sqlite3") {
		logrus.Debugf("use Sqlite3 as db dsn: %s", Sqlite3Setting.Dsn())
		db, err = gorm.Open(sqlite.Open(Sqlite3Setting.Dsn()), config)
	} else {
		logrus.Debugln("use default of MySQL as db")
		if db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config); err == nil {
			db.Use(plugin)
		}
	}

	return db, err
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
