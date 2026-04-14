// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"

	"github.com/rocboss/paopao-ce/pkg/version"
)

type VersionResp struct {
	BuildInfo *version.BuildInfo `json:"build_info"`
}

type SiteProfileResp struct {
	UseFriendship             bool   `json:"use_friendship"`
	EnableTrendsBar           bool   `json:"enable_trends_bar"`
	EnableWallet              bool   `json:"enable_wallet"`
	AllowTweetAttachment      bool   `json:"allow_tweet_attachment"`
	AllowTweetAttachmentPrice bool   `json:"allow_tweet_attachment_price"`
	AllowTweetVideo           bool   `json:"allow_tweet_video"`
	AllowUserRegister         bool   `json:"allow_user_register"`
	AllowPhoneBind            bool   `json:"allow_phone_bind"`
	DefaultTweetMaxLength     int    `json:"default_tweet_max_length"`
	TweetWebEllipsisSize      int    `json:"tweet_web_ellipsis_size"`
	TweetMobileEllipsisSize   int    `json:"tweet_mobile_ellipsis_size"`
	DefaultTweetVisibility    string `json:"default_tweet_visibility"`
	DefaultMsgLoopInterval    int    `json:"default_msg_loop_interval"`
	CopyrightTop              string `json:"copyright_top"`
	CopyrightLeft             string `json:"copyright_left"`
	CopyrightLeftLink         string `json:"copyright_left_link"`
	CopyrightRight            string `json:"copyright_right"`
	CopyrightRightLink        string `json:"copyright_right_link"`
}

type SiteSettingsReq struct {
	BaseInfo `json:"-" binding:"-"`

	UseFriendship             *bool   `json:"use_friendship" binding:"required"`
	EnableTrendsBar           *bool   `json:"enable_trends_bar" binding:"required"`
	EnableWallet              *bool   `json:"enable_wallet" binding:"required"`
	AllowTweetAttachment      *bool   `json:"allow_tweet_attachment" binding:"required"`
	AllowTweetAttachmentPrice *bool   `json:"allow_tweet_attachment_price" binding:"required"`
	AllowTweetVideo           *bool   `json:"allow_tweet_video" binding:"required"`
	DefaultTweetMaxLength     *int    `json:"default_tweet_max_length" binding:"required,gte=1,lte=2000"`
	TweetWebEllipsisSize      *int    `json:"tweet_web_ellipsis_size" binding:"required,gte=1,lte=2000"`
	TweetMobileEllipsisSize   *int    `json:"tweet_mobile_ellipsis_size" binding:"required,gte=1,lte=2000"`
	DefaultTweetVisibility    *string `json:"default_tweet_visibility" binding:"required,oneof=public following friend private"`
	DefaultMsgLoopInterval    *int    `json:"default_msg_loop_interval" binding:"required,gte=1000,lte=60000"`
	CopyrightTop              *string `json:"copyright_top" binding:"required,max=255"`
	CopyrightLeft             *string `json:"copyright_left" binding:"required,max=255"`
	CopyrightLeftLink         *string `json:"copyright_left_link" binding:"omitempty,max=255,url"`
	CopyrightRight            *string `json:"copyright_right" binding:"required,max=255"`
	CopyrightRightLink        *string `json:"copyright_right_link" binding:"omitempty,max=255,url"`
}

type SiteSettingsResp struct {
	SiteProfileResp
	ReadonlyFields []string `json:"readonly_fields"`
}

type AdminSettingValueInput struct {
	Key   string          `json:"key" binding:"required"`
	Value json.RawMessage `json:"value"`
}

type AdminSettingsSaveReq struct {
	BaseInfo `json:"-" binding:"-"`

	Items []AdminSettingValueInput `json:"items" binding:"required"`
}

type AdminSettingSchemaItem struct {
	Key                 string `json:"key"`
	Group               string `json:"group"`
	Section             string `json:"section"`
	Type                string `json:"type"`
	Label               string `json:"label"`
	Description         string `json:"description"`
	ApplyMode           string `json:"apply_mode"`
	Secret              bool   `json:"secret"`
	Readonly            bool   `json:"readonly"`
	Active              bool   `json:"active"`
	BootstrapValue      any    `json:"bootstrap_value,omitempty"`
	BootstrapConfigured bool   `json:"bootstrap_configured,omitempty"`
	Options             any    `json:"options,omitempty"`
}

type AdminSettingsSchemaResp struct {
	Items []AdminSettingSchemaItem `json:"items"`
}

type AdminSettingValue struct {
	Key            string `json:"key"`
	Value          any    `json:"value,omitempty"`
	EffectiveValue any    `json:"effective_value,omitempty"`
	Source         string `json:"source"`
	PendingRestart bool   `json:"pending_restart"`
	Configured     bool   `json:"configured,omitempty"`
	Active         bool   `json:"active"`
}

type AdminSettingsValuesResp struct {
	Items             []AdminSettingValue `json:"items"`
	HasPendingRestart bool                `json:"has_pending_restart"`
}

type AdminSettingsSaveResp struct {
	Items             []AdminSettingValue `json:"items"`
	UpdatedKeys       []string            `json:"updated_keys"`
	HasPendingRestart bool                `json:"has_pending_restart"`
}
