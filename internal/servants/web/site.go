// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/pkg/version"
)

var _ api.Site = (*siteSrv)(nil)

type siteSrv struct {
	api.UnimplementedSiteServant
	*base.BaseServant
}

func (*siteSrv) Profile() (*web.SiteProfileResp, error) {
	return conf.WebProfileSetting, nil
}

func (*siteSrv) Version() (*web.VersionResp, error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

func newSiteSrv() api.Site {
	return &siteSrv{
		BaseServant: base.NewBaseServant(),
	}
}
