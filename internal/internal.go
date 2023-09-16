// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/rocboss/paopao-ce/internal/metrics"
	"github.com/rocboss/paopao-ce/internal/migration"
)

func Initial() {
	// migrate database if needed
	migration.Run()
	// event manager system initialize
	events.Initial()
	// metric manager system initialize
	metrics.Initial()
}
