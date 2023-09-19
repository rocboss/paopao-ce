// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mobile

import (
	api "github.com/rocboss/paopao-ce/auto/rpc/greet/v1"
	"google.golang.org/grpc"
)

func RegisterServants(s *grpc.Server) {
	api.RegisterGreetServiceServer(s, newGreetServiceServer())
}
