// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package space

import (
	api "github.com/rocboss/paopao-ce/auto/api/x/v1"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

var _ api.User = (*userSrv)(nil)

type userSrv struct {
	*base.BaseServant
	api.UnimplementedUserServant
}

func newUserSrv() api.User {
	return &userSrv{
		BaseServant: base.NewBaseServant(),
	}
}
