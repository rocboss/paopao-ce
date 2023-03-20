// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// nolint
var (
	_errUserHasBeenBanned  = xerror.NewError(20006, "该账户已被封停")
	_errAccountNoPhoneBind = xerror.NewError(20013, "拒绝操作: 账户未绑定手机号")
	_errNoAdminPermission  = xerror.NewError(20022, "无管理权限")
)
