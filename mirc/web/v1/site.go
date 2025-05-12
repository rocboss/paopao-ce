package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Site 站点本身相关的信息服务
type Site struct {
	Schema `mir:"v1"`

	// Version 获取后台版本信息
	Version func(Get) web.VersionResp `mir:"site/version"`

	// Profile 站点配置概要信息
	Profile func(Get) web.SiteProfileResp `mir:"site/profile"`
}
