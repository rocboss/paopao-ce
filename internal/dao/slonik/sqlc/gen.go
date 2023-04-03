// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"os"

	sqlc "github.com/kyleconroy/sqlc/pkg/cli"
)

//go:generate go run $GOFILE generate -x
func main() {
	os.Exit(sqlc.Run(os.Args[1:]))
}
