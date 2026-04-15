// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/infra"
)

func Initial(ctx context.Context) {
	// infrastruct component initialize
	infra.Initial(ctx)
}
