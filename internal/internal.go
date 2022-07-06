package internal

import (
	"github.com/rocboss/paopao-ce/internal/migration"
	"github.com/rocboss/paopao-ce/internal/routers/api"
	"github.com/rocboss/paopao-ce/internal/service"
)

func Initialize() {
	// migrate database if needed
	migration.Run()

	// initialize service
	service.Initialize()
	api.Initialize()
}
