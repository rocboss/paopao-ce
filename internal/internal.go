// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/infra"
)

// Initial 内部功能模块初始化
// Note: 目前函数很简单，只有infra需要初始化，暂时保留此函数，后续可能会
// 有其他大的服务模块放到internal包中，需要统一在此函数中做初始化动作，而
// 且对外保留一个internal的统一初始化接口符合直觉。
func Initial(ctx context.Context) {
	// infrastruct component initialize
	infra.Initial(ctx)
}
