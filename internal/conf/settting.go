// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"embed"
	"fmt"
	"strings"
	"time"

	"github.com/pyroscope-io/client/pyroscope"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

//go:embed config.yaml
var files embed.FS

type Setting struct {
	vp *viper.Viper
}

type PyroscopeSettingS struct {
	AppName   string
	Endpoint  string
	AuthToken string
	Logger    string
}

type LoggerSettingS struct {
	Level string
}

type LoggerFileSettingS struct {
	SavePath string
	FileName string
	FileExt  string
}

type LoggerZincSettingS struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool
}

type LoggerMeiliSettingS struct {
	Host         string
	Index        string
	ApiKey       string
	Secure       bool
	MaxLogBuffer int
	MinWorker    int
}

type HttpServerSettingS struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GRPCServerSettingS struct {
	Host string
	Port string
}

type AppSettingS struct {
	RunMode               string
	MaxCommentCount       int64
	AttachmentIncomeRate  float64
	DefaultContextTimeout time.Duration
	DefaultPageSize       int
	MaxPageSize           int
}

type CacheIndexSettingS struct {
	MaxUpdateQPS int
	MinWorker    int
}

type SimpleCacheIndexSettingS struct {
	MaxIndexSize       int
	CheckTickDuration  time.Duration
	ExpireTickDuration time.Duration
}

type BigCacheIndexSettingS struct {
	MaxIndexPage     int
	HardMaxCacheSize int
	ExpireInSecond   time.Duration
	Verbose          bool
}

type RedisCacheIndexSettingS struct {
	ExpireInSecond time.Duration
	Verbose        bool
}

type AlipaySettingS struct {
	AppID             string
	PrivateKey        string
	RootCertFile      string
	PublicCertFile    string
	AppPublicCertFile string
	InProduction      bool
}

type SmsJuheSettings struct {
	Gateway string
	Key     string
	TplID   string
	TplVal  string
}

type TweetSearchS struct {
	MaxUpdateQPS int
	MinWorker    int
}

type ZincSettingS struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool
}

type MeiliSettingS struct {
	Host   string
	Index  string
	ApiKey string
	Secure bool
}

type DatabaseSetingS struct {
	TablePrefix string
	LogLevel    string
}

type MySQLSettingS struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type PostgresSettingS map[string]string

type Sqlite3SettingS struct {
	Path string
}

type ObjectStorageS struct {
	RetainInDays int
	TempDir      string
}

type MinIOSettingS struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type S3SettingS struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type AliOSSSettingS struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Domain          string
}

type COSSettingS struct {
	SecretID  string
	SecretKey string
	Region    string
	Bucket    string
	Domain    string
}

type HuaweiOBSSettingS struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Domain    string
}

type LocalOSSSettingS struct {
	SavePath string
	Secure   bool
	Bucket   string
	Domain   string
}

type RedisSettingS struct {
	InitAddress      []string
	Username         string
	Password         string
	SelectDB         int
	ConnWriteTimeout time.Duration
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func NewSetting() (*Setting, error) {
	cfgFile, err := files.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.AddConfigPath("custom/")
	vp.SetConfigType("yaml")
	if err = vp.ReadConfig(cfgFile); err != nil {
		return nil, err
	}
	if err = vp.MergeInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *Setting) Unmarshal(objects map[string]any) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Setting) featuresInfoFrom(k string) (map[string][]string, map[string]string) {
	sub := s.vp.Sub(k)
	keys := sub.AllKeys()

	suites := make(map[string][]string)
	kv := make(map[string]string, len(keys))
	for _, key := range sub.AllKeys() {
		val := sub.Get(key)
		switch v := val.(type) {
		case string:
			kv[key] = v
		case []any:
			suites[key] = sub.GetStringSlice(key)
		}
	}
	return suites, kv
}

func (s *HttpServerSettingS) GetReadTimeout() time.Duration {
	return s.ReadTimeout * time.Second
}

func (s *HttpServerSettingS) GetWriteTimeout() time.Duration {
	return s.WriteTimeout * time.Second
}

func (s *MySQLSettingS) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.UserName,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime,
	)
}

func (s PostgresSettingS) Dsn() string {
	var params []string
	for k, v := range s {
		if len(v) == 0 {
			continue
		}
		lk := strings.ToLower(k)
		tv := strings.Trim(v, " ")
		switch lk {
		case "schema":
			params = append(params, "search_path="+tv)
		case "applicationname":
			params = append(params, "application_name="+tv)
		default:
			params = append(params, lk+"="+tv)
		}
	}
	return strings.Join(params, " ")
}

func (s *Sqlite3SettingS) Dsn(driverName string) string {
	pragmas := "_foreign_keys=1&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=8000"
	if driverName == "sqlite" {
		pragmas = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	}
	return fmt.Sprintf("file:%s?%s", s.Path, pragmas)
}

func (s *DatabaseSetingS) logLevel() logger.LogLevel {
	switch strings.ToLower(s.LogLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Error
	}
}

func (s *LoggerSettingS) logLevel() logrus.Level {
	switch strings.ToLower(s.Level) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}

func (s *LoggerZincSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func (s *LoggerMeiliSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func (s *LoggerMeiliSettingS) minWork() int {
	if s.MinWorker < 5 {
		return 5
	} else if s.MinWorker > 100 {
		return 100
	}
	return s.MinWorker
}

func (s *LoggerMeiliSettingS) maxLogBuffer() int {
	if s.MaxLogBuffer < 10 {
		return 10
	} else if s.MaxLogBuffer > 1000 {
		return 1000
	}
	return s.MaxLogBuffer
}

func (s *ObjectStorageS) TempDirSlash() string {
	return strings.Trim(s.TempDir, " /") + "/"
}

func (s *ZincSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func (s *MeiliSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func (s *DatabaseSetingS) TableNames() (res TableNameMap) {
	tableNames := []string{
		TableAnouncement,
		TableAnouncementContent,
		TableAttachment,
		TableCaptcha,
		TableComment,
		TableCommentContent,
		TableCommentReply,
		TableContact,
		TableContactGroup,
		TableMessage,
		TablePost,
		TablePostAttachmentBill,
		TablePostCollection,
		TablePostContent,
		TablePostStar,
		TableTag,
		TableUser,
		TableWalletRecharge,
		TableWalletStatement,
	}
	res = make(TableNameMap, len(tableNames))
	for _, name := range tableNames {
		res[name] = s.TablePrefix + name
	}
	return
}

func (s *PyroscopeSettingS) GetLogger() (logger pyroscope.Logger) {
	switch strings.ToLower(s.Logger) {
	case "standard":
		logger = pyroscope.StandardLogger
	case "logrus":
		logger = logrus.StandardLogger()
	}
	return
}

func endpoint(host string, secure bool) string {
	schema := "http"
	if secure {
		schema = "https"
	}
	return schema + "://" + host
}
