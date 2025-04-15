// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !(slim && embed)
// +build !slim !embed

package web

import (
	"embed"
	"io/fs"
	"net/http"

	embedfs "github.com/alimy/tryst/embed"
	"github.com/rocboss/paopao-ce/pkg/version"
)

//go:embed all:dist
var files embed.FS

// NewFileSystem get an embed static assets http.FileSystem instance.
func NewFileSystem() http.FileSystem {
	subfs, _ := fs.Sub(files, "dist")
	// add custom mod time for embed fs
	timefs := embedfs.NewFS(subfs, version.BuildTime)
	return http.FS(timefs)
}
