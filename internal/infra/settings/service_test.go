package settings

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestGetProfileUsesBootstrapDefaultsWhenNoOverride(t *testing.T) {
	svc := newTestService(t)

	profile, err := svc.GetProfile(context.Background())
	if err != nil {
		t.Fatalf("GetProfile() error = %v", err)
	}
	if !profile.AllowUserRegister {
		t.Fatalf("AllowUserRegister = false, want bootstrap true")
	}
	if profile.DefaultTweetVisibility != "friend" {
		t.Fatalf("DefaultTweetVisibility = %q, want friend", profile.DefaultTweetVisibility)
	}
	if profile.CopyrightRight != "fallback-right" {
		t.Fatalf("CopyrightRight = %q, want fallback-right", profile.CopyrightRight)
	}
}

func TestUpdateEditableProfilePersistsOnlyEditableKeys(t *testing.T) {
	svc := newTestService(t)

	profile, err := svc.UpdateEditableProfile(context.Background(), EditableProfile{
		UseFriendship:             false,
		EnableTrendsBar:           true,
		EnableWallet:              true,
		AllowTweetAttachment:      false,
		AllowTweetAttachmentPrice: false,
		AllowTweetVideo:           false,
		DefaultTweetMaxLength:     1200,
		TweetWebEllipsisSize:      300,
		TweetMobileEllipsisSize:   200,
		DefaultTweetVisibility:    "public",
		DefaultMsgLoopInterval:    3000,
		CopyrightTop:              "top",
		CopyrightLeft:             "left",
		CopyrightLeftLink:         "https://left.example.com",
		CopyrightRight:            "right",
		CopyrightRightLink:        "https://right.example.com",
	})
	if err != nil {
		t.Fatalf("UpdateEditableProfile() error = %v", err)
	}
	if !profile.AllowUserRegister {
		t.Fatalf("AllowUserRegister = false, want bootstrap true")
	}
	if !profile.AllowPhoneBind {
		t.Fatalf("AllowPhoneBind = false, want bootstrap true")
	}
	if profile.DefaultTweetVisibility != "public" {
		t.Fatalf("DefaultTweetVisibility = %q, want public", profile.DefaultTweetVisibility)
	}
	values, err := svc.GetValues(context.Background())
	if err != nil {
		t.Fatalf("GetValues() error = %v", err)
	}
	if !hasValue(values.Items, "web_profile.enable_wallet", true, false) {
		t.Fatalf("web_profile.enable_wallet override not found")
	}
}

func TestSaveValuesEncryptsSecretsAtRest(t *testing.T) {
	svc := newTestService(t)
	conf.AdminSettingsSetting.EncryptionKey = "bootstrap-test-encryption-key"
	svc.codec = newSecretCodec()

	_, err := svc.SaveValues(context.Background(), []web.AdminSettingValueInput{{
		Key:   "meili.api_key",
		Value: []byte(`"top-secret-key"`),
	}})
	if err != nil {
		t.Fatalf("SaveValues() error = %v", err)
	}
	var record settingRecord
	if err := svc.db.WithContext(context.Background()).First(&record, "key = ?", "meili.api_key").Error; err != nil {
		t.Fatalf("load record error = %v", err)
	}
	if !record.IsEncrypted {
		t.Fatalf("IsEncrypted = false, want true")
	}
	if strings.Contains(record.Value, "top-secret-key") {
		t.Fatalf("record.Value stored plaintext = %q", record.Value)
	}
}

func TestRestartRequiredValuesReportPendingRestart(t *testing.T) {
	svc := newTestService(t)

	resp, err := svc.SaveValues(context.Background(), []web.AdminSettingValueInput{{
		Key:   "meili.host",
		Value: []byte(`"pending-restart:7700"`),
	}})
	if err != nil {
		t.Fatalf("SaveValues() error = %v", err)
	}
	if !resp.HasPendingRestart {
		t.Fatal("HasPendingRestart = false, want true")
	}
	if !hasValue(resp.Items, "meili.host", "pending-restart:7700", true) {
		t.Fatalf("meili.host pending_restart not reported")
	}
}

func TestSaveValuesRejectsInvalidSingleIntUpdate(t *testing.T) {
	svc := newTestService(t)

	_, err := svc.SaveValues(context.Background(), []web.AdminSettingValueInput{{
		Key:   "web_profile.default_tweet_max_length",
		Value: []byte(`120`),
	}})
	if err == nil {
		t.Fatal("SaveValues() error = nil, want invalid params error")
	}
}

func TestSaveValuesAcceptsCoupledProfileUpdate(t *testing.T) {
	svc := newTestService(t)

	resp, err := svc.SaveValues(context.Background(), []web.AdminSettingValueInput{
		{Key: "web_profile.default_tweet_max_length", Value: []byte(`120`)},
		{Key: "web_profile.tweet_web_ellipsis_size", Value: []byte(`120`)},
		{Key: "web_profile.tweet_mobile_ellipsis_size", Value: []byte(`120`)},
	})
	if err != nil {
		t.Fatalf("SaveValues() error = %v", err)
	}
	if !hasValue(resp.Items, "web_profile.default_tweet_max_length", 120, false) {
		t.Fatalf("web_profile.default_tweet_max_length override not found")
	}
	if conf.WebProfileSetting.DefaultTweetMaxLength != 120 {
		t.Fatalf("DefaultTweetMaxLength = %d, want 120", conf.WebProfileSetting.DefaultTweetMaxLength)
	}
}

func TestBootstrapAppliesPersistedOverrides(t *testing.T) {
	svc := newTestService(t)
	conf.AdminSettingsSetting.EncryptionKey = "bootstrap-test-encryption-key"
	svc.codec = newSecretCodec()

	_, err := svc.SaveValues(context.Background(), []web.AdminSettingValueInput{{
		Key:   "web_profile.enable_wallet",
		Value: []byte(`true`),
	}})
	if err != nil {
		t.Fatalf("SaveValues() error = %v", err)
	}
	conf.WebProfileSetting.EnableWallet = false
	Bootstrap(context.Background(), svc.db)
	if !conf.WebProfileSetting.EnableWallet {
		t.Fatal("Bootstrap() did not apply persisted web_profile.enable_wallet override")
	}
}

func newTestService(t *testing.T) *Service {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error = %v", err)
	}
	root := filepath.Clean(filepath.Join(wd, "../../.."))
	if err := os.Chdir(root); err != nil {
		t.Fatalf("os.Chdir(%q) error = %v", root, err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
	conf.Initial(nil, false)
	conf.WebProfileSetting = &conf.WebProfileConf{
		UseFriendship:             true,
		EnableTrendsBar:           false,
		EnableWallet:              false,
		AllowTweetAttachment:      true,
		AllowTweetAttachmentPrice: true,
		AllowTweetVideo:           true,
		AllowUserRegister:         true,
		AllowPhoneBind:            true,
		DefaultTweetMaxLength:     2000,
		TweetWebEllipsisSize:      400,
		TweetMobileEllipsisSize:   300,
		DefaultTweetVisibility:    "friend",
		DefaultMsgLoopInterval:    5000,
		CopyrightTop:              "fallback-top",
		CopyrightLeft:             "fallback-left",
		CopyrightLeftLink:         "",
		CopyrightRight:            "fallback-right",
		CopyrightRightLink:        "https://fallback.example.com",
	}
	bootstrapConfig = nil
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "p_", SingularTable: true}})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}
	if err := db.AutoMigrate(&settingRecord{}); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}
	return NewService(db)
}

func hasValue(items []web.AdminSettingValue, key string, expected any, pending bool) bool {
	for _, item := range items {
		if item.Key != key {
			continue
		}
		return item.Value == expected && item.PendingRestart == pending
	}
	return false
}
