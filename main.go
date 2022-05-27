package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/routers"
)

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           global.ServerSetting.HttpIp + ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("\x1b[36m%s\x1b[0m\n", "PaoPao service listen on http://"+global.ServerSetting.HttpIp+":"+global.ServerSetting.HttpPort)
	s.ListenAndServe()
}
