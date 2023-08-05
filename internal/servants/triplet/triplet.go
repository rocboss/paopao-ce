// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package triplet

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	api "github.com/rocboss/paopao-ce/auto/connect/core/v1/corev1connect"
)

func RegisterAuthenticateHandler(h func(string, http.Handler), opts ...connect.HandlerOption) {
	h(api.NewAuthenticateServiceHandler(&authenticateSrv{}, opts...))
}
