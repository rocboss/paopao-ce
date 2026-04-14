// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"context"

	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/sitesetting"
	"github.com/rocboss/paopao-ce/pkg/version"
)

type siteSrv struct {
	api.UnimplementedSiteServant
	*base.BaseServant
	settings *sitesetting.Service
}

func (s *siteSrv) Profile() (*web.SiteProfileResp, error) {
	return s.settings.GetProfile(context.Background())
}

func (*siteSrv) Version() (*web.VersionResp, error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

func newSiteSrv(settings *sitesetting.Service) api.Site {
	return &siteSrv{
		BaseServant: base.NewBaseServant(),
		settings:    settings,
	}
}
