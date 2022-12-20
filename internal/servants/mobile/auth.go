package mobile

import (
	api "github.com/rocboss/paopao-ce/auto/rpc/v1"
)

var (
	_ api.AuthenticateServer = (*authenticateServant)(nil)
)

type authenticateServant struct {
	api.UnimplementedAuthenticateServer
}

func newAuthenticateServer() *authenticateServant {
	return &authenticateServant{}
}
