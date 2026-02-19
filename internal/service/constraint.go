//go:build constraint

package service

var (
	_ server = (*grpcServer)(nil)
	_ server = (*httpServer)(nil)

	_ Service = (*adminService)(nil)
	_ Service = (*botService)(nil)
	_ Service = (*frontendWebService)(nil)
	_ Service = (*localossService)(nil)
	_ Service = (*metricsService)(nil)
	_ Service = (*pprofService)(nil)
	_ Service = (*mobileService)(nil)
	_ Service = (*spaceXService)(nil)
	_ Service = (*webService)(nil)
)
