//go:build migration
// +build migration

package migration

import (
	"embed"
)

//go:embed **/*
var Files embed.FS
