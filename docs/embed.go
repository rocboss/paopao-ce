//go:build docs
// +build docs

package docs

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
