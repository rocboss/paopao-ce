package setting

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Setting struct {
	vp *viper.Viper
}

type LogType string

const LogFileType LogType = "file"
const LogZincType LogType = "zinc"

type LoggerSettingS struct {
	LogType         LogType
	LogFileSavePath string
	LogFileName     string
	LogFileExt      string
	LogZincHost     string
	LogZincIndex    string
	LogZincUser     string
	LogZincPassword string
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
	SmsJuheKey            string
	SmsJuheTplID          string
	SmsJuheTplVal         string
	AlipayAppID           string
	AlipayPrivateKey      string
}

type RuntimeSettingS struct {
	DisablePhoneVerify bool
}

type SearchSettingS struct {
	ZincHost     string
	ZincIndex    string
	ZincUser     string
	ZincPassword string
}

type DatabaseSettingS struct {
	DBType       string
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
