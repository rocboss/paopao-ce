// Copyright 2025 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package infra infrastructure components.
// 该包可以包含的功能模块：
// - 一些app的基础功能相关的模块；
// - 一些单一、较大的功能模块；
// - 在想放到internal包中但是又不知道放那的功能包，那就放到这里，尽量保持internal包的一级子包一眼就可以看出app内部功能的主要结构；

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
