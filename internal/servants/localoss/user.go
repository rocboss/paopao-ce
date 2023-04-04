// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package localoss

import (
	api "github.com/rocboss/paopao-ce/auto/api/s/v1"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

var (
	_ api.User        = (*userSrv)(nil)
	_ api.UserBinding = (*userBinding)(nil)
	_ api.UserRender  = (*userRender)(nil)
)

type userSrv struct {
	base.BaseServant
	api.UnimplementedUserServant
}

type userBinding struct {
	base.BaseBinding
	*api.UnimplementedUserBinding
}

type userRender struct {
	base.BaseRender
	*api.UnimplementedUserRender
}

func newUserSrv() api.User {
	return &userSrv{}
}

func newUserBinding() api.UserBinding {
	return &userBinding{
		UnimplementedUserBinding: &api.UnimplementedUserBinding{
			BindAny: base.NewBindAnyFn(),
		},
	}
}

func newUserRender() api.UserRender {
	return &userRender{
		UnimplementedUserRender: &api.UnimplementedUserRender{
			RenderAny: base.RenderAny,
		},
	}
}
