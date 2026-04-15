// Copyright 2025 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package infra infrastructure components.
// 该包可以包含的功能模块：
// - 一些app的基础功能相关的模块；
// - 一些单一、较大的功能模块；
// - 尽量保持internal包的清晰结构关系，可以将其他不知道放哪的包先放到infra中，后续再视情况调整；

package infra

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/infra/events"
	"github.com/rocboss/paopao-ce/internal/infra/metrics"
	"github.com/rocboss/paopao-ce/internal/infra/migration"
	"github.com/rocboss/paopao-ce/internal/infra/settings"
)

func Initial(ctx context.Context) {
	// migrate database if needed
	migration.Run()
	// event manager system initialize
	events.Initial()
	// metric manager system initialize
	metrics.Initial()
	// settings manager system bootstrap
	settings.Bootstrap(ctx, conf.MustGormDB())
}
