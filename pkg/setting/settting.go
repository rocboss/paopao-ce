package setting

import (
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Setting struct {
	vp *viper.Viper
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
}

type ServerSettingS struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	BarkToken             string
	MaxCommentCount       int64
	AttachmentIncomeRate  float64
	DefaultContextTimeout time.Duration
	DefaultPageSize       int
	MaxPageSize           int
	IsShastaTestnet       bool
	TronApiKeys           []string
}

type AlipaySettingS struct {
	AppID      string
	PrivateKey string
}

type SmsJuheSettings struct {
	Key    string
	TplID  string
	TplVal string
}

type FeaturesSettingS struct {
	kv       map[string]string
	suites   map[string][]string
	features map[string]string
}

type ZincSettingS struct {
	Host     string
	Index    string
	User     string
	Password string
}

type MySQLSettingS struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	LogLevel     logger.LogLevel
	MaxIdleConns int
	MaxOpenConns int
}

type AliossSettingS struct {
	AliossAccessKeyID     string
	AliossAccessKeySecret string
	AliossEndpoint        string
	AliossBucket          string
	AliossDomain          string
}

type AliOSSSettingS struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Domain          string
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
		feature = strings.ToLower(feature)
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
		if items, exist := f.suites[s[0]]; exist {
			s = append(s, items...)
		}
		features = append(features, s[0])
		s[0] = s[len(s)-1]
	}
	return features
}

// Cfg get value by key if exsit
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
