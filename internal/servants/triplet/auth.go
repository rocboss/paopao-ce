// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package triplet

import (
	api "github.com/rocboss/paopao-ce/auto/connect/core/v1/corev1connect"
)

var (
	_ api.AuthenticateServiceHandler = (*authenticateSrv)(nil)
)

type authenticateSrv struct {
	api.UnimplementedAuthenticateServiceHandler
}
