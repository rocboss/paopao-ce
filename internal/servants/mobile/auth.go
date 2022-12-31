// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
