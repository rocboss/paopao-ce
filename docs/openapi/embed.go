// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build docs
// +build docs

package openapi

import (
	"embed"
	"net/http"
)

//go:embed index.html openapi.json **/*
var files embed.FS

// NewFileSystem get an embed static assets http.FileSystem instance.
func NewFileSystem() http.FileSystem {
	return http.FS(files)
}
