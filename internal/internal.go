package internal

import (
	"github.com/rocboss/paopao-ce/internal/routers/api"
	"github.com/rocboss/paopao-ce/internal/service"
)

func Initialize() {
	service.Initialize()
	api.Initialize()
}
