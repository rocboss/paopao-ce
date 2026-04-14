//go:build constraint

package space

import (
	api "github.com/rocboss/paopao-ce/auto/api/x/v1"
)

var _ api.User = (*userSrv)(nil)
