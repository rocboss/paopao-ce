package conf

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Setting struct {
	vp *viper.Viper
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

type ServerSettingS struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
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
	MaxIndexPage   int
	ExpireInSecond time.Duration
	Verbose        bool
}

type AlipaySettingS struct {
	AppID      string
	PrivateKey string
}

type SmsJuheSettings struct {
	Gateway string
	Key     string
	TplID   string
	TplVal  string
}

type FeaturesSettingS struct {
	kv       map[string]string
	suites   map[string][]string
	features map[string]string
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

type LocalOSSSettingS struct {
	SavePath string
	Secure   bool
	Bucket   string
	Domain   string
}

type RedisSettingS struct {
	Host     string
	Password string
	DB       int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) Unmarshal(objects map[string]interface{}) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Setting) FeaturesFrom(k string) *FeaturesSettingS {
	sub := s.vp.Sub(k)
	keys := sub.AllKeys()

	suites := make(map[string][]string)
	kv := make(map[string]string, len(keys))
	for _, key := range sub.AllKeys() {
		val := sub.Get(key)
		switch v := val.(type) {
		case string:
			kv[key] = v
		case []interface{}:
			suites[key] = sub.GetStringSlice(key)
		}
	}
	return newFeatures(suites, kv)
}

func newFeatures(suites map[string][]string, kv map[string]string) *FeaturesSettingS {
	features := &FeaturesSettingS{
		suites:   suites,
		kv:       kv,
		features: make(map[string]string),
	}
	features.UseDefault()
	return features
}

func (f *FeaturesSettingS) UseDefault() {
	f.Use([]string{"default"}, true)
}

func (f *FeaturesSettingS) Use(suite []string, noDefault bool) error {
	if noDefault && len(f.features) != 0 {
		f.features = make(map[string]string)
	}
	features := f.flatFeatures(suite)
	for _, feature := range features {
		if len(feature) == 0 {
			continue
		}
		f.features[feature] = f.kv[feature]
	}
	return nil
}

func (f *FeaturesSettingS) flatFeatures(suite []string) []string {
	features := make([]string, 0, len(suite)+10)
	for s := suite[:]; len(s) > 0; s = s[:len(s)-1] {
		item := strings.TrimSpace(strings.ToLower(s[0]))
		if len(item) > 0 {
			if items, exist := f.suites[item]; exist {
				s = append(s, items...)
			}
			features = append(features, item)
		}
		s[0] = s[len(s)-1]
	}
	return features
}

// Cfg get value by key if exist
func (f *FeaturesSettingS) Cfg(key string) (string, bool) {
	key = strings.ToLower(key)
	value, exist := f.features[key]
	return value, exist
}

// CfgIf check expression is true. if expression just have a string like
// `Sms` is mean `Sms` whether define in suite feature settings. expression like
// `Sms = SmsJuhe` is mean whether `Sms` define in suite feature settings and value
// is `SmsJuhe``
func (f *FeaturesSettingS) CfgIf(expression string) bool {
	kv := strings.Split(expression, "=")
	key := strings.Trim(strings.ToLower(kv[0]), " ")
	v, ok := f.features[key]
	if len(kv) == 2 && ok && strings.Trim(kv[1], " ") == v {
		return true
	} else if len(kv) == 1 && ok {
		return true
	}
	return false
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
		if len(v) > 0 {
			params = append(params, strings.ToLower(k)+"="+v)
		}
	}
	return strings.Join(params, " ")
}

func (s *Sqlite3SettingS) Dsn() string {
	return "file:" + s.Path + "?cache=shared&mode=rwc"
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

func (s *ZincSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func (s *MeiliSettingS) Endpoint() string {
	return endpoint(s.Host, s.Secure)
}

func endpoint(host string, secure bool) string {
	schema := "http"
	if secure {
		schema = "https"
	}
	return schema + "://" + host
}
