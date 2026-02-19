//go:build constraint

package bot

import (
	api "github.com/rocboss/paopao-ce/auto/api/r/v1"
)

var _ api.User = (*userSrv)(nil)
