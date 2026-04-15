// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/infra/settings"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

type adminSrv struct {
	api.UnimplementedAdminServant
	*base.DaoServant
	wc           core.WebCache
	settings     *settings.Service
	serverUpTime int64
}

func (s *adminSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT(), chain.Admin()}
}

func (s *adminSrv) ChangeUserStatus(req *web.ChangeUserStatusReq) error {
	user, err := s.Ds.GetUserByID(req.ID)
	if err != nil || user.Model == nil || user.ID <= 0 {
		return web.ErrNoExistUsername
	}
	// 执行更新
	user.Status = req.Status
	if err := s.Ds.UpdateUser(user); err != nil {
		return xerror.ServerError
	}
	return nil
}

func (s *adminSrv) SiteInfo(req *web.SiteInfoReq) (*web.SiteInfoResp, error) {
	res, err := &web.SiteInfoResp{ServerUpTime: s.serverUpTime}, error(nil)
	res.RegisterUserCount, err = s.Ds.GetRegisterUserCount()
	if err != nil {
		logrus.Errorf("get SiteInfo[1] occurs error: %s", err)
	}
	onlineUserKeys, xerr := s.wc.Keys(conf.PrefixOnlineUser + "*")
	if xerr == nil {
		res.OnlineUserCount = len(onlineUserKeys)
		if res.HistoryMaxOnline, err = s.wc.PutHistoryMaxOnline(res.OnlineUserCount); err != nil {
			logrus.Errorf("get Siteinfo[3] occurs error: %s", err)
		}
	} else {
		logrus.Errorf("get Siteinfo[2] occurs error: %s", err)
	}
	// 错误进行宽松赦免处理
	return res, nil
}

func (s *adminSrv) GetSiteSettings() (*web.SiteSettingsResp, error) {
	profile, err := s.settings.GetProfile(context.Background())
	if err != nil {
		return nil, err
	}
	return &web.SiteSettingsResp{
		SiteProfileResp: *profile,
		ReadonlyFields:  settings.CloneReadonlyFields(),
	}, nil
}

func (s *adminSrv) UpdateSiteSettings(req *web.SiteSettingsReq) (*web.SiteSettingsResp, error) {
	profile, err := s.settings.UpdateEditableProfile(context.Background(), settings.EditableFromRequest(req))
	if err != nil {
		return nil, err
	}
	return &web.SiteSettingsResp{
		SiteProfileResp: *profile,
		ReadonlyFields:  settings.CloneReadonlyFields(),
	}, nil
}

func (s *adminSrv) GetSettingsSchema() (*web.AdminSettingsSchemaResp, error) {
	return s.settings.GetSchema()
}

func (s *adminSrv) GetSettingsValues() (*web.AdminSettingsValuesResp, error) {
	return s.settings.GetValues(context.Background())
}

func (s *adminSrv) SaveSettings(req *web.AdminSettingsSaveReq) (*web.AdminSettingsSaveResp, error) {
	return s.settings.SaveValues(context.Background(), req.Items)
}

func newAdminSrv(s *base.DaoServant, wc core.WebCache, settings *settings.Service) api.Admin {
	return &adminSrv{
		DaoServant:   s,
		wc:           wc,
		settings:     settings,
		serverUpTime: time.Now().Unix(),
	}
}
