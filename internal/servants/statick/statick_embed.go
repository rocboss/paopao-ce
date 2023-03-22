// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build embed
// +build embed

package statick

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/web"
)

// RegisterWebStatick register web static assets route
func RegisterWebStatick(e *gin.Engine) {
	routeWebStatic(e, "/", "/index.html", "/favicon.ico", "/logo.png", "/sw.js", "/manifest.json", "/assets/*filepath")
}

func routeWebStatic(e *gin.Engine, paths ...string) {
	staticHandler := http.FileServer(web.NewFileSystem())
	handler := func(c *gin.Context) {
		staticHandler.ServeHTTP(c.Writer, c.Request)
	}
	for _, path := range paths {
		e.GET(path, handler)
		e.HEAD(path, handler)
	}
}
