//go:build constraint

package mobile

import (
	api "github.com/rocboss/paopao-ce/auto/rpc/greet/v1"
)

var _ api.GreetServiceServer = (*greetServiceSrv)(nil)
