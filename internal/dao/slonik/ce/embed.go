// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build migration
// +build migration

package main

import (
	"embed"
)

//go:embed postgres/schema
var Files embed.FS
