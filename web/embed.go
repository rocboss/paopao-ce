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
)

//go:embed dist/*
var files embed.FS

// NewFileSystem get an embed static assets http.FileSystem instance.
func NewFileSystem() http.FileSystem {
	subfs, _ := fs.Sub(files, "dist")
	return http.FS(subfs)
}
