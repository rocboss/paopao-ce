package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Admin 运维相关服务
type Admin struct {
	Schema `mir:"v1,chain"`

	// ChangeUserStatus 管理·禁言/解封用户
	ChangeUserStatus   func(Post, web.ChangeUserStatusReq)                            `mir:"admin/user/status"`
	SiteInfo           func(Get, web.SiteInfoReq) web.SiteInfoResp                    `mir:"admin/site/status"`
	GetSiteSettings    func(Get) web.SiteSettingsResp                                 `mir:"admin/site/profile"`
	UpdateSiteSettings func(Post, web.SiteSettingsReq) web.SiteSettingsResp           `mir:"admin/site/profile"`
	GetSettingsSchema  func(Get) web.AdminSettingsSchemaResp                          `mir:"admin/settings/schema"`
	GetSettingsValues  func(Get) web.AdminSettingsValuesResp                          `mir:"admin/settings/values"`
	SaveSettings       func(Post, web.AdminSettingsSaveReq) web.AdminSettingsSaveResp `mir:"admin/settings/save"`
}
