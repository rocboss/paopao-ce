//go:build constraint

package admin

import (
	api "github.com/rocboss/paopao-ce/auto/api/m/v1"
)

var _ api.User = (*userSrv)(nil)
