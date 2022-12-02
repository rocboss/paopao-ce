// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build migration
// +build migration

package migration

import (
	"embed"
)

//go:embed **/*
var Files embed.FS
