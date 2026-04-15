// Copyright 2026 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package settings

import (
	"context"
	stdjson "encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EditableProfile struct {
	UseFriendship             bool
	EnableTrendsBar           bool
	EnableWallet              bool
	AllowTweetAttachment      bool
	AllowTweetAttachmentPrice bool
	AllowTweetVideo           bool
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

type settingRecord struct {
	Key         string `gorm:"column:key;type:varchar(191);primaryKey"`
	Value       string `gorm:"column:value;type:text;not null"`
	IsEncrypted bool   `gorm:"column:is_encrypted;not null"`
	CreatedOn   int64  `gorm:"column:created_on;not null"`
	ModifiedOn  int64  `gorm:"column:modified_on;not null"`
	DeletedOn   int64  `gorm:"column:deleted_on;not null"`
	IsDel       int8   `gorm:"column:is_del;not null"`
}

func (settingRecord) TableName() string {
	if conf.DatabaseSetting == nil {
		return "p_" + conf.TableSiteSettings
	}
	return conf.DatabaseSetting.TablePrefix + conf.TableSiteSettings
}

func (r *settingRecord) BeforeCreate(_ *gorm.DB) error {
	now := time.Now().Unix()
	r.CreatedOn = now
	r.ModifiedOn = now
	return nil
}

func (r *settingRecord) BeforeUpdate(tx *gorm.DB) error {
	if !tx.Statement.Changed("modified_on") {
		r.ModifiedOn = time.Now().Unix()
	}
	return nil
}

type Service struct {
	db       *gorm.DB
	registry map[string]Definition
	codec    *secretCodec
}

func NewService(db *gorm.DB) *Service {
	ensureBootstrapSnapshot()
	return &Service{db: db, registry: registryMap(), codec: newSecretCodec()}
}

func Bootstrap(ctx context.Context, db *gorm.DB) {
	ensureBootstrapSnapshot()
	if db == nil {
		return
	}
	if err := NewService(db).ApplyPersistedOverrides(ctx); err != nil {
		logrus.WithError(err).Warn("sitesetting: bootstrap override load failed; using bootstrap config only")
	}
}

func (s *Service) ApplyPersistedOverrides(ctx context.Context) error {
	records, err := s.loadOverridesLenient(ctx)
	if err != nil {
		return err
	}
	for key, record := range records {
		def, ok := s.registry[key]
		if !ok || def.BootstrapOverride == nil || def.ApplyMode == ApplyModeBootstrapOnly {
			continue
		}
		value, err := s.parseRecord(def, record)
		if err != nil {
			logBootstrapApplyError(def, err)
			continue
		}
		def.BootstrapOverride(value)
	}
	return nil
}

func (s *Service) GetProfile(_ctx context.Context) (*web.SiteProfileResp, error) {
	return &web.SiteProfileResp{
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
	}, nil
}

func (s *Service) UpdateEditableProfile(ctx context.Context, input EditableProfile) (*web.SiteProfileResp, error) {
	if err := ValidateEditableProfile(input); err != nil {
		return nil, err
	}
	items := []web.AdminSettingValueInput{
		{Key: "web_profile.use_friendship", Value: boolRaw(input.UseFriendship)},
		{Key: "web_profile.enable_trends_bar", Value: boolRaw(input.EnableTrendsBar)},
		{Key: "web_profile.enable_wallet", Value: boolRaw(input.EnableWallet)},
		{Key: "web_profile.allow_tweet_attachment", Value: boolRaw(input.AllowTweetAttachment)},
		{Key: "web_profile.allow_tweet_attachment_price", Value: boolRaw(input.AllowTweetAttachmentPrice)},
		{Key: "web_profile.allow_tweet_video", Value: boolRaw(input.AllowTweetVideo)},
		{Key: "web_profile.default_tweet_max_length", Value: intRaw(input.DefaultTweetMaxLength)},
		{Key: "web_profile.tweet_web_ellipsis_size", Value: intRaw(input.TweetWebEllipsisSize)},
		{Key: "web_profile.tweet_mobile_ellipsis_size", Value: intRaw(input.TweetMobileEllipsisSize)},
		{Key: "web_profile.default_tweet_visibility", Value: stringRaw(input.DefaultTweetVisibility)},
		{Key: "web_profile.default_msg_loop_interval", Value: intRaw(input.DefaultMsgLoopInterval)},
		{Key: "web_profile.copyright_top", Value: stringRaw(input.CopyrightTop)},
		{Key: "web_profile.copyright_left", Value: stringRaw(input.CopyrightLeft)},
		{Key: "web_profile.copyright_left_link", Value: stringRaw(input.CopyrightLeftLink)},
		{Key: "web_profile.copyright_right", Value: stringRaw(input.CopyrightRight)},
		{Key: "web_profile.copyright_right_link", Value: stringRaw(input.CopyrightRightLink)},
	}
	if _, err := s.SaveValues(ctx, items); err != nil {
		return nil, err
	}
	return s.GetProfile(ctx)
}

func (s *Service) GetSchema() (*web.AdminSettingsSchemaResp, error) {
	items := make([]web.AdminSettingSchemaItem, 0, len(s.registry))
	for _, def := range Registry() {
		item := web.AdminSettingSchemaItem{
			Key:         def.Key,
			Group:       def.Group,
			Section:     def.Section,
			Type:        string(def.Type),
			Label:       def.Label,
			Description: def.Description,
			ApplyMode:   string(def.ApplyMode),
			Secret:      def.Secret,
			Readonly:    def.Readonly,
			Active:      activeState(def),
			Options:     def.Options,
		}
		bootstrap := def.BootstrapDefault()
		if def.Secret {
			item.BootstrapConfigured = configuredValue(bootstrap)
		} else {
			item.BootstrapValue = bootstrap
		}
		items = append(items, item)
	}
	return &web.AdminSettingsSchemaResp{Items: items}, nil
}

func (s *Service) GetValues(ctx context.Context) (*web.AdminSettingsValuesResp, error) {
	records, err := s.loadOverridesLenient(ctx)
	if err != nil {
		return nil, err
	}
	items, pending := s.buildValueItems(records)
	return &web.AdminSettingsValuesResp{Items: items, HasPendingRestart: pending}, nil
}

func (s *Service) SaveValues(ctx context.Context, inputs []web.AdminSettingValueInput) (*web.AdminSettingsSaveResp, error) {
	if len(inputs) == 0 {
		return nil, xerror.InvalidParams.WithDetails("items must not be empty")
	}
	prepared := make([]preparedValue, 0, len(inputs))
	seen := make(map[string]struct{}, len(inputs))
	for _, input := range inputs {
		if len(input.Value) == 0 {
			return nil, xerror.InvalidParams.WithDetails("missing setting value: " + input.Key)
		}
		def, ok := s.registry[input.Key]
		if !ok {
			return nil, xerror.InvalidParams.WithDetails("unknown setting key: " + input.Key)
		}
		if def.Readonly || def.ApplyMode == ApplyModeBootstrapOnly || def.BootstrapOverride == nil {
			return nil, xerror.InvalidParams.WithDetails("setting is not editable: " + input.Key)
		}
		if _, ok := seen[input.Key]; ok {
			return nil, xerror.InvalidParams.WithDetails("duplicate setting key: " + input.Key)
		}
		seen[input.Key] = struct{}{}
		value, err := parseRequestValue(def, input.Value)
		if err != nil {
			return nil, err
		}
		prepared = append(prepared, preparedValue{Definition: def, Value: value})
	}
	if err := validatePreparedBatch(prepared); err != nil {
		return nil, err
	}
	if err := s.persistPrepared(ctx, prepared); err != nil {
		return nil, err
	}
	for _, item := range prepared {
		if item.ApplyMode == ApplyModeLive {
			item.BootstrapOverride(item.Value)
		}
	}
	values, err := s.GetValues(ctx)
	if err != nil {
		return nil, err
	}
	updatedKeys := make([]string, 0, len(prepared))
	for _, item := range prepared {
		updatedKeys = append(updatedKeys, item.Key)
	}
	return &web.AdminSettingsSaveResp{Items: values.Items, UpdatedKeys: updatedKeys, HasPendingRestart: values.HasPendingRestart}, nil
}

type preparedValue struct {
	Definition
	Value any
}

func validatePreparedBatch(prepared []preparedValue) error {
	defaultTweetMaxLength := preparedInt(prepared, "web_profile.default_tweet_max_length", conf.WebProfileSetting.DefaultTweetMaxLength)
	tweetWebEllipsisSize := preparedInt(prepared, "web_profile.tweet_web_ellipsis_size", conf.WebProfileSetting.TweetWebEllipsisSize)
	tweetMobileEllipsisSize := preparedInt(prepared, "web_profile.tweet_mobile_ellipsis_size", conf.WebProfileSetting.TweetMobileEllipsisSize)
	defaultPageSize := preparedInt(prepared, "app.default_page_size", conf.AppSetting.DefaultPageSize)
	maxPageSize := preparedInt(prepared, "app.max_page_size", conf.AppSetting.MaxPageSize)

	if tweetWebEllipsisSize < 1 || tweetWebEllipsisSize > defaultTweetMaxLength {
		return xerror.InvalidParams.WithDetails("tweet_web_ellipsis_size must be between 1 and default_tweet_max_length")
	}
	if tweetMobileEllipsisSize < 1 || tweetMobileEllipsisSize > defaultTweetMaxLength {
		return xerror.InvalidParams.WithDetails("tweet_mobile_ellipsis_size must be between 1 and default_tweet_max_length")
	}
	if defaultPageSize < 1 || defaultPageSize > maxPageSize {
		return xerror.InvalidParams.WithDetails("default_page_size must be between 1 and max_page_size")
	}
	return nil
}

func preparedInt(prepared []preparedValue, key string, fallback int) int {
	for _, item := range prepared {
		if item.Key == key {
			if value, ok := item.Value.(int); ok {
				return value
			}
			break
		}
	}
	return fallback
}

func (s *Service) persistPrepared(ctx context.Context, prepared []preparedValue) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range prepared {
			serialized := serializeValue(item.Definition, item.Value)
			if valuesEqual(item.BootstrapDefault(), item.Value) {
				if err := tx.Where("key = ?", item.Key).Delete(&settingRecord{}).Error; err != nil {
					return err
				}
				continue
			}
			storedValue := serialized
			isEncrypted := false
			if item.Secret {
				encrypted, err := s.codec.Encrypt(serialized)
				if err != nil {
					return err
				}
				storedValue = encrypted
				isEncrypted = true
			}
			record := settingRecord{Key: item.Key, Value: storedValue, IsEncrypted: isEncrypted}
			if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "key"}}, DoUpdates: clause.AssignmentColumns([]string{"value", "is_encrypted", "modified_on", "deleted_on", "is_del"})}).Create(&record).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Service) buildValueItems(records map[string]settingRecord) ([]web.AdminSettingValue, bool) {
	items := make([]web.AdminSettingValue, 0, len(s.registry))
	hasPendingRestart := false
	for _, def := range Registry() {
		current := def.CurrentValue()
		item := web.AdminSettingValue{Key: def.Key, Source: "bootstrap", Active: activeState(def)}
		if def.Secret {
			item.Configured = configuredValue(current)
		} else {
			item.Value = current
			item.EffectiveValue = current
		}
		if record, ok := records[def.Key]; ok {
			storedValue, err := s.parseRecord(def, record)
			if err == nil {
				item.Source = "override"
				if def.Secret {
					item.Configured = configuredValue(storedValue)
				} else {
					item.Value = storedValue
					item.EffectiveValue = current
				}
				if def.ApplyMode == ApplyModeRestartRequired && !valuesEqual(storedValue, current) {
					item.PendingRestart = true
					hasPendingRestart = true
				}
			}
		}
		if def.Secret && item.Source == "bootstrap" {
			item.Configured = configuredValue(def.BootstrapDefault())
		}
		items = append(items, item)
	}
	return items, hasPendingRestart
}

func (s *Service) parseRecord(def Definition, record settingRecord) (any, error) {
	raw := record.Value
	if record.IsEncrypted {
		decrypted, err := s.codec.Decrypt(record.Value)
		if err != nil {
			return nil, err
		}
		raw = decrypted
	}
	return parseStoredValue(def, raw)
}

func (s *Service) loadOverrides(ctx context.Context) (map[string]settingRecord, error) {
	var records []settingRecord
	err := s.db.WithContext(ctx).Where("is_del = ?", 0).Find(&records).Error
	if err != nil {
		return nil, err
	}
	res := make(map[string]settingRecord, len(records))
	for _, record := range records {
		res[record.Key] = record
	}
	return res, nil
}

func (s *Service) loadOverridesLenient(ctx context.Context) (map[string]settingRecord, error) {
	records, err := s.loadOverrides(ctx)
	if err != nil {
		if isMissingTableError(err) {
			logrus.WithError(err).Warn("sitesetting: overrides table unavailable; falling back to bootstrap values")
			return map[string]settingRecord{}, nil
		}
		return nil, err
	}
	return records, nil
}

func isMissingTableError(err error) bool {
	if err == nil {
		return false
	}
	text := strings.ToLower(err.Error())
	return strings.Contains(text, "no such table") || strings.Contains(text, "doesn't exist") || strings.Contains(text, "does not exist") || errors.Is(err, gorm.ErrRecordNotFound)
}

func configuredValue(v any) bool {
	switch value := v.(type) {
	case nil:
		return false
	case string:
		return strings.TrimSpace(value) != ""
	default:
		return true
	}
}

func boolRaw(v bool) []byte {
	b, _ := stdjson.Marshal(v)
	return b
}

func intRaw(v int) []byte {
	b, _ := stdjson.Marshal(v)
	return b
}

func stringRaw(v string) []byte {
	b, _ := stdjson.Marshal(strings.TrimSpace(v))
	return b
}

func ValidateEditableProfile(input EditableProfile) error {
	return validateProfileInput(input)
}

func validateProfileInput(input EditableProfile) error {
	if err := between(input.DefaultTweetMaxLength, 1, 2000, "default_tweet_max_length"); err != nil {
		return err
	}
	if err := between(input.TweetWebEllipsisSize, 1, input.DefaultTweetMaxLength, "tweet_web_ellipsis_size"); err != nil {
		return err
	}
	if err := between(input.TweetMobileEllipsisSize, 1, input.DefaultTweetMaxLength, "tweet_mobile_ellipsis_size"); err != nil {
		return err
	}
	if err := between(input.DefaultMsgLoopInterval, 1000, 60000, "default_msg_loop_interval"); err != nil {
		return err
	}
	if err := validateVisibility(strings.TrimSpace(input.DefaultTweetVisibility)); err != nil {
		return err
	}
	if err := validateRequiredTrimmed("copyright_top", 255)(input.CopyrightTop); err != nil {
		return err
	}
	if err := validateRequiredTrimmed("copyright_left", 255)(input.CopyrightLeft); err != nil {
		return err
	}
	if err := validateRequiredTrimmed("copyright_right", 255)(input.CopyrightRight); err != nil {
		return err
	}
	if err := validateOptionalURL("copyright_left_link", 255)(input.CopyrightLeftLink); err != nil {
		return err
	}
	if err := validateOptionalURL("copyright_right_link", 255)(input.CopyrightRightLink); err != nil {
		return err
	}
	return nil
}

func EditableFromRequest(req *web.SiteSettingsReq) EditableProfile {
	return EditableProfile{
		UseFriendship:             *req.UseFriendship,
		EnableTrendsBar:           *req.EnableTrendsBar,
		EnableWallet:              *req.EnableWallet,
		AllowTweetAttachment:      *req.AllowTweetAttachment,
		AllowTweetAttachmentPrice: *req.AllowTweetAttachmentPrice,
		AllowTweetVideo:           *req.AllowTweetVideo,
		DefaultTweetMaxLength:     *req.DefaultTweetMaxLength,
		TweetWebEllipsisSize:      *req.TweetWebEllipsisSize,
		TweetMobileEllipsisSize:   *req.TweetMobileEllipsisSize,
		DefaultTweetVisibility:    *req.DefaultTweetVisibility,
		DefaultMsgLoopInterval:    *req.DefaultMsgLoopInterval,
		CopyrightTop:              *req.CopyrightTop,
		CopyrightLeft:             *req.CopyrightLeft,
		CopyrightLeftLink:         stringValue(req.CopyrightLeftLink),
		CopyrightRight:            *req.CopyrightRight,
		CopyrightRightLink:        stringValue(req.CopyrightRightLink),
	}
}

func stringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
