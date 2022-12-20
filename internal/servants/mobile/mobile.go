package mobile

import (
	api "github.com/rocboss/paopao-ce/auto/rpc/v1"
	"google.golang.org/grpc"
)

func RegisterServants(s *grpc.Server) {
	api.RegisterAuthenticateServer(s, newAuthenticateServer())
}
