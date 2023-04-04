// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/rocboss/paopao-ce/internal/migration"
)

func Initial() {
	// migrate database if needed
	migration.Run()
}
