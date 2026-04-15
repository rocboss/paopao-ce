// Copyright 2026 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package settings

import (
	stdjson "encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/alimy/tryst/cfg"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

type (
	ValueType string
	ApplyMode string
)

const (
	TypeBool   ValueType = "bool"
	TypeInt    ValueType = "int"
	TypeFloat  ValueType = "float"
	TypeString ValueType = "string"

	ApplyModeLive            ApplyMode = "live"
	ApplyModeRestartRequired ApplyMode = "restart_required"
	ApplyModeBootstrapOnly   ApplyMode = "bootstrap_only"
)

type bootstrapSnapshot struct {
	App           confAppSnapshot
	TweetSearch   confTweetSearchSnapshot
	Zinc          confZincSnapshot
	Meili         confMeiliSnapshot
	ObjectStorage confObjectStorageSnapshot
	AliOSS        confAliOSSSnapshot
	COS           confCOSSnapshot
	HuaweiOBS     confHuaweiOBSSnapshot
	MinIO         confMinIOSnapshot
	S3            confS3Snapshot
	LocalOSS      confLocalOSSSnapshot
	SmsJuhe       confSmsJuheSnapshot
	Alipay        confAlipaySnapshot
	WebProfile    confWebProfileSnapshot
}

type confAppSnapshot struct {
	AttachmentIncomeRate float64
	MaxCommentCount      int64
	MaxWhisperDaily      int64
	MaxCaptchaTimes      int
	DefaultPageSize      int
	MaxPageSize          int
}

type confTweetSearchSnapshot struct {
	MaxUpdateQPS int
	MinWorker    int
}

type confZincSnapshot struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool
}

type confMeiliSnapshot struct {
	Host   string
	Index  string
	ApiKey string
	Secure bool
}

type confObjectStorageSnapshot struct {
	RetainInDays int
	TempDir      string
}

type confAliOSSSnapshot struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Domain          string
}

type confCOSSnapshot struct {
	SecretID  string
	SecretKey string
	Region    string
	Bucket    string
	Domain    string
}

type confHuaweiOBSSnapshot struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Domain    string
}

type confMinIOSnapshot struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type confS3Snapshot struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type confLocalOSSSnapshot struct {
	SavePath string
	Secure   bool
	Bucket   string
	Domain   string
}

type confSmsJuheSnapshot struct {
	Gateway string
	Key     string
	TplID   string
	TplVal  string
}

type confAlipaySnapshot struct {
	AppID             string
	PrivateKey        string
	RootCertFile      string
	PublicCertFile    string
	AppPublicCertFile string
	InProduction      bool
}

type confWebProfileSnapshot struct {
	UseFriendship             bool
	EnableTrendsBar           bool
	EnableWallet              bool
	AllowTweetAttachment      bool
	AllowTweetAttachmentPrice bool
	AllowTweetVideo           bool
	AllowUserRegister         bool
	AllowPhoneBind            bool
	DefaultTweetMaxLength     int
	TweetWebEllipsisSize      int
	TweetMobileEllipsisSize   int
	DefaultTweetVisibility    string
	DefaultMsgLoopInterval    int
	CopyrightTop              string
	CopyrightLeft             string
	CopyrightLeftLink         string
	CopyrightRight            string
	CopyrightRightLink        string
}

type Option struct {
	Label string `json:"label"`
	Value any    `json:"value"`
}

type Definition struct {
	Key               string
	Group             string
	Section           string
	Type              ValueType
	Label             string
	Description       string
	ApplyMode         ApplyMode
	Secret            bool
	Readonly          bool
	Options           []Option
	Active            func() bool
	CurrentValue      func() any
	BootstrapDefault  func() any
	ValidateValue     func(any) (any, error)
	BootstrapOverride func(any)
}

var (
	ReadonlyFields   = []string{"allow_user_register", "allow_phone_bind"}
	bootstrapConfig  *bootstrapSnapshot
	visibilityOption = []Option{{Label: "Public", Value: "public"}, {Label: "Following", Value: "following"}, {Label: "Friend", Value: "friend"}, {Label: "Private", Value: "private"}}
)

func ensureBootstrapSnapshot() {
	if bootstrapConfig != nil {
		return
	}
	bootstrapConfig = &bootstrapSnapshot{}
	if conf.AppSetting != nil {
		bootstrapConfig.App = confAppSnapshot{
			AttachmentIncomeRate: conf.AppSetting.AttachmentIncomeRate,
			MaxCommentCount:      conf.AppSetting.MaxCommentCount,
			MaxWhisperDaily:      conf.AppSetting.MaxWhisperDaily,
			MaxCaptchaTimes:      conf.AppSetting.MaxCaptchaTimes,
			DefaultPageSize:      conf.AppSetting.DefaultPageSize,
			MaxPageSize:          conf.AppSetting.MaxPageSize,
		}
	}
	if conf.TweetSearchSetting != nil {
		bootstrapConfig.TweetSearch = confTweetSearchSnapshot{MaxUpdateQPS: conf.TweetSearchSetting.MaxUpdateQPS, MinWorker: conf.TweetSearchSetting.MinWorker}
	}
	if conf.ZincSetting != nil {
		bootstrapConfig.Zinc = confZincSnapshot{Host: conf.ZincSetting.Host, Index: conf.ZincSetting.Index, User: conf.ZincSetting.User, Password: conf.ZincSetting.Password, Secure: conf.ZincSetting.Secure}
	}
	if conf.MeiliSetting != nil {
		bootstrapConfig.Meili = confMeiliSnapshot{Host: conf.MeiliSetting.Host, Index: conf.MeiliSetting.Index, ApiKey: conf.MeiliSetting.ApiKey, Secure: conf.MeiliSetting.Secure}
	}
	if conf.ObjectStorage != nil {
		bootstrapConfig.ObjectStorage = confObjectStorageSnapshot{RetainInDays: conf.ObjectStorage.RetainInDays, TempDir: conf.ObjectStorage.TempDir}
	}
	if conf.AliOSSSetting != nil {
		bootstrapConfig.AliOSS = confAliOSSSnapshot{Endpoint: conf.AliOSSSetting.Endpoint, AccessKeyID: conf.AliOSSSetting.AccessKeyID, AccessKeySecret: conf.AliOSSSetting.AccessKeySecret, Bucket: conf.AliOSSSetting.Bucket, Domain: conf.AliOSSSetting.Domain}
	}
	if conf.COSSetting != nil {
		bootstrapConfig.COS = confCOSSnapshot{SecretID: conf.COSSetting.SecretID, SecretKey: conf.COSSetting.SecretKey, Region: conf.COSSetting.Region, Bucket: conf.COSSetting.Bucket, Domain: conf.COSSetting.Domain}
	}
	if conf.HuaweiOBSSetting != nil {
		bootstrapConfig.HuaweiOBS = confHuaweiOBSSnapshot{AccessKey: conf.HuaweiOBSSetting.AccessKey, SecretKey: conf.HuaweiOBSSetting.SecretKey, Endpoint: conf.HuaweiOBSSetting.Endpoint, Bucket: conf.HuaweiOBSSetting.Bucket, Domain: conf.HuaweiOBSSetting.Domain}
	}
	if conf.MinIOSetting != nil {
		bootstrapConfig.MinIO = confMinIOSnapshot{AccessKey: conf.MinIOSetting.AccessKey, SecretKey: conf.MinIOSetting.SecretKey, Secure: conf.MinIOSetting.Secure, Endpoint: conf.MinIOSetting.Endpoint, Bucket: conf.MinIOSetting.Bucket, Domain: conf.MinIOSetting.Domain}
	}
	if conf.S3Setting != nil {
		bootstrapConfig.S3 = confS3Snapshot{AccessKey: conf.S3Setting.AccessKey, SecretKey: conf.S3Setting.SecretKey, Secure: conf.S3Setting.Secure, Endpoint: conf.S3Setting.Endpoint, Bucket: conf.S3Setting.Bucket, Domain: conf.S3Setting.Domain}
	}
	if conf.LocalOSSSetting != nil {
		bootstrapConfig.LocalOSS = confLocalOSSSnapshot{SavePath: conf.LocalOSSSetting.SavePath, Secure: conf.LocalOSSSetting.Secure, Bucket: conf.LocalOSSSetting.Bucket, Domain: conf.LocalOSSSetting.Domain}
	}
	if conf.SmsJuheSetting != nil {
		bootstrapConfig.SmsJuhe = confSmsJuheSnapshot{Gateway: conf.SmsJuheSetting.Gateway, Key: conf.SmsJuheSetting.Key, TplID: conf.SmsJuheSetting.TplID, TplVal: conf.SmsJuheSetting.TplVal}
	}
	if conf.AlipaySetting != nil {
		bootstrapConfig.Alipay = confAlipaySnapshot{AppID: conf.AlipaySetting.AppID, PrivateKey: conf.AlipaySetting.PrivateKey, RootCertFile: conf.AlipaySetting.RootCertFile, PublicCertFile: conf.AlipaySetting.PublicCertFile, AppPublicCertFile: conf.AlipaySetting.AppPublicCertFile, InProduction: conf.AlipaySetting.InProduction}
	}
	if conf.WebProfileSetting != nil {
		bootstrapConfig.WebProfile = confWebProfileSnapshot{
			UseFriendship:             conf.WebProfileSetting.UseFriendship,
			EnableTrendsBar:           conf.WebProfileSetting.EnableTrendsBar,
			EnableWallet:              conf.WebProfileSetting.EnableWallet,
			AllowTweetAttachment:      conf.WebProfileSetting.AllowTweetAttachment,
			AllowTweetAttachmentPrice: conf.WebProfileSetting.AllowTweetAttachmentPrice,
			AllowTweetVideo:           conf.WebProfileSetting.AllowTweetVideo,
			AllowUserRegister:         conf.WebProfileSetting.AllowUserRegister,
			AllowPhoneBind:            conf.WebProfileSetting.AllowPhoneBind,
			DefaultTweetMaxLength:     conf.WebProfileSetting.DefaultTweetMaxLength,
			TweetWebEllipsisSize:      conf.WebProfileSetting.TweetWebEllipsisSize,
			TweetMobileEllipsisSize:   conf.WebProfileSetting.TweetMobileEllipsisSize,
			DefaultTweetVisibility:    conf.WebProfileSetting.DefaultTweetVisibility,
			DefaultMsgLoopInterval:    conf.WebProfileSetting.DefaultMsgLoopInterval,
			CopyrightTop:              conf.WebProfileSetting.CopyrightTop,
			CopyrightLeft:             conf.WebProfileSetting.CopyrightLeft,
			CopyrightLeftLink:         conf.WebProfileSetting.CopyrightLeftLink,
			CopyrightRight:            conf.WebProfileSetting.CopyrightRight,
			CopyrightRightLink:        conf.WebProfileSetting.CopyrightRightLink,
		}
	}
}

func CloneReadonlyFields() []string {
	res := make([]string, len(ReadonlyFields))
	copy(res, ReadonlyFields)
	return res
}

func Registry() []Definition {
	ensureBootstrapSnapshot()
	return []Definition{
		boolDef("web_profile.use_friendship", "web", "profile", "Use friendship", "Switch the frontend friendship model.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.UseFriendship }, func() any { return bootstrapConfig.WebProfile.UseFriendship }, func(v any) { conf.WebProfileSetting.UseFriendship = v.(bool) }),
		boolDef("web_profile.enable_trends_bar", "web", "profile", "Enable trends bar", "Show the trends sidebar in the web UI.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.EnableTrendsBar }, func() any { return bootstrapConfig.WebProfile.EnableTrendsBar }, func(v any) { conf.WebProfileSetting.EnableTrendsBar = v.(bool) }),
		boolDef("web_profile.enable_wallet", "web", "profile", "Enable wallet", "Enable wallet-related frontend features.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.EnableWallet }, func() any { return bootstrapConfig.WebProfile.EnableWallet }, func(v any) { conf.WebProfileSetting.EnableWallet = v.(bool) }),
		boolDef("web_profile.allow_tweet_attachment", "web", "profile", "Allow attachments", "Allow file attachments on posts.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.AllowTweetAttachment }, func() any { return bootstrapConfig.WebProfile.AllowTweetAttachment }, func(v any) { conf.WebProfileSetting.AllowTweetAttachment = v.(bool) }),
		boolDef("web_profile.allow_tweet_attachment_price", "web", "profile", "Allow paid attachments", "Allow post attachments to have a price.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.AllowTweetAttachmentPrice }, func() any { return bootstrapConfig.WebProfile.AllowTweetAttachmentPrice }, func(v any) { conf.WebProfileSetting.AllowTweetAttachmentPrice = v.(bool) }),
		boolDef("web_profile.allow_tweet_video", "web", "profile", "Allow video posts", "Allow video uploads on posts.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.AllowTweetVideo }, func() any { return bootstrapConfig.WebProfile.AllowTweetVideo }, func(v any) { conf.WebProfileSetting.AllowTweetVideo = v.(bool) }),
		boolDef("web_profile.allow_user_register", "web", "profile", "Allow user registration", "Bootstrap-only registration gate from YAML/features.", ApplyModeBootstrapOnly, false, false, func() any { return conf.WebProfileSetting.AllowUserRegister }, func() any { return bootstrapConfig.WebProfile.AllowUserRegister }, nil),
		boolDef("web_profile.allow_phone_bind", "web", "profile", "Allow phone binding", "Bootstrap-only phone binding gate from YAML/features.", ApplyModeBootstrapOnly, false, false, func() any { return conf.WebProfileSetting.AllowPhoneBind }, func() any { return bootstrapConfig.WebProfile.AllowPhoneBind }, nil),
		intDef("web_profile.default_tweet_max_length", "web", "profile", "Default tweet max length", "Maximum allowed tweet length.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.DefaultTweetMaxLength }, func() any { return bootstrapConfig.WebProfile.DefaultTweetMaxLength }, func(v int) error { return between(v, 1, 2000, "default_tweet_max_length") }, func(v any) { conf.WebProfileSetting.DefaultTweetMaxLength = v.(int) }),
		intDef("web_profile.tweet_web_ellipsis_size", "web", "profile", "Web ellipsis size", "Truncated feed length on web.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.TweetWebEllipsisSize }, func() any { return bootstrapConfig.WebProfile.TweetWebEllipsisSize }, func(v int) error {
			return between(v, 1, conf.WebProfileSetting.DefaultTweetMaxLength, "tweet_web_ellipsis_size")
		}, func(v any) { conf.WebProfileSetting.TweetWebEllipsisSize = v.(int) }),
		intDef("web_profile.tweet_mobile_ellipsis_size", "web", "profile", "Mobile ellipsis size", "Truncated feed length on mobile.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.TweetMobileEllipsisSize }, func() any { return bootstrapConfig.WebProfile.TweetMobileEllipsisSize }, func(v int) error {
			return between(v, 1, conf.WebProfileSetting.DefaultTweetMaxLength, "tweet_mobile_ellipsis_size")
		}, func(v any) { conf.WebProfileSetting.TweetMobileEllipsisSize = v.(int) }),
		stringDef("web_profile.default_tweet_visibility", "web", "profile", "Default tweet visibility", "Default visibility for new posts.", ApplyModeLive, false, true, visibilityOption, func() any { return conf.WebProfileSetting.DefaultTweetVisibility }, func() any { return bootstrapConfig.WebProfile.DefaultTweetVisibility }, validateVisibility, func(v any) { conf.WebProfileSetting.DefaultTweetVisibility = v.(string) }),
		intDef("web_profile.default_msg_loop_interval", "web", "profile", "Message polling interval", "Polling interval for unread messages in milliseconds.", ApplyModeLive, false, true, func() any { return conf.WebProfileSetting.DefaultMsgLoopInterval }, func() any { return bootstrapConfig.WebProfile.DefaultMsgLoopInterval }, func(v int) error { return between(v, 1000, 60000, "default_msg_loop_interval") }, func(v any) { conf.WebProfileSetting.DefaultMsgLoopInterval = v.(int) }),
		stringDef("web_profile.copyright_top", "web", "profile", "Copyright top", "Top copyright line displayed in the UI.", ApplyModeLive, false, true, nil, func() any { return conf.WebProfileSetting.CopyrightTop }, func() any { return bootstrapConfig.WebProfile.CopyrightTop }, validateRequiredTrimmed("copyright_top", 255), func(v any) { conf.WebProfileSetting.CopyrightTop = v.(string) }),
		stringDef("web_profile.copyright_left", "web", "profile", "Copyright left", "Left footer label.", ApplyModeLive, false, true, nil, func() any { return conf.WebProfileSetting.CopyrightLeft }, func() any { return bootstrapConfig.WebProfile.CopyrightLeft }, validateRequiredTrimmed("copyright_left", 255), func(v any) { conf.WebProfileSetting.CopyrightLeft = v.(string) }),
		stringDef("web_profile.copyright_left_link", "web", "profile", "Copyright left link", "Optional left footer link.", ApplyModeLive, false, true, nil, func() any { return conf.WebProfileSetting.CopyrightLeftLink }, func() any { return bootstrapConfig.WebProfile.CopyrightLeftLink }, validateOptionalURL("copyright_left_link", 255), func(v any) { conf.WebProfileSetting.CopyrightLeftLink = v.(string) }),
		stringDef("web_profile.copyright_right", "web", "profile", "Copyright right", "Right footer label.", ApplyModeLive, false, true, nil, func() any { return conf.WebProfileSetting.CopyrightRight }, func() any { return bootstrapConfig.WebProfile.CopyrightRight }, validateRequiredTrimmed("copyright_right", 255), func(v any) { conf.WebProfileSetting.CopyrightRight = v.(string) }),
		stringDef("web_profile.copyright_right_link", "web", "profile", "Copyright right link", "Optional right footer link.", ApplyModeLive, false, true, nil, func() any { return conf.WebProfileSetting.CopyrightRightLink }, func() any { return bootstrapConfig.WebProfile.CopyrightRightLink }, validateOptionalURL("copyright_right_link", 255), func(v any) { conf.WebProfileSetting.CopyrightRightLink = v.(string) }),

		int64Def("app.max_comment_count", "app", "general", "Max comment count", "Maximum comments allowed on a post.", ApplyModeLive, func() any { return conf.AppSetting.MaxCommentCount }, func() any { return bootstrapConfig.App.MaxCommentCount }, func(v int64) error { return betweenInt64(v, 1, 100000, "max_comment_count") }, func(v any) { conf.AppSetting.MaxCommentCount = v.(int64) }),
		floatDef("app.attachment_income_rate", "app", "general", "Attachment income rate", "Revenue share for paid attachments.", ApplyModeLive, func() any { return conf.AppSetting.AttachmentIncomeRate }, func() any { return bootstrapConfig.App.AttachmentIncomeRate }, func(v float64) error {
			if v < 0 || v > 1 {
				return xerror.InvalidParams.WithDetails("attachment_income_rate must be between 0 and 1")
			}
			return nil
		}, func(v any) { conf.AppSetting.AttachmentIncomeRate = v.(float64) }),
		intDef("app.default_page_size", "app", "general", "Default page size", "Default pagination size.", ApplyModeLive, false, true, func() any { return conf.AppSetting.DefaultPageSize }, func() any { return bootstrapConfig.App.DefaultPageSize }, func(v int) error { return between(v, 1, conf.AppSetting.MaxPageSize, "default_page_size") }, func(v any) { conf.AppSetting.DefaultPageSize = v.(int) }),
		intDef("app.max_page_size", "app", "general", "Max page size", "Maximum pagination size.", ApplyModeLive, false, true, func() any { return conf.AppSetting.MaxPageSize }, func() any { return bootstrapConfig.App.MaxPageSize }, func(v int) error {
			return between(v, maxInt(conf.AppSetting.DefaultPageSize, 1), 1000, "max_page_size")
		}, func(v any) { conf.AppSetting.MaxPageSize = v.(int) }),
		int64Def("app.max_whisper_daily", "app", "limits", "Max daily whispers", "Daily whisper send limit. Restart required because servants cache it at startup.", ApplyModeRestartRequired, func() any { return conf.AppSetting.MaxWhisperDaily }, func() any { return bootstrapConfig.App.MaxWhisperDaily }, func(v int64) error { return betweenInt64(v, 1, 1000000, "max_whisper_daily") }, func(v any) { conf.AppSetting.MaxWhisperDaily = v.(int64) }),
		intDef("app.max_captcha_times", "app", "limits", "Max captcha times", "Max captcha request count cached at startup by web servants.", ApplyModeRestartRequired, false, true, func() any { return conf.AppSetting.MaxCaptchaTimes }, func() any { return bootstrapConfig.App.MaxCaptchaTimes }, func(v int) error { return between(v, 1, 1000, "max_captcha_times") }, func(v any) { conf.AppSetting.MaxCaptchaTimes = v.(int) }),

		intDef("tweet_search.max_update_qps", "search", "bridge", "Search update QPS", "Buffered update throughput for search indexing.", ApplyModeRestartRequired, false, true, func() any { return conf.TweetSearchSetting.MaxUpdateQPS }, func() any { return bootstrapConfig.TweetSearch.MaxUpdateQPS }, func(v int) error { return between(v, 10, 10000, "max_update_qps") }, func(v any) { conf.TweetSearchSetting.MaxUpdateQPS = v.(int) }),
		intDef("tweet_search.min_worker", "search", "bridge", "Search worker count", "Background worker count for search indexing.", ApplyModeRestartRequired, false, true, func() any { return conf.TweetSearchSetting.MinWorker }, func() any { return bootstrapConfig.TweetSearch.MinWorker }, func(v int) error { return between(v, 5, 1000, "min_worker") }, func(v any) { conf.TweetSearchSetting.MinWorker = v.(int) }),

		stringDefWithActive("meili.host", "search", "meili", "Meili host", "Meilisearch host.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Meili") }, func() any { return conf.MeiliSetting.Host }, func() any { return bootstrapConfig.Meili.Host }, validateTrimmedMax("meili.host", 255), func(v any) { conf.MeiliSetting.Host = v.(string) }),
		stringDefWithActive("meili.index", "search", "meili", "Meili index", "Meilisearch index name.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Meili") }, func() any { return conf.MeiliSetting.Index }, func() any { return bootstrapConfig.Meili.Index }, validateTrimmedMax("meili.index", 255), func(v any) { conf.MeiliSetting.Index = v.(string) }),
		stringDefWithActive("meili.api_key", "search", "meili", "Meili API key", "Meilisearch API key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("Meili") }, func() any { return conf.MeiliSetting.ApiKey }, func() any { return bootstrapConfig.Meili.ApiKey }, validateTrimmedMax("meili.api_key", 512), func(v any) { conf.MeiliSetting.ApiKey = v.(string) }),
		boolDefWithActive("meili.secure", "search", "meili", "Meili secure", "Use HTTPS for Meilisearch.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("Meili") }, func() any { return conf.MeiliSetting.Secure }, func() any { return bootstrapConfig.Meili.Secure }, func(v any) { conf.MeiliSetting.Secure = v.(bool) }),

		stringDefWithActive("zinc.host", "search", "zinc", "Zinc host", "Zinc host.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Zinc") }, func() any { return conf.ZincSetting.Host }, func() any { return bootstrapConfig.Zinc.Host }, validateTrimmedMax("zinc.host", 255), func(v any) { conf.ZincSetting.Host = v.(string) }),
		stringDefWithActive("zinc.index", "search", "zinc", "Zinc index", "Zinc index name.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Zinc") }, func() any { return conf.ZincSetting.Index }, func() any { return bootstrapConfig.Zinc.Index }, validateTrimmedMax("zinc.index", 255), func(v any) { conf.ZincSetting.Index = v.(string) }),
		stringDefWithActive("zinc.user", "search", "zinc", "Zinc user", "Zinc username.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Zinc") }, func() any { return conf.ZincSetting.User }, func() any { return bootstrapConfig.Zinc.User }, validateTrimmedMax("zinc.user", 255), func(v any) { conf.ZincSetting.User = v.(string) }),
		stringDefWithActive("zinc.password", "search", "zinc", "Zinc password", "Zinc password.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("Zinc") }, func() any { return conf.ZincSetting.Password }, func() any { return bootstrapConfig.Zinc.Password }, validateTrimmedMax("zinc.password", 512), func(v any) { conf.ZincSetting.Password = v.(string) }),
		boolDefWithActive("zinc.secure", "search", "zinc", "Zinc secure", "Use HTTPS for Zinc.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("Zinc") }, func() any { return conf.ZincSetting.Secure }, func() any { return bootstrapConfig.Zinc.Secure }, func(v any) { conf.ZincSetting.Secure = v.(bool) }),

		intDef("object_storage.retain_in_days", "storage", "common", "Object retention days", "Retention window for temporary objects.", ApplyModeRestartRequired, false, true, func() any { return conf.ObjectStorage.RetainInDays }, func() any { return bootstrapConfig.ObjectStorage.RetainInDays }, func(v int) error { return between(v, 0, 3650, "retain_in_days") }, func(v any) { conf.ObjectStorage.RetainInDays = v.(int) }),
		stringDef("object_storage.temp_dir", "storage", "common", "Object temp dir", "Temporary directory/prefix for objects.", ApplyModeRestartRequired, false, true, nil, func() any { return conf.ObjectStorage.TempDir }, func() any { return bootstrapConfig.ObjectStorage.TempDir }, validateTrimmedMax("temp_dir", 255), func(v any) { conf.ObjectStorage.TempDir = v.(string) }),

		stringDefWithActive("local_oss.save_path", "storage", "local_oss", "Local OSS save path", "Filesystem path for local object storage.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("LocalOSS") }, func() any { return conf.LocalOSSSetting.SavePath }, func() any { return bootstrapConfig.LocalOSS.SavePath }, validateTrimmedMax("local_oss.save_path", 1024), func(v any) { conf.LocalOSSSetting.SavePath = v.(string) }),
		boolDefWithActive("local_oss.secure", "storage", "local_oss", "Local OSS secure", "Generate HTTPS URLs for local object storage.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("LocalOSS") }, func() any { return conf.LocalOSSSetting.Secure }, func() any { return bootstrapConfig.LocalOSS.Secure }, func(v any) { conf.LocalOSSSetting.Secure = v.(bool) }),
		stringDefWithActive("local_oss.bucket", "storage", "local_oss", "Local OSS bucket", "Bucket folder name for local storage.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("LocalOSS") }, func() any { return conf.LocalOSSSetting.Bucket }, func() any { return bootstrapConfig.LocalOSS.Bucket }, validateTrimmedMax("local_oss.bucket", 255), func(v any) { conf.LocalOSSSetting.Bucket = v.(string) }),
		stringDefWithActive("local_oss.domain", "storage", "local_oss", "Local OSS domain", "Public domain for local object storage.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("LocalOSS") }, func() any { return conf.LocalOSSSetting.Domain }, func() any { return bootstrapConfig.LocalOSS.Domain }, validateTrimmedMax("local_oss.domain", 255), func(v any) { conf.LocalOSSSetting.Domain = v.(string) }),

		stringDefWithActive("minio.access_key", "storage", "minio", "MinIO access key", "MinIO access key.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.AccessKey }, func() any { return bootstrapConfig.MinIO.AccessKey }, validateTrimmedMax("minio.access_key", 255), func(v any) { conf.MinIOSetting.AccessKey = v.(string) }),
		stringDefWithActive("minio.secret_key", "storage", "minio", "MinIO secret key", "MinIO secret key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.SecretKey }, func() any { return bootstrapConfig.MinIO.SecretKey }, validateTrimmedMax("minio.secret_key", 512), func(v any) { conf.MinIOSetting.SecretKey = v.(string) }),
		boolDefWithActive("minio.secure", "storage", "minio", "MinIO secure", "Use HTTPS for MinIO.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.Secure }, func() any { return bootstrapConfig.MinIO.Secure }, func(v any) { conf.MinIOSetting.Secure = v.(bool) }),
		stringDefWithActive("minio.endpoint", "storage", "minio", "MinIO endpoint", "MinIO endpoint.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.Endpoint }, func() any { return bootstrapConfig.MinIO.Endpoint }, validateTrimmedMax("minio.endpoint", 255), func(v any) { conf.MinIOSetting.Endpoint = v.(string) }),
		stringDefWithActive("minio.bucket", "storage", "minio", "MinIO bucket", "MinIO bucket name.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.Bucket }, func() any { return bootstrapConfig.MinIO.Bucket }, validateTrimmedMax("minio.bucket", 255), func(v any) { conf.MinIOSetting.Bucket = v.(string) }),
		stringDefWithActive("minio.domain", "storage", "minio", "MinIO domain", "MinIO public domain.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("MinIO") }, func() any { return conf.MinIOSetting.Domain }, func() any { return bootstrapConfig.MinIO.Domain }, validateTrimmedMax("minio.domain", 255), func(v any) { conf.MinIOSetting.Domain = v.(string) }),

		stringDefWithActive("s3.access_key", "storage", "s3", "S3 access key", "Amazon S3 access key.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.AccessKey }, func() any { return bootstrapConfig.S3.AccessKey }, validateTrimmedMax("s3.access_key", 255), func(v any) { conf.S3Setting.AccessKey = v.(string) }),
		stringDefWithActive("s3.secret_key", "storage", "s3", "S3 secret key", "Amazon S3 secret key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.SecretKey }, func() any { return bootstrapConfig.S3.SecretKey }, validateTrimmedMax("s3.secret_key", 512), func(v any) { conf.S3Setting.SecretKey = v.(string) }),
		boolDefWithActive("s3.secure", "storage", "s3", "S3 secure", "Use HTTPS for S3.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.Secure }, func() any { return bootstrapConfig.S3.Secure }, func(v any) { conf.S3Setting.Secure = v.(bool) }),
		stringDefWithActive("s3.endpoint", "storage", "s3", "S3 endpoint", "Amazon S3 endpoint.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.Endpoint }, func() any { return bootstrapConfig.S3.Endpoint }, validateTrimmedMax("s3.endpoint", 255), func(v any) { conf.S3Setting.Endpoint = v.(string) }),
		stringDefWithActive("s3.bucket", "storage", "s3", "S3 bucket", "Amazon S3 bucket name.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.Bucket }, func() any { return bootstrapConfig.S3.Bucket }, validateTrimmedMax("s3.bucket", 255), func(v any) { conf.S3Setting.Bucket = v.(string) }),
		stringDefWithActive("s3.domain", "storage", "s3", "S3 domain", "Amazon S3 public domain.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("S3") }, func() any { return conf.S3Setting.Domain }, func() any { return bootstrapConfig.S3.Domain }, validateTrimmedMax("s3.domain", 255), func(v any) { conf.S3Setting.Domain = v.(string) }),

		stringDefWithActive("alioss.endpoint", "storage", "alioss", "AliOSS endpoint", "AliOSS endpoint.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("AliOSS") }, func() any { return conf.AliOSSSetting.Endpoint }, func() any { return bootstrapConfig.AliOSS.Endpoint }, validateTrimmedMax("alioss.endpoint", 255), func(v any) { conf.AliOSSSetting.Endpoint = v.(string) }),
		stringDefWithActive("alioss.access_key_id", "storage", "alioss", "AliOSS access key ID", "AliOSS access key ID.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("AliOSS") }, func() any { return conf.AliOSSSetting.AccessKeyID }, func() any { return bootstrapConfig.AliOSS.AccessKeyID }, validateTrimmedMax("alioss.access_key_id", 255), func(v any) { conf.AliOSSSetting.AccessKeyID = v.(string) }),
		stringDefWithActive("alioss.access_key_secret", "storage", "alioss", "AliOSS access key secret", "AliOSS secret key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("AliOSS") }, func() any { return conf.AliOSSSetting.AccessKeySecret }, func() any { return bootstrapConfig.AliOSS.AccessKeySecret }, validateTrimmedMax("alioss.access_key_secret", 512), func(v any) { conf.AliOSSSetting.AccessKeySecret = v.(string) }),
		stringDefWithActive("alioss.bucket", "storage", "alioss", "AliOSS bucket", "AliOSS bucket name.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("AliOSS") }, func() any { return conf.AliOSSSetting.Bucket }, func() any { return bootstrapConfig.AliOSS.Bucket }, validateTrimmedMax("alioss.bucket", 255), func(v any) { conf.AliOSSSetting.Bucket = v.(string) }),
		stringDefWithActive("alioss.domain", "storage", "alioss", "AliOSS domain", "AliOSS public domain.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("AliOSS") }, func() any { return conf.AliOSSSetting.Domain }, func() any { return bootstrapConfig.AliOSS.Domain }, validateTrimmedMax("alioss.domain", 255), func(v any) { conf.AliOSSSetting.Domain = v.(string) }),

		stringDefWithActive("cos.secret_id", "storage", "cos", "COS secret ID", "Tencent COS secret ID.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("COS") }, func() any { return conf.COSSetting.SecretID }, func() any { return bootstrapConfig.COS.SecretID }, validateTrimmedMax("cos.secret_id", 255), func(v any) { conf.COSSetting.SecretID = v.(string) }),
		stringDefWithActive("cos.secret_key", "storage", "cos", "COS secret key", "Tencent COS secret key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("COS") }, func() any { return conf.COSSetting.SecretKey }, func() any { return bootstrapConfig.COS.SecretKey }, validateTrimmedMax("cos.secret_key", 512), func(v any) { conf.COSSetting.SecretKey = v.(string) }),
		stringDefWithActive("cos.region", "storage", "cos", "COS region", "Tencent COS region.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("COS") }, func() any { return conf.COSSetting.Region }, func() any { return bootstrapConfig.COS.Region }, validateTrimmedMax("cos.region", 255), func(v any) { conf.COSSetting.Region = v.(string) }),
		stringDefWithActive("cos.bucket", "storage", "cos", "COS bucket", "Tencent COS bucket.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("COS") }, func() any { return conf.COSSetting.Bucket }, func() any { return bootstrapConfig.COS.Bucket }, validateTrimmedMax("cos.bucket", 255), func(v any) { conf.COSSetting.Bucket = v.(string) }),
		stringDefWithActive("cos.domain", "storage", "cos", "COS domain", "Tencent COS public domain.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("COS") }, func() any { return conf.COSSetting.Domain }, func() any { return bootstrapConfig.COS.Domain }, validateTrimmedMax("cos.domain", 255), func(v any) { conf.COSSetting.Domain = v.(string) }),

		stringDefWithActive("huawei_obs.access_key", "storage", "huawei_obs", "Huawei OBS access key", "Huawei OBS access key.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("HuaweiOBS") }, func() any { return conf.HuaweiOBSSetting.AccessKey }, func() any { return bootstrapConfig.HuaweiOBS.AccessKey }, validateTrimmedMax("huawei_obs.access_key", 255), func(v any) { conf.HuaweiOBSSetting.AccessKey = v.(string) }),
		stringDefWithActive("huawei_obs.secret_key", "storage", "huawei_obs", "Huawei OBS secret key", "Huawei OBS secret key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("HuaweiOBS") }, func() any { return conf.HuaweiOBSSetting.SecretKey }, func() any { return bootstrapConfig.HuaweiOBS.SecretKey }, validateTrimmedMax("huawei_obs.secret_key", 512), func(v any) { conf.HuaweiOBSSetting.SecretKey = v.(string) }),
		stringDefWithActive("huawei_obs.endpoint", "storage", "huawei_obs", "Huawei OBS endpoint", "Huawei OBS endpoint.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("HuaweiOBS") }, func() any { return conf.HuaweiOBSSetting.Endpoint }, func() any { return bootstrapConfig.HuaweiOBS.Endpoint }, validateTrimmedMax("huawei_obs.endpoint", 255), func(v any) { conf.HuaweiOBSSetting.Endpoint = v.(string) }),
		stringDefWithActive("huawei_obs.bucket", "storage", "huawei_obs", "Huawei OBS bucket", "Huawei OBS bucket.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("HuaweiOBS") }, func() any { return conf.HuaweiOBSSetting.Bucket }, func() any { return bootstrapConfig.HuaweiOBS.Bucket }, validateTrimmedMax("huawei_obs.bucket", 255), func(v any) { conf.HuaweiOBSSetting.Bucket = v.(string) }),
		stringDefWithActive("huawei_obs.domain", "storage", "huawei_obs", "Huawei OBS domain", "Huawei OBS public domain.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("HuaweiOBS") }, func() any { return conf.HuaweiOBSSetting.Domain }, func() any { return bootstrapConfig.HuaweiOBS.Domain }, validateTrimmedMax("huawei_obs.domain", 255), func(v any) { conf.HuaweiOBSSetting.Domain = v.(string) }),

		stringDefWithActive("sms_juhe.gateway", "notifications", "sms_juhe", "SMS gateway", "Juhe SMS gateway URL.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Sms") }, func() any { return conf.SmsJuheSetting.Gateway }, func() any { return bootstrapConfig.SmsJuhe.Gateway }, validateTrimmedMax("sms_juhe.gateway", 255), func(v any) { conf.SmsJuheSetting.Gateway = v.(string) }),
		stringDefWithActive("sms_juhe.key", "notifications", "sms_juhe", "SMS key", "Juhe SMS key.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("Sms") }, func() any { return conf.SmsJuheSetting.Key }, func() any { return bootstrapConfig.SmsJuhe.Key }, validateTrimmedMax("sms_juhe.key", 255), func(v any) { conf.SmsJuheSetting.Key = v.(string) }),
		stringDefWithActive("sms_juhe.tpl_id", "notifications", "sms_juhe", "SMS template ID", "Juhe SMS template ID.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Sms") }, func() any { return conf.SmsJuheSetting.TplID }, func() any { return bootstrapConfig.SmsJuhe.TplID }, validateTrimmedMax("sms_juhe.tpl_id", 255), func(v any) { conf.SmsJuheSetting.TplID = v.(string) }),
		stringDefWithActive("sms_juhe.tpl_val", "notifications", "sms_juhe", "SMS template value", "Juhe SMS template value format.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Sms") }, func() any { return conf.SmsJuheSetting.TplVal }, func() any { return bootstrapConfig.SmsJuhe.TplVal }, validateTrimmedMax("sms_juhe.tpl_val", 255), func(v any) { conf.SmsJuheSetting.TplVal = v.(string) }),

		stringDefWithActive("alipay.app_id", "payments", "alipay", "Alipay app ID", "Alipay application ID.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.AppID }, func() any { return bootstrapConfig.Alipay.AppID }, validateTrimmedMax("alipay.app_id", 255), func(v any) { conf.AlipaySetting.AppID = v.(string) }),
		stringDefWithActive("alipay.private_key", "payments", "alipay", "Alipay private key", "Alipay private key PEM content.", ApplyModeRestartRequired, true, true, nil, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.PrivateKey }, func() any { return bootstrapConfig.Alipay.PrivateKey }, validateTrimmedMax("alipay.private_key", 8192), func(v any) { conf.AlipaySetting.PrivateKey = v.(string) }),
		stringDefWithActive("alipay.root_cert_file", "payments", "alipay", "Alipay root cert file", "Path to the Alipay root certificate file.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.RootCertFile }, func() any { return bootstrapConfig.Alipay.RootCertFile }, validateTrimmedMax("alipay.root_cert_file", 1024), func(v any) { conf.AlipaySetting.RootCertFile = v.(string) }),
		stringDefWithActive("alipay.public_cert_file", "payments", "alipay", "Alipay public cert file", "Path to the Alipay public certificate file.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.PublicCertFile }, func() any { return bootstrapConfig.Alipay.PublicCertFile }, validateTrimmedMax("alipay.public_cert_file", 1024), func(v any) { conf.AlipaySetting.PublicCertFile = v.(string) }),
		stringDefWithActive("alipay.app_public_cert_file", "payments", "alipay", "Alipay app public cert file", "Path to the app public certificate file.", ApplyModeRestartRequired, false, true, nil, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.AppPublicCertFile }, func() any { return bootstrapConfig.Alipay.AppPublicCertFile }, validateTrimmedMax("alipay.app_public_cert_file", 1024), func(v any) { conf.AlipaySetting.AppPublicCertFile = v.(string) }),
		boolDefWithActive("alipay.in_production", "payments", "alipay", "Alipay production mode", "Use Alipay production environment.", ApplyModeRestartRequired, false, true, func() bool { return cfg.If("Alipay") }, func() any { return conf.AlipaySetting.InProduction }, func() any { return bootstrapConfig.Alipay.InProduction }, func(v any) { conf.AlipaySetting.InProduction = v.(bool) }),
	}
}

func registryMap() map[string]Definition {
	defs := Registry()
	res := make(map[string]Definition, len(defs))
	for _, def := range defs {
		res[def.Key] = def
	}
	return res
}

func boolDef(key, group, section, label, description string, applyMode ApplyMode, secret, editable bool, current func() any, bootstrap func() any, apply func(any)) Definition {
	return boolDefWithActive(key, group, section, label, description, applyMode, secret, editable, func() bool { return true }, current, bootstrap, apply)
}

func boolDefWithActive(key, group, section, label, description string, applyMode ApplyMode, secret, editable bool, active func() bool, current func() any, bootstrap func() any, apply func(any)) Definition {
	return Definition{Key: key, Group: group, Section: section, Type: TypeBool, Label: label, Description: description, ApplyMode: applyMode, Secret: secret, Readonly: !editable, Active: active, CurrentValue: current, BootstrapDefault: bootstrap, ValidateValue: func(v any) (any, error) {
		b, ok := v.(bool)
		if !ok {
			return nil, invalidType(key, "bool")
		}
		return b, nil
	}, BootstrapOverride: apply}
}

func intDef(key, group, section, label, description string, applyMode ApplyMode, secret, editable bool, current func() any, bootstrap func() any, validator func(int) error, apply func(any)) Definition {
	return Definition{Key: key, Group: group, Section: section, Type: TypeInt, Label: label, Description: description, ApplyMode: applyMode, Secret: secret, Readonly: !editable, Active: func() bool { return true }, CurrentValue: current, BootstrapDefault: bootstrap, ValidateValue: func(v any) (any, error) {
		i, ok := asInt(v)
		if !ok {
			return nil, invalidType(key, "int")
		}
		if validator != nil {
			if err := validator(i); err != nil {
				return nil, err
			}
		}
		return i, nil
	}, BootstrapOverride: apply}
}

func int64Def(key, group, section, label, description string, applyMode ApplyMode, current func() any, bootstrap func() any, validator func(int64) error, apply func(any)) Definition {
	return Definition{Key: key, Group: group, Section: section, Type: TypeInt, Label: label, Description: description, ApplyMode: applyMode, Active: func() bool { return true }, CurrentValue: current, BootstrapDefault: bootstrap, ValidateValue: func(v any) (any, error) {
		i, ok := asInt64(v)
		if !ok {
			return nil, invalidType(key, "int")
		}
		if validator != nil {
			if err := validator(i); err != nil {
				return nil, err
			}
		}
		return i, nil
	}, BootstrapOverride: apply}
}

func floatDef(key, group, section, label, description string, applyMode ApplyMode, current func() any, bootstrap func() any, validator func(float64) error, apply func(any)) Definition {
	return Definition{Key: key, Group: group, Section: section, Type: TypeFloat, Label: label, Description: description, ApplyMode: applyMode, Active: func() bool { return true }, CurrentValue: current, BootstrapDefault: bootstrap, ValidateValue: func(v any) (any, error) {
		f, ok := asFloat64(v)
		if !ok {
			return nil, invalidType(key, "float")
		}
		if validator != nil {
			if err := validator(f); err != nil {
				return nil, err
			}
		}
		return f, nil
	}, BootstrapOverride: apply}
}

func stringDef(key, group, section, label, description string, applyMode ApplyMode, secret, editable bool, options []Option, current func() any, bootstrap func() any, validator func(string) error, apply func(any)) Definition {
	return stringDefWithActive(key, group, section, label, description, applyMode, secret, editable, options, func() bool { return true }, current, bootstrap, validator, apply)
}

func stringDefWithActive(key, group, section, label, description string, applyMode ApplyMode, secret, editable bool, options []Option, active func() bool, current func() any, bootstrap func() any, validator func(string) error, apply func(any)) Definition {
	return Definition{Key: key, Group: group, Section: section, Type: TypeString, Label: label, Description: description, ApplyMode: applyMode, Secret: secret, Readonly: !editable, Options: options, Active: active, CurrentValue: current, BootstrapDefault: bootstrap, ValidateValue: func(v any) (any, error) {
		s, ok := v.(string)
		if !ok {
			return nil, invalidType(key, "string")
		}
		s = strings.TrimSpace(s)
		if validator != nil {
			if err := validator(s); err != nil {
				return nil, err
			}
		}
		return s, nil
	}, BootstrapOverride: apply}
}

func validateVisibility(v string) error {
	for _, option := range visibilityOption {
		if option.Value == v {
			return nil
		}
	}
	return xerror.InvalidParams.WithDetails("default_tweet_visibility must be one of public/following/friend/private")
}

func validateRequiredTrimmed(name string, max int) func(string) error {
	return func(v string) error {
		if v == "" {
			return xerror.InvalidParams.WithDetails(name + " must not be empty")
		}
		if len(v) > max {
			return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s length must be <= %d", name, max))
		}
		return nil
	}
}

func validateTrimmedMax(name string, max int) func(string) error {
	return func(v string) error {
		if len(v) > max {
			return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s length must be <= %d", name, max))
		}
		return nil
	}
}

func validateOptionalURL(name string, max int) func(string) error {
	return func(v string) error {
		if len(v) > max {
			return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s length must be <= %d", name, max))
		}
		if v == "" {
			return nil
		}
		parsed, err := url.ParseRequestURI(v)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return xerror.InvalidParams.WithDetails(name + " must be a valid URL")
		}
		return nil
	}
}

func between(v, min, max int, name string) error {
	if v < min || v > max {
		return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s must be between %d and %d", name, min, max))
	}
	return nil
}

func betweenInt64(v, min, max int64, name string) error {
	if v < min || v > max {
		return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s must be between %d and %d", name, min, max))
	}
	return nil
}

func invalidType(key, want string) error {
	return xerror.InvalidParams.WithDetails(fmt.Sprintf("%s must be %s", key, want))
}

func asInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	case float64:
		if n != float64(int(n)) {
			return 0, false
		}
		return int(n), true
	default:
		return 0, false
	}
}

func asInt64(v any) (int64, bool) {
	switch n := v.(type) {
	case int:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	case float64:
		if n != float64(int64(n)) {
			return 0, false
		}
		return int64(n), true
	default:
		return 0, false
	}
}

func asFloat64(v any) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int:
		return float64(n), true
	case int64:
		return float64(n), true
	default:
		return 0, false
	}
}

func parseRequestValue(def Definition, raw stdjson.RawMessage) (any, error) {
	var target any
	switch def.Type {
	case TypeBool:
		var v bool
		if err := stdjson.Unmarshal(raw, &v); err != nil {
			return nil, invalidType(def.Key, "bool")
		}
		target = v
	case TypeInt:
		var v float64
		if err := stdjson.Unmarshal(raw, &v); err != nil {
			return nil, invalidType(def.Key, "int")
		}
		target = v
	case TypeFloat:
		var v float64
		if err := stdjson.Unmarshal(raw, &v); err != nil {
			return nil, invalidType(def.Key, "float")
		}
		target = v
	case TypeString:
		var v string
		if err := stdjson.Unmarshal(raw, &v); err != nil {
			return nil, invalidType(def.Key, "string")
		}
		target = v
	default:
		return nil, xerror.ServerError.WithDetails("unsupported settings type")
	}
	return def.ValidateValue(target)
}

func parseStoredValue(def Definition, raw string) (any, error) {
	switch def.Type {
	case TypeBool:
		return strconv.ParseBool(raw)
	case TypeInt:
		if i, err := strconv.ParseInt(raw, 10, 64); err == nil {
			return def.ValidateValue(i)
		} else {
			return nil, err
		}
	case TypeFloat:
		if f, err := strconv.ParseFloat(raw, 64); err == nil {
			return def.ValidateValue(f)
		} else {
			return nil, err
		}
	case TypeString:
		return def.ValidateValue(raw)
	default:
		return nil, xerror.ServerError.WithDetails("unsupported settings type")
	}
}

func serializeValue(def Definition, value any) string {
	switch def.Type {
	case TypeBool:
		return strconv.FormatBool(value.(bool))
	case TypeInt:
		switch v := value.(type) {
		case int:
			return strconv.Itoa(v)
		case int64:
			return strconv.FormatInt(v, 10)
		default:
			return fmt.Sprintf("%v", value)
		}
	case TypeFloat:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	default:
		return value.(string)
	}
}

func valuesEqual(left, right any) bool {
	return fmt.Sprintf("%v", left) == fmt.Sprintf("%v", right)
}

func activeState(def Definition) bool {
	if def.Active == nil {
		return true
	}
	return def.Active()
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func logBootstrapApplyError(def Definition, err error) {
	logrus.WithError(err).WithField("key", def.Key).Warn("sitesetting: skip invalid persisted override")
}
