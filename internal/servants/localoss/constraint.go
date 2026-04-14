//go:build constraint

package localoss

import (
	api "github.com/rocboss/paopao-ce/auto/api/s/v1"
)

var _ api.User = (*userSrv)(nil)
