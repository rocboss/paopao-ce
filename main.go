// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/rocboss/paopao-ce/cmd"
	_ "github.com/rocboss/paopao-ce/cmd/migrate"
	_ "github.com/rocboss/paopao-ce/cmd/serve"
)

func main() {
	cmd.Execute()
}
